package response

type CreateUserResponse struct {
	Name  string `json:name`
	Email string `json:email`
}
