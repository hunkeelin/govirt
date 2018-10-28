package controller

import (
	"sync"
	"time"
)

type Conn struct {
	cb     []byte
	kb     []byte
	tb     []byte
	postMu sync.Mutex
	authcb []byte
	authkb []byte
	authtb []byte
	Ixml   map[string][]byte
	rmap   map[string]rlimit
}
type rlimit struct {
	cpu       int       // vcpu
	mem       int       // mem in GB
	timelimit time.Time // in hours.h
}
