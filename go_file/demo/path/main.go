/**
 * @Author : jinchunguang
 * @Date : 19-10-31 上午10:59
 * @Project : sty
 */
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	fmt.Println(filepath.Dir("/home/polaris/studygolang/src/logic/topic.go"))
	fmt.Println(filepath.Base("/home/polaris/studygolang/src/logic/topic.go"))

	fmt.Println(filepath.Abs("./main.go"))
	fmt.Println(filepath.IsAbs("/home/polaris/studygolang/src/logic/topic.go"))

	fmt.Println(os.Getwd())
	//     fmt.Println(filepath.Rel("/home/polaris/studygolang", "/home/polaris/studygolang/src/logic/topic.go"))
	//     fmt.Println(filepath.Rel("/home/polaris/studygolang", "/data/studygolang"))

	fmt.Println(filepath.Split("/home/polaris/studygolang"))

	// fmt.Println(filepath.EvalSymlinks("symlink/studygolang.txt"))
	// fmt.Println(os.Readlink("symlink/studygolang.txt"))

	/*  files,_:=filepath.Glob("/home/jinchunguang/data/jiayunhui/*")
	    for k,v := range files {
	        fmt.Println(k,v) //prints 0, 1, 2
	    }*/

	log.Println(filepath.Walk("/home/jinchunguang/data/jiayunhui/", func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			log.Println(info.Name())
		} else {
			log.Println(path)
		}
		return nil
	}))

}
