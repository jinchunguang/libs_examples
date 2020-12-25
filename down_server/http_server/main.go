package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const defaultPort = 7788

// 文件下载
func download(w http.ResponseWriter, r *http.Request) {

	// 解析参数，默认是不会解析的
	r.ParseForm()
	log.Println("Recv:", r.RemoteAddr)

	pwd, _ := os.Getwd()
	des := pwd + string(os.PathSeparator) + r.URL.Path[1:len(r.URL.Path)]
	log.Println("Des:", des)
	desStat, err := os.Stat(des)
	if err != nil {
		log.Println("File Not Exit", des)
		http.NotFoundHandler().ServeHTTP(w, r)
	}else if(desStat.IsDir()) {
		log.Println("File Is Dir", des)
		http.NotFoundHandler().ServeHTTP(w, r)
	}else{
		fileData, err := ioutil.ReadFile(des)
		if err != nil {
			log.Println("Read File Err:", err.Error())
		} else {
			log.Println("Send File:", des)
			w.Write(fileData)
		}
	}
}

func main() {

	port := flag.Int("p", defaultPort, "Set The Http Port")
	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Listen On Port:%d pwd:%s\n", *port, pwd)

	http.HandleFunc("/download", download)
	err = http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if nil != err {
		log.Fatalln("Get Dir Err", err.Error())
	}

}
