package api

import (
	db "GoBank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server 提供http请求
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer 创建一个新Server，设置所有http路由
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// 添加路由    传入一个或多个函数，若多个，最后一个为处理，其他为中间件
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server

}

// Start 在输入地址上运行http服务器，监听请求
func (server Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse 接收发生的错误并进行处理
func errorResponse(err error) gin.H {
	return gin.H{"error": err}
}
