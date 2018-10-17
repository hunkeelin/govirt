# Introduction #
This is the govirthost.


## GET ##
```
{ 
    "target": $item
    "domain": $domain
}
```
### item = vm ###
Will return a json as follows. A list of domains in the machine with it's state and domain. 
```
type ReturnPayload struct {
    Domains        []DomainInfo            `json:"domains"`
}
type DomainInfo struct {
    Domain libvirt.Domain `json:"domain"`
    State  string         `json:"state"`
}
```
### item = xml ###
Required Fields: 
    - domain 
returns the xml in []byte of the specific domain
### item = metal ###
Will return the vmhost system information. 
```
type ReturnPayload struct {
    HostMemoryInfo *sysinfo.SI             `json:"hostmemoryInfo"`
}
type SI struct {
    Uptime       time.Duration // time since boot
    Loads        [3]float64    // 1, 5, and 15 minute load averages, see e.g. UPTIME(1)
    Procs        uint64        // number of current processes
    TotalRam     uint64        // total usable main memory size [kB]
    FreeRam      uint64        // available memory size [kB]
    SharedRam    uint64        // amount of shared memory [kB]
    BufferRam    uint64        // memory used by buffers [kB]
    TotalSwap    uint64        // total swap space size [kB]
    FreeSwap     uint64        // swap space still available [kB]
    TotalHighRam uint64        // total high memory size [kB]
    FreeHighRam  uint64        // available high memory size [kB]
    // contains filtered or unexported fields
}
```

## POST ##
```
type PostPayload struct {
    Domain        string         `json:"domain"` // the domain to action should apply to :govirthost
    Action        string         `json:"action"` // what action right now only create, reset. shutdown, start. :All
    Target        string         `json:"target"`
    VmForm        CreateVmForm   `json:"createvmform"`  // the object for the govirthost that contains vminfo to create the virtual machine :govirthost

}
```
### Action = start/shutdown/reset/destroy/define/undefine ###
Required fields: 
- `Domain`: the vm you want to start/shutdown/reset/destroy/define/undefine

### Action = create ###
Required fields: 
- VmForm: 
```
type CreateVmForm struct {
    Hostname    string `json:"hostname"`
    VmMac       string `json:"mac"`         // the mac addressA
    VmIp        string `json:"ip"`          // the ip address of the virutal machine
    CpuCount    int    `json:"cpucount"`    // cpucount
    MemoryCount int    `json:"memorycount"` // in GB
    Image       string `json:"image"`       // the image to use
    Xml         []byte `json:"xml"`         // the xml in bytes
    Leasetime   int    `json:"leasetime"`   // this is for godhcp
}
```
- CreateVmForm.Xml: this is the []byte of the xml that will get defined and start. 

### Action = migrate ###
Require fields: 
- Target: The destination vmhost. 
- Domain: the Vm you want to migrate. 

## DELETE ##
```
type PostPayload struct {
    Domain        string         `json:"domain"` // the domain to action should apply to :govirthost
    Action        string         `json:"action"` // what action right now only create, reset. shutdown, start. :All
}
```

Require fields: 
- Domain: the vm you want to undefine
