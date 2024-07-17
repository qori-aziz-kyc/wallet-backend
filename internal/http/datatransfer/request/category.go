package request

type CategoryRequest struct {
	ID         int    `json:"id" form:"id" uri:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	CategoryID []int  `form:"category_id"`
}
