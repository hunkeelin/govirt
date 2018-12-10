package controller

import (
	"bytes"
	"errors"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/klinutils"
	"math/rand"
	"strconv"
	"time"
)

func (c *Conn) Migratehost(ori, dest, vm string) error {
	err := c.migrate(ori, dest, vm)
	if err != nil {
		return err
	}
	xml, err := c.Getxml(vm, ori)
	if err != nil {
		return err
	}
	err = c.Statevm("unDefine", vm, ori)
	if err != nil {
		return err
	}
	err = c.Define(xml, dest)
	if err != nil {
		return err
	}
	return nil
}
func checkVmForm(v govirtlib.CreateVmForm) error {
	switch {
	case v.Hostname == "":
		return errors.New("Please specify hostname")
	case v.Uuid == "":
		return errors.New("Please specifiy uuid")
	case v.MemoryCount == 0:
		return errors.New("Please specify memorycount")
	case v.CpuCount == 0:
		return errors.New("Please specify memorycount")
	case !klinutils.Is_mac(v.VmMac):
		return errors.New("Please specify a valid mac address")
	case v.Vlan == "":
		return errors.New("Please specify Vlan for network")
	case !klinutils.Is_ipv4(v.VmIp):
		return errors.New("Please speceify a valid IP")
	default:
		return nil
	}
	return nil
}
func (c *Conn) CreateNewVm(v govirtlib.PostPayload) error {
	err := checkVmForm(v.VmForm)
	if err != nil {
		return err
	}
	if c.Ixml[v.VmForm.Image] == nil {
		return errors.New("No image for : " + v.VmForm.Image)
	}
	m, err := Parse("config")
	if err != nil {
		panic(err)
	}
	err = c.edithost(m[v.Cluster].Godhcp, v, false)
	if err != nil {
		panic(err)
	}
	err = c.setimage(m[v.Cluster].Storage, v.VmForm.Image, v.VmForm.Hostname)
	if err != nil {
		panic(err)
	}
	xml := c.Ixml[v.VmForm.Image]
	xml = bytes.Replace(xml, []byte("name_replace"), []byte(v.VmForm.Hostname), -1)
	xml = bytes.Replace(xml, []byte("uuid_replace"), []byte(v.VmForm.Uuid), -1)
	xml = bytes.Replace(xml, []byte("memory_replace"), []byte(strconv.Itoa(v.VmForm.MemoryCount)), -1)
	xml = bytes.Replace(xml, []byte("cpu_replace"), []byte(strconv.Itoa(v.VmForm.CpuCount)), -1)
	xml = bytes.Replace(xml, []byte("imagedir_replace"), []byte("/data/govirt/storage"), -1)
	xml = bytes.Replace(xml, []byte("mac_replace"), []byte(v.VmForm.VmMac), -1)
	xml = bytes.Replace(xml, []byte("vlan_replace"), []byte(v.VmForm.Vlan), -1)
	rand.Seed(time.Now().UTC().UnixNano())
	randhostint := randInt(0, len(m[v.Cluster].Govirt))
	err = c.Define(xml, m[v.Cluster].Govirt[randhostint])
	if err != nil {
		panic(err)
	}
	err = c.Statevm("start", v.VmForm.Hostname, m[v.Cluster].Govirt[randhostint])
	if err != nil {
		panic(err)
	}
	return nil
}
