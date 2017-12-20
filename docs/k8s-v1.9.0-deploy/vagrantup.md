# DevOps

Update box
```
fanhonglingdeMacBook-Pro:~ fanhongling$ vagrant box list
coreos-alpha        (virtualbox, 969.0.0)
fedora26-cloud-base (virtualbox, 0)
fedora_inst         (virtualbox, 0)
ubuntu/xenial64     (virtualbox, 20171122.0.0)
```

```
fanhonglingdeMacBook-Pro:~ fanhongling$ vagrant box outdated --global
* 'ubuntu/xenial64' is outdated! Current: 20171122.0.0. Latest: 20171212.0.0
* 'fedora_inst' wasn't added from a catalog, no version information
* 'fedora26-cloud-base' wasn't added from a catalog, no version information
* 'coreos-alpha' (v969.0.0) is up to date
```

```
fanhonglingdeMacBook-Pro:~ fanhongling$ vagrant box update --box ubuntu/xenial64
Checking for updates to 'ubuntu/xenial64'
Latest installed version: 20171122.0.0
Version constraints: > 20171122.0.0
Provider: virtualbox
Updating 'ubuntu/xenial64' with provider 'virtualbox' from version
'20171122.0.0' to '20171212.0.0'...
Loading metadata for box 'https://atlas.hashicorp.com/ubuntu/xenial64'
Adding box 'ubuntu/xenial64' (v20171212.0.0) for provider: virtualbox
Downloading: https://vagrantcloud.com/ubuntu/boxes/xenial64/versions/20171212.0.0/providers/virtualbox.box
Successfully added box 'ubuntu/xenial64' (v20171212.0.0) for 'virtualbox'!
```

```
fanhonglingdeMacBook-Pro:~ fanhongling$ vagrant box list
coreos-alpha        (virtualbox, 969.0.0)
fedora26-cloud-base (virtualbox, 0)
fedora_inst         (virtualbox, 0)
ubuntu/xenial64     (virtualbox, 20171122.0.0)
ubuntu/xenial64     (virtualbox, 20171212.0.0)
```

```
fanhonglingdeMacBook-Pro:~ fanhongling$ vagrant box remove ubuntu/xenial64 --box-version=20171122.0.0
Removing box 'ubuntu/xenial64' (v20171122.0.0) with provider 'virtualbox'...
```

```
fanhonglingdeMacBook-Pro:~ fanhongling$ vagrant box list
coreos-alpha        (virtualbox, 969.0.0)
fedora26-cloud-base (virtualbox, 0)
fedora_inst         (virtualbox, 0)
ubuntu/xenial64     (virtualbox, 20171212.0.0)
```

workdir
```
fanhonglingdeMacBook-Pro:k8s-v1.9.0-deploy fanhongling$ pwd
/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/docs/k8s-v1.9.0-deploy
```

up
```
fanhonglingdeMacBook-Pro:k8s-v1.9.0-deploy fanhongling$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'ubuntu/xenial64'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'ubuntu/xenial64' is up to date...
==> default: Setting the name of the VM: k8s-v190-deploy_default_1513633316589_53408
==> default: Fixed port collision for 22 => 2222. Now on port 2200.
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
    default: Adapter 3: hostonly
==> default: Forwarding ports...
    default: 22 (guest) => 2200 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2200
    default: SSH username: ubuntu
    default: SSH auth method: password
    default: Warning: Connection reset. Retrying...
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Connection reset. Retrying...
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Connection reset. Retrying...
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Connection reset. Retrying...
    default: Warning: Authentication failure. Retrying...
    default: 
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
    default: The guest additions on this VM do not match the installed version of
    default: VirtualBox! In most cases this is fine, but in rare cases it can
    default: prevent things such as shared folders from working properly. If you see
    default: shared folder errors, please make sure the guest additions within the
    default: virtual machine match the version of VirtualBox you have installed on
    default: your host and reload your VM.
    default: 
    default: Guest Additions Version: 5.0.40
    default: VirtualBox Version: 5.1
==> default: Mounting shared folders...
    default: /Users/fanhongling/go => /Users/fanhongling/go
    default: /Users/fanhongling/Downloads => /Users/fanhongling/Downloads
```

Enter
```
fanhonglingdeMacBook-Pro:k8s-v1.9.0-deploy fanhongling$ vagrant ssh
Welcome to Ubuntu 16.04.3 LTS (GNU/Linux 4.4.0-103-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

  Get cloud support with Ubuntu Advantage Cloud Guest:
    http://www.ubuntu.com/business/services/cloud

0 packages can be updated.
0 updates are security updates.
```

appen into /etc/network/interfaces
```
auto enp0s8
iface enp0s8 inet static
  address 172.17.4.55
  netmask 255.255.255.0

auto enp0s9
iface enp0s9 inet dhcp
```

```
ubuntu@ubuntu-xenial:~$ sudo ip link set dev enp0s8 up
```

```
ubuntu@ubuntu-xenial:~$ sudo ip link set dev enp0s9 up
```

```
ubuntu@ubuntu-xenial:~$ sudo systemctl restart networking.service
```

```
ubuntu@ubuntu-xenial:~$ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 02:63:8e:2c:ef:cd brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global enp0s3
       valid_lft forever preferred_lft forever
    inet6 fe80::63:8eff:fe2c:efcd/64 scope link 
       valid_lft forever preferred_lft forever
3: enp0s8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:de:0a:b5 brd ff:ff:ff:ff:ff:ff
    inet 172.17.4.55/24 brd 172.17.4.255 scope global enp0s8
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fede:ab5/64 scope link 
       valid_lft forever preferred_lft forever
4: enp0s9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:f6:61:b0 brd ff:ff:ff:ff:ff:ff
    inet 172.28.128.4/24 brd 172.28.128.255 scope global enp0s9
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fef6:61b0/64 scope link 
       valid_lft forever preferred_lft forever
```