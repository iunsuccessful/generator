package ddl

import (
	"github.com/iunsuccessful/generator/rules"
)

type TableInfo struct {
	TableName      string // 原始表名
	BeanName	   string // Java bean Name
	TableComment   string
	Columns		   []ColumnInfo
}

func (t *TableInfo) Name() string {
	if len(t.BeanName) > 0 {
		return rules.UpperCamelCase(t.BeanName)
	}
	return rules.UpperCamelCase(t.TableName)
}

type ColumnInfo struct {
	ColumnName        string
	JDBCType  		  string
	ColumnComment     string
	ColumnKey		  string // 可判断是不是主键 “PRI” 为主键
}

func (c *ColumnInfo) UpperCamelName() string {
	return rules.UpperCamelCase(c.ColumnName)
}

func (c *ColumnInfo) LowerCamelName() string {
	return rules.LowerCamelCase(c.ColumnName)
}

func (c *ColumnInfo) DataType() string {
	return rules.TypeCase(c.JDBCType)
}

func (c *ColumnInfo) StdJDBCType() string {
	return rules.JDBCTypeCase(c.JDBCType)
}