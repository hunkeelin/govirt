package storagehost

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
	switch strings.ToLower(p.Action) {
	case "host":
		err := c.delhost(p.Target)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
	case "image":
		err := c.deltemplate(p.Target)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
	default:
		return errors.New("Invalid Action")
	}
	return nil
}
