package comm

type Pagination struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	TotalPage   int `json:"total_page"`
	CurrentPage int `json:"current_page"`
}
