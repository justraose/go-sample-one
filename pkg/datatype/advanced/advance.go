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
	// 已关闭的通道不能再添加值，会发生运行时恐慌
	myChan1 <- "test two"
	close(myChan1)
	// close(myChan1) 对通道值的重复关闭会引发运行时恐慌
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

type Animal interface {
	Grow()
	Move(new string) (old string)
}

type Cat struct {
	little  string
	num     int16
	address string
}

// 值方法的 对象， 是内存副本
func (cat Cat) Grow() {
	fmt.Println("Grow func")
	cat.num++
}

func (cat *Cat) Move(new string) string {
	old := cat.address
	cat.address = new
	cat.num++
	return old
}

func TestInterface() {
	// 如果一个数据类型所拥有的方法集合中包含了某一个接口类型中的所有方法声明的实现，
	// 那么就可以说这个数据类型实现了那个接口类型。
	// “方法集合为其超集”

	// Go语言的类型转换规则定义了是否能够以及怎样可以把一个类型的值转换另一个类型的值。
	// 另一方面，所谓空接口类型即是不包含任何方法声明的接口类型，用interface{}表示，常简称为空接口。
	// 正因为空接口的定义，Go语言中的包含预定义的任何数据类型都可以被看做是空接口的实现。
	var cat Cat = Cat{"1", 2, "3"}
	var v interface{} = interface{}(&cat)
	//  在这之后，我们就可以在v上应用类型断言了，即：
	animal, ok := v.(Animal)
	//  类型断言表达式v.(Animal)的求值结果可以有两个。
	// 第一个结果是被转换后的那个目标类型（这里是Animal）的值，
	// 而第二个结果则是转换操作成功与否的标志。
	// 显然，ok代表了一个bool类型的值。它也是这里判定实现关系的重要依据。
	fmt.Println(animal, " ", ok)

	// question
	// 为什么只有*Person类型才实现了Animal接口?
}

func TestPointer() {
	// *Person是Person的指针类型
	// 表达式&p的求值结果是p的指针
	// 例如，*[3]string是数组类型[3]string的指针类型，而[3]string是*[3]string的基底类型
	cat := Cat{"test", 1, "xiamen"}

	// 之所以选择表达式person.Age成立，
	// 是因为如果Go语言发现person是指针并且指向的那个值有Age字段，
	// 那么就会把该表达式视为(*person).Age。
	// 其实，这时的person.Age正是(*person).Age的速记法。
	fmt.Println("===指针测试===")
	// 寻址(指针) &
	var catp *Cat = &cat
	catp.Grow()
	catp.Move("new catp")
	fmt.Println(catp)
	// 取值 *
	// 该处取值是生成新的副本
	var cata Cat = *catp
	cata.Grow()
	cata.Move("new cata")
	fmt.Println(cata)

	cat.Grow()
	fmt.Println(cat)

	// 多态
	// 要注意指针方法与值方法的区别
	// 拥有指针方法Grow和Move的指针类型*Person是接口类型Animal的实现类型，但是它的基底类型Person却不是
	// 隐藏规则： 一个指针类型拥有以它以及以它的基底类型为接收者类型的所有方法，而它的基底类型却只拥有以它本身为接收者类型的方法。
	animal, ok := interface{}(cat).(Animal)
	an, ok1 := interface{}(&cat).(Animal)
	fmt.Println(animal, ok, an, ok1)

	// 我们在基底类型的值上仍然可以调用它的指针方法。
	// 例如，若我们有一个Person类型的变量bp，则调用表达式bp.Grow()是合法的。
	// 这是因为，如果Go语言发现我们调用的Grow方法是bp的指针方法，那么它会把该调用表达式视为(&bp).Grow()。
	// 实际上，这时的bp.Grow()是(&bp).Grow()的速记法

	// golang的方法调用，参数是值传递，传递的是对象副本，需要操作原对象内存，则使用指针
	newCat := Cat{"zz", 0, "xx"}
	addCat(newCat)
	fmt.Println(newCat)

}

func addCat(cat Cat) {
	cat.num++
}
