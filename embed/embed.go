package main

import "fmt"

type notifer interface{
	notify()
}

type user struct{
	name string
	email string
}

type admin struct{
	user
	level string
}

func (user user) notify(){
	fmt.Printf("Send user email to %s<%s>",
		user.name,
		user.email)
}

func main(){
	ad := admin {
		user:user{
			name:"linjinzao",
			email:"123456789@qq.com",
		},
		level:"super",
	}

	sendNotifier(ad)
}

func sendNotifier(n notifer){
	n.notify()
}