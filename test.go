package main

//func main() {
//	//fmt.Println(minFlips(8, 3, 5))
//	fmt.Println(minimumDifference([]int{1, 2, 1, 2}, 2))
//}
//
//func minimumDifference(nums []int, k int) int {
//	ret := jdz(k, nums[0])
//	save := []int{}
//	for _, i := range nums {
//		if jdz(k, i) < ret {
//			ret = jdz(k, i)
//		}
//		for j := 0; j < len(save); j++ {
//			save[j] = save[j] & i
//			//fmt.Println(save[j], jdz(k, save[j]), ret)
//			if jdz(k, save[j]) < ret {
//				ret = jdz(k, save[j])
//			}
//		}
//		save = append(save, i)
//		save = removeDuplicates(save)
//		//fmt.Println(save)
//	}
//	return ret
//}
//func jdz(a, b int) int {
//	if a > b {
//		return a - b
//	}
//	return b - a
//}
//func removeDuplicates(list []int) []int {
//	seen := make(map[int]bool)
//	result := []int{}
//
//	for _, item := range list {
//		if !seen[item] {
//			seen[item] = true
//			result = append(result, item)
//		}
//	}
//	return result
//}
