package config

import (
	"os"
	"fmt"
)

type Argument struct {
	TableName      string // 原始表名
	BeanName	   string // Java bean Name
	DeleteFolder   bool // 是否清空文件夹
}

func New(args []string) *Argument {
	if len(os.Args) < 2 {
		// print help
		usage()
	}

	argument := &Argument{ DeleteFolder: false }
	for i := 1; i < len(args); i++ {
		if (args[i][0] == '-') || (args[i][0] == '/') {
			switch args[i][1] {
				case 'd': argument.DeleteFolder = true
				case 'h': usage()
			}
		} else if len(argument.TableName) == 0 {
			argument.TableName = args[i]
		} else if len(argument.BeanName) == 0 {
			argument.BeanName = args[i]
		}
	}
	return argument
}

func usage() {
	fmt.Println()
	fmt.Printf("Usage:\n\n %s [options] <table name> [alias]\n", "generator")
	fmt.Println()
	fmt.Printf("\t -h \t\thelp\n")
	fmt.Printf("\t -d \t\tdelete old generator files.\n")
	fmt.Println()
	os.Exit(0)
}