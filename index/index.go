package index

import (
	"fmt"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/meeron/silkania/models"
)

type Index struct {
	bleve bleve.Index
}

func (ix *Index) IndexDocument(id string, doc any) error {
	return ix.bleve.Index(id, doc)
}

func (ix *Index) Search(req *models.SearchReq) ([]any, uint64, error) {
	if req.Q == "" {
		return make([]any, 0), 0, nil
	}

	sortBy := make([]string, 0)

	if req.SortBy == "" {
		req.SortBy = "-_score"
	}

	sorts := strings.Split(req.SortBy, ",")
	for _, s := range sorts {
		sortBy = append(sortBy, strings.TrimSpace(s))
	}

	q := bleve.NewQueryStringQuery(req.Q)

	searchRequest := bleve.NewSearchRequest(q)
	searchRequest.Fields = []string{"*"}
	searchRequest.Size = req.ItemsPerPage
	searchRequest.From = (req.Page - 1) * req.ItemsPerPage
	searchRequest.SortBy(sortBy)

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

func (ix *Index) GetStats() map[string]any {
	stats := ix.bleve.StatsMap()

	docCount, _ := ix.bleve.DocCount()

	stats["DocCount"] = docCount

	return stats
}
