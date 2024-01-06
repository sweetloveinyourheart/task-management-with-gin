package dtos

type NewBoardDTO struct {
	Title string `validate:"required,min=3,max=200" json:"title"`
}

type AddBoardMembers struct {
	Members []uint `validate:"required" json:"members"`
}
