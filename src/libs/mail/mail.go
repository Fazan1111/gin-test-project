package mailOTP

import (
	"fmt"
	envconfig "learnGin/src/common/envConfig"
	"learnGin/src/common/util"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMailOTP(to string) (string, error) {
	otp, err := util.GenerateOTP()
	if err != nil {
		fmt.Println("Error generating OTP:", err)
	}

	// Sender
	from := envconfig.GetEnv("MAIL_ADDRESS")
	smtpPassword := envconfig.GetEnv("MAIL_PASSWORD")
	smtpHost := envconfig.GetEnv("MAIL_HOST")
	smtpPort, _ := strconv.Atoi(envconfig.GetEnv("MAIL_PORT"))
	fmt.Println("to mail", map[string]interface{}{
		"otp":  otp,
		"to":   to,
		"host": smtpHost,
		"port": smtpPort,
	})
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your verify OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s", otp))
	d := gomail.NewDialer(smtpHost, smtpPort, from, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		return "", nil
	}
	return otp, nil
}
