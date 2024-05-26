package token

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"time"
)

type PasetoMaker struct {
	paseto      *paseto.V2
	symmericKey []byte //对称密钥
}

// NewPasetoMaker 创建一个PASETOmaker 作为Maker进行管理token操作
func NewPasetoMaker(symmericKey string) (Maker, error) {
	if len(symmericKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("密钥尺寸无效，需要%d大小", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto:      paseto.NewV2(),
		symmericKey: []byte(symmericKey),
	}
	return maker, nil
}

// CreateToken 对指定用户创建时长为duration的token
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}
	token, err := maker.paseto.Encrypt(maker.symmericKey, payload, nil)
	return token, payload, err
}

// VerifyToken 检查输入令牌是否有效
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmericKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	// 检查token是否有效
	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
