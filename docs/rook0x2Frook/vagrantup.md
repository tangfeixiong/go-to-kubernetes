## DevOps

### Table of content

1. Debian 9 for mac
1. Fedora 27 for mac
1. Fedora 27 for windows 10

### Debian 9 for mac

file
```
fanhonglingdeMacBook-Pro:stretch64 fanhongling$ pwd
/Users/fanhongling/Downloads/https0x3A0x2F0x2Fwww.vagrantup.com/https0x3A0x2F0x2Fapp.vagrantup.com0x2Fboxes0x2Fsearch/debian/stretch64
```

```
fanhonglingdeMacBook-Pro:stretch64 fanhongling$ VBoxManage list runningvms
"go-to-kubernetes_default_1513507231694_54213" {be44b6c5-a7e9-4afc-ae0d-0e3a7acce03a}
"rook0x2Frook_default_1513669158492_88742" {ab852988-2130-4c39-a5ae-93e5028ec8bb}
```

up
```
fanhonglingdeMacBook-Pro:stretch64 fanhongling$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Box 'debian/stretch64' could not be found. Attempting to find and install...
    default: Box Provider: virtualbox
    default: Box Version: 9.3.0
==> default: Loading metadata for box 'debian/stretch64'
    default: URL: https://vagrantcloud.com/debian/stretch64
==> default: Adding box 'debian/stretch64' (v9.3.0) for provider: virtualbox
    default: Downloading: https://vagrantcloud.com/debian/boxes/stretch64/versions/9.3.0/providers/virtualbox.box
==> default: Successfully added box 'debian/stretch64' (v9.3.0) for 'virtualbox'!
==> default: Importing base box 'debian/stretch64'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'debian/stretch64' is up to date...
==> default: Setting the name of the VM: stretch64_default_1513689816764_73850
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

mount: unknown filesystem type 'vboxsf'

```

```
fanhonglingdeMacBook-Pro:stretch64 fanhongling$ vagrant ssh
Linux stretch 4.9.0-4-amd64 #1 SMP Debian 4.9.65-3 (2017-12-03) x86_64

The programs included with the Debian GNU/Linux system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Debian GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent
permitted by applicable law.
```

```
vagrant@stretch:~$ cat .ssh/authorized_keys 
ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEA6NF8iallvQVp22WDkTkyrtvp9eWW6A8YVr+kz4TjGYe7gHzIw+niNltGEFHzD8+v1I2YJ6oXevct1YeS0o9HZyN1Q9qgCgzUFtdOKLv6IedplqoPkcmF0aYet2PkEDo3MlTBckFXPITAMzF8dJSIFo9D8HfdOV0IAdx4O7PtixWKn5y2hMNG0zQPyUecp4pzC6kivAIhyfHilFR61RGL+GPXQ2MWZWFYbAGjyiYJnAmCP3NOTd0jMZEnDkbUvxhMmBYSdETk1rRgm+R4LOzFUGaHqHDLKLX+FIPKcF96hrucXzcWyLbIbEgE98OHlnVYCzRdK8jlqm8tehUc9c9WhQ== vagrant insecure public key
```

```
vagrant@stretch:~$ sudo apt update && sudo apt-cache search linux-headers
Ign:1 http://deb.debian.org/debian stretch InRelease                            
Get:2 http://security.debian.org/debian-security stretch/updates InRelease [63.0 kB]
Hit:3 http://deb.debian.org/debian stretch Release                                         
Get:5 http://security.debian.org/debian-security stretch/updates/main Sources [93.0 kB]         
Get:7 http://security.debian.org/debian-security stretch/updates/main amd64 Packages [249 kB]
Reading package lists... Done  
linux-headers-4.9.0-4-all - All header files for Linux 4.9 (meta-package)
linux-headers-4.9.0-4-all-amd64 - All header files for Linux 4.9 (meta-package)
linux-headers-4.9.0-4-amd64 - Header files for Linux 4.9.0-4-amd64
linux-headers-4.9.0-4-common - Common header files for Linux 4.9.0-4
linux-headers-4.9.0-4-common-rt - Common header files for Linux 4.9.0-4-rt
linux-headers-4.9.0-4-rt-amd64 - Header files for Linux 4.9.0-4-rt-amd64
linux-headers-amd64 - Header files for Linux amd64 configuration (meta-package)
linux-headers-rt-amd64 - Header files for Linux rt-amd64 configuration (meta-package)
linux-headers-4.9.0-3-all - All header files for Linux 4.9 (meta-package)
linux-headers-4.9.0-3-all-amd64 - All header files for Linux 4.9 (meta-package)
linux-headers-4.9.0-3-amd64 - Header files for Linux 4.9.0-3-amd64
linux-headers-4.9.0-3-common - Common header files for Linux 4.9.0-3
linux-headers-4.9.0-3-common-rt - Common header files for Linux 4.9.0-3-rt
linux-headers-4.9.0-3-rt-amd64 - Header files for Linux 4.9.0-3-rt-amd64
```

```
vagrant@stretch:~$ uname -r
4.9.0-4-amd64
```

