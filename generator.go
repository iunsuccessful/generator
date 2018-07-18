package main

import (
	"os"
	"log"
	"github.com/iunsuccessful/generator/template"
	"github.com/iunsuccessful/generator/config"
)

/**
 * 程序入口
 *
 * 1. 读取配置文件
 * 2. 连接数据库读取 DDL
 * 3. 读取模板目录生成目标文件
 *
 */
func main() {

	args := config.New(os.Args)

	log.Println("Begin generator...")
	template.Render(args)
	log.Println("success.")

}



