package main

import (
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "math/big"
)

func main() {
    //1、Int
    n, err := rand.Int(rand.Reader, big.NewInt(128))
    if err == nil {
        fmt.Println("rand.Int：", n, n.BitLen())
    }
    //2、Prime
    p, err := rand.Prime(rand.Reader, 5)
    if err == nil {
        fmt.Println("rand.Prime：", p)
    }
    //3、Read
    b := make([]byte, 32)
    m, err := rand.Read(b)
    if err == nil {
        fmt.Println("rand.Read：", b[:m])
        fmt.Println("rand.Read：", base64.URLEncoding.EncodeToString(b))
    }
    // rand.Int： 46 6
    // rand.Prime： 31
    // rand.Read： [209 194 83 164 157 235 15 88 17 246 237 180 44 139 179 59 31 121 210 249 188 36 189 217 112 36 168 56 212 177 162 80]
    // rand.Read： 0cJTpJ3rD1gR9u20LIuzOx950vm8JL3ZcCSoONSxolA=
}