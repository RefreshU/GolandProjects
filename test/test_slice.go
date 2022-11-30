package main

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"math"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

/*func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path,'/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB
	dir1 = append(dir1,"suffix"...)
	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB
}*/
//type Country struct {
//	Name string
//}
//
//type City struct {
//	Name string
//}
//
//type Stringable interface {
//	ToString() string
//}
//
//func (c Country) ToString() string {
//	return "Country = " + c.Name
//}
//
//func (c City) ToString() string {
//	return "City = " + c.Name
//}
//func PrintStr(s Stringable) {
//	fmt.Println(s.ToString())
//}
//
//func main () {
//	d1 := Country {"USA"}
//	d2 := City {"Los Angeles"}
//
//	PrintStr(d1)
//	PrintStr(d2)
//}
//func main(){
//	type favContextKey	string
//	f := func(ctx context.Context, k favContextKey){
//		if v := ctx.Value(k); v != nil {
//			fmt.Println("found value:", v)
//			return
//		}
//		fmt.Println("key not found:", k)
//	}
//
//	k := favContextKey("language")
//	ctx := context.WithValue(context.Background(), k, "Go")
//
//	f(ctx, k)
//	f(ctx, favContextKey("color"))
//}
/**
	等待组
 */
//var wg sync.WaitGroup
//func run(i int){
//	fmt.Println("start 任务ID：", i)
//	time.Sleep(time.Second * 3)
//	wg.Done() //每个goroutine运行结束后就释放等待计数器
//}
//
//func main(){
//	countThread := 2 //runtime.NumCPU()
//	for i := 0; i < countThread; i++ {
//		go run(i) //LIFO
//	}
//	wg.Add(countThread) //需要开启的goroutine等待组的计数器
//	fmt.Println("等待中......")
//	wg.Wait()
//	fmt.Println("全部任务退出")
//}
/**
  *  通道 + select
 */
//func run(stop chan bool){
//	for{
//		select {
//			case <-stop:
//				fmt.Println("task 1 end")
//				return
//			default:
//				fmt.Println("task 1 running")
//				time.Sleep(time.Second * 1)
//		}
//	}
//}
//func main(){
//	stop := make(chan bool)
//	go run(stop)
//	time.Sleep(time.Second * 10)
//	fmt.Println("stopping task 1...")
//	stop <- true
//	time.Sleep(time.Second * 3)
//	return
//}
/**
 * common context
 */
//func run(ctx context.Context, id int){
//	for{
//		select{
//			case <-ctx.Done():
//				fmt.Printf("task %v end \n", id)
//				return
//			default:
//				fmt.Printf("task %v running .... \n", id)
//				time.Sleep(time.Second * 2)
//		}
//	}
//}
//
//func main(){
//	//init context
//	ctx, cancel := context.WithCancel(context.Background())
//	//start multi goroutine
//	go run(ctx, 1)
//	go run(ctx, 2)
//	time.Sleep(time.Second * 10)
//	fmt.Println("stop task......")
//	//using cancel() stop goroutine
//	cancel()
//	time.Sleep(time.Second * 1)
//	return
//}
/*
 * Context Timeout
 */
/*func coroutine(ctx context.Context, duration time.Duration, id int, wg *sync.WaitGroup){
	for{
		select {
		case <-ctx.Done():
			fmt.Printf("coroutine %d out\n", id)
			wg.Done()
			return
		case <-time.After(duration):
			fmt.Printf("message from coroutine %d \n", id)
			
		}
	}
}
func main(){
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go coroutine(ctx, 3*time.Second, i, wg)
	}
	wg.Wait()
}*/
/**
 * Context Pass Metadata
 */
