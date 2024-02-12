package emailsender

import (
	"crypto/tls"
	"net/smtp"
)

func Sender(recipient, text string) error {
	config := LoadConfig()
	message := []byte(
		"Subject: Новый коментарий\r\n" + // Установка темы письма
			"\r\n" + // Пустая строка для разделения заголовков и тела письма
			"Новый коментарий от",
	)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         config.Host,
	}

	conn, err := tls.Dial("tcp", config.Address, tlsConfig)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, config.Host)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("tcp", config.From, config.Password, config.Host)
	if err := client.Auth(auth); err != nil {
		return err
	}

	if err := client.Mail(config.From); err != nil {
		return err
	}

	if err := client.Rcpt(recipient); err != nil {
		return err
	}

	wc, err := client.Data()
	if err != nil {
		return err
	}

	_, err = wc.Write(message)
	if err != nil {
		return err
	}

	err = wc.Close()
	if err != nil {
		return err
	}

	return nil
}
