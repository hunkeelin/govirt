package govirthost

import (
	"errors"
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
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
	switch strings.ToLower(p.Action) {
	case "start":
		c.postMu.Lock()
		err := start(p.Domain, c.L)
		if err != nil {
			return err
		}
		c.postMu.Unlock()
	case "shutdown":
		c.postMu.Lock()
		err = shutdown(p.Domain, c.L)
		if err != nil {
			return err
		}
		c.postMu.Unlock()
	case "reset":
		c.postMu.Lock()
		err = reset(p.Domain, c.L, true)
		if err != nil {
			return err
		}
		c.postMu.Unlock()
	case "create":
		c.postMu.Lock()
		err := create(p.VmForm, c.L)
		if err != nil {
			return err
		}
		c.postMu.Unlock()
	case "destroy":
		c.postMu.Lock()
		err := destroy(p.Domain, c.L)
		if err != nil {
			return err
		}
		c.postMu.Unlock()
	case "migrate":
		c.postMu.Lock()
		err := migratev2(p.Domain, p.Target)
		if err != nil {
			return err
		}
		c.postMu.Unlock()
	case "define":
		c.postMu.Lock()
		err := define(p.Xml, c.L)
		if err != nil {
			return err
		}
		c.postMu.Unlock()
	case "undefine":
		c.postMu.Lock()
		err := undefine(p.Domain, c.L)
		if err != nil {
			return err
		}
		c.postMu.Unlock()
	default:
		return errors.New("Invalid Action")
	}
	return nil
}
