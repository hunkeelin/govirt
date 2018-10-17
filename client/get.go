package govirtclient

import (
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/mtls/klinreq"
	"io/ioutil"
)

func get() {
	payload := &govirtlib.GetPayload{
		Target: "vm",
	}
	i := &klinreq.ReqInfo{
		Dest:    "hades.klin-pro.com",
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
	for _, i := range p.Domains {
		fmt.Printf("%s\t%s%s", i.Domain.Name, "----> ", i.State)
		fmt.Println(i.Domain.UUID)
	}
	//	var jj []map[string]interface{}
	//	for _, i := range p.ReturnObj.([]interface{}) {
	//		j := i.(map[string]interface{})
	//		jj = append(jj, j)
	//	}
	//	for _, i := range jj {
	//		fmt.Println(i["state"])
	//	}
}
