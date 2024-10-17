package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"errors"
	"gopkg.in/gomail.v2"
	"os"
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
	m.SetHeader("To", intention.CreateUser.Email)
	m.SetHeader("Subject", "Welcome to Matcha!")
	m.SetBody(
		"text/plain",
		"Welcome aboard!\n\nClick here to activate your account: "+os.Getenv("FRONTEND_URL")+"/confirm/"+intention.SlugID)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
