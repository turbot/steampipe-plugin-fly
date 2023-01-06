package apiClient

type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`

	// When paginating forwards, the cursor to continue.
	EndCursor string `json:"endCursor"`
}
