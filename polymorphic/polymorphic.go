package main

import "fmt"

//声明一个notifier的接口
type notifier interface{
	notify()
}

//声明 user ,admin 类型
type user struct{
	name string
	email string
}

type admin struct{
	name string 
	email string
}

//值接收者实现user类型的接口
func (user user) notify(){
	fmt.Printf("Send user emial to %s<%s>",
		user.name,
		user.email)
}

//指针接收者实现admin类型的接口
func (admin *admin) notify(){
	fmt.Printf("Send admin email to %s<%s>\r\n",
		admin.name,
		admin.email)
}

func main(){
	bill := user{"bill","123456789@qq.com"}
	sendNotifition(bill)

	lisa := admin{"lisa","987654321@qq.com"}
	sendNotifition(&lisa)
}

func sendNotifition(n notifier){
	n.notify()
}