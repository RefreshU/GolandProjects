package geerpc

import (
	"awesomeProject/codec/codec"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

// Call represents an active RPC
type Call struct {
	Seq uint64
	ServiceMethod string // format "<service>.<method>"
	Args interface{}     // arguments to the functions
	Reply interface{}    //reply from the function
	Error error          // if error occurs, it will be set
	Done chan *Call      // Strobes when is complete
}

func (call *Call) done(){
	call.Done <-call
}

// Client represents an RPC Client
// THere may be multiple outstanding Calls associated
// with a single Client, and a Client may be used by
// multiple goroutines simultaneously.
//cc 是消息的编解码器，和服务端类似，用来序列化将要发送出去的请求，以及反序列化接收到的响应。
//sending 是一个互斥锁，和服务端类似，为了保证请求的有序发送，即防止出现多个请求报文混淆。
//header 是每个请求的消息头，header 只有在请求发送时才需要，而请求发送是互斥的，因此每个客户端只需要一个，声明在 Client 结构体中可以复用。
//seq 用于给发送的请求编号，每个请求拥有唯一编号。
//pending 存储未处理完的请求，键是编号，值是 Call 实例。
//closing 和 shutdown 任意一个值置为 true，则表示 Client 处于不可用的状态，但有些许的差别，closing 是用户主动关闭的，即调用 Close 方法，而 shutdown 置为 true 一般是有错误发生。
type Client struct {
	cc codec.Codec
	opt *Option
	sending sync.Mutex //protect following
	header codec.Header
	mu sync.Mutex
	seq uint64
	pending map[uint64]*Call
	closing bool // user has called Close
	shutdown bool // server has told us to stop
}

var _ io.Closer = (*Client)(nil)

var ErrShutdown = errors.New("connection is shutdown")

// Close the function
func (client *Client) Close()error {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing {
		return ErrShutdown
	}
	client.closing = true
	return client.cc.Close()
}

// IsAvailable return true if the client does work
func (client *Client) IsAvailable() bool {
	client.mu.Lock()
	defer client.mu.Unlock()
	return !client.shutdown && !client.closing
}

func (client *Client) registerCall(call *Call)(uint64 , error){
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing || client.shutdown {
		return 0, ErrShutdown
	}
	call.Seq = client.seq
	client.pending[call.Seq] = call
	client.seq++
	return call.Seq, nil
}

func (client *Client) removingCall(seq uint64) *Call {
	client.mu.Lock()
	defer client.mu.Unlock()
	call := client.pending[seq]
	delete(client.pending, seq)
	return call
}

func (client *Client) terminateCall(err error){
	client.sending.Lock()
	defer client.sending.Unlock()
	client.mu.Lock()
	defer client.mu.Unlock()
	client.shutdown = true
	for _, call := range client.pending {
		call.Error = err
		call.done()
	}
}

func (client *Client) receive(){
	var err error
	for err == nil {
		var h codec.Header
		if err = client.cc.ReadHeader(&h); err != nil {
			break
		}
		call := client.removingCall(h.Seq)
		switch {
		case call == nil :
			// it usually means that Write partially failed
			// and call was already removed
			err = client.cc.ReadBody(nil)
		case h.Error != "":
			call.Error = fmt.Errorf(h.Error)
			err = client.cc.ReadBody(nil)
			call.done()
		default:
			err = client.cc.ReadBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
		// error occurs, so terminateCalls pending calls
		client.terminateCall(err)
	}
}

func NewClient(conn net.Conn, opt *Option)(*Client, error){
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		err := fmt.Errorf("invalid codec type %s", opt.CodecType)
		log.Println("rpc client: codec error:", err)
		return nil, err
	}
	//send options with server
	if err := json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client: option error:", err)
		_ = conn.Close()
		return nil, err
	}
	return newClientCodec(f(conn), opt), nil
}

func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq: 1, //seq starts with 1, 0 means invalid call
		cc: cc,
		opt: opt,
		pending: make(map[uint64]*Call),
	}
	go client.receive()
	return client
}

func parseOptions(opts ...*Option)(*Option, error){
	//if opts is nil or pass nil as parameter
	if len(opts) == 0 || opts[0] == nil{
		return DefaultOption, nil
	}

	if len(opts) != 1 {
		return nil, errors.New("number of options more than 1")
	}

	opt := opts[0]
	opt.MagicNumber = DefaultOption.MagicNumber
	if opt.CodecType == ""{
		opt.CodecType = DefaultOption.CodecType
	}
	return opt, nil
}

// Dail connects to to-an RPC server at the specified network address
func Dail(network, address string, opts ...*Option)(client *Client, err error){
	opt, err := parseOptions(opts...)
	if err != nil {
		return nil, err
	}
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	// close the connections if client is nil
	defer func() {
		if client == nil {
			_ = conn.Close()
		}
	}()
	return NewClient(conn, opt)
}

func (client *Client)send(call *Call){
	//make sure that the client will send a complete request
	client.sending.Lock()
	defer client.sending.Unlock()

	//register this call
	seq, err := client.registerCall(call)
	if err != nil {
		call.Error = err
		call.done()
		return
	}

	// prepare request header
	client.header.ServiceMethod = call.ServiceMethod
	client.header.Seq = seq
	client.header.Error = ""

	// encode and send the request
	if err := client.cc.Write(&client.header, call.Args); err != nil {
		call := client.removingCall(seq)
		// call may be nil, it usually means that Write partially failed.
		// client has received the response and handled
		if call != nil {
			call.Error = err
			call.done()
		}
	}

}
// GO invokes the functions asynchronously
// It returns the call structure representing the invocation
func (client *Client) Go(serviceMethod string, args, reply interface{}, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0{
		log.Panic("rpc client: done channel is unbuffered")
	}
	call := &Call{
		ServiceMethod: serviceMethod,
		Args: args,
		Reply: reply,
		Done: done,
	}
	client.send(call)
	return call
}

// Call invokes the named function, waits for it to complete
// and returns its error status
func (client *Client) Call(serviceMethod string, args, reply interface{}) error {
	call := <-client.Go(serviceMethod, args, reply, make(chan *Call, 1)).Done
	return call.Error
}
