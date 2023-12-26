package requestbody

type Users struct {
	Username string
	Email    string
	Password string
}

type UpdateUsersUsername struct {
	Username string
}
