
# websocket

import "github.com/gorilla/websocket"

包websocket实现RFC 6455中定义的WebSocket协议。

## 概观
Conn类型表示WebSocket连接。服务器应用程序从HTTP请求处理程序调用Upgrader.Upgrade方法以获取* Conn:

    var 
        upgrader = websocket.Upgrader { ReadBufferSize:1024,
        WriteBufferSize:1024,
    } 
    
    func handler（w http.ResponseWriter,r * http.Request）{ 
        conn,err:= upgrader.Upgrade （w,r,nil）
        if err！= nil { 
            log.Println（err）
            return 
        } 
        ...使用conn发送和接收消息。
    }

调用连接的WriteMessage和ReadMessage方法以一片字节发送和接收消息。这段代码显示了如何使用以下方法回显消息:

    for { 
        messageType,p,err:= conn.ReadMessage()
        if err！= nil { 
            log.Println（err）
            return 
        } 
        if err:= conn.WriteMessage（messageType,p）; err！= nil { 
            log.Println（err）
            return 
        } 
    }
    
在上面的代码片段中,p是一个[]字节,messageType是一个值为websocket.BinaryMessage或websocket.TextMessage的int。

应用程序还可以使用io.WriteCloser和io.Reader接口发送和接收消息。要发送消息,请调用连接NextWriter方法以获取io.WriteCloser,将消息写入writer并在完成后关闭writer。要接收消息,请调用连接NextReader方法以获取io.Reader并读取,直到返回io.EOF。此代码段显示了如何使用NextWriter和NextReader方法回显消息:

    for { 
        messageType,r,err:= conn.NextReader()
        if err！= nil { 
            return 
        } 
        w,err:= conn.NextWriter（messageType）
        if err！= nil { 
            return err 
        } 
        if _,err:= io。复制（w,r）; err！= nil { 
            return err 
        } 
        if err:= w.Close(); 错误！= nil { 
            return err 
        } 
    }
    
## 数据消息

WebSocket协议区分文本和二进制数据消息。

文本消息被解释为UTF-8编码的文本。二进制消息的解释留给应用程序。

使用`TextMessage`和`BinaryMessage`整数常量来标识两种数据消息类型。

ReadMessage和NextReader方法返回接收消息的类型。 WriteMessage和NextWriter方法的messageType参数指定已发送消息的类型。


应用程序有责任确保文本消息是有效的UTF-8编码文本。

## 控制消息

WebSocket协议定义了三种类型的控制消息:close,ping和pong。调用连接WriteControl,WriteMessage或NextWriter方法向对等方发送控制消息。

Connections通过调用使用SetCloseHandler方法设置的处理函数并从NextReader,ReadMessage或消息Read方法返回* CloseError来处理收到的关闭消息。默认关闭处理程序向对等方发送关闭消息。

Connections通过使用SetPingHandler方法调用处理函数集来处理收到的ping消息。默认的ping处理程序向对等方发送pong消息。

Connections通过调用使用SetPongHandler方法设置的处理函数来处理收到的pong消息。默认的pong处理程序什么都不做。如果应用程序发送ping消息,则应用程序应设置pong处理程序以接收相应的pong。

从NextReader,ReadMessage和消息读取器读取方法调用控制消息处理函数。当处理程序写入连接时,默认的关闭和ping处理程序可以在短时间内阻止这些方法。

应用程序必须读取连接以处理从对等方发送的close,ping和pong消息。 如果应用程序对来自对等方的消息不感兴趣,那么应用程序应该启动goroutine来读取和丢弃来自对等方的消息。一个简单的例子是:

    func readLoop（c * websocket.Conn）{ 
        for { 
            if _,_,err:= c.NextReader(); err！= nil { 
                c.Close()
                break 
            } 
        } 
    }

## 并发

Connections支持一个并发读者和一个并发编写器。

