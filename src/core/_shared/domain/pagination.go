package domain

type Pagination struct {
	Total        int `json:"total"`
	PerPage      int `json:"per_page"`
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"previous_page"`
}
