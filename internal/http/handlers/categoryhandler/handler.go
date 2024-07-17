package categoryhandler

import (
	"github.com/gin-gonic/gin"
	categoryusecase "github.com/qori-aziz-kyc/wallet-backend/internal/business/usecases/categoryUsecase"
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
