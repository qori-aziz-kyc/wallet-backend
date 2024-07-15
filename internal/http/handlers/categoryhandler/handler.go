package categoryhandler

import (
	categoryusecase "github.com/qori-aziz-kyc/wallet-backend/internal/business/usecases/categoryUsecase"
)

type CategoryHandler struct {
	usecase categoryusecase.CategoryUsecase
}

func NewCategoryHandler(usecase categoryusecase.CategoryUsecase) CategoryHandler {
	return CategoryHandler{
		usecase: usecase,
	}
}
