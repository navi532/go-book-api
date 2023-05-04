package requests

type CreateUpdateBookRequest struct {
	Title     string `json:"title" validate:"required"`
	Author    string `json:"author" validate:"required"`
	PageCount int    `json:"page_count" validate:"required"`
}
