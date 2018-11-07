package main

import (
	"fmt"
	"luzj.com/go-sample-one/pkg/datatype/advanced"
	"luzj.com/go-sample-one/pkg/datatype/basic"
)

// 常量只能被赋予基本数据类型的值本身
const NAME string = "basic type first"

func main() {
	fmt.Println("GoLang 开始 ：" + NAME)

	// 基础数据类型
	basic.TestInt()
	basic.TestFloat()
	basic.TestComplex()
	basic.TestStr()

	// type 类型定义
	advanced.Exec(func(str string) {
		fmt.Println("type 类型定义")
	})

	// 高级数据类型
	advanced.TestArray()
	advanced.TestSlice()
	advanced.TestMap()
	advanced.TestChannel()

	// 高级数据类型 >>> 2
	advanced.TestFunc()
	advanced.TestStruct()
	advanced.TestInterface()

	// 指针
	advanced.TestPointer()
}
