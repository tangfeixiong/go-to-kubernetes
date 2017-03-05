Hadoop installation
======================

Prerequisites
---------------

downoad-oracle-java-8.sh

```
#! /bin/bash
curl -LsO 'http://download.oracle.com/otn-pub/java/jdk/8u73-b15/jdk-8u73-linux-x64.tar.gz' -H 'Cookie: oraclelicense=accept-securebackup-cookie'
```

Extrat to __/opt__

```
$ sudo tar -C /opt -zxf jdk-8u73-linux-x64.tar.gz
$ sudo chown -R ubuntu: /opt/jdk1.8.0_73

```

download-apache-hadoop.sh

```
#! /bin/bash
curl -LsO http://mirrors.cnnic.cn/apache/hadoop/common/current/hadoop-3.0.0-alpha2.tar.gz
```

Extract

```
$ sudo tar -C /opt -zxf hadoop-3.0.0-alpha2.tar.gz
$ sudo chown -R ubuntu: /opt/hadoop-3.0.0-alpha2

```

Add Java environment into profile

```
$ echo 'export PATH=$PATH:/opt/jdk1.8.0_73/bin' | sudo tee -a /etc/profile.d/java.sh
$ echo 'export JAVA_HOME=/opt/jdk1.8.0_73' | sudo tee -a /etc/profile.d/java.sh
$ echo 'export HADOOP_HOME=/opt/hadoop-3.0.0-alpha2' | sudo tee -a /etc/profile.d/hadoop.sh
$ . /etc/profile.d/java.sh
$ . /etc/profile.d/hadoop.sh

```

Network Configuration
-----------------------

namenode

```
ubuntu@datanode125:~$ cat /etc/network/interfaces
# This file describes the network interfaces available on your system
# and how to activate them. For more information, see interfaces(5).

# The loopback network interface
auto lo
iface lo inet loopback

# The primary network interface
auto eth0
iface eth0 inet dhcp

auto eth1
iface eth1 inet static
address 10.64.33.125
netmask 255.255.255.0

ubuntu@datanode125:~$ cat /etc/hostname
datanode125

ubuntu@datanode125:~$ sudo hostname datanode125

ubuntu@datanode126:~$ cat /etc/hosts
127.0.0.1	localhost
127.0.1.1	ubuntu-14-04-3-server-amd64
10.64.33.125	datanode125
10.64.33.126	datanode126

# The following lines are desirable for IPv6 capable hosts
::1     localhost ip6-localhost ip6-loopback
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters

```

datanode

```
ubuntu@datanode126:~$ cat /etc/network/interfaces
# This file describes the network interfaces available on your system
# and how to activate them. For more information, see interfaces(5).

# The loopback network interface
auto lo
iface lo inet loopback

# The primary network interface
auto eth0
iface eth0 inet dhcp

auto eth1
iface eth1 inet static
address 10.64.33.126
netmask 255.255.255.0

ubuntu@datanode126:~$ cat /etc/hostname
datanode126

ubuntu@datanode125:~$ sudo hostname datanode126

ubuntu@datanode126:~$ cat /etc/hosts
127.0.0.1	localhost
127.0.1.1	ubuntu-14-04-3-server-amd64
10.64.33.125	datanode125
10.64.33.126	datanode126

# The following lines are desirable for IPv6 capable hosts
::1     localhost ip6-localhost ip6-loopback
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters

```

