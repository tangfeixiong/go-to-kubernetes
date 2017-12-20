# DevOps

## Docker
```
vagrant@stretch:~$ sudo apt install -y docker-engine
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following packages were automatically installed and are no longer required:
  libglib2.0-0 libglib2.0-data shared-mime-info xdg-user-dirs
Use 'sudo apt autoremove' to remove them.
The following additional packages will be installed:
  apparmor aufs-dkms aufs-tools cgroupfs-mount dkms fakeroot git git-man libapparmor-perl liberror-perl libfakeroot libltdl7
  linux-headers-amd64 make patch rsync
Suggested packages:
  apparmor-profiles apparmor-profiles-extra apparmor-utils aufs-dev python3-apport menu git-daemon-run | git-daemon-sysvinit git-doc git-el
  git-email git-gui gitk gitweb git-arch git-cvs git-mediawiki git-svn make-doc ed diffutils-doc
The following NEW packages will be installed:
  apparmor aufs-dkms aufs-tools cgroupfs-mount dkms docker-engine fakeroot git git-man libapparmor-perl liberror-perl libfakeroot libltdl7
  linux-headers-amd64 make patch rsync
0 upgraded, 17 newly installed, 0 to remove and 2 not upgraded.
Need to get 22.4 MB of archives.
After this operation, 111 MB of additional disk space will be used.
Get:1 http://deb.debian.org/debian stretch/main amd64 liberror-perl all 0.17024-1 [26.9 kB]
Get:2 http://deb.debian.org/debian stretch/main amd64 git-man all 1:2.11.0-3+deb9u2 [1,432 kB]
Get:3 http://deb.debian.org/debian stretch/main amd64 git amd64 1:2.11.0-3+deb9u2 [4,160 kB]
Get:4 http://security.debian.org/debian-security stretch/updates/main amd64 rsync amd64 3.1.2-1+deb9u1 [393 kB]
Get:6 http://deb.debian.org/debian stretch/main amd64 make amd64 4.1-9.1 [302 kB]
Get:7 http://deb.debian.org/debian stretch/main amd64 patch amd64 2.7.5-1+b2 [112 kB]
Get:8 http://deb.debian.org/debian stretch/main amd64 dkms all 2.3-2 [74.8 kB]
Get:9 http://deb.debian.org/debian stretch/main amd64 aufs-dkms amd64 4.9+20161219-1 [169 kB]
Get:10 http://deb.debian.org/debian stretch/main amd64 aufs-tools amd64 1:4.1+20161219-1 [102 kB]
Get:11 http://deb.debian.org/debian stretch/main amd64 cgroupfs-mount all 1.3 [5,716 B]
Get:12 http://deb.debian.org/debian stretch/main amd64 libltdl7 amd64 2.4.6-2 [389 kB]
Get:13 http://deb.debian.org/debian stretch/main amd64 libfakeroot amd64 1.21-3.1 [45.7 kB]
Get:14 http://deb.debian.org/debian stretch/main amd64 fakeroot amd64 1.21-3.1 [85.6 kB]
Get:15 http://deb.debian.org/debian stretch/main amd64 linux-headers-amd64 amd64 4.9+80+deb9u2 [5,880 B]
Get:16 http://deb.debian.org/debian stretch/main amd64 libapparmor-perl amd64 2.11.0-3 [82.1 kB]
Get:5 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 docker-engine amd64 1.11.2-0~xenial [14.5 MB]                         
Get:17 http://deb.debian.org/debian stretch/main amd64 apparmor amd64 2.11.0-3 [525 kB]                                                        
Fetched 22.4 MB in 15s (1,433 kB/s)                                                                                                            
Preconfiguring packages ...
Selecting previously unselected package liberror-perl.
(Reading database ... 46969 files and directories currently installed.)
Preparing to unpack .../00-liberror-perl_0.17024-1_all.deb ...
Unpacking liberror-perl (0.17024-1) ...
Selecting previously unselected package git-man.
Preparing to unpack .../01-git-man_1%3a2.11.0-3+deb9u2_all.deb ...
Unpacking git-man (1:2.11.0-3+deb9u2) ...
Selecting previously unselected package git.
Preparing to unpack .../02-git_1%3a2.11.0-3+deb9u2_amd64.deb ...
Unpacking git (1:2.11.0-3+deb9u2) ...
Selecting previously unselected package make.
Preparing to unpack .../03-make_4.1-9.1_amd64.deb ...
Unpacking make (4.1-9.1) ...
Selecting previously unselected package patch.
Preparing to unpack .../04-patch_2.7.5-1+b2_amd64.deb ...
Unpacking patch (2.7.5-1+b2) ...
Selecting previously unselected package dkms.
Preparing to unpack .../05-dkms_2.3-2_all.deb ...
Unpacking dkms (2.3-2) ...
Selecting previously unselected package aufs-dkms.
Preparing to unpack .../06-aufs-dkms_4.9+20161219-1_amd64.deb ...
Unpacking aufs-dkms (4.9+20161219-1) ...
Selecting previously unselected package aufs-tools.
Preparing to unpack .../07-aufs-tools_1%3a4.1+20161219-1_amd64.deb ...
Unpacking aufs-tools (1:4.1+20161219-1) ...
Selecting previously unselected package cgroupfs-mount.
Preparing to unpack .../08-cgroupfs-mount_1.3_all.deb ...
Unpacking cgroupfs-mount (1.3) ...
Selecting previously unselected package libltdl7:amd64.
Preparing to unpack .../09-libltdl7_2.4.6-2_amd64.deb ...
Unpacking libltdl7:amd64 (2.4.6-2) ...
Selecting previously unselected package docker-engine.
Preparing to unpack .../10-docker-engine_1.11.2-0~xenial_amd64.deb ...
Unpacking docker-engine (1.11.2-0~xenial) ...
Selecting previously unselected package libfakeroot:amd64.
Preparing to unpack .../11-libfakeroot_1.21-3.1_amd64.deb ...
Unpacking libfakeroot:amd64 (1.21-3.1) ...
Selecting previously unselected package fakeroot.
Preparing to unpack .../12-fakeroot_1.21-3.1_amd64.deb ...
Unpacking fakeroot (1.21-3.1) ...
Selecting previously unselected package linux-headers-amd64.
Preparing to unpack .../13-linux-headers-amd64_4.9+80+deb9u2_amd64.deb ...
Unpacking linux-headers-amd64 (4.9+80+deb9u2) ...
Selecting previously unselected package rsync.
Preparing to unpack .../14-rsync_3.1.2-1+deb9u1_amd64.deb ...
Unpacking rsync (3.1.2-1+deb9u1) ...
Selecting previously unselected package libapparmor-perl.
Preparing to unpack .../15-libapparmor-perl_2.11.0-3_amd64.deb ...
Unpacking libapparmor-perl (2.11.0-3) ...
Selecting previously unselected package apparmor.
Preparing to unpack .../16-apparmor_2.11.0-3_amd64.deb ...
Unpacking apparmor (2.11.0-3) ...
Setting up aufs-tools (1:4.1+20161219-1) ...
Setting up linux-headers-amd64 (4.9+80+deb9u2) ...
Setting up git-man (1:2.11.0-3+deb9u2) ...
Setting up make (4.1-9.1) ...
Setting up liberror-perl (0.17024-1) ...
Setting up cgroupfs-mount (1.3) ...
Setting up rsync (3.1.2-1+deb9u1) ...
Created symlink /etc/systemd/system/multi-user.target.wants/rsync.service → /lib/systemd/system/rsync.service.
Setting up patch (2.7.5-1+b2) ...
Setting up dkms (2.3-2) ...
Processing triggers for libc-bin (2.24-11+deb9u1) ...
Setting up libfakeroot:amd64 (1.21-3.1) ...
Processing triggers for systemd (232-25+deb9u1) ...
Setting up libltdl7:amd64 (2.4.6-2) ...
Processing triggers for man-db (2.7.6.1-2) ...
Setting up libapparmor-perl (2.11.0-3) ...
Setting up git (1:2.11.0-3+deb9u2) ...
Setting up docker-engine (1.11.2-0~xenial) ...
Created symlink /etc/systemd/system/multi-user.target.wants/docker.service → /lib/systemd/system/docker.service.
Created symlink /etc/systemd/system/sockets.target.wants/docker.socket → /lib/systemd/system/docker.socket.
Setting up aufs-dkms (4.9+20161219-1) ...
Loading new aufs-4.9+20161219 DKMS files...
Building for 4.9.0-4-amd64
Building initial module for 4.9.0-4-amd64
Done.

aufs:
Running module version sanity check.
 - Original module
   - No original module exists within this kernel
 - Installation
   - Installing to /lib/modules/4.9.0-4-amd64/updates/dkms/

depmod....

DKMS: install completed.
Setting up apparmor (2.11.0-3) ...
Created symlink /etc/systemd/system/sysinit.target.wants/apparmor.service → /lib/systemd/system/apparmor.service.
update-rc.d: warning: start and stop actions are no longer supported; falling back to defaults
diff: /var/lib/apparmor/profiles/.apparmor.md5sums: No such file or directory
Setting up fakeroot (1.21-3.1) ...
update-alternatives: using /usr/bin/fakeroot-sysv to provide /usr/bin/fakeroot (fakeroot) in auto mode
Processing triggers for libc-bin (2.24-11+deb9u1) ...
Processing triggers for systemd (232-25+deb9u1) ...
```