```
vagrant@stretch:~$ sudo apt install -y linux-headers-`uname -r` gcc
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  binutils cpp cpp-6 gcc-6 libasan3 libatomic1 libc-dev-bin libc6-dev libcc1-0 libcilkrts5 libgcc-6-dev libgomp1 libisl15 libitm1 liblsan0
  libmpc3 libmpfr4 libmpx2 libquadmath0 libtsan0 libubsan0 linux-compiler-gcc-6-x86 linux-headers-4.9.0-4-common linux-kbuild-4.9
  linux-libc-dev manpages-dev
Suggested packages:
  binutils-doc cpp-doc gcc-6-locales gcc-multilib make autoconf automake libtool flex bison gdb gcc-doc gcc-6-multilib gcc-6-doc libgcc1-dbg
  libgomp1-dbg libitm1-dbg libatomic1-dbg libasan3-dbg liblsan0-dbg libtsan0-dbg libubsan0-dbg libcilkrts5-dbg libmpx2-dbg libquadmath0-dbg
  glibc-doc
The following NEW packages will be installed:
  binutils cpp cpp-6 gcc gcc-6 libasan3 libatomic1 libc-dev-bin libc6-dev libcc1-0 libcilkrts5 libgcc-6-dev libgomp1 libisl15 libitm1 liblsan0
  libmpc3 libmpfr4 libmpx2 libquadmath0 libtsan0 libubsan0 linux-compiler-gcc-6-x86 linux-headers-4.9.0-4-amd64 linux-headers-4.9.0-4-common
  linux-kbuild-4.9 linux-libc-dev manpages-dev
0 upgraded, 28 newly installed, 0 to remove and 2 not upgraded.
Need to get 37.1 MB of archives.
After this operation, 169 MB of additional disk space will be used.
Get:1 http://deb.debian.org/debian stretch/main amd64 binutils amd64 2.28-5 [3,770 kB]
Get:2 http://deb.debian.org/debian stretch/main amd64 libisl15 amd64 0.18-1 [564 kB]
Get:3 http://deb.debian.org/debian stretch/main amd64 libmpfr4 amd64 3.1.5-1 [556 kB]
Get:4 http://deb.debian.org/debian stretch/main amd64 libmpc3 amd64 1.0.3-1+b2 [39.9 kB]
Get:5 http://deb.debian.org/debian stretch/main amd64 cpp-6 amd64 6.3.0-18 [6,579 kB]
Get:6 http://deb.debian.org/debian stretch/main amd64 cpp amd64 4:6.3.0-4 [18.7 kB]                                                            
Get:7 http://deb.debian.org/debian stretch/main amd64 libcc1-0 amd64 6.3.0-18 [30.6 kB]                                                        
Get:8 http://deb.debian.org/debian stretch/main amd64 libgomp1 amd64 6.3.0-18 [73.2 kB]                                                        
Get:9 http://deb.debian.org/debian stretch/main amd64 libitm1 amd64 6.3.0-18 [27.4 kB]                                                         
Get:10 http://deb.debian.org/debian stretch/main amd64 libatomic1 amd64 6.3.0-18 [8,920 B]                                                     
Get:11 http://deb.debian.org/debian stretch/main amd64 libasan3 amd64 6.3.0-18 [311 kB]                                                        
Get:12 http://deb.debian.org/debian stretch/main amd64 liblsan0 amd64 6.3.0-18 [115 kB]                                                        
Get:13 http://deb.debian.org/debian stretch/main amd64 libtsan0 amd64 6.3.0-18 [257 kB]                                                        
Get:14 http://deb.debian.org/debian stretch/main amd64 libubsan0 amd64 6.3.0-18 [106 kB]                                                       
Get:15 http://deb.debian.org/debian stretch/main amd64 libcilkrts5 amd64 6.3.0-18 [40.4 kB]                                                    
Get:16 http://deb.debian.org/debian stretch/main amd64 libmpx2 amd64 6.3.0-18 [11.2 kB]                                                        
Get:17 http://deb.debian.org/debian stretch/main amd64 libquadmath0 amd64 6.3.0-18 [131 kB]                                                    
Get:18 http://deb.debian.org/debian stretch/main amd64 libgcc-6-dev amd64 6.3.0-18 [2,297 kB]                                                  
Get:19 http://deb.debian.org/debian stretch/main amd64 gcc-6 amd64 6.3.0-18 [6,895 kB]                                                         
Get:20 http://deb.debian.org/debian stretch/main amd64 gcc amd64 4:6.3.0-4 [5,196 B]                                                           
Get:21 http://deb.debian.org/debian stretch/main amd64 libc-dev-bin amd64 2.24-11+deb9u1 [259 kB]                                              
Get:22 http://deb.debian.org/debian stretch/main amd64 linux-libc-dev amd64 4.9.65-3 [1,298 kB]                                                
Get:23 http://deb.debian.org/debian stretch/main amd64 libc6-dev amd64 2.24-11+deb9u1 [2,365 kB]                                               
Get:24 http://deb.debian.org/debian stretch/main amd64 linux-compiler-gcc-6-x86 amd64 4.9.65-3 [507 kB]                                        
Get:25 http://deb.debian.org/debian stretch/main amd64 linux-headers-4.9.0-4-common all 4.9.65-3 [7,521 kB]                                    
Get:26 http://deb.debian.org/debian stretch/main amd64 linux-kbuild-4.9 amd64 4.9.65-3 [714 kB]                                                
Get:27 http://deb.debian.org/debian stretch/main amd64 linux-headers-4.9.0-4-amd64 amd64 4.9.65-3 [449 kB]                                     
Get:28 http://deb.debian.org/debian stretch/main amd64 manpages-dev all 4.10-2 [2,145 kB]                                                      
Fetched 37.1 MB in 21s (1,702 kB/s)                                                                                                            
Selecting previously unselected package binutils.
(Reading database ... 26760 files and directories currently installed.)
Preparing to unpack .../00-binutils_2.28-5_amd64.deb ...
Unpacking binutils (2.28-5) ...
Selecting previously unselected package libisl15:amd64.
Preparing to unpack .../01-libisl15_0.18-1_amd64.deb ...
Unpacking libisl15:amd64 (0.18-1) ...
Selecting previously unselected package libmpfr4:amd64.
Preparing to unpack .../02-libmpfr4_3.1.5-1_amd64.deb ...
Unpacking libmpfr4:amd64 (3.1.5-1) ...
Selecting previously unselected package libmpc3:amd64.
Preparing to unpack .../03-libmpc3_1.0.3-1+b2_amd64.deb ...
Unpacking libmpc3:amd64 (1.0.3-1+b2) ...
Selecting previously unselected package cpp-6.
Preparing to unpack .../04-cpp-6_6.3.0-18_amd64.deb ...
Unpacking cpp-6 (6.3.0-18) ...
Selecting previously unselected package cpp.
Preparing to unpack .../05-cpp_4%3a6.3.0-4_amd64.deb ...
Unpacking cpp (4:6.3.0-4) ...
Selecting previously unselected package libcc1-0:amd64.
Preparing to unpack .../06-libcc1-0_6.3.0-18_amd64.deb ...
Unpacking libcc1-0:amd64 (6.3.0-18) ...
Selecting previously unselected package libgomp1:amd64.
Preparing to unpack .../07-libgomp1_6.3.0-18_amd64.deb ...
Unpacking libgomp1:amd64 (6.3.0-18) ...
Selecting previously unselected package libitm1:amd64.
Preparing to unpack .../08-libitm1_6.3.0-18_amd64.deb ...
Unpacking libitm1:amd64 (6.3.0-18) ...
Selecting previously unselected package libatomic1:amd64.
Preparing to unpack .../09-libatomic1_6.3.0-18_amd64.deb ...
Unpacking libatomic1:amd64 (6.3.0-18) ...
Selecting previously unselected package libasan3:amd64.
Preparing to unpack .../10-libasan3_6.3.0-18_amd64.deb ...
Unpacking libasan3:amd64 (6.3.0-18) ...
Selecting previously unselected package liblsan0:amd64.
Preparing to unpack .../11-liblsan0_6.3.0-18_amd64.deb ...
Unpacking liblsan0:amd64 (6.3.0-18) ...
Selecting previously unselected package libtsan0:amd64.
Preparing to unpack .../12-libtsan0_6.3.0-18_amd64.deb ...
Unpacking libtsan0:amd64 (6.3.0-18) ...
Selecting previously unselected package libubsan0:amd64.
Preparing to unpack .../13-libubsan0_6.3.0-18_amd64.deb ...
Unpacking libubsan0:amd64 (6.3.0-18) ...
Selecting previously unselected package libcilkrts5:amd64.
Preparing to unpack .../14-libcilkrts5_6.3.0-18_amd64.deb ...
Unpacking libcilkrts5:amd64 (6.3.0-18) ...
Selecting previously unselected package libmpx2:amd64.
Preparing to unpack .../15-libmpx2_6.3.0-18_amd64.deb ...
Unpacking libmpx2:amd64 (6.3.0-18) ...
Selecting previously unselected package libquadmath0:amd64.
Preparing to unpack .../16-libquadmath0_6.3.0-18_amd64.deb ...
Unpacking libquadmath0:amd64 (6.3.0-18) ...
Selecting previously unselected package libgcc-6-dev:amd64.
Preparing to unpack .../17-libgcc-6-dev_6.3.0-18_amd64.deb ...
Unpacking libgcc-6-dev:amd64 (6.3.0-18) ...
Selecting previously unselected package gcc-6.
Preparing to unpack .../18-gcc-6_6.3.0-18_amd64.deb ...
Unpacking gcc-6 (6.3.0-18) ...
Selecting previously unselected package gcc.
Preparing to unpack .../19-gcc_4%3a6.3.0-4_amd64.deb ...
Unpacking gcc (4:6.3.0-4) ...
Selecting previously unselected package libc-dev-bin.
Preparing to unpack .../20-libc-dev-bin_2.24-11+deb9u1_amd64.deb ...
Unpacking libc-dev-bin (2.24-11+deb9u1) ...
Selecting previously unselected package linux-libc-dev:amd64.
Preparing to unpack .../21-linux-libc-dev_4.9.65-3_amd64.deb ...
Unpacking linux-libc-dev:amd64 (4.9.65-3) ...
Selecting previously unselected package libc6-dev:amd64.
Preparing to unpack .../22-libc6-dev_2.24-11+deb9u1_amd64.deb ...
Unpacking libc6-dev:amd64 (2.24-11+deb9u1) ...
Selecting previously unselected package linux-compiler-gcc-6-x86.
Preparing to unpack .../23-linux-compiler-gcc-6-x86_4.9.65-3_amd64.deb ...
Unpacking linux-compiler-gcc-6-x86 (4.9.65-3) ...
Selecting previously unselected package linux-headers-4.9.0-4-common.
Preparing to unpack .../24-linux-headers-4.9.0-4-common_4.9.65-3_all.deb ...
Unpacking linux-headers-4.9.0-4-common (4.9.65-3) ...
Selecting previously unselected package linux-kbuild-4.9.
Preparing to unpack .../25-linux-kbuild-4.9_4.9.65-3_amd64.deb ...
Unpacking linux-kbuild-4.9 (4.9.65-3) ...
Selecting previously unselected package linux-headers-4.9.0-4-amd64.
Preparing to unpack .../26-linux-headers-4.9.0-4-amd64_4.9.65-3_amd64.deb ...
Unpacking linux-headers-4.9.0-4-amd64 (4.9.65-3) ...
Selecting previously unselected package manpages-dev.
Preparing to unpack .../27-manpages-dev_4.10-2_all.deb ...
Unpacking manpages-dev (4.10-2) ...
Setting up libquadmath0:amd64 (6.3.0-18) ...
Setting up libgomp1:amd64 (6.3.0-18) ...
Setting up libatomic1:amd64 (6.3.0-18) ...
Setting up libcc1-0:amd64 (6.3.0-18) ...
Setting up libasan3:amd64 (6.3.0-18) ...
Setting up libcilkrts5:amd64 (6.3.0-18) ...
Setting up libubsan0:amd64 (6.3.0-18) ...
Setting up libtsan0:amd64 (6.3.0-18) ...
Setting up linux-libc-dev:amd64 (4.9.65-3) ...
Setting up liblsan0:amd64 (6.3.0-18) ...
Setting up libmpx2:amd64 (6.3.0-18) ...
Setting up libisl15:amd64 (0.18-1) ...
Setting up linux-kbuild-4.9 (4.9.65-3) ...
Processing triggers for libc-bin (2.24-11+deb9u1) ...
Setting up libmpfr4:amd64 (3.1.5-1) ...
Setting up linux-headers-4.9.0-4-common (4.9.65-3) ...
Processing triggers for man-db (2.7.6.1-2) ...
Setting up libmpc3:amd64 (1.0.3-1+b2) ...
Setting up binutils (2.28-5) ...
Setting up cpp-6 (6.3.0-18) ...
Setting up libc-dev-bin (2.24-11+deb9u1) ...
Setting up manpages-dev (4.10-2) ...
Setting up libc6-dev:amd64 (2.24-11+deb9u1) ...
Setting up libitm1:amd64 (6.3.0-18) ...
Setting up cpp (4:6.3.0-4) ...
Setting up libgcc-6-dev:amd64 (6.3.0-18) ...
Setting up gcc-6 (6.3.0-18) ...
Setting up gcc (4:6.3.0-4) ...
Setting up linux-compiler-gcc-6-x86 (4.9.65-3) ...
Setting up linux-headers-4.9.0-4-amd64 (4.9.65-3) ...
Processing triggers for libc-bin (2.24-11+deb9u1) ...
```

