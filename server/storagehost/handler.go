package storagehost

import (
	"fmt"
	"github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (c *Conn) MainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		err := c.get(w, r)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
	case "POST":
		err := c.post(w, r)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
	case "DELETE":
		err := c.del(w, r)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
	default:
		fmt.Println("invalid method")
		w.WriteHeader(500)
	}
}
