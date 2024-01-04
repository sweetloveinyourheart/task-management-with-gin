package dtos

type RegisterDTO struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6,max=200" json:"password"`
	FullName string `validate:"omitempty,min=3,max=200" json:"full_name"`
}

type SignInDTO struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6,max=200" json:"password"`
}

type UpdateProfileDTO struct {
	FullName string `validate:"omitempty,min=3,max=200" json:"full_name"`
	Username string `validate:"omitempty,min=3,max=200" json:"username"`
}
