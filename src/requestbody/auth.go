package requestbody

type Register struct {
	Username   string `json:"username" validate:"required,min=3,max=20"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6,max=225"`
	Created_at string `json:"created_at"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}