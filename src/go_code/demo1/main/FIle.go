package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 读取文件
	file, err := os.Open("/Users/Rain/Desktop/pay/src/go_code/demo1/main/FIle.go")
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	var strSlice []byte

	for {
		var temp = make([]byte, 128)
		n, err := file.Read(temp)
		if err == io.EOF {
			fmt.Println("读取完成")
			break
		}
		if err != nil {
			fmt.Println("读取文件失败", err)
			return
		}
		strSlice = append(strSlice, temp[:n]...)
		//fmt.Printf("读取到%v 字节",n)
		//fmt.Println(temp)
		//fmt.Println(string(temp))
	}
	fmt.Println(string(strSlice))
	read, err := ioutil.ReadFile("/Users/Rain/Desktop/pay/src/go_code/demo1/main/FIle.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(read))

}
