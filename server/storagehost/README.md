# Introduction #
This is the storage server. This server should be sitting on top of one of the gluster host. 

## GET 
``` 
{
    "Target": $target
}
```

### target = images ###
This will list all the templates in the storage directory. 

## POST 
```
type PostPayload struct {
    Domain        string         `json:"domain"` // the domain to action should apply to :govirthost
    Action        string         `json:"action"` // what action right now only create, reset. shutdown, start. :All
    VmForm        CreateVmForm   `json:"createvmform"`  // the object for the govirthost that contains vminfo to create the virtual machine :govirthost
    DuplicateInfo map[string]int `json:"duplicateinfo"` // for storage, string is name of the image and the int is how many copy of the image :storagehost
    //  Hostinfo      HostInfo       `json:"hostinfo"`      // the hostinfo to add to godhcp. :godhcp
    AddStrgInfo StrgInfo `json:"storageinfo"`
}
```

### Action = dup ###
This is a performance booster. The software retains copies aka cache of each image, it will instead rename the image for each host instead of copying it.
Require fields:
    - DuplicateInfo: It's a map of string and int. The string is the image that you want to duplicate, the int is the copies of it. 
### Action = setimage ###
This is basically renaming the images copies or copying from origin if no copies to the hsotname you want to set. The most basic function of imaging. 
Require fields:
    - VmForm.Image: The image you want to you.
    - VmForm.Hostname: The hostname you want it to set to. 
### Action = addstorage ###
This is adding storage, equi of qemu-img create qcow2. With a specifc name ${hostname}_storage_${size}G.
Require fields:
    - AddStrgInfo.Hostname: The hostname.
    - AddStrgInfo.Size: The size. 

## DELETE 
```
type PostPayload struct {
    Action        string         `json:"action"` // what action right now only create, reset. shutdown, start. :All
    Target        string         `json:"target"`
    //  Hostinfo      HostInfo       `json:"hostinfo"`      // the hostinfo to add to godhcp. :godhcp
}
```
### Action = host ###
Deletes a specific host 
Require fields:
    - Target: hostname

### Action = image ###
Deletes a specific template 
Require fields:
    - Target: name of the image


