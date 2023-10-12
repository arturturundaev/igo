package dto

type UserShortRequest struct {
	Limit int `validate:"required,gte=0,lte=130"`
	Page  int `validate:"required,gte=0"`
}
