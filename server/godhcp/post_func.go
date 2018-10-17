package godhcp

import (
	"errors"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/klinutils"
)

func checkValid(h govirtlib.CreateVmForm) error {
	if h.Hostname == "" {
		return errors.New("Please specific hostname")
	}
	if !klinutils.Is_mac(h.VmMac) {
		return errors.New("Mac not valid")
	}
	if !klinutils.Is_ipv4(h.VmIp) {
		return errors.New("invalid Ip Address")
	}
	return nil
}
func (c *Conn) addhost(h govirtlib.CreateVmForm) error {
	err := checkValid(h)
	if err != nil {
		return err
	}
	for _, i := range c.Hostmapinfo {
		if i.Hostname == h.Hostname {
			return errors.New("host already exist in dhcp configuration " + i.Hostname)
		}
		if i.VmMac == h.VmMac {
			return errors.New("Mac is already in used by " + i.Hostname)
		}
		if i.VmIp == h.VmIp {
			return errors.New("Ip is already in used by " + i.Hostname)
		}
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
func checkValidnet(n govirtlib.Network) error {
	if !klinutils.Is_ipv4(n.Subnet) {
		return errors.New("invalid Subnet")
	}
	if !klinutils.Is_ipv4(n.Netmask) {
		return errors.New("invalid netmask")
	}
	for _, i := range n.Dns {
		if !klinutils.Is_ipv4(i) {
			return errors.New("invalid dns ip address")
		}
	}
	for _, i := range n.Iprange {
		if !klinutils.Is_ipv4(i) {
			return errors.New("invalid iprange ip address")
		}
	}
	if !klinutils.Is_ipv4(n.Router) {
		return errors.New("invalid Router ip address")
	}
	return nil
}
func (c *Conn) addnet(h govirtlib.Network) error {
	err := checkValidnet(h)
	if err != nil {
		return err
	}
	for i, _ := range c.Netmapinfo {
		if i == h.Subnet {
			return errors.New("Subnet already in used")
		}
	}
	c.Netmapinfo[h.Subnet] = h
	err = c.write("network")
	if err != nil {
		return err
	}
	return nil
}
