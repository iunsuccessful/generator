package path

import (
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

/**
 * 获取文件执行路径
 */
func GetExecPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return pwd
}

func GetBinFilePath() string {
	//return "E:\\Anonymous\\Documents\\golang\\Generator\\src\\github.com\\iunsuccessful\\generator\\test";
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return filepath.Dir(path)
}

/**
 * 应用程序路径
 */
func GetBasePath() string {
	return GetExecPath()
	//return GetBinFilePath()
}

func GetTemplatePath() string {
	return filepath.Join(GetBasePath(), "template")
}

func GetTargetPath() string {
	return filepath.Join(GetBasePath(), "target")
}

/**
 * 递归获取目录中所有文件路径
 */
func GetAllFiles() []string {
	return getFiles(filepath.Join(GetBasePath(), "template"))
}

/**
 * 获取目录下的所有文件
 */
func getFiles(dirPath string) []string {
	filePaths := make([]string, 0)
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	return filePaths
}

func Create(targetPath string) *os.File {
	targetPath = strings.Replace(targetPath, "\\", "/", -1)
	err := os.MkdirAll(path.Dir(targetPath), os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	newFile, err := os.Create(targetPath)
	if err != nil {
		log.Println(err)
	}
	return newFile
}
