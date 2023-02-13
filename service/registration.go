package service

import (
	"Go-Handson/config"
	"Go-Handson/entity"
	"github.com/google/uuid"
)

type Registration struct {
	Repo Registrationer
}

func (r *Registration) GenerateUserID() (string, error) {
	// ユーザーIDの決定
	userID := uuid.New().String() // 例: 4b3403665fea6
	return userID, nil
}

func (r *Registration) CreateUserProfile(userID string) (*entity.UserProfile, error) {
	// 初期データの設定
	cfg, _ := config.New()
	userProfile := entity.UserProfile{
		UserID:           userID,
		UserName:         "user name",
		Crystal:          cfg.CrystalDefault,
		CrystalFree:      cfg.CrystalFreeDefault,
		FriendCoin:       cfg.FriendCoinDefault,
		TutorialProgress: cfg.TutorialStart,
	}

	err := r.Repo.SaveUserProfile(userProfile)
	if err != nil {

	}
	return &userProfile, err
}
