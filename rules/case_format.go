package rules

import "strings"

/*
 * 1. 类型检查，判断 string 是什么命名法
 * 2. 命名之间互相转换
 */
func UpperCamelCase(s string) string {
	return camelCase(s, true)
}

func LowerCamelCase(s string) string {
	return camelCase(s, false)
}

// TODO 转换中线线，URL 中用
func LowerHyphenCase(s string) string {
	return s;
}

func camelCase(s string, upper bool) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return ""
	}
	buffer := make([]rune, 0, len(s))

	if strings.ContainsAny(s, "_-") {
		var prev rune
		for _, curr := range s {
			if !isDelimiter(curr) {
				if isDelimiter(prev) || (upper && prev == 0) {
					buffer = append(buffer, toUpper(curr))
				} else {
					buffer = append(buffer, toLower(curr))
				}
			}
			prev = curr
		}

		return string(buffer)
	}

	if upper {
		return strings.ToUpper(s[0:1]) + s[1:]
	} else {
		return strings.ToLower(s[0:1]) + s[1:]
	}

}

// isSpace checks if a character is some kind of whitespace.
func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// isDelimiter checks if a character is some kind of whitespace or '_' or '-'.
func isDelimiter(ch rune) bool {
	return ch == '-' || ch == '_' || isSpace(ch)
}

func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

func toUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}
	return ch
}

//func Create(chaosName string)  {
//
//}
//
//type CaseFormat struct {
//
//}
//
//func (c *CaseFormat) convert(format CaseFormat, s string)  {
//
//}
//
//func (c *CaseFormat) NormalizeWord() string {
//	return ""
//}
//
//func (c *CaseFormat) NormalizeFirstWord(word string) string {
//	fmt.Println("test")
//	return ""
//}
//
///**
// * e.g., "lowerCamel".
// */
//type LOWER_CAMEL struct {
//	CaseFormat
//}
//
///**
// * e.g., "UpperCamel"
// */
//type UPPER_CAMEL struct {
//	CaseFormat
//}
//
///**
// * e.g., "UPPER_UNDERSCORE"
// */
//type UPPER_UNDERSCORE struct {
//	CaseFormat
//}
//
///**
// * e.g., "lower_underscore"
// */
//type LOWER_UNDERSCORE struct {
//	CaseFormat
//}

