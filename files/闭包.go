package files

import (
	"errors"
	"fmt"
	"log"
)

//func main8() {
//	// 可以叠加
//	//a := sumMiddleWare(sumMiddleWare(Sum))
//	//a := sumMiddleWare(Sum)
//	//fmt.Println(a(1, 2))
//	//
//	//fmt.Println(a.Accu(1, 2, 3, 4, 5))
//
//	//a := tool() // 声明的时候会带上闭包的数字
//	//b := tool()
//	//fmt.Println(a())
//	//fmt.Println(a())
//	//fmt.Println(a())
//	//fmt.Println(a())
//	//fmt.Println(b())
//	//fmt.Println(b())
//	//fmt.Println(b())
//	//fmt.Println(b())
//
//	xiec()
//	time.Sleep(time.Second * 10)
//}

func Sum(a, b int) (sum int, err error) {
	if a <= 0 && b <= 0 {
		err := errors.New("err")
		return 0, err
	}
	return a + b, nil
}

// 函数固定为类型
type SumFunc func(a, b int) (int, error)

func sumMiddleWare(in SumFunc) SumFunc {

	// 返回的函数为闭包函数
	return func(a, b int) (int, error) {
		log.Println("日志中间件输入", a, b)
		return in(a, b)
	}
}

// 累加  函数类型对象的方法
func (sum SumFunc) Accu(list ...int) (int, error) {
	sumR := 0
	var err error
	for _, dt := range list {
		sumR, err = sum(sumR, dt)
		if err != nil {
			return 0, err
		}
	}
	return sumR, nil
}

// 闭包斐波那契数列
func tool() func() int {
	var x0 = 0
	var x1 = 1
	var x2 = 0
	return func() int {
		x2 = x1 + x0
		x0 = x1
		x1 = x2
		return x2
	}
}

func xiec() {
	for i := 0; i < 10; i++ {
		// 传参会正确接收
		go func(num int) {
			fmt.Println(num)
		}(i)
		// 直接调用外部会闭包出错
		go func() {
			fmt.Println(i)
		}()
	}
}
