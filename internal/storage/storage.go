package storage

type storage interface {
	ReadTfState(id string, env string) ([]byte, error)
	WriteTfState(id string, env string, content []byte) error
}

type localStorage struct {
	path string
}
