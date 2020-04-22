package session

import (
	"fmt"
	"testing"
)

func TestRecord_Insert(t *testing.T) {
	engine, _ := NewEngine("mysql", "root:root@tcp(localhost:3306)/go_test?charset=utf8&parseTime=True&loc=Local")
	user := User{
		Name: "xiaoli",
		Age:  20,
	}
	_, _ = engine.NewSession().Insert(&user)
}

func TestRecord_Find(t *testing.T) {
	engine, _ := NewEngine("mysql", "root:root@tcp(localhost:3306)/go_test?charset=utf8&parseTime=True&loc=Local")
	user := [] User{}
	_ = engine.NewSession().Find(&user)
	fmt.Println(user)
}
