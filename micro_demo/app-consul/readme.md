```
✗ go get -u -v github.com/lucas-clemente/quic-go/internal/handshake
go: finding golang.org/x/sys latest
go: finding golang.org/x/crypto latest
github.com/lucas-clemente/quic-go/internal/handshake
# github.com/lucas-clemente/quic-go/internal/handshake
../../../../../go/pkg/mod/github.com/lucas-clemente/quic-go@v0.14.1/internal/handshake/crypto_setup.go:433:40: not enough arguments in call to h.conn.GetSessionTicket
        have ()
        want ([]byte)

```

go 语言版本太高了
