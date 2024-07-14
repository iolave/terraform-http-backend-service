package storage

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

var _ storage = NewLocalStorage("")

func NewLocalStorage(path string) storage {
	return &localStorage{
		path: path,
	}
}

func (storge *localStorage) ReadTfState(id string, env string) ([]byte, error) {
	tfStatePath := storge.getTfStatePath(id, env)

	content, err := os.ReadFile(tfStatePath)

	if os.IsNotExist(err) {
		return nil, nil
	}

	return content, err
}

func (storage *localStorage) WriteTfState(id string, env string, content []byte) error {
	tfStateDir := storage.getTfStateDir(id, env)
	tfStatePath := storage.getTfStatePath(id, env)

	os.MkdirAll(tfStateDir, fs.ModePerm)

	prevContent, err := os.ReadFile(tfStatePath)

	if os.IsNotExist(err) {
		return os.WriteFile(tfStatePath, content, 0666)
	}
	err = os.WriteFile(fmt.Sprintf("%s.bak", tfStatePath), prevContent, 0666)

	if err != nil {
		fmt.Println("unable to write .bak file")
	}

	return os.WriteFile(tfStatePath, content, 0666)
}

func (storge *localStorage) getTfStatePath(id string, env string) string {
	return path.Join(storge.getTfStateDir(id, env), "terraform.tfstate")
}

func (storge *localStorage) getTfStateDir(id string, env string) string {
	return path.Join(storge.path, id, env)
}
