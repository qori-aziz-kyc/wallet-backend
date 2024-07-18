package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/helper"
)

func (handler *userHandler) LoginHandler(ctx *gin.Context) {
	var request request.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	statusCode, resp, err := handler.usecase.Login(ctx.Request.Context(), request)
	if err != nil {
		helper.NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	helper.NewSuccessResponse(ctx, statusCode, "login success", resp)
}
