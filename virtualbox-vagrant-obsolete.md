# Instruction

__Table of contents__


Fedora 26 VirtualBox for Mac

CentOS 7.4 VirtualBox for Windows 10

__Prerequisites__

VirtualBox 5.1.30 for Mac
```
fanhonglingdeMacBook-Pro:~ fanhongling$ curl -jkSL http://download.virtualbox.org/virtualbox/5.1.30/VirtualBox-5.1.30-118389-OSX.dmg -o Downloads/99-mirror/darwin-packages/VirtualBox-5.1.30-118389-OSX.dmg
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 90.2M  100 90.2M    0     0   700k      0  0:02:11  0:02:11 --:--:--  839k
```

Vagrant 2.0.0 for Mac
```
fanhonglingdeMacBook-Pro:~ fanhongling$ curl -jkSL https://releases.hashicorp.com/vagrant/2.0.0/vagrant_2.0.0_x86_64.dmg -o Downloads/99-mirror/darwin-packages/vagrant_2.0.0_x86_64.dmg
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 72.0M  100 72.0M    0     0  1175k      0  0:01:02  0:01:02 --:--:--  971k
```

vagrant
```
fanhonglingdeMacBook-Pro:origin fanhongling$ vagrant --version
Vagrant 1.8.1
```

virtualbox
```
fanhonglingdeMacBook-Pro:origin fanhongling$ /Applications/VirtualBox.app/Contents/MacOS/VBoxManage --version
5.0.20r106931
```

## Fedora 26 VirtualBox for Mac

### Up

Common problem: _eth0_ or _enp0s3_ configured incorrectly when set auto_config with true in Vagrantfile 
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Box 'fedora26-cloud-base' could not be found. Attempting to find and install...
    default: Box Provider: virtualbox
    default: Box Version: >= 0
==> default: Box file was not detected as metadata. Adding it directly...
==> default: Adding box 'fedora26-cloud-base' (v0) for provider: virtualbox
    default: Downloading: https://download.fedoraproject.org/pub/fedora/linux/releases/26/CloudImages/x86_64/images/Fedora-Cloud-Base-Vagrant-26-1.5.x86_64.vagrant-virtualbox.box
==> default: Successfully added box 'fedora26-cloud-base' (v0) for 'virtualbox'!
==> default: Importing base box 'fedora26-cloud-base'...
==> default: Matching MAC address for NAT networking...
==> default: Setting the name of the VM: go-to-kubernetes_default_1513507231694_54213
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
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Remote connection disconnect. Retrying...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
    default: No guest additions were detected on the base box for this VM! Guest
    default: additions are required for forwarded ports, shared folders, host only
    default: networking, and more. If SSH fails on this machine, please install
    default: the guest additions and repackage the box to continue.
    default: 
    default: This is not an error message; everything may continue to work properly,
    default: in which case you may ignore this message.
==> default: Configuring and enabling network interfaces...
^C==> default: Waiting for cleanup before exiting...
^C==> default: Exiting immediately, without cleanup!
```
Enter into VirtualBox Guest window then manually configure it, set auto_config with false

Common problem: _eth0_ or _enp0s3_ configured incorrectly 
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant reload
==> default: Attempting graceful shutdown of VM...
    default: Guest communication could not be established! This is usually because
    default: SSH is not running, the authentication information was changed,
    default: or some other networking issue. Vagrant will force halt, if
    default: capable.
==> default: Forcing shutdown of VM...
==> default: Clearing any previously set forwarded ports...
==> default: Fixed port collision for 22 => 2222. Now on port 2201.
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
    default: Adapter 3: hostonly
==> default: Forwarding ports...
    default: 22 (guest) => 2201 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2201
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Remote connection disconnect. Retrying...
Timed out while waiting for the machine to boot. This means that
Vagrant was unable to communicate with the guest machine within
the configured ("config.vm.boot_timeout" value) time period.

If you look above, you should be able to see the error(s) that
Vagrant had when attempting to connect to the machine. These errors
are usually good hints as to what may be wrong.

If you're using a custom box, make sure that networking is properly
working and you're able to connect to the machine. It is a common
problem that networking isn't setup properly in these boxes.
Verify that authentication configurations are also setup properly,
as well.

If the box appears to be booting properly, you may want to increase
the timeout ("config.vm.boot_timeout") value.
```
Enter into VirtualBox Guest window then manually configure it, set auto_config with false

vboxsf had not installed, 
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant reload
==> default: Attempting graceful shutdown of VM...
    default: Guest communication could not be established! This is usually because
    default: SSH is not running, the authentication information was changed,
    default: or some other networking issue. Vagrant will force halt, if
    default: capable.
==> default: Forcing shutdown of VM...
==> default: Clearing any previously set forwarded ports...
==> default: Fixed port collision for 22 => 2222. Now on port 2201.
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
    default: Adapter 3: hostonly
==> default: Forwarding ports...
    default: 22 (guest) => 2201 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2201
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Remote connection disconnect. Retrying...
Timed out while waiting for the machine to boot. This means that
Vagrant was unable to communicate with the guest machine within
the configured ("config.vm.boot_timeout" value) time period.

If you look above, you should be able to see the error(s) that
Vagrant had when attempting to connect to the machine. These errors
are usually good hints as to what may be wrong.

If you're using a custom box, make sure that networking is properly
working and you're able to connect to the machine. It is a common
problem that networking isn't setup properly in these boxes.
Verify that authentication configurations are also setup properly,
as well.

If the box appears to be booting properly, you may want to increase
the timeout ("config.vm.boot_timeout") value.
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant ssh
ssh_exchange_identification: Connection closed by remote host
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant reload
==> default: Clearing any previously set forwarded ports...
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
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Remote connection disconnect. Retrying...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
    default: No guest additions were detected on the base box for this VM! Guest
    default: additions are required for forwarded ports, shared folders, host only
    default: networking, and more. If SSH fails on this machine, please install
    default: the guest additions and repackage the box to continue.
    default: 
    default: This is not an error message; everything may continue to work properly,
    default: in which case you may ignore this message.
==> default: Mounting shared folders...
    default: /Users/fanhongling/go => /Users/fanhongling/go
