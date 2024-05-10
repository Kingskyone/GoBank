package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ContextCase2()

	chok := make(chan os.Signal, 0)
	signal.Notify(chok, os.Interrupt, os.Kill)
	<-chok
}

// context 上下文信号传递

func ContextCase() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "desc", "ContextCase")
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	done := make(chan struct{})
	go f1(done)
	go f1(done)
	go f1(done)
	time.Sleep(time.Second)
	close(done)
	time.Sleep(time.Second)
}

func f1(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

func sum(a, b int) int {
	return a + b
}

func multi(a, b int) int {
	time.Sleep(5 * time.Second)
	return a * b
}

func ContextCase2() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "desc", "ContextCase2")
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	data := [][]int{{1, 2}, {3, 2}}

	ch := make(chan []int)
	go caculate(ctx, ch)
	ch <- data[0]
	ch <- data[1]

}

func caculate(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			ctx := context.WithValue(ctx, "desc", "caculate")

			ch := make(chan []int)
			go sumContext(ctx, ch)
			ch <- item

			ch2 := make(chan []int)
			go multiContext(ctx, ch2)
			ch2 <- item

			fmt.Println(item)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Println("caculate   Done    ", desc, "   ", ctx.Err())
			return
		}
	}
}

func sumContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			a, b := item[0], item[1]
			res := sum(a, b)
			fmt.Println(a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Println("sumContext   Done    ", desc, "   ", ctx.Err())
			return
		}
	}
}

func multiContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			a, b := item[0], item[1]
			res := multi(a, b)
			fmt.Println(a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Println("multiContext   Done    ", desc, "   ", ctx.Err())
			return
		}
	}
}
