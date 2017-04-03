// main package adds core functions to jstpl
package main

import (
	"fmt"
	"reflect"
)

func init() {
	templateFuncs["list"] = list
}

func list(s interface{}) string {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		return ""
	}

	v := reflect.ValueOf(s)

	var out string
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		out += fmt.Sprintf("%v\n", elem.Interface())
	}
	return out
}
