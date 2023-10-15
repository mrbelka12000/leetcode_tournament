package models

type PaginationParams struct {
	Offset int64 `json:"page"`
	Limit  int64 `json:"page_size"`
}

type PaginatedListRepSt struct {
	Page       int64       `json:"page"`
	PageSize   int64       `json:"page_size"`
	TotalCount int64       `json:"total_count"`
	Results    interface{} `json:"results"`
}
