package categoryhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/qori-aziz-kyc/wallet-backend/internal/business/usecases/categoryusecase"
)

type CategoryHandler interface {
	CreateHandler(ctx *gin.Context)
	UpdateHandler(ctx *gin.Context)
	FindHandler(ctx *gin.Context)
}

type categoryHandler struct {
	usecase categoryusecase.CategoryUsecase
}

func NewCategoryHandler(usecase categoryusecase.CategoryUsecase) CategoryHandler {
	return &categoryHandler{
		usecase: usecase,
	}
}
