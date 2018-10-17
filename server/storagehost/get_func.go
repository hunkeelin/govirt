package storagehost

import (
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"net/http"
	"path/filepath"
	"strings"
)

func (c *Conn) getImages(w http.ResponseWriter, r *http.Request) error {
	files, err := filepath.Glob(c.StorageLocation + "*" + c.TemplateRegex)
	if err != nil {
		fmt.Println("unable to get storage location", c.StorageLocation)
		return err
	}
	for i, location := range files {
		tmpstring := strings.Replace(location, c.StorageLocation, "", -1)
		files[i] = strings.Replace(tmpstring, c.TemplateRegex, "", -1)
	}
	p := govirtlib.ReturnPayload{
		Images: files,
	}
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		fmt.Println("unable to encode json")
		return err
	}
	return nil
}
