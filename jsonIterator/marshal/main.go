package main

import (
    "encoding/json"
    "fmt"
    "github.com/json-iterator/go"
)

func main() {
    type ColorGroup struct {
        ID     int
        Name   string
        Colors []string
    }
    group := ColorGroup{
        ID:     1,
        Name:   "Reds",
        Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
    }

    b, err := json.Marshal(group)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(string(b))

    var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
    b, err = json_iterator.Marshal(group)
    fmt.Println(string(b))
}