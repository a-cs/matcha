package dto

type JwtResponseDto struct {
	Token         string `json:"token"`
	ProfileStatus string `json:"profile_status"`
	UserID        uint64 `json:"user_id"`
}