Failed to mount folders in Linux guest. This is usually because
the "vboxsf" file system is not available. Please verify that
the guest additions are properly installed in the guest and
can work properly. The command attempted was:

mount -t vboxsf -o uid=`id -u vagrant`,gid=`getent group vagrant | cut -d: -f3` Users_fanhongling_go /Users/fanhongling/go
mount -t vboxsf -o uid=`id -u vagrant`,gid=`id -g vagrant` Users_fanhongling_go /Users/fanhongling/go

The error output from the last command was:

mount: unknown filesystem type 'vboxsf'
```
Mount VBoxGuestAdditions.ISO from VirtualBox Window, then login into VM and make installation

__vboxsf__

halt
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant halt
==> default: Attempting graceful shutdown of VM...
``` 

link
```
fanhonglingdeMacBook-Pro:~ fanhongling$ ln -s /Applications/VirtualBox.app/Contents/MacOS/VBoxGuestAdditions.iso 
```

Mount ISO like ![屏幕快照 2017-12-17 上午5.47.12.png](./屏幕快照%202017-12-17%20上午5.47.12.png)


Re-enter with Warning
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Clearing any previously set forwarded ports...
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
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Remote connection disconnect. Retrying...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
    default: No guest additions were detected on the base box for this VM! Guest
    default: additions are required for forwarded ports, shared folders, host only
    default: networking, and more. If SSH fails on this machine, please install
    default: the guest additions and repackage the box to continue.
    default: 
    default: This is not an error message; everything may continue to work properly,
    default: in which case you may ignore this message.
==> default: Mounting shared folders...
    default: /Users/fanhongling/go => /Users/fanhongling/go
Failed to mount folders in Linux guest. This is usually because
the "vboxsf" file system is not available. Please verify that
the guest additions are properly installed in the guest and
can work properly. The command attempted was:

mount -t vboxsf -o uid=`id -u vagrant`,gid=`getent group vagrant | cut -d: -f3` Users_fanhongling_go /Users/fanhongling/go
mount -t vboxsf -o uid=`id -u vagrant`,gid=`id -g vagrant` Users_fanhongling_go /Users/fanhongling/go

The error output from the last command was:

mount: unknown filesystem type 'vboxsf'
```

ssh
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant ssh
Last login: Sun Dec 17 13:13:13 2017
```

kernel
```
[vagrant@localhost ~]$ uname -r
4.11.8-300.fc26.x86_64
```

kernel-devel
```
[vagrant@localhost ~]$ sudo dnf --showduplicates search kernel-devel
Failed to set locale, defaulting to C
Fedora 26 - x86_64 - Updates                                                                                    1.5 MB/s |  18 MB     00:12    
Fedora 26 - x86_64                                                                                              638 kB/s |  53 MB     01:25    
Last metadata expiration check: 0:00:11 ago on Sun Dec 17 13:53:12 2017.
====================================================== Name Exactly Matched: kernel-devel ======================================================
kernel-devel-4.14.5-200.fc26.x86_64 : Development package for building kernel modules to match the kernel
kernel-devel-4.11.8-300.fc26.x86_64 : Development package for building kernel modules to match the kernel
```

```
[vagrant@localhost ~]$ sudo dnf install -yq kernel-devel-4.11.8-300.fc26
Failed to set locale, defaulting to C
```

For detailed content, see CentOS section

kernel-headers
```
[vagrant@localhost ~]$ ls /usr/src/kernels           
4.11.8-300.fc26.x86_64
```

gcc
```
[vagrant@localhost ~]$ sudo dnf install -yq gcc                         
Failed to set locale, defaulting to C
```

For detailed content, see CentOS section

Mount ISO
```
[vagrant@localhost ~]$ sudo mount /dev/sr0 /media/
mount: /dev/sr0 is write-protected, mounting read-only
```

Issue
```
[vagrant@localhost media]$ sudo -i
[root@localhost ~]# cd /media
[root@localhost media]# ./VBoxLinuxAdditions.run --nox11
Verifying archive integrity... All good.
Uncompressing VirtualBox 5.0.20 Guest Additions for Linux............
VirtualBox Guest Additions installer
Removing installed version 5.0.20 of VirtualBox Guest Additions...
Removing existing VirtualBox non-DKMS kernel modules[  OK  ]
Copying additional installer modules ...
Installing additional modules ...
Removing existing VirtualBox non-DKMS kernel modules[  OK  ]
Building the VirtualBox Guest Additions kernel modules
Building the main Guest Additions module[FAILED]
(Look at /var/log/vboxadd-install.log to find out what went wrong)
Doing non-kernel setup of the Guest Additions[  OK  ]
```

```
[root@localhost media]# dnf install -yq dkms
Failed to set locale, defaulting to C
```

Even failed to install, may be required latest VirtualBox version
```
fanhonglingdeMacBook-Pro:~ fanhongling$ /Applications/VirtualBox.app/Contents/MacOS/VBoxManage --version
5.1.30r118389
```

```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant --version
Vagrant 2.0.0
```

up
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Clearing any previously set forwarded ports...
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
    default: Adapter 3: hostonly
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
    default: No guest additions were detected on the base box for this VM! Guest
    default: additions are required for forwarded ports, shared folders, host only
    default: networking, and more. If SSH fails on this machine, please install
    default: the guest additions and repackage the box to continue.
    default: 
    default: This is not an error message; everything may continue to work properly,
    default: in which case you may ignore this message.
==> default: Mounting shared folders...
    default: /Users/fanhongling/go => /Users/fanhongling/go
Vagrant was unable to mount VirtualBox shared folders. This is usually
because the filesystem "vboxsf" is not available. This filesystem is
made available via the VirtualBox Guest Additions and kernel module.
Please verify that these guest additions are properly installed in the
guest. This is not a bug in Vagrant and is usually caused by a faulty
Vagrant box. For context, the command attempted was:

mount -t vboxsf -o uid=1000,gid=1000 Users_fanhongling_go /Users/fanhongling/go

The error output from the command was:

/sbin/mount.vboxsf: mounting failed with the error: No such device

