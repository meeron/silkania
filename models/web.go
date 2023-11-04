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

type ServerStats struct {
	Uptime  string                    `json:"upTime"`
	Indexes map[string]map[string]any `json:"indexes"`
}
