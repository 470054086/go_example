package session

import (
	"database/sql"
	"lorm/clause"
	"lorm/dialect"
	"lorm/log"
	"lorm/schema"
	"strings"
)

type Session struct {
	db      *sql.DB         //sql链接
	sql     strings.Builder //解析sql字符串
	refTable *schema.Schema // 表解析结构
	clause clause.Clause
	sqlVars []interface{}   //sql的参数
	dialect dialect.Dialect
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{db: db, dialect: dialect}
}

// 清除sql和参数
func (s *Session) clear() {
	s.sqlVars = nil
	s.sql.Reset()
}

// 返回一个db链接
func (s *Session) DB() *sql.DB {
	return s.db
}

// 增加sql和参数
func (s *Session) Raw(sql string, vars ...interface{}) *Session {
	s.sqlVars = append(s.sqlVars, vars...)
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	return s
}

// 执行sql语句
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
		return nil, nil
	}
	return
}

// 返回一行
func (s *Session) QueryRow() *sql.Row {
	defer s.clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// 返回多行
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
		return
	}
	return
}
