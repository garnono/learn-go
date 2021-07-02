package basic

import "fmt"

func TestArray() {
	PrintStart("array")

	// 特点：类型相同，长度固定，值可变

	// 声明：var variable_name [SIZE] variable_type
	var arr1 [3]int
	printArrayInt(arr1, 1)

	// 默认值不是nil
	//if arr1 == nil {
	//	fmt.Println("arr 1 is nil")
	//}

	// 赋值
	arr1[0] = 1
	printArrayInt(arr1, 1)

	// 声明 & 赋值
	var arr2 = [3]int{1, 2, 3}
	printArrayInt(arr2, 2)

	// 声明 & 赋值：省略长度
	var arr3 = []int{1, 2, 3}
	// error: cannot use arr3 (type []int) as type [3]int in argument to printArrayInt
	//printArrayInt(arr3, 3)
	fmt.Println("arr 3", "value:", arr3, "len:", len(arr3), "cap:", cap(arr3))

	//声明 & 赋值：省略长度 & 内部
	arr4 := []int{1, 2, 3}
	fmt.Println("arr 4", "value:", arr4, "len:", len(arr4), "cap:", cap(arr4))

	arr5 := [3]int{1, 2, 3}
	printArrayInt(arr5, 5)

	// 注意：
}

func printArrayInt(arr [3]int, num int) {
	fmt.Println("arr", num, "value:", arr, "len:", len(arr), "cap:", cap(arr))
}
