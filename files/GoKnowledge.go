package files //package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

import (
	"flag" //接收解析参数
	"fmt"  //IO包
	"time"
) //fmt 包实现了格式化 IO（输入/输出）的函数。

var name string

//	func init() {
//		//var name = flag.String("Name-命令参数名称","默认名称","对参数的说明")
//		flag.StringVar(&name, "Name-命令参数名称", "默认名称", "对参数的说明")
//	}
func GoKnowledgeMain() { //  重点：“ { ” 不能单独一行
	/* zhushi */
	flag.Parse()
	var testNumber string //变量定义，函数内不被使用的变量会报错，全局变量不会
	testNumber = "A line"
	var Aa = 1
	const Bb int = 2 //  const  常量   可以 const b = "abc" 省略类型
	var (            //多个声明，常用于全局变量
		a1 int
		a2 string
	)
	a2 += string(a1)
	Aa += Bb

	Cc, Dd := 3, 4 //不用声明，系统自动定义使用 " := "   但左侧必须为未声明的变量
	_, Cc = 5, 6   //    _  为只写变量，不可获得它的值
	Cc += Dd
	fmt.Println("hello,world,  " + testNumber + "\n") //Println输出会自动加\n,  fmt.Print()函数不会
	sd := &Cc                                         //  &指针取地址
	fmt.Println(sd)
	fmt.Println(name)
}

func Ts(a int, b string) (int, string) { //  传入值， return值
	var balance [10]float32 // 数组
	//var balance = [10]float32{1000.0, 2.0, 3.4, 7.0, 50.0}   初始定义的数组
	balance[0] = 1

	if balance[0] == 1 { //if 语句
		fmt.Println("as")
	} else {
		fmt.Println("sa")
	}
	switch balance[0] { //switch
	case 0:
		a = 1
	case 1:
		a = 2
		fallthrough // fallthrough表示继续执行下“ 一 ”个case
	default:
		a = 0
	}
	switch { //另一种switch
	case balance[0] == 0:
		a = 1
	case balance[0] == 1:
		a = 2
	default:
		a = 0
	}
	for ia := 0; ia < 10; ia++ { // for
		balance[ia] = 0
	}
	for true { // 类似while
		break // continue
	}

	var ip *float32 = &balance[0] // 指针
	fmt.Println(*ip)              // 指针取值

	type Books struct { // 结构体
		title   string
		author  string
		subject string
		book_id int
	}
	var bok Books      // 实例化结构体
	bok.title = "NMSL" // 取值

	var numbers = make([]int, 3, 5)                                 // 可变数组  参数：类型， 初始长度， 最大长度（可选）
	fmt.Printf("%d, %d, %v\n", len(numbers), cap(numbers), numbers) // cap最大长度   len当前长度   %v
	fmt.Println("numbers[1:4] ==", numbers[1:4])                    //切片
	numbers = append(numbers, 0, 1, 2)                              // append添加
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)         // 建一个更长的
	copy(numbers1, numbers)                                         /* 拷贝 numbers 的内容到 numbers1 */

	sum := 0
	for _, num := range numbers { //range可以遍历，需要两个来接收，第一个为第几个
		sum += num
	}

	//var cMap map[string]string /*创建集合 */
	cMap := make(map[string]string)
	cMap["tes"] = "TES"             // 赋值
	capital, ok := cMap["American"] // 可以获取第二个参数ok 为是否存在
	if ok {
		capital += "1"
	} else {
		capital += "2"
	}
	delete(cMap, "tes") //  删除一个集合项

	return a, b
}

type Phone interface { //总接口，有call函数
	call()
}

type NokiaPhone struct { //实例化结构体
}

func (nokiaPhone NokiaPhone) call() { //将结构体作为传参，绑定
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main2() {
	var phone Phone //实例化接口

	phone = new(NokiaPhone) // 接口联到一个结构体中
	phone.call()

	phone = new(IPhone)
	phone.call()

}

func sayYes0(Num int) {
	for ia := 0; ia < Num; ia++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(ia)
	}
}

func tongdao() {
	var v int = 1
	ch := make(chan int) //声明通道
	ch <- v              // 把 v 发送到通道 ch
	v = <-ch             // 从 ch 接收数据并把值赋给 v

	ch2 := make(chan int, 100) //有缓冲区，不需要必须等待接收
	ch2 <- v

	close(ch) //关闭通道
}

func yunxing() {
	fmt.Println("M")
	//不产生可执行文件直接运行：  go run *.go
	//产生: go build *.go  此方法需要清除  go clean
}

func TryPanic() {
	defer func() { //  * defer方法在return之后
		err := recover() //正常panic会中断程序，但若defer里面有recover则会恢复现场
		if err != nil {
			fmt.Println(err) // 此处为  message
		}
	}()
	getPanic()
}

func getPanic() {
	panic("message")
}