```
vagrant@stretch:~$ sudo apt remove -y docker-engine
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following packages were automatically installed and are no longer required:
  apparmor aufs-dkms aufs-tools cgroupfs-mount dkms fakeroot git git-man libapparmor-perl liberror-perl libfakeroot libglib2.0-0
  libglib2.0-data libltdl7 linux-headers-amd64 patch rsync shared-mime-info xdg-user-dirs
Use 'sudo apt autoremove' to remove them.
The following packages will be REMOVED:
  docker-engine
0 upgraded, 0 newly installed, 1 to remove and 2 not upgraded.
After this operation, 73.4 MB disk space will be freed.
(Reading database ... 48478 files and directories currently installed.)
Removing docker-engine (1.11.2-0~xenial) ...
Processing triggers for man-db (2.7.6.1-2) ...
```

```
vagrant@stretch:~$ sudo apt-get install -y software-properties-common
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following packages were automatically installed and are no longer required:
  apparmor aufs-dkms aufs-tools cgroupfs-mount dkms fakeroot git git-man libapparmor-perl liberror-perl libfakeroot libltdl7
  linux-headers-amd64 patch rsync
Use 'sudo apt autoremove' to remove them.
The following additional packages will be installed:
  gir1.2-glib-2.0 gir1.2-packagekitglib-1.0 libcap2-bin libdbus-glib-1-2 libgirepository-1.0-1 libglib2.0-bin libgstreamer1.0-0
  libpackagekit-glib2-18 libpam-cap libpolkit-agent-1-0 libpolkit-backend-1-0 libpolkit-gobject-1-0 packagekit packagekit-tools policykit-1
  python3-dbus python3-gi python3-software-properties unattended-upgrades
Suggested packages:
  gstreamer1.0-tools appstream python-dbus-doc python3-dbus-dbg bsd-mailx mail-transport-agent needrestart
The following NEW packages will be installed:
  gir1.2-glib-2.0 gir1.2-packagekitglib-1.0 libcap2-bin libdbus-glib-1-2 libgirepository-1.0-1 libglib2.0-bin libgstreamer1.0-0
  libpackagekit-glib2-18 libpam-cap libpolkit-agent-1-0 libpolkit-backend-1-0 libpolkit-gobject-1-0 packagekit packagekit-tools policykit-1
  python3-dbus python3-gi python3-software-properties software-properties-common unattended-upgrades
0 upgraded, 20 newly installed, 0 to remove and 2 not upgraded.
Need to get 5,819 kB of archives.
After this operation, 14.2 MB of additional disk space will be used.
Get:1 http://deb.debian.org/debian stretch/main amd64 libgirepository-1.0-1 amd64 1.50.0-1+b1 [89.0 kB]
Get:2 http://deb.debian.org/debian stretch/main amd64 gir1.2-glib-2.0 amd64 1.50.0-1+b1 [139 kB]
Get:3 http://deb.debian.org/debian stretch/main amd64 libpackagekit-glib2-18 amd64 1.1.5-2 [114 kB]
Get:4 http://deb.debian.org/debian stretch/main amd64 gir1.2-packagekitglib-1.0 amd64 1.1.5-2 [34.6 kB]
Get:5 http://deb.debian.org/debian stretch/main amd64 libcap2-bin amd64 1:2.25-1 [26.5 kB]
Get:6 http://deb.debian.org/debian stretch/main amd64 libdbus-glib-1-2 amd64 0.108-2 [206 kB]
Get:7 http://deb.debian.org/debian stretch/main amd64 libglib2.0-bin amd64 2.50.3-2 [1,615 kB]
Get:8 http://deb.debian.org/debian stretch/main amd64 libgstreamer1.0-0 amd64 1.10.4-1 [1,962 kB]
Get:9 http://deb.debian.org/debian stretch/main amd64 libpam-cap amd64 1:2.25-1 [13.5 kB]
Get:10 http://deb.debian.org/debian stretch/main amd64 libpolkit-gobject-1-0 amd64 0.105-18 [43.6 kB]
Get:11 http://deb.debian.org/debian stretch/main amd64 libpolkit-agent-1-0 amd64 0.105-18 [24.3 kB]
Get:12 http://deb.debian.org/debian stretch/main amd64 libpolkit-backend-1-0 amd64 0.105-18 [45.8 kB]
Get:13 http://deb.debian.org/debian stretch/main amd64 policykit-1 amd64 0.105-18 [63.4 kB]
Get:14 http://deb.debian.org/debian stretch/main amd64 packagekit amd64 1.1.5-2 [546 kB]
Get:15 http://deb.debian.org/debian stretch/main amd64 packagekit-tools amd64 1.1.5-2 [45.2 kB]
Get:16 http://deb.debian.org/debian stretch/main amd64 python3-dbus amd64 1.2.4-1+b1 [184 kB]
Get:17 http://deb.debian.org/debian stretch/main amd64 python3-gi amd64 3.22.0-2 [473 kB]
Get:18 http://deb.debian.org/debian stretch/main amd64 python3-software-properties all 0.96.20.2-1 [49.5 kB]
Get:19 http://deb.debian.org/debian stretch/main amd64 software-properties-common all 0.96.20.2-1 [83.6 kB]
Get:20 http://deb.debian.org/debian stretch/main amd64 unattended-upgrades all 0.93.1+nmu1 [61.7 kB]
Fetched 5,819 kB in 3s (1,564 kB/s)        
Preconfiguring packages ...
Selecting previously unselected package libgirepository-1.0-1:amd64.
(Reading database ... 48379 files and directories currently installed.)
Preparing to unpack .../00-libgirepository-1.0-1_1.50.0-1+b1_amd64.deb ...
Unpacking libgirepository-1.0-1:amd64 (1.50.0-1+b1) ...
Selecting previously unselected package gir1.2-glib-2.0:amd64.
Preparing to unpack .../01-gir1.2-glib-2.0_1.50.0-1+b1_amd64.deb ...
Unpacking gir1.2-glib-2.0:amd64 (1.50.0-1+b1) ...
Selecting previously unselected package libpackagekit-glib2-18:amd64.
Preparing to unpack .../02-libpackagekit-glib2-18_1.1.5-2_amd64.deb ...
Unpacking libpackagekit-glib2-18:amd64 (1.1.5-2) ...
Selecting previously unselected package gir1.2-packagekitglib-1.0.
Preparing to unpack .../03-gir1.2-packagekitglib-1.0_1.1.5-2_amd64.deb ...
Unpacking gir1.2-packagekitglib-1.0 (1.1.5-2) ...
Selecting previously unselected package libcap2-bin.
Preparing to unpack .../04-libcap2-bin_1%3a2.25-1_amd64.deb ...
Unpacking libcap2-bin (1:2.25-1) ...
Selecting previously unselected package libdbus-glib-1-2:amd64.
Preparing to unpack .../05-libdbus-glib-1-2_0.108-2_amd64.deb ...
Unpacking libdbus-glib-1-2:amd64 (0.108-2) ...
Selecting previously unselected package libglib2.0-bin.
Preparing to unpack .../06-libglib2.0-bin_2.50.3-2_amd64.deb ...
Unpacking libglib2.0-bin (2.50.3-2) ...
Selecting previously unselected package libgstreamer1.0-0:amd64.
Preparing to unpack .../07-libgstreamer1.0-0_1.10.4-1_amd64.deb ...
Unpacking libgstreamer1.0-0:amd64 (1.10.4-1) ...
Selecting previously unselected package libpam-cap:amd64.
Preparing to unpack .../08-libpam-cap_1%3a2.25-1_amd64.deb ...
Unpacking libpam-cap:amd64 (1:2.25-1) ...
Selecting previously unselected package libpolkit-gobject-1-0:amd64.
Preparing to unpack .../09-libpolkit-gobject-1-0_0.105-18_amd64.deb ...
Unpacking libpolkit-gobject-1-0:amd64 (0.105-18) ...
Selecting previously unselected package libpolkit-agent-1-0:amd64.
Preparing to unpack .../10-libpolkit-agent-1-0_0.105-18_amd64.deb ...
Unpacking libpolkit-agent-1-0:amd64 (0.105-18) ...
Selecting previously unselected package libpolkit-backend-1-0:amd64.
Preparing to unpack .../11-libpolkit-backend-1-0_0.105-18_amd64.deb ...
Unpacking libpolkit-backend-1-0:amd64 (0.105-18) ...
Selecting previously unselected package policykit-1.
Preparing to unpack .../12-policykit-1_0.105-18_amd64.deb ...
Unit polkit.service does not exist, proceeding anyway.
Created symlink /run/systemd/system/polkit.service → /dev/null.
Unpacking policykit-1 (0.105-18) ...
Selecting previously unselected package packagekit.
Preparing to unpack .../13-packagekit_1.1.5-2_amd64.deb ...
Unpacking packagekit (1.1.5-2) ...
Selecting previously unselected package packagekit-tools.
Preparing to unpack .../14-packagekit-tools_1.1.5-2_amd64.deb ...
Unpacking packagekit-tools (1.1.5-2) ...
Selecting previously unselected package python3-dbus.
Preparing to unpack .../15-python3-dbus_1.2.4-1+b1_amd64.deb ...
Unpacking python3-dbus (1.2.4-1+b1) ...
Selecting previously unselected package python3-gi.
Preparing to unpack .../16-python3-gi_3.22.0-2_amd64.deb ...
Unpacking python3-gi (3.22.0-2) ...
Selecting previously unselected package python3-software-properties.
Preparing to unpack .../17-python3-software-properties_0.96.20.2-1_all.deb ...
Unpacking python3-software-properties (0.96.20.2-1) ...
Selecting previously unselected package software-properties-common.
Preparing to unpack .../18-software-properties-common_0.96.20.2-1_all.deb ...
Unpacking software-properties-common (0.96.20.2-1) ...
Selecting previously unselected package unattended-upgrades.
Preparing to unpack .../19-unattended-upgrades_0.93.1+nmu1_all.deb ...
Unpacking unattended-upgrades (0.93.1+nmu1) ...
Setting up libpam-cap:amd64 (1:2.25-1) ...
Setting up libcap2-bin (1:2.25-1) ...
Setting up libdbus-glib-1-2:amd64 (0.108-2) ...
Setting up libpackagekit-glib2-18:amd64 (1.1.5-2) ...
Setting up libgirepository-1.0-1:amd64 (1.50.0-1+b1) ...
Setting up gir1.2-glib-2.0:amd64 (1.50.0-1+b1) ...
Processing triggers for libc-bin (2.24-11+deb9u1) ...
Processing triggers for systemd (232-25+deb9u1) ...
Setting up unattended-upgrades (0.93.1+nmu1) ...

Creating config file /etc/apt/apt.conf.d/20auto-upgrades with new version

Creating config file /etc/apt/apt.conf.d/50unattended-upgrades with new version
Created symlink /etc/systemd/system/multi-user.target.wants/unattended-upgrades.service → /lib/systemd/system/unattended-upgrades.service.
Synchronizing state of unattended-upgrades.service with SysV service script with /lib/systemd/systemd-sysv-install.
Executing: /lib/systemd/systemd-sysv-install enable unattended-upgrades
Processing triggers for man-db (2.7.6.1-2) ...
Setting up gir1.2-packagekitglib-1.0 (1.1.5-2) ...
Setting up python3-software-properties (0.96.20.2-1) ...
Processing triggers for dbus (1.10.24-0+deb9u1) ...
Setting up libglib2.0-bin (2.50.3-2) ...
Setting up libpolkit-gobject-1-0:amd64 (0.105-18) ...
Setting up python3-dbus (1.2.4-1+b1) ...
Setting up libgstreamer1.0-0:amd64 (1.10.4-1) ...
Setcap worked! gst-ptp-helper is not suid!
Setting up libpolkit-agent-1-0:amd64 (0.105-18) ...
Setting up python3-gi (3.22.0-2) ...
Setting up libpolkit-backend-1-0:amd64 (0.105-18) ...
Setting up software-properties-common (0.96.20.2-1) ...
Setting up policykit-1 (0.105-18) ...
Removed /run/systemd/system/polkit.service.
polkit.service is a disabled or a static unit not running, not starting it.
Setting up packagekit (1.1.5-2) ...
Setting up packagekit-tools (1.1.5-2) ...
Processing triggers for systemd (232-25+deb9u1) ...
Processing triggers for libc-bin (2.24-11+deb9u1) ...
Processing triggers for dbus (1.10.24-0+deb9u1) ...
```

