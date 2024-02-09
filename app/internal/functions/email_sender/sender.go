package emailsender

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func Sender(recipient, text string) error {
	err := godotenv.Load("internal/functions/email_sender/.env")
	if err != nil {
		log.Fatal(err)
	}

	from := os.Getenv("EMAIL_FROM")
	if from == "" {
		log.Fatal("Can`t find email in env")
	}

	password := os.Getenv("EMAIL_PASS")
	if password == "" {
		log.Fatal("Can`t find password in env")
	}

	// Receiver email address.
	to := []string{
		recipient,
	}

	// smtp server configuration.
	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost == "" {
		log.Fatal("Can`t find smtp host in env")
	}
	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort == "" {
		log.Fatal("Can`t find smtp port in env")
	}

	// Message.
	message := []byte(text)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
