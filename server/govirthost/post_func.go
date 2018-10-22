package govirthost

import (
	"errors"
	"fmt"
	"github.com/digitalocean/go-libvirt"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/klinutils"
)

func shutdown(hostname string, l *libvirt.Libvirt) error {
	var flags libvirt.DomainShutdownFlagValues
	if err := l.Shutdown(hostname, flags); err != nil {
		return err
	}
	return nil
}
func reset(hostname string, l *libvirt.Libvirt, force bool) error {
	if force == true {
		if err := l.Reset(hostname); err != nil {
			return err
		}
	}
	return nil
}
func start(hostname string, l *libvirt.Libvirt) error {
	d, err := l.DomainLookupByName(hostname)
	if err != nil {
		return err
	}
	var flags uint32
	if _, err := l.DomainCreateWithFlags(d, flags); err != nil {
		return err
	}
	return nil
}
func create(f govirtlib.CreateVmForm, l *libvirt.Libvirt) error {
	_, err := l.DomainLookupByName(f.Hostname)
	if err == nil {
		return errors.New("host already exist: " + f.Hostname)
	}
	if f.Hostname == "" || !klinutils.Is_mac(f.VmMac) || f.CpuCount <= 0 || f.MemoryCount <= 1 || f.Image == "" {
		return errors.New("Errors with create vm form please check params")
	}
	return define(f.Xml, l)
}
func destroy(domain string, l *libvirt.Libvirt) error {
	var flags libvirt.DomainDestroyFlagsValues
	if err := l.Destroy(domain, flags); err != nil {
		return err
	}
	return nil
}
func define(f []byte, l *libvirt.Libvirt) error {
	var flags libvirt.DomainDefineFlags
	return l.DefineXML(f, flags)
}
func undefine(domain string, l *libvirt.Libvirt) error {
	var flags libvirt.DomainUndefineFlagsValues
	state, err := domainstate(domain, l)
	if err != nil {
		fmt.Println("unable to get domain state")
		return err
	}
	if state != "shutoff" {
		return errors.New("You cannot undefine a running vm, please shut it off first. The current state of " + domain + " is " + state)
	}
	if err := l.Undefine(domain, flags); err != nil {
		return err
	}
	return nil
}

func migratev2(domain, target string) error {
	cmd := "virsh"
	args := []string{"migrate", "--live", domain, "qemu+tls://" + target + "/system"}
	err := klinutils.Runshellv2(cmd, args)
	if err != nil {
		return errors.New("Unable to migrate " + domain + " to " + target)
	}
	return nil
}
func migrate(domain, target string, l *libvirt.Libvirt) error {
	if domain == "" || target == "" {
		return errors.New("Please specify domain and target")
	}
	var flags libvirt.DomainMigrateFlags
	flags = libvirt.MigrateLive |
		libvirt.MigratePeer2peer |
		libvirt.MigratePersistDest |
		libvirt.MigrateChangeProtection |
		libvirt.MigrateAbortOnError |
		libvirt.MigrateAutoConverge |
		libvirt.MigrateNonSharedDisk

	dom, err := l.DomainLookupByName(domain)
	if err != nil {
		fmt.Println("does", domain, "exist in this host?")
		return err
	}
	dconnuri := []string{"qemu+tls://" + target + "/system"}
	if _, err := l.DomainMigratePerform3Params(dom, dconnuri,
		[]libvirt.TypedParam{}, []byte{}, flags); err != nil {
		fmt.Println("Unable to migrate", domain, "to", target)
		return err
	}
	return nil
}
