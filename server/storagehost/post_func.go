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
	c.storageMu.Lock()
	defer c.storageMu.Unlock()
	for i := 1; i < 11; i++ {
		possible_dup_name := c.StorageLocation + image + "_dup_" + strconv.Itoa(i) + "_ready"
		_, err := os.Stat(possible_dup_name)
		if err == nil {
			os.Rename(possible_dup_name, c.StorageLocation+hostname+".qcow2")
			return nil
		}
	}
	go func() {
		dupimage, err := os.OpenFile(newimage, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Unable to create create image", err)
		}
		todup, err := ioutil.ReadFile(c.StorageLocation + image + "_template.img")
		if err != nil {
			fmt.Println("Unable to read file", c.StorageLocation+image+"_template.img", err)
		}
		dupimage.Write(todup)
		dupimage.Close()
	}()
	return errors.New("No dup image is avalible, copying from original, will take a while")
}
func (c *Conn) duplicate(lookup map[string]int) error {
	files, err := filepath.Glob(c.StorageLocation + "*" + c.TemplateRegex)
	if err != nil {
		fmt.Println("unable to get storage location", c.StorageLocation)
		return err
	}
	go func() {
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
				if klinutils.Exist(c.StorageLocation + image + "_dup_" + strconv.Itoa(j) + "_ready") {
					continue
				}
				dupimage, err := os.OpenFile(c.StorageLocation+image+"_dup_"+strconv.Itoa(j)+"copying", os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0644)
				if err != nil {
					fmt.Println("Unable to create duplicates ", c.StorageLocation+image+"_dup_"+string(j))
					continue
				}
				dupimage.Write(todup)
				dupimage.Close()
				c.storageMu.Lock()
				os.Rename(c.StorageLocation+image+"_dup_"+strconv.Itoa(j)+"copying", c.StorageLocation+image+"_dup_"+strconv.Itoa(j)+"_ready")
				c.storageMu.Unlock()
			}
			fmt.Println("finished duplicating image:", image)
		}
	}()
	return nil
}
