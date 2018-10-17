package godhcp

import (
	"github.com/hunkeelin/govirt/govirtlib"
	"sync"
)

type Conn struct {
	NetConfig    string
	HostConfig   string
	Hostmapinfo  map[string]govirtlib.CreateVmForm
	Netmapinfo   map[string]govirtlib.Network
	hostConfigMu sync.Mutex
	netConfigMu  sync.Mutex
	Trustbytes   []byte
	Certbytes    []byte
	Keybytes     []byte
	ReserveIps   []string
}
