package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

func main() {
    // httpclient := &http.Client{}
    // endpoint := fmt.Sprintf("http://127.0.0.1:%d/%s?topic=%s", port, method, topic)
    // req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
    // resp, err := httpclient.Do(req)
    // if err != nil {
    //     t.Fatalf(err.Error())
    //     return
    // }
    // if resp.StatusCode != 200 {
    //     t.Fatalf("%s status code: %d", method, resp.StatusCode)
    // }
    // resp.Body.Close()

    // 生成client 参数为默认
    // client := &http.Client{}
    //
    // // 生成要访问的url
    // url := "http://127.0.0.1:4151/pub?topic=test"
    //
    // // 提交请求
    // reqest, err := http.NewRequest("GET", url, nil)
    //
    // if err != nil {
    //     panic(err)
    // }

    //
    // // 处理返回结果
    // response, _ := client.Do(reqest)
    //
    // // 将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
    // stdout := os.Stdout
    // _, err = io.Copy(stdout, response.Body)
    //
    // // 返回的状态码
    // status := response.StatusCode
    //
    // fmt.Println(status)

    url := "http://127.0.0.1:4151/pub?topic=test"
    data := `haha`
    request, _ := http.NewRequest("POST", url, strings.NewReader(data))

    // post数据并接收http响应
    resp, err := http.DefaultClient.Do(request)
    defer resp.Body.Close()
    if err != nil {
        fmt.Printf("post data error:%v\n", err)
        return
    }

    fmt.Println("post a data successful.")
    respBody, _ := ioutil.ReadAll(resp.Body)
    fmt.Printf("response data:%v\n", string(respBody))

}
