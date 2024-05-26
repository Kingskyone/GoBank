package token

import "time"

// Maker 管理token
type Maker interface {
	// CreateToken 对指定用户创建时长为duration的token
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerifyToken 检查输入令牌是否有效
	VerifyToken(token string) (*Payload, error)
}
