package handler

import (
	"Go-Handson/config"
	"Go-Handson/service"
	"encoding/json"
	"net/http"
)

type RegistrationResponse struct {
	UserProfile *service.UserProfile
}

func Registration(w http.ResponseWriter, r *http.Request) {
	userID, err := service.GenerateUserID()
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// 初期データの設定
	cfg, err := config.New()
	if err != nil {
		//config error

	}
	userProfile, err := service.CreateUserProfile(userID, cfg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// クライアントへのレスポンス
	response := &RegistrationResponse{
		UserProfile: userProfile,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
