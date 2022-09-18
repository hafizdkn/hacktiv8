package params

type UserCreateRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
