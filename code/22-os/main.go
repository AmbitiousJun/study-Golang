package main

import (
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 在当前目录下创建一个名为fileName的文件
func CreateFile(fileName string) {
	_, err := os.Create(fileName)
	check(err)
}

// 创建一个目录
func MakeDir(path string) {
	err := os.Mkdir(path, os.ModePerm)
	check(err)
}

// 将path沿途的文件夹都创建出来
func MakeDirAll(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	check(err)
}

// 删除文件夹
func RemoveDir(path string) {
	err := os.Remove(path)
	check(err)
}

// 将path沿途的文件夹都删除
func RemoveDirAll(path string) {
	err := os.RemoveAll(path)
	check(err)
}

// 获取当前的工作目录
func GetWd() {
	dir, err := os.Getwd()
	check(err)
	fmt.Println(dir)
}

// 读取文件
func ReadFile() {
	b, err := os.ReadFile("./22-os/test.txt")
	check(err)
	fmt.Println(string(b))
}

// 写入文件
func WriteFile() {
	str := "Hello World!"
	err := os.WriteFile("./22-os/test.txt", []byte(str), os.ModePerm)
	check(err)
}

func main() {
	// CreateFile("./22-os/create_test.txt")
	// MakeDir("./22-os/a")
	// MakeDirAll("./22-os/a/b/c")
	// GetWd()
	// ReadFile()
	WriteFile()
}
