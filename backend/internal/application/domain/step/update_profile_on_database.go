package step

import (
	"backend/internal/adpter/dto"
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

func UpdateProfileOnDatabaseStep(e interface{}) error {
	profileIntention, ok := e.(*entity.ProfileIntention)
	if ok {
		return updateProfile(&profileIntention.Profile, &profileIntention.ProfileRequest)
	}
	return errors.New(defines.CannotUpdateProfile)
}

func updateProfile(previousProfile *entity.Profile, newProfile *dto.ProfileFrontDto) error {
	updatedProfile := createUpdatedProfile(newProfile)
	updatedProfile.ID = previousProfile.ID
	updatedProfile.UserID = previousProfile.UserID

	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"))
	db, connErr := sql.Open("postgres", dataSourceName)
	if connErr != nil {
		return errors.New(defines.CannotEstablishDatabaseConnection)
	}
	defer db.Close()

	genderID, getGenderIDError := getGenderID(newProfile.Gender, db)
	if getGenderIDError != nil {
		return getGenderIDError
	}
	sexualPreferenceID, getSexualPreferenceIDError := getSexualPreferenceID(newProfile.SexualPreference, db)
	if getSexualPreferenceIDError != nil {
		return getSexualPreferenceIDError
	}
	updatedProfile.GenderID = genderID
	updatedProfile.SexualPreferenceID = sexualPreferenceID

	query := `UPDATE profile SET first_name = $1, last_name = $2, location = $3, likes_counter = $4, gender_id = $5, tags_list = $6, biography = $7, sexual_preference_id = $8, pictures = $9, view_counter = $10, is_online = $11, last_online_at = $12, account_status = $13 WHERE id = $14`
	_, execErr := db.Exec(
		query,
		updatedProfile.FirstName,
		updatedProfile.LastName,
		updatedProfile.Location,
		updatedProfile.LikesCounter,
		updatedProfile.GenderID,
		updatedProfile.TagsList,
		updatedProfile.Biography,
		updatedProfile.SexualPreferenceID,
		updatedProfile.Pictures,
		updatedProfile.ViewCounter,
		updatedProfile.IsOnline,
		updatedProfile.LastOnlineAt,
		updatedProfile.AccountStatus,
		updatedProfile.ID,
	)
	if execErr != nil {
		return errors.New(defines.CannotUpdateProfile)
	}

	return nil
}

func createUpdatedProfile(profile *dto.ProfileFrontDto) *dto.ProfileDto {
	res := &dto.ProfileDto{
		FirstName:     profile.FirstName,
		LastName:      profile.LastName,
		Location:      profile.Location,
		LikesCounter:  profile.LikesCounter,
		Biography:     profile.Biography,
		ViewCounter:   profile.ViewCounter,
		IsOnline:      profile.IsOnline,
		LastOnlineAt:  profile.LastOnlineAt,
		AccountStatus: profile.AccountStatus,
	}

	tagsList, _ := json.Marshal(entity.TagsList{Tags: profile.TagsList})
	pictures, _ := json.Marshal(entity.Pictures{Picture: profile.Pictures})

	res.TagsList = tagsList
	res.Pictures = pictures

	return res
}

func getGenderID(genderValue string, db *sql.DB) (uint64, error) {
	query := `SELECT * FROM gender WHERE type = $1 LIMIT 1`
	row := db.QueryRow(query, genderValue)

	var gender dto.GenderDto
	err := row.Scan(
		&gender.ID,
		&gender.Type,
	)
	if err != nil {
		return 0, errors.New(defines.CannotGetGenderID)
	}

	return uint64(gender.ID), nil
}

func getSexualPreferenceID(sexualPreferenceList []string, db *sql.DB) (uint64, error) {
	sort.Strings(sexualPreferenceList)
	sexualPreferenceType := strings.Join(sexualPreferenceList, defines.SexualPreferenceSeparator)

	query := `SELECT * FROM sexual_preference WHERE option = $1 LIMIT 1`
	row := db.QueryRow(query, sexualPreferenceType)

	var sexualPreference dto.SexualPreferenceDto
	err := row.Scan(
		&sexualPreference.ID,
		&sexualPreference.Option,
	)
	if err != nil {
		return 0, errors.New(defines.CannotGetSexualPreferenceID)
	}

	return uint64(sexualPreference.ID), nil
}
