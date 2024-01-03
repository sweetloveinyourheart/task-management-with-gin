package responses

type SignInResponse struct {
	AccessToken  string
	RefreshToken string
}

type UserProfile struct {
	Id       uint
	Email    string
	Username string
	FullName string
}
