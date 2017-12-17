
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