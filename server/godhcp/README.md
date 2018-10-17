### Introduction ###
isc-dhcp api. For host and network add/update

### Security ###
Can host via http or https

### Usage ###

If you are querying the API. uses govirtlib.PostPayload and GetPayload

- Method `GET`
```
{
    target: $target // host,network
}
    
```

- `POST` `PATCH` `DELETE`
```
{
    target: $target, // host,network
    hostinfo: { // govirtlib.CreateVmForm relevant if target=host
        Hostname:  string `json:"hostname"`
        Mac:       string `json:"mac"`
        Ip:        string `json:"ip"`
        Leasetime: int    `json:"leasetime"`
    },
    network: { //govirtlib.Network relevant if target=network
        Subnet:  string   `json:"subnet"`
        Netmask: string   `json:"netmask"`
        Dns:     []string `json:"dns"`
        Router:  string   `json:"router"`
        Iprange: []string `json:"iprange"`
    },
}
```

#### Todo ####
- Post network add and update
- Post network delete
