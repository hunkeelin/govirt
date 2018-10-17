package main

import (
	"flag"
	"github.com/hunkeelin/SuperCAclient/lib"
	"github.com/hunkeelin/govirt/server/storagehost"
	"github.com/hunkeelin/klinutils"
	"github.com/hunkeelin/mtls/klinserver"
	"github.com/hunkeelin/pki"
	"net/http"
)

var (
	strloc        = flag.String("strloc", "/data/govirt/storage/", "the location of the storage")
	templateregex = flag.String("templateregex", "_template.img", "the regex for the template")
	insecure      = flag.Bool("insecure", false, "By default it uses mtls, this is for debug purpose. ")
	CA            = flag.String("CA", "superca", "The SuperCA server")
	CAport        = flag.String("CAport", "2018", "The SuperCA server port")
	Trust         = flag.String("Trust", "govirt", "Trust all the certs that's signed by this CA, read superca documentation")
	RequestCA     = flag.String("RequestCA", "govirt", "The requested CA to sign the csr when starting up the server")
)

func main() {
	flag.Parse()
	cc := storagehost.Conn{
		StorageLocation: *strloc,
		TemplateRegex:   *templateregex,
	}
	con := http.NewServeMux()
	con.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cc.MainHandler(w, r)
	})
	j := &klinserver.ServerConfig{
		BindPort: klinutils.Stringtoport("storagehost"),
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
