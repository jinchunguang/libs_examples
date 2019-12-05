package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func main() {

    /*
    os 包规定为所有操作系统实现的接口都是一致的。有一些某个系统特定的功能，需要使用 syscall 获取。实际上，os 依赖于 syscall。在实际编程中，我们应该总是优先使用 os 中提供的功能，而不是 syscall。
     */

    /*
    文件 I/O

    OpenFile 是一个更一般性的文件打开函数，大多数调用者都应用 Open 或 Create 代替本函数。


   位掩码参数 flag 用于指定文件的访问模式，可用的值在 os 中定义为常量（以下值并非所有操作系统都可用）：

   const (
       O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
       O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
       O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
       O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
       O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
       O_EXCL   int = syscall.O_EXCL   // 和 O_CREATE 配合使用，文件必须不存在
       O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步 I/O
       O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
   )


    注意：

    Write 调用成功并不能保证数据已经写入磁盘，因为内核会缓存磁盘的 I/O 操作。如果希望立刻将数据写入磁盘（一般场景不建议这么做，因为会影响性能），

    有两种办法：

   1. 打开文件时指定 `os.O_SYNC`；
   2. 调用 `File.Sync()` 方法。
   说明：File.Sync() 底层调用的是 fsync 系统调用，这会将数据和元数据都刷到磁盘；如果只想刷数据到磁盘（比如，文件大小没变，只是变了文件数据），需要自己封装，调用 fdatasync 系统调用。（syscall.Fdatasync）


    改变文件偏移量：Seek


     */

    /*
    截断文件


     */

    fmt.Println("--------------------------- 文件属性--------------------------------------")

    /*
    文件属性，也即文件元数据。在 Go 中，文件属性具体信息通过 os.FileInfo 接口获取。函数 Stat、Lstat 和 File.Stat 可以得到该接口的实例。这三个函数对应三个系统调用：stat、lstat 和 fstat。

    stat 会返回所命名文件的相关信息。
    lstat 与 stat 类似，区别在于如果文件是符号链接，那么所返回的信息针对的是符号链接自身（而非符号链接所指向的文件）。
    fstat 则会返回由某个打开文件描述符（Go 中则是当前打开文件 File）所指代文件的相关信息。

    FileInfo 接口如下：

    type FileInfo interface {
        Name() string       // 文件的名字（不含扩展名）
        Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
        Mode() FileMode     // 文件的模式位
        ModTime() time.Time // 文件的修改时间
        IsDir() bool        // 等价于 Mode().IsDir()
        Sys() interface{}   // 底层数据来源（可以返回 nil）
    }
    Sys() 底层数据的 C 语言 结构 statbuf 格式如下：

    struct stat {
        dev_t    st_dev;    // 设备 ID
        ino_t    st_ino;    // 文件 i 节点号
        mode_t    st_mode;    // 位掩码，文件类型和文件权限
        nlink_t    st_nlink;    // 硬链接数
        uid_t    st_uid;    // 文件属主，用户 ID
        gid_t    st_gid;    // 文件属组，组 ID
        dev_t    st_rdev;    // 如果针对设备 i 节点，则此字段包含主、辅 ID
        off_t    st_size;    // 常规文件，则是文件字节数；符号链接，则是链接所指路径名的长度，字节为单位；对于共享内存对象，则是对象大小
        blksize_t    st_blsize;    // 分配给文件的总块数，块大小为 512 字节
        blkcnt_t    st_blocks;    // 实际分配给文件的磁盘块数量
        time_t    st_atime;        // 对文件上次访问时间
        time_t    st_mtime;        // 对文件上次修改时间
        time_t    st_ctime;        // 文件状态发生改变的上次时间
    }
     */

    file, err := os.Create("studygolang.txt")
    if err != nil {
        log.Fatal("error:", err)
    }
    defer file.Close()

    fileMode := getFileMode(file)
    log.Println("file mode:", fileMode)
    file.Chmod(fileMode | os.ModeSticky)

    log.Println("change after, file mode:", getFileMode(file))

    fmt.Println("-----------------------------path/filepath — 兼容操作系统的文件路径操作----------------------------------")

    // 解析路径名字符串
    fmt.Println(filepath.Dir("/home/polaris/studygolang.go"))
    fmt.Println(filepath.Base("/home/polaris/studygolang.go"))
    fmt.Println(filepath.Ext("/home/polaris/studygolang.go"))

    // 相对路径和绝对路径
    // Abs 函数返回 path 代表的绝对路径 Rel 函数返回一个相对路径
    fmt.Println(filepath.Rel("/home/polaris/studygolang", "/home/polaris/studygolang/src/logic/topic.go"))
    fmt.Println(filepath.Rel("/home/polaris/studygolang", "/data/studygolang"))

    // 路径的切分和拼接
    // dir == /home/polaris/，file == studygolang
    fmt.Println(filepath.Split("/home/polaris/studygolang"))

    // dir == /home/polaris/studygolang/，file == ""
    fmt.Println(filepath.Split("/home/polaris/studygolang/"))

    /*
    Split 函数根据最后一个路径分隔符将路径 path 分隔为目录和文件名两部分（dir 和 file）。如果路径中没有路径分隔符，函数返回值 dir 为空字符串，file 等于 path；反之，如果路径中最后一个字符是 /，则 dir 等于 path，file 为空字符串。返回值满足 path == dir+file。dir 非空时，最后一个字符总是 /。
     */
    // dir == ""，file == studygolang
    fmt.Println(filepath.Split("studygolang"))

    // 文件路径匹配
    pth,_:=filepath.Glob("/home/jinchunguang/work/golang/libs_examples/libs/os/*.txt")
    fmt.Println(pth)

    // 遍历目录


}

func getFileMode(file *os.File) os.FileMode {
    fileInfo, err := file.Stat()
    if err != nil {
        log.Fatal("file stat error:", err)
    }

    return fileInfo.Mode()
}
