package ddl

import (
	"fmt"
	"container/list"
)

/**
 *	获取表的信息
 *
 * select table_name, table_comment, create_time, update_time
 *		from information_schema.tables
 *		where table_comment != '' and table_schema = (select database())
 *		and table_name = 'u_user'
 */
func GetTableInfo(tableName string) *TableInfo {
	sql := fmt.Sprintf("select table_name, table_comment from information_schema.tables where table_schema = (select database()) and table_name = '%s'", tableName)
	tableInfo := &TableInfo{}
	selectTableInfo(tableInfo, sql)

	list := GetColumnInfo(tableName)

	//columns := make([]ColumnInfo, list.Len())
	columns := make([]ColumnInfo, 0)

	for e := list.Front(); e != nil; e = e.Next() {
		columns = append(columns, *e.Value.(*ColumnInfo))
	}

	tableInfo.Columns = columns;

	return tableInfo;
}

/**
 * 获取表的结构信息
 *
 * select column_name, data_type, column_comment from information_schema.columns
 *     where table_name = #{tableName} and table_schema = (select database()) order by ordinal_position
 *
 * 增加 column_key 标识主键
 */
func GetColumnInfo(tableName string) *list.List {
	sql := fmt.Sprintf("select column_name, data_type, column_comment, column_key from information_schema.columns where table_name = '%s' and table_schema = (select database()) order by ordinal_position", tableName)
	return selectColumnInfo(sql)
}