Exit, modify Vagrantfile to disable netowrk auto_config

mount VBoxGuestAdditions.iso and re-enter
```
vagrant@stretch:~$ sudo mount /dev/sr0 /media/
mount: /dev/sr0 is write-protected, mounting read-only
```

```
vagrant@stretch:~$ sudo /media/VBoxLinuxAdditions.run --nox11
Verifying archive integrity... All good.
Uncompressing VirtualBox 5.1.30 Guest Additions for Linux...........
VirtualBox Guest Additions installer
Copying additional installer modules ...
Installing additional modules ...
vboxadd.sh: Starting the VirtualBox Guest Additions.

Could not find the X.Org or XFree86 Window System, skipping.
```

```
fanhonglingdeMacBook-Pro:stretch64 fanhongling$ vagrant reload
==> default: Attempting graceful shutdown of VM...
==> default: Checking if box 'debian/stretch64' is up to date...
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
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /Users/fanhongling/go => /Users/fanhongling/go
    default: /Users/fanhongling/Downloads => /Users/fanhongling/Downloads

==> default: Machine 'default' has a post `vagrant up` message. This is a message
==> default: from the creator of the Vagrantfile, and not from Vagrant itself:
==> default: 
==> default: Vanilla Debian box. See https://app.vagrantup.com/debian for help and bug reports
```

```
fanhonglingdeMacBook-Pro:stretch64 fanhongling$ ssh -i ~/.ssh/vagrant vagrant@172.17.4.65
The authenticity of host '172.17.4.65 (172.17.4.65)' can't be established.
RSA key fingerprint is 70:0b:98:27:0c:75:24:19:65:78:de:27:ac:91:c4:12.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '172.17.4.65' (RSA) to the list of known hosts.
Linux stretch 4.9.0-4-amd64 #1 SMP Debian 4.9.65-3 (2017-12-03) x86_64

The programs included with the Debian GNU/Linux system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Debian GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent
permitted by applicable law.
Last login: Tue Dec 19 14:34:38 2017 from 10.0.2.2
```

```
vagrant@stretch:~$ ls /Users/fanhongling/go/
bin  pkg  src
```

### Mac

version
```
fanhonglingdeMacBook-Pro:rook0x2Frook fanhongling$ VBoxManage --version
5.1.30r118389
```

```
fanhonglingdeMacBook-Pro:rook0x2Frook fanhongling$ vagrant --version
Vagrant 2.0.0
```

up
```
fanhonglingdeMacBook-Pro:rook0x2Frook fanhongling$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Box 'fedora27-cloud-base' could not be found. Attempting to find and install...
    default: Box Provider: virtualbox
    default: Box Version: >= 0
==> default: Box file was not detected as metadata. Adding it directly...
==> default: Adding box 'fedora27-cloud-base' (v0) for provider: virtualbox
    default: Downloading: https://download.fedoraproject.org/pub/fedora/linux/releases/27/CloudImages/x86_64/images/Fedora-Cloud-Base-Vagrant-27-1.6.x86_64.vagrant-virtualbox.box
==> default: Successfully added box 'fedora27-cloud-base' (v0) for 'virtualbox'!
==> default: Importing base box 'fedora27-cloud-base'...
==> default: Matching MAC address for NAT networking...
==> default: Setting the name of the VM: rook0x2Frook_default_1513669158492_88742
==> default: Fixed port collision for 22 => 2222. Now on port 2201.
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
    default: Adapter 3: intnet
==> default: Forwarding ports...
    default: 22 (guest) => 2201 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2201
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
==> default: Configuring and enabling network interfaces...
    default: SSH address: 127.0.0.1:2201
    default: SSH username: vagrant
    default: SSH auth method: private key
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

mount: /Users/fanhongling/go: unknown filesystem type 'vboxsf'.

```

vboxsf by default is not installed
```
fanhonglingdeMacBook-Pro:rook0x2Frook fanhongling$ vagrant halt
==> default: Attempting graceful shutdown of VM...
```

Mount VBoxGuestAdditions.ISO like ![屏幕快照 2017-12-18 下午11.47.20.png](./屏幕快照%202017-12-18%20下午11.47.20.png)

Modify Vagrantfile to disable network auto_config 
```
fanhonglingdeMacBook-Pro:rook0x2Frook fanhongling$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Clearing any previously set forwarded ports...
==> default: Fixed port collision for 22 => 2222. Now on port 2201.
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
    default: Adapter 3: intnet
==> default: Forwarding ports...
    default: 22 (guest) => 2201 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2201
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Connection reset. Retrying...
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Connection reset. Retrying...
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

mount: /Users/fanhongling/go: unknown filesystem type 'vboxsf'.
```

Enter
```
fanhonglingdeMacBook-Pro:rook0x2Frook fanhongling$ ssh vagrant@172.17.4.63
The authenticity of host '172.17.4.63 (172.17.4.63)' can't be established.
RSA key fingerprint is 52:75:21:48:04:cf:94:18:5f:ec:2b:51:96:01:2a:de.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '172.17.4.63' (RSA) to the list of known hosts.
vagrant@172.17.4.63's password: 
Permission denied, please try again.
vagrant@172.17.4.63's password: 
Last failed login: Tue Dec 19 08:00:35 UTC 2017 from 172.17.4.1 on ssh:notty
There was 1 failed login attempt since the last successful login.
```

Change pasword
```
[vagrant@localhost ~]$ passwd    
Changing password for user vagrant.
Current password: 
New password: 
Retype new password: 
passwd: all authentication tokens updated successfully.
```

vboxsf
```
[vagrant@localhost ~]$ uname -r 
4.13.9-300.fc27.x86_64
```

