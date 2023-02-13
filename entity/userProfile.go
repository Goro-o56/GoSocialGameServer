package entity

type UserProfile struct {
	UserID           string `json:"user_id"`
	UserName         string `json:"user_name"`
	Crystal          int    `json:"crystal"`
	CrystalFree      int    `json:"crystal_free"`
	FriendCoin       int    `json:"friend_coin"`
	TutorialProgress int    `json:"tutorial_progress"`
}
