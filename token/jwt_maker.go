package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const minSecretKeySize = 32

// JWTMaker JSON Web Token Maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker 创建一个JWTMaker，作为Maker进行管理token操作
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("长度不足，至少为 %d", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken 对指定用户创建时长为duration的token
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}
	// 创建jwtToken， 输入算法和payload
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// 生成token字符串
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err

}

// VerifyToken 检查输入令牌是否有效
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// 判断算法是否匹配
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		// jwt中将err使用ValidationError结构进行储存，需要进行获取判断具体类型
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
