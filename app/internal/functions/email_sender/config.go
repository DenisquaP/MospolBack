package emailsender

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	From     string
	Password string
	Host     string
	Address  string
}

func LoadConfig() Config {
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

	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost == "" {
		log.Fatal("Can`t find smtp host in env")
	}

	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort == "" {
		log.Fatal("Can`t find smtp port in env")
	}

	return Config{
		From:     from,
		Password: password,
		Host:     smtpHost,
		Address:  smtpHost + ":" + smtpPort,
	}
}
