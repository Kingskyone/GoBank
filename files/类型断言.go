package main

import "fmt"

func main() {
	// 编译器推断
	fmt.Println(max(3, 4))

	// 指定
	fmt.Println(max[int](3, 4))

	fmt.Println(getComparable(3, 3))

	//fxPrintLn([]int{})

	//a.Get()

}

// 泛型  一般情况不需要interface括起来，当出现指针*时避免出现与乘法混淆
func max[T interface{ int | float64 }](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func max2[T NumT](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

// 通过接口定义泛型
// ~ 该类型及其衍生类型
type NumT interface {
	// 支持的类型
	uint8 | int32 | float64 | ~int64
}

type NumT2[T comparable] interface {
	// 支持的类型
	any
	Get() T
}

//func listToMap[k comparable, T NumT2[k]](list []T) map[k]T {
//
//}

// comparable 内置的泛型类型 ，只支持 == 和 != 操作
func getComparable[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

// any 内置的泛型类型:任意类型
func getAny[T any](a T) {
	fmt.Println(a)
}

// 集合转列表  泛型
func mapToList[k comparable, T any](mp map[k]T) []T {
	list := make([]T, len(mp))
	var i int = 0
	for _, data := range mp {
		list[i] = data
		i++
	}
	return list
}

func fxPrintLn[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}

// 泛型结构体
type FX[T interface{ *int | *string }] struct {
	Name string
	Data T
}

// 泛型receiver
func (receiver FX[T]) GetData(data T) T {
	return receiver.Data
}
