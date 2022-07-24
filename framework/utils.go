package framework

import "os"

func GetExecDirectory() string {
	file, err := os.Getwd()
	if err == nil {
		return file + "/"
	}
	return ""
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
