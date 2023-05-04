package requests

type UpdateRequest struct {
	Title     string `json:"title" validate:"omitempty"`
	Author    string `json:"author" validate:"omitempty"`
	PageCount int    `json:"page_count" validate:"omitempty"`
}
