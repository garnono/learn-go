package learn

import "fmt"

func TestSlice() {
	PrintStart("slice")

	// 特点：动态数组——长度可变，内容可变
	// 声明：var identifier []type
	var slice1 []int
	printSlice(slice1, 1)

	// 默认值
	if slice1 == nil {
		fmt.Println("slice 1 is nil")
	}

	// 声明 & 赋值：等号左侧的类型可省略
	var slice2 []int = make([]int, 3)
	if slice2 == nil {
		fmt.Println("slice 2 is nil")
	}
	printSlice(slice2, 2)

	slice3 := make([]int, 3, 6)
	printSlice(slice3, 3)

	slice4 := []int{1, 2, 3}
	printSlice(slice4, 4)

	// 数组的声明和slice的相同了
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(arr, 0)

	slice5 := arr[:]
	printSlice(slice5, 5)

	// 第一个参数会影响cap
	slice6 := arr[1:5]
	printSlice(slice6, 6)

	// append & copy
	slice6 = append(slice6, 0, 1, 3)
	printSlice(slice6, 6)

	copy(slice6, slice2)
	printSlice(slice6, 6)

	// 注意：
}

func printSlice(slice []int, num int) {
	fmt.Println("slice", num, ":", slice, "len:", len(slice), "cap:", cap(slice))
}
