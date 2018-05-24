package main

import (
	"log"
	"github.com/iunsuccessful/generator/template"
	"os"
	"fmt"
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

	if len(os.Args) < 2 || os.Args[1] == "-h" {
		// print help
		usage()
	}

	log.Println("Begin generator...")

	if len(os.Args) == 2 {
		template.Render(os.Args[1], "")
	} else {
		template.Render(os.Args[1], os.Args[2])
	}
	log.Println("success.")
}

func usage() {
	fmt.Printf("Usage: %s <table name> [alias]\n", "generator");
	fmt.Println()
	fmt.Printf("\t -h \t\thelp\n");
	os.Exit(0)
}

