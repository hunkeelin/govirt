package govirthost

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
		fmt.Println("unable to read response body del")
		return err
	}
	err = json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println("unable to unmarshal json get")
		return err
	}
	switch strings.ToLower(p.Target) {
	case "vm":
		err = getDomains(c.L, w)
		if err != nil {
			fmt.Println("Unable to get domains")
			return err
		}
	case "metal":
		err = getHostMemory(w)
		if err != nil {
			fmt.Println("Unable to get Memory")
			return err
		}
	case "xml":
		err = getxml(p.Domain, w, c.L)
		if err != nil {
			fmt.Println("Unable to get xml")
			return err
		}
	default:
		return errors.New("Invalid Action")
	}
	return nil
}
