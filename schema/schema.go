package schema

import (
	"gee-orm/dialect"
	"go/ast"
	"reflect"
)

type Field struct {
	Name string
	Type string
	Tag  string // 约束条件
}

type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (s *Schema) GetField(name string) *Field {
	return s.fieldMap[name]
}

func (s *Schema) RecordValues(dest interface{}) []interface{} { // 将某个struct中的字段按照s.Fields的顺序构造一个参数列表返回
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range s.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if p.Anonymous || !ast.IsExported(p.Name) { // 该field匿名 或者 该field没有导出
			continue
		}
		field := &Field{
			Name: p.Name,
			Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
		}
		if v, ok := p.Tag.Lookup("geeorm"); ok {
			field.Tag = v
		}
		schema.Fields = append(schema.Fields, field)
		schema.FieldNames = append(schema.FieldNames, p.Name)
		schema.fieldMap[p.Name] = field
	}
	return schema
}
