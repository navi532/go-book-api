package requests

type DeleteBookRequest struct {
	IDS []string `json:"ids" validate:"required,gt=0" `
}
