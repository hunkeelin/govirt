package govirtclient

import (
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/mtls/klinreq"
	"io/ioutil"
	"strconv"
)

func godhcpget() {
	payload := &govirtlib.GetPayload{
		Target: "network",
	}
	i := &klinreq.ReqInfo{
		Dest:    "test2.klin-pro.com",
		Dport:   "2020",
		Method:  "GET",
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
	var p govirtlib.ReturnPayload
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}
	for _, i := range p.NetInfos {
		fmt.Println("Subnet ", i.Subnet)
		fmt.Println("Netmask ", i.Netmask)
		fmt.Println("Dns ", i.Dns)
		fmt.Println("Router ", i.Router)
		fmt.Println("Iprange ", i.Iprange)
	}
}
func godhcpgethosts() {
	payload := &govirtlib.GetPayload{
		Target: "host",
	}
	i := &klinreq.ReqInfo{
		Dest:    "test2.klin-pro.com",
		Dport:   "2020",
		Method:  "GET",
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
	var p govirtlib.ReturnPayload
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}
	for _, i := range p.HostInfos {
		fmt.Println("host:" + i.Hostname)
		fmt.Println("    - ip:" + i.VmIp)
		fmt.Println("    - mac:" + i.VmMac)
		fmt.Println("    - Leasetime:", strconv.Itoa(i.Leasetime))
	}
}
