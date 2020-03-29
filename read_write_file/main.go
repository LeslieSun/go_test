package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 判断文件/文件夹是否存在
func Exists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//判断是否为文件夹
func IsDir(path string) (bool, error) {
	s, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return s.IsDir(), nil
}

//判断是否为文件
func IsFile(file string) (bool, error) {
	res, err := IsDir(file)
	return !res, err
}
func main() {
	fileName := "./test"
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	contentByte, err := ioutil.ReadAll(f)
	fmt.Printf("%s", string(contentByte))
	_, err = f.Write([]byte("test write to file\n"))
}
