package responses

type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type UserProfile struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}