```

ssh
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ ssh vagrant@172.17.4.59
The authenticity of host '172.17.4.59 (172.17.4.59)' can't be established.
RSA key fingerprint is 5e:ef:bc:e8:ff:d2:4d:28:2a:ef:33:0a:2a:5d:ba:b1.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '172.17.4.59' (RSA) to the list of known hosts.
vagrant@172.17.4.59's password: 
Last login: Sun Dec 17 13:51:32 2017 from 10.0.2.2
```

mount new ISO
```
[vagrant@localhost ~]$ sudo mount /dev/sr0 /media
mount: /dev/sr0 is write-protected, mounting read-only
```

Install
```
[vagrant@localhost ~]$ sudo /media/VBoxLinuxAdditions.run --nox11
Verifying archive integrity... All good.
Uncompressing VirtualBox 5.1.30 Guest Additions for Linux...........
VirtualBox Guest Additions installer
Removing installed version 5.1.30 of VirtualBox Guest Additions...
Copying additional installer modules ...
Installing additional modules ...
vboxadd.sh: Starting the VirtualBox Guest Additions.

Could not find the X.Org or XFree86 Window System, skipping.
```

Exit
```
[vagrant@localhost ~]$ exit
logout
Connection to 172.17.4.59 closed.
```

Reload
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant reload
==> default: Attempting graceful shutdown of VM...
==> default: Clearing any previously set forwarded ports...
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
    default: Adapter 3: hostonly
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /Users/fanhongling/go => /Users/fanhongling/go
    default: /Users/fanhongling/Downloads => /Users/fanhongling/Downloads
```

ssh
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ ssh vagrant@172.17.4.59
vagrant@172.17.4.59's password: 
Last login: Sun Dec 17 14:44:58 2017 from 172.17.4.1
```

test
```
[vagrant@localhost ~]$ ls -l /Users/fanhongling/go/
total 0
drwxr-xr-x. 1 vagrant vagrant 748 Dec 17 08:27 bin
drwxr-xr-x. 1 vagrant vagrant 204 Dec 14 01:26 pkg
drwxr-xr-x. 1 vagrant vagrant 306 Oct 14 20:07 src
```

git
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo dnf install -yq git
Failed to set locale, defaulting to C
```

```
[vagrant@kubedev-172-17-4-59 ~]$ git --version
git version 2.13.6
```

change password
```
[vagrant@localhost ~]$ passwd    
Changing password for user vagrant.
Current password: 
New password: 
Retype new password: 
passwd: all authentication tokens updated successfully.
```

## CentOS 7.4 VirtualBox for Windows 10

vagrant
```
tangf@DESKTOP-H68OQDV /cygdrive/g/work/src/github.com/tangfeixiong/go-to-kubernetes
$ vagrant --version
Vagrant 1.8.6
```

virtualbox
```
tangf@DESKTOP-H68OQDV ~
$ /cygdrive/c/Program\ Files/Oracle/VirtualBox/VBoxManage.exe --version
5.1.30r118389
```

### Up

spawn
```
tangf@DESKTOP-H68OQDV /cygdrive/g/work/src/github.com/tangfeixiong/go-to-kubernetes
$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Checking if box 'centos/7' is up to date...
==> default: Clearing any previously set forwarded ports...
==> default: Fixed port collision for 22 => 2222. Now on port 2200.
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
    default: Adapter 3: hostonly
==> default: Forwarding ports...
    default: 8080 (guest) => 18080 (host) (adapter 1)
    default: 443 (guest) => 1443 (host) (adapter 1)
    default: 22 (guest) => 2200 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
### snippets ###
```

enter
```
tangf@DESKTOP-H68OQDV /cygdrive/g/work/src/github.com/tangfeixiong/go-to-kubernetes
$ vagrant ssh
### snippets ###
```

### Reload

After network interface configured for all ethX, and lot of works have done...

modify Vagrantfile network setting to disable auto_config, because it may work confused with NetworkManager.service
```
tangf@DESKTOP-H68OQDV /cygdrive/g/work/src/github.com/tangfeixiong/go-to-kubernetes
$ vagrant reload
==> default: Attempting graceful shutdown of VM...
==> default: Checking if box 'centos/7' is up to date...
==> default: Clearing any previously set forwarded ports...
==> default: Fixed port collision for 22 => 2222. Now on port 2200.
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
    default: Adapter 3: hostonly
==> default: Forwarding ports...
    default: 8080 (guest) => 18080 (host) (adapter 1)
    default: 443 (guest) => 1443 (host) (adapter 1)
    default: 22 (guest) => 2200 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2200
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Remote connection disconnect. Retrying...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
    default: No guest additions were detected on the base box for this VM! Guest
    default: additions are required for forwarded ports, shared folders, host only
    default: networking, and more. If SSH fails on this machine, please install
    default: the guest additions and repackage the box to continue.
    default:
    default: This is not an error message; everything may continue to work properly,
    default: in which case you may ignore this message.
==> default: Machine already provisioned. Run `vagrant provision` or use the `--provision`
==> default: flag to force provisioning. Provisioners marked to run always will still run.

```

### vboxsf

kernel version
```
[vagrant@kubedev-10-64-33-82 ~]$ uname -r
3.10.0-693.5.2.el7.x86_64
```

kernel-devel packages
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo yum search --showduplicates kernel-devel
Loaded plugins: fastestmirror
Repodata is over 2 weeks old. Install yum-cron? Or run: yum makecache fast
Determining fastest mirrors
 * base: centos.ustc.edu.cn
 * extras: mirrors.aliyun.com
 * updates: centos.ustc.edu.cn
========================== N/S matched: kernel-devel ===========================
kernel-devel-3.10.0-693.el7.x86_64 : Development package for building kernel
                                   : modules to match the kernel
kernel-devel-3.10.0-693.1.1.el7.x86_64 : Development package for building kernel
                                       : modules to match the kernel
kernel-devel-3.10.0-693.2.1.el7.x86_64 : Development package for building kernel
                                       : modules to match the kernel
kernel-devel-3.10.0-693.2.2.el7.x86_64 : Development package for building kernel
                                       : modules to match the kernel
kernel-devel-3.10.0-693.5.2.el7.x86_64 : Development package for building kernel
                                       : modules to match the kernel

  Name and summary matches only, use "search all" for everything.
```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo yum install -y kernel-devel-3.10.0-693.5.2.el7
Loaded plugins: fastestmirror
base                                                     | 3.6 kB     00:00
extras                                                   | 3.4 kB     00:00
updates                                                  | 3.4 kB     00:00
(1/2): extras/7/x86_64/primary_db                          | 145 kB   00:01
(2/2): updates/7/x86_64/primary_db                         | 4.5 MB   00:01
Loading mirror speeds from cached hostfile
 * base: centos.ustc.edu.cn
 * extras: mirrors.aliyun.com
 * updates: centos.ustc.edu.cn
