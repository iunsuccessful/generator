package main

import (
	"github.com/iunsuccessful/generator/config"
	"github.com/iunsuccessful/generator/path"
	"github.com/iunsuccessful/generator/template"
	"log"
	"os"
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
	log.Println("bin file path: ", path.GetBinFilePath())
	log.Println("exec file path: ", path.GetExecPath())
	log.Println("db path: ", config.GetFileConfig().Url)
	log.Println("begin generator...")
	template.Render(args)
	log.Println("success.")
}
