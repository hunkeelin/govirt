package govirthost

import (
	"github.com/digitalocean/go-libvirt"
	"sync"
)

type Conn struct {
	L          *libvirt.Libvirt // listening socket
	postMu     sync.Mutex
	getMu      sync.Mutex
	Trustbytes []byte
	Certbytes  []byte
	Keybytes   []byte
}
