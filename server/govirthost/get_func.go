package govirthost

import (
	"errors"
	"github.com/capnm/sysinfo"
	"github.com/digitalocean/go-libvirt"
	"github.com/hunkeelin/govirt/govirtlib"
	"net/http"
)

func getDomains(l *libvirt.Libvirt, w http.ResponseWriter) error {
	domains, err := listdomains(l)
	if err != nil {
		return errors.New("Unable to list domains")
	}
	p := govirtlib.ReturnPayload{
		Domains: domains,
		//		ReturnObj: domains,
	}
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		return errors.New("Unable to encode object to json")
	}
	return nil
}
func getHostMemory(w http.ResponseWriter) error {
	si := sysinfo.Get()
	p := govirtlib.ReturnPayload{
		HostMemoryInfo: si,
		//		ReturnObj:      si,
	}
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		return errors.New("Unable to encode object to json")
	}
	return nil
}
func getxml(hostname string, w http.ResponseWriter, l *libvirt.Libvirt) error {
	var f libvirt.DomainXMLFlags
	xml, err := l.XML(hostname, f)
	if err != nil {
		return err
	}
	p := govirtlib.ReturnPayload{
		Xml: xml,
	}
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		return errors.New("Unable to encode object to json")
	}
	return nil
}
func getMemory(host string, l *libvirt.Libvirt) ([]libvirt.DomainMemoryStat, error) {
	var domainMemStat []libvirt.DomainMemoryStat
	d, err := l.DomainLookupByName(host)
	if err != nil {
		return domainMemStat, err
	}

	domainMemStat, err = l.DomainMemoryStats(d, 20, 0)
	if err != nil {
		return domainMemStat, err
	}

	if len(domainMemStat) == 0 {
		return domainMemStat, errors.New("No memory return")
	}
	return domainMemStat, nil
}