```
[vagrant@localhost ~]$ sudo dnf search --showduplicates kernel-devel
Failed to set locale, defaulting to C
Fedora 27 - x86_64 - Updates                                                                                    528 kB/s |  13 MB     00:25    
Fedora 27 - x86_64                                                                                              984 kB/s |  58 MB     01:00    
Last metadata expiration check: 0:00:13 ago on Tue Dec 19 08:07:22 2017.
====================================================== Name Exactly Matched: kernel-devel ======================================================
kernel-devel-4.14.6-300.fc27.x86_64 : Development package for building kernel modules to match the kernel
kernel-devel-4.13.9-300.fc27.x86_64 : Development package for building kernel modules to match the kernel
```

```
[vagrant@localhost ~]$ sudo dnf install -y kernel-devel-4.13.9-300.fc27
Failed to set locale, defaulting to C
Last metadata expiration check: 0:01:25 ago on Tue Dec 19 08:07:22 2017.
Dependencies resolved.
================================================================================================================================================
 Package                                    Arch                       Version                                 Repository                  Size
================================================================================================================================================
Installing:
 kernel-devel                               x86_64                     4.13.9-300.fc27                         fedora                      11 M
Installing dependencies:
 perl-Carp                                  noarch                     1.42-394.fc27                           fedora                      28 k
 perl-Errno                                 x86_64                     1.28-401.fc27                           fedora                      71 k
 perl-Exporter                              noarch                     5.72-395.fc27                           fedora                      32 k
 perl-File-Path                             noarch                     2.15-1.fc27                             fedora                      37 k
 perl-IO                                    x86_64                     1.38-401.fc27                           fedora                     136 k
 perl-PathTools                             x86_64                     3.67-395.fc27                           fedora                      88 k
 perl-Scalar-List-Utils                     x86_64                     3:1.48-1.fc27                           fedora                      65 k
 perl-Socket                                x86_64                     4:2.024-5.fc27                          fedora                      56 k
 perl-Text-Tabs+Wrap                        noarch                     2013.0523-394.fc27                      fedora                      23 k
 perl-Unicode-Normalize                     x86_64                     1.25-395.fc27                           fedora                      80 k
 perl-constant                              noarch                     1.33-395.fc27                           fedora                      24 k
 perl-interpreter                           x86_64                     4:5.26.1-401.fc27                       fedora                     6.2 M
 perl-libs                                  x86_64                     4:5.26.1-401.fc27                       fedora                     1.5 M
 perl-macros                                x86_64                     4:5.26.1-401.fc27                       fedora                      67 k
 perl-parent                                noarch                     1:0.236-394.fc27                        fedora                      18 k
 perl-threads                               x86_64                     1:2.16-4.fc27                           fedora                      59 k
 perl-threads-shared                        x86_64                     1.57-4.fc27                             fedora                      46 k

Transaction Summary
================================================================================================================================================
Install  18 Packages

Total download size: 20 M
Installed size: 64 M
Downloading Packages:
(1/18): perl-Carp-1.42-394.fc27.noarch.rpm                                                                       55 kB/s |  28 kB     00:00    
(2/18): perl-Exporter-5.72-395.fc27.noarch.rpm                                                                  223 kB/s |  32 kB     00:00    
(3/18): perl-File-Path-2.15-1.fc27.noarch.rpm                                                                   245 kB/s |  37 kB     00:00    
(4/18): perl-IO-1.38-401.fc27.x86_64.rpm                                                                        308 kB/s | 136 kB     00:00    
(5/18): perl-PathTools-3.67-395.fc27.x86_64.rpm                                                                 266 kB/s |  88 kB     00:00    
(6/18): perl-Scalar-List-Utils-1.48-1.fc27.x86_64.rpm                                                           470 kB/s |  65 kB     00:00    
(7/18): perl-Text-Tabs+Wrap-2013.0523-394.fc27.noarch.rpm                                                       146 kB/s |  23 kB     00:00    
(8/18): perl-Unicode-Normalize-1.25-395.fc27.x86_64.rpm                                                         218 kB/s |  80 kB     00:00    
(9/18): perl-constant-1.33-395.fc27.noarch.rpm                                                                   15 kB/s |  24 kB     00:01    
(10/18): perl-libs-5.26.1-401.fc27.x86_64.rpm                                                                   380 kB/s | 1.5 MB     00:04    
(11/18): perl-macros-5.26.1-401.fc27.x86_64.rpm                                                                 312 kB/s |  67 kB     00:00    
(12/18): perl-parent-0.236-394.fc27.noarch.rpm                                                                  169 kB/s |  18 kB     00:00    
(13/18): perl-threads-2.16-4.fc27.x86_64.rpm                                                                    186 kB/s |  59 kB     00:00    
(14/18): perl-threads-shared-1.57-4.fc27.x86_64.rpm                                                             283 kB/s |  46 kB     00:00    
(15/18): perl-Errno-1.28-401.fc27.x86_64.rpm                                                                    211 kB/s |  71 kB     00:00    
(16/18): perl-Socket-2.024-5.fc27.x86_64.rpm                                                                    185 kB/s |  56 kB     00:00    
(17/18): perl-interpreter-5.26.1-401.fc27.x86_64.rpm                                                            620 kB/s | 6.2 MB     00:10    
(18/18): kernel-devel-4.13.9-300.fc27.x86_64.rpm                                                                715 kB/s |  11 MB     00:16    
------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                           1.1 MB/s |  20 MB     00:17     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                        1/1 
  Installing       : perl-Exporter-5.72-395.fc27.noarch                                                                                    1/18 
  Installing       : perl-libs-4:5.26.1-401.fc27.x86_64                                                                                    2/18 
  Running scriptlet: perl-libs-4:5.26.1-401.fc27.x86_64                                                                                    2/18 
  Installing       : perl-Carp-1.42-394.fc27.noarch                                                                                        3/18 
  Installing       : perl-Scalar-List-Utils-3:1.48-1.fc27.x86_64                                                                           4/18 
  Installing       : perl-macros-4:5.26.1-401.fc27.x86_64                                                                                  5/18 
  Installing       : perl-Text-Tabs+Wrap-2013.0523-394.fc27.noarch                                                                         6/18 
  Installing       : perl-Unicode-Normalize-1.25-395.fc27.x86_64                                                                           7/18 
  Installing       : perl-File-Path-2.15-1.fc27.noarch                                                                                     8/18 
  Installing       : perl-PathTools-3.67-395.fc27.x86_64                                                                                   9/18 
  Installing       : perl-constant-1.33-395.fc27.noarch                                                                                   10/18 
  Installing       : perl-parent-1:0.236-394.fc27.noarch                                                                                  11/18 
  Installing       : perl-threads-shared-1.57-4.fc27.x86_64                                                                               12/18 
  Installing       : perl-threads-1:2.16-4.fc27.x86_64                                                                                    13/18 
  Installing       : perl-Errno-1.28-401.fc27.x86_64                                                                                      14/18 
  Installing       : perl-IO-1.38-401.fc27.x86_64                                                                                         15/18 
  Installing       : perl-interpreter-4:5.26.1-401.fc27.x86_64                                                                            16/18 
  Installing       : perl-Socket-4:2.024-5.fc27.x86_64                                                                                    17/18 
  Installing       : kernel-devel-4.13.9-300.fc27.x86_64                                                                                  18/18 
  Running scriptlet: kernel-devel-4.13.9-300.fc27.x86_64                                                                                  18/18 
  Verifying        : kernel-devel-4.13.9-300.fc27.x86_64                                                                                   1/18 
  Verifying        : perl-interpreter-4:5.26.1-401.fc27.x86_64                                                                             2/18 
  Verifying        : perl-Carp-1.42-394.fc27.noarch                                                                                        3/18 
  Verifying        : perl-Exporter-5.72-395.fc27.noarch                                                                                    4/18 
  Verifying        : perl-File-Path-2.15-1.fc27.noarch                                                                                     5/18 
  Verifying        : perl-IO-1.38-401.fc27.x86_64                                                                                          6/18 
  Verifying        : perl-PathTools-3.67-395.fc27.x86_64                                                                                   7/18 
  Verifying        : perl-Scalar-List-Utils-3:1.48-1.fc27.x86_64                                                                           8/18 
  Verifying        : perl-Text-Tabs+Wrap-2013.0523-394.fc27.noarch                                                                         9/18 
  Verifying        : perl-Unicode-Normalize-1.25-395.fc27.x86_64                                                                          10/18 
  Verifying        : perl-constant-1.33-395.fc27.noarch                                                                                   11/18 
  Verifying        : perl-libs-4:5.26.1-401.fc27.x86_64                                                                                   12/18 
  Verifying        : perl-macros-4:5.26.1-401.fc27.x86_64                                                                                 13/18 
  Verifying        : perl-parent-1:0.236-394.fc27.noarch                                                                                  14/18 
  Verifying        : perl-threads-1:2.16-4.fc27.x86_64                                                                                    15/18 
  Verifying        : perl-threads-shared-1.57-4.fc27.x86_64                                                                               16/18 
  Verifying        : perl-Errno-1.28-401.fc27.x86_64                                                                                      17/18 
  Verifying        : perl-Socket-4:2.024-5.fc27.x86_64                                                                                    18/18 

Installed:
  kernel-devel.x86_64 4.13.9-300.fc27               perl-Carp.noarch 1.42-394.fc27                  perl-Errno.x86_64 1.28-401.fc27           
  perl-Exporter.noarch 5.72-395.fc27                perl-File-Path.noarch 2.15-1.fc27               perl-IO.x86_64 1.38-401.fc27              
  perl-PathTools.x86_64 3.67-395.fc27               perl-Scalar-List-Utils.x86_64 3:1.48-1.fc27     perl-Socket.x86_64 4:2.024-5.fc27         
  perl-Text-Tabs+Wrap.noarch 2013.0523-394.fc27     perl-Unicode-Normalize.x86_64 1.25-395.fc27     perl-constant.noarch 1.33-395.fc27        
  perl-interpreter.x86_64 4:5.26.1-401.fc27         perl-libs.x86_64 4:5.26.1-401.fc27              perl-macros.x86_64 4:5.26.1-401.fc27      
  perl-parent.noarch 1:0.236-394.fc27               perl-threads.x86_64 1:2.16-4.fc27               perl-threads-shared.x86_64 1.57-4.fc27    

Complete!
```

