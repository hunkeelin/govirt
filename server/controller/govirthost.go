package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/klinutils"
	"github.com/hunkeelin/mtls/klinreq"
	"io/ioutil"
)

func (c *Conn) Getxml(vm, host string) ([]byte, error) {
	var r []byte
	p := &govirtlib.GetPayload{
		Target: "xml",
		Domain: vm,
	}
	i := &klinreq.ReqInfo{
		Dest:       host,
		Dport:      klinutils.Stringtoport("govirthost"),
		Method:     "GET",
		Payload:    p,
		TrustBytes: c.Tb,
		CertBytes:  c.Cb,
		KeyBytes:   c.Kb,
	}
	resp, err := klinreq.SendPayload(i)
	if err != nil {
		return r, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return r, err
	}
	resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println(string(body))
		return r, errors.New("Failed, check logs on the govirthost server" + host)
	}
	var tmpr govirtlib.ReturnPayload
	err = json.Unmarshal(body, &tmpr)
	if err != nil {
		return r, err
	}
	return tmpr.Xml, err
}
func (c *Conn) Statevm(state, vm, host string) error {
	p := &govirtlib.PostPayload{
		Action: state,
		Domain: vm,
	}
	i := &klinreq.ReqInfo{
		Dest:       host,
		Dport:      klinutils.Stringtoport("govirthost"),
		Method:     "POST",
		Payload:    p,
		TrustBytes: c.Tb,
		CertBytes:  c.Cb,
		KeyBytes:   c.Kb,
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
		return errors.New("Failed, check logs on the govirthost server")
	}
	return nil
}
func (c *Conn) migrate(ori, dest, vm string) error {
	p := &govirtlib.PostPayload{
		Action: "Migrate",
		Target: dest,
		Domain: vm,
	}
	i := &klinreq.ReqInfo{
		Dest:       ori,
		Dport:      klinutils.Stringtoport("govirthost"),
		Method:     "POST",
		Payload:    p,
		TrustBytes: c.Tb,
		CertBytes:  c.Cb,
		KeyBytes:   c.Kb,
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
		return errors.New("Failed, check logs on the govirthost server")
	}
	return nil
}
func (c *Conn) Getvms(hosts []string) ([]govirtlib.ReturnPayload, error) {
	var r []govirtlib.ReturnPayload
	p := &govirtlib.GetPayload{
		Target: "vm",
	}
	for _, host := range hosts {
		i := &klinreq.ReqInfo{
			Dest:       host,
			Dport:      klinutils.Stringtoport("govirthost"),
			Method:     "GET",
			Payload:    p,
			TrustBytes: c.Tb,
			CertBytes:  c.Cb,
			KeyBytes:   c.Kb,
		}
		resp, err := klinreq.SendPayload(i)
		if err != nil {
			return r, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return r, err
		}
		resp.Body.Close()
		if resp.StatusCode != 200 {
			fmt.Println(string(body))
			return r, errors.New("Failed, check logs on the govirthost server" + host)
		}
		var tmpr govirtlib.ReturnPayload
		err = json.Unmarshal(body, &tmpr)
		if err != nil {
			return r, err
		}
		tmpr.Parent = host
		r = append(r, tmpr)
	}
	return r, nil
}
func (c *Conn) Define(xml []byte, dest string) error {
	p := govirtlib.PostPayload{
		Action: "Define",
		Xml:    xml,
	}
	i := &klinreq.ReqInfo{
		Dest:       dest,
		Dport:      klinutils.Stringtoport("govirthost"),
		Method:     "POST",
		Payload:    p,
		TrustBytes: c.Tb,
		CertBytes:  c.Cb,
		KeyBytes:   c.Kb,
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
		return errors.New("Failed, check logs on the govirthost server")
	}
	return nil
}