Resolving Dependencies
--> Running transaction check
---> Package kernel-devel.x86_64 0:3.10.0-693.5.2.el7 will be installed
--> Processing Dependency: perl for package: kernel-devel-3.10.0-693.5.2.el7.x86_64
--> Running transaction check
---> Package perl.x86_64 4:5.16.3-292.el7 will be installed
--> Processing Dependency: perl-libs = 4:5.16.3-292.el7 for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Socket) >= 1.3 for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Scalar::Util) >= 1.10 for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl-macros for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl-libs for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(threads::shared) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(threads) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(constant) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Time::Local) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Time::HiRes) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Storable) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Socket) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Scalar::Util) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Pod::Simple::XHTML) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Pod::Simple::Search) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Getopt::Long) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Filter::Util::Call) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(File::Temp) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(File::Spec::Unix) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(File::Spec::Functions) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(File::Spec) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(File::Path) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Exporter) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Cwd) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Carp) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: libperl.so()(64bit) for package: 4:perl-5.16.3-292.el7.x86_64
--> Running transaction check
---> Package perl-Carp.noarch 0:1.26-244.el7 will be installed
---> Package perl-Exporter.noarch 0:5.68-3.el7 will be installed
---> Package perl-File-Path.noarch 0:2.09-2.el7 will be installed
---> Package perl-File-Temp.noarch 0:0.23.01-3.el7 will be installed
---> Package perl-Filter.x86_64 0:1.49-3.el7 will be installed
---> Package perl-Getopt-Long.noarch 0:2.40-2.el7 will be installed
--> Processing Dependency: perl(Pod::Usage) >= 1.14 for package: perl-Getopt-Long-2.40-2.el7.noarch
--> Processing Dependency: perl(Text::ParseWords) for package: perl-Getopt-Long-2.40-2.el7.noarch
---> Package perl-PathTools.x86_64 0:3.40-5.el7 will be installed
---> Package perl-Pod-Simple.noarch 1:3.28-4.el7 will be installed
--> Processing Dependency: perl(Pod::Escapes) >= 1.04 for package: 1:perl-Pod-Simple-3.28-4.el7.noarch
--> Processing Dependency: perl(Encode) for package: 1:perl-Pod-Simple-3.28-4.el7.noarch
---> Package perl-Scalar-List-Utils.x86_64 0:1.27-248.el7 will be installed
---> Package perl-Socket.x86_64 0:2.010-4.el7 will be installed
---> Package perl-Storable.x86_64 0:2.45-3.el7 will be installed
---> Package perl-Time-HiRes.x86_64 4:1.9725-3.el7 will be installed
---> Package perl-Time-Local.noarch 0:1.2300-2.el7 will be installed
---> Package perl-constant.noarch 0:1.27-2.el7 will be installed
---> Package perl-libs.x86_64 4:5.16.3-292.el7 will be installed
---> Package perl-macros.x86_64 4:5.16.3-292.el7 will be installed
---> Package perl-threads.x86_64 0:1.87-4.el7 will be installed
---> Package perl-threads-shared.x86_64 0:1.43-6.el7 will be installed
--> Running transaction check
---> Package perl-Encode.x86_64 0:2.51-7.el7 will be installed
---> Package perl-Pod-Escapes.noarch 1:1.04-292.el7 will be installed
---> Package perl-Pod-Usage.noarch 0:1.63-3.el7 will be installed
--> Processing Dependency: perl(Pod::Text) >= 3.15 for package: perl-Pod-Usage-1.63-3.el7.noarch
--> Processing Dependency: perl-Pod-Perldoc for package: perl-Pod-Usage-1.63-3.el7.noarch
---> Package perl-Text-ParseWords.noarch 0:3.29-4.el7 will be installed
--> Running transaction check
---> Package perl-Pod-Perldoc.noarch 0:3.20-4.el7 will be installed
--> Processing Dependency: perl(parent) for package: perl-Pod-Perldoc-3.20-4.el7.noarch
--> Processing Dependency: perl(HTTP::Tiny) for package: perl-Pod-Perldoc-3.20-4.el7.noarch
---> Package perl-podlators.noarch 0:2.5.1-3.el7 will be installed
--> Running transaction check
---> Package perl-HTTP-Tiny.noarch 0:0.033-3.el7 will be installed
---> Package perl-parent.noarch 1:0.225-244.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package                    Arch       Version                Repository   Size
================================================================================
Installing:
 kernel-devel               x86_64     3.10.0-693.5.2.el7     updates      14 M
