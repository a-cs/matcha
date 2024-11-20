package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"errors"
	"gopkg.in/gomail.v2"
	"os"
)

func SendEmailStep(e interface{}) error {
	createUserIntention, ok := e.(*entity.CreateUserIntention)
	if ok {
		err := sendEmail(
			createUserIntention.CreateUser.Email,
			defines.ActiveAccountEmailSubject,
			defines.ActiveAccountEmailContent+os.Getenv("FRONTEND_URL")+"/confirm/"+createUserIntention.SlugID)
		if err != nil {
			createUserIntention.StepError = err
		}
		return nil
	}
	passwordIntention, ok := e.(*entity.PasswordIntention)
	if ok {
		return sendEmail(
			passwordIntention.User.Email,
			defines.RecoverPasswordEmailSubject,
			defines.RecoverPasswordEmailContent+os.Getenv("FRONTEND_URL")+"/recover/"+passwordIntention.User.RecoverPasswordSlugID)
	}

	return errors.New(defines.CannotSendEmail)
}

func sendEmail(email, subject, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "noreply@matcha.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", message)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return errors.New(defines.CannotSendEmail)
	}

	return nil
}
