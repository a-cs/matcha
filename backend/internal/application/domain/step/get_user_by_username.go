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

func GetUserByUsernameStep(e interface{}) error {
	profileIntention, ok := e.(*entity.ProfileIntention)
	if ok {
		user, err := getUserByUsername(profileIntention.Username)
		if err != nil {
			return err
		}
		profileIntention.User = *user
		return nil
	}
	changePasswordIntention, ok := e.(*entity.PasswordIntention)
	if ok {
		user, err := getUserByUsername(changePasswordIntention.JwtObj.Username)
		if err != nil {
			return err
		}
		changePasswordIntention.User = *user
		return nil
	}

	return errors.New(defines.CannotGetUserByUsername)
}

func getUserByUsername(username string) (*entity.User, error) {
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

	query := `SELECT * FROM users WHERE username = $1 LIMIT 1`
	row := db.QueryRow(query, username)

	var user dto.UserDto
	err = row.Scan(&user.ID, &user.Email, &user.Password, &user.Username, &user.ActiveMatches, &user.AccountStatus, &user.SlugID, &user.RecoverPasswordSlugID)
	if err != nil {
		return nil, errors.New(defines.CannotGetUserByUsername)
	}

	return mapper.ToUser(&user), nil
}
