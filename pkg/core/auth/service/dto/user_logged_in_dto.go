package dto_service_auth

type UserDTOForLogin struct {
	ID string
}

type UserLoggedInDTO struct {
	User  UserDTOForLogin
	Token string
}
