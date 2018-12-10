package govirtlib

import (
	"github.com/capnm/sysinfo"
	"github.com/digitalocean/go-libvirt"
)

type DomainInfo struct {
	Domain libvirt.Domain `json:"domain"`
	State  string         `json:"state"`
}
type CreateVmForm struct {
	Hostname    string `json:"hostname"`
	VmMac       string `json:"mac"` // the mac addressA
	Uuid        string `json:"uuid"`
	VmIp        string `json:"ip"`          // the ip address of the virutal machine
	CpuCount    int    `json:"cpucount"`    // cpucount
	MemoryCount int    `json:"memorycount"` // in GB
	Image       string `json:"image"`       // the image to use
	Xml         []byte `json:"xml"`         // the xml in bytes
	Leasetime   int    `json:"leasetime"`   // this is for godhcp
	Vlan        string `json:"vlan"`        // the vlan
}
type PostPayload struct {
	Domain        string         `json:"domain"` // the domain to action should apply to :govirthost
	Action        string         `json:"action"` // what action right now only create, reset. shutdown, start. :All
	Target        string         `json:"target"`
	VmForm        CreateVmForm   `json:"createvmform"`  // the object for the govirthost that contains vminfo to create the virtual machine :govirthost
	DuplicateInfo map[string]int `json:"duplicateinfo"` // for storage, string is name of the image and the int is how many copy of the image :storagehost
	//	Hostinfo      HostInfo       `json:"hostinfo"`      // the hostinfo to add to godhcp. :godhcp
	Netinfo     Network  `json:"netinfo"` // when you want to add or delete network for. :godhcp
	AddStrgInfo StrgInfo `json:"storageinfo"`
	Xml         []byte   `json:"xml"`
	Cluster     string   `json:"cluster"` // the name of the cluster
}
type StrgInfo struct {
	Hostname string `json:"hostname"`
	Size     int    `json:"size"`
}
type GetPayload struct {
	Target string `json:"target"` // whether you are targeting vm or the host
	Domain string `json:"domain"`
}
type Network struct { // for godhcp
	Subnet   string   `json:"subnet"`
	Netmask  string   `json:"netmask"`
	Dns      []string `json:"dns"`
	Router   string   `json:"router"`
	Iprange  []string `json:"iprange"`
	Lease    string   `json:"lease"`
	Maxlease string   `json:"maxlease"`
}

//type HostInfo struct {
//	Hostname  string `json:"hostname"`
//	Mac       string `json:"mac"`
//	Ip        string `json:"ip"`
//	Leasetime int    `json:"leasetime"`
//}
type AddHostDhcpForm struct {
	hostname string `json:"hostname"`
	mac      string `json:"mac"`
}
type ReturnPayload struct {
	Images         []string                `json:"images"`
	HostMemoryInfo *sysinfo.SI             `json:"hostmemoryInfo"`
	Domains        []DomainInfo            `json:"domains"`
	Networks       []Network               `json:"network"`
	HostInfos      map[string]CreateVmForm `json:"hostinfos"`
	NetInfos       map[string]Network      `json:"netinfo"`
	Xml            []byte                  `json:"xml"`
	Parent         string                  `json:"parent"`
	listvms        map[string][]DomainInfo `json:"listvms"`
}
