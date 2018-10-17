package storagehost

import (
	"fmt"
	"github.com/hunkeelin/mtls/server"
	"net/http"
	"testing"
)

func TestDup(t *testing.T) {
	c := Conn{
		StorageLocation: "/home/bgops/files/golesson/govirt/server/storagehost/teststorage/",
		TemplateRegex:   "_template.img",
		Config:          "config",
	}
	m := map[string]int{
		"sfhit": 4,
	}
	err := c.duplicate(m)
	if err != nil {
		panic(err)
	}
}
func TestAddstorage(t *testing.T) {
	fmt.Println("testing add storage")
	c := Conn{
		StorageLocation: "/home/bgops/files/golesson/govirt/server/storagehost/teststorage/",
		TemplateRegex:   "_template.img",
		Config:          "config",
	}
	err := c.addstorage("noobshit", 5)
	if err != nil {
		panic(err)
	}
}
func TestServer(t *testing.T) {
	fmt.Println("Testing Storage server")
	finish := make(chan bool)
	cc := Conn{
		StorageLocation: "/home/bgops/files/golesson/govirt/server/storagehost/teststorage/",
		TemplateRegex:   "_template.img",
	}
	go func() {
		con := http.NewServeMux()
		con.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			cc.MainHandler(w, r)
		})
		j := &klinserver.ServerConfig{
			BindPort: "2021",
			ServeMux: con,
		}
		panic(klinserver.Server(j))
	}()
	//	go func() {
	//		panic(cc.duplicate())
	//	}()
	<-finish
}
