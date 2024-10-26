package dto

import "backend/internal/utils"

type ProfileDto struct {
	ID                 uint64      `json:"id,omitempty"`
	UserID             uint64      `json:"user_id,omitempty"`
	FirstName          string      `json:"first_name,omitempty"`
	LastName           string      `json:"last_name,omitempty"`
	Location           string      `json:"location,omitempty"`
	LikesCounter       uint64      `json:"likes_counter,omitempty"`
	GenderID           uint64      `json:"gender_id,omitempty"`
	TagsList           interface{} `json:"tags_list,omitempty"`
	Biography          string      `json:"biography,omitempty"`
	SexualPreferenceID uint64      `json:"sexual_preference_id,omitempty"`
	Pictures           interface{} `json:"pictures,omitempty"`
	ViewCounter        uint64      `json:"view_counter,omitempty"`
	IsOnline           bool        `json:"is_online,omitempty"`
	LastOnlineAt       string      `json:"last_online_at,omitempty"`
	AccountStatus      string      `json:"account_status,omitempty"`
}

type ConfirmAccountDto struct {
	SlugID string `json:"slugId"`
}

func (c ConfirmAccountDto) IsValid() bool {
	return len(c.SlugID) > 0 && !utils.IsSQLInjection(c.SlugID)
}