```
vagrant@stretch:~$ sudo add-apt-repository "deb  https://apt.dockerproject.org/repo/ debian-$(lsb_release -cs) main"
```

```
vagrant@stretch:~$ wget https://apt.dockerproject.org/gpg -O- | sudo apt-key add -
--2017-12-19 14:12:07--  https://apt.dockerproject.org/gpg
Resolving apt.dockerproject.org (apt.dockerproject.org)... 13.33.14.239
Connecting to apt.dockerproject.org (apt.dockerproject.org)|13.33.14.239|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 1648 (1.6K) [application/pgp-keys]
Saving to: ‘STDOUT’

-                                   100%[===================================================================>]   1.61K  --.-KB/s    in 0s      

2017-12-19 14:12:08 (11.4 MB/s) - written to stdout [1648/1648]

OK
```

```
vagrant@stretch:~$ sudo apt-get update && sudo apt-cache show docker-engine=1.13.1-0~debian-stretch
Hit:1 http://security.debian.org/debian-security stretch/updates InRelease
Ign:2 http://deb.debian.org/debian stretch InRelease                                    
Hit:3 http://deb.debian.org/debian stretch Release                                      
Hit:5 https://apt.dockerproject.org/repo debian-stretch InRelease
Hit:6 https://packages.cloud.google.com/apt kubernetes-xenial InRelease
Reading package lists... Done
Package: docker-engine
Priority: optional
Section: admin
Installed-Size: 102933
Maintainer: Docker <support@docker.com>
Architecture: amd64
Version: 1.13.1-0~debian-stretch
Depends: iptables, init-system-helpers (>= 1.18~), libapparmor1 (>= 2.6~devel), libc6 (>= 2.17), libdevmapper1.02.1 (>= 2:1.02.97), libltdl7 (>= 2.4.6), libseccomp2 (>= 2.1.0), libsystemd0
Recommends: aufs-tools, ca-certificates, cgroupfs-mount | cgroup-lite, git, xz-utils
Conflicts: docker (<< 1.5~), docker-engine-cs, docker.io, lxc-docker, lxc-docker-virtual-package
Filename: pool/main/d/docker-engine/docker-engine_1.13.1-0~debian-stretch_amd64.deb
Size: 21280658
MD5sum: 9915bb2927073a109e3b505607b257de
SHA1: 19296514610aa2e5efddade5222cafae7894a689
SHA256: 3c9bc151978011249e714abae0f13034692807072233d24e2776d4b95ac598d7
SHA512: 48b75a03171ef53be28226303565afdbf5c6b07f3899249bb8545c4b879d9fb8cd281dc8d61f6ecc5cf308cb0d7e7867ff40ec138959f152783d57ee6dbdd633
Description: Docker: the open-source application container engine
 Docker is an open source project to build, ship and run any application as a
 lightweight container
 .
 Docker containers are both hardware-agnostic and platform-agnostic. This means
 they can run anywhere, from your laptop to the largest EC2 compute instance and
 everything in between - and they don't require you to use a particular
 language, framework or packaging system. That makes them great building blocks
 for deploying and scaling web apps, databases, and backend services without
 depending on a particular stack or provider.
Description-md5: 081115dab45344eb40852d38ab20e419
Homepage: https://dockerproject.org
```

