## DevOps

### Cloudimg 1.13.1

install
```
ubuntu@ubuntu-xenial:~$ sudo apt update && sudo apt install -y docker.io
Hit:1 http://archive.ubuntu.com/ubuntu xenial InRelease
Get:2 http://security.ubuntu.com/ubuntu xenial-security InRelease [102 kB]
Get:3 http://archive.ubuntu.com/ubuntu xenial-updates InRelease [102 kB]          
Get:4 http://archive.ubuntu.com/ubuntu xenial-backports InRelease [102 kB]         
Get:5 http://security.ubuntu.com/ubuntu xenial-security/main Sources [104 kB]
Get:6 http://archive.ubuntu.com/ubuntu xenial/main Sources [868 kB]  
Get:7 http://security.ubuntu.com/ubuntu xenial-security/restricted Sources [2,600 B]
Get:8 http://security.ubuntu.com/ubuntu xenial-security/universe Sources [48.9 kB]
Get:9 http://security.ubuntu.com/ubuntu xenial-security/multiverse Sources [1,520 B]
Get:10 http://security.ubuntu.com/ubuntu xenial-security/main amd64 Packages [408 kB]
Get:11 http://archive.ubuntu.com/ubuntu xenial/restricted Sources [4,808 B]
Get:12 http://archive.ubuntu.com/ubuntu xenial/universe Sources [7,728 kB]      
Get:13 http://security.ubuntu.com/ubuntu xenial-security/main Translation-en [179 kB]
Get:14 http://security.ubuntu.com/ubuntu xenial-security/universe amd64 Packages [190 kB]                                                      
Get:15 http://security.ubuntu.com/ubuntu xenial-security/universe Translation-en [99.0 kB]                                                     
Get:16 http://security.ubuntu.com/ubuntu xenial-security/multiverse amd64 Packages [3,208 B]                                                   
Get:17 http://security.ubuntu.com/ubuntu xenial-security/multiverse Translation-en [1,408 B]                                                   
Get:18 http://archive.ubuntu.com/ubuntu xenial/multiverse Sources [179 kB]                                                                     
Get:19 http://archive.ubuntu.com/ubuntu xenial/universe amd64 Packages [7,532 kB]                                                              
Get:20 http://archive.ubuntu.com/ubuntu xenial/universe Translation-en [4,354 kB]                                                              
Get:21 http://archive.ubuntu.com/ubuntu xenial/multiverse amd64 Packages [144 kB]                                                              
Get:22 http://archive.ubuntu.com/ubuntu xenial/multiverse Translation-en [106 kB]                                                              
Get:23 http://archive.ubuntu.com/ubuntu xenial-updates/main Sources [286 kB]                                                                   
Get:24 http://archive.ubuntu.com/ubuntu xenial-updates/restricted Sources [3,404 B]                                                            
Get:25 http://archive.ubuntu.com/ubuntu xenial-updates/universe Sources [185 kB]                                                               
Get:26 http://archive.ubuntu.com/ubuntu xenial-updates/multiverse Sources [7,968 B]                                                            
Get:27 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 Packages [681 kB]                                                            
Get:28 http://archive.ubuntu.com/ubuntu xenial-updates/main Translation-en [285 kB]                                                            
Get:29 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 Packages [565 kB]                                                        
Get:30 http://archive.ubuntu.com/ubuntu xenial-updates/universe Translation-en [229 kB]                                                        
Get:31 http://archive.ubuntu.com/ubuntu xenial-updates/multiverse amd64 Packages [16.2 kB]                                                     
Get:32 http://archive.ubuntu.com/ubuntu xenial-updates/multiverse Translation-en [8,052 B]                                                     
Get:33 http://archive.ubuntu.com/ubuntu xenial-backports/main Sources [3,428 B]                                                                
Get:34 http://archive.ubuntu.com/ubuntu xenial-backports/universe Sources [4,908 B]                                                            
Get:35 http://archive.ubuntu.com/ubuntu xenial-backports/main amd64 Packages [4,860 B]                                                         
Get:36 http://archive.ubuntu.com/ubuntu xenial-backports/main Translation-en [3,220 B]                                                         
Get:37 http://archive.ubuntu.com/ubuntu xenial-backports/universe amd64 Packages [6,612 B]                                                     
Get:38 http://archive.ubuntu.com/ubuntu xenial-backports/universe Translation-en [3,768 B]                                                     
Fetched 24.6 MB in 57s (425 kB/s)                                                                                                              
Reading package lists... Done
Building dependency tree       
Reading state information... Done
8 packages can be upgraded. Run 'apt list --upgradable' to see them.
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  bridge-utils cgroupfs-mount containerd runc ubuntu-fan
Suggested packages:
  mountall aufs-tools debootstrap docker-doc rinse zfs-fuse | zfsutils
The following NEW packages will be installed:
  bridge-utils cgroupfs-mount containerd docker.io runc ubuntu-fan
0 upgraded, 6 newly installed, 0 to remove and 8 not upgraded.
Need to get 17.5 MB of archives.
After this operation, 90.5 MB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu xenial/main amd64 bridge-utils amd64 1.5-9ubuntu1 [28.6 kB]
Get:2 http://archive.ubuntu.com/ubuntu xenial/universe amd64 cgroupfs-mount all 1.2 [4,970 B]
Get:3 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 runc amd64 1.0.0~rc2+docker1.13.1-0ubuntu1~16.04.1 [1,488 kB]
Get:4 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 containerd amd64 0.2.5-0ubuntu1~16.04.1 [4,041 kB]
Get:5 http://archive.ubuntu.com/ubuntu xenial-updates/universe amd64 docker.io amd64 1.13.1-0ubuntu1~16.04.2 [11.9 MB]                         
Get:6 http://archive.ubuntu.com/ubuntu xenial-updates/main amd64 ubuntu-fan all 0.12.8~16.04.1 [35.6 kB]                                       
Fetched 17.5 MB in 24s (713 kB/s)                                                                                                              
Selecting previously unselected package bridge-utils.
(Reading database ... 53972 files and directories currently installed.)
Preparing to unpack .../bridge-utils_1.5-9ubuntu1_amd64.deb ...
Unpacking bridge-utils (1.5-9ubuntu1) ...
Selecting previously unselected package cgroupfs-mount.
Preparing to unpack .../cgroupfs-mount_1.2_all.deb ...
Unpacking cgroupfs-mount (1.2) ...
Selecting previously unselected package runc.
Preparing to unpack .../runc_1.0.0~rc2+docker1.13.1-0ubuntu1~16.04.1_amd64.deb ...
Unpacking runc (1.0.0~rc2+docker1.13.1-0ubuntu1~16.04.1) ...
Selecting previously unselected package containerd.
Preparing to unpack .../containerd_0.2.5-0ubuntu1~16.04.1_amd64.deb ...
Unpacking containerd (0.2.5-0ubuntu1~16.04.1) ...
Selecting previously unselected package docker.io.
Preparing to unpack .../docker.io_1.13.1-0ubuntu1~16.04.2_amd64.deb ...
Unpacking docker.io (1.13.1-0ubuntu1~16.04.2) ...
Selecting previously unselected package ubuntu-fan.
Preparing to unpack .../ubuntu-fan_0.12.8~16.04.1_all.deb ...
Unpacking ubuntu-fan (0.12.8~16.04.1) ...
Processing triggers for man-db (2.7.5-1) ...
Processing triggers for ureadahead (0.100.0-19) ...
Processing triggers for systemd (229-4ubuntu21) ...
Setting up bridge-utils (1.5-9ubuntu1) ...
Setting up cgroupfs-mount (1.2) ...
Setting up runc (1.0.0~rc2+docker1.13.1-0ubuntu1~16.04.1) ...
Setting up containerd (0.2.5-0ubuntu1~16.04.1) ...
Setting up docker.io (1.13.1-0ubuntu1~16.04.2) ...
Adding group `docker' (GID 117) ...
Done.
Setting up ubuntu-fan (0.12.8~16.04.1) ...
Processing triggers for systemd (229-4ubuntu21) ...
Processing triggers for ureadahead (0.100.0-19) ...
```

status
```
ubuntu@ubuntu-xenial:~$ sudo systemctl -l status --no-pager docker.service
● docker.service - Docker Application Container Engine
   Loaded: loaded (/lib/systemd/system/docker.service; enabled; vendor preset: enabled)
   Active: active (running) since Mon 2017-12-18 22:24:38 UTC; 3min 34s ago
     Docs: https://docs.docker.com
 Main PID: 3643 (dockerd)
   CGroup: /system.slice/docker.service
           ├─3643 /usr/bin/dockerd -H fd://
           └─3648 containerd -l unix:///var/run/docker/libcontainerd/docker-containerd.sock --metrics-interval=0 --start-timeout 2m --state-dir 

