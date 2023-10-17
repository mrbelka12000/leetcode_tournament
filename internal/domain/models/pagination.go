package models

type PaginationParams struct {
	Offset int64 `json:"page" schema:"page"`
	Limit  int64 `json:"page_size" schema:"page_size"`
}

type PaginatedListRepSt struct {
	Page       int64       `json:"page" schema:"page"`
	PageSize   int64       `json:"page_size" schema:"page_size"`
	TotalCount int64       `json:"total_count" schema:"total_count"`
	Results    interface{} `json:"results" schema:"results"`
}
