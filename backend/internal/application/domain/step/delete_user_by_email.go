package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func DeleteUserByEmailStep(e interface{}) error {
	createUserIntention, ok := e.(*entity.CreateUserIntention)
	if ok {
		if createUserIntention.StepError == nil {
			return nil
		}
		err := deleteUserByEmail(createUserIntention.CreateUser.Email)
		if err != nil {
			return err
		}
		return createUserIntention.StepError
	}

	return errors.New(defines.CannotDeleteUser)
}

func deleteUserByEmail(email string) error {
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

	query := `DELETE FROM users WHERE email = $1`
	_, execErr := db.Exec(query, email)
	if execErr != nil {
		return errors.New(defines.CannotDeleteUser)
	}
	return nil
}
