package basic

import "fmt"

func TestInt() {
	// var
	var (
		num1 int8  = 1
		num2 int16 = 2
		num3 int32 = 3
	)

	var num4, num5 int64 = 4, 5
	var num6 uint64 = 6

	// 强制转换
	num7 := num6 + uint64(num5)

	fmt.Println(num1, num2, num3, num4, num5, num6, num7)

	// 声明一个整数类型变量并赋值
	var num8 int = -0x1000

	// 这里用到了字符串格式化函数。其中，%X用于以16进制显示整数类型值，%d用于以10进制显示整数类型值。
	fmt.Printf("16进制数 %X 表示的是 %d。\n", num8, num8)

	// byte与rune类型有一个共性，即：它们都属于别名类型。byte是uint8的别名类型，而rune则是int32的别名类型。
	// 一个rune类型的值即可表示一个Unicode字符

	// 声明一个rune类型变量并赋值
	var char1 rune = '赞'

	// 这里用到了字符串格式化函数。其中，%c用于显示rune类型值代表的字符。
	fmt.Printf("字符 '%c' 的Unicode代码点是 %s。\n", char1, "U+8D5E")

	divide()
}

func TestFloat() {
	// 浮点数类型有两个，即float32和float64
	// 3.7E-2表示浮点数0.037。又比如，3.7E+1表示浮点数37
	// 有一点需要注意，在Go语言里，浮点数的相关部分只能由10进制表示法表示，
	// 而不能由8进制表示法或16进制表示法表示。比如，03.7表示的一定是浮点数3.7

	// 可以在变量声明并赋值的语句中，省略变量的类型部分。
	// 不过别担心，Go语言可以推导出该变量的类型。
	var num1 = 1.22
	var num2 = 5.89E-4

	// 这里用到了字符串格式化函数。其中，%E用于以带指数部分的表示法显示浮点数类型值，%f用于以通常的方法显示浮点数类型值。
	fmt.Printf("浮点数 %E 表示的是 %f。\n", num2, num2)
	fmt.Printf("%E, %f\n", num1, num1)

	divide()
}

/**
 * 复数
 */
func TestComplex() {
	// 复数类型同样有两个，即complex64和complex128。
	// complex64类型的值会由两个float32类型的值分别表示复数的实数部分和虚数部分。
	// 而complex128类型的值会由两个float64类型的值分别表示复数的实数部分和虚数部分。

	// 复数类型的值一般由浮点数表示的实数部分、加号“+”、浮点数表示的虚数部分，以及小写字母“i”组成。
	// 比如，3.7E+1 + 5.98E-2i

	var num4 = 3.7E+1 + 5.98E-2i
	fmt.Printf("浮点数 %E 表示的是 %f。\n", num4, num4)

	divide()
}

func TestStr() {
	// 声明一个string类型变量并赋值
	var str1 string = "\\\""

	// 这里用到了字符串格式化函数。其中，%q用于显示字符串值的表象值并用双引号包裹。
	// %q 表示不进行转义的字符串
	fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 %s。\n", str1, str1)

	divide()
}

func divide() {
	fmt.Println("===================================================")
	fmt.Print("\n \n")
}
