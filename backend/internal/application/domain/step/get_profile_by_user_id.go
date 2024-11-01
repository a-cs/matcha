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
)

func GetProfileByUserIDStep(e interface{}) error {
	loginIntention, ok := e.(*entity.LoginIntention)
	if ok {
		profile, err := getProfileByUserID(loginIntention.User.ID)
		if err != nil {
			return err
		}
		loginIntention.Profile = *profile
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
	return mapper.ToProfile(&profile), nil
}
