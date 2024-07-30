package authDto

type RegisterDto struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" binding:"required"`
}

type VerifyMailOTP struct {
	Email string `json:"email" validate:"email,required"`
}

type VerifyPhoneOTP struct {
	PhoneNumber string `json:"phoneNumber" `
}
