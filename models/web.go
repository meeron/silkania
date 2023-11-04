package models

type CreateIndexReq struct {
	Name    string
	Mapping IndexMapping
}

type SearchResult struct {
	Total uint64 `json:"total"`
	Items []any  `json:"items"`
}

type BatchReq struct {
	IdField string
	Items   []map[string]any
}
