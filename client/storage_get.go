package govirtclient

import (
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/mtls/klinreq"
	"io/ioutil"
)

func storageget() {
	payload := &govirtlib.GetPayload{
		Target: "images",
	}
	i := &klinreq.ReqInfo{
		Dest:    "hades.klin-pro.com",
		Dport:   "2021",
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
	for _, i := range p.Images {
		fmt.Println(i)
	}
}
