package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	file.Close()
}

/*
func Create(name string) (file *File, err Error)
根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666
func NewFile(fd uintptr, name string) *File
根据文件描述符创建相应的文件，返回一个文件对象
func Open(name string) (file *File, err Error)
只读方式打开一个名称为name的文件
func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限
func (file *File) Write(b []byte) (n int, err Error)
写入byte类型的信息到文件
func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
在指定位置开始写入byte类型的信息
func (file *File) WriteString(s string) (ret int, err Error)
写入string信息到文件
func (file *File) Read(b []byte) (n int, err Error)
读取数据到b中
func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
从off开始读取数据到b中
func Remove(name string) Error
删除文件名为name的文件



*/
