package requests

type CreateRequest struct {
	Title     string   `json:"title" validate:"required"`
	Authors   []Author `json:"authors" validate:"required,gt=0,dive,unique=Email"`
	PageCount int      `json:"page_count" validate:"required"`
}