```
vagrant@stretch:~$ sudo apt-get install -y docker-engine=1.13.1-0~debian-stretch
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following packages were automatically installed and are no longer required:
  apparmor libapparmor-perl
Use 'sudo apt autoremove' to remove them.
The following NEW packages will be installed:
  docker-engine
0 upgraded, 1 newly installed, 0 to remove and 2 not upgraded.
Need to get 21.3 MB of archives.
After this operation, 105 MB of additional disk space will be used.
Get:1 https://apt.dockerproject.org/repo debian-stretch/main amd64 docker-engine amd64 1.13.1-0~debian-stretch [21.3 MB]
Fetched 21.3 MB in 12s (1,677 kB/s)                                                                                                            
Selecting previously unselected package docker-engine.
(Reading database ... 48926 files and directories currently installed.)
Preparing to unpack .../docker-engine_1.13.1-0~debian-stretch_amd64.deb ...
Unpacking docker-engine (1.13.1-0~debian-stretch) ...
Setting up docker-engine (1.13.1-0~debian-stretch) ...
Installing new version of config file /etc/default/docker ...
Installing new version of config file /etc/init.d/docker ...
Installing new version of config file /etc/init/docker.conf ...
Processing triggers for systemd (232-25+deb9u1) ...
Processing triggers for man-db (2.7.6.1-2) ...
```

