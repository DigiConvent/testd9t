package pagination

type Page[T any] struct {
	Page         int `json:"page"`
	ItemsCount   int `json:"items_count"`
	ItemsPerPage int `json:"items_per_page"`
	Items        []T `json:"items"`
}