Installing for dependencies:
 perl                       x86_64     4:5.16.3-292.el7       base        8.0 M
 perl-Carp                  noarch     1.26-244.el7           base         19 k
 perl-Encode                x86_64     2.51-7.el7             base        1.5 M
 perl-Exporter              noarch     5.68-3.el7             base         28 k
 perl-File-Path             noarch     2.09-2.el7             base         26 k
 perl-File-Temp             noarch     0.23.01-3.el7          base         56 k
 perl-Filter                x86_64     1.49-3.el7             base         76 k
 perl-Getopt-Long           noarch     2.40-2.el7             base         56 k
 perl-HTTP-Tiny             noarch     0.033-3.el7            base         38 k
 perl-PathTools             x86_64     3.40-5.el7             base         82 k
 perl-Pod-Escapes           noarch     1:1.04-292.el7         base         51 k
 perl-Pod-Perldoc           noarch     3.20-4.el7             base         87 k
 perl-Pod-Simple            noarch     1:3.28-4.el7           base        216 k
 perl-Pod-Usage             noarch     1.63-3.el7             base         27 k
 perl-Scalar-List-Utils     x86_64     1.27-248.el7           base         36 k
 perl-Socket                x86_64     2.010-4.el7            base         49 k
 perl-Storable              x86_64     2.45-3.el7             base         77 k
 perl-Text-ParseWords       noarch     3.29-4.el7             base         14 k
 perl-Time-HiRes            x86_64     4:1.9725-3.el7         base         45 k
 perl-Time-Local            noarch     1.2300-2.el7           base         24 k
 perl-constant              noarch     1.27-2.el7             base         19 k
 perl-libs                  x86_64     4:5.16.3-292.el7       base        688 k
 perl-macros                x86_64     4:5.16.3-292.el7       base         43 k
 perl-parent                noarch     1:0.225-244.el7        base         12 k
 perl-podlators             noarch     2.5.1-3.el7            base        112 k
 perl-threads               x86_64     1.87-4.el7             base         49 k
 perl-threads-shared        x86_64     1.43-6.el7             base         39 k

Transaction Summary
================================================================================
Install  1 Package (+27 Dependent packages)

Total download size: 26 M
Installed size: 72 M
Downloading packages:
updates/7/x86_64/prestodelta                               | 589 kB   00:00
(1/28): perl-Carp-1.26-244.el7.noarch.rpm                  |  19 kB   00:01
(2/28): perl-Exporter-5.68-3.el7.noarch.rpm                |  28 kB   00:01
(3/28): perl-File-Path-2.09-2.el7.noarch.rpm               |  26 kB   00:00
(4/28): kernel-devel-3.10.0-693.5.2.el7.x86_64.rpm         |  14 MB   00:05
(5/28): perl-Getopt-Long-2.40-2.el7.noarch.rpm             |  56 kB   00:00
(6/28): perl-HTTP-Tiny-0.033-3.el7.noarch.rpm              |  38 kB   00:00
(7/28): perl-PathTools-3.40-5.el7.x86_64.rpm               |  82 kB   00:00
(8/28): perl-Pod-Escapes-1.04-292.el7.noarch.rpm           |  51 kB   00:00
(9/28): perl-Pod-Perldoc-3.20-4.el7.noarch.rpm             |  87 kB   00:00
(10/28): perl-Pod-Simple-3.28-4.el7.noarch.rpm             | 216 kB   00:00
(11/28): perl-Pod-Usage-1.63-3.el7.noarch.rpm              |  27 kB   00:00
(12/28): perl-Scalar-List-Utils-1.27-248.el7.x86_64.rpm    |  36 kB   00:00
(13/28): perl-File-Temp-0.23.01-3.el7.noarch.rpm           |  56 kB   00:04
(14/28): perl-Socket-2.010-4.el7.x86_64.rpm                |  49 kB   00:00
(15/28): perl-Text-ParseWords-3.29-4.el7.noarch.rpm        |  14 kB   00:00
(16/28): perl-Time-HiRes-1.9725-3.el7.x86_64.rpm           |  45 kB   00:00
(17/28): perl-Time-Local-1.2300-2.el7.noarch.rpm           |  24 kB   00:00
(18/28): perl-constant-1.27-2.el7.noarch.rpm               |  19 kB   00:00
(19/28): perl-5.16.3-292.el7.x86_64.rpm                    | 8.0 MB   00:08
(20/28): perl-macros-5.16.3-292.el7.x86_64.rpm             |  43 kB   00:00
(21/28): perl-parent-0.225-244.el7.noarch.rpm              |  12 kB   00:00
(22/28): perl-Storable-2.45-3.el7.x86_64.rpm               |  77 kB   00:01
(23/28): perl-podlators-2.5.1-3.el7.noarch.rpm             | 112 kB   00:00
(24/28): perl-threads-shared-1.43-6.el7.x86_64.rpm         |  39 kB   00:00
(25/28): perl-threads-1.87-4.el7.x86_64.rpm                |  49 kB   00:00
(26/28): perl-libs-5.16.3-292.el7.x86_64.rpm               | 688 kB   00:01
(27/28): perl-Filter-1.49-3.el7.x86_64.rpm                 |  76 kB   00:09
(28/28): perl-Encode-2.51-7.el7.x86_64.rpm                 | 1.5 MB   00:11
--------------------------------------------------------------------------------
Total                                              2.2 MB/s |  26 MB  00:11
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : 1:perl-parent-0.225-244.el7.noarch                          1/28
  Installing : perl-HTTP-Tiny-0.033-3.el7.noarch                           2/28
  Installing : perl-podlators-2.5.1-3.el7.noarch                           3/28
  Installing : perl-Pod-Perldoc-3.20-4.el7.noarch                          4/28
  Installing : 1:perl-Pod-Escapes-1.04-292.el7.noarch                      5/28
  Installing : perl-Text-ParseWords-3.29-4.el7.noarch                      6/28
  Installing : perl-Encode-2.51-7.el7.x86_64                               7/28
  Installing : perl-Pod-Usage-1.63-3.el7.noarch                            8/28
  Installing : 4:perl-macros-5.16.3-292.el7.x86_64                         9/28
  Installing : 4:perl-libs-5.16.3-292.el7.x86_64                          10/28
  Installing : perl-Storable-2.45-3.el7.x86_64                            11/28
  Installing : perl-Exporter-5.68-3.el7.noarch                            12/28
  Installing : perl-constant-1.27-2.el7.noarch                            13/28
  Installing : perl-Time-Local-1.2300-2.el7.noarch                        14/28
  Installing : perl-Socket-2.010-4.el7.x86_64                             15/28
  Installing : perl-Carp-1.26-244.el7.noarch                              16/28
  Installing : 4:perl-Time-HiRes-1.9725-3.el7.x86_64                      17/28
  Installing : perl-PathTools-3.40-5.el7.x86_64                           18/28
  Installing : perl-Scalar-List-Utils-1.27-248.el7.x86_64                 19/28
  Installing : perl-File-Temp-0.23.01-3.el7.noarch                        20/28
  Installing : perl-File-Path-2.09-2.el7.noarch                           21/28
  Installing : perl-threads-shared-1.43-6.el7.x86_64                      22/28
  Installing : perl-threads-1.87-4.el7.x86_64                             23/28
  Installing : perl-Filter-1.49-3.el7.x86_64                              24/28
  Installing : 1:perl-Pod-Simple-3.28-4.el7.noarch                        25/28
  Installing : perl-Getopt-Long-2.40-2.el7.noarch                         26/28
  Installing : 4:perl-5.16.3-292.el7.x86_64                               27/28
  Installing : kernel-devel-3.10.0-693.5.2.el7.x86_64                     28/28
  Verifying  : perl-HTTP-Tiny-0.033-3.el7.noarch                           1/28
  Verifying  : perl-threads-shared-1.43-6.el7.x86_64                       2/28
  Verifying  : perl-Storable-2.45-3.el7.x86_64                             3/28
  Verifying  : perl-Exporter-5.68-3.el7.noarch                             4/28
  Verifying  : perl-constant-1.27-2.el7.noarch                             5/28
  Verifying  : perl-PathTools-3.40-5.el7.x86_64                            6/28
  Verifying  : 4:perl-macros-5.16.3-292.el7.x86_64                         7/28
  Verifying  : 1:perl-parent-0.225-244.el7.noarch                          8/28
  Verifying  : 4:perl-5.16.3-292.el7.x86_64                                9/28
  Verifying  : perl-File-Temp-0.23.01-3.el7.noarch                        10/28
  Verifying  : 1:perl-Pod-Simple-3.28-4.el7.noarch                        11/28
  Verifying  : perl-Time-Local-1.2300-2.el7.noarch                        12/28
  Verifying  : 4:perl-libs-5.16.3-292.el7.x86_64                          13/28
  Verifying  : perl-Pod-Perldoc-3.20-4.el7.noarch                         14/28
  Verifying  : perl-Socket-2.010-4.el7.x86_64                             15/28
  Verifying  : kernel-devel-3.10.0-693.5.2.el7.x86_64                     16/28
  Verifying  : perl-Carp-1.26-244.el7.noarch                              17/28
  Verifying  : 4:perl-Time-HiRes-1.9725-3.el7.x86_64                      18/28
  Verifying  : perl-Scalar-List-Utils-1.27-248.el7.x86_64                 19/28
  Verifying  : 1:perl-Pod-Escapes-1.04-292.el7.noarch                     20/28
  Verifying  : perl-Pod-Usage-1.63-3.el7.noarch                           21/28
  Verifying  : perl-Encode-2.51-7.el7.x86_64                              22/28
  Verifying  : perl-podlators-2.5.1-3.el7.noarch                          23/28
  Verifying  : perl-Getopt-Long-2.40-2.el7.noarch                         24/28
  Verifying  : perl-File-Path-2.09-2.el7.noarch                           25/28
  Verifying  : perl-threads-1.87-4.el7.x86_64                             26/28
  Verifying  : perl-Filter-1.49-3.el7.x86_64                              27/28
  Verifying  : perl-Text-ParseWords-3.29-4.el7.noarch                     28/28

