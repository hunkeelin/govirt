package controller

import (
	"fmt"
    "bytes"
	"github.com/hunkeelin/govirt/govirtlib"
    "testing"
    "github.com/hunkeelin/klinutils"
    "io/ioutil"
)

func TestParse(t *testing.T) {
	m, err := parse("config")
	if err != nil {
		panic(err)
	}
	for _, i := range m {
		fmt.Println(i.ClusterName)
		fmt.Println(i.Godhcp)
		fmt.Println(i.Govirt)
		fmt.Println(i.Storage)
	}
}
func TestEditNet(t *testing.T) {
	fmt.Println("testing patch network")
	n := govirtlib.Network{
		Subnet:   "10.180.250.0",
		Netmask:  "255.255.254.0",
		Dns:      []string{"10.181.35.100", "10.181.35.101"},
		Router:   "10.180.250.1",
		Iprange:  []string{"10.180.250.160", "10.180.250.200"},
		Lease:    "3600",
		Maxlease: "7200",
	}
	m, err := parse("config")
	if err != nil {
		panic(err)
	}
	err = editnetwork(m["sf_deploy"].Godhcp, n,true)
	if err != nil {
		panic(err)
	}
}
func TestCreatevm(t *testing.T){
    fmt.Println("testing create vm")
    mac, err := klinutils.Genmac()
    if err != nil {
        panic(err)
    }
    uuid, err := klinutils.Genuuid()
    if err != nil {
        panic(err)
    }
    template, err := ioutil.ReadFile("ctemplate.xml")
    if err != nil {
        panic(err)
    }
    template = bytes.Replace(template,[]byte("name_replace"),[]byte("centostest"),-1)
    template = bytes.Replace(template,[]byte("uuid_replace"),uuid,-1)
    template = bytes.Replace(template,[]byte("memory_replace"),[]byte("4"),-1)
    template = bytes.Replace(template,[]byte("cpu_replace"),[]byte("2"),-1)
    template = bytes.Replace(template,[]byte("imagedir_replace"),[]byte("/data/govirt/storage"),-1)
    template = bytes.Replace(template,[]byte("mac_replace"),mac,-1)
    template = bytes.Replace(template,[]byte("vlan_replace"),[]byte("govirtmgmt"),-1)
    fmt.Println(string(template))
    err = createvm(template,"sf01-lab-netsrv-2.squaretrade.com")
    if err != nil {
        panic(err)
    }
}
func TestSetimage(t *testing.T){
    fmt.Println("testing set image")
	m, err := parse("config")
	if err != nil {
		panic(err)
    }
    err =  SetImage(m["sf_deploy"].Storage,"centos","centostest")
    if err != nil {
        panic(err)
    }
}
func TestDup(t *testing.T){
    fmt.Println("testing dup")
    d := make(map[string]int)
    d["centos"] = 3
    d["ubuntu"] = 5
	m, err := parse("config")
	if err != nil {
		panic(err)
    }
    err = storagedup(m["sf_deploy"].Storage,d)
    if err != nil {
        panic(err)
    }
}
func TestEditHost(t *testing.T) {
	fmt.Println("testing patch host")
	h := govirtlib.CreateVmForm{
		Hostname:  "sf01-lab-netsrv-2.squaretrade.com",
		VmMac:     "d4:ae:52:88:7f:73",
		VmIp:      "10.180.250.123",
		Leasetime: 15000,
	}
	m, err := parse("config")
	if err != nil {
		panic(err)
	}
	err = edithost(m["sf_deploy"].Godhcp, h,true)
	if err != nil {
		panic(err)
	}
}