//var key string = "name"
//func run(ctx context.Context){
//	for{
//		select {
//		case <-ctx.Done():
//			fmt.Printf("task %v ending \n", ctx.Value(key))
//			return
//		default:
//			fmt.Printf("task %v is running \n", ctx.Value(key))
//			time.Sleep(time.Second * 2)
//		}
//	}
//}
//
//func main(){
//	ctx, cancel := context.WithCancel(context.Background())
//
//	valuectx := context.WithValue(ctx, key, "【monitor】")
//
//	go run(valuectx)
//
//	time.Sleep(time.Second * 10)
//	fmt.Println("Stopping Task")
//
//	//using the cancel() cancel context to stopping goroutine
//	cancel()
//	//
//	time.Sleep(time.Second * 3)
//}
func Json_encode(data interface{}) (string, error){
	jsons, err := json.Marshal(data)
	return string(jsons), err
}
func ArrayMerge(data ...[]interface{}) (interface{}){
	n := 0
	for _, v := range data {
		n += len(v)
	}
	newArray := make([]interface{}, 0, n)
	for _, v := range data {
		newArray = append(newArray, v...)
	}
	return newArray
}
//func main(){
//	var person = map[string]string{
//		"name" :    "Anmy",
//		"sex":     "man",
//		"company": "Apple",
//		"school":  "Tshing Hua",
//	}
//	//var Job = map[int]string{
//	//	0: "programer",
//	//	1: "Assist",
//	//}
//	//var other = map[int]int {
//	//	0: 1000,
//	//	2: 1200,
//	//}
//	fmt.Println(person["name"])
//	arrListJson, err := Json_encode(person)
//	if err != nil {
//		fmt.Println("failed to map to json")
//	}
//	fmt.Printf("Json string: %s", arrListJson)
//}

//type Person struct {
//	Name string
//	Age int
//}
//
//func MapToStruct() {
//	mapInstance := make(map[string]interface{})
//	mapInstance["Name"] = "liang637210"
//	mapInstance["Age"] = 28
//
//	var person Person
//	//将 map 转换为指定的结构体
//	if err := mapstructure.Decode(mapInstance, &person); err != nil {
//		fmt.Println(err)
//	}
//	fmt.Printf("map2struct后得到的 struct 内容为:%v", person)
//}
//
//func main(){
//	MapToStruct()
//}

/**
	类型的定义和类型别名的区别
 */
//类型定义
//type NewInt int
//
////类型别名
////type MyInt = int
////func main() {
////	var a NewInt
////	var b MyInt
////
////	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
////	fmt.Printf("type of b:%T\n", b) //type of b:int
//}

/**
	匿名结构体
 */
//func main(){
//	var User struct{Name string; Age int;Sex string}
//	User.Name = "Little Prince"
//	User.Age = 20
//	fmt.Printf("%#v\n", User)
//}
// 使用值接收者

//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

//func main() {
//	p1 := NewPerson("小王子", 25)
//	p1.Dream()
//	fmt.Println(p1.age) // 25
//	p1.SetAge2(30) // (*p1).SetAge2(30)
//	fmt.Println(p1.age) // 25
//}
/**
	创建指针类型结构体
 */
//type person struct {
//	name string
//	city string
//	age  int8
//}
//func main(){
//	//通过new关键字对结构体进行实例化，得到的是结构的地址
//	var p2 = new(person)
//	p2.name = "Little Prince"
//	p2.city = "SZ"
//	p2.age = 40
//	fmt.Printf("p2=%#v\n", p2)
//	//取结构体的地址实例化
//	//使用&对结构体取地址操作相当于对该结构体进行了一次new实例化操作
//
//	p3 := &person{}
//	p3.name = "seven"
//	p3.city = "SZ"
//	fmt.Printf("p3=%#v\n", p3)
//}
/*
	go编程 修饰器
 */
//func decorator(f func(s string)) func(s string){
//	return func(s string){
//		fmt.Println("Startd")
//		f(s)
//		fmt.Println("Done")
//	}
//}
//
//func Hello(s string){
//	fmt.Println(s)
//}
//
//func main() {
//	//1. 直接调用
//	decorator(Hello)("Hello, world!")
//
//	//2. 先修饰函数，然后再调用
//	hello := decorator(Hello)
//	hello("Hello, World")
//}
/*
	示例代码 1
 */
type SumFunc func(int64, int64) int64

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timedSumFunc(f SumFunc) SumFunc {
	return func(start, end int64) int64 {
		defer func(t time.Time){
			fmt.Printf("--- Time Elapsed (%s): %v ---\n", getFunctionName(f), time.Since(t))
		}(time.Now())
		return f(start,end)
	}
}
func Sum1(start, end int64) int64{
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i < end; i++ {
		sum += i
	}
	return sum
}

func Sum2(start, end int64) int64{
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}

//func main(){
//	sum1 := timedSumFunc(Sum1)
//	sum2 := timedSumFunc(Sum2)
//
//	fmt.Printf("%d, %d\n", sum1(-10000, 1000000), sum2(-10000, 100000))
//}
/**
	示例2 HTTP Server
 */
