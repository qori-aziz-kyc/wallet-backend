package recordhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qori-aziz-kyc/wallet-backend/internal/constants"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/helper"
)

func (handler *recordHandler) UpdateHandler(ctx *gin.Context) {
	var req request.RecordCreateRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	uid, err := strconv.Atoi(ctx.GetString(constants.CtxAuthenticatedUserKey))
	if err != nil {
		helper.NewErrorResponse(ctx, 400, err.Error())
		return
	}
	req.UserID = uid

	statusCode, record, err := handler.usecase.Update(ctx.Request.Context(), req)
	if err != nil {
		helper.NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	helper.NewSuccessResponse(ctx, statusCode, "update success", record)
}
