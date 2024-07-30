package customJWT

type SignJwtResp struct {
	AccessToken  string
	RefreshToken string
}

type VerifyJwtResp struct {
	Id    string
	Name  string
	Email string
}
