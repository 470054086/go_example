package dialect

import "reflect"

var dialectMap = map[string]Dialect{}

// Dialect接口类
type Dialect interface {
	// 将Go语言类型为该数据库的数据类型
	DataTypeOf(typ reflect.Value) string
	// 返回某个表是否存在的 SQL 语句，参数是表名(table)
	TableExistSql(tableName string) (string, []interface{})
}
// 注册Dialect
func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}
// 获取Dialect
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
