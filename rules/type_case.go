package rules

import "strings"

func TypeCase(s string) string {

	switch strings.ToLower(s) {
		case "date", "time", "datetime", "timestamp": return "Date"
		case "float", "double", "numeric", "decimal": return "Double"
		case "int", "integer", "smallint", "tinyint", "bigint": return "Integer"
		case "varchar", "char", "longvarchar": fallthrough
		default: return "String"
	}

}

func JDBCTypeCase(s string) string {

	switch strings.ToLower(s) {
		case "date", "time", "datetime", "timestamp": return "TIMESTAMP"
		case "float", "double", "numeric", "decimal": return "DOUBLE"
		case "int", "integer", "smallint", "tinyint", "bigint": return "INTEGER"
		case "varchar", "char", "longvarchar": fallthrough
		default: return "VARCHAR"
	}

}
