package funcs

import (
	"text/template"
	"github.com/iunsuccessful/generator/rules"
	"reflect"
	"time"
)

func CustomFuncMap() template.FuncMap {

	return template.FuncMap{
		"lowerCamelCase": rules.LowerCamelCase,
		"upperCamelCase": rules.UpperCamelCase,
		"lowerHyphenCase": rules.LowerHyphenCase,
		"datetime": func() string {
			return time.Now().Format("2006/1/2 15:04:05")
		},
		// {{range  $i, $e := .}}{{$e}}{{if last $i $e | not}}, {{end}}{{end}}.
		// https://play.golang.org/p/Y-CZ2TBSGU
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len() - 1
		},
	}

}
