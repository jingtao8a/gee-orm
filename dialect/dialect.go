package dialect

import "reflect"

var dialectsMap = make(map[string]Dialect)

type Dialect interface {
	DataTypeOf(typ reflect.Value) string                    // 将Go语言的类型转为该数据库的数据类型
	TableExistSql(tableName string) (string, []interface{}) // 返回某个表是否存在的SQL语句，参数是表名(TABLE)
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (Dialect, bool) {
	dialect, ok := dialectsMap[name]
	return dialect, ok
}
