package model

type PaginationOptions struct {
	Limit  int
	Offset int
}

type PaginatedOutput[T interface{}] struct {
	Rows  []T `json:"rows"`
	Count int `json:"count"`
}
