package gapi

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	UserAgentHeader            = "user-agent"
	xForwardForHeader          = "x-forwarded-for"
)

// Metadata 存储从上下文中提取的元数据
type Metadata struct {
	UserAgent string
	ClientIp  string
}

// 解析上下文
func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// 使用grpcGateway的key获取
		if UserAgents := md.Get(grpcGatewayUserAgentHeader); len(UserAgents) > 0 {
			mtdt.UserAgent = UserAgents[0]
		}
		if ClientIps := md.Get(xForwardForHeader); len(ClientIps) > 0 {
			mtdt.ClientIp = ClientIps[0]
		}

		// 使用grpc获取
		if UserAgents := md.Get(UserAgentHeader); len(UserAgents) > 0 {
			mtdt.UserAgent = UserAgents[0]
		}
	}

	//grpc获取IP
	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIp = p.Addr.String()
	}

	return mtdt
}
