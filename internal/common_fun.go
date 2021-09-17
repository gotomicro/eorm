package internal

import (
	"reflect"
	"unicode"
)

func TableName(table interface{}) string {
	t := reflect.TypeOf(table)
	tableName := ""
	if _, ok := t.MethodByName("TableName"); ok {
		tableName = reflect.ValueOf(table).MethodByName("TableName").Call(nil)[0].String()
	} else {
		tableName = underscoreName(t.Elem().Name())
	}
	return tableName
}

func underscoreName(tableName string) string {
	buf := []byte{}
	for i, v := range tableName {
		if unicode.IsUpper(v) {
			if i != 0 {
				buf = append(buf, '_')
			}
			buf = append(buf, byte(unicode.ToLower(v)))
		} else {
			buf = append(buf, byte(v))
		}

	}
	return string(buf)
}
