package main

import "fmt"


//动态规划

func fib3(n int) int{
	var f,g int = 0,1
	for index := 0; index <= n-2; index++ {
		f,g = g,g+f
	}
	return g
}
//回归迭代
func fib2(n int) int{
	var fibArr = []int{0,1}
	if(2 > n){
		return fibArr[n]
	}
	for index := 2; index <= n; index++ {
		fibArr = append(fibArr,fibArr[index-1]+fibArr[index-2])
	}
	return fibArr[n];
}

func fib1(n int) int{
	if( 2 > n ){
		return n
	}
	return fib1(n-1) + fib1(n-2)
}
func main(){
	var n int = 30;
	fmt.Printf("fib(%d) = %d",n,fib1(n))
	fmt.Printf("fib(%d) = %d",n,fib2(n))
	fmt.Printf("fib(%d) = %d",n,fib3(n))
}