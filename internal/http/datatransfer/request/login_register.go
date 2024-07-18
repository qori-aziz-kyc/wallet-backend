package request

type RegisterRequest struct {
	ID       int    `json:"id" form:"id" uri:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
