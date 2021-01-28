package email

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail2(email, otp string) {

	// Sender data.
	from := "susyadavb12@gmail.com"
	password := os.Getenv("SMTP_PASS")
	email = email
	// Receiver email address.
	to := []string{
		"susyadavb12@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(otp)
	// fmt.Println(message)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
