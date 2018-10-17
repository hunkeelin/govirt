package godhcp

import (
	"errors"
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Conn) del(w http.ResponseWriter, r *http.Request) error {
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
	switch strings.ToLower(p.Target) {
	case "network":
		println("deleting ", p.Netinfo.Subnet)
		err = c.delnet(p.Netinfo.Subnet)
		if err != nil {
			return err
		}
	case "host":
		err = c.delhost(p.Domain)
		if err != nil {
			return err
		}
	default:
		return errors.New("Invalid Target Please specify host or network")
	}
	return nil
}
