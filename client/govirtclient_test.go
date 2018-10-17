package govirtclient

import (
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/json-iterator/go"
	"testing"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func TestServerGet(t *testing.T) {
	fmt.Println("testing ServerGet")
	get()
}
func TestDhcpNetwork(t *testing.T) {
	fmt.Println("testing dhcpGet network")
	godhcpget()
	fmt.Println("testing end dhcpGet network")
}
func TestDhcpHosts(t *testing.T) {
	fmt.Println("testing dhcpGethost")
	godhcpgethosts()
}

func TestStorageGet(t *testing.T) {
	fmt.Println("testing StorageGet")
	storageget()
}
func TestStoragePost(t *testing.T) {
	fmt.Println("testing StoragePOST")
	d := make(map[string]int)
	d["shit"] = 3
	d["fuck"] = 2
	storagepost(d)
}
func TestPostdhcphost(t *testing.T) {
	fmt.Println("testing postdhcphost")
	hostin := govirtlib.CreateVmForm{
		Hostname: "pdns-rec-7",
		VmMac:    "00:15:5d:32:02:4f",
		VmIp:     "192.168.38.202",
	}
	postdhcphost(hostin)
}
func TestPostdhcpnet(t *testing.T) {
	fmt.Println("testing postdhcpnet")
	hostin := govirtlib.Network{
		Subnet:  "10.181.35.0",
		Netmask: "255.255.255.0",
		Iprange: []string{"10.181.35.100", "10.181.35.200"},
		Dns:     []string{"10.181.35.10", "10.181.35.11"},
		Router:  "10.181.35.1",
	}
	postdhcpnet(hostin)
}
func TestDeldhcphost(t *testing.T) {
	fmt.Println("testing postdhcphost")
	hostin := govirtlib.HostInfo{
		Hostname: "pdns-rec-7",
		Mac:      "00:15:5d:32:02:4e",
		Ip:       "192.168.38.139",
	}
	deldhcphost(hostin)
}
