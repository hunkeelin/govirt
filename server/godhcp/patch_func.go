package godhcp

import (
	"github.com/hunkeelin/govirt/govirtlib"
)

func (c *Conn) patchhost(h govirtlib.CreateVmForm) error {
	err := checkValid(h)
	if err != nil {
		return err
	}
	if h.Leasetime == 0 {
		h.Leasetime = 16000
	}
	c.Hostmapinfo[h.Hostname] = h
	err = c.write("host")
	if err != nil {
		return err
	}
	return nil
}
func (c *Conn) patchnet(h govirtlib.Network) error {
	err := checkValidnet(h)
	if err != nil {
		return err
	}
	c.Netmapinfo[h.Subnet] = h
	err = c.write("network")
	if err != nil {
		return err
	}
	return nil
}
