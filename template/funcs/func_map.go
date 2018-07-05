package funcs

import (
	"text/template"
	"github.com/iunsuccessful/generator/rules"
	"reflect"
)

func CustomFuncMap() template.FuncMap {

	return template.FuncMap{
		"lowerCamelCase": rules.LowerCamelCase,
		"upperCamelCase": rules.UpperCamelCase,
		"lowerHyphenCase": rules.LowerHyphenCase,
		// {{range  $i, $e := .}}{{$e}}{{if last $i $e | not}}, {{end}}{{end}}.
		// https://play.golang.org/p/Y-CZ2TBSGU
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len() - 1
		},
	}

}
