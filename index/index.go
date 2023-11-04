package index

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/meeron/silkania/models"
)

var ixs map[string]*Index = make(map[string]*Index)
var indexPath string

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

func Get(name string) *Index {
	return ixs[name]
}

func Load(basePath string) error {
	entries, err := os.ReadDir(basePath)
	_, ok := err.(*fs.PathError)
	if ok {
		err = os.Mkdir(basePath, 0777)
	}

	if err != nil {
		return err
	}

	indexPath = basePath

	for _, entry := range entries {
		name := entry.Name()
		dbPath := path.Join(basePath, name)

		bleveIx, err := bleve.Open(dbPath)

		if err != nil {
			return err
		}

		ixs[name] = &Index{
			bleve: bleveIx,
		}
	}

	return nil
}

func Create(name string, mapping models.IndexMapping) error {
	indexMapping := bleve.NewIndexMapping()

	rootDocumentMapping := bleve.NewDocumentMapping()

	mapFields(rootDocumentMapping, mapping.Fields)

	indexMapping.AddDocumentMapping(mapping.DocumentType, rootDocumentMapping)

	path := path.Join(indexPath, name)
	fmt.Printf("%s\n", path)

	ix, err := bleve.New(path, indexMapping)
	if err != nil {
		return err
	}

	ixs[name] = &Index{
		bleve: ix,
	}

	return nil

}

func Drop(name string) error {
	ix := ixs[name]
	if ix == nil {
		return errors.New("index not found")
	}

	if err := ix.bleve.Close(); err != nil {
		return err
	}
	delete(ixs, name)

	path := path.Join(indexPath, name)
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	return nil
}

func mapFields(document *mapping.DocumentMapping, fields map[string]models.FieldMapping) {
	for name, v := range fields {

		if len(v.Fields) == 0 {
			field := bleve.NewTextFieldMapping()
			field.Analyzer = v.Lang
			field.Name = name

			document.AddFieldMapping(field)
			continue
		}

		subDocument := bleve.NewDocumentMapping()

		mapFields(subDocument, v.Fields)

		document.AddSubDocumentMapping(name, subDocument)
	}
}
