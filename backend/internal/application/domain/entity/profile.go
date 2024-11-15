package entity

import "backend/internal/adpter/dto"

type ProfileIntention struct {
	Username        string
	User            User
	Profile         Profile
	ProfileRequest  dto.ProfileFrontDto
	JwtToken        string
	JwtObj          JwtObj
	ProfileResponse dto.ProfileFrontDto
}

type Profile struct {
	ID                 uint64
	UserID             uint64
	FirstName          string
	LastName           string
	Location           string
	LikesCounter       uint64
	GenderID           uint64
	Gender             string
	TagsList           TagsList
	Biography          string
	SexualPreferenceID uint64
	SexualPreference   []string
	Pictures           Pictures
	ViewCounter        uint64
	IsOnline           bool
	LastOnlineAt       string
	AccountStatus      string
}

type TagsList struct {
	Tags []string
}

type Pictures struct {
	Picture []string
}
