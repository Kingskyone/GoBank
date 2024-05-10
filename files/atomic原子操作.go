package files

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main1() {
	AtomicCase2()
}

func AtomicCase() {
	var count int64 = 10
	atomic.StoreInt64(&count, 5)
	fmt.Println(atomic.LoadInt64(&count))
	atomic.AddInt64(&count, 10)
	fmt.Println(atomic.LoadInt64(&count))
	atomic.SwapInt64(&count, 1)
	fmt.Println(atomic.LoadInt64(&count))
	// 防止协程覆盖的情况 相同则替换
	atomic.CompareAndSwapInt64(&count, 1, 100)
}

// 计数器结构体
type atomicCounter struct {
	count int64
}

func (a *atomicCounter) Inc() {
	atomic.AddInt64(&a.count, 1)
}

func (a *atomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.count)
}

// 计数器 多少携程
func AtomicCase1() {
	var count int64 = 0
	// 一种方法  用锁保证
	//locker := sync.Mutex{}
	// 另一种，用atomic
	atom := atomicCounter{count: 0}
	wg := sync.WaitGroup{}
	for i := 0; i < 2001; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//locker.Lock()
			count += 1
			atom.Inc()
			//locker.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count, atom.Load())
}

// 统计列表中数字出现次数
func AtomicCase2() {
	list := []int{1, 2, 3, 4, 5, 6}
	// 定义一个原子值
	stoMp := atomic.Value{}
	mp := map[int]int{}
	// 存储
	stoMp.Store(&mp)

	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		atomicLabel:
			m := stoMp.Load().(*map[int]int)
			m1 := map[int]int{}
			for k, v := range *m {
				m1[k] = v
			}
			for _, item := range list {
				_, ok := m1[item]
				if !ok {
					m1[item] = 0
				}
				m1[item] += 1
			}
			swap := stoMp.CompareAndSwap(m, &m1)
			if !swap {
				goto atomicLabel
			}
		}()
	}
	wg.Wait()
	fmt.Println(stoMp.Load())

}
