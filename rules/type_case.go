package rules

import "strings"

func TypeCase(s string) string {

	switch strings.ToLower(s) {
		case "date", "time", "datetime", "timestamp": return "Date"
		case "float", "double", "numeric", "decimal": return "Double"
		case "int", "integer", "smallint", "tinyint": return "Integer"
		case "bigint": return "Long"
		case "varchar", "char", "longvarchar": fallthrough
		default: return "String"
	}

}

func JDBCTypeCase(s string) string {

	switch strings.ToLower(s) {
		case "date", "time", "datetime", "timestamp": return "TIMESTAMP"
		case "float", "double", "numeric", "decimal": return "DOUBLE"
		case "int", "integer", "smallint", "tinyint": return "INTEGER"
        case "bigint": return "BIGINT"
		case "varchar", "char", "longvarchar": fallthrough
		default: return "VARCHAR"
	}

}
