package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Score int64
}

//模拟远程调用数据
func DoData(user *User){
	user.Score = int64(len(user.Name) * 10)
}

func CreateData() []*User{
	users := make([]*User, 0, 10)
	for i := 0; i < 10; i++ {
		u := &User{
			Name: fmt.Sprintf("name%d", i),
			Score: 2,
		}
		users = append(users, u)
	}
	return users
}
/*func main(){
	s := CreateData()
	var wg  sync.WaitGroup
	wg.Add(len(s))
	for _, item := range s {
		go func(i *User){
			defer wg.Done()
			DoData(i)
		}(item)
	}
	wg.Wait()
	//获取数据后下一步处理
	for _, val := range s {
		fmt.Println(val.Score)
	}
}*/
func Handler(number int, wg *sync.WaitGroup, s []*User, workFun func(user *User)){
	inch := make(chan *User, 0)
	//go-routine1 把需要处理的数据写入inch
	go func(){
		for _, item := range s{
			inch <- item
		}
		close(inch)
	}()
	//go-routine2: 开启number个协程， 同时读取inch的参数
	for i := 0; i < number; i++ {
		go func() {
			defer wg.Done()
			for item := range inch {
				workFun(item)
			}
		}()
	}
}
func main(){
	s := CreateData()
	var wg sync.WaitGroup
	wg.Add(len(s))
	Handler(len(s), &wg, s, DoData)
	wg.Wait()
	//得到后的数据后下一步处理
	for _, val := range s {
		fmt.Println(val.Score)
	}
}