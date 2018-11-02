package advanced

import "fmt"

// type关键字介绍
// https://blog.csdn.net/hzwy23/article/details/79890778
//type有如下几种用法：
//1. 定义结构体
//2. 定义接口
//3. 类型定义 用在func上，例子如下： handle和exec
//4. 类型别名
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

	// 切面扩容
	var slice11 []int
	fmt.Printf("len:%d, cap:%d \n", len(slice11), cap(slice11))
	slice11 = append(slice11, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("len:%d, cap:%d \n", len(slice11), cap(slice11))
}

func TestMap() {

}
