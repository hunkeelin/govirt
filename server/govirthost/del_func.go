package govirthost

import (
	"fmt"
	"github.com/digitalocean/go-libvirt"
)

func delvm(domain string, l *libvirt.Libvirt) error {
	var flags libvirt.DomainUndefineFlagsValues
	state, err := domainstate(domain, l)
	if err != nil {
		fmt.Println("unable to get domain state")
		return err
	}
	if state != "shutoff" {
		err = destroy(domain, l)
		if err != nil {
			fmt.Println("Unable to destroy vm")
			return err
		}
	}
	if err := l.Undefine(domain, flags); err != nil {
		return err
	}
	return nil
}
