package godhcp

import (
	"errors"
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Conn) get(w http.ResponseWriter, r *http.Request) error {
	var p govirtlib.GetPayload
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
		err = c.getnetinfo(w)
		if err != nil {
			fmt.Println(err)
			return err
		}
	case "host":
		err = c.gethostinfo(w)
		if err != nil {
			fmt.Println(err)
			return err
		}
	default:
		return errors.New("Invalid Target please specify host or network")
	}
	return nil
}
