package models

import "errors"

type IndexMapping struct {
	DocumentType string
	Fields       map[string]FieldMapping
}

type FieldMapping struct {
	Lang   string
	Fields map[string]FieldMapping
}

type SearchReq struct {
	Q            string
	SortBy       string
	Page         int
	ItemsPerPage int
}

func (s *SearchReq) Validate() error {
	if s.Page <= 0 {
		return errors.New("'Page' must be greater or equal than 1")
	}

	if s.ItemsPerPage <= 0 {
		return errors.New("'ItemsPerPage' must be greater or equal than 1")
	}

	if s.ItemsPerPage > 100 {
		return errors.New("'ItemsPerPage' must be lower than 100")
	}

	return nil
}
