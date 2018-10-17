package govirthost

import (
	"fmt"
	"github.com/digitalocean/go-libvirt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/mtls/server"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestMigrate(t *testing.T) {
	err := migratev2("utemplate", "sf01-lab-netsrv-2.squaretrade.com")
	if err != nil {
		panic(err)
	}
	fmt.Println("migrate successful")
}
func TestDel(t *testing.T) {
	fmt.Println("testing del")
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		panic(err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		panic(err)
	}
	err = delvm("fuck", l)
	if err != nil {
		panic(err)
	}
	fmt.Println("finish")
}
func TestShutdown(t *testing.T) {
	fmt.Println("testing shutdown")
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		panic(err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		panic(err)
	}
	err = shutdown("template", l)
	if err != nil {
		panic(err)
	}
}
func TestXML(t *testing.T) {
	fmt.Println("testing virshdump")
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		panic(err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		panic(err)
	}
	var f libvirt.DomainXMLFlags
	xml, err := l.XML("template", f)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(string(xml))
}
func TestCreate(t *testing.T) {
	fmt.Println("testing create")
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		panic(err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		panic(err)
	}
	f := govirtlib.CreateVmForm{
		Hostname: "shit",
	}
	err = create(f, l)
	if err != nil {
		panic(err)
	}
}

//func TestHostMemory(t *testing.T) {
//	j := getHostMemory()
//	fmt.Println(j.FreeRam)
//}
func TestGetMemory(t *testing.T) {
	fmt.Println("testing get memory stat")
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		panic(err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		panic(err)
	}
	d, err := getMemory("template", l)
	if err != nil {
		panic(err)
	}
	for _, i := range d {
		if i.Tag == 4 {
			fmt.Println("Unused memory is ", i.Val, "kb")
		}
		if i.Tag == 5 {
			fmt.Println("Avalible memory is ", i.Val, "kb")
		}
	}
}
func TestServer(t *testing.T) {
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		panic(err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		panic(err)
	}
	cc := Conn{
		L: l,
	}
	con := http.NewServeMux()
	con.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cc.MainHandler(w, r)
	})
	j := &klinserver.ServerConfig{
		BindPort: "2020",
		ServeMux: con,
	}
	panic(klinserver.Server(j))
	if err := l.Disconnect(); err != nil {
		panic(err)
	}
}
