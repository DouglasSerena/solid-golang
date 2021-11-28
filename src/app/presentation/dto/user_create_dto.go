package dto

type UserCreateRequestDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserCreateResponseDto struct {
	Token string `json:"token"`
}
