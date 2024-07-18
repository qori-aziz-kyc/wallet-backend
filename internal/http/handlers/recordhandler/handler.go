package recordhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/qori-aziz-kyc/wallet-backend/internal/business/usecases/recordusecase"
)

type RecordHandler interface {
	CreateHandler(ctx *gin.Context)
	UpdateHandler(ctx *gin.Context)
	FindHandler(ctx *gin.Context)
}

type recordHandler struct {
	usecase recordusecase.RecordUsecase
}

func NewRecordHandler(usecase recordusecase.RecordUsecase) RecordHandler {
	return &recordHandler{
		usecase: usecase,
	}
}
