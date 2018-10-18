package controller

import (
	"fmt"
    "bytes"
    "math/rand"
    "time"
    "io/ioutil"
    "github.com/hunkeelin/govirt/govirtlib"
    "testing"
    "github.com/hunkeelin/klinutils"
)
func TestDup(t *testing.T){
    fmt.Println("testing dup")
    var err error
    c := Conn{}
    c.cb,err = ioutil.ReadFile("cert")
    if err != nil{
        panic(err)
    }
    c.kb,err = ioutil.ReadFile("key")
    if err != nil{
        panic(err)
    }
    c.tb,err = ioutil.ReadFile("govirt.crt")
    if err != nil{
        panic(err)
    }
    d := make(map[string]int)
    d["centos"] = 3
    d["ubuntu"] = 5
	m, err := parse("config")
	if err != nil {
		panic(err)
    }
    err = c.storagedup(m["sf_deploy"].Storage,d)
    if err != nil {
        panic(err)
    }
}
func TestDelHost(t *testing.T) {
	fmt.Println("testing delete host")
    var err error
    c := Conn{}
    c.cb,err = ioutil.ReadFile("cert")
    if err != nil{
        panic(err)
    }
    c.kb,err = ioutil.ReadFile("key")
    if err != nil{
        panic(err)
    }
    c.tb,err = ioutil.ReadFile("govirt.crt")
    if err != nil{
        panic(err)
    }
	m, err := parse("config")
	if err != nil {
		panic(err)
	}
	err = c.delhost(m["sf_deploy"].Godhcp, "sf01-test2.squaretrade.com")
	if err != nil {
		panic(err)
	}
}
func TestEditHost(t *testing.T) {
	fmt.Println("testing patch host")
    var err error
    c := Conn{}
    c.cb,err = ioutil.ReadFile("cert")
    if err != nil{
        panic(err)
    }
    c.kb,err = ioutil.ReadFile("key")
    if err != nil{
        panic(err)
    }
    c.tb,err = ioutil.ReadFile("govirt.crt")
    if err != nil{
        panic(err)
    }
	h := govirtlib.CreateVmForm{
		Hostname:  "sf01-test2.squaretrade.com",
		VmMac:     "d4:ae:52:88:7f:74",
		VmIp:      "10.180.250.40",
		Leasetime: 15000,
	}
	m, err := parse("config")
	if err != nil {
		panic(err)
	}
	err = c.edithost(m["sf_deploy"].Godhcp, h,false)
	if err != nil {
		panic(err)
	}
}
func TestSetimage(t *testing.T){
    fmt.Println("testing set image")
    var err error
    c := Conn{}
    c.cb,err = ioutil.ReadFile("cert")
    if err != nil{
        panic(err)
    }
    c.kb,err = ioutil.ReadFile("key")
    if err != nil{
        panic(err)
    }
    c.tb,err = ioutil.ReadFile("govirt.crt")
    if err != nil{
        panic(err)
    }
	m, err := parse("config")
	if err != nil {
		panic(err)
    }
    err =  c.setimage(m["sf_deploy"].Storage,"ubuntu","ubuntu3")
    if err != nil {
        panic(err)
    }
}
func TestDefinevm(t *testing.T){
    fmt.Println("testing create vm")
    var err error
    c := Conn{}
    c.cb,err = ioutil.ReadFile("cert")
    if err != nil{
        panic(err)
    }
    c.kb,err = ioutil.ReadFile("key")
    if err != nil{
        panic(err)
    }
    c.tb,err = ioutil.ReadFile("govirt.crt")
    if err != nil{
        panic(err)
    }
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
    template = bytes.Replace(template,[]byte("name_replace"),[]byte("cent2"),-1)
    template = bytes.Replace(template,[]byte("uuid_replace"),uuid,-1)
    template = bytes.Replace(template,[]byte("memory_replace"),[]byte("4"),-1)
    template = bytes.Replace(template,[]byte("cpu_replace"),[]byte("2"),-1)
    template = bytes.Replace(template,[]byte("imagedir_replace"),[]byte("/data/govirt/storage"),-1)
    template = bytes.Replace(template,[]byte("mac_replace"),mac,-1)
    template = bytes.Replace(template,[]byte("vlan_replace"),[]byte("govirtmgmt"),-1)
    fmt.Println(string(template))
	m, err := parse("config")
	if err != nil {
		panic(err)
	}
    rand.Seed(time.Now().UTC().UnixNano())
    randhostint := randInt(0,len(m["sf_deploy"].Govirt))
    err = c.Define(template,m["sf_deploy"].Govirt[randhostint])
    if err != nil {
        panic(err)
    }
}
func TestMigrate(t *testing.T){
    var err error
    c := Conn{}
    c.cb,err = ioutil.ReadFile("cert")
    if err != nil{
        panic(err)
    }
    c.kb,err = ioutil.ReadFile("key")
    if err != nil{
        panic(err)
    }
    c.tb,err = ioutil.ReadFile("govirt.crt")
    if err != nil{
        panic(err)
    }
    err = c.Migrate("sf01-lab-2.squaretrade.com","sf01-lab-netsrv-2.squaretrade.com","cent1")
    if err != nil {
        panic(err)
    }
    b,err := c.Getxml("cent1","sf01-lab-netsrv-2.squaretrade.com")
    if err != nil {
        panic(err)
    }
    err = c.Statevm("unDefine","cent1","sf01-lab-2.squaretrade.com")
    if err != nil {
        panic(err)
    }
    err = c.Define(b,"sf01-lab-netsrv-2.squaretrade.com")
    if err != nil {
        panic(err)
    }
}
func TestGetvm(t *testing.T){
    var err error
    c := Conn{}
    c.cb,err = ioutil.ReadFile("cert")
    if err != nil{
        panic(err)
    }
    c.kb,err = ioutil.ReadFile("key")
    if err != nil{
        panic(err)
    }
    c.tb,err = ioutil.ReadFile("govirt.crt")
    if err != nil{
        panic(err)
    }
    fmt.Println("getting a list of vm")
    m, err := parse("config")
    if err != nil {
        panic(err)
    }
    p, err := c.Getvms(m["sf_deploy"].Govirt)
    if err != nil {
        panic(err)
    }
    for _,hostsd:= range p {
        fmt.Println("ParentHost:",hostsd.Parent)
        for _,i := range hostsd.Domains{
            fmt.Printf("\t%s\t%s%s\n", i.Domain.Name, "----> ", i.State)
        }
    }
}
func TestStatevm(t *testing.T){
    fmt.Println("testing start vm with https")
    var err error
    c := Conn{}
    c.cb,err = ioutil.ReadFile("cert")
    if err != nil{
        panic(err)
    }
    c.kb,err = ioutil.ReadFile("key")
    if err != nil{
        panic(err)
    }
    c.tb,err = ioutil.ReadFile("govirt.crt")
    if err != nil{
        panic(err)
    }
    err = c.Statevm("start","cent1","sf01-lab-2.squaretrade.com")
    if err != nil {
        panic(err)
    }
}
func TestEditNet(t *testing.T) {
    fmt.Println("testing patch network")
    var err error
    c := Conn{}
    c.cb,err = ioutil.ReadFile("cert")
    if err != nil{
        panic(err)
    }
    c.kb,err = ioutil.ReadFile("key")
    if err != nil{
        panic(err)
    }
    c.tb,err = ioutil.ReadFile("govirt.crt")
    if err != nil{
        panic(err)
    }
    n := govirtlib.Network{
        Subnet:   "10.180.250.0",
        Netmask:  "255.255.254.0",
        Dns:      []string{"10.181.35.100", "10.181.35.101"},
        Router:   "10.180.250.1",
        Iprange:  []string{"10.180.250.160", "10.180.250.200"},
        Lease:    "3601",
        Maxlease: "7200",
    }
    m, err := parse("config")
    if err != nil {
        panic(err)
    }
    err = c.editnetwork(m["sf_deploy"].Godhcp, n,true)
    if err != nil {
        panic(err)
    }
}

