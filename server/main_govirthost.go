package main

import (
	"flag"
	"github.com/digitalocean/go-libvirt"
	"github.com/hunkeelin/SuperCAclient/lib"
	"github.com/hunkeelin/govirt/server/govirthost"
	"github.com/hunkeelin/klinutils"
	"github.com/hunkeelin/mtls/klinserver"
	"github.com/hunkeelin/pki"
	"net"
	"net/http"
	"time"
)

var (
	insecure  = flag.Bool("insecure", false, "By default it uses mtls, this is for debug purpose. ")
	CA        = flag.String("CA", "superca", "The SuperCA server")
	CAport    = flag.String("CAport", "2018", "The SuperCA server port")
	Trust     = flag.String("Trust", "govirt", "Trust all the certs that's signed by this CA, read superca documentation")
	RequestCA = flag.String("RequestCA", "govirt", "The requested CA to sign the csr when starting up the server")
)

func main() {
	flag.Parse()
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		panic(err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		panic(err)
	}
	cc := govirthost.Conn{
		L: l,
	}
	con := http.NewServeMux()
	con.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cc.MainHandler(w, r)
	})
	j := &klinserver.ServerConfig{
		BindPort: klinutils.Stringtoport("govirthost"),
		ServeMux: con,
	}
	if !*insecure {
		r := klinutils.WgetInfo{
			Dest:  *CA,
			Dport: *CAport,
			Route: "cacerts/rootca.crt",
		}
		cab, err := klinutils.Wget(r)
		if err != nil {
			panic(err)
		}
		r = klinutils.WgetInfo{
			Dest:  *CA,
			Dport: *CAport,
			Route: "cacerts/" + *Trust + ".crt",
		}
		cc.Trustbytes, err = klinutils.Wget(r)
		if err != nil {
			panic(err)
		}
		w := client.WriteInfo{
			CAName:  *CA,
			CABytes: cab,
			CAport:  *CAport,
			Chain:   true,
			CSRConfig: &klinpki.CSRConfig{
				RsaBits: 2048,
			},
			SignCA: *RequestCA,
		}
		cc.Certbytes, cc.Keybytes, err = client.Getkeycrtbyte(w)
		if err != nil {
			panic(err)
		}
		j.CertBytes = cc.Certbytes
		j.KeyBytes = cc.Keybytes
		j.Https = true
		j.Verify = true
		j.TrustBytes = cc.Trustbytes
	}
	panic(klinserver.Server(j))
}