//func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		log.Println("--->WithServerHeader()")
//		w.Header().Set("Server", "HelloServer v0.0.1")
//		h(w, r)
//	}
//}
//func WithAuthCookie(h http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request){
//		log.Println("--->WithServerHeader()")
//		cookie := &http.Cookie{Name: "Auth", Value: "Pass", Path: "/"}
//		http.SetCookie(w, cookie)
//		h(w, r)
//	}
//}
//func WithBasicAuth(h http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		log.Println("--->WithBasicAuth")
//		cookie, err := r.Cookie("Auth")
//		if err != nil || cookie.Value != "Pass" {
//			w.WriteHeader(http.StatusForbidden)
//			return
//		}
//		h(w, r)
//	}
//}
//
//func WithDebugLog(h http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		log.Println("--->WithDebugLog")
//		r.ParseForm()
//		log.Println(r.Form)
//		log.Println("path", r.URL.Path)
//		log.Println("scheme", r.URL.Scheme)
//		log.Println(r.Form["url_long"])
//		for k, v := range r.Form {
//			log.Println("key:", k)
//			log.Println("val:", strings.Join(v, ""))
//		}
//		h(w, r)
//	}
//}
//func hello(w http.ResponseWriter, r *http.Request){
//	log.Printf("Recived Request %s from %s.\n", r.URL.Path, r.RemoteAddr)
//	fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
//}
//
//type HttpHandlerDecorator func(http.HandlerFunc) http.HandlerFunc
//
//func Handler(h http.HandlerFunc, decors ...HttpHandlerDecorator) http.HandlerFunc {
//	for i := range decors {
//		d := decors[len(decors) - 1 - i]
//		h = d(h)
//	}
//	return h
//}
//func main(){
//	//http.HandleFunc("/v1/hello", WithServerHeader(hello))
//	//http.HandleFunc("/v1/hello", WithServerHeader(WithAuthCookie(hello)))
//	//http.HandleFunc("/v2/hello", WithServerHeader(WithBasicAuth(hello)))
//	//http.HandleFunc("/v3/hello", WithServerHeader(WithBasicAuth(WithDebugLog(hello))))
//	http.HandleFunc("/v4/hello", Handler(hello,
//		WithServerHeader, WithBasicAuth, WithDebugLog))
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		log.Fatal("ListenAndServer: ", err)
//	}
//}
/*
	Visitor 模式
 */
type Info struct {
	Namespace string
	Name string
	OtherThing string
}
//定义一个VisitorFunc的函数类型
type VisitorFunc func(*Info, error) error

// define Visitor Interface
type Visitor interface {
	Visit(VisitorFunc) error
}
//为Info 实现 Visitor 接口中的 Visit() 方法
func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}
/*
	Name Visitor
 */
/**
 声明了一个 NameVisitor 的结构体，这个结构体里有一个 Visitor 接口成员，这里意味着多态；
 */
type NameVisitor struct {
	visitor Visitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error  {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("NameVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("===> Name=%s, Namespace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function")
		return err
	})
}
/*
	Other Visitor
 */
//声明了一个 OtherVisitor 的结构体，这个结构体里有一个 Visitor 接口成员，这里意味着多态；
type OtherVisitor struct {
	visitor Visitor
}

func (v OtherVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("OtherVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("===> OtherThing=%s\n", info.OtherThing)
		}
		fmt.Println("OtherVisitor() after call function")
		return err
	})
}
/*
	Log Visitor
 */
type LogVisitor struct {
	visitor Visitor
}
func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}

//func main(){
//	info := Info{}
//	var v Visitor = &info
//	v = LogVisitor{v}
//	v = NameVisitor{v}
//	v = OtherVisitor{v}
//
//	loadFile := func(info *Info, err error) error {
//		info.Name = "Name"
//		info.Namespace = "Namespace"
//		info.OtherThing = "we are running a visitor"
//		return err
//	}
//	v.Visit(loadFile)
//}

