package schema

type SignUpRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type SignUpResponse struct {
	Success bool   `json:"success" validate:"required"`
	Msg     string `json:"msg" validate:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Success bool   `json:"success" validate:"required"`
	Token   string `json:"token" validate:"required"`
	Email   string `json:"email" validate:"required"`
}
