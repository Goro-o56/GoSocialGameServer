package store

import (
	"Go-Handson/config"
	"database/sql"
	"errors"
	"log"
)

type UserProfile struct {
	UserID           string `json:"user_id"`
	UserName         string `json:"user_name"`
	Crystal          int    `json:"crystal"`
	CrystalFree      int    `json:"crystal_free"`
	FriendCoin       int    `json:"friend_coin"`
	TutorialProgress int    `json:"tutorial_progress"`
}

func SaveUserProfile(userProfile *UserProfile, cfg *config.Config) error {
	// データの書き込み
	db, err := sql.Open("mysql", cfg.DBUser+":"+cfg.DBPassword+"@/"+cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO user_profile (user_id, user_name, crystal, crystal_free, friend_coin, tutorial_progress) VALUES (?, ?, ?, ?, ?, ?)",
		userProfile.UserID, userProfile.UserName, userProfile.Crystal, userProfile.CrystalFree, userProfile.FriendCoin, userProfile.TutorialProgress)
	if err != nil {
		log.Println(err)
		return errors.New(cfg.ErrorDbUpdate)
	}
	return nil
}
