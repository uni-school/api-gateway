package response_controller_auth

type UserResponseForLogin struct {
	ID string `json:"id"`
}

type UserLoggedInResponse struct {
	User  UserResponseForLogin `json:"user"`
	Token string               `json:"token"`
}
