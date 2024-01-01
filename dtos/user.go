package dtos

type RegisterDTO struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6,max=200" json:"password"`
	FullName string `validate:"omitempty,min=3,max=200" json:"full_name"`
}
