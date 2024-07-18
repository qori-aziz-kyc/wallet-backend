package userusecase

import (
	"context"
	"fmt"

	"github.com/qori-aziz-kyc/wallet-backend/internal/http/datatransfer/request"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/helper"
	"github.com/qori-aziz-kyc/wallet-backend/internal/library/models"
	"gorm.io/gorm"
)

func (uc *userUsecase) Register(ctx context.Context, req request.RegisterRequest) (int, *models.User, error) {
	if len(req.Password) < 4 {
		return 422, nil, fmt.Errorf("password must be 4 character or more")
	}

	// Find Username
	userFind, err := uc.userRepo.FindOneBy(map[string]interface{}{
		"username": req.Username,
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		return 422, nil, fmt.Errorf("failed register : %v", err)
	}

	if userFind != nil {
		return 422, nil, fmt.Errorf("username already used")
	}

	// Create user and hashing
	hashedPassword, err := helper.GenerateHash(req.Password)
	if err != nil {
		return 422, nil, fmt.Errorf("failed register : %v", err)
	}

	user := &models.User{
		Username: req.Username,
		Password: hashedPassword,
		IsAdmin:  false,
	}

	tx := uc.db.Begin()
	defer tx.Rollback()
	users, err := uc.userRepo.Create([]*models.User{user}, tx)
	if err != nil {
		return 422, nil, fmt.Errorf("failed register : %v", err)
	}

	tx.Commit()
	return 200, users[0], nil
}
