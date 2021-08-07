package domain

type MetaPagination struct {
	PerPage     int    `json:"perPage"`
	CurrentPage int    `json:"currentPage"`
	TotalPage   int    `json:"totalPage"`
	Total       int64  `json:"total"`
	Offset      int    `json:"offset"`
	Sort        string `json:"sort"`
}