Installed:
  kernel-devel.x86_64 0:3.10.0-693.5.2.el7

Dependency Installed:
  perl.x86_64 4:5.16.3-292.el7
  perl-Carp.noarch 0:1.26-244.el7
  perl-Encode.x86_64 0:2.51-7.el7
  perl-Exporter.noarch 0:5.68-3.el7
  perl-File-Path.noarch 0:2.09-2.el7
  perl-File-Temp.noarch 0:0.23.01-3.el7
  perl-Filter.x86_64 0:1.49-3.el7
  perl-Getopt-Long.noarch 0:2.40-2.el7
  perl-HTTP-Tiny.noarch 0:0.033-3.el7
  perl-PathTools.x86_64 0:3.40-5.el7
  perl-Pod-Escapes.noarch 1:1.04-292.el7
  perl-Pod-Perldoc.noarch 0:3.20-4.el7
  perl-Pod-Simple.noarch 1:3.28-4.el7
  perl-Pod-Usage.noarch 0:1.63-3.el7
  perl-Scalar-List-Utils.x86_64 0:1.27-248.el7
  perl-Socket.x86_64 0:2.010-4.el7
  perl-Storable.x86_64 0:2.45-3.el7
  perl-Text-ParseWords.noarch 0:3.29-4.el7
  perl-Time-HiRes.x86_64 4:1.9725-3.el7
  perl-Time-Local.noarch 0:1.2300-2.el7
  perl-constant.noarch 0:1.27-2.el7
  perl-libs.x86_64 4:5.16.3-292.el7
  perl-macros.x86_64 4:5.16.3-292.el7
  perl-parent.noarch 1:0.225-244.el7
  perl-podlators.noarch 0:2.5.1-3.el7
  perl-threads.x86_64 0:1.87-4.el7
  perl-threads-shared.x86_64 0:1.43-6.el7

