package gapi

import (
	db "GoBank/db/sqlc"
	"GoBank/pb"
	"GoBank/token"
	"GoBank/util"
	"fmt"
)

// Server 提供grpc请求
type Server struct {
	// 向前兼容
	pb.UnimplementedGoBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer 创建一个新gRPC Server，
func NewServer(config util.Config, store db.Store) (*Server, error) {
	// 对称密钥
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("无法创建token")
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	return server, nil
}
