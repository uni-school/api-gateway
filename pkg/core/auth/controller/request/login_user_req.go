package request_controller_auth

type LoginUserRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginUserRequestBody) CustomValidate() error {
	return nil
}
