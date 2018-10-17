package godhcp

import (
	"errors"
)

func (c *Conn) delhost(h string) error {
	if h == "" {
		return errors.New("Please specify hostname")
	}
	delete(c.Hostmapinfo, h)
	err := c.write("host")
	if err != nil {
		return err
	}
	return nil
}
func (c *Conn) delnet(h string) error {
	if h == "" {
		return errors.New("Please specify subnet")
	}
	delete(c.Netmapinfo, h)
	err := c.write("network")
	if err != nil {
		return err
	}
	return nil
}