gcc
```
[vagrant@localhost ~]$ sudo dnf install -y gcc                         
Failed to set locale, defaulting to C
Last metadata expiration check: 0:02:58 ago on Tue Dec 19 08:07:22 2017.
Dependencies resolved.
================================================================================================================================================
 Package                              Arch                         Version                                  Repository                     Size
================================================================================================================================================
Installing:
 gcc                                  x86_64                       7.2.1-2.fc27                             fedora                         21 M
Installing dependencies:
 binutils                             x86_64                       2.29-6.fc27                              fedora                        5.9 M
 cpp                                  x86_64                       7.2.1-2.fc27                             fedora                        9.2 M
 glibc-devel                          x86_64                       2.26-15.fc27                             fedora                        985 k
 glibc-headers                        x86_64                       2.26-15.fc27                             fedora                        500 k
 isl                                  x86_64                       0.16.1-3.fc27                            fedora                        835 k
 kernel-headers                       x86_64                       4.14.6-300.fc27                          updates                       1.2 M
 libmpc                               x86_64                       1.0.2-8.fc27                             fedora                         56 k

Transaction Summary
================================================================================================================================================
Install  8 Packages

Total download size: 40 M
Installed size: 110 M
Downloading Packages:
(1/8): binutils-2.29-6.fc27.x86_64.rpm                                                                          320 kB/s | 5.9 MB     00:18    
(2/8): isl-0.16.1-3.fc27.x86_64.rpm                                                                             773 kB/s | 835 kB     00:01    
(3/8): libmpc-1.0.2-8.fc27.x86_64.rpm                                                                            69 kB/s |  56 kB     00:00    
(4/8): glibc-devel-2.26-15.fc27.x86_64.rpm                                                                      433 kB/s | 985 kB     00:02    
(5/8): glibc-headers-2.26-15.fc27.x86_64.rpm                                                                    652 kB/s | 500 kB     00:00    
(6/8): kernel-headers-4.14.6-300.fc27.x86_64.rpm                                                                182 kB/s | 1.2 MB     00:06    
(7/8): cpp-7.2.1-2.fc27.x86_64.rpm                                                                              282 kB/s | 9.2 MB     00:33    
(8/8): gcc-7.2.1-2.fc27.x86_64.rpm                                                                              499 kB/s |  21 MB     00:43    
------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                           885 kB/s |  40 MB     00:45     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                        1/1 
  Installing       : libmpc-1.0.2-8.fc27.x86_64                                                                                             1/8 
  Running scriptlet: libmpc-1.0.2-8.fc27.x86_64                                                                                             1/8 
  Installing       : cpp-7.2.1-2.fc27.x86_64                                                                                                2/8 
  Running scriptlet: cpp-7.2.1-2.fc27.x86_64                                                                                                2/8 
  Installing       : kernel-headers-4.14.6-300.fc27.x86_64                                                                                  3/8 
  Running scriptlet: glibc-headers-2.26-15.fc27.x86_64                                                                                      4/8 
  Installing       : glibc-headers-2.26-15.fc27.x86_64                                                                                      4/8 
  Installing       : glibc-devel-2.26-15.fc27.x86_64                                                                                        5/8 
  Running scriptlet: glibc-devel-2.26-15.fc27.x86_64                                                                                        5/8 
  Installing       : isl-0.16.1-3.fc27.x86_64                                                                                               6/8 
  Running scriptlet: isl-0.16.1-3.fc27.x86_64                                                                                               6/8 
  Installing       : binutils-2.29-6.fc27.x86_64                                                                                            7/8 
  Running scriptlet: binutils-2.29-6.fc27.x86_64                                                                                            7/8 
  Installing       : gcc-7.2.1-2.fc27.x86_64                                                                                                8/8 
  Running scriptlet: gcc-7.2.1-2.fc27.x86_64                                                                                                8/8 
  Verifying        : gcc-7.2.1-2.fc27.x86_64                                                                                                1/8 
  Verifying        : binutils-2.29-6.fc27.x86_64                                                                                            2/8 
  Verifying        : cpp-7.2.1-2.fc27.x86_64                                                                                                3/8 
  Verifying        : isl-0.16.1-3.fc27.x86_64                                                                                               4/8 
  Verifying        : libmpc-1.0.2-8.fc27.x86_64                                                                                             5/8 
  Verifying        : glibc-devel-2.26-15.fc27.x86_64                                                                                        6/8 
  Verifying        : glibc-headers-2.26-15.fc27.x86_64                                                                                      7/8 
  Verifying        : kernel-headers-4.14.6-300.fc27.x86_64                                                                                  8/8 

Installed:
  gcc.x86_64 7.2.1-2.fc27             binutils.x86_64 2.29-6.fc27   cpp.x86_64 7.2.1-2.fc27                 glibc-devel.x86_64 2.26-15.fc27  
  glibc-headers.x86_64 2.26-15.fc27   isl.x86_64 0.16.1-3.fc27      kernel-headers.x86_64 4.14.6-300.fc27   libmpc.x86_64 1.0.2-8.fc27       

Complete!
```

ISO
```
[vagrant@localhost ~]$ sudo mount /dev/sr0 /media
mount: /media: WARNING: device write-protected, mounted read-only.
```

Install
```
[vagrant@localhost ~]$ sudo /media/VBoxLinuxAdditions.run --nox11
Verifying archive integrity... All good.
Uncompressing VirtualBox 5.1.30 Guest Additions for Linux...........
VirtualBox Guest Additions installer
Copying additional installer modules ...
Installing additional modules ...
vboxadd.sh: Starting the VirtualBox Guest Additions.

Could not find the X.Org or XFree86 Window System, skipping.
```

Exit
```
[vagrant@localhost ~]$ exit
logout
Connection to 172.17.4.63 closed.
```

Reload
```
fanhonglingdeMacBook-Pro:rook0x2Frook fanhongling$ vagrant reload
==> default: Attempting graceful shutdown of VM...
==> default: Clearing any previously set forwarded ports...
==> default: Fixed port collision for 22 => 2222. Now on port 2201.
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
    default: Adapter 3: intnet
==> default: Forwarding ports...
    default: 22 (guest) => 2201 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2201
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Connection reset. Retrying...
    default: Warning: Remote connection disconnect. Retrying...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /Users/fanhongling/go => /Users/fanhongling/go
    default: /Users/fanhongling/Downloads => /Users/fanhongling/Downloads
```

