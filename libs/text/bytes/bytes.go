package main

import (
	"bytes"
	"fmt"
)

func main() {

	fmt.Println("------------------Reader 类型---------------------------")
	x := []byte("你好，世界")

	r1 := bytes.NewReader(x)
	d1 := make([]byte, len(x))
	n, _ := r1.Read(d1)
	fmt.Println(n, string(d1))

	r2 := bytes.Reader{}
	r2.Reset(x)
	d2 := make([]byte, len(x))
	n, _ = r2.Read(d2)
	fmt.Println(n, string(d2))

	x = []byte("你好，世界")
	r1 = bytes.NewReader(x)

	ch, size, _ := r1.ReadRune()
	fmt.Println(size, string(ch))
	_ = r1.UnreadRune()
	ch, size, _ = r1.ReadRune()
	fmt.Println(size, string(ch))
	_ = r1.UnreadRune()

	by, _ := r1.ReadByte()
	fmt.Println(by)
	_ = r1.UnreadByte()
	by, _ = r1.ReadByte()
	fmt.Println(by)
	_ = r1.UnreadByte()

	d1 = make([]byte, 6)
	n, _ = r1.Read(d1)
	fmt.Println(n, string(d1))

	d2 = make([]byte, 6)
	n, _ = r1.ReadAt(d2, 0)
	fmt.Println(n, string(d2))

	w1 := &bytes.Buffer{}
	_, _ = r1.Seek(0, 0)
	_, _ = r1.WriteTo(w1)
	fmt.Println(w1.String())

	fmt.Println("------------------Buffer 类型---------------------------")
	a := bytes.NewBufferString("Hello World")
	b := bytes.NewBuffer([]byte("Hello World"))
	c := bytes.Buffer{}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
