package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "qwertyuiopasdfghjklzxcvbnm"
)

// init 将在第一次使用包时调用
func init() {
	// 设置rand的种子为当前时间的Unix时间戳
	rand.Seed(time.Now().UnixNano())
}

// RandomInt 生成范围在min和max间的随机Int64数
func RandomInt(min, max int64) int64 {
	// Int63n返回0到n-1的随机Int64数
	return min + rand.Int63n(max-min+1)
}

// RandomString 生成长度为n的随机字符串
func RandomString(n int) string {
	// strings.Builder 高效构建字符串
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner 生成随机Owner
func RandomOwner() string {
	return RandomString(6)
}

// RandomBalance 生成随机Balance
func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency 生成随机Currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CNY"}
	m := len(currencies)
	return currencies[rand.Intn(m)]
}

// RandomAmount 生成随机Amount
func RandomAmount() int64 {
	return RandomInt(-100, 100)
}
