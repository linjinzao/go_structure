package vector

import (
	"reflect"
	"fmt"
)

type vector struct{
	values []interface{}  //类似泛型
}

//初始化
func New(cap int) *vector {
    this := new(vector)
    this.values = make([]interface{}, 0, cap)
 
    return this
}

//是否为空
func (this *vector) IsEmpty() bool{
	return len(this.values) == 0
}

//元素数量
func (this *vector) Size() int {
	return len(this.values)
}

// 追加单个元素
func (this *vector) Append(value interface{}) bool {
    this.values = append(this.values, value)
    return true
}

//添加元素--多个元素
func (this *vector) AppendAll(values []interface{}) bool {
    if values == nil || len(values) < 1 {
        return false
    }
    this.values = append(this.values, values...)
    return true
}
//插入元素
func (this *vector) Insert(rank int, value interface{}) bool {
    if rank < 0 || rank >= len(this.values) {
        return false
    }
    this.values = append(this.values[:rank], append([]interface{}{value}, this.values[rank:]...)...)
    return true
}
 
// 插入
func (this *vector) InsertAll(rank int, values []interface{}) bool {
    if rank < 0 || rank >= len(this.values) || values == nil || len(values) < 1 {
        return false
    }
    this.values = append(this.values[:rank], append(values, this.values[rank:]...)...)
    return true
}

func (this *vector) Remove(rank int) bool {
	if(rank < 0 || rank >= len(this.values)){
		return  false
	}
	//重置为Nil 防止内存泄漏
	this.values[rank] = nil
	this.values = append(this.values[:rank],this.values[rank+1:]...)
	return true
}

//移除范围内的值
func (this *vector) RemoveRange(lo,hi int) bool {
	if lo <0 || lo >= len(this.values) || hi > len(this.values) || lo > hi {
		return false
	}
	for i := lo; i < hi; i++ {
		this.values[i] = nil
	}

	this.values = append(this.values[:lo],this.values[hi:]...)
	return true
}

//循秩访问
func (this *vector) GetRank(value interface{}) int {
	for i := 0; i < len(this.values); i++ {
		if reflect.DeepEqual(this.values[i],value){
			return i
		}
	}

	return -1
}
//循秩访问
func (this *vector) Find(value interface{},lo,hi int) int {
	if(lo > hi){
		return -1
	}
	for i := lo; i < hi; i++ {
		if reflect.DeepEqual(this.values[i],value){
			return i
		}
	}

	return -1
}
func (this *vector) GetValue(rank int) interface{} {
	if rank < 0 || rank >= len(this.values) {
		return nil
	}

	return this.values[rank]
}

//去除重复元素
func (this *vector) Deduplicate() {
	for i := 1; i < len(this.values); i++ {
		rank := this.Find(this.values[i],0,i)
		if( rank >= 0){
			this.Remove(i)
		}
	}
}

//是否是有序向量--返回逆序对个数
func (this *vector) Disordered() int {
	var n,i int
	n = 0
	for i = 1; i < len(this.values); i++ {
		var a,b int
		a = this.values[i-1].(int)
		b = this.values[i].(int)
		if a > b  {
			n += 1
		}
	}
	return n
}


//有序向量去重
func (this *vector) Uniquify(){  //低效算法
	for i := 0;i < len(this.values) -1; i++ {
		if(reflect.DeepEqual(this.values[i],this.values[i+1])){
			this.Remove(i+1)
		}
	}
}

//二分查找
//转向左边比较一次，转向右边比较两次  
func (this *vector) BinSearch(target,lo,hi int) int {
START:
	if lo < hi {
		mi := (lo + hi) >> 1
		switch  {
		case target < this.values[mi].(int):
			hi = mi
		case this.values[mi].(int) < target:
			lo = mi
		default:
			return mi
		}
		goto START
	}
	return -1
		
}

//二分查找b
//平衡左右两边的比较次数，但需要最终才有结果
func (this *vector) BinSearchB(target,lo,hi int) int {
	START:
		if 1 < hi - lo {
			mi := (lo + hi) >> 1
			if target < this.values[mi].(int) {
				hi = mi
			}else{
				lo = mi + 1
			}
			goto START
		}
	
		if(target == this.values[lo].(int)){
			return lo
		}
		lo -= 1
		return lo
	}

//二分查找c
//平衡左右两边的比较次数，但需要最终才有结果  
//当lo=hi时，最后一次赋值的mi = lo + 1,所以返回lo-1相当于是中间的节点，只需要在最后一次判断mi是否等于target
//根据语义，返回不大于target的秩，所以可以直接返回lo-1
func (this *vector) BinSearchC(target,lo,hi int) int {
START:
	if lo < hi {
		mi := (lo + hi) >> 1
		if target < this.values[mi].(int) {
			hi = mi
		}else{
			lo = mi + 1
		}
		goto START
	}
	lo -= 1
	return lo
}


//冒泡排序
func (this *vector) BubbleSort(lo,hi int) {
START:
	hi -= 1
	if(!this.bubble(lo,hi)){
		goto START
	}
}

func (this *vector) bubble(lo,hi int) bool {
	var sorted bool = true
	START:
		lo += 1
		if(lo < hi){
			if(this.values[lo - 1].(int) > this.values[lo].(int)){
				sorted = false
				this.values[lo - 1],this.values[lo] = this.values[lo],this.values[lo - 1]
			}
			goto START
		}
	
	return sorted
}








