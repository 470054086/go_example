package lodash

import (
	"math"
	"reflect"
)

type ASlice []interface{}
type fun func(interface{}) bool
type funMap func(v interface{}) interface{}
type funRedure func(res interface{},v interface{},k interface{}) interface{}
func NewSliceInterface(s []interface{}) ASlice {
	var aSlice ASlice
	for _, val := range s {
		aSlice = append(aSlice, val)
	}
	return aSlice
}

func NewSliceString(s []string) ASlice  {
	var aSlice ASlice
	for _, val := range s {
		aSlice = append(aSlice, val)
	}
	return aSlice
}

func NewSliceInt(s []int) ASlice  {
	var aSlice ASlice
	for _, val := range s {
		aSlice = append(aSlice, val)
	}
	return aSlice
}
func NewSliceFloat32(s []float32) ASlice  {
	var aSlice ASlice
	for _, val := range s {
		aSlice = append(aSlice, val)
	}
	return aSlice
}

func NewSliceFloat64(s []float64) ASlice  {
	var aSlice ASlice
	for _, val := range s {
		aSlice = append(aSlice, val)
	}
	return aSlice
}

func (s ASlice) IsEmpty() bool {
	return s == nil || len(s) == 0
}

// 切割两个函数
func (s ASlice) Chunk(i int) []ASlice {
	res := []ASlice{}
	skip := int(math.Floor(float64(len(s)) / float64(i)))
	for j := 0; j <= len(s); j = j + skip {
		end := j + skip
		if end > len(s) {
			end = len(s)
		}
		res = append(res, s[j:end])
	}
	return res
}

//  判断i是否存在
func (s ASlice) In(i interface{}) bool {
	flag := false
	for _, val := range s {
		if i == val {
			flag = true
			break
		}
	}
	return flag
}

//  连接传进来的变量
func (s ASlice) Concat(i ...interface{}) ASlice {
	return append(s,i...)
}

// 创建一个切片数组，去除array前面的n个元素 n默认值为1
func (s ASlice) Drop(i int) ASlice  {
	return s[i:]
}
// 创建一个切片数组，去除array尾部的n个元素
func (s ASlice) DropRight(i int) ASlice  {
	return s[0:len(s)-i]
}
// 根据key删除某个位置
func (s *ASlice) RemoveIndex(i int) interface{}  {
	remove := (*s)[i]
	if i == 0  {
		*s = (*s)[1:]
	}else if i ==  len(*s) {
		*s = (*s)[:len(*s)-1]
	}else {
		first := (*s)[0:i]
		end := (*s)[i+1:len(*s)]
		*s = append(first,end...)
	}
	return remove
}

// 查找value出现的位置 未查询到的话 返回-1
func (s ASlice) FindIndex(v interface{}) int {
	k:= -1;
	for key,val := range s {
		if(val == v) {
			k = key
			break;
		}
	}
	return k;
}
//  获取头部元素
func (s ASlice) First() interface{}  {
	return s[0];
}

// Filter函数
func (s *ASlice) Filter(f fun ) ASlice {
	// 如果b存在的话,就从s 中去掉这个元素
	for key,val := range *s {
		if b := f(val); b {
			s.RemoveIndex(key)
		}
	}
	return *s;
}

// Map函数
func (s *ASlice) Map(f funMap)  {
	for key,val := range *s  {
		(*s)[key] = f(val)
	}
}
// reduce 函数
func (s *ASlice) Reduce(f funRedure,sum interface{}) interface{} {
	of := reflect.TypeOf(sum)
	switch of.Kind() {
	case reflect.Int:
		sum = sum.(int)
	case reflect.String:
		sum = sum.(string)
	default:
		panic("error")
	}
	for key,val := range *s {
		sum = f(sum,val,key)
	}
	return sum
}



