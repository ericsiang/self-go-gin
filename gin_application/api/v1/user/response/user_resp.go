package response

type CreateUserResponse struct {
	UsersId uint   `json:"id"`
	Account string `json:"account"`
}
