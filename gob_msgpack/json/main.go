package json

import (
	"encoding/json"
	"fmt"
)

type s struct {
	data map[string]interface{}
}

func jsonDemo() {

	var s1 = s{
		data: make(map[string]interface{}, 8),
	}
	s1.data["count"] = "1"
	ret, err := json.Marshal(s1.data)
	if err != nil {
		fmt.Println("marshal failed", err)
	}
	fmt.Printf("%#v\n", string(ret))

	var s2 = s{
		data: make(map[string]interface{}, 8),
	}
	err = json.Unmarshal(ret, &s2.data)
	if err != nil {
		fmt.Println("unmarshal failed", err)
	}
	fmt.Println(s2)
	for _, v := range s2.data {
		fmt.Printf("value:%v, type:%T\n", v, v)
	}
}

func main() {

	jsonDemo()
}
