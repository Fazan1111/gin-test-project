package twilioLip

import (
	"fmt"
	envconfig "learnGin/src/common/envConfig"
	"learnGin/src/common/util"
	"os"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(to string, msg string) string {
	otp, _ := util.GenerateOTP()
	client := twilio.NewRestClient()
	twilioId := envconfig.GetEnv("TWILIO_ID")
	params := &api.CreateMessageParams{}
	params.SetBody(fmt.Sprintf("Your OTP code is: %s", otp))
	params.SetMessagingServiceSid(twilioId)
	params.SetTo(to)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return otp
}
