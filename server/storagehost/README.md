# Introduction #
This is the storage server. This server should be sitting on top of one of the gluster host. 

## GET ##
``` 
{
    "Target": $target
}
```

### target = images ###
This will list all the templates in the storage directory. 

## POST ##
```
type PostPayload struct {
    Domain        string         `json:"domain"` // the domain to action should apply to :govirthost
    Action        string         `json:"action"` // what action right now only create, reset. shutdown, start. :All
    Target        string         `json:"target"`
    VmForm        CreateVmForm   `json:"createvmform"`  // the object for the govirthost that contains vminfo to create the virtual machine :govirthost
    DuplicateInfo map[string]int `json:"duplicateinfo"` // for storage, string is name of the image and the int is how many copy of the image :storagehost
    //  Hostinfo      HostInfo       `json:"hostinfo"`      // the hostinfo to add to godhcp. :godhcp
    Netinfo     Network  `json:"netinfo"` // when you want to add or delete network for. :godhcp
    AddStrgInfo StrgInfo `json:"storageinfo"`
    Xml         []byte   `json:"xml"`
}
```

