package storage

import (
	"os"
	"path"
	"testing"

	"github.com/google/uuid"
)

const (
	PROJECT_ID  = "0"
	ENV         = "test"
	EMPTY_STATE = "{}"
)

type tfStateTmpPath struct {
	path string
}

func newTfStateTmpPath() *tfStateTmpPath {
	tmpDir := os.TempDir()
	uuid := uuid.NewString()

	return &tfStateTmpPath{path: path.Join(tmpDir, uuid)}
}

func (tmp *tfStateTmpPath) remove() {
	os.RemoveAll(tmp.path)

}

func TestWriteState(t *testing.T) {
	tmp := newTfStateTmpPath()
	defer tmp.remove()

	storage := NewLocalStorage(tmp.path)

	err := storage.WriteTfState(PROJECT_ID, ENV, []byte("{}"))

	if err != nil {
		t.Fatalf(`WriteTfState(%s) = %v, <nil>`, tmp.path, err)
	}
}

func TestGetStateWhenStateDoesNotExist(t *testing.T) {
	tmp := newTfStateTmpPath()
	defer tmp.remove()

	storage := NewLocalStorage(tmp.path)

	content, err := storage.ReadTfState(PROJECT_ID, ENV)

	if err != nil || content != nil {
		t.Fatalf(`GetTfState(%s) = %q, %v, want "", <nil>`, tmp.path, content, err)
	}
}

func TestGetStateWhenStateDoesExist(t *testing.T) {
	tmp := newTfStateTmpPath()
	defer tmp.remove()

	storage := NewLocalStorage(tmp.path)
	storage.WriteTfState(PROJECT_ID, ENV, []byte(EMPTY_STATE))

	content, err := storage.ReadTfState(PROJECT_ID, ENV)

	if content == nil {
		t.Fatalf(`GetTfState(%s) = %q, %v, want "%s", <nil>`, tmp.path, content, err, EMPTY_STATE)
	}
}
