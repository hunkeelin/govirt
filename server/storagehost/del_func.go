package storagehost

import (
	"os"
    "fmt"
	"path/filepath"
)

func (c *Conn) delhost(hostname string) error {
	err := os.Remove(c.StorageLocation + hostname + ".qcow2")
    if err != nil {
        fmt.Println("error while removing")
        return err
    }
    return nil
}
func (c *Conn) deltemplate(template string) error {
	matches, err := filepath.Glob(c.StorageLocation + template + "*" + c.TemplateRegex)
	if err != nil {
		return err
	}
	for _, i := range matches {
		err = os.Remove(i)
		if err != nil {
			return err
		}
	}
	return nil
}
