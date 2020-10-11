package main

import (
	"fmt"
	_ "fmt"
	_ "github.com/gin-gonic/gin"
	"os"
	_ "os"
	"time"
)

/**
 * https://juejin.im/post/6844904166985318414
 */

func main() {
	fmt.Println("go go go")
	fmt.Println(os.Getenv("GOPATH"))
	types()
	arrays()
	pointer()
	structs()
	slice()
	slice1()
	ranges()
	maps()
	mapDelete()
	//递归
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
	//斐波那契数列
	var ii int
	for ii = 0; ii < 10; ii++ {
		fmt.Printf("%d\t", fibonacci(ii))
	}

	typeChange()

	//goroutine
	go say("world")
	say("hello")

	//channel
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)

	//Go 遍历通道与关闭通道
	c1 := make(chan int, 10)
	go fibonacci1(cap(c1), c1)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func fibonacci1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

/**
* int int8 int16 int32 int64
* uint uint8 uint16 uint32 uint64 uintptr
* float32 float64 complex128 complex64
* bool byte rune string error
 */
func types() {
	type ProgramType string // 定义新类型 ProgramType
	var t1 ProgramType = "Golang"
	var t2 ProgramType = "C++"
	fmt.Println(t1, t2)

	type CompanyType string // 定义新类型 ProgramType
	var t3 ProgramType = "C++"
	//var c3, c3 CompanyType = "Google", "Tencent"
	fmt.Println(t1 + t3) // 同类型相加合法
	//fmt.Println(t3+c3)      // 不同类型相加非法

	var varint int = 66
	var varf float32 = float32(varint)                       // int 转换 float32
	fmt.Printf("%T %v %T %v \n", varint, varint, varf, varf) // %T输出值的类型  %v输出值

	var imVar1, imVar2, imVar3 int // 一口气声明了三个 int 类型的变量

	fmt.Println(imVar1 + imVar2 + imVar3) // 同类型相加合法
}

func arrays() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

func pointer() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	//nil 指针也称为空指针。
	//nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
	var ptr *int
	if ptr != nil /* ptr 不是空指针 */ {
		fmt.Printf("ptr1 的值为 : %x\n", ptr)
	}
	if ptr == nil /* ptr 是空指针 */ {
		fmt.Printf("ptr2 的值为 : %x\n", ptr)

	}
	fmt.Printf("ptr3 的值为 : %x\n", ptr)
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func structs() {
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})
	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)

	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)

	/* 打印 Book1 信息 */
	printBookPointer(&Book1)

	/* 打印 Book2 信息 */
	printBookPointer(&Book2)
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBookPointer(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func slice() {
	//https://www.runoob.com/go/go-slice.html

	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
}

func slice1() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func ranges() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func maps() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}

func mapDelete() {
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func typeChange() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}
