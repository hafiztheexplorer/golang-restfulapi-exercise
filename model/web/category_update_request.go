package web

type CategoryUpdateRequest struct {
	Id           int    `validate:"required" json:"id"`
	Namakategori string `validate:"required,max=255,min=1" json:"nomakategori"`
}
