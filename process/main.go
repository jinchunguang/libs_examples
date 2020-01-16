/**
 * @Author : jinchunguang
 * @Date : 19-10-31 下午12:21
 * @Project : sty
 */
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func main() {
	/*log.Println(os.Getpid())
	  log.Println(os.Getuid())
	  log.Println(os.Getgid())
	  time.Sleep(1000*time.Second)*/

	file, err := os.Open("./my_shadow.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("my_shadow:%s\n", data)

	fmt.Println(user.Current())
	fmt.Println(user.Lookup("jinchunguang"))
	fmt.Println(user.LookupId("0"))
	fmt.Println(os.Getwd())
	fmt.Println(os.Environ())
	fmt.Println(os.Getenv)
	fmt.Println(os.LookupEnv("SHELL"))
}
