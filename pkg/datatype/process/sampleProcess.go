package process

import (
	"errors"
	"fmt"
	"io/ioutil"
	"luzj.com/go-sample-one/pkg/datatype/advanced"
	"math/rand"
	"os"
)

func TestIf() {
	var number int = 5
	if number += 4; 10 > number {
		number += 27
		number += 3
		fmt.Print(number)
	} else if 10 < number {
		number -= 2
		fmt.Print(number)
	}
	fmt.Println(number)
}

func TestSwitch() {
	// 1
	ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
	switch v := ia[rand.Intn(4)]; interface{}(v).(type) {
	case interface{}:
		fmt.Printf("Case A.")
	case byte:
		fmt.Printf("Case B.")
	default:
		fmt.Println("Unknown!")
	}

	// 2
	var name string = "Golang"
	// 省略若干条语句
	switch name {
	case "Golang":
		fmt.Println("A programming language from Google.")
	case "Rust":
		fmt.Println("A programming language from Mozilla.")
	default:
		fmt.Println("Unknown!")
	}

	// 3
	v := 11
	switch i := interface{}(v).(type) {
	case int, int8, int16, int32, int64:
		fmt.Printf("A signed integer: %d. The type is %T. \n", i, i)
	case uint, uint8, uint16, uint32, uint64:
		fmt.Printf("A unsigned integer: %d. The type is %T. \n", i, i)
	default:
		fmt.Println("Unknown!")
	}
}

func TestFor() {
	// map
	// 关键字range
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
	for k, v := range map1 {
		fmt.Printf("%d: %s\n", k, v)
	}

	// num
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Printf("\n")

	// string
	for i, v := range "Go语言" {
		fmt.Printf("%d: %c\n", i, v)
	}

	// array
	var myArr [5]int = [5]int{0, 1, 2, 3, 4}
	for i, num := range myArr {
		fmt.Println(i, ":", num)
	}
}

func TestSelect() {
	// select语句属于条件分支流程控制方法，不过它只能用于通道。
	// 请注意，如果一条select语句中不存在default case，
	// 并且在被执行时其中的所有case都不满足执行条件，那么它的执行将会被阻塞！

	ch1 := make(chan int, 1)
	ch1 <- 1
	ch2 := make(chan int, 1)
	ch2 <- 2
	// 只会进入一个case
	select {
	case e1 := <-ch1:
		fmt.Printf("1th case is selected. e1=%v.\n", e1)
	case e2 := <-ch2:
		fmt.Printf("2th case is selected. e2=%v.\n", e2)
	default:
		fmt.Println("No data!")
	}

	//
	ch3 := make(chan int, 100)
	// 随机插入 1 或 2
	for i := 0; i < 10; i++ {
		select {
		case ch3 <- 1:
			fmt.Printf("Sent %d\n", 1)
		case ch3 <- 2:
			fmt.Printf("Sent %d\n", 2)
		default:
			fmt.Println("Full channel!")
		}
	}

	//
	ch4 := make(chan int, 1)
	for i := 0; i < 4; i++ {
		select {
		case e, ok := <-ch4:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Println(e)
			close(ch4)
		default:
			fmt.Println("No Data!")
			ch4 <- 1
		}
	}
}

func TestDefer() {
	// 作用类似于java中的catch和finally

	// 1. 关键字defer：仅能被放置在函数或方法中,
	// 注意，当这条defer语句被执行的时候，其中的这条表达式语句并不会被立即执行。
	// 它的确切的执行时机是在其所属的函数（这里是readFile）的执行即将结束的那个时刻。
	// 也就是说，在readFile函数真正结束执行的前一刻，file.Close()才会被执行
	// 2. 更为关键的是，无论readFile函数正常地返回了结果还是由于在其执行期间有运行时恐慌发生而被剥夺了流程控制权，
	// 其中的file.Close()都会在该函数即将退出那一刻被执行。这就更进一步地保证了资源的及时释放。

	// defer关键字 栈这种数据结构存储，后声明先执行
	defer func() {
		fmt.Print(1)
	}()
	defer func() {
		fmt.Print(2)
	}()
	defer func() {
		fmt.Print(3)
	}()

	// 证明 f(i)先执行，然后defer延迟执行
	// 输出：1 2 3 4 40 30 20 10
	f := func(i int) int {
		fmt.Printf("%d ", i)
		return i * 10
	}
	for i := 1; i < 5; i++ {
		defer fmt.Printf("%d ", f(i))
	}
}

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

func TestError() (*advanced.Cat, error) {
	// 创造错误。没错，在很多时候，我们需要创造出错误（即error类型的值）并把它传递给上层程序。
	path := ""
	if path == "" {
		// 对于nil，一般通常  <指针类型>  和  <interface类型>  可以使用这样的返回值
		return nil, errors.New("The parameter is invalid!")
	}
	return nil, nil
}

func TestPanicAndRecover() {
	// panic可被意译为运行时恐慌。因为它只有在程序运行的时候才会被“抛出来”。
	// 实际上，内建函数panic和recover是天生的一对。前者用于产生运行时恐慌，而后者用于“恢复”它。

	// 不过要注意，recover函数必须要在defer语句中调用才有效。
	// 因为一旦有运行时恐慌发生，当前函数以及在调用栈上的所有代码都是失去对流程的控制权。
	// 只有defer语句携带的函数中的代码才可能在运行时恐慌迅速向调用栈上层蔓延时“拦截到”它。

	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Fatal error: %s\n", p)
		}

		fmt.Println("after recover...")
	}()

	fmt.Println("Enter TestPanicAndRecover")
	outerFunc()
	fmt.Println("Quit TestPanicAndRecover")
}

func innerFunc() {
	fmt.Println("Enter innerFunc")
	panic(errors.New("Occur a panic!"))
	fmt.Println("Quit innerFunc")
}

func outerFunc() {
	fmt.Println("Enter outerFunc")
	innerFunc()
	fmt.Println("Quit outerFunc")
}
