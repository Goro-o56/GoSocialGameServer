package handler

import (
	"Go-Handson/config"
	"Go-Handson/service"
	"encoding/json"
	"net/http"
)

type Registration struct {
	Service RegistrationService
}

func (rg *Registration) Registration(w http.ResponseWriter, r *http.Request) {
	userID, err := rg.Service.GenerateUserID()
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
	userProfile, err := rg.Service.CreateUserProfile(userID, cfg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// クライアントへのレスポンス
	type RegistrationResponse struct {
		UserProfile *service.UserProfile
	}

	response := &RegistrationResponse{
		UserProfile: &userProfile,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
