package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"
)

func CreateProfileOnDatabaseStep(e interface{}) error {
	createUserIntention, ok := e.(*entity.CreateUserIntention)
	if ok {
		err := createProfileOnDatabase(createUserIntention)
		if err != nil {
			createUserIntention.StepError = err
		}
		return nil
	}

	return errors.New(defines.CannotCreateProfileOnDatabase)
}

func createProfileOnDatabase(intention *entity.CreateUserIntention) error {
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

	query := `INSERT INTO profile (user_id, first_name, last_name, location, likes_counter, gender_id, tags_list, biography, sexual_preference_id, pictures, view_counter, is_online, last_online_at, account_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	_, execErr := db.Exec(
		query,
		intention.User.ID,
		intention.CreateUser.FirstName,
		intention.CreateUser.LastName,
		defines.EmptyString,
		0,
		1,
		defines.EmptyJson,
		defines.EmptyString,
		1,
		defines.EmptyJson,
		0,
		false,
		time.Now().Format(time.RFC3339),
		defines.PendingStatus,
	)
	if execErr != nil {
		return errors.New(defines.CannotCreateProfileOnDatabase)
	}

	return nil
}
