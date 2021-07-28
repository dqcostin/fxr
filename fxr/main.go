package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	content := getContent("./result.txt")
	//removeDuplicate(content)
	str := CutValue(content)
	//fmt.Println(str)
	remove := removeDuplicate(str)
	//fmt.Println(remove)
	var filename = "./test.txt"
	var f *os.File
	for _, in := range remove {
		// fmt.Println(in)
		if checkFileIsExist(filename) { //如果文件存在
			f, _ = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
			fmt.Println("文件存在")
		} else {
			f, _ = os.Create(filename) //创建文件
			fmt.Println("文件不存在")
		}
		defer f.Close()
		n, err1 := io.WriteString(f, in+"\r\n") //写入文件(字符串)
		if err1 != nil {
			panic(err1)
		}

		fmt.Printf("写入 %d 个字节n", n)
	}
	xray()
}

func getContent(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}

	return string(fd)

}

func CutValue(a string) []string {
	regexp, _ := regexp.Compile("[a-z]*://[0-9]*[.][0-9]*[.][0-9]*[.][0-9]*:[0-9]*")
	str1 := regexp.FindAllString(a, -1)
	return str1
}

func removeDuplicate(a []string) (ret []string) {

	len1 := len(a)
	for i := 0; i < len1; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
