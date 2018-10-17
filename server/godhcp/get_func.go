package godhcp

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/hunkeelin/govirt/govirtlib"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func Getmaphost(config string) (map[string]govirtlib.CreateVmForm, error) {
	hostinfomap := make(map[string]govirtlib.CreateVmForm)
	file, err := ioutil.ReadFile(config)
	if err != nil {
		return hostinfomap, err
	}
	bytelines := bytes.Split(file, []byte("\n"))
	hostname := regexp.MustCompile("host (.*) {")
	mac := regexp.MustCompile("hardware ethernet (.*);")
	ip := regexp.MustCompile("fixed-address (.*);")
	leasetime := regexp.MustCompile("default-lease-time (.*);")
	var currenthost string
	for _, i := range bytelines {
		var hostinfoobj govirtlib.CreateVmForm
		hostmatch := hostname.FindStringSubmatch(string(i))
		if len(hostmatch) > 1 {
			currenthost = hostmatch[1]
			hostinfoobj.Hostname = hostmatch[1]
			hostinfomap[currenthost] = hostinfoobj
		}
		macmatch := mac.FindStringSubmatch(string(i))
		if len(macmatch) > 1 {
			tmp := hostinfomap[currenthost]
			tmp.VmMac = macmatch[1]
			hostinfomap[currenthost] = tmp
		}
		ipmatch := ip.FindStringSubmatch(string(i))
		if len(ipmatch) > 1 {
			tmp := hostinfomap[currenthost]
			tmp.VmIp = ipmatch[1]
			hostinfomap[currenthost] = tmp
		}
		routermatch := leasetime.FindStringSubmatch(string(i))
		if len(routermatch) > 1 {
			tmp := hostinfomap[currenthost]
			leaseint, err := strconv.Atoi(routermatch[1])
			if err != nil {
				return hostinfomap, err
			}
			tmp.Leasetime = leaseint
			hostinfomap[currenthost] = tmp
		}
	}
	return hostinfomap, nil
}
func (c *Conn) gethostinfo(w http.ResponseWriter) error {
	returnPayload := govirtlib.ReturnPayload{
		HostInfos: c.Hostmapinfo,
	}
	err := json.NewEncoder(w).Encode(returnPayload)
	if err != nil {
		fmt.Println("unable to encode json")
		return err
	}
	return nil
}
func (c *Conn) getnetinfo(w http.ResponseWriter) error {
	returnPayload := govirtlib.ReturnPayload{
		NetInfos: c.Netmapinfo,
	}
	err := json.NewEncoder(w).Encode(returnPayload)
	if err != nil {
		fmt.Println("unable to encode json")
		return err
	}
	return nil
}
func Getmapnet(config string) (map[string]govirtlib.Network, error) {
	netmap := make(map[string]govirtlib.Network)
	file, err := ioutil.ReadFile(config)
	if err != nil {
		return netmap, err
	}
	bytelines := bytes.Split(file, []byte("\n"))
	subnetre := regexp.MustCompile("subnet (.*)\\sn")
	netmaskre := regexp.MustCompile("netmask (.*)\\s")
	dnsre := regexp.MustCompile("domain-name-servers[\\s|\\t]{1,}(.*);")
	routerre := regexp.MustCompile("routers[\\s|\\t]{1,}(.*);")
	iprangere := regexp.MustCompile("range[\\s|\\t]{1,}(.*);")
	leasere := regexp.MustCompile("default-lease-time[\\s|\\t]{1,}(.*);")
	maxleasere := regexp.MustCompile("max-lease-time[\\s|\\t]{1,}(.*);")
	var currentnet string
	for _, i := range bytelines {
		var netobj govirtlib.Network
		subnetmatch := subnetre.FindStringSubmatch(string(i))
		if len(subnetmatch) > 1 {
			currentnet = subnetmatch[1]
			netobj.Subnet = subnetmatch[1]
			netmap[currentnet] = netobj
		}
		netmaskmatch := netmaskre.FindStringSubmatch(string(i))
		if len(netmaskmatch) > 1 {
			tmp := netmap[currentnet]
			tmp.Netmask = netmaskmatch[1]
			netmap[currentnet] = tmp
		}
		dnsmatch := dnsre.FindStringSubmatch(string(i))
		if len(dnsmatch) > 1 {
			tmp := netmap[currentnet]
			tmp.Dns = strings.Split(dnsmatch[1], ",")
			netmap[currentnet] = tmp
		}
		routermatch := routerre.FindStringSubmatch(string(i))
		if len(routermatch) > 1 {
			tmp := netmap[currentnet]
			tmp.Router = routermatch[1]
			netmap[currentnet] = tmp
		}
		iprangematch := iprangere.FindStringSubmatch(string(i))
		if len(iprangematch) > 1 {
			tmp := netmap[currentnet]
			tmprange := strings.Split(iprangematch[1], " ")
			if len(tmprange) != 2 {
				return netmap, errors.New("syntax error at range")
			}
			tmp.Iprange = tmprange
			netmap[currentnet] = tmp
		}
		leasematch := leasere.FindStringSubmatch(string(i))
		if len(leasematch) > 1 {
			tmp := netmap[currentnet]
			tmp.Lease = leasematch[1]
			netmap[currentnet] = tmp
		}
		maxleasematch := maxleasere.FindStringSubmatch(string(i))
		if len(maxleasematch) > 1 {
			tmp := netmap[currentnet]
			tmp.Maxlease = maxleasematch[1]
			netmap[currentnet] = tmp
		}
	}
	return netmap, nil
}