//func main(){
//	// 示例1。
//	contents := "ab"
//	buffer1 := bytes.NewBufferString(contents)
//
//	unreadBytes := buffer1.Bytes()
//	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
//	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
//		contents, buffer1.Cap())
//	fmt.Println()
//
//	contents = "cdefg"
//	fmt.Printf("Write contents %q ...\n", contents)
//	buffer1.WriteString(contents)
//	fmt.Printf("The capacity of buffer:%d\n", buffer1.Cap())
//	fmt.Println()
//
//	// 只要扩充一下之前拿到的未读字节切片unreadBytes，
//	// 就可以用它来读取甚至修改buffer中的后续内容。
//	unreadBytes = unreadBytes[:cap(unreadBytes)]
//	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
//	fmt.Println()
//
//	value := byte('X')
//	fmt.Printf("Set a byte in the unread bytes to %v ...\n", value)
//	unreadBytes[len(unreadBytes)-2] = value
//	fmt.Printf("The unread off point %d", buffer1.Bytes())
//	fmt.Printf("The unread bytes of the buffer: %v\n", buffer1.Bytes())
//	fmt.Println()
//
//	// 不过，在buffer的内容容器真正扩容之后就无法这么做了。
//	contents = "hijklmn"
//	fmt.Printf("Write contents %q ...\n", contents)
//	buffer1.WriteString(contents)
//	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
//	fmt.Println()
//
//	unreadBytes = unreadBytes[:cap(unreadBytes)]
//	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
//	fmt.Print("\n\n")
//}
type Catalogue struct {
	CatalogueName string `json:"catalogue_name"`
	CatalogueParentId int64 `json:"catalogue_parent_id"`
	CatalogueId int64 `json:"catalogue_id"`
	CatalogueLevel int64 `json:"catalogue_level"`
}
type CatalogueTree struct {
	CatalogueName string `json:"catalogue_name"`
	CatalogueParentId int64 `json:"catalogue_parent_id"`
	CatalogueId int64 `json:"catalogue_id"`
	CatalogueLevel int64 `json:"catalogue_level"`
	Child []*CatalogueTree `json:"child"`
}
func insert_catalogue_tree(catalogues []Catalogue, Catalogue_parent_id int64, tree *CatalogueTree) error {
	catalogue_childs := getCatalogueChild(catalogues, Catalogue_parent_id)
	for _, catalogue := range catalogue_childs {
		child := CatalogueTree{catalogue.CatalogueName, catalogue.CatalogueParentId, catalogue.CatalogueId, catalogue.CatalogueLevel,[]*CatalogueTree{}}
		tree.Child = append(tree.Child, &child)
		_ = insert_catalogue_tree(catalogues, catalogue.CatalogueId, &child)
	}
	return nil
}
func getCatalogueChild(catalogues []Catalogue, catalogue_parent_id int64) ([]Catalogue){
	var cata_arr []Catalogue
	for _, catalogue := range catalogues {
		if catalogue.CatalogueParentId == catalogue_parent_id {
			cata_arr = append(cata_arr, catalogue)
		}
	}
	return cata_arr
}
//func main(){
//	catalogues := []Catalogue{
//		{CatalogueName: "catalogue1", CatalogueParentId: 0, CatalogueId: 1, CatalogueLevel: 1},
//		{CatalogueName: "catalogue1.1", CatalogueParentId: 1, CatalogueId: 2, CatalogueLevel: 2},
//		{CatalogueName: "catalogue1.2", CatalogueParentId: 1, CatalogueId: 3, CatalogueLevel: 2},
//		{CatalogueName: "catalogue2", CatalogueParentId: 0, CatalogueId: 4, CatalogueLevel: 1},
//		{CatalogueName: "catalogue2.1", CatalogueParentId: 4, CatalogueId: 5, CatalogueLevel: 2},
//		{CatalogueName: "catalogue2.2", CatalogueParentId: 4, CatalogueId: 6, CatalogueLevel: 2},
//	}
//	var cata_tree = CatalogueTree{"catalogue0", 0, 0,0, []*CatalogueTree{}}
//
//
//	_ = insert_catalogue_tree(catalogues, 0, &cata_tree)
//	cata_tree_json, err := json.Marshal(cata_tree)
//	if err != nil {
//		fmt.Println("json marshal failed")
//	}
//	fmt.Printf("%s", cata_tree_json)
//	//for _, cata := range cata_tree.Child {
//	//	fmt.Println(cata.CatalogueName)
//	//	//if len(cata.Child) > 0 {
//	//	//
//	//	//}
//	//}
//}
func B2S(bs []uint8) string {
	// 将字节切片转换成字符串
	ba := make([]byte, 0)  // 创建一个字节切片，默认容量0
	fmt.Println(ba, reflect.TypeOf(ba))  // 字节切片: [] 每一个元素是：uint8 类型: []uint8
	for i, v := range bs {
		fmt.Println(i, v)
		ba = append(ba, v)
	}
	return string(ba)
}

