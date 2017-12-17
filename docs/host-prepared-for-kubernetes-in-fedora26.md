### DevOps

Looking at network scripts 
```
[vagrant@kubedev-172-17-4-59 ~]$ cat /etc/sysconfig/network-scripts/ifcfg-eth0
DEVICE="eth0"
BOOTPROTO="dhcp"
ONBOOT="yes"
TYPE="Ethernet"
PERSISTENT_DHCLIENT="yes"
```

```
[vagrant@kubedev-172-17-4-59 ~]$ cat /etc/sysconfig/network-scripts/ifcfg-eth1
NM_CONTROLLED=no
#BOOTPROTO=dhcp
BOOTPROTO=none
ONBOOT=yes
IPADDR=172.17.4.59
NETMASK=255.255.255.0
DEVICE=eth1
PEERDNS=no
```

```
[vagrant@kubedev-172-17-4-59 ~]$ cat /etc/sysconfig/network-scripts/ifcfg-eth2
NM_CONTROLLED=no
BOOTPROTO=dhcp
#BOOTPROTO=none
ONBOOT=yes
#IPADDR=172.17.4.59
#NETMASK=255.255.255.0
DEVICE=eth2
PEERDNS=no
```

Hostname
```
[vagrant@localhost ~]$ name="kubedev";pos=2;suffix=$(hostname -I | cut -d' ' -f$pos) && suffix=${suffix//\./-} && echo "$name-$suffix" | sudo tee /etc/hostname
kubedev-172-17-4-59
```

```
[vagrant@localhost ~]$ sudo hostname --file=/etc/hostname 
```

hosts
```
[vagrant@localhost ~]$ echo -e "\n$(hostname -I | cut -d' ' -f2)    `hostname`" | sudo tee -a /etc/hosts

172.17.4.59    kubedev-172-17-4-59
```

Selinux
```
[vagrant@localhost ~]$ if [[ $(getenforce) == "Enforcing" ]] ;then sudo setenforce 0 ; sudo sed -i 's/SELINUX=enforcing/SELINUX=permissive/' /etc/sysconfig/selinux ; fi
```

```
[vagrant@kubedev-172-17-4-59 ~]$ sudo sed -i 's/SELINUX=enforcing/SELINUX=permissive/' /etc/selinux/config 
```

Firewall
```
[vagrant@localhost ~]$ if [[ $(systemctl is-active firewalld.service) == "active" ]] ;then sudo systemctl stop firewalld.service && sudo systemctl disable firewalld.service ; fi
```
