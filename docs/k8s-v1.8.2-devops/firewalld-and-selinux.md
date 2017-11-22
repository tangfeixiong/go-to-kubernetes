### Firwalld

It is stopped
```
[vagrant@openshiftdev ~]$ sudo systemctl is-active firewalld.service
unknown
[vagrant@openshiftdev ~]$ sudo systemctl status firewalld.service
● firewalld.service - firewalld - dynamic firewall daemon
   Loaded: loaded (/usr/lib/systemd/system/firewalld.service; disabled; vendor preset: enabled)
   Active: inactive (dead)
```

Command to stop and disable auto-started
```
[tangfx@localhost ~]$ sudo systemctl stop firewalld.service
[tangfx@localhost ~]$ sudo systemctl disable firewalld.service
```


### Selinux

It is not enforcing
```
[vagrant@openshiftdev ~]$ sudo getenforce
Permissive
```

Command to change
```
[vagrant@openshiftdev ~]$ sudo getenforce
Enforcing
[vagrant@openshiftdev ~]$ sudo setenforce 0
[vagrant@openshiftdev ~]$ sudo getenforce
Permissive
```

Command to disable auto-started
```
[tangfx@localhost ~]$ sudo sed -i 's/SELINUX=.*/SELINUE=permissive/' /etc/sysconfig/selinux
```