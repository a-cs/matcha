package dto

import (
	"backend/internal/utils"
	"encoding/json"
)

type ProfileDto struct {
	ID                 uint64          `json:"id,omitempty"`
	UserID             uint64          `json:"user_id,omitempty"`
	FirstName          string          `json:"first_name,omitempty"`
	LastName           string          `json:"last_name,omitempty"`
	Location           string          `json:"location,omitempty"`
	LikesCounter       uint64          `json:"likes_counter,omitempty"`
	GenderID           uint64          `json:"gender_id,omitempty"`
	TagsList           json.RawMessage `json:"tags_list,omitempty"`
	Biography          string          `json:"biography,omitempty"`
	SexualPreferenceID uint64          `json:"sexual_preference_id,omitempty"`
	Pictures           json.RawMessage `json:"pictures,omitempty"`
	ViewCounter        uint64          `json:"view_counter,omitempty"`
	IsOnline           bool            `json:"is_online,omitempty"`
	LastOnlineAt       string          `json:"last_online_at,omitempty"`
	AccountStatus      string          `json:"account_status,omitempty"`
}

type ProfileFrontDto struct {
	ID               uint64   `json:"id"`
	UserID           uint64   `json:"user_id"`
	FirstName        string   `json:"first_name"`
	LastName         string   `json:"last_name"`
	Location         string   `json:"location"`
	LikesCounter     uint64   `json:"likes_counter"`
	Gender           string   `json:"gender"`
	TagsList         []string `json:"tags_list"`
	Biography        string   `json:"biography"`
	SexualPreference []string `json:"sexual_preference"`
	Pictures         []string `json:"pictures"`
	ViewCounter      uint64   `json:"view_counter"`
	IsOnline         bool     `json:"is_online"`
	LastOnlineAt     string   `json:"last_online_at"`
	AccountStatus    string   `json:"account_status"`
}

type ConfirmAccountDto struct {
	SlugID string `json:"slugId"`
}

func (c ConfirmAccountDto) IsValid() bool {
	return len(c.SlugID) > 0 && !utils.IsSQLInjection(c.SlugID)
}

func (p ProfileFrontDto) IsValid() bool {
	if utils.IsSQLInjection(p.FirstName, p.LastName, p.Location, p.Gender, p.Biography, p.LastOnlineAt, p.AccountStatus) {
		return false
	}
	for _, tag := range p.TagsList {
		if utils.IsSQLInjection(tag) {
			return false
		}
	}
	for _, preference := range p.SexualPreference {
		if utils.IsSQLInjection(preference) {
			return false
		}
	}
	for _, picture := range p.Pictures {
		if utils.IsSQLInjection(picture) {
			return false
		}
	}

	return p.ID > 0 && p.UserID > 0
}
