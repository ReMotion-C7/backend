package request

type AddProgressDto struct {
	Date string `json:"date" validate:"required"`
}
