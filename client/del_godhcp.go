package govirtclient

import (
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/mtls/klinreq"
	"io/ioutil"
)

func deldhcphost(hostin govirtlib.HostInfo) {
	payload := &govirtlib.PostPayload{
		Target:   "host",
		Hostinfo: hostin,
	}
	i := &klinreq.ReqInfo{
		Dest:    "test2.klin-pro.com",
		Dport:   "2020",
		Method:  "DELETE",
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
