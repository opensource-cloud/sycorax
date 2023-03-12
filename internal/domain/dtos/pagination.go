package dtos

type (
	PaginationMetadata struct {
		Count int `json:"count"`
	}
	Pagination struct {
		Items    []interface{}      `json:"items"`
		Metadata PaginationMetadata `json:"metadata"`
	}
)

// NewPagination returns a new instance of Pagination
func NewPagination(items []interface{}) *Pagination {
	count := len(items)
	return &Pagination{
		Items: items,
		Metadata: PaginationMetadata{
			Count: count,
		},
	}
}
