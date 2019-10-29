package main;

import "fmt"

//迭代求解
func sumI(arr [5] int,n int) int {
	var sum int = 0
	s := arr[:n]
	for _,val := range s {
		sum += val
	}
	return sum
}
//递归求解
func sumArr(arr [5] int ,n int) int {
	if(0>n){
		return 0
	}
	return sumArr(arr,n-1) + arr[n]
}

//反转数组
func reverse(arr *[5] int,lo,hi int) {
	if(lo < hi){
		var temp int 
		temp = (*arr)[lo]
		(*arr)[lo] = (*arr)[hi]
		(*arr)[hi] = temp
		reverse(&(*arr),lo+1 ,hi-1)
	}
}

func main(){
	 var arr = [5] int {10,20,30,4,1}
	 sum1 := sumI(arr,5)
	fmt.Printf("sum1 = %d",sum1)
	sum2 := sumArr(arr,4)
	fmt.Printf("sun2 = %d",sum2)

	reverse(&arr,0,4)

	for key,val := range arr{
		fmt.Printf("index %d is %d",key,val)
	}
}