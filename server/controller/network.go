package controller

import (
	"errors"
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/klinutils"
	"github.com/hunkeelin/mtls/klinreq"
	"io/ioutil"
)

func editnetwork(godhcp string, n govirtlib.Network) error {
	p := govirtlib.PostPayload{
		Target:  "network",
		Netinfo: n,
	}
	i := &klinreq.ReqInfo{
		Dest:    godhcp,
		Dport:   klinutils.Stringtoport("godhcp"),
		Method:  "PATCH",
		Payload: p,
		Http:    true,
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
func edithost(godhcp string, n govirtlib.CreateVmForm) error {
	p := govirtlib.PostPayload{
		Target: "host",
		VmForm: n,
	}
	i := &klinreq.ReqInfo{
		Dest:    godhcp,
		Dport:   klinutils.Stringtoport("godhcp"),
		Method:  "POST",
		Payload: p,
		Http:    true,
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
