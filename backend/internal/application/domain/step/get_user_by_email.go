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
	intention, ok := e.(*entity.CreateUserIntention)
	if !ok {
		return errors.New(defines.CannotGetUserByEmail)
	}

	// TODO: validate database error
	_ = getUserByEmail(intention)
	return nil
}

func getUserByEmail(intention *entity.CreateUserIntention) error {
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
	row := db.QueryRow(query, intention.CreateUser.Email)

	var user dto.UserDto
	err = row.Scan(&user.ID, &user.Email, &user.Password, &user.Username, &user.ActiveMatches, &user.AccountStatus, &user.SlugID)
	if err != nil {
		return err
	}
	intention.User = *mapper.ToUser(&user)

	return nil
}
