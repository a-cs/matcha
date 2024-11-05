package mapper

import (
	"backend/internal/adpter/dto"
	"backend/internal/application/domain/entity"
	"encoding/json"
)

func ToProfile(profileDto *dto.ProfileDto) *entity.Profile {
	var tagsList []entity.Tags
	var pictures []entity.Picture

	if err := json.Unmarshal(profileDto.TagsList, &tagsList); err != nil {
		tagsList = []entity.Tags{}
	}
	if err := json.Unmarshal(profileDto.Pictures, &pictures); err != nil {
		pictures = []entity.Picture{}
	}

	return &entity.Profile{
		ID:                 profileDto.ID,
		UserID:             profileDto.UserID,
		FirstName:          profileDto.FirstName,
		LastName:           profileDto.LastName,
		Location:           profileDto.Location,
		LikesCounter:       profileDto.LikesCounter,
		GenderID:           profileDto.GenderID,
		TagsList:           tagsList,
		Biography:          profileDto.Biography,
		SexualPreferenceID: profileDto.SexualPreferenceID,
		Pictures:           pictures,
		ViewCounter:        profileDto.ViewCounter,
		IsOnline:           profileDto.IsOnline,
		LastOnlineAt:       profileDto.LastOnlineAt,
		AccountStatus:      profileDto.AccountStatus,
	}
}

func ToProfileDto(profile *entity.Profile) *dto.ProfileDto {
	if profile == nil {
		return nil
	}

	tagsList, err := json.Marshal(profile.TagsList)
	if err != nil {
		tagsList = []byte("[]")
	}
	pictures, err := json.Marshal(profile.Pictures)
	if err != nil {
		pictures = []byte("[]")
	}

	return &dto.ProfileDto{
		ID:                 profile.ID,
		UserID:             profile.UserID,
		FirstName:          profile.FirstName,
		LastName:           profile.LastName,
		Location:           profile.Location,
		LikesCounter:       profile.LikesCounter,
		GenderID:           profile.GenderID,
		TagsList:           tagsList,
		Biography:          profile.Biography,
		SexualPreferenceID: profile.SexualPreferenceID,
		Pictures:           pictures,
		ViewCounter:        profile.ViewCounter,
		IsOnline:           profile.IsOnline,
		LastOnlineAt:       profile.LastOnlineAt,
		AccountStatus:      profile.AccountStatus,
	}
}
