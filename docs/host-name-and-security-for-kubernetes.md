## DevOps

### Table of content

1. Fedora 26
1. Ubuntu 16.04

### Fedora 26

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

### Ubuntu 16.04

address
```
ubuntu@ubuntu-xenial:~$ hostname -I
10.0.2.15 172.17.4.55 172.28.128.4 172.18.0.1 
```

```
ubuntu@ubuntu-xenial:~$ pos=2; hostname -I | awk -v i=$pos '{print $i}'
172.17.4.55
```

Hostname
```
ubuntu@ubuntu-xenial:~$ echo "kubedev-172-17-4-55" | sudo tee /etc/hostname
kubedev-172-17-4-55
```

```
ubuntu@ubuntu-xenial:~$ sudo hostname --file /etc/hostname 
```

```
ubuntu@ubuntu-xenial:~$ sudo hostnamectl
sudo: unable to resolve host kubedev-172-17-4-55
   Static hostname: kubedev-172-17-4-55
         Icon name: computer-vm
           Chassis: vm
        Machine ID: 170d649385c24a6f80d0b74cee4516b7
           Boot ID: 601b6401f99f4a4ebbd48371c3a1da38
    Virtualization: oracle
  Operating System: Ubuntu 16.04.3 LTS
            Kernel: Linux 4.4.0-103-generic
      Architecture: x86-64
```

hosts
```
ubuntu@ubuntu-xenial:~$ echo "172.17.4.55    kubedev-172-17-4-55" | sudo tee -a /etc/hosts
sudo: unable to resolve host kubedev-172-17-4-55
172.17.4.55    kubedev-172-17-4-55
```

```
ubuntu@ubuntu-xenial:~$ ping -c3  kubedev-172-17-4-55
PING kubedev-172-17-4-55 (172.17.4.55) 56(84) bytes of data.
64 bytes from kubedev-172-17-4-55 (172.17.4.55): icmp_seq=1 ttl=64 time=0.034 ms
64 bytes from kubedev-172-17-4-55 (172.17.4.55): icmp_seq=2 ttl=64 time=0.033 ms
64 bytes from kubedev-172-17-4-55 (172.17.4.55): icmp_seq=3 ttl=64 time=0.063 ms

--- kubedev-172-17-4-55 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 1998ms
rtt min/avg/max/mdev = 0.033/0.043/0.063/0.014 ms
```

Firewall
```
ubuntu@ubuntu-xenial:~$ sudo ufw status
Status: inactive
```