work
```
fanhonglingdeMacBook-Pro:rook0x2Frook fanhongling$ ssh vagrant@172.17.4.63
vagrant@172.17.4.63's password: 
Last login: Tue Dec 19 08:00:38 2017 from 172.17.4.1
```

```
[vagrant@localhost ~]$ ls /Users/fanhongling/go/ 
bin  pkg  src
```

### Windows 10
Up
```
tangf@DESKTOP-H68OQDV /cygdrive/g/work/src/github.com/rook/rook
$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Box 'fedora27-cloud-base' could not be found. Attempting to find and install...
    default: Box Provider: virtualbox
    default: Box Version: >= 0
==> default: Box file was not detected as metadata. Adding it directly...
==> default: Adding box 'fedora27-cloud-base' (v0) for provider: virtualbox
    default: Downloading: https://download.fedoraproject.org/pub/fedora/linux/releases/27/CloudImages/x86_64/images/Fedora-Cloud-Base-Vagrant-27-1.6.x86_64.vagrant-virtualbox.box
    default:
==> default: Successfully added box 'fedora27-cloud-base' (v0) for 'virtualbox'!
==> default: Importing base box 'fedora27-cloud-base'...
==> default: Matching MAC address for NAT networking...
==> default: Setting the name of the VM: rook_default_1513236922901_62976
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
==> default: Configuring and enabling network interfaces...
```

Box
```
tangf@DESKTOP-H68OQDV /cygdrive/g/work/src/github.com/rook/rook
$ vagrant box list
centos/7                      (virtualbox, 1708.01)
centos/7                      (virtualbox, 1710.01)
centos7_inst                  (virtualbox, 0)
fedora/25-cloud-base          (virtualbox, 20161122)
fedora/26-atomic-host         (virtualbox, 26.20171030.0)
fedora27-cloud-base           (virtualbox, 0)
fedora_inst                   (virtualbox, 0)
openshift3_fedora_21-nov-2016 (virtualbox, 0)
opensuse_13-2                 (virtualbox, 0)
ubuntu/trusty64               (virtualbox, 20161111.0.0)
ubuntu/xenial64               (virtualbox, 20171122.0.0)
```

File
```
tangf@DESKTOP-H68OQDV /cygdrive/g/work/src/github.com/rook/rook
$ ls /cygdrive/c/Users/tangf/.vagrant.d/boxes/fedora27-cloud-base/0/virtualbox/
box.ovf                                       metadata.json
Fedora-Cloud-Base-Vagrant-27-1.6.x86_64.vmdk  Vagrantfile
```

vboxsf
```
[root@localhost media]# uname -r
4.13.9-300.fc27.x86_64
```

```
[root@localhost media]# dnf search --showduplicates kernel-devel
Last metadata expiration check: 0:08:33 ago on Thu 14 Dec 2017 08:28:03 AM UTC.
====================== Name Exactly Matched: kernel-devel ======================
kernel-devel-4.14.3-300.fc27.x86_64 : Development package for building kernel
                                    : modules to match the kernel
kernel-devel-4.14.3-300.fc27.x86_64 : Development package for building kernel
                                    : modules to match the kernel
kernel-devel-4.13.9-300.fc27.x86_64 : Development package for building kernel
                                    : modules to match the kernel
[root@localhost media]# dnf search --showduplicates kernel-headers
Last metadata expiration check: 0:12:11 ago on Thu 14 Dec 2017 08:28:03 AM UTC.
===================== Name Exactly Matched: kernel-headers =====================
kernel-headers-4.14.3-300.fc27.x86_64 : Header files for the Linux kernel for
                                      : use by glibc
kernel-headers-4.13.9-300.fc27.x86_64 : Header files for the Linux kernel for
                                      : use by glibc
```

