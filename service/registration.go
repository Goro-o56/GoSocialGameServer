package service

import (
	"Go-Handson/config"
	"github.com/google/uuid"
)

type UserProfile struct {
	UserID           string `json:"user_id"`
	UserName         string `json:"user_name"`
	Crystal          int    `json:"crystal"`
	CrystalFree      int    `json:"crystal_free"`
	FriendCoin       int    `json:"friend_coin"`
	TutorialProgress int    `json:"tutorial_progress"`
}

type Registration struct {
}

func (r *Registration) GenerateUserID() (string, error) {
	// ユーザーIDの決定
	userID := uuid.New().String() // 例: 4b3403665fea6
	return userID, nil
}

func (r *Registration) CreateUserProfile(userID string, cfg *config.Config) (*UserProfile, error) {
	// 初期データの設定
	userProfile := &UserProfile{
		UserID:           userID,
		UserName:         "user name",
		Crystal:          cfg.CrystalDefault,
		CrystalFree:      cfg.CrystalFreeDefault,
		FriendCoin:       cfg.FriendCoinDefault,
		TutorialProgress: cfg.TutorialStart,
	}
	return userProfile, nil
}
