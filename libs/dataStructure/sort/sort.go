package main

import (
	"fmt"
	"sort"
)

// 学生成绩结构体
type StuScore struct {
	name  string // 姓名
	score int    // 成绩
}

type StuScores []StuScore

// Len()
func (s StuScores) Len() int {
	return len(s)
}

// Less(): 成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

// Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	fmt.Println("-----------------------数据集合排序----------------------------")
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	// 打印未排序的 stus 数据
	fmt.Println("Default:\n\t", stus)
	// StuScores 已经实现了 sort.Interface 接口 , 所以可以调用 Sort 函数进行排序
	sort.Sort(stus)
	// 判断是否已经排好顺序，将会打印 true
	fmt.Println("IS Sorted?\n\t", sort.IsSorted(stus))
	// 打印排序后的 stus 数据
	fmt.Println("Sorted:\n\t", stus)
	// 逆向排序
	sort.Sort(sort.Reverse(stus))
	fmt.Println("Sorted Reverse:\n\t", stus)

	fmt.Println("-----------------------二分查找----------------------------")
	/*
	   Search() 函数一个常用的使用方式是搜索元素 x 是否在已经升序排好的切片 s 中
	*/
	x := 11
	s := []int{3, 6, 8, 11, 45} // 注意已经升序排序
	pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
	if pos < len(s) && s[pos] == x {
		fmt.Println(x, " 在 s 中的位置为：", pos)
	} else {
		fmt.Println("s 不包含元素 ", x)
	}

	fmt.Println("----------------------- 猜数字的小程序 ----------------------------")
	// GuessingGame()

	fmt.Println("-----------------------sort包已经支持的内部数据类型排序----------------------------")

	/*
	   IntSlice 类型及[]int 排序

	   sort包原生支持[]int、[]float64 和[]string 三种内建数据类型切片的排序操作，即不必我们自己实现相关的 Len()、Less() 和 Swap() 方法。
	*/
	s = []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s) // 对[]int 切片排序更常使用 sort.Ints()，而不是直接使用 IntSlice 类型
	fmt.Println(s)

	// 使用降序排序
	s = []int{5, 2, 6, 3, 1, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)

	// SearchInts() 的使用条件为：切片 a 已经升序排序 以下是一个错误使用的例子

	s = []int{5, 2, 6, 3, 1, 4}        // 未排序的切片数据
	fmt.Println(sort.SearchInts(s, 2)) // 将会输出 0 而不是 1

	/*

	   Float64Slice 类型及[]float64 排序

	    Float64Slice 类型定义的 Less 方法中，有一个内部函数 isNaN()。 isNaN() 与math包中 IsNaN() 实现完全相同，sort包之所以不使用 math.IsNaN()，完全是基于包依赖性的考虑，应当看到，sort包的实现不依赖与其他任何包。

	*/

	/*
	   StringSlice 类型及[]string 排序
	   两个 string 对象之间的大小比较是基于“字典序”的
	*/

	fmt.Println("----------------------- []interface 排序与查找 ----------------------------")

	/*
		    为什么 sort 包可以完成 []int 的排序，而不能完成 []struct 的排序

		    因为排序涉及到比较两个变量的值，而 struct 可能包含多个属性，程序并不知道你想以哪一个属性或哪几个属性作为衡量大小的标准。如果你能帮助程序完成比较，并将结果返回， sort 包内的方法就可以完成排序，判断，查找等。sort 包提供了以下函数：

		    func Slice(slice interface{}, less func(i, j int) bool)
		    func SliceStable(slice interface{}, less func(i, j int) bool)
		    func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool
		    func Search(n int, f func(int) bool) int

		通过函数签名可以看到，排序相关的三个函数都接收 []interface，并且需要传入一个比较函数，用于为程序比较两个变量的大小，因为函数签名和作用域的原因，这个函数只能是 匿名函数。

	*/

	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	// sort.Slice []interface 的排序
	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age }) // 按年龄升序排序
	fmt.Println("Sort by age:", people)

	// sort.SliceStable  []interface 的稳定排序
	sort.SliceStable(people, func(i, j int) bool { return people[i].Age > people[j].Age }) // 按年龄降序排序
	fmt.Println("Sort by age:", people)

	// sort.SliceIsSorted  []interface 是否为有序
	sort.Slice(people, func(i, j int) bool { return people[i].Age > people[j].Age }) // 按年龄降序排序
	fmt.Println("Sort by age:", people)
	fmt.Println("Sorted:", sort.SliceIsSorted(people, func(i, j int) bool { return people[i].Age < people[j].Age }))

	// sort.Search  []interface 是否存在指定元素
	a := []int{2, 3, 4, 200, 100, 21, 234, 56}
	x = 21

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })           // 升序排序
	index := sort.Search(len(a), func(i int) bool { return a[i] >= x }) // 查找元素

	if index < len(a) && a[index] == x {
		fmt.Printf("found %d at index %d in %v\n", x, index, a)
	} else {
		fmt.Printf("%d not found in %v,index:%d\n", x, a, index)
	}
	fmt.Println("-----------------------  ----------------------------")

}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

/*
   该包实现了四种基本排序算法：
   插入排序、归并排序、堆排序和快速排序。

   但是这四种排序方法是不公开的，它们只被用于 sort 包内部使用。所以在对数据集合排序时不必考虑应当选择哪一种排序方法

   只要实现了 sort.Interface 定义的三个方法：
   获取数据集合长度的 Len() 方法、
   比较两个元素大小的 Less() 方法
   交换两个元素位置的 Swap() 方法

   就可以顺利对数据集合进行排序。

   sort 包会根据实际数据自动选择高效的排序算法。

   为了方便对常用数据类型的操作，sort 包提供了对[]int 切片、[]float64 切片和[]string 切片完整支持，主要包括：

   type Interface interface {
           // 获取数据集合元素个数
           Len() int
           // 如果 i 索引的数据小于 j 索引的数据，返回 true，且不会调用下面的 Swap()，即数据升序排序。
           Less(i, j int) bool
           // 交换 i 和 j 索引的两个元素的位置
           Swap(i, j int)
   }


*/
