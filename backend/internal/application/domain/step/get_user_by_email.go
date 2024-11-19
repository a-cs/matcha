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

func GetUserByEmailStep(e interface{}) error {
	createUserIntention, ok := e.(*entity.CreateUserIntention)
	if ok {
		user, err := getUserByEmail(createUserIntention.CreateUser.Email)
		if err != nil {
			return err
		}
		createUserIntention.User = *user
		return nil
	}
	loginIntention, ok := e.(*entity.LoginIntention)
	if ok {
		user, err := getUserByEmail(loginIntention.Email)
		if err != nil {
			return err
		}
		loginIntention.User = *user
		return nil
	}
	passwordIntention, ok := e.(*entity.PasswordIntention)
	if ok {
		user, err := getUserByEmail(passwordIntention.Email)
		if err != nil {
			return err
		}
		passwordIntention.User = *user
		return nil
	}

	return errors.New(defines.CannotGetUserByEmail)
}

func getUserByEmail(email string) (*entity.User, error) {
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

	query := `SELECT * FROM users WHERE email = $1 LIMIT 1`
	row := db.QueryRow(query, email)

	var user dto.UserDto
	err = row.Scan(&user.ID, &user.Email, &user.Password, &user.Username, &user.ActiveMatches, &user.AccountStatus, &user.SlugID, &user.RecoverPasswordSlugID)
	if err != nil {
		return nil, err
	}

	return mapper.ToUser(&user), nil
}
