package main

import (
	"fmt"
	"sync/atomic"
)

/*

http://wjp2013.github.io/go/go-atomic/
https://www.kancloud.cn/digest/batu-go/153537
https://docs.kilvn.com/The-Golang-Standard-Library-by-Example/chapter16/16.02.html

atomic包（sync/atomic）提供了底层的原子级内存操作。

共有五种操作：增减， 比较并交换， 载入， 存储，交换（T代表int32、int64、uint32、uint64、unitptr、pointer（没有增减操作））

	func LoadT(addr *T) (val T) ：读取；
	func StoreT(addr *T, val T)：写入；
	func AddT(addr *T, delta T) (new T)：增减，返回新的值；
	func SwapT(addr *T, new T) (old T)：交换，并返回旧值；
	func CompareAndSwapT(addr *T, old, new T) (swapped bool)：比较addr中保存的值是否与old相等，若相等则替换为新值，并返回true；否则，直接返回false。
	 */
func main() {

	// 增加或减少 Add
	// 函数会直接在传递的地址上进行修改操作，此外函数会返回修改之后的新值。
	// 需要注意的是当你处理 unint32 和 unint64 时，由于 delta 参数类型被限定，不能直接传输负数，所以需要利用二进制补码机制，其中 N 为需要减少的正整数值。
	var a uint32 = 10
	atomic.AddUint32(&a, 10)
	fmt.Println(a)
	atomic.AddUint32(&a, ^uint32(2 - 1)) // 等价于 b -= 10
	fmt.Println(a)

	// 比较并交换 CAS
	// CAS 的意思是判断内存中的某个值是否等于 old 值，如果是的话，则赋 new 值给这块内存。
	// 调用函数后，会先判断参数 addr 指向的被操作值与参数 old 的值是否相等
	// 仅当此判断得到肯定的结果之后，才会用参数 new 代表的新值替换掉原先的旧值，否则操作就会被忽略。

	// 读取和写入 Load and Store
	// 当我们要读取一个变量的时候，很有可能这个变量正在被写入，这时我们就很有可能读取到写到一半的数据，所以读取操作是需要一个原子行为的。如果有多个 CPU 往内存中一个数据块写入数据的时候，可能导致这个写入的数据不完整。
	// 在原子地存储某个值的过程中，任何 CPU 都不会进行针对同一个值的读或写操作。
	// 原子的值存储操作总会成功。

	var money int64 = 2000
	for i := 0; i < 1000; i++ {
		val := atomic.LoadInt64(&money)
		atomic.StoreInt64(&money, val+10)
	}
	fmt.Println(money)



}

func incNumAtomic(money *int64) {

}

