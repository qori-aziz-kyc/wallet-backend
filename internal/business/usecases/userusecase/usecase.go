package userusecase

import (
	"context"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/response"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/jwt"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/service"
	"gorm.io/gorm"
)

type UserUsecase interface {
	Login(ctx context.Context, req request.RegisterRequest) (int, *response.LoginResponse, error)
	Register(ctx context.Context, req request.RegisterRequest) (int, *models.User, error)
}

type userUsecase struct {
	jwtService jwt.JWTService
	db         *gorm.DB
	userRepo   service.UserRepo
}

func NewUserUsecase(jwtService jwt.JWTService, db *gorm.DB, userRepo service.UserRepo) UserUsecase {
	return &userUsecase{
		jwtService: jwtService,
		db:         db,
		userRepo:   userRepo,
	}
}
