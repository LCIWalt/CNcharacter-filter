package main

import (
	//"unsafe"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	//"path"
	"log"
	"path/filepath"
	"strings"
)

var lock sync.Mutex

//import "github.com/gin-gonic/gin"
var relativePath string= "E:/newE"
func main() {
	
	
	filepath.Walk(relativePath, func(path string, info os.FileInfo, err error) error {
		//fmt.Println(filepath.Ext("./path.exe")) //查看文件后缀
		ok, err := filepath.Match(`*.txt`, info.Name())

		if ok {
			fmt.Println(filepath.Dir(path), info.Name())
			// 遇到 txt 文件则继续处理所在目录的下一个目录
			// 注意会跳过子目录
			var realroad string
			realroad = filepath.Dir(path) + "\\" + info.Name()
			fmt.Println(realroad)
			//打开文件io
			relativePathword, err := os.Open(realroad)
			if err != nil {
				fmt.Println("the file read wrong")
			}
			if err == nil {
				fmt.Println("the file read ojbk")
			}
			defer relativePathword.Close()
			if err == io.EOF {
				fmt.Printf("sth going wr")
			}
			readbystr(realroad)
			//uu(realroad)
		}
		//lock.Unlock()
		if err!= nil{
			log.Fatal(err)
		}
		//fmt.Println(path) //打印path信息 能显示子目录名字
		//fmt.Println(filepath.Dir(relativePath))  这个不行
		//fmt.Println(info.Name()) //打印文件或目录名
		return nil
	})

	return
}
func listFiles(dirname string, level int) {
	// level用来记录当前递归的层次
	// 生成有层次感的空格
	//列出文件
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|   " + s
	}
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			//继续遍历fi这个目录
			listFiles(filename, level+1)
		}
	}
}

func readTxt(r io.Reader) ([]string, error) {
	reader := bufio.NewReader(r)

	l := make([]string, 0, 64)

	// 按行读取
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		l = append(l, strings.Trim(string(line), " "))
	}

	return l, nil
}
func readbystr(str string) {
	buf, err := ioutil.ReadFile(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lock.Lock()
	in := string(buf)
	if in = strings.Replace(in, "。", ".", -1); err != nil {
		fmt.Println(err)
	}
	in = strings.Replace(in, "，", ",", -1)
	in = strings.Replace(in, "”", "\"", -1)
	in = strings.Replace(in, "“", "\"", -1)
	in = strings.Replace(in, "：", ":", -1)
	in = strings.Replace(in, "？", "?", -1)
	fmt.Println(in)
	s := bufio.NewScanner(strings.NewReader(in))
	s.Split(bufio.ScanRunes)
	lock.Unlock()
	for s.Scan() {
		if s.Text() == "。" {

			ty := s.Text()
			ty = "."
			//dd := bufio.New
			fmt.Println(ty)
			continue
		}
		fmt.Println(s.Text())
	}
	if ioutil.WriteFile(str, []byte(in), 0777); err != nil {
		fmt.Println(err)
	}

	//fmt.Println(in)
}
func uu(s string) {
	s = os.Args[1]
	in, err := os.Open(s)
	br := bufio.NewReader(in)

	out, err := os.OpenFile(s+".mdf", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("Open write file fail:", err)
		os.Exit(-1)
	}
	defer out.Close()
	index := 1
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(-1)
		}
		newLine := strings.Replace(string(line), os.Args[2], os.Args[3], -1)
		_, err = out.WriteString(newLine + "\n")
		if err != nil {
			fmt.Println("write to file fail:", err)
			os.Exit(-1)
		}
		fmt.Println("done ", index)
		index++
	}
	fmt.Println("FINISH!")

}
