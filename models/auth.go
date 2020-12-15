package models

//AuthCredentials for login
type AuthCredentials struct {
	Email    string `form:"email" example:"mail@mail.com"`
	Password string `form:"password" example:"123"`
}
