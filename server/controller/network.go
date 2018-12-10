package controller

import (
	"errors"
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/klinutils"
	"github.com/hunkeelin/mtls/klinreq"
	"io/ioutil"
)

func (c *Conn)editnetwork(godhcp string, n govirtlib.Network, overwrite bool) error {
    method := "POST"
    if overwrite {
        method = "PATCH"
    }
	p := govirtlib.PostPayload{
		Target:  "network",
		Netinfo: n,
	}
	i := &klinreq.ReqInfo{
		Dest:    godhcp,
		Dport:   klinutils.Stringtoport("godhcp"),
		Method:  method,
		Payload: p,
        CertBytes: c.cb,
        KeyBytes: c.kb,
        TrustBytes: c.tb,
	}
	resp, err := klinreq.SendPayload(i)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println(string(body))
		return errors.New("Failed, check logs on the godhcp server")
	}
	return nil
}
func (c *Conn) edithost(godhcp string, n govirtlib.PostPayload,overwrite bool) error {
    method := "POST"
    if overwrite {
        method = "PATCH"
    }
	p := govirtlib.PostPayload{
		Target: "host",
		VmForm: n.VmForm,
	}
	i := &klinreq.ReqInfo{
		Dest:    godhcp,
		Dport:   klinutils.Stringtoport("godhcp"),
		Method:  method,
		Payload: p,
        CertBytes: c.cb,
        KeyBytes: c.kb,
        TrustBytes: c.tb,
	}
	resp, err := klinreq.SendPayload(i)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println(string(body))
		return errors.New("Failed, check logs on the godhcp server")
	}
	return nil
}
func (c *Conn) delhost_network(godhcp string, host string) error {
    method := "DELETE"
	p := govirtlib.PostPayload{
		Target: "host",
        Domain: host,
	}
	i := &klinreq.ReqInfo{
		Dest:    godhcp,
		Dport:   klinutils.Stringtoport("godhcp"),
		Method:  method,
		Payload: p,
        CertBytes: c.cb,
        KeyBytes: c.kb,
        TrustBytes: c.tb,
	}
	resp, err := klinreq.SendPayload(i)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println(string(body))
		return errors.New("Failed, check logs on the godhcp server")
	}
	return nil
}
