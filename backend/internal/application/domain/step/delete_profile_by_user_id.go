package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func DeleteProfileByUserIDStep(e interface{}) error {
	createUserIntention, ok := e.(*entity.CreateUserIntention)
	if ok {
		if createUserIntention.StepError == nil {
			return nil
		}
		err := deleteProfileByUserID(createUserIntention.User.ID)
		if err != nil {
			createUserIntention.StepError = err
		}
		return nil
	}

	return errors.New(defines.CannotDeleteProfile)
}

func deleteProfileByUserID(userID uint64) error {
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

	query := `DELETE FROM profile WHERE user_id = $1`
	_, execErr := db.Exec(query, userID)
	if execErr != nil {
		return errors.New(defines.CannotDeleteProfile)
	}
	return nil
}
