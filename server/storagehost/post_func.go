package storagehost

import (
	"errors"
	"fmt"
	"github.com/hunkeelin/klinutils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (c *Conn) addstorage(hostname string, sizeint int) error {
	cmd := "qemu-img"
	var id int
	size := strconv.Itoa(sizeint) + "G"
	imgloc := c.StorageLocation + hostname + "_storage_" + size
	for klinutils.Exist(imgloc + "_" + strconv.Itoa(id)) {
		id++
	}
	imgloc = imgloc + "_" + strconv.Itoa(id)
	args := []string{"create", "-f", "qcow2", imgloc, size}
	fmt.Println(args)
	return klinutils.Runshellv2(cmd, args)
}
func (c *Conn) setimage(image, hostname string) error {
	newimage := c.StorageLocation + hostname + ".qcow2"
	_, err := os.Stat(newimage)
	if err == nil {
		return errors.New(hostname + ".qcow2 already exist")
	}
	if _, err := os.Stat(c.StorageLocation + image + "_template.img"); os.IsNotExist(err) {
		return err
	}
	for i := 1; i < 11; i++ {
		possible_dup_name := c.StorageLocation + image + "_dup_" + strconv.Itoa(i) + "_ready"
		_, err := os.Stat(possible_dup_name)
		if err == nil {
			c.storageMu.Lock()
			os.Rename(possible_dup_name, c.StorageLocation+hostname+".qcow2")
			c.storageMu.Unlock()
			return nil
		}
	}
	dupimage, err := os.OpenFile(newimage, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Unable to create create image")
		return err
	}
	todup, err := ioutil.ReadFile(c.StorageLocation + image + "_template.img")
	if err != nil {
		fmt.Println("Unable to read file", c.StorageLocation+image+"_template.img")
		return err
	}
	dupimage.Write(todup)
	dupimage.Close()
	return nil
}
func (c *Conn) duplicate(lookup map[string]int) error {
	files, err := filepath.Glob(c.StorageLocation + "*" + c.TemplateRegex)
	if err != nil {
		fmt.Println("unable to get storage location", c.StorageLocation)
		return err
	}
	go func() {
		c.storageMu.Lock()
		for _, location := range files {
			tmpstring := strings.Replace(location, c.StorageLocation, "", -1)
			image := strings.Replace(tmpstring, c.TemplateRegex, "", -1)
			if lookup[image] == 0 {
				continue // don't read something you dn't need to read.
			}
			fmt.Println("duplicating image:", image)
			todup, err := ioutil.ReadFile(location)
			if err != nil {
				fmt.Println("Unable to read image ", location)
				continue
			}
			if lookup[image] > 10 {
				lookup[image] = 10
			}
			for j := 1; j < lookup[image]+1; j++ {
				dupimage, err := os.OpenFile(c.StorageLocation+image+"_dup_"+strconv.Itoa(j)+"copying", os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0644)
				if err != nil {
					fmt.Println("Unable to create duplicates ", c.StorageLocation+image+"_dup_"+string(j))
					continue
				}
				dupimage.Write(todup)
				dupimage.Close()
				os.Rename(c.StorageLocation+image+"_dup_"+strconv.Itoa(j)+"copying", c.StorageLocation+image+"_dup_"+strconv.Itoa(j)+"_ready")
			}
			fmt.Println("finished duplicating image:", image)
		}
		c.storageMu.Unlock()
	}()
	return nil
}
