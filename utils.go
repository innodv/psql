package psql

import (
	"fmt"
	"reflect"
)

func GetSQLColumns[T any]() []string {
	var columns []string
	structFields := reflect.VisibleFields(reflect.TypeOf((*T)(nil)).Elem())
	for _, field := range structFields {
		tag, ok := field.Tag.Lookup("db")
		if ok && tag != "-" && len(tag) > 0 {
			columns = append(columns, tag)
		}
	}
	return columns
}

func GetSQLColumnsQuoted[T any]() []string {
	cols := GetSQLColumns[T]()
	for i, col := range cols {
		cols[i] = fmt.Sprintf("\"%s\"", col)
	}
	return cols
}

func GetPostgreSQLPreparedArgList(cols []string) string {
	out := ""
	for i := range cols {
		if i == 0 {
			out += "$1"
		} else {
			out += fmt.Sprintf(",$%d", i+1)
		}
	}
	return out
}
