package gapi

import (
	"GoBank/token"
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

// 在rpc中加入验证访问令牌
func (server *Server) authorizeUser(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("缺少元数据metadata")
	}
	value := md.Get(authorizationHeader)
	if len(value) == 0 {
		return nil, fmt.Errorf("缺少授权标头authorization")
	}

	authHeader := value[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("授权标头authorization格式错误")
	}
	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("不支持的授权类型")
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("访问令牌无效")
	}
	return payload, err
}
