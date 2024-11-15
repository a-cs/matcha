package mapper

import (
	"backend/internal/adpter/dto"
	"backend/internal/application/domain/entity"
	"encoding/json"
)

func FromProfileDtoToProfile(source *dto.ProfileDto) *entity.Profile {
	if source == nil {
		return nil
	}

	dest := &entity.Profile{
		ID:                 source.ID,
		UserID:             source.UserID,
		FirstName:          source.FirstName,
		LastName:           source.LastName,
		Location:           source.Location,
		LikesCounter:       source.LikesCounter,
		GenderID:           source.GenderID,
		Biography:          source.Biography,
		SexualPreferenceID: source.SexualPreferenceID,
		ViewCounter:        source.ViewCounter,
		IsOnline:           source.IsOnline,
		LastOnlineAt:       source.LastOnlineAt,
		AccountStatus:      source.AccountStatus,
	}

	var tagsList entity.TagsList
	var pictures entity.Pictures

	if err := json.Unmarshal(source.TagsList, &tagsList); err == nil {
		dest.TagsList = tagsList
	}
	if err := json.Unmarshal(source.Pictures, &pictures); err == nil {
		dest.Pictures = pictures
	}

	return dest
}

func FromProfileToProfileFrontDto(source *entity.Profile) *dto.ProfileFrontDto {
	if source == nil {
		return nil
	}

	return &dto.ProfileFrontDto{
		ID:               source.ID,
		UserID:           source.UserID,
		FirstName:        source.FirstName,
		LastName:         source.LastName,
		Location:         source.Location,
		LikesCounter:     source.LikesCounter,
		Gender:           source.Gender,
		TagsList:         source.TagsList.Tags,
		Biography:        source.Biography,
		SexualPreference: source.SexualPreference,
		Pictures:         source.Pictures.Picture,
		ViewCounter:      source.ViewCounter,
		IsOnline:         source.IsOnline,
		LastOnlineAt:     source.LastOnlineAt,
		AccountStatus:    source.AccountStatus,
	}
}
