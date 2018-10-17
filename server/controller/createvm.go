package controller

import (
	"bytes"
	"github.com/hunkeelin/govirt/govirtlib"
	"github.com/hunkeelin/klinutils"
	"io/ioutil"
	"strconv"
)

func fillinform(file string, v govirtlib.CreateVmForm) ([]byte, error) {
	// need to get the vmform in xml
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return f, err
	}
	f = bytes.Replace(f, []byte("name_replace"), []byte(v.Hostname), -1)
	uuid, err := klinutils.Genuuid()
	if err != nil {
		panic(err)
	}
	f = bytes.Replace(f, []byte("uuid_replace"), uuid, -1)
	f = bytes.Replace(f, []byte("memory_replace"), []byte(strconv.Itoa(v.MemoryCount)), -1)
	f = bytes.Replace(f, []byte("cpu_replace"), []byte(strconv.Itoa(v.CpuCount)), -1)
	if err != nil {
		panic(err)
	}
	f = bytes.Replace(f, []byte("mac_replace"), []byte(v.VmMac), -1)
	f = bytes.Replace(f, []byte("vlan_replace"), []byte(v.Vlan), -1)
	return f, nil
}

//func main() {
//	mac, err := klinutils.Genmac()
//	if err != nil {
//		panic(err)
//	}
//	v := govirtlib.CreateVmForm{
//		CpuCount:    2,
//		MemoryCount: 4,
//		VmMac:       string(mac),
//		Hostname:    "noobfuck",
//		Vlan:        "virbr0",
//		Image:       "ctemplate",
//		VmIp:        "192.168.38.243",
//	}
//	f, err := fillinform("ctemplate.xml", v)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(f))
//}