```
vagrant@stretch:~$ sudo usermod -aG docker vagrant
```

```
vagrant@stretch:~$ exit
logout
Connection to 127.0.0.1 closed.
```

```
fanhonglingdeMacBook-Pro:stretch64 fanhongling$ ssh -i ~/.ssh/vagrant vagrant@172.17.4.65
Linux stretch 4.9.0-4-amd64 #1 SMP Debian 4.9.65-3 (2017-12-03) x86_64

The programs included with the Debian GNU/Linux system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Debian GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent
permitted by applicable law.
Last login: Tue Dec 19 14:34:38 2017 from 10.0.2.2
```

```
vagrant@stretch:~$ docker info
Containers: 0
 Running: 0
 Paused: 0
 Stopped: 0
Images: 0
Server Version: 1.13.1
Storage Driver: devicemapper
 Pool Name: docker-8:1-399420-pool
 Pool Blocksize: 65.54 kB
 Base Device Size: 10.74 GB
 Backing Filesystem: ext4
 Data file: /dev/loop0
 Metadata file: /dev/loop1
 Data Space Used: 240.1 MB
 Data Space Total: 107.4 GB
 Data Space Available: 7.355 GB
 Metadata Space Used: 696.3 kB
 Metadata Space Total: 2.147 GB
 Metadata Space Available: 2.147 GB
 Thin Pool Minimum Free Space: 10.74 GB
 Udev Sync Supported: true
 Deferred Removal Enabled: false
 Deferred Deletion Enabled: false
 Deferred Deleted Device Count: 0
 Data loop file: /var/lib/docker/devicemapper/devicemapper/data
 WARNING: Usage of loopback devices is strongly discouraged for production use. Use `--storage-opt dm.thinpooldev` to specify a custom block storage device.
 Metadata loop file: /var/lib/docker/devicemapper/devicemapper/metadata
 Library Version: 1.02.137 (2016-11-30)
Logging Driver: json-file
Cgroup Driver: cgroupfs
Plugins: 
 Volume: local
 Network: bridge host macvlan null overlay
Swarm: inactive
Runtimes: runc
Default Runtime: runc
Init Binary: docker-init
containerd version: aa8187dbd3b7ad67d8e5e3a15115d3eef43a7ed1
runc version: 9df8b306d01f59d3a8029be411de015b7304dd8f
init version: 949e6fa
Security Options:
 seccomp
  Profile: default
Kernel Version: 4.9.0-4-amd64
Operating System: Debian GNU/Linux 9 (stretch)
OSType: linux
Architecture: x86_64
CPUs: 1
Total Memory: 1.958 GiB
Name: stretch
ID: 3YIZ:II3Q:LPXL:R5UR:JYPY:PHFE:RVNB:OFSI:UR22:ZM5K:CILM:CTEE
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

### Hostname, hosts, security

host
```
vagrant@stretch:~$ hostname -I
10.0.2.15 172.17.4.65 172.28.128.5 172.18.0.1 
```

```
vagrant@stretch:~$ addr=$(hostname -I | awk '{print $2}'); echo $addr
172.17.4.65
```

```
vagrant@stretch:~$ name="kubedev-${addr//\./-}"; echo $name
kubedev-172-17-4-65
```

```
vagrant@stretch:~$ sudo hostnamectl set-hostname $name
```

```
vagrant@stretch:~$ hostname
kubedev-172-17-4-65
```

```
vagrant@stretch:~$ echo "$addr    $name" | sudo tee -a /etc/hosts
sudo: unable to resolve host kubedev-172-17-4-65
172.17.4.65    kubedev-172-17-4-65
```

```
vagrant@stretch:~$ sudo systemctl is-active firewall.service
inactive
```

## kubernetes

repo
```
vagrant@stretch:~$ echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
```

```
vagrant@stretch:~$ sudo apt update && sudo apt-cache show kubeadm
Ign:1 http://deb.debian.org/debian stretch InRelease                            
Get:2 http://security.debian.org/debian-security stretch/updates InRelease [63.0 kB]
Hit:3 http://deb.debian.org/debian stretch Release                                         
Get:5 http://security.debian.org/debian-security stretch/updates/main Sources [93.0 kB]         
Get:7 http://security.debian.org/debian-security stretch/updates/main amd64 Packages [249 kB]
Reading package lists... Done  
E: The method driver /usr/lib/apt/methods/https could not be found.
N: Is the package apt-transport-https installed?
```

```
vagrant@stretch:~$ sudo apt install -yq apt-transport-https
Reading package lists...
Building dependency tree...
Reading state information...
The following NEW packages will be installed:
  apt-transport-https
