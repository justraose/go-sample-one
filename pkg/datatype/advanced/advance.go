package advanced

import (
	"fmt"
	"time"
)

// type关键字介绍
// https://blog.csdn.net/hzwy23/article/details/79890778
//type有如下几种用法：
//1. 定义结构体
//2. 定义接口
//3. 类型定义 用在func上，例子如下： handle和exec
//4. 类型别名4
//5. 类型查询
// 		如果使用.(type)查询类型的变量不是interface{}类型，则在编译时会报如下错误
// 			cannot type switch on non-interface value a (type string)
// 		所以，使用type进行类型查询时，只能在switch中使用，且使用类型查询的变量类型必须是interface{}

// 定义一个接收一个字符串类型参数的函数类型
type handle func(str string)

// exec函数，接收handle类型的参数
func Exec(f handle) {
	f("hello")
}

func TestArray() {
	// [3]int{1, 2, 3}
	// var numbers = [...]int{1, 2, 3}
	type MyNumbers [5]int
	var numbers2 MyNumbers
	numbers2[0] = 2
	numbers2[3] = numbers2[0] - 3
	numbers2[1] = numbers2[2] + 5
	numbers2[4] = len(numbers2)
	sum := 11

	// “==”用于两个值的相等性判断
	fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))
}

/**
 * 切片可以起到自动扩容数组的效果
 */
func TestSlice() {
	//  可以看到，它们与数组的类型字面量的唯一不同是不包含代表其长度的信息。
	// 因此，不同长度的切片值是有可能属于同一个类型的。
	// 相对的，不同长度的数组值必定属于不同类型
	// 别名: type MySlice []int
	// 切片表达式的求值结果相当于以元素下界索引和元素上界索引作为依据从被操作对象上“切下”而形成的新值。注意，被“切下”的部分不包含元素上界索引指向的元素。
	// 一个切片值的容量即为它的第一个元素值在其底层数组中的索引值与该数组长度的差值的绝对值。
	var numbers3 = [5]int{1, 2, 3, 4, 5}
	var slice3 = numbers3[2:len(numbers3)]
	length := 3
	capacity := 3
	fmt.Printf("%v, %v\n", (length == len(slice3)), (capacity == cap(slice3)))

	// numbers3[1:4:4]
	//  这第三个正整数被称为容量上界索引。它的意义在于可以把作为结果的切片值的容量设置得更小。
	// 换句话说，它可以限制我们通过这个切片值对其底层数组中的更多元素的访问。

	// 为了防止切片访问超过长度: slice1 = slice1[:cap(slice1)]
	// 可以设置容量上界索引：var slice1 = numbers3[1:4:4]

	// 虽然切片值在上述方面受到了其容量的限制，但是我们却可以通过另外一种手段对其进行不受任何限制地扩展。
	// 这需要使用到内建函数append。append会对切片值进行扩展并返回一个新的切片值。使用方法如下：
	// slice1 = append(slice1, 6, 7)

	// 初始化有一定容量的切面
	var array12 = [5]int{}
	var slice12 = array12[1:4]
	fmt.Printf("len:%d, cap:%d \n", len(slice12), cap(slice12))

	// 切面扩容， 声明切片
	var slice11 []int
	fmt.Printf("len:%d, cap:%d \n", len(slice11), cap(slice11))
	slice11 = append(slice11, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("len:%d, cap:%d \n", len(slice11), cap(slice11))
}

func TestMap() {
	// map[K]T
	// 请注意，字典的键类型必须是可比较的，否则会引起错误。也就是说，它不能是切片、字典或函数类型。
	var myMap map[int]string
	// 上句是声明，若无下句初始化，则 panic: assignment to entry in nil map
	// 数组例外，声明即创建
	myMap = map[int]string{}
	//myMap
	myMap[4] = "测试"
	fmt.Println(myMap[4])
	fmt.Println(myMap)

	myMap2 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(myMap2[2])
	fmt.Println(myMap2)

	// 在不知道mm的确切值的情况下，我们无法得知mm[5]的求值结果意味着什么？它意味着5对应的值就是一个空字符串？
	// 还是说mm中根本就没有键为5的键值对？这无所判别。为了解决这个问题，Go语言为我们提供了另外一个写法，即：
	//e, ok := myMap2[5]

	// delete
	delete(myMap2, 1)
	fmt.Println(myMap2)
}

func TestChannel() {
	// 声明方法 chan T
	// 带缓冲 阻塞队列，读写都阻塞
	var myChan1 chan string
	// 与其它的数据类型不同，我们无法表示一个通道类型的值。因此，我们也无法用字面量来为通道类型的变量赋值。
	// 我们只能通过调用内建函数make来达到目的。
	// make函数也可以被用来初始化切片类型或字典类型的值
	// 初始化一个长度为5且元素类型为int的通道值
	myChan1 = make(chan string, 5)
	myChan1 <- "test chan"
	fmt.Println("channel value = " + <-myChan1)

	// 变量ok的值同样是bool类型的。它代表了通道值的状态，true代表通道值有效，而false则代表通道值已无效（或称已关闭）。
	// 如果通道内部还有数据，即使close(chan)，仍返回true，当数据都取出后，才返回false
	myChan1 <- "test two"
	close(myChan1)
	//close(myChan1) 对通道值的重复关闭会引发运行时恐慌
	value, ok := <-myChan1
	if ok == true {
		fmt.Println("channel有数据，数据为：" + value)
	} else {
		fmt.Println("通道myChan1已关闭")
	}

	// 不带缓冲 阻塞队列， 类似于java的SynchronousQueue
	type Sender chan<- int
	type Receiver <-chan int

	// 类似于java的SynchronousQueue
	var myChannel = make(chan int, 0)
	var number = 6
	go func() {
		var sender Sender = myChannel
		//time.Sleep(30*time.Second)
		sender <- number
		fmt.Println("Sent!")
	}()
	go func() {
		var receiver Receiver = myChannel
		//time.Sleep(10*time.Second)
		fmt.Println("Received!", <-receiver)
	}()
	// 让main函数执行结束的时间延迟1秒，
	// 以使上面两个代码块有机会被执行。
	time.Sleep(time.Second)
}

func TestFunc() {
	// 在Go语言中，函数是一等（first-class）类型。这意味着，我们可以把函数作为 值和参数 来传递和使用。
	// 函数类型的声明和使用
	var myF func(string, int) (string, int)
	myF = func(str string, num int) (string, int) {
		fmt.Println("I am 1 " + str + " " + string(num))
		return "hi", 2
	}

	var myF2 MyFunc
	myF2 = func(str string, num int) (string, int) {
		fmt.Println("I am 2 " + str + " " + string(num))
		return "hi", 3
	}

	// myF的类型等价于myF2
	exec(myF)
	exec(myF2)
}

// 函数类型声明
type MyFunc func(str string, num int) (string, int)

func exec(myfunc MyFunc) {
	s, i := myfunc("hello", 1)
	fmt.Println(s, " ", i)
}

type person struct {
	Name   string
	Gender string
	Age    int8
}

func (p *person) test() {
	fmt.Println("我是指针方法")
}

/**
封装
*/
func TestStruct() {
	// Struct
	// 声明 var x structName
	// 可以通过在结构体类型的声明中添加匿名字段（或称嵌入类型）来模仿继承。
	var p1 person = person{}
	fmt.Println(p1)
	p1.test()

	personSample := person{"a", "v", 1}
	fmt.Println(personSample)

	// 匿名结构体
	p := struct {
		Name   string
		Gender string
		Age    uint8
	}{"Robert", "Male", 33}
	fmt.Println(p)
}

func TestInterface() {

}
