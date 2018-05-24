package template

import (
	"github.com/iunsuccessful/generator/path"
	"fmt"
	"github.com/iunsuccessful/generator/ddl"
	"os"
	"text/template"
	"bytes"
	"github.com/iunsuccessful/generator/rules"
	"path/filepath"
	"github.com/iunsuccessful/generator/template/funcs"
)

/*
 * 1. 根据路径读取文件
 * 2. 填充文件内容
 * 3. 写入 target 对应目录
 */
func Render(tableName, alias string)  {

	// 加载表格信息
	tableInfo := ddl.GetTableInfo(tableName)
	tableInfo.BeanName = rules.UpperCamelCase(alias)

	// 先删除 target 目录
	os.RemoveAll(path.GetTargetPath())
	// TODO 这里可用 t.ExecuteTemplate(w, tmpl, data) 改写
	for _, filePath := range path.GetAllFiles() {

		if ignoreFiles(filePath) {
			continue
		}

		// 增加自定义方法
		temp := template.New(filepath.Base(filePath)).Funcs(funcs.CustomFuncMap())
		if temp, err := temp.ParseFiles(filePath); err == nil {
			targetPath := path.GetTargetPath() + filePath[len(path.GetTemplatePath()):]
			targetPath = templateDir(targetPath, tableInfo)
			newFile := path.Create(targetPath)
			defer newFile.Close()
			if err := temp.Execute(newFile, tableInfo); err != nil {
				fmt.Println("There was an error:", err.Error())
			}
		}
	}

}

/**
 * 把路径也通过模板转换一下
 */
func templateDir(dir string, data interface{}) string {
	t, _ := template.New("path-template").Parse(dir)
	b := bytes.NewBuffer(make([]byte, 0))
	t.Execute(b, data)
	return b.String()
}

func ignoreFiles(filePath string) bool {
	return filepath.Base(filePath) == "README.md";
}