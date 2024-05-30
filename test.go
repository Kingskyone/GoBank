package main

//func main() {
//	fmt.Println(minFlips(8, 3, 5))
//	//fmt.Println(finMax([]int{5, 2, 1}, -1))
//}
//
//func minFlips(a int, b int, c int) int {
//	res := 0
//	aTail := 0
//	bTail := 0
//	cTail := 0
//	for a != 0 || b != 0 || c != 0 {
//		aTail = a % 2
//		bTail = b % 2
//		cTail = c % 2
//		a = a >> 1
//		b = b >> 1
//		c = c >> 1
//		if aTail|bTail != cTail {
//			if cTail == 1 {
//				res++
//			} else {
//				if aTail == 1 {
//					res++
//				}
//				if bTail == 1 {
//					res++
//				}
//			}
//		}
//	}
//	return res
//}