应用程序负责确保:
不超过一个goroutine同时调用write方法（NextWriter,SetWriteDeadline,WriteMessage,WriteJSON,EnableWriteCompression,SetCompressionLevel）同时发生。
不超过一个goroutine调用read方法（NextReader,SetReadDeadline,ReadMessage,ReadJSON,SetPongHandler） ,SetPingHandler）同时发生。

Close和WriteControl方法可以与所有其他方法同时调用。

## 起源考虑因素

Web浏览器允许Javascript应用程序打开与任何主机的WebSocket连接。服务器可以使用浏览器发送的Origin请求标头来强制执行原始策略。

Upgrader调用`CheckOrigin`字段中指定的函数来检查原点。如果CheckOrigin函数返回false,则Upgrade方法使HTTP状态`403`的WebSocket握手失败。

如果CheckOrigin字段为nil,则Upgrader使用安全默认值:如果Origin请求标头存在且Origin主机不等于Host请求标头,则握手失败。

不推荐使用的包级别升级功能不执行原始检查。应用程序负责在调用Upgrade函数之前检查Origin标头。

## 缓冲区

连接缓冲网络输入和输出,以减少读取或写入消息时的系统调用次数。

写缓冲区也用于构造WebSocket帧。有关消息框架的讨论,请参阅RFC 6455,第5节。每次将写入缓冲区刷新到网络时,都会将WebSocket帧头写入网络。减小写缓冲区的大小会增加连接上的成帧开销量。

缓冲区大小（以字节为单位）由Dialer和Upgrader中的`ReadBufferSize`和`WriteBufferSize`字段指定。
当缓冲区大小字段设置为零时,Dialer使用`默认大小4096`。当缓冲区大小字段设置为零时,Upgrader重用HTTP服务器创建的缓冲区。在撰写本文时,`HTTP服务器缓冲区的大小为4096`。

缓冲区大小不限制可由连接读取或写入的消息的大小。

默认情况下,缓冲区在连接的生命周期内保留。如果设置了Dialer或Upgrader WriteBufferPool字段,则只有在写入消息时,连接才会保留写入缓冲区。

应用程序应调整缓冲区大小以平衡内存使用和性能。增加缓冲区大小会占用更多内存,但可以减少读取或写入网络的系统调用次数。在写入的情况下,增加缓冲区大小可以减少写入网络的帧头的数量。

设置缓冲区参数的一些准则是:

- 将缓冲区大小限制为最大预期消息大小。大于最大消息的缓冲区不提供任何好处。

- 根据消息大小的分布,将缓冲区大小设置为小于最大预期消息大小的值可以大大减少内存使用,同时对性能影响很小。下面是一个示例:如果99％的消息小于256字节且最大消息大小为512字节,则256字节的缓冲区大小将导致1.01个系统调用比512字节的缓冲区大小多。节省的内存为50％。

- 当应用程序具有适当数量的大量连接写入时,写缓冲池很有用。当缓冲池合并时,较大的缓冲区大小对总内存使用的影响较小,并具有减少系统调用和帧开销的好处。

## 压缩实验

此程序包以有限的容量通过实验支持 每个邮件压缩扩展（RFC 7692）。在Dialer或Upgrader中将EnableCompression选项设置为true将尝试协商每个消息deflate支持。

    var upgrader = websocket.Upgrader { 
        EnableCompression:true,
    }

如果与连接的对等方成功协商压缩,则将自动解压缩以压缩形式接收的任何消息。所有Read方法都将返回未压缩的字节。

通过调用相应的Conn方法,可以启用或禁用写入连接的消息的每个消息压缩:

conn.EnableWriteCompression（假）
目前,此程序包不支持使用“上下文接管”进行压缩。这意味着必须单独压缩和解压缩消息,而不保留消息中的滑动窗口或字典状态。有关更多详细信息,请参阅RFC 7692。

压缩的使用是实验性的,可能导致性能下降。
