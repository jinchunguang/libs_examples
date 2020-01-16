package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {

	httpclient := &http.Client{}
	data := `haha`

	endpoint := fmt.Sprintf("http://127.0.0.1:%d/%s?topic=%s", 4151, "pub", "test")
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(data)))
	resp, err := httpclient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	if resp.StatusCode != 200 {
		fmt.Printf("%s status code: %d", "pub", resp.StatusCode)
	}
	defer resp.Body.Close()

}
