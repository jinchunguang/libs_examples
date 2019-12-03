/**
 * @Author : jinchunguang
 * @Date : 19-10-31 上午10:46
 * @Project : sty
 */
package main

import (
    "log"
    "os"
)

func main() {
    file, err := os.Create("./studygolang.txt")
    if err != nil {
        log.Println(err)
        // 错误处理，一般会阻止程序往下执行
        return
    }


    fileMode := getFileMode(file)
    log.Println("file mode:", fileMode)
    file.Chmod(fileMode | os.ModeSticky)

    log.Println("change after, file mode:", getFileMode(file))



}

func getFileMode(file *os.File) os.FileMode {
    fileInfo, err := file.Stat()
    if err != nil {
        log.Fatal("file stat error:", err)
    }

    return fileInfo.Mode()
}
