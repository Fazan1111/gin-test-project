package vonageAPI

import (
	"fmt"
	envconfig "learnGin/src/common/envConfig"
	"learnGin/src/common/util"

	"github.com/vonage/vonage-go-sdk"
)

func SendSMS(to string) string {
	otp, _ := util.GenerateOTP()
	API_KEY := envconfig.GetEnv("VONAGE_KEY")
	API_SECRET := envconfig.GetEnv("VONAGE_SECRET")

	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	smsClient := vonage.NewSMSClient(auth)
	response, _, _ := smsClient.Send("+85590805680", to, fmt.Sprintf("Your OTP code is: %s", otp), vonage.SMSOpts{})
	fmt.Println("response sms", response)
	return otp
}
