package godhcp

import (
	"errors"
	"fmt"
	"github.com/hunkeelin/klinutils"
	"os"
	"strconv"
	"strings"
)

func (c *Conn) write(t string) error {
	if t == "host" {
		err := c.writeHostConfig()
		if err != nil {
			return err
		}
		cmd := "systemctl"
		args := []string{"restart", "dhcpd"}
		err = klinutils.Runshellv2(cmd, args)
		if err != nil {
			fmt.Println("unable to restart dhcpd, encountered some errors")
			return err
		}
		return nil
	}
	if t == "network" {
		err := c.writeNetConfig()
		if err != nil {
			return err
		}
		cmd := "systemctl"
		args := []string{"restart", "dhcpd"}
		err = klinutils.Runshellv2(cmd, args)
		if err != nil {
			fmt.Println("unable to restart dhcpd, encountered some errors")
			return err
		}
		return nil
	}
	return errors.New("Please specify a proper write parameter")
}

func (c *Conn) writeHostConfig() error {
	c.hostConfigMu.Lock()
	target := c.HostConfig + "_tmp"
	f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return errors.New("Server Error check server logs")
	}
	for _, i := range c.Hostmapinfo {
		var leasestring, text string
		if i.Leasetime < 1 {
			leasestring = ""
		} else {
			leasetime := strconv.Itoa(i.Leasetime)
			leasestring = "    max-lease-time " + leasetime + ";\n" + "    default-lease-time " + leasetime + ";\n"
		}
		text = "host " + i.Hostname + " {\n" + "    hardware ethernet " + i.VmMac + ";\n" + "    fixed-address " + i.VmIp + ";\n"
		text = text + leasestring + "}\n"
		if _, err = f.WriteString(text); err != nil {
			fmt.Println(err)
			return errors.New("Server Error check server logs")
		}
	}
	f.Close()
	os.Rename(target, c.HostConfig)
	c.hostConfigMu.Unlock()
	return nil
}
func (c *Conn) writeNetConfig() error {
	c.hostConfigMu.Lock()
	target := c.NetConfig + "_tmp"
	f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return errors.New("Server Error check server logs")
	}
	text := "authoritative;\n" + "include \"" + c.HostConfig + "\";\n"
	if _, err = f.WriteString(text); err != nil {
		fmt.Println(err)
		return errors.New("Server Error check server logs")
	}
	for _, i := range c.Netmapinfo {
		var dnsstr, iprangestr string
		if len(i.Dns) > 1 { // check if there is even an option for it.
			dnsstr = strings.Join(i.Dns, ",")
		}
		iprangestr = strings.Join(i.Iprange, " ")
		text = "subnet " + i.Subnet + " netmask " + i.Netmask + " {\n" + "    option domain-name-servers    " + dnsstr + ";\n" + "    range " + iprangestr + ";\n" + "    option routers " + i.Router + ";\n" + "    default-lease-time " + i.Lease + ";\n" + "    max-lease-time " + i.Maxlease + ";\n" + "}\n"
		if _, err = f.WriteString(text); err != nil {
			fmt.Println(err)
			return errors.New("Server Error check server logs")
		}
	}
	f.Close()
	os.Rename(target, c.NetConfig)
	c.hostConfigMu.Unlock()
	return nil
}
