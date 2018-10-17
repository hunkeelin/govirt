package storagehost

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
	case "dup":
		err := c.duplicate(p.DuplicateInfo)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
	case "setimage":
		err := c.setimage(p.VmForm.Image, p.VmForm.Hostname)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
	case "addstorage":
		err := c.addstorage(p.AddStrgInfo.Hostname, p.AddStrgInfo.Size)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
	default:
		return errors.New("Invalid Action")
	}
	return nil
}
