package godhcp

import (
	"fmt"
	"github.com/hunkeelin/mtls/klinserver"
	"net/http"
	"testing"
)

func TestNetmap(t *testing.T) {
	netmap, err := Getmapnet("dhcpd.conf")
	if err != nil {
		panic(err)
	}
	fmt.Println(netmap)
}
func TestServer(t *testing.T) {
	fmt.Println("Testing godhcp server")
	finish := make(chan bool)
	hostmap, err := Getmaphost("dhcpd-hosts.conf")
	if err != nil {
		panic(err)
	}
	netmap, err := Getmapnet("dhcpd.conf")
	if err != nil {
		panic(err)
	}
	cc := Conn{
		NetConfig:   "dhcpd.conf",
		HostConfig:  "dhcpd-hosts.conf",
		Hostmapinfo: hostmap,
		Netmapinfo:  netmap,
		ReserveIps:  []string{"192.168.38.200", "192.168.38.201"},
	}
	go func() {
		con := http.NewServeMux()
		con.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			cc.MainHandler(w, r)
		})
		j := &klinserver.ServerConfig{
			BindPort: "2020",
			ServeMux: con,
		}
		panic(klinserver.Server(j))
	}()
	<-finish
}
