package index

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/blevesearch/bleve/v2"
)

var ixs map[string]*Index = make(map[string]*Index)
var indexPath string

type Index struct {
	ix bleve.Index
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

		ix, err := bleve.Open(dbPath)

		if err != nil {
			return err
		}

		ixs[name] = &Index{
			ix: ix,
		}
	}

	return nil
}

func Create(name string, lang string) error {
	mapping := bleve.NewIndexMapping()

	path := path.Join(indexPath, name)
	fmt.Printf("%s\n", path)

	ix, err := bleve.New(path, mapping)
	if err != nil {
		return err
	}

	ixs[name] = &Index{
		ix: ix,
	}

	return nil

}
