package api

import (
	"GoBank/token"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey  = "authorization"         //请求头
	authorizationTypeBearer = "bearer"                // 令牌类型 不记名令牌
	authorizationPayloadKey = "authorization_payload" //有效的payload使用该key在上下文进行存储
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	// 返回一个认证中间件函数
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("未提供 authorizationHeader")
			// 终止请求
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		// 验证格式
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("无效的 authorizationHeader 格式")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		// 验证格式是否支持
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf(" authorizationHeader 格式不支持 %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		// 验证令牌
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		// 将该令牌存储到ctx中
		ctx.Set(authorizationPayloadKey, payload)
		// 进行下一步
		ctx.Next()
	}
}
