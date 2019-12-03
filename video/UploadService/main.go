package main

import (
    "crypto/md5"
    "fmt"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "html/template"
    "io"
    "net/http"
    "os"
    "strconv"
    "time"
)


func upload(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("upload.tpl")
        t.Execute(w, token)
    }
    if r.Method == "POST" {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        fmt.Fprintf(w, "%v", handler.Header)
        f, err := os.OpenFile("./files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)
    }

}

func ossUpload(w http.ResponseWriter, r *http.Request) {

    endpoint := ""
    accessKeyID := ""
    accessKeySecret := ""
    bucketName := ""

    // 创建OSSClient实例。
    client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)
    }

    // 获取存储空间。
    bucket, err := client.Bucket(bucketName)
    if err != nil {
        fmt.Println("Error:", err)
    }

    // 上传本地文件
    err = bucket.PutObjectFromFile("video/oss.mp4", "./files/oss.mp4")
    if err != nil {
        fmt.Println("Error:", err)
    }

    w.Write([]byte("Success"))
}


func main() {
    http.HandleFunc("/ossUpload", ossUpload)
    http.HandleFunc("/upload", upload)
    http.ListenAndServe(":10003", nil)
}
