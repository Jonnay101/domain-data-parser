package persist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Persist interface {
	Save(filename string, i interface{}) error
	Get(filename string) ([]byte, error)
}

type persist struct {
	path string
}

func New(path string) (Persist, error) {
	err := os.MkdirAll(path, os.ModePerm)
	return &persist{path}, err
}

// Save will write the provided interface to disk
func (p *persist) Save(filename string, i interface{}) error {
	fullpath := fmt.Sprintf("%s/%s", p.path, filename)

	data, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fullpath, data, os.ModePerm)
}

// Get will retrieve the files data from disk
func (p *persist) Get(filename string) ([]byte, error) {
	fullpath := fmt.Sprintf("%s/%s", p.path, filename)

	return ioutil.ReadFile(fullpath)
}
