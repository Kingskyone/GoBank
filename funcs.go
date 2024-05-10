package main

import "runtime"

// GetCurrentFilePath 获取当前文件路径
func GetCurrentFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}
