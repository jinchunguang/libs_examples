package main

import (
    "log"
    "os"
)

func main() {

    // log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger示例。

    /*
    New创建一个Logger对象。其中，参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。
     */
    logger := log.New(os.Stdout, "[report] ", log.Lshortfile|log.Ldate|log.Ltime)
    logger.Println("这是自定义的logger记录的日志。")

}
