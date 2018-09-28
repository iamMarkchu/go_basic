package main

import "fmt"

func main() {
	fmt.Println("go的基础类型介绍")

	/*
		基础类型
		bool
		int(rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64), complex128, complex64
		float32, float64
	   	string
	   	error
	   	array
	   	slice
	    map
	   	struct
		pointer
	*/

	var isSaved bool
	fmt.Println("是否保存:", isSaved)

	var intA int = 32
	var intB int32 = 57
	fmt.Println("A + B =", intA + int(intB))  // 不同类型的数值不能做运算

	var productPrice float64 = 799.99
	fmt.Println("商品的价格是: ", productPrice)

}
