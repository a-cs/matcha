package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"errors"
	"gopkg.in/gomail.v2"
)

func SendEmailStep(e interface{}) error {
	intention, ok := e.(*entity.CreateUserIntention)
	if !ok {
		return errors.New(defines.CannotSendEmail)
	}

	return sendEmail(intention)
}

func sendEmail(intention *entity.CreateUserIntention) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "noreply@matcha.com")
	m.SetHeader("To", intention.User.Email)
	m.SetHeader("Subject", "Welcome to Matcha!")
	m.SetBody(
		"text/plain",
		"Welcome aboard!\n\nClick here to activate your account: http://localhost:8000/confirm/"+intention.SlugID)

	d := gomail.NewDialer("smtp.gmail.com", 587, "rfslonghi@gmail.com", "usfb gewv hvun sedc")

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
