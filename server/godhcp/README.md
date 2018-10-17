### Introduction ###
isc-dhcp api. For host and network add/update

## GET

```
{
    target: $target
}
    
```
### target = host ###
Returns: 
```
govirtlib.ReturnPayload {
    HostInfos
}
```
    
### target = network ###
Returns: 
```
govirtlib.ReturnPayload {
    NetInfos
}
```

## POST
```
type PostPayload struct {
    Domain        string
    Target        string         `json:"target"`
    VmForm        CreateVmForm   `json:"createvmform"`  // the object for the govirthost that contains vminfo to create the virtual machine :govirthost
    //  Hostinfo      HostInfo       `json:"hostinfo"`      // the hostinfo to add to godhcp. :godhcp
    Netinfo     Network  `json:"netinfo"` // when you want to add or delete network for. :godhcp
}
```
### target = network ###
Post the network you want to add. You must fill in a valid govirtlib.Network object before doing so. 
Require fields:
    - Netinfo

### target = host ###
Post the network you want to add. You must fill in a valid govirtlib.Network object before doing so. 
Require fields:
    - VmForm

## PATCH
same as POST but being able to update existing records. 

## DELETE

### target = network ###
Delete the network 
Require fields: 
    - Netinfo.Subnet

### target = host ###
Delete the host 
Require fields: 
    - Domain 

