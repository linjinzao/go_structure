package vector

import (
    "reflect"
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
func (this *vector) Insert(rank int,values interface{}) bool{
	if rank < 0 || rank >= len(this.values) {
		return false
	}
	this.values = append(this.values[:rank], append([]interface{}{values}, this.values[rank:]...)...)
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

func (this *vector) GetValue(rank int) interface{} {
	if rank < 0 || rank >= len(this.values) {
		return nil
	}

	return this.values[rank]
}










