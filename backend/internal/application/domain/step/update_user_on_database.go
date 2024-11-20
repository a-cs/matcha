package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func UpdateUserOnDatabaseStep(e interface{}) error {
	confirmAccountIntention, ok := e.(*entity.ConfirmAccountIntention)
	if ok {
		confirmAccountIntention.User.AccountStatus = defines.ActiveStatus
		return updateUser(&confirmAccountIntention.User)
	}
	passwordIntention, ok := e.(*entity.PasswordIntention)
	if ok {
		return updateUser(&passwordIntention.User)
	}

	return errors.New(defines.CannotUpdateUser)
}

func updateUser(user *entity.User) error {
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

	query := `UPDATE users SET username = $1, active_matches = $2, account_status = $3, recovery_slug_id = $4 WHERE id = $5`
	_, execErr := db.Exec(
		query,
		user.Username,
		user.ActiveMatches,
		user.AccountStatus,
		user.RecoverPasswordSlugID,
		user.ID,
	)
	if execErr != nil {
		return errors.New(defines.CannotUpdateUser)
	}

	return nil
}
