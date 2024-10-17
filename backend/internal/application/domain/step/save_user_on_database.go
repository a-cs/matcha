package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func SaveUserOnDatabaseStep(e interface{}) error {
	intention, ok := e.(*entity.CreateUserIntention)
	if !ok {
		return errors.New(defines.CannotSaveUserOnDatabase)
	}

	// TODO: validate database error
	_ = saveUserOnDatabase(intention)
	return nil
}

func saveUserOnDatabase(intention *entity.CreateUserIntention) error {
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

	query := `INSERT INTO users (email, password, username, active_matches, account_status, slug_id) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = db.Exec(
		query,
		intention.User.Email,
		intention.HashedPassword,
		intention.User.Username,
		defines.EmptyJson,
		defines.PendingStatus,
		intention.SlugID,
	)
	if err != nil {
		return err
	}

	return nil
}
