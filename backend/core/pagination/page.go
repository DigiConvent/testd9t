package pagination

type Page[T any] struct {
	Page         int `json:"page"`
	TotalPages   int `json:"total_pages"`
	ItemsPerPage int `json:"items_per_page"`
	Items        []T `json:"items"`
}