```
[root@localhost media]# dnf install -y kernel-devel-4.13.9-300.fc27
Last metadata expiration check: 0:13:25 ago on Thu 14 Dec 2017 08:28:03 AM UTC.
Dependencies resolved.
================================================================================
 Package                    Arch       Version                 Repository  Size
================================================================================
Installing:
 kernel-devel               x86_64     4.13.9-300.fc27         fedora      11 M
Installing dependencies:
 perl-Carp                  noarch     1.42-394.fc27           fedora      28 k
 perl-Errno                 x86_64     1.28-401.fc27           fedora      71 k
 perl-Exporter              noarch     5.72-395.fc27           fedora      32 k
 perl-File-Path             noarch     2.15-1.fc27             fedora      37 k
 perl-IO                    x86_64     1.38-401.fc27           fedora     136 k
 perl-PathTools             x86_64     3.67-395.fc27           fedora      88 k
 perl-Scalar-List-Utils     x86_64     3:1.48-1.fc27           fedora      65 k
 perl-Socket                x86_64     4:2.024-5.fc27          fedora      56 k
 perl-Text-Tabs+Wrap        noarch     2013.0523-394.fc27      fedora      23 k
 perl-Unicode-Normalize     x86_64     1.25-395.fc27           fedora      80 k
 perl-constant              noarch     1.33-395.fc27           fedora      24 k
 perl-interpreter           x86_64     4:5.26.1-401.fc27       fedora     6.2 M
 perl-libs                  x86_64     4:5.26.1-401.fc27       fedora     1.5 M
 perl-macros                x86_64     4:5.26.1-401.fc27       fedora      67 k
 perl-parent                noarch     1:0.236-394.fc27        fedora      18 k
 perl-threads               x86_64     1:2.16-4.fc27           fedora      59 k
 perl-threads-shared        x86_64     1.57-4.fc27             fedora      46 k

Transaction Summary
================================================================================
Install  18 Packages

Total download size: 20 M
Installed size: 64 M
Downloading Packages:
(1/18): perl-Carp-1.42-394.fc27.noarch.rpm       61 kB/s |  28 kB     00:00
(2/18): perl-Exporter-5.72-395.fc27.noarch.rpm  301 kB/s |  32 kB     00:00
(3/18): perl-File-Path-2.15-1.fc27.noarch.rpm    87 kB/s |  37 kB     00:00
(4/18): perl-interpreter-5.26.1-401.fc27.x86_64 2.7 MB/s | 6.2 MB     00:02
(5/18): perl-IO-1.38-401.fc27.x86_64.rpm        104 kB/s | 136 kB     00:01
(6/18): perl-PathTools-3.67-395.fc27.x86_64.rpm 253 kB/s |  88 kB     00:00
(7/18): perl-Scalar-List-Utils-1.48-1.fc27.x86_  89 kB/s |  65 kB     00:00
(8/18): perl-Text-Tabs+Wrap-2013.0523-394.fc27.  32 kB/s |  23 kB     00:00
(9/18): perl-constant-1.33-395.fc27.noarch.rpm  263 kB/s |  24 kB     00:00
(10/18): perl-Unicode-Normalize-1.25-395.fc27.x 148 kB/s |  80 kB     00:00
(11/18): perl-libs-5.26.1-401.fc27.x86_64.rpm   3.3 MB/s | 1.5 MB     00:00
(12/18): perl-macros-5.26.1-401.fc27.x86_64.rpm 159 kB/s |  67 kB     00:00
(13/18): perl-parent-0.236-394.fc27.noarch.rpm  108 kB/s |  18 kB     00:00
(14/18): perl-threads-shared-1.57-4.fc27.x86_64 780 kB/s |  46 kB     00:00
(15/18): perl-threads-2.16-4.fc27.x86_64.rpm    376 kB/s |  59 kB     00:00
(16/18): perl-Errno-1.28-401.fc27.x86_64.rpm    798 kB/s |  71 kB     00:00
(17/18): perl-Socket-2.024-5.fc27.x86_64.rpm     56 kB/s |  56 kB     00:01
[MIRROR] kernel-devel-4.13.9-300.fc27.x86_64.rpm: Curl error (28): Timeout was reached for http://mirrors.ustc.edu.cn/fedora/releases/27/Everything/x86_64/os/Packages/k/kernel-devel-4.13.9-300.fc27.x86_64.rpm [Connection timed out after 30002 milliseconds]
(18/18): kernel-devel-4.13.9-300.fc27.x86_64.rp 359 kB/s |  11 MB     00:32
--------------------------------------------------------------------------------
Total                                           527 kB/s |  20 MB     00:38
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                        1/1
  Installing       : perl-Exporter-5.72-395.fc27.noarch                    1/18
  Installing       : perl-libs-4:5.26.1-401.fc27.x86_64                    2/18
  Running scriptlet: perl-libs-4:5.26.1-401.fc27.x86_64                    2/18
  Installing       : perl-Carp-1.42-394.fc27.noarch                        3/18
  Installing       : perl-Scalar-List-Utils-3:1.48-1.fc27.x86_64           4/18
  Installing       : perl-macros-4:5.26.1-401.fc27.x86_64                  5/18
  Installing       : perl-Text-Tabs+Wrap-2013.0523-394.fc27.noarch         6/18
  Installing       : perl-Unicode-Normalize-1.25-395.fc27.x86_64           7/18
  Installing       : perl-File-Path-2.15-1.fc27.noarch                     8/18
  Installing       : perl-PathTools-3.67-395.fc27.x86_64                   9/18
  Installing       : perl-constant-1.33-395.fc27.noarch                   10/18
  Installing       : perl-parent-1:0.236-394.fc27.noarch                  11/18
  Installing       : perl-threads-shared-1.57-4.fc27.x86_64               12/18
  Installing       : perl-threads-1:2.16-4.fc27.x86_64                    13/18
  Installing       : perl-Errno-1.28-401.fc27.x86_64                      14/18
  Installing       : perl-IO-1.38-401.fc27.x86_64                         15/18
  Installing       : perl-interpreter-4:5.26.1-401.fc27.x86_64            16/18
  Installing       : perl-Socket-4:2.024-5.fc27.x86_64                    17/18
  Installing       : kernel-devel-4.13.9-300.fc27.x86_64                  18/18
  Running scriptlet: kernel-devel-4.13.9-300.fc27.x86_64                  18/18
  Verifying        : kernel-devel-4.13.9-300.fc27.x86_64                   1/18
  Verifying        : perl-interpreter-4:5.26.1-401.fc27.x86_64             2/18
  Verifying        : perl-Carp-1.42-394.fc27.noarch                        3/18
  Verifying        : perl-Exporter-5.72-395.fc27.noarch                    4/18
  Verifying        : perl-File-Path-2.15-1.fc27.noarch                     5/18
  Verifying        : perl-IO-1.38-401.fc27.x86_64                          6/18
  Verifying        : perl-PathTools-3.67-395.fc27.x86_64                   7/18
  Verifying        : perl-Scalar-List-Utils-3:1.48-1.fc27.x86_64           8/18
  Verifying        : perl-Text-Tabs+Wrap-2013.0523-394.fc27.noarch         9/18
  Verifying        : perl-Unicode-Normalize-1.25-395.fc27.x86_64          10/18
  Verifying        : perl-constant-1.33-395.fc27.noarch                   11/18
  Verifying        : perl-libs-4:5.26.1-401.fc27.x86_64                   12/18
  Verifying        : perl-macros-4:5.26.1-401.fc27.x86_64                 13/18
  Verifying        : perl-parent-1:0.236-394.fc27.noarch                  14/18
  Verifying        : perl-threads-1:2.16-4.fc27.x86_64                    15/18
  Verifying        : perl-threads-shared-1.57-4.fc27.x86_64               16/18
  Verifying        : perl-Errno-1.28-401.fc27.x86_64                      17/18
  Verifying        : perl-Socket-4:2.024-5.fc27.x86_64                    18/18

Installed:
  kernel-devel.x86_64 4.13.9-300.fc27
  perl-Carp.noarch 1.42-394.fc27
  perl-Errno.x86_64 1.28-401.fc27
  perl-Exporter.noarch 5.72-395.fc27
  perl-File-Path.noarch 2.15-1.fc27
  perl-IO.x86_64 1.38-401.fc27
  perl-PathTools.x86_64 3.67-395.fc27
  perl-Scalar-List-Utils.x86_64 3:1.48-1.fc27
  perl-Socket.x86_64 4:2.024-5.fc27
  perl-Text-Tabs+Wrap.noarch 2013.0523-394.fc27
  perl-Unicode-Normalize.x86_64 1.25-395.fc27
  perl-constant.noarch 1.33-395.fc27
  perl-interpreter.x86_64 4:5.26.1-401.fc27
  perl-libs.x86_64 4:5.26.1-401.fc27
  perl-macros.x86_64 4:5.26.1-401.fc27
  perl-parent.noarch 1:0.236-394.fc27
  perl-threads.x86_64 1:2.16-4.fc27
  perl-threads-shared.x86_64 1.57-4.fc27

Complete!
```

```
[root@localhost media]# ls /usr/src/kernels/
4.13.9-300.fc27.x86_64
```

```
[root@localhost media]# sudo dnf install -y gcc
Last metadata expiration check: 0:17:02 ago on Thu 14 Dec 2017 08:28:03 AM UTC.
Dependencies resolved.
================================================================================
 Package              Arch         Version                  Repository     Size
================================================================================
Installing:
 gcc                  x86_64       7.2.1-2.fc27             fedora         21 M
Installing dependencies:
 binutils             x86_64       2.29-6.fc27              fedora        5.9 M
 cpp                  x86_64       7.2.1-2.fc27             fedora        9.2 M
 glibc-devel          x86_64       2.26-15.fc27             fedora        985 k
 glibc-headers        x86_64       2.26-15.fc27             fedora        500 k
 isl                  x86_64       0.16.1-3.fc27            fedora        835 k
 kernel-headers       x86_64       4.14.3-300.fc27          updates       1.2 M
 libmpc               x86_64       1.0.2-8.fc27             fedora         56 k

Transaction Summary
================================================================================
Install  8 Packages

Total download size: 40 M
Installed size: 110 M
Downloading Packages:
(1/8): binutils-2.29-6.fc27.x86_64.rpm          1.7 MB/s | 5.9 MB     00:03
(2/8): isl-0.16.1-3.fc27.x86_64.rpm             1.6 MB/s | 835 kB     00:00
(3/8): libmpc-1.0.2-8.fc27.x86_64.rpm           781 kB/s |  56 kB     00:00
(4/8): glibc-devel-2.26-15.fc27.x86_64.rpm      1.2 MB/s | 985 kB     00:00
(5/8): glibc-headers-2.26-15.fc27.x86_64.rpm    1.0 MB/s | 500 kB     00:00
(6/8): cpp-7.2.1-2.fc27.x86_64.rpm              1.6 MB/s | 9.2 MB     00:05
(7/8): kernel-headers-4.14.3-300.fc27.x86_64.rp 882 kB/s | 1.2 MB     00:01
(8/8): gcc-7.2.1-2.fc27.x86_64.rpm              2.2 MB/s |  21 MB     00:09
--------------------------------------------------------------------------------
Total                                           2.3 MB/s |  40 MB     00:17
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                        1/1
  Installing       : libmpc-1.0.2-8.fc27.x86_64                             1/8
  Running scriptlet: libmpc-1.0.2-8.fc27.x86_64                             1/8
  Installing       : cpp-7.2.1-2.fc27.x86_64                                2/8
  Running scriptlet: cpp-7.2.1-2.fc27.x86_64                                2/8
  Installing       : kernel-headers-4.14.3-300.fc27.x86_64                  3/8
  Running scriptlet: glibc-headers-2.26-15.fc27.x86_64                      4/8
  Installing       : glibc-headers-2.26-15.fc27.x86_64                      4/8
  Installing       : glibc-devel-2.26-15.fc27.x86_64                        5/8
  Running scriptlet: glibc-devel-2.26-15.fc27.x86_64                        5/8
  Installing       : isl-0.16.1-3.fc27.x86_64                               6/8
  Running scriptlet: isl-0.16.1-3.fc27.x86_64                               6/8
  Installing       : binutils-2.29-6.fc27.x86_64                            7/8
  Running scriptlet: binutils-2.29-6.fc27.x86_64                            7/8
  Installing       : gcc-7.2.1-2.fc27.x86_64                                8/8
  Running scriptlet: gcc-7.2.1-2.fc27.x86_64                                8/8
  Verifying        : gcc-7.2.1-2.fc27.x86_64                                1/8
  Verifying        : binutils-2.29-6.fc27.x86_64                            2/8
  Verifying        : cpp-7.2.1-2.fc27.x86_64                                3/8
  Verifying        : isl-0.16.1-3.fc27.x86_64                               4/8
  Verifying        : libmpc-1.0.2-8.fc27.x86_64                             5/8
  Verifying        : glibc-devel-2.26-15.fc27.x86_64                        6/8
  Verifying        : glibc-headers-2.26-15.fc27.x86_64                      7/8
  Verifying        : kernel-headers-4.14.3-300.fc27.x86_64                  8/8

Installed:
  gcc.x86_64 7.2.1-2.fc27                   binutils.x86_64 2.29-6.fc27
  cpp.x86_64 7.2.1-2.fc27                   glibc-devel.x86_64 2.26-15.fc27
  glibc-headers.x86_64 2.26-15.fc27         isl.x86_64 0.16.1-3.fc27
  kernel-headers.x86_64 4.14.3-300.fc27     libmpc.x86_64 1.0.2-8.fc27

Complete!
```

