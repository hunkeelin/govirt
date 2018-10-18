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
		defer c.postMu.Unlock()
		err := start(p.Domain, c.L)
		if err != nil {
			return err
		}
	case "shutdown":
		c.postMu.Lock()
		defer c.postMu.Unlock()
		err = shutdown(p.Domain, c.L)
		if err != nil {

			return err
		}

	case "reset":
		c.postMu.Lock()
		defer c.postMu.Unlock()
		err = reset(p.Domain, c.L, true)
		if err != nil {

			return err
		}

	case "destroy":
		c.postMu.Lock()
		defer c.postMu.Unlock()
		err := destroy(p.Domain, c.L)
		if err != nil {

			return err
		}

	case "migrate":
		c.postMu.Lock()
		defer c.postMu.Unlock()
		err := migratev2(p.Domain, p.Target)
		if err != nil {

			return err
		}

	case "define":
		err := define(p.Xml, c.L)
		if err != nil {

			return err
		}

	case "undefine":
		c.postMu.Lock()
		defer c.postMu.Unlock()
		err := undefine(p.Domain, c.L)
		if err != nil {

			return err
		}
	default:
		return errors.New("Invalid Action")
	}
	return nil
}
