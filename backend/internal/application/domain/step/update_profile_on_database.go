package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func UpdateProfileOnDatabaseStep(e interface{}) error {
	return errors.New(defines.CannotUpdateProfile)
}

func updateProfile(profile *entity.Profile) error {
	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"))
	db, connErr := sql.Open("postgres", dataSourceName)
	if connErr != nil {
		panic(connErr.Error())
	}
	defer db.Close()

	query := `UPDATE profile SET first_name = $1, last_name = $2, location = $3, likes_counter = $4, gender_id = $5, tags_list = $6, biography = $7, sexual_preference_id = $8, pictures = $9, view_counter = $10, is_online = $11, last_online_at = $12, account_status = $13 WHERE id = $14`
	_, execErr := db.Exec(
		query,
		profile.FirstName,
		profile.LastName,
		profile.Location,
		profile.LikesCounter,
		profile.GenderID,
		profile.TagsList,
		profile.Biography,
		profile.SexualPreferenceID,
		profile.Pictures,
		profile.ViewCounter,
		profile.IsOnline,
		profile.LastOnlineAt,
		profile.AccountStatus,
		profile.ID,
	)
	if execErr != nil {
		return execErr
	}

	return nil
}
