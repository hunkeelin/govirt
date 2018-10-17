package storagehost

import (
	"sync"
)

type Conn struct {
	StorageLocation string
	Config          string
	TemplateRegex   string
	storageMu       sync.Mutex
	Trustbytes      []byte
	Certbytes       []byte
	Keybytes        []byte
}
