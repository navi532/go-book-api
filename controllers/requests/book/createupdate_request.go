package requests

type CreateUpdateBookRequest struct {
	Title     string   `json:"title" validate:"required"`
	Authors   []Author `json:"authors" validate:"dive,gt=0,required"`
	PageCount int      `json:"page_count" validate:"required"`
}
