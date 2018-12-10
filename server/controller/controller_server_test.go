package controller

import (
    "github.com/hunkeelin/SuperCAclient/lib"
    "testing"
    "github.com/hunkeelin/pki"
    "github.com/hunkeelin/klinutils"
)
func TestServer(t *testing.T){
    c := Conn{}
    var err error
    r := klinutils.WgetInfo{
        Dest:  "ec2-superca-prod-1.squaretrade.com",
        Dport: "2018",
        Route: "cacerts/rootca.crt",
    }
    r.Route = "cacerts/superauth.crt"
    c.authtb, err = klinutils.Wget(r)
    if err != nil {
        panic(err)
    }
    w := client.WriteInfo{
        CAName:  "ec2-superca-prod-1.squaretrade.com",
        CAport:  "2018",
        Chain:   true,
        CSRConfig: &klinpki.CSRConfig{
            RsaBits: 2048,
        },
        SignCA: "nginx",
    }
    w.CABytes, err = klinutils.Wget(r)
    c.Cb, c.Kb, err = client.Getkeycrtbyte(w)
    if err != nil {
        panic(err)
    }
    w.SignCA = "superauth"
    c.authcb,c.authkb,err = client.Getkeycrtbyte(w)
}
