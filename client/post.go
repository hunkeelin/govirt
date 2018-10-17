package govirtclient

import (
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/mtls/klinreq"
	"io/ioutil"
)

func post() {
	payload := &govirtlib.PostPayload{
		Domain: "template",
		Action: "shutdown",
	}
	i := &klinreq.ReqInfo{
		Dest:    "hades.klin-pro.com",
		Dport:   "2020",
		Method:  "POST",
		Payload: payload,
		Http:    true,
	}
	resp, err := klinreq.SendPayload(i)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body), resp.StatusCode)
}
func postdhcphost(hostin govirtlib.CreateVmForm) {
	payload := &govirtlib.PostPayload{
		Target: "host",
		VmForm: hostin,
	}
	i := &klinreq.ReqInfo{
		Dest:    "test2.klin-pro.com",
		Dport:   "2020",
		Method:  "POST",
		Payload: payload,
		Http:    true,
	}
	resp, err := klinreq.SendPayload(i)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body), resp.StatusCode)
}
func postdhcpnet(hostin govirtlib.Network) {
	payload := &govirtlib.PostPayload{
		Target:  "network",
		Netinfo: hostin,
	}
	i := &klinreq.ReqInfo{
		Dest:    "test3.klin-pro.com",
		Dport:   "2020",
		Method:  "POST",
		Payload: payload,
		Http:    true,
	}
	resp, err := klinreq.SendPayload(i)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body), resp.StatusCode)
}
