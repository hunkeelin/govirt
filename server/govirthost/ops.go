package govirthost

import (
	"errors"
	"github.com/digitalocean/go-libvirt"
	"github.com/hunkeelin/govirt/govirtlib"
)

func domainstate(hostname string, l *libvirt.Libvirt) (string, error) {
	gotState, err := l.DomainState(hostname)
	if err != nil {
		return "", err
	}
	switch gotState {
	case libvirt.DomainNostate:
		return "nostate", nil
	case libvirt.DomainRunning:
		return "running", nil
	case libvirt.DomainBlocked:
		return "blocked", nil
	case libvirt.DomainPaused:
		return "paused", nil
	case libvirt.DomainShutdown:
		return "shutdown", nil
	case libvirt.DomainShutoff:
		return "shutoff", nil
	case libvirt.DomainCrashed:
		return "crashed", nil
	case libvirt.DomainPmsuspended:
		return "suspended", nil
	default:
		return "", errors.New("Unable to get state")
	}
	return "", nil
}
func listdomains(l *libvirt.Libvirt) ([]govirtlib.DomainInfo, error) {
	var toReturn []govirtlib.DomainInfo
	domains, err := l.Domains()
	if err != nil {
		return toReturn, err
	}
	for _, domain := range domains {
		state, err := domainstate(domain.Name, l)
		if err != nil {
			return toReturn, errors.New("unable to get state for domain " + domain.Name)
		}
		vm := govirtlib.DomainInfo{
			Domain: domain,
			State:  state,
		}
		toReturn = append(toReturn, vm)
	}
	return toReturn, nil
}
