package ddl

import (
	"container/list"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iunsuccessful/generator/config"
	"log"
)

func open() *sql.DB {
	cfg := config.GetFileConfig()
	dsn := fmt.Sprintf("%s:%s@%s?charset=utf8", cfg.Username, cfg.Password, cfg.Url)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	return db
}

/**
 * 查询封装到对象
 * 先一个一个写，后面再写通用的吧
 */
func selectTableInfo(tableInfo *TableInfo, selectSql string)  {
	//查询数据
	db := open()
	defer db.Close()
	rows, err := db.Query(selectSql)

	if err != nil {
		log.Fatalln(err)
	}

	//for rows.Next() {
	if rows.Next() {
		err = rows.Scan(&tableInfo.TableName, &tableInfo.TableComment)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func selectColumnInfo(selectSql string) *list.List {
	//查询数据
	db := open()
	defer db.Close()
	rows, err := db.Query(selectSql)

	if err != nil {
		log.Println(err)
	}

	columnList := list.New()

	for rows.Next() {
		columnInfo := &ColumnInfo{}
		err = rows.Scan(&columnInfo.ColumnName, &columnInfo.JDBCType, &columnInfo.ColumnComment, &columnInfo.ColumnKey)
		if err != nil {
			log.Println(err)
		}
		columnList.PushBack(columnInfo)
	}

	return columnList
}