Complete!
```

```
[vagrant@kubedev-10-64-33-82 ~]$ ls /usr/src/kernels/`uname -r`
arch     firmware  ipc      Makefile        net       sound       virt
block    fs        Kconfig  Makefile.qlock  samples   System.map  vmlinux.id
crypto   include   kernel   mm              scripts   tools
drivers  init      lib      Module.symvers  security  usr
```

gcc
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo yum install -y gcc
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: centos.ustc.edu.cn
 * extras: mirrors.aliyun.com
 * updates: centos.ustc.edu.cn
Resolving Dependencies
--> Running transaction check
---> Package gcc.x86_64 0:4.8.5-16.el7_4.1 will be installed
--> Processing Dependency: libgomp = 4.8.5-16.el7_4.1 for package: gcc-4.8.5-16.el7_4.1.x86_64
--> Processing Dependency: cpp = 4.8.5-16.el7_4.1 for package: gcc-4.8.5-16.el7_4.1.x86_64
--> Processing Dependency: libgcc >= 4.8.5-16.el7_4.1 for package: gcc-4.8.5-16.el7_4.1.x86_64
--> Processing Dependency: glibc-devel >= 2.2.90-12 for package: gcc-4.8.5-16.el7_4.1.x86_64
--> Processing Dependency: libmpfr.so.4()(64bit) for package: gcc-4.8.5-16.el7_4.1.x86_64
--> Processing Dependency: libmpc.so.3()(64bit) for package: gcc-4.8.5-16.el7_4.1.x86_64
--> Running transaction check
---> Package cpp.x86_64 0:4.8.5-16.el7_4.1 will be installed
---> Package glibc-devel.x86_64 0:2.17-196.el7_4.2 will be installed
--> Processing Dependency: glibc-headers = 2.17-196.el7_4.2 for package: glibc-devel-2.17-196.el7_4.2.x86_64
--> Processing Dependency: glibc = 2.17-196.el7_4.2 for package: glibc-devel-2.17-196.el7_4.2.x86_64
--> Processing Dependency: glibc-headers for package: glibc-devel-2.17-196.el7_4.2.x86_64
---> Package libgcc.x86_64 0:4.8.5-16.el7 will be updated
---> Package libgcc.x86_64 0:4.8.5-16.el7_4.1 will be an update
---> Package libgomp.x86_64 0:4.8.5-16.el7 will be updated
---> Package libgomp.x86_64 0:4.8.5-16.el7_4.1 will be an update
---> Package libmpc.x86_64 0:1.0.1-3.el7 will be installed
---> Package mpfr.x86_64 0:3.1.1-4.el7 will be installed
--> Running transaction check
---> Package glibc.x86_64 0:2.17-196.el7 will be updated
--> Processing Dependency: glibc = 2.17-196.el7 for package: glibc-common-2.17-196.el7.x86_64
---> Package glibc.x86_64 0:2.17-196.el7_4.2 will be an update
---> Package glibc-headers.x86_64 0:2.17-196.el7_4.2 will be installed
--> Processing Dependency: kernel-headers >= 2.2.1 for package: glibc-headers-2.17-196.el7_4.2.x86_64
--> Processing Dependency: kernel-headers for package: glibc-headers-2.17-196.el7_4.2.x86_64
--> Running transaction check
---> Package glibc-common.x86_64 0:2.17-196.el7 will be updated
---> Package glibc-common.x86_64 0:2.17-196.el7_4.2 will be an update
---> Package kernel-headers.x86_64 0:3.10.0-693.11.1.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package             Arch        Version                     Repository    Size
================================================================================
Installing:
 gcc                 x86_64      4.8.5-16.el7_4.1            updates       16 M
Installing for dependencies:
 cpp                 x86_64      4.8.5-16.el7_4.1            updates      5.9 M
 glibc-devel         x86_64      2.17-196.el7_4.2            updates      1.1 M
 glibc-headers       x86_64      2.17-196.el7_4.2            updates      676 k
 kernel-headers      x86_64      3.10.0-693.11.1.el7         updates      6.0 M
 libmpc              x86_64      1.0.1-3.el7                 base          51 k
 mpfr                x86_64      3.1.1-4.el7                 base         203 k
Updating for dependencies:
 glibc               x86_64      2.17-196.el7_4.2            updates      3.6 M
 glibc-common        x86_64      2.17-196.el7_4.2            updates       11 M
 libgcc              x86_64      4.8.5-16.el7_4.1            updates       98 k
 libgomp             x86_64      4.8.5-16.el7_4.1            updates      154 k

Transaction Summary
================================================================================
Install  1 Package  (+6 Dependent packages)
Upgrade             ( 4 Dependent packages)

Total download size: 45 M
Downloading packages:
Delta RPMs reduced 3.8 M of updates to 726 k (81% saved)
(1/11): glibc-2.17-196.el7_2.17-196.el7_4.2.x86_64.drpm    | 643 kB   00:00
(2/11): libgomp-4.8.5-16.el7_4.8.5-16.el7_4.1.x86_64.drpm  |  41 kB   00:01
/usr/share/doc/glibc-2.17/INSTALL: No such file or directory
cannot reconstruct rpm from disk files
/usr/share/doc/libgomp-4.8.5/ChangeLog.bz2: No such file or directory
cannot reconstruct rpm from disk files
(3/11): glibc-common-2.17-196.el7_4.2.x86_64.rpm           |  11 MB   00:02
(4/11): libgcc-4.8.5-16.el7_4.8.5-16.el7_4.1.x86_64.drpm   |  41 kB   00:03
/usr/share/doc/libgcc-4.8.5/COPYING: No such file or directory
cannot reconstruct rpm from disk files
(5/11): glibc-devel-2.17-196.el7_4.2.x86_64.rpm            | 1.1 MB   00:01
(6/11): libmpc-1.0.1-3.el7.x86_64.rpm                      |  51 kB   00:00
(7/11): glibc-headers-2.17-196.el7_4.2.x86_64.rpm          | 676 kB   00:00
(8/11): mpfr-3.1.1-4.el7.x86_64.rpm                        | 203 kB   00:00
(9/11): cpp-4.8.5-16.el7_4.1.x86_64.rpm                    | 5.9 MB   00:07
kernel-headers-3.10.0-693.11.1 FAILED
http://mirror.lzu.edu.cn/centos/7.4.1708/updates/x86_64/Packages/kernel-headers-3.10.0-693.11.1.el7.x86_64.rpm: [Errno 14] curl#56 - "Recv failure: Connection reset by peer"
Trying other mirror.
(10/11): kernel-headers-3.10.0-693.11.1.el7.x86_64.rpm     | 6.0 MB   00:01
(11/11): gcc-4.8.5-16.el7_4.1.x86_64.rpm                   |  16 MB   00:15
Some delta RPMs failed to download or rebuild. Retrying..
(1/3): libgcc-4.8.5-16.el7_4.1.x86_64.rpm                  |  98 kB   00:01
(2/3): glibc-2.17-196.el7_4.2.x86_64.rpm                   | 3.6 MB   00:01
(3/3): libgomp-4.8.5-16.el7_4.1.x86_64.rpm                 | 154 kB   00:01
--------------------------------------------------------------------------------
Total                                              2.6 MB/s |  46 MB  00:17
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Updating   : libgcc-4.8.5-16.el7_4.1.x86_64                              1/15
  Updating   : glibc-common-2.17-196.el7_4.2.x86_64                        2/15
  Updating   : glibc-2.17-196.el7_4.2.x86_64                               3/15
  Installing : mpfr-3.1.1-4.el7.x86_64                                     4/15
  Installing : libmpc-1.0.1-3.el7.x86_64                                   5/15
  Installing : cpp-4.8.5-16.el7_4.1.x86_64                                 6/15
  Updating   : libgomp-4.8.5-16.el7_4.1.x86_64                             7/15
  Installing : kernel-headers-3.10.0-693.11.1.el7.x86_64                   8/15
  Installing : glibc-headers-2.17-196.el7_4.2.x86_64                       9/15
  Installing : glibc-devel-2.17-196.el7_4.2.x86_64                        10/15
  Installing : gcc-4.8.5-16.el7_4.1.x86_64                                11/15
  Cleanup    : libgomp-4.8.5-16.el7.x86_64                                12/15
  Cleanup    : glibc-common-2.17-196.el7.x86_64                           13/15
  Cleanup    : glibc-2.17-196.el7.x86_64                                  14/15
  Cleanup    : libgcc-4.8.5-16.el7.x86_64                                 15/15
  Verifying  : cpp-4.8.5-16.el7_4.1.x86_64                                 1/15
  Verifying  : kernel-headers-3.10.0-693.11.1.el7.x86_64                   2/15
  Verifying  : glibc-devel-2.17-196.el7_4.2.x86_64                         3/15
  Verifying  : mpfr-3.1.1-4.el7.x86_64                                     4/15
  Verifying  : libgomp-4.8.5-16.el7_4.1.x86_64                             5/15
  Verifying  : libgcc-4.8.5-16.el7_4.1.x86_64                              6/15
  Verifying  : glibc-common-2.17-196.el7_4.2.x86_64                        7/15
  Verifying  : glibc-headers-2.17-196.el7_4.2.x86_64                       8/15
  Verifying  : glibc-2.17-196.el7_4.2.x86_64                               9/15
  Verifying  : gcc-4.8.5-16.el7_4.1.x86_64                                10/15
  Verifying  : libmpc-1.0.1-3.el7.x86_64                                  11/15
  Verifying  : libgcc-4.8.5-16.el7.x86_64                                 12/15
  Verifying  : glibc-2.17-196.el7.x86_64                                  13/15
  Verifying  : glibc-common-2.17-196.el7.x86_64                           14/15
  Verifying  : libgomp-4.8.5-16.el7.x86_64                                15/15

Installed:
  gcc.x86_64 0:4.8.5-16.el7_4.1

Dependency Installed:
  cpp.x86_64 0:4.8.5-16.el7_4.1
  glibc-devel.x86_64 0:2.17-196.el7_4.2
  glibc-headers.x86_64 0:2.17-196.el7_4.2
  kernel-headers.x86_64 0:3.10.0-693.11.1.el7
  libmpc.x86_64 0:1.0.1-3.el7
  mpfr.x86_64 0:3.1.1-4.el7

Dependency Updated:
  glibc.x86_64 0:2.17-196.el7_4.2     glibc-common.x86_64 0:2.17-196.el7_4.2
  libgcc.x86_64 0:4.8.5-16.el7_4.1    libgomp.x86_64 0:4.8.5-16.el7_4.1

Complete!
```

