package govirthost

import (
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"io/ioutil"
	"net/http"
)

func (c *Conn) del(w http.ResponseWriter, r *http.Request) error {
	var p govirtlib.PostPayload
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("unable to read response body del")
		return err
	}
	err = json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println("unable to unmarshal json del")
		return err
	}
	c.postMu.Lock()
	err = delvm(p.Domain, c.L)
	if err != nil {
		return err
	}
	c.postMu.Unlock()
	return nil
}
