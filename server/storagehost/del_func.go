package storagehost

import (
	"os"
	"path/filepath"
)

func (c *Conn) delhost(hostname string) error {
	return os.Remove(c.StorageLocation + hostname + ".qcow2")
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
