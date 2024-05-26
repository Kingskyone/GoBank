package main

//
//func main() {
//	fmt.Println(mostCompetitive([]int{3, 5, 2, 6}, 2))
//	//fmt.Println(finMax([]int{5, 2, 1}, -1))
//}
//
//func mostCompetitive(nums []int, k int) []int {
//	res := make([]int, k)
//	head := 0
//	for ind, i := range nums {
//		for true {
//			if (len(nums) - ind) == (k - head) {
//				res = append(res[:head], nums[ind:]...)
//				return res
//			}
//			if head == 0 {
//				res[head] = i
//				head++
//				break
//			} else if head != 0 && i >= res[head-1] {
//				if head < k {
//					res[head] = i
//					head++
//					break
//				} else {
//					break
//				}
//			} else {
//				head--
//				continue
//			}
//		}
//	}
//	return res
//}
