package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"lorm/session"
)

func main() {
	engine, _ := session.NewEngine("mysql", "root:root@tcp(localhost:3306)/go_test?charset=utf8&parseTime=True&loc=Local")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
