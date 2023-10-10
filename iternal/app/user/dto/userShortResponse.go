package dto

type UserShortResponse struct {
	Data []UserShort `json:"data"`
}

func (response *UserShortResponse) SetData(UserShort []UserShort) {
	response.Data = UserShort
}
