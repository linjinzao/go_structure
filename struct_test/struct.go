package main

import "fmt"

//声明user类型
//每个类型相当于php中的类，类中可以有属性，可以是方法，可以是属性
type user struct{
	name string
	email string
}

//为user类型添加方法
//不需要修改，不用穿指针类型
func(user user) notify(){
	fmt.Printf("Send User Email To %s<%s>",
		user.name,
		user.email)
}
//需要修改元数据，传指针类型
func(user *user) changeEmail(email string){
	user.email = email
}

func main(){
	//赋值一个类型后相当于new一个对象，可以使用这个对象的属性跟方法
	bill := user{"bill","526649566@qq.com"}
	//触发方法
	bill.notify()
	
	bill.changeEmail("15112360658@163.com")

	bill.notify()

}