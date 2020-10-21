package web

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func BugFix() error {
	var GOPATH string

	for _, x := range os.Environ() {
		if strings.Contains(x, "GOPATH=") {
			GOPATH = strings.Split(x, "GOPATH=")[1]
			fmt.Println(GOPATH)

			if err := os.MkdirAll(GOPATH+"/pkg/mod/github.com/micro/go-plugins/",0777); err != nil {
				log.Println("创建文件失败，", err)
				return err
			}else {
				//将 这个文件-->https://github.com/asim/go-plugins/archive/wrapper/validator/v2.9.1.zip
				//下载下来，解压更名为v2@v2.0.0  然后将srcPath更新为其路径
				srcPath := ""
				if srcPath == "" {
					log.Fatal("srcPath 没更新")
				}
				if err := os.Rename(srcPath, GOPATH+"/pkg/mod/github.com/micro/go-plugins/"); err != nil {
					log.Println("移动资源文件失败，", err)
					return err
				}
			}
			return nil
		}
	}


	return nil
}
