## Introduction 
This is my attempt on writing my own version of sdn for vm deployment. I am basically jumping right into the fire and get myself burn.

### Workflow 
1. Admin send vmInfo to controller 
2. Controller send the storagehost to rename the cache image. 
3. Controller will generate the xml in `[]byte` and send it to the designated govirthost. 
4. The govirthost will define the virtula machine using the `[]byte` xml it got from the controller. 

### libvirtd Configuration 
This software heavily depends on libvirtd. Security is handle via mtls. To configure mtls you will have to do the following

1. edit `/etc/sysconfig/libvirtd` uncomment `LIBVIRTD_ARGS="--listen"`
2. edit `/etc/libvirt/libvirtd.conf` with the following configuration
3. edit `/etc/libvirt/qemu.conf`. All qcow2 needs to be own have root. 
```
user = "root"
group = "root"
dynamic_ownership = 0
```
```
listen_tls = 1
listen_addr = $listen_Addr 
key_file = $keyfile_for_mtls //suggestion to put it as /etc/pki/libvirt/private/clientkey.pem
cert_file = $certfile_for_mtls // suggest to put it in /etc/pki/libvirt/clientcert.pem
ca_file = $cafile_for_trust // suggest to put it in /etc/pki/CA/cacert.pem
log_level = 3
```
The suggestion on the cert,key,ca file location is because for mtls to work, the client need to send it's cert as well and the location I listed is where libvirt will use to talk to libvirt server. There are ways to change the server's key,cert,ca location but not the client side and since govirthost needs to talk to each other might as well use the same cert,key,ca to save the time and complexity. 
