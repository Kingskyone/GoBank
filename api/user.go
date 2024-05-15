package api

import (
	db "GoBank/db/sqlc"
	"GoBank/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

// 接收request的参数   alphanum代表没有字符
type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// 返回的信息，去掉密码信息再返回
type createUserResponse struct {
	Username string `json:"username"`
	//HashedPassword   string             `json:"hashed_password"`
	FullName         string             `json:"full_name"`
	Email            string             `json:"email"`
	PasswordChangeAt pgtype.Timestamptz `json:"password_change_at"`
	CreatedAt        pgtype.Timestamptz `json:"created_at"`
}

// 实现创建用户  gin中的处理函数必须带有context输入
func (server Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	// 验证回调数据是否符合
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	// 插入数据
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		// 判断错误是否为外键错误、不可重复键错误
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := createUserResponse{
		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		PasswordChangeAt: user.PasswordChangeAt,
		CreatedAt:        user.CreatedAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
