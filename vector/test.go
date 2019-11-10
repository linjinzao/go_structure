package main

import (
	"fmt"
	"vector"
)
func main(){

    v := vector.New(5)
    for i := 0; i < 6; i++ {
        v.Append(i)
    }
    fmt.Println(*v)
    fmt.Println(v.IsEmpty())
    v.Insert(1, 10)
    v.Insert(20, 20)
    fmt.Println(*v)
    v.Remove(-1)
    v.Remove(1)
	v.Remove(100)
	fmt.Println(*v)
    fmt.Println(v.GetValue(1))
    fmt.Println(v.GetValue(100))
}