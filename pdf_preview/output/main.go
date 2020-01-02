package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/pv", zipHandler)
    http.HandleFunc("/", index)
    log.Println("Listening......")
    http.ListenAndServe(":8080", nil)
}
func index(rw http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html"))
    msg := "PDF预览测试"
    data := map[string]string{
        "msg": msg,
    }
    tmpl.Execute(rw, data)
}
func zipHandler(rw http.ResponseWriter, r *http.Request) {
    rw.Header().Set("Content-Type", "application/pdf")
    f, err := os.Open("./ali_java.pdf");
    defer f.Close()
    if err == nil {
        for {
            buf := make([]byte, 1024)
            _, err := f.Read(buf);
            if err == nil {
                rw.Write(buf)
                // 流控
                // time.Sleep(10000 * time.Microsecond)
            } else {
                return
            }

        }
    }

}
