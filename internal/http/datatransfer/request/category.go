package request

type CategoryRequest struct {
	ID    int    `json:"id" form:"id" uri:"id"`
	Name  string `json:"string"`
	Color string `json:"color"`
}
