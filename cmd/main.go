package main

import (
    "fmt"
    "os"
)

func main() {
    pid, err := os.StartProcess("/bin/ps", []string{"ps", "-ef"},nil )

    if err != nil {
        fmt.Printf("Error %v starting process!", err)  //
        os.Exit(1)
    }

    fmt.Printf("The process id is %v", pid)
}
