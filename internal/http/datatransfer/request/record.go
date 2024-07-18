package request

type RecordCreateRequest struct {
	ID         int     `json:"id" uri:"id"`
	UserID     int     `json:"-"`
	CategoryID int     `json:"category_id"`
	Amount     float64 `json:"amount"`
	Date       string  `json:"date"`
	Type       string  `json:"type"`
	Note       string  `json:"note"`
}

type RecordFindRequest struct {
	UserID     int    `json:"id" form:"id" uri:"id"`
	CategoryID []int  `json:"category_id" form:"category_id"`
	StartFrom  string `json:"start_from" form:"start_from"`
	To         string `json:"to" form:"to"`
}