0 upgraded, 1 newly installed, 0 to remove and 1 not upgraded.
Need to get 171 kB of archives.
After this operation, 243 kB of additional disk space will be used.
Get:1 http://deb.debian.org/debian stretch/main amd64 apt-transport-https amd64 1.4.8 [171 kB]
Fetched 171 kB in 2s (75.4 kB/s)
Selecting previously unselected package apt-transport-https.
(Reading database ... 26643 files and directories currently installed.)
Preparing to unpack .../apt-transport-https_1.4.8_amd64.deb ...
Unpacking apt-transport-https (1.4.8) ...
Setting up apt-transport-https (1.4.8) ...
```

```
vagrant@stretch:~$ sudo apt update && sudo apt-cache show kubeadm=1.9.0-00
Ign:1 http://deb.debian.org/debian stretch InRelease
Hit:2 http://deb.debian.org/debian stretch Release    
Hit:3 http://security.debian.org/debian-security stretch/updates InRelease
Get:5 https://packages.cloud.google.com/apt kubernetes-stretch InRelease [6,301 B]
Err:5 https://packages.cloud.google.com/apt kubernetes-stretch InRelease
  The following signatures couldn't be verified because the public key is not available: NO_PUBKEY 3746C208A7317B0F
Reading package lists... Done
W: GPG error: https://packages.cloud.google.com/apt kubernetes-stretch InRelease: The following signatures couldn't be verified because the public key is not available: NO_PUBKEY 3746C208A7317B0F
E: The repository 'https://apt.kubernetes.io kubernetes-stretch InRelease' is not signed.
N: Updating from such a repository can't be done securely, and is therefore disabled by default.
N: See apt-secure(8) manpage for repository creation and user configuration details.
```

```
vagrant@stretch:~$ wget https://packages.cloud.google.com/apt/doc/apt-key.gpg -O- | sudo apt-key add -
--2017-12-19 13:42:40--  https://packages.cloud.google.com/apt/doc/apt-key.gpg
Resolving packages.cloud.google.com (packages.cloud.google.com)... 216.58.197.206, 2404:6800:4004:80e::200e
Connecting to packages.cloud.google.com (packages.cloud.google.com)|216.58.197.206|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 663 [application/octet-stream]
Saving to: ‘STDOUT’

-                                   100%[===================================================================>]     663  --.-KB/s    in 0s      

2017-12-19 13:42:41 (18.7 MB/s) - written to stdout [663/663]

