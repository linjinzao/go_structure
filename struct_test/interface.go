package main

import "fmt"


//声明一个notify行为的类型
type notifier interface{
	notify()
}

//声明一个user类型
type user struct{
	name string
	email string
}

func (user *user) notify(){
	fmt.Printf("Sending user email to %s<%s>",
	user.name,
	user.email)
}
//两个方法的差异 ：如果使用指针接受着来实现一个接口，那么只有指向那个类型的指针才能够实现对应的接口
//如果使用值接受者来实现一个接口，那么那个类型的值和指针都能够实现对应的接口
// func (user user) notify(){
// 	fmt.Printf("Sending user email to %s<%s>",
// 	user.name,
// 	user.email)
// }

func main(){
	bill := user{"linjinzao","526649566@qq.com"}
	//notify的实现是指针传递参数，所以只能指向那个类型的指针才有实现对应接口，故参数只能传递地址
	sendNotification(&bill)

}

//方法接受一个已经实现Notifier接口的值
func sendNotification(n notifier){
	n.notify()
}