package main

type 是 int

type intAlias = int

func main() {
	// 自定义类型myInt，基本类型是int
	type myInt int

	//将 int 类型取一个别名intAlias
	type intAlias = int
}
