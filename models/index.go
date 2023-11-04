package models

type IndexMapping struct {
	DocumentType string
	Fields       map[string]FieldMapping
}

type FieldMapping struct {
	Lang   string
	Fields map[string]FieldMapping
}