func Char2S(cs []int32) string {
	return string(cs)  // 将字符切片直接转换成字符串
}

//func main() {
	// 1. 字符串和字节切片的互转
	//bs := []byte("马亚南")
	//fmt.Println(bs)
	//s1 := B2S(bs)
	//fmt.Println(s1)

	//2. 字符切片和字符串的互转
	//charSlice := make([]int32, 0)
	//a := "马亚南"
	//fmt.Println(len(a))
	//for i, v := range a {  // 注意：range 也是赋值拷贝
	//	fmt.Println(i, v)
	//	v = int32(20122)  // a不会改变，还是马亚南，因为range是赋值拷贝
	//	charSlice = append(charSlice, v)  // 由于charSlice是新构建的字符切片，所以它都是一个值 20122
	//}
	//fmt.Println(a)
	//s2 := Char2S(charSlice)
	//fmt.Println(s2)
	//
	//var ages = []int64{123, 122, 112}
	//var ageStrs []string
	//for _, age := range ages {
	//	ageStrs = append(ageStrs, strconv.FormatInt(age, 10))
	//}
	//fmt.Println(ageStrs)
	//fmt.Printf("%#v\n", ages)
	//fmt.Printf("%#v\n", ageStrs)

//}

//func main() {
//	var uri url.URL
//	q := uri.Query()
//	q.Add("name", "张三")
//	q.Add("age", "20")
//	q.Add("sex", "1")
//	queryStr := q.Encode()
//	fmt.Println(queryStr)
//	api := fmt.Sprintf("api:%s",USER_INFO_URI)
//	fmt.Println(api)
//	//timesteamp := time.Now().Unix()
//	timesteamp := fmt.Sprintf("%d", time.Now().Unix())
//	fmt.Println(timesteamp)
//}


