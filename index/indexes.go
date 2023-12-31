package index

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/meeron/silkania/models"

	// to support standard set of build tags
	_ "github.com/blevesearch/bleve/v2/config"
)

var ixs map[string]*Index = make(map[string]*Index)
var indexPath string

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
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		dbPath := path.Join(basePath, name)

		bleveIx, err := bleve.Open(dbPath)

		if err != nil {
			log.Printf("WARN: Cannot load '%s': %v", name, err)
			continue
		}

		ixs[name] = &Index{
			bleve: bleveIx,
		}
	}

	return nil
}

func Create(name string, defaultLang string, mapping models.IndexMapping) error {
	indexMapping := bleve.NewIndexMapping()
	indexMapping.DefaultAnalyzer = defaultLang

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

func All() map[string]*Index {
	return ixs
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
