package main

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

//二分递归
func sum2Arr(arr [5] int ,lo,hi int) int {
	if(lo == hi) {
		return arr[lo]
	}
	var mi int = (lo + hi) >> 1
	return sum2Arr(arr,lo,mi) + sum2Arr(arr,mi+1,hi)
}

func swap(a,b *int){
	*a,*b = *b,*a
}
//反转数组
func reverse(arr *[5] int,lo,hi int) {
	if(lo < hi){
		swap(&((*arr)[lo]),&((*arr)[hi]))
		reverse(&(*arr),lo+1 ,hi-1)
	}
}

//max2 获取数组最大的两个元素  方法1
func max2one(arr [5] int,lo,hi int)(int,int){
	var x1,x2 int
	x1 = lo
	x2 = lo+1
	if(arr[x1] < arr[x2]){
		swap(&x1,&x2)
	}
	for index := lo + 2; index < hi; index++ {
		if(arr[x2] < arr[x1]){
			x2 = index
			if(arr[x1] < arr[x2]){
				swap(&x1,&x2)
			}
		}
	}
	return x1,x2
}

//递归+分治
func max2two(arr [5] int,lo,hi int)(int,int){
	var x1,x2 int
	x1 = lo
	x2 = lo + 1
	if(lo + 1 == hi){
		if(arr[x1] < arr[x2]){
			x1,x2 = hi,lo
		}
		return x1,x2
	}
	//todo 三个元素的情况
	if(lo + 2 == hi){
		if(arr[x1] < arr[x2]){
			x1,x2 = lo+1,lo
		}
		if(arr[x2]<arr[lo+2]){
			x2 = lo+2
			if(arr[x1] < arr[x2]){
				x1,x2 = x2,x1
			}
		}
	}
	mi := (lo + hi)/ 2
	fmt.Printf("lo:%d,hi:%d,middle:%d \r\n",lo,hi,mi)
	x1L,x2L := max2two(arr,lo,mi)
	x1R,x2R := max2two(arr,mi,hi)
	if(arr[x1L] > arr[x1R]){
		x1 = x1L
		x2 = x2L
		if(arr[x2L] < arr[x1R]){
			x2 = x1R
		}
	}else{
		x1 = x1R
		x2 = x2R
		if(arr[x1L] > arr[x2R]){
			x2 = x2R
		}
	}
	return x1,x2
}
func main(){
	var arr = [5] int {10,20,30,4,1}
	sum1 := sumI(arr,5)

	fmt.Printf("sum1 = %d\r\n",sum1)
	sum2 := sumArr(arr,4)

	fmt.Printf("sun2 = %d\r\n",sum2)
	sum3 := sum2Arr(arr,0,len(arr) -1)

	fmt.Printf("sun3 = %d\r\n",sum3)
	reverse(&arr,0,4)

	for key,val := range arr{
		fmt.Printf("index %d is %d\r\n",key,val)
	}

	x1,x2 := max2one(arr,0,len(arr) -1 )
	fmt.Printf("first:%d,sec:%d\r\n",x1,x2)

	x11,x22 := max2two(arr,0,len(arr) - 1)
	fmt.Printf("first:%d,sec:%d\r\n",x11,x22)
}