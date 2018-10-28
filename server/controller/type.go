package controller

import (
	"sync"
)

type Conn struct {
	cb     []byte
	kb     []byte
	tb     []byte
	postMu sync.Mutex
	authcb []byte
	authkb []byte
	authtb []byte
}
