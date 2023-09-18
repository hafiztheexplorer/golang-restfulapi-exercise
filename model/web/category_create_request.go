package web

type CategoryCreateRequest struct {
	Namakategori string `validate:"required,max=255,min=1" json:"namakategori"`
}
