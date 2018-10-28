package controller

import (
    "github.com/hunkeelin/govirt/govirtlib"
    "strconv"
    "time"
    "github.com/hunkeelin/klinutils"
    "bytes"
    "math/rand"
    "errors"
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
func (c *Conn) CreateNewVm(cluster string, v govirtlib.CreateVmForm) error {
    err := checkVmForm(v)
    if err != nil {
        return err
    }
    if c.Ixml[v.Image] == nil {
        return errors.New("No image for : " + v.Image)
    }
    m, err := parse("config")
    if err != nil {
        panic(err)
    }
    err = c.edithost(m[cluster].Godhcp,v,false)
    if err != nil {
        panic(err)
    }
    err = c.setimage(m[cluster].Storage,v.Image,v.Hostname)
    if err != nil {
        panic(err)
    }
    xml := c.Ixml[v.Image]
    xml = bytes.Replace(xml, []byte("name_replace"), []byte(v.Hostname), -1)
    xml = bytes.Replace(xml, []byte("uuid_replace"), []byte(v.Uuid), -1)
    xml = bytes.Replace(xml, []byte("memory_replace"), []byte(strconv.Itoa(v.MemoryCount)), -1)
    xml = bytes.Replace(xml, []byte("cpu_replace"), []byte(strconv.Itoa(v.CpuCount)), -1)
    xml = bytes.Replace(xml, []byte("imagedir_replace"), []byte("/data/govirt/storage"), -1)
    xml = bytes.Replace(xml, []byte("mac_replace"), []byte(v.VmMac), -1)
    xml = bytes.Replace(xml, []byte("vlan_replace"), []byte(v.Vlan), -1)
    rand.Seed(time.Now().UTC().UnixNano())
    randhostint := randInt(0, len(m[cluster].Govirt))
    err = c.Define(xml,m[cluster].Govirt[randhostint])
    if err != nil {
        panic(err)
    }
    err = c.Statevm("start",v.Hostname,m[cluster].Govirt[randhostint])
    if err != nil {
        panic(err)
    }
    return nil
}
