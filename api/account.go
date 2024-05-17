package api

import (
	db "GoBank/db/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 接收request的参数，验证
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

// 实现创建账号   gin中的处理函数必须带有context输入
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	// 验证回调数据是否符合
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	}

	// 插入数据
	account, err := server.store.CreateAccount(ctx, arg)
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

	ctx.JSON(http.StatusOK, account)
}

// 使用uri绑定get请求里的参数
type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// 实现获取账号信息
func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	// 验证回调数据是否符合    ShouldBindUri处理get
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// 使用form绑定get请求里的参数
type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// 实现批量获取账号信息
func (server *Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest
	// 验证回调数据是否符合    ShouldBindQuery处理form
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListAccountParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	accounts, err := server.store.ListAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
