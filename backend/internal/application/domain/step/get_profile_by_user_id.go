package step

import (
	"backend/internal/adpter/dto"
	"backend/internal/application/domain/entity"
	"backend/internal/application/mapper"
	"backend/internal/defines"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetProfileByUserIDStep(e interface{}) error {
	loginIntention, ok := e.(*entity.LoginIntention)
	if ok {
		profile, err := getProfileByUserID(loginIntention.User.ID)
		if err != nil {
			return err
		}
		loginIntention.ProfileStatus = profile.AccountStatus
		return nil
	}
	profileIntention, ok := e.(*entity.ProfileIntention)
	if ok {
		var userID uint64
		if profileIntention.User.ID != 0 {
			userID = profileIntention.User.ID
		} else {
			userID = profileIntention.JwtObj.UserID
		}
		profile, err := getProfileByUserID(userID)
		if err != nil {
			return err
		}
		profileIntention.Profile = *profile
		return nil
	}

	return errors.New(defines.CannotGetProfileByUserID)
}

func getProfileByUserID(userID uint64) (*entity.Profile, error) {
	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"))
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query := `SELECT * FROM profile WHERE user_id = $1 LIMIT 1`
	row := db.QueryRow(query, userID)

	var profile dto.ProfileDto
	err = row.Scan(
		&profile.ID,
		&profile.UserID,
		&profile.FirstName,
		&profile.LastName,
		&profile.Location,
		&profile.LikesCounter,
		&profile.GenderID,
		&profile.TagsList,
		&profile.Biography,
		&profile.SexualPreferenceID,
		&profile.Pictures,
		&profile.ViewCounter,
		&profile.IsOnline,
		&profile.LastOnlineAt,
		&profile.AccountStatus,
	)
	if err != nil {
		return nil, err
	}

	res := mapper.FromProfileDtoToProfile(&profile)
	gender, getGenderErr := getGender(profile.GenderID, db)
	if getGenderErr != nil {
		return nil, getGenderErr
	}
	res.Gender = gender
	sexualPreference, getSexualPreferenceErr := getSexualPreferenceList(profile.SexualPreferenceID, db)
	if getSexualPreferenceErr != nil {
		return nil, getSexualPreferenceErr
	}
	res.SexualPreference = sexualPreference

	return res, nil
}

func getGender(genderID uint64, db *sql.DB) (string, error) {
	query := `SELECT * FROM gender WHERE id = $1 LIMIT 1`
	row := db.QueryRow(query, genderID)

	var gender dto.GenderDto
	err := row.Scan(
		&gender.ID,
		&gender.Type,
	)
	if err != nil {
		return defines.EmptyString, err
	}

	return gender.Type, nil
}

func getSexualPreferenceList(sexualPreferenceID uint64, db *sql.DB) ([]string, error) {
	query := `SELECT * FROM sexual_preference WHERE id = $1 LIMIT 1`
	row := db.QueryRow(query, sexualPreferenceID)

	var sexualPreference dto.SexualPreferenceDto
	err := row.Scan(
		&sexualPreference.ID,
		&sexualPreference.Option,
	)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(sexualPreference.Option, defines.SexualPreferenceSeparator), nil
}
