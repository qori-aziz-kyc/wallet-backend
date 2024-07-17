package categoryhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/helper"
)

func (handler *categoryHandler) CreateHandler(ctx *gin.Context) {
	var categoryRequest request.CategoryRequest
	if err := ctx.ShouldBindJSON(&categoryRequest); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	statusCode, category, err := handler.usecase.Create(ctx.Request.Context(), categoryRequest)
	if err != nil {
		helper.NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	helper.NewSuccessResponse(ctx, statusCode, "create category success", category)
}
