package main

import (
	db "GoBank/db/sqlc"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func main() {
	testDB, err := pgx.Connect(context.Background(), "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	store := db.NewStore(testDB)

	a, err0 := store.TransferTx(context.Background(), db.TransferTxParams{
		FromAccountID: 1,
		ToAccountID:   2,
		Amount:        100,
	})

	fmt.Println(a, err0)
}

//
//func oneRoutine() {
//	mp := make(map[string]int)
//	list := []string{"A", "B", "C", "D"}
//	for i := 0; i < 20; i++ {
//		for _, item := range list {
//			_, ok := mp[item]
//			if !ok {
//				mp[item] = 0
//			}
//			mp[item]++
//		}
//	}
//	fmt.Println(mp)
//}
//
//type safeMap struct {
//	data   map[string]int
//	locker sync.Mutex
//}
//
//func manyRoutine() {
//	mp := safeMap{
//		data:   make(map[string]int),
//		locker: sync.Mutex{},
//	}
//	list := []string{"A", "B", "C", "D"}
//
//	wg := sync.WaitGroup{}
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			mp.locker.Lock()
//			defer mp.locker.Unlock()
//			for _, item := range list {
//				_, ok := mp.data[item]
//				if !ok {
//					mp.data[item] = 0
//				}
//				mp.data[item]++
//			}
//		}()
//
//	}
//	wg.Wait()
//	fmt.Println(mp)
//}
//
//// 协程安全的Map 针对读多写少场景
//func MapCase() {
//	mp := sync.Map{}
//	// 存
//	mp.Store("id", 1)
//	mp.Store("val", 1)
//
//	// 查，不存在则设置,存在则不设置并返回true  false不存在
//	fmt.Println(mp.LoadOrStore("id", 12))
//
//	// 查
//	fmt.Println(mp.Load("id"))
//}
