package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "strings"
)

// 内网转发
func proxyHandler(w http.ResponseWriter, r *http.Request) {
    realServer := "http://localhost:10001"
    url, err := url.Parse(realServer)
    if err != nil {
        log.Println(err)
        return
    }
    proxy := httputil.NewSingleHostReverseProxy(url)
    proxy.ServeHTTP(w, r)
}

// api透传
func apiHandler(w http.ResponseWriter, r *http.Request) {

    err := r.ParseForm()
    if err != nil {
        log.Println(err)
        return
    }

    // post参数组装
    postData := ""
    formData := make(map[string]interface{})
    json.NewDecoder(r.Body).Decode(&formData)
    for key, value := range formData {
        postData += key + "=" + value.(string) + "&"
    }

    client := &http.Client{}
    req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader(postData))
    if err != nil {
        log.Println(err)
        return
    }
    req.Header = r.Header
    resp, err := client.Do(req)

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err)
        return
    }

    w.Write([]byte(body))
}

func main() {
    // http://localhost:10002/video
    http.HandleFunc("/video", proxyHandler)
    // curl http://localhost:10002/accept -X POST -H "Content-type:application/json" -d '{ "username": "admin", "password": "123456", "captcha": "", "report_date": "2019-08"}'
    http.HandleFunc("/accept", apiHandler)
    http.ListenAndServe(":10002", nil)
}