```
[root@localhost media]# ./VBoxLinuxAdditions.run --nox11
Verifying archive integrity... All good.
Uncompressing VirtualBox 5.1.30 Guest Additions for Linux...........
VirtualBox Guest Additions installer
Removing installed version 5.1.30 of VirtualBox Guest Additions...
Copying additional installer modules ...
Installing additional modules ...
vboxadd.sh: Starting the VirtualBox Guest Additions.

Could not find the X.Org or XFree86 Window System, skipping.
```

```
tangf@DESKTOP-H68OQDV /cygdrive/g/work/src/github.com/rook/rook
$ vagrant reload
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
    default: Warning: Remote connection disconnect. Retrying...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Configuring and enabling network interfaces...
==> default: Mounting shared folders...
    default: /windows10_drive_f => F:/
    default: /windows10_drive_g => G:/
==> default: Machine already provisioned. Run `vagrant provision` or use the `--provision`
==> default: flag to force provisioning. Provisioners marked to run always will still run.

tangf@DESKTOP-H68OQDV /cygdrive/g/work/src/github.com/rook/rook
```

```
$ vagrant ssh
Last login: Thu Dec 14 08:20:01 2017 from 10.0.2.2
```

```
[vagrant@localhost ~]$ ls /windows10_drive_g/work
bin  pkg  src
```

### Ubuntu 17.04 for mac

file
```
fanhonglingdeMacBook-Pro:zesty fanhongling$ pwd
/Users/fanhongling/Downloads/https0x3A0x2F0x2Fwww.vagrantup.com/https0x3A0x2F0x2Fcloud-images.ubuntu.com/zesty
```

up
```
fanhonglingdeMacBook-Pro:zesty fanhongling$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Box 'ubuntu17.04-cloud' could not be found. Attempting to find and install...
    default: Box Provider: virtualbox
    default: Box Version: >= 0
==> default: Box file was not detected as metadata. Adding it directly...
==> default: Adding box 'ubuntu17.04-cloud' (v0) for provider: virtualbox
    default: Downloading: https://cloud-images.ubuntu.com/zesty/20171219/zesty-server-cloudimg-amd64-vagrant.box
==> default: Successfully added box 'ubuntu17.04-cloud' (v0) for 'virtualbox'!
==> default: Importing base box 'ubuntu17.04-cloud'...
==> default: Matching MAC address for NAT networking...
==> default: Setting the name of the VM: zesty_default_1513686186345_35988
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
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Mounting shared folders...
    default: /Users/fanhongling/go => /Users/fanhongling/go
    default: /Users/fanhongling/Downloads => /Users/fanhongling/Downloads
```

```
fanhonglingdeMacBook-Pro:origin fanhongling$ cat ~/.vagrant.d/boxes/ubuntu17.04-cloud/0/virtualbox/Vagrantfile 
# Front load the includes
include_vagrantfile = File.expand_path("../include/_Vagrantfile", __FILE__)
load include_vagrantfile if File.exist?(include_vagrantfile)

Vagrant.configure("2") do |config|
  config.vm.base_mac = "02F2D5CBEDDC"
  config.ssh.username = "ubuntu"
  config.ssh.password = "a011032d2d8ed7993c979e48"

  config.vm.provider "virtualbox" do |vb|
     vb.customize [ "modifyvm", :id, "--uart1", "0x3F8", "4" ]
     vb.customize [ "modifyvm", :id, "--uartmode1", "file", File.join(Dir.pwd, "ubuntu-zesty-17.04-cloudimg-console.log") ]
  end
end
```

```
fanhonglingdeMacBook-Pro:zesty fanhongling$ vagrant ssh
==> default: The machine you're attempting to SSH into is configured to use
==> default: password-based authentication. Vagrant can't script entering the
==> default: password for you. If you're prompted for a password, please enter
==> default: the same password you have configured in the Vagrantfile.
ubuntu@127.0.0.1's password: 
Welcome to Ubuntu 17.04 (GNU/Linux 4.10.0-42-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage


  Get cloud support with Ubuntu Advantage Cloud Guest:
    http://www.ubuntu.com/business/services/cloud

0 packages can be updated.
0 updates are security updates.

New release '17.10' available.
Run 'do-release-upgrade' to upgrade to it.


Last login: Tue Dec 19 12:37:49 2017 from 10.0.2.2
```

```
fanhonglingdeMacBook-Pro:origin fanhongling$ cat ~/.ssh/vagrant.pub 
ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEA6NF8iallvQVp22WDkTkyrtvp9eWW6A8YVr+kz4TjGYe7gHzIw+niNltGEFHzD8+v1I2YJ6oXevct1YeS0o9HZyN1Q9qgCgzUFtdOKLv6IedplqoPkcmF0aYet2PkEDo3MlTBckFXPITAMzF8dJSIFo9D8HfdOV0IAdx4O7PtixWKn5y2hMNG0zQPyUecp4pzC6kivAIhyfHilFR61RGL+GPXQ2MWZWFYbAGjyiYJnAmCP3NOTd0jMZEnDkbUvxhMmBYSdETk1rRgm+R4LOzFUGaHqHDLKLX+FIPKcF96hrucXzcWyLbIbEgE98OHlnVYCzRdK8jlqm8tehUc9c9WhQ== vagrant insecure public key
```

```
ubuntu@ubuntu-zesty:~$ echo "ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEA6NF8iallvQVp22WDkTkyrtvp9eWW6A8YVr+kz4TjGYe7gHzIw+niNltGEFHzD8+v1I2YJ6oXevct1YeS0o9HZyN1Q9qgCgzUFtdOKLv6IedplqoPkcmF0aYet2PkEDo3MlTBckFXPITAMzF8dJSIFo9D8HfdOV0IAdx4O7PtixWKn5y2hMNG0zQPyUecp4pzC6kivAIhyfHilFR61RGL+GPXQ2MWZWFYbAGjyiYJnAmCP3NOTd0jMZEnDkbUvxhMmBYSdETk1rRgm+R4LOzFUGaHqHDLKLX+FIPKcF96hrucXzcWyLbIbEgE98OHlnVYCzRdK8jlqm8tehUc9c9WhQ== vagrant insecure public key" >> .ssh/authorized_keys 
```

```
ubuntu@ubuntu-zesty:~$ exit
logout
Connection to 127.0.0.1 closed.
```

```
fanhonglingdeMacBook-Pro:zesty fanhongling$ vagrant ssh -- -i ~/.ssh/vagrant
==> default: The machine you're attempting to SSH into is configured to use
==> default: password-based authentication. Vagrant can't script entering the
==> default: password for you. If you're prompted for a password, please enter
==> default: the same password you have configured in the Vagrantfile.
Welcome to Ubuntu 17.04 (GNU/Linux 4.10.0-42-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage


  Get cloud support with Ubuntu Advantage Cloud Guest:
    http://www.ubuntu.com/business/services/cloud

0 packages can be updated.
0 updates are security updates.

New release '17.10' available.
Run 'do-release-upgrade' to upgrade to it.


Last login: Tue Dec 19 12:41:27 2017 from 10.0.2.2
```