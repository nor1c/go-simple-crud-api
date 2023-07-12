package validators

type CustomerRegister struct {
	FullName string `json:"fullname" validate:"required"`
	Username string `json:"username" validate:"required,lowercase,gte=6"`
	Password string `json:"password" validate:"required,gte=8"`
}
