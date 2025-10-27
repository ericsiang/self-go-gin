package request

type CreateUserRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUsersByIDRequest struct {
	FilterUsersId string `form:"filterUsersId" json:"filterUsersId" binding:"required"`
}