Dec 18 22:24:37 ubuntu-xenial dockerd[3643]: time="2017-12-18T22:24:37.940281304Z" level=info msg="Firewalld running: false"
Dec 18 22:24:38 ubuntu-xenial dockerd[3643]: time="2017-12-18T22:24:38.016989978Z" level=info msg="Default bridge (docker0) is assigned with an 
Dec 18 22:24:38 ubuntu-xenial dockerd[3643]: time="2017-12-18T22:24:38.075716831Z" level=info msg="Loading containers: done."
Dec 18 22:24:38 ubuntu-xenial dockerd[3643]: time="2017-12-18T22:24:38.078415963Z" level=warning msg="Couldn't run auplink before unmount /var/l
Dec 18 22:24:38 ubuntu-xenial dockerd[3643]: time="2017-12-18T22:24:38.085342675Z" level=warning msg="failed to retrieve runc version: unknown o
Dec 18 22:24:38 ubuntu-xenial dockerd[3643]: time="2017-12-18T22:24:38.086167733Z" level=warning msg="failed to retrieve docker-init version: un
Dec 18 22:24:38 ubuntu-xenial dockerd[3643]: time="2017-12-18T22:24:38.088148139Z" level=info msg="Daemon has completed initialization"
Dec 18 22:24:38 ubuntu-xenial dockerd[3643]: time="2017-12-18T22:24:38.088377614Z" level=info msg="Docker daemon" commit=092cba3 graphdriver=auf
Dec 18 22:24:38 ubuntu-xenial systemd[1]: Started Docker Application Container Engine.
Dec 18 22:24:38 ubuntu-xenial dockerd[3643]: time="2017-12-18T22:24:38.100776324Z" level=info msg="API listen on /var/run/docker.sock"
```

```
ubuntu@ubuntu-xenial:~$ sudo usermod -aG docker ubuntu
```

Enter
```
fanhonglingdeMacBook-Pro:k8s-v1.9.0-deploy fanhongling$ ssh -i .vagrant/machines/default/virtualbox/private_key ubuntu@172.17.4.55
The authenticity of host '172.17.4.55 (172.17.4.55)' can't be established.
RSA key fingerprint is 34:4d:ca:3a:b8:57:39:45:dc:93:32:20:3f:7a:bc:98.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '172.17.4.55' (RSA) to the list of known hosts.
Welcome to Ubuntu 16.04.3 LTS (GNU/Linux 4.4.0-103-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

  Get cloud support with Ubuntu Advantage Cloud Guest:
    http://www.ubuntu.com/business/services/cloud

11 packages can be updated.
8 updates are security updates.


Last login: Mon Dec 18 21:43:53 2017 from 10.0.2.2
```

```
ubuntu@ubuntu-xenial:~$ docker info
Containers: 0
 Running: 0
 Paused: 0
 Stopped: 0
Images: 0
Server Version: 1.13.1
Storage Driver: aufs
 Root Dir: /var/lib/docker/aufs
 Backing Filesystem: extfs
 Dirs: 0
 Dirperm1 Supported: true
Logging Driver: json-file
Cgroup Driver: cgroupfs
Plugins: 
 Volume: local
 Network: bridge host macvlan null overlay
Swarm: inactive
Runtimes: runc
Default Runtime: runc
Init Binary: docker-init
containerd version:  (expected: aa8187dbd3b7ad67d8e5e3a15115d3eef43a7ed1)
runc version: N/A (expected: 9df8b306d01f59d3a8029be411de015b7304dd8f)
init version: N/A (expected: 949e6facb77383876aeff8a6944dde66b3089574)
Security Options:
 apparmor
 seccomp
  Profile: default
Kernel Version: 4.4.0-103-generic
Operating System: Ubuntu 16.04.3 LTS
OSType: linux
Architecture: x86_64
CPUs: 1
Total Memory: 1.953 GiB
Name: ubuntu-xenial
ID: FE3L:4HRR:BOAM:F7UB:XYQ6:VOTJ:2UKK:BMKL:SH4Q:MI5F:XF5P:W7TI
Docker Root Dir: /var/lib/docker
Debug Mode (client): false
Debug Mode (server): false
Registry: https://index.docker.io/v1/
WARNING: No swap limit support
Experimental: false
Insecure Registries:
 127.0.0.0/8
Live Restore Enabled: false
```
