package main

import(
	"fmt"
	"strings"
)



func LCS1(s1,s2 []string) string {
	//#1.递归基
	if(len(s1) == 0 || len(s2) == 0){
		return ""
	}
	var s1Len,s2Len int
	s1Len = len(s1)
	s2Len = len(s2)
	//如果末尾元素相同  删掉末尾元素并返回
	if(s1[s1Len - 1] == s2[s2Len - 1]){
		return s1[s1Len - 1] + LCS1(s1[0:s1Len - 1],s2[0:s2Len - 1])
	}
	if(s1Len > s2Len){
		return LCS1(s1[0:s1Len - 1],s2)
	}
	return LCS1(s1,s2[0:s2Len - 1])
}
func main(){
	var str1,str2 string
	str1 = "ascxzcsqwesd"
	str2 = "cxvxcsdfasd"
	s1Arr := strings.Split(str1,"")
	s2Arr := strings.Split(str2,"")
	lcs := LCS1(s1Arr,s2Arr)
	fmt.Println(lcs)
}