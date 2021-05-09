package main

import (
	"fmt"
	"log"
	"net/smtp"
)

var (
	from       = "sethiharsh0811@gmail.com"
	recipients = []string{"arshsethi08@gmail.com"}
)

func Notify(msg string) {
	hostname := "smtp.gmail.com"

	auth := smtp.PlainAuth("", "sethiharsh0811@gmail.com", "password", hostname)

	message := []byte(fmt.Sprintf("To: recipient@example.net\r\n"+
		"Subject: CoWin Results!\r\n"+
		"\r\n"+
		"%s.\r\n", msg))

	err := smtp.SendMail(hostname+":587", auth, from, recipients, message)
	if err != nil {
		log.Fatal(err)
	}
}
