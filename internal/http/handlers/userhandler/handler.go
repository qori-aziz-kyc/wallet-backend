package userhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/qori-aziz-kyc/wallet-backend/internal/business/usecases/userusecase"
)

type UserHandler interface {
	LoginHandler(ctx *gin.Context)
	RegisterHandler(ctx *gin.Context)
}

type userHandler struct {
	usecase userusecase.UserUsecase
}

func NewUserHandler(usecase userusecase.UserUsecase) UserHandler {
	return &userHandler{
		usecase: usecase,
	}
}
