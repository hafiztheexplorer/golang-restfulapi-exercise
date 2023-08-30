package web

type CategoryUpdateRequest struct {
	Id           int64  `validate:"required"`
	Namakategori string `validate:"required,max=255,min=1"`
}
