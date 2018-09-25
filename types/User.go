package types

type User struct {
	ID int					`json:"id" form:"id"`
	Username string	`json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}