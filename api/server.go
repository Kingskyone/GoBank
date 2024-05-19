package api

import (
	db "GoBank/db/sqlc"
	"GoBank/token"
	"GoBank/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server 提供http请求
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer 创建一个新Server，设置所有http路由
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

	//获取验证器引擎，转换为 *validator.Validate 类型
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证currency,可以在参数binding中使用
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

// 绑定路由单独封装为函数
func (server *Server) setupRouter() {
	router := gin.Default()
	// 添加路由    传入一个或多个函数，若多个，最后一个为处理，其他为中间件
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	// 设置路由组，用于绑定中间件
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start 在输入地址上运行http服务器，监听请求
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse 接收发生的错误并进行处理
func errorResponse(err error) gin.H {
	return gin.H{"error": err}
}