//func handleError(err error) {
//	fmt.Println("Error:", err)
//	os.Exit(-1)
//}
//var secretid string = "AKIDGUjcKsgFE7lGbnh4LYgpLjx4zkmot7zg"
//var secretkey string = "mIYYHkJbomVyCB0sSmqeCKHyjW6HCrG2"
var cos_url string ="http://wechatappdev-10011692.picsh.myqcloud.com"
//func main()  {
//	//u, _ := url.Parse("http://wechatapppro-1252524126.picsh.myqcloud.com")
//	u, _ := url.Parse("cos_url")
//	b := &cos.BaseURL{BucketURL: u}
//	c := cos.NewClient(b, &http.Client{
//		Transport: &cos.AuthorizationTransport{
//			SecretID:  "AKIDA7PXqMW183llaEk4GF5WkkUzJQHxnA0a",
//			SecretKey: "U9RTJzZ4MSM0QgjeKxTmxucy6yb2D9zb",
//		},
//	})
//
//	s,_,err := c.Service.Get(context.Background())
//	if err != nil{
//		panic(err)
//	}
//
//	for _,b := range s.Buckets {
//		fmt.Println("%#v\n",b)
//	}
//
//	//name := "/video/source/app_1.csv"
//	//// 1. 通过字符串上传对象
//	//f := strings.NewReader("test")
//	////
//	//_,err := c.Object.Put(context.Background(),name,f,nil)
//	//if err != nil{
//	//	panic(err)
//	//}
//	//
//	//// 上传本地文件
//	//_, err = c.Object.PutFromFile(context.Background(),name,"./app_1.csv",nil)
//	//if err != nil {
//	//	panic(err)
//	//}
//}
func Split(s, sep string)(result []string){
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
type KnowledgeImportError struct {
	Row int    `json:"row"`
	Msg string `json:"msg"`
}
func ValidateRowDtaa(row []string) (bool, string){
	for column, fieldValue := range row {
		switch column {
		case 0:
			if utf8.RuneCountInString(fieldValue) > 10 {
				return false, fmt.Sprintf("第%d列字段长度超过10", column + 1)
			}
			break
		case 1:
			if utf8.RuneCountInString(fieldValue) <= 0 {
				return false, fmt.Sprintf("第%d列字段值不能为空", column + 1)
			}
			break
		default:
			return true, ""
		}
	}
	return true, ""
}
func handleRow(wg *sync.WaitGroup,importRow chan<- []string, errorRow chan<-  KnowledgeImportError,rows [][]string){
	defer wg.Done()
	for rowNumber,row := range rows{
		if ok, errMsg := ValidateRowDtaa(row); !ok {
			errorRow <- KnowledgeImportError{Row:rowNumber+1,Msg:errMsg}
			continue
		}
		importRow <- row
		fmt.Printf("row data %#v \n", row)
	}
	close(importRow)
	close(errorRow)
}

func importRow(wg *sync.WaitGroup, rows <-chan []string){
	defer wg.Done()
	for row := range rows {
		fmt.Printf("imort data: %#v\n", row)
	}
}

func handleRowErrorMsg(wg *sync.WaitGroup, errors <-chan KnowledgeImportError){
	defer wg.Done()
	for err := range errors {
		fmt.Printf("第%d行：%s\n", err.Row, err.Msg)
	}
}

//func main() {
//	f, err := excelize.OpenFile("./greeting_knowledge_data.xlsx")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer func() {
//		if err := f.Close(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//	rowImportChan := make(chan []string, 20)
//	rowErrorChan := make(chan KnowledgeImportError, 100)
//	// 获取工作表中指定单元格的值
//	cell, err := f.GetCellValue("Sheet1", "B2")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Printf(cell)
//	// 获取 Sheet1 上所有单元格
//	rows, err := f.GetRows("Sheet1")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//for i, row := range rows{
//	//	fmt.Printf("row:%d data:%#v\n", i, row)
//	//}
//	var wg sync.WaitGroup
//	//var rowChunk = make([][]string,0)
//	//totalRows := len(rows)
//	//for rowNumber,row := range rows {
//	//	if rowNumber+1 % 10 != 0 {
//	//		rowChunk = append(rowChunk, row)
//	//	} else if totalRows - 1 == rowNumber || rowNumber+1 % 10 == 0{
//	//		fmt.Println(rowChunk)
//	//		//wg.Add(1)
//	//		//go handleRow(&wg,rowImportChan, rowErrorChan, rowChunk)
//	//		rowChunk = make([][]string,0)
//	//	}
//	//}
//	wg.Add(3)
//	go handleRow(&wg,rowImportChan, rowErrorChan, rows)
//	go importRow(&wg,rowImportChan)
//	go handleRowErrorMsg(&wg, rowErrorChan)
//	wg.Wait()
//	//time.Sleep(time.Second * 10)
//}

func main()  {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan *http.Response{
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:

				}
			}
		}()
		return responses
	}
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range checkStatus(done, urls...){
		fmt.Printf("Response: %v\n", response.Status)
	}
}

// php-csc32
func php_crc32(){
	data := []byte("appbOVM8zDL7885")
	//norm := math.Fmod(float64(*d), 360)
	mod := crc32.ChecksumIEEE(data)
	norm := math.Mod(float64(mod), 20)
	fmt.Println(norm)

}

// StringInt create a type alias for type int
type StringInt int

// UnmarshalJSON create a custom unmarshal for the StringInt
/// this helps us check the type of our value before unmarshalling it

func (st *StringInt) UnmarshalJSON(b []byte) error {
	//convert the bytes into an interface
	//this will help us check the type of our value
	//if it is a string that can be converted into an int we convert it
	///otherwise we return an error
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = StringInt(v)
	case float64:
		*st = StringInt(int(v))
	case string:
		///here convert the string into
		///an integer
		if v == "" {
			*st = 0
		} else {
			i, err := strconv.Atoi(v)
			if err != nil {
				///the string might not be of integer type
				///so return an error
				*st = 0
				return err
			}
			*st = StringInt(i)
		}
	}
	return nil
}
type KafkaEceEventMsg struct {
	UrlTimeUnix   int64  `json:"url_time_unix"`
	TrackId       string `json:"track_id"`
	PCUserId      string `json:"p_c_user_id"`
	PUnionId      string `json:"p_union_id"`
	PAppId        string `json:"p_app_id"`
	PPageId       string `json:"p_page_id"`
	PProductId    string `json:"p_product_id"`
	PResourceId   string `json:"p_resource_id"`
	PResourceType StringInt `json:"p_resource_type"`
	GeneratedAt   string `json:"generated_at"`
}

