package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func UpdateUserPasswordStep(e interface{}) error {
	changePasswordIntention, ok := e.(*entity.PasswordIntention)
	if ok {
		return updatePasswordUser(changePasswordIntention.User.ID, changePasswordIntention.HashedPassword)
	}

	return errors.New(defines.CannotUpdateUserPassword)
}

func updatePasswordUser(userID uint64, newPassword string) error {
	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"))
	db, connErr := sql.Open("postgres", dataSourceName)
	if connErr != nil {
		return errors.New(defines.CannotUpdateUserPassword)
	}
	defer db.Close()

	query := `UPDATE users SET password = $1 WHERE id = $2`
	_, execErr := db.Exec(
		query,
		newPassword,
		userID,
	)
	if execErr != nil {
		return errors.New(defines.CannotUpdateUserPassword)
	}

	return nil
}
