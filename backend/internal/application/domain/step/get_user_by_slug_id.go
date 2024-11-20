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

func GetUserBySlugIDStep(e interface{}) error {
	confirmAccountIntention, ok := e.(*entity.ConfirmAccountIntention)
	if ok {
		user, err := getUserBySlugID(confirmAccountIntention.SlugID)
		if err != nil {
			return err
		}
		if user.AccountStatus == defines.ActiveStatus {
			return errors.New(defines.AccountAlreadyActive)
		}
		confirmAccountIntention.User = *user
		return nil
	}

	return errors.New(defines.CannotGetUserBySlugID)
}

func getUserBySlugID(slugID string) (*entity.User, error) {
	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"))
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, errors.New(defines.CannotEstablishDatabaseConnection)
	}
	defer db.Close()

	query := `SELECT * FROM users WHERE slug_id = $1 LIMIT 1`
	row := db.QueryRow(query, slugID)

	var user dto.UserDto
	err = row.Scan(&user.ID, &user.Email, &user.Password, &user.Username, &user.ActiveMatches, &user.AccountStatus, &user.SlugID, &user.RecoverPasswordSlugID)
	if err != nil {
		return nil, errors.New(defines.CannotGetUserBySlugID)
	}
	return mapper.ToUser(&user), nil
}
