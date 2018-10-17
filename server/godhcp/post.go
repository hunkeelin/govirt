package godhcp

import (
	"errors"
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/klinutils"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Conn) post(w http.ResponseWriter, r *http.Request) error {
	var p govirtlib.PostPayload
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("unable to read response body post")
		return err
	}
	err = json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println("unable to unmarshal json post")
		return err
	}
	if klinutils.StringInSlice(p.VmForm.VmIp, c.ReserveIps) {
		return errors.New("Ip requested is in reserve ip list")
	}
	switch strings.ToLower(p.Target) {
	case "network":
		err = c.addnet(p.Netinfo)
		if err != nil {
			return err
		}
	case "host":
		err = c.addhost(p.VmForm)
		if err != nil {
			return err
		}
	default:
		return errors.New("Invalid Target Please specify host or network")
	}
	return nil
}
