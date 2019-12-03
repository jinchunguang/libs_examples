package main

import (
    "fmt"
    "github.com/json-iterator/go"
    "github.com/tidwall/gjson"
)

type ColorGroup struct {
    ID     int
    Name   string
    Colors []string
}

func main() {
    // group := ColorGroup{
    //     ID:     1,
    //     Name:   "Reds",
    //     Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
    // }
    // b, _ := jsoniter.Marshal(group)
    // fmt.Println(string(b))
    //
    // val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
    //
    // colors := jsoniter.Get(val, "Colors").ToString()
    // fmt.Println(colors)
    //
    // colors = jsoniter.Get(val, "Colors", 0).ToString()
    // fmt.Println(colors)

    tempJson := `{
    "name": {"first": "Tom", "last": "Anderson"},
    "age":37,
    "children": ["Sara","Alex","Jack"],
    "fav.movie": "Deer Hunter",
    "friends": [
        {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
        {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
        {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
    ]
    }`
    value := gjson.Get(tempJson, "friends")
    fmt.Println(value.String())

    colors := jsoniter.Get([]byte(tempJson), "friends").ToString()
    fmt.Println(colors)
}
