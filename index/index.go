package index

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
)

type Index struct {
	bleve bleve.Index
}

func (ix *Index) IndexDocument(id string, doc any) error {
	return ix.bleve.Index(id, doc)
}

func (ix *Index) Search(query string) ([]any, uint64, error) {
	if query == "" {
		return make([]any, 0), 0, nil
	}

	q := bleve.NewQueryStringQuery(query)
	searchRequest := bleve.NewSearchRequest(q)
	searchRequest.Fields = []string{"*"}

	result, err := ix.bleve.Search(searchRequest)
	if err != nil {
		return make([]any, 0), 0, err
	}

	items := make([]any, 0)
	for _, d := range result.Hits {
		items = append(items, d.Fields)
	}

	return items, result.Total, nil
}

func (ix *Index) GetDocument(id string) (any, error) {
	query := bleve.NewDocIDQuery([]string{id})

	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"*"}

	result, err := ix.bleve.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	if result.Total == 0 {
		return nil, nil
	}

	return result.Hits[0].Fields, nil
}

func (ix *Index) DeleteDocument(id string) error {
	return ix.bleve.Delete(id)
}

func (ix *Index) Batch(idField string, items []map[string]any) error {
	batch := ix.bleve.NewBatch()

	for _, itm := range items {
		id, ok := itm[idField].(string)
		if !ok {
			return fmt.Errorf("id must be string (id=%v)", itm[idField])
		}

		if err := batch.Index(id, itm); err != nil {
			return err
		}
	}

	return ix.bleve.Batch(batch)
}
