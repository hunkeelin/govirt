package govirtclient

import (
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/mtls/klinreq"
	"io/ioutil"
)

func storagepost(d map[string]int) {
	payload := &govirtlib.PostPayload{
		Action:        "dup",
		DuplicateInfo: d,
	}
	i := &klinreq.ReqInfo{
		Dest:    "test1.klin-pro.com",
		Dport:   "2021",
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