Mount VBoxGuestAdditions.iso
```
[vagrant@kubedev-10-64-33-82 ~]$ sudo mount /dev/sr0 /media
```

```
[vagrant@kubedev-10-64-33-82 ~]$ sudo -i
```

```
[root@kubedev-10-64-33-82 ~]# /media/VBoxLinuxAdditions.run --nox11
Verifying archive integrity... All good.
Uncompressing VirtualBox 5.1.30 Guest Additions for Linux...........
VirtualBox Guest Additions installer
Removing installed version 5.1.30 of VirtualBox Guest Additions...
Copying additional installer modules ...
Installing additional modules ...
vboxadd.sh: Starting the VirtualBox Guest Additions.

Could not find the X.Org or XFree86 Window System, skipping.
```

Add synced_folder into Vagrantfile
```
  config.vm.synced_folder "f:/", "/windows10_drive_f"
  config.vm.synced_folder "g:/", "/windows10_drive_g"
```

Exit
```
[root@kubedev-10-64-33-82 ~]# exit
logout
```

```
[vagrant@kubedev-10-64-33-82 ~]$ exit
logout
Connection to 127.0.0.1 closed.
```

Reload
```
tangf@DESKTOP-H68OQDV /cygdrive/g/work/src/github.com/tangfeixiong/go-to-kubernetes
$ vagrant reload
==> default: Attempting graceful shutdown of VM...
==> default: Checking if box 'centos/7' is up to date...
==> default: Clearing any previously set forwarded ports...
==> default: Fixed port collision for 22 => 2222. Now on port 2200.
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
    default: Adapter 3: hostonly
==> default: Forwarding ports...
    default: 8080 (guest) => 18080 (host) (adapter 1)
    default: 443 (guest) => 1443 (host) (adapter 1)
    default: 22 (guest) => 2200 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2200
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Remote connection disconnect. Retrying...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /windows10_drive_f => F:/
    default: /windows10_drive_g => G:/
==> default: Machine already provisioned. Run `vagrant provision` or use the `--provision`
==> default: flag to force provisioning. Provisioners marked to run always will still run.

```
tangf@DESKTOP-H68OQDV /cygdrive/g/work/src/github.com/tangfeixiong/go-to-kubernetes
$ ssh -i ~/vagrant vagrant@10.64.33.82
Last login: Sat Dec 16 00:29:11 2017 from 10.64.33.1
```

```
[vagrant@kubedev-10-64-33-82 ~]$ ls -ld /windows10_drive_g/work/src/github.com/tangfeixiong/go-to-kubernetes/
drwxrwxrwx. 1 vagrant vagrant 4096 Dec 15 23:48 /windows10_drive_g/work/src/github.com/tangfeixiong/go-to-kubernetes/
```