OK
```

```
vagrant@stretch:~$ sudo apt update && sudo apt-cache show kubeadm=1.9.0-00
Ign:1 http://deb.debian.org/debian stretch InRelease
Hit:2 http://deb.debian.org/debian stretch Release
Hit:4 http://security.debian.org/debian-security stretch/updates InRelease
Get:5 https://packages.cloud.google.com/apt kubernetes-xenial InRelease [8,951 B]
Get:6 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 Packages [12.1 kB]
Fetched 21.0 kB in 2s (8,169 B/s)
Reading package lists... Done
Building dependency tree       
Reading state information... Done
2 packages can be upgraded. Run 'apt list --upgradable' to see them.
Package: kubeadm
Version: 1.9.0-00
Installed-Size: 147193
Maintainer: Kubernetes Authors <kubernetes-dev+release@googlegroups.com>
Architecture: amd64
Depends: kubelet (>= 1.6.0), kubectl (>= 1.6.0)
Description: Kubernetes Cluster Bootstrapping Tool
 The Kubernetes command line tool for bootstrapping a Kubernetes cluster.
Description-md5: bb3c7836839894793de38af875e01b30
Homepage: https://kubernetes.io
Filename: pool/kubeadm_1.9.0-00_amd64_191bd1d8a63d263cdb318f77b03fad6c8387a912cf16a21b9b24d7e9108b4911.deb
Priority: optional
SHA256: 191bd1d8a63d263cdb318f77b03fad6c8387a912cf16a21b9b24d7e9108b4911
Section: misc
Size: 20044720
```

```
vagrant@stretch:~$ sudo apt -y install kubeadm
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following additional packages will be installed:
  ebtables ethtool kubelet kubernetes-cni socat
The following NEW packages will be installed:
  ebtables ethtool kubeadm kubelet kubernetes-cni socat
0 upgraded, 6 newly installed, 0 to remove and 2 not upgraded.
Need to get 47.0 MB of archives.
After this operation, 346 MB of additional disk space will be used.
Get:1 http://deb.debian.org/debian stretch/main amd64 ebtables amd64 2.0.10.4-3.5+b1 [85.5 kB]
Get:2 http://deb.debian.org/debian stretch/main amd64 ethtool amd64 1:4.8-1+b1 [113 kB]
Get:3 http://deb.debian.org/debian stretch/main amd64 socat amd64 1.7.3.1-2+deb9u1 [353 kB]
Get:4 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubernetes-cni amd64 0.6.0-00 [5,910 kB]
Get:5 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubelet amd64 1.9.0-00 [20.5 MB]
Get:6 https://packages.cloud.google.com/apt kubernetes-xenial/main amd64 kubeadm amd64 1.9.0-00 [20.0 MB]                                      
Fetched 47.0 MB in 36s (1,279 kB/s)                                                                                                            
Selecting previously unselected package ebtables.
(Reading database ... 26650 files and directories currently installed.)
Preparing to unpack .../0-ebtables_2.0.10.4-3.5+b1_amd64.deb ...
Unpacking ebtables (2.0.10.4-3.5+b1) ...
Selecting previously unselected package ethtool.
Preparing to unpack .../1-ethtool_1%3a4.8-1+b1_amd64.deb ...
Unpacking ethtool (1:4.8-1+b1) ...
Selecting previously unselected package kubernetes-cni.
Preparing to unpack .../2-kubernetes-cni_0.6.0-00_amd64.deb ...
Unpacking kubernetes-cni (0.6.0-00) ...
Selecting previously unselected package socat.
Preparing to unpack .../3-socat_1.7.3.1-2+deb9u1_amd64.deb ...
Unpacking socat (1.7.3.1-2+deb9u1) ...
Selecting previously unselected package kubelet.
Preparing to unpack .../4-kubelet_1.9.0-00_amd64.deb ...
Unpacking kubelet (1.9.0-00) ...
Selecting previously unselected package kubeadm.
Preparing to unpack .../5-kubeadm_1.9.0-00_amd64.deb ...
Unpacking kubeadm (1.9.0-00) ...
Setting up kubernetes-cni (0.6.0-00) ...
Setting up socat (1.7.3.1-2+deb9u1) ...
Processing triggers for systemd (232-25+deb9u1) ...
Setting up ebtables (2.0.10.4-3.5+b1) ...
Created symlink /etc/systemd/system/multi-user.target.wants/ebtables.service → /lib/systemd/system/ebtables.service.
update-rc.d: warning: start and stop actions are no longer supported; falling back to defaults
Processing triggers for man-db (2.7.6.1-2) ...
Setting up ethtool (1:4.8-1+b1) ...
Setting up kubelet (1.9.0-00) ...
Created symlink /etc/systemd/system/multi-user.target.wants/kubelet.service → /lib/systemd/system/kubelet.service.
Setting up kubeadm (1.9.0-00) ...
Processing triggers for systemd (232-25+deb9u1) ...
```

### install

admin.conf
```
vagrant@stretch:~$ mkdir .kube; scp vagrant@172.17.4.59:/home/vagrant/.kube/config .kube
The authenticity of host '172.17.4.59 (172.17.4.59)' can't be established.
ECDSA key fingerprint is SHA256:VV533NeJHFykYUa3CBNDm0zDI4HSnEaMKcQSKwtvLGw.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '172.17.4.59' (ECDSA) to the list of known hosts.
vagrant@172.17.4.59's password: 
config                                                                                                        100% 5451     4.0MB/s   00:00    
```

```
vagrant@stretch:~$ kubectl version
Client Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.0", GitCommit:"925c127ec6b946659ad0fd596fa959be43f0cc05", GitTreeState:"clean", BuildDate:"2017-12-15T21:07:38Z", GoVersion:"go1.9.2", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"9", GitVersion:"v1.9.0", GitCommit:"925c127ec6b946659ad0fd596fa959be43f0cc05", GitTreeState:"clean", BuildDate:"2017-12-15T20:55:30Z", GoVersion:"go1.9.2", Compiler:"gc", Platform:"linux/amd64"}
```

```
vagrant@stretch:~$ sudo systemctl enable kubelet.service && sudo systemctl start kubelet.service
```

```
vagrant@stretch:~$ sudo kubeadm token list --kubeconfig=.kube/config
TOKEN                     TTL       EXPIRES                USAGES                   DESCRIPTION   EXTRA GROUPS
86257a.d063ca0eedd9e165   19h       2017-12-20T10:40:36Z   authentication,signing   <none>        system:bootstrappers:kubeadm:default-node-token
```




```
vagrant@stretch:~$ sudo sed -i 's/\(ExecStart=.*ARGS\)/\1 --fail-swap-on=false/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf 
```

```
vagrant@stretch:~$ sudo kubeadm join --token=86257a.d063ca0eedd9e165 172.17.4.59 --discovery-token-unsafe-skip-ca-verification
[preflight] Running pre-flight checks.
	[WARNING FileExisting-crictl]: crictl not found in system path
