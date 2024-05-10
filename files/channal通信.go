package files

import (
	"fmt"
	"strconv"
	"time"
)

//func main2() {
//	NotifMulti()
//
//	chok := make(chan os.Signal, 0)
//	signal.Notify(chok, os.Interrupt, os.Kill)
//	<-chok
//
//}

// 协程间通讯
func Communication() {
	// 可读写通道
	ch := make(chan int, 0)
	go communicationF1(ch)
	go communicationF2(ch)

}

// 输入只写通道
func communicationF1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

// 输入只读通道
func communicationF2(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

// 并发场景同步机制
func CommSync() {
	// 带缓冲
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

}

// select通知协程退出、多路复用
func NotifMulti() {
	intCh := make(chan int, 0)
	strCh := make(chan string, 0)
	structCh := make(chan struct{}, 0)
	go notifMultiF1(intCh)
	go notifMultiF2(strCh)
	go notifMultiF3(intCh, strCh, structCh)
	time.Sleep(5 * time.Second)
	// 关闭channal 发送一个零值
	close(structCh)
}

func notifMultiF1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}
func notifMultiF2(ch chan<- string) {
	for i := 0; i < 100; i++ {
		ch <- "数字" + strconv.Itoa(i)
	}
}
func notifMultiF3(intCh <-chan int, strCh <-chan string, structCh <-chan struct{}) {
	i := 0
	for {
		select {
		case intData := <-intCh:
			fmt.Println(intData)
		case strData := <-strCh:
			fmt.Println(strData)
		case struData := <-structCh:
			fmt.Println(struData)
			fmt.Println("结束  次数：", i)
			return
		}
		i++

	}
}
