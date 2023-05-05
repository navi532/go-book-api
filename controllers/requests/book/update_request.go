package requests

type UpdateRequest struct {
	Title     string   `json:"title" validate:"omitempty"`
	Authors   []Author `json:"authors" validate:"omitempty,dive"`
	PageCount int      `json:"page_count" validate:"omitempty"`
}
