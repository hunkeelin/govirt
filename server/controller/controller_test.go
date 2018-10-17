package controller
import(
    "testing"
    "fmt"
    "github.com/hunkeelin/govirt/govirtlib"
)
func TestPatchNetwork(t *testing.T){
    fmt.Println("testing patch network")
    n := govirtlib.Network {
        Subnet: "10.180.250.0",
        Netmask: "255.255.255.0",
        Dns: []string{"10.181.35.100","10.181.35.101"},
        Router: "10.181.250.1",
        Iprange: []string{"10.180.250.160","10.180.250.200"},
        Lease: "3600",
        Maxlease: "7200",
    }
    err := editnetwork(n)
    if err != nil {
        panic(err)
    }
}
func TestPatchHost(t *testing.T){
    fmt.Println("testing patch host")
    h := govirtlib.CreateVmForm {
        Hostname: "sf01-lab-netsrv-2.squaretrade.com",
        VmMac: "d4:ae:52:88:7f:73",
        VmIp: "10.180.250.123",
        Leasetime: 16000,
    }
    err := edithost(h)
    if err != nil {
        panic(err)
    }
}
