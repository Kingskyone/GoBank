package api

import (
	db "GoBank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	//获取验证器引擎，转换为 *validator.Validate 类型
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证currency,可以在参数binding中使用
		v.RegisterValidation("currency", validCurrency)
	}

	// 添加路由    传入一个或多个函数，若多个，最后一个为处理，其他为中间件
	router.POST("/users", server.createUser)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

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