[preflight] Some fatal errors occurred:
	[ERROR Swap]: running with swap on is not supported. Please disable swap
[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`
```

```
vagrant@stretch:~$ sudo cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf 
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_SYSTEM_PODS_ARGS=--pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true"
Environment="KUBELET_NETWORK_ARGS=--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
Environment="KUBELET_DNS_ARGS=--cluster-dns=10.96.0.10 --cluster-domain=cluster.local"
Environment="KUBELET_AUTHZ_ARGS=--authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt"
Environment="KUBELET_CADVISOR_ARGS=--cadvisor-port=0"
Environment="KUBELET_CERTIFICATE_ARGS=--rotate-certificates=true --cert-dir=/var/lib/kubelet/pki"
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS --fail-swap-on=false
```

```
vagrant@stretch:~$ sudo systemctl daemon-reload && sudo systemctl restart kubelet.service
```

```
vagrant@stretch:~$ sudo kubeadm join --token=86257a.d063ca0eedd9e165 172.17.4.59 --discovery-token-unsafe-skip-ca-verification
[preflight] Running pre-flight checks.
	[WARNING FileExisting-crictl]: crictl not found in system path
[preflight] Some fatal errors occurred:
	[ERROR Swap]: running with swap on is not supported. Please disable swap
[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`
```

```
vagrant@stretch:~$ sudo kubeadm join --token=86257a.d063ca0eedd9e165 172.17.4.59:6443 --discovery-token-unsafe-skip-ca-verification --ignore-preflight-errors=Swap
[preflight] Running pre-flight checks.
	[WARNING Swap]: running with swap on is not supported. Please disable swap
	[WARNING FileExisting-crictl]: crictl not found in system path
[discovery] Trying to connect to API Server "172.17.4.59:6443"
[discovery] Created cluster-info discovery client, requesting info from "https://172.17.4.59:6443"
[discovery] Cluster info signature and contents are valid and no TLS pinning was specified, will use API Server "172.17.4.59:6443"
[discovery] Successfully established connection with API Server "172.17.4.59:6443"

This node has joined the cluster:
* Certificate signing request was sent to master and a response
  was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the master to see this node join the cluster.
```

```
vagrant@stretch:~$ kubectl get nodes
NAME                  STATUS     ROLES     AGE       VERSION
kubedev-172-17-4-55   NotReady   <none>    12h       v1.9.0
kubedev-172-17-4-59   Ready      master    1d        v1.9.0
kubedev-172-17-4-65   Ready      <none>    5m        v1.9.0
rookdev-172-17-4-63   Ready      <none>    4h        v1.9.0
```

```
vagrant@stretch:~$ kubectl get pods -o wide -n kube-system
NAME                                          READY     STATUS     RESTARTS   AGE       IP            NODE
etcd-kubedev-172-17-4-59                      1/1       Running    0          1d        172.17.4.59   kubedev-172-17-4-59
kube-apiserver-kubedev-172-17-4-59            1/1       Running    0          1d        172.17.4.59   kubedev-172-17-4-59
kube-controller-manager-kubedev-172-17-4-59   1/1       Running    0          1d        172.17.4.59   kubedev-172-17-4-59
kube-dns-6f4fd4bdf-r8c7h                      3/3       Running    0          1d        10.244.0.2    kubedev-172-17-4-59
kube-flannel-ds-29rcz                         1/1       Running    0          6m        172.17.4.65   kubedev-172-17-4-65
kube-flannel-ds-7pvdh                         1/1       Running    0          12h       172.17.4.59   kubedev-172-17-4-59
kube-flannel-ds-89hwq                         1/1       Running    0          4h        172.17.4.63   rookdev-172-17-4-63
kube-flannel-ds-x2l9q                         1/1       NodeLost   1          12h       172.17.4.55   kubedev-172-17-4-55
kube-proxy-fbbdd                              1/1       Running    0          4h        172.17.4.63   rookdev-172-17-4-63
kube-proxy-g6qr8                              1/1       Running    0          6m        172.17.4.65   kubedev-172-17-4-65
kube-proxy-ls4nx                              1/1       NodeLost   1          12h       172.17.4.55   kubedev-172-17-4-55
kube-proxy-xj2vb                              1/1       Running    0          1d        172.17.4.59   kubedev-172-17-4-59
kube-scheduler-kubedev-172-17-4-59            1/1       Running    0          1d        172.17.4.59   kubedev-172-17-4-59
```

```
vagrant@stretch:~$ ip a show cni0
7: cni0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP group default qlen 1000
    link/ether 0a:58:0a:f4:04:01 brd ff:ff:ff:ff:ff:ff
    inet 10.244.4.1/24 scope global cni0
       valid_lft forever preferred_lft forever
    inet6 fe80::819:e8ff:fe58:7d35/64 scope link 
       valid_lft forever preferred_lft forever
```