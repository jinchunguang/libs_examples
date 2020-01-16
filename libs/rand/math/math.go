package math

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//不设置随机种子，每次运行结果都一样
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Float64())
	//设置随机种子
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Float64())
	//随机数切片
	fmt.Println(rand.Perm(5))

}
