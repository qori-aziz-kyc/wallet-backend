package response

type LoginResponse struct {
	UserID      int    `json:"user_id"`
	AccessToken string `json:"access_token"`
}