type KafkaTOrdersMsg struct {
	//Data []byte `json:"data"`
	Data OrderData `json:"data"`
}

type OrderData struct {
	AppId        string    `json:"app_id"`
	OrderId      string    `json:"order_id"`
	UserId       string    `json:"user_id"`
	ResourceId   string    `json:"resource_id"`
	ResourceType StringInt `json:"resource_type"`
	ProductId    string    `json:"product_id"`
	DiscountId   string    `json:"discount_id"`
	PaymentType  int       `json:"payment_type"`
	CuId         string    `json:"cu_id"`
	CouPrice     int       `json:"cou_price"`
	Source       int       `json:"source"`
	IsRenew      int       `json:"is_renew"`
	WxAppType    int       `json:"wx_app_type"`
	OrderState   int       `json:"order_state"`
	CreatedAt    string    `json:"created_at"`
	GeneratedAt  string    `json:"generated_at"`
}
//func main() {
//	//JSON反序列化：JSON格式的字符串-->结构体
//	str := `{"url_time_unix":1662691129730,"p__city":"宜春市","project":"c_production","type":"track","p_resource_type":"''","p__receive_time":1662691129735,"p_browser_env":"wechat","p__url":"https://apprpcdvfmj9458.h5.xiaoeknow.com/v1/goods/goods_detail/p_62760eabe4b09dda1264dfb1?type=3","p__model":"RMX3031-BUILD","p__url_path":"/v1/goods/goods_detail/p_62760eabe4b09dda1264dfb1","day":19244,"p_spu_id":"p_62760eabe4b09dda1264dfb1","login_id":"u_62b83ec1b0933_UhH0aRIAuY","p__ip":"182.102.160.213","create_time":1662691129735,"p_user_agent":"Mozilla/5.0 (Linux; Android 12; RMX3031 Build/SP1A.210812.016; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/4313 MMWEBSDK/20220805 Mobile Safari/537.36 MMWEBID/4557 MicroMessenger/8.0.27.2220(0x28001B53) WeChat/arm64 Weixin NetType/5G Language/zh_CN ABI/arm64","p_platform":"h5","p__browser_version":"8.0.27.2220","p_from_share":0,"p__lib_version":"1.29.2","p__manufacturer":"","p__is_first_day":true,"p__lib":"js","p_c_user_id":"u_62b83ec1b0933_UhH0aRIAuY","p__timezone_offset":-480,"distinct_id":"u_62b83ec1b0933_UhH0aRIAuY","p_shop_version_type":4,"track_id":"839869480","p_session_id":"sid_183201c59214942c0fb1fcf4befa29d34862","p__lib_method":"code","p__network_type":"5g","p_page_path":"C端#H5#商品#商品详情页","month_id":632,"p_sharer_id":"","p_resource_id":"p_62760eabe4b09dda1264dfb1","p__referrer":"https://apprpcdvfmj9458.h5.xiaoeknow.com/p/decorate/page/eyJpZCI6MTc4MDM4MSwiY2hhbm5lbF9pZCI6IiJ9","p__browser":"wchat","p_spu_type":"SPC","p__latest_traffic_source_type":"直接流量","p_pv_id":"pv_183201d3668984c0aa5103d3745bba431ca0","url_ip":"182.102.160.213","p_page_name":"商品详情页","p__province":"江西省","event":"view_page","p_user_channel":"","key":"view_page_C#h5#goods#goods_detail_nullff9ed771-057f-4c92-ac5c-3a8fc5358747","p__screen_height":889,"p_abtest_cookie":0,"p__os":"Android","p__latest_search_keyword":"未取到值_直接打开","p__latest_referrer":"","p__referrer_host":"apprpcdvfmj9458.h5.xiaoeknow.com","week_id":2749,"p_page_id":"C#h5#goods#goods_detail","p__country":"中国","p_l_program":"knowledge_shop","p_app_id":"apprPCDvfmj9458","_track_id":839869480,"url_project":"c_production","p_page_module":"H5","p__os_version":"12","anonymous_id":"18320092aec212-0480b4e96c7b83-7b7a3a7b-355600-18320092aee1fd","p_introduce_way":"-1","time":1662691129730,"original_id":"","p__screen_width":400}`
//	c1 := &KafkaEceEventMsg{}
//	err := json.Unmarshal([]byte(str), c1)
//	if err != nil {
//		fmt.Printf("json unmarshal failed!: %s", err)
//		return
//	}
//	fmt.Printf("%#v\n", c1)
//	//JSON序列化：结构体-->JSON格式的字符串
//	data, err := json.Marshal(c1)
//	if err != nil {
//		fmt.Println("json marshal failed")
//		return
//	}
//	fmt.Printf("json:%s\n", data)
//
//	order_str := `{"data":{"app_id":"appibud2bsi9969","content_app_id":null,"order_id":"o_1662691208_631aa788ea1e2_03898497","user_id":"u_62c00a3e4f09c_PfZKd5idcP","pay_way":0,"payment_type":3,"resource_type":6,"resource_id":"p_62f7dc45e4b0c94264874692","product_id":"p_62f7dc45e4b0c94264874692","count":1,"channel_id":"","channel_info":"","share_user_id":"u_61b5cdf75c8fc_PlAIJDyqpj","share_type":5,"purchase_name":"Y20造句场景篇（新版）","img_url":"https://wechatapppro-1252524126.file.myqcloud.com/appibud2bsi9969/image/b_u_61b480c52c675_HgS56Cr8/l2beo49k0317.jpeg","cu_id":"","cou_price":0,"discount_id":"","discount_price":0,"price":9000,"order_state":0,"goods_type":0,"ship_state":0,"out_order_id":null,"transaction_id":null,"wx_app_type":1,"period":null,"use_collection":2,"settle_status":0,"distribute_type":1,"que_check_state":0,"distribute_price":null,"distribute_percent":null,"superior_distribute_user_id":null,"superior_distribute_price":null,"superior_distribute_percent":null,"related_id":"","is_renew":0,"created_at":"2022-09-09 10:40:08","updated_at":"2022-09-09 10:40:09","source":0,"agent":"Mozilla/5.0 (Linux; Android 10; WLZ-AL10 Build/HUAWEIWLZ-AL10; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/4313 MMWEBSDK/20220805 Mobile Safari/537.36 MMWEBID/9870 MicroMessenger/8.0.27.2220(0x28001B3F) WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64","pay_time":"1970-01-01 08:00:00","settle_time":"1970-01-01 08:00:00","refund_time":"1970-01-01 08:00:00","refund_money":null,"invalid_time":"2022-09-09 12:40:08"},"op":"-U"}`
//	o1 := &KafkaTOrdersMsg{}
//	err1 := json.Unmarshal([]byte(order_str), o1)
//	if err1 != nil {
//		fmt.Println("json unmarshal failed!")
//		return
//	}
//	fmt.Printf("%#v\n", o1)
//	//JSON序列化：结构体-->JSON格式的字符串
//	order_data, err := json.Marshal(o1.Data)
//	if err != nil {
//		fmt.Println("json marshal failed")
//		return
//	}
//	fmt.Printf("json:%s\n", order_data)
//
//	//string 转 时间戳
//
//	stringTime := "2022-09-23 15:06:00"
//
//	loc, _ := time.LoadLocation("Local")
//
//	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
//	now_time := time.Now().Unix()
//	//_3_day_ahead_time := time.Now().Add(-time.Hour * 72).Unix()
//
//	if err == nil {
//
//		unix_time := the_time.Unix() //1504082441
//
//		fmt.Printf("unix_time: %d\n",unix_time)
//		fmt.Printf("3 day ahead of unix time: %d\n", the_time.Add(- time.Hour * 72).Unix())
//
//	}
//	fmt.Printf("now unix time: %d\n", now_time)
//}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

//func main(){
//	prefix := "ma_gc"
//	discount_id := "ma_g232"
//	//if strings.HasPrefix(discount_id, prefix){
//	//	fmt.Println(discount_id)
//	//} else {
//	//	fmt.Println(prefix)
//	//}
//	if HasPrefix(discount_id, prefix) {
//		fmt.Println(discount_id)
//	} else {
//		fmt.Println(prefix)
//	}
//	time := time.Now().Format("2006-01-02 15:04:05")
//	fmt.Println(time)
//}