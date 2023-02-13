package store

import (
	"Go-Handson/config"
	"Go-Handson/entity"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func (r *Repository) SaveUserProfile(userProfile entity.UserProfile) error {
	// データの書き込み
	cfg, _ := config.New()
	db, err := sql.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.DBUser, cfg.DBPassword,
			cfg.DBHost, cfg.DBPort,
			cfg.DBName,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO user_profile (user_id, user_name, crystal, crystal_free, friend_coin, tutorial_progress) VALUES (?, ?, ?, ?, ?, ?)",
		userProfile.UserID, userProfile.UserName, userProfile.Crystal, userProfile.CrystalFree, userProfile.FriendCoin, userProfile.TutorialProgress)
	if err != nil {
		log.Println(err)
		log.Println("args:", userProfile.UserID, userProfile.UserName, userProfile.Crystal, userProfile.CrystalFree, userProfile.FriendCoin, userProfile.TutorialProgress)
		return errors.New(cfg.ErrorDbUpdate)
	}
	return nil
}
