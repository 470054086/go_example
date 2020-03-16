package test

import (
	"strconv"
	"testing"
	"time"
)


func Benchmark_GoCachePut( b *testing.B)  {
	for i:=0;i<=b.N;i++ {
		C.Put("name"+strconv.Itoa(i),"value"+strconv.Itoa(i))
	}
}
func Benchmark_GoCachePutExpire( b *testing.B)  {
	for i:=0;i<=b.N;i++ {
		C.PutExpire("name"+strconv.Itoa(i),"value"+strconv.Itoa(i),time.Duration(i))
	}
}


func Benchmark_GoCacheGet( b *testing.B)  {
	for i:=0;i<=b.N;i++ {
		C.Get("name"+strconv.Itoa(i))
	}
}