# Provide VM with VirtualBox and Vagrant

VirtualBox
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ /Applications/VirtualBox.app/Contents/MacOS/VBoxManage --version
5.2.36r135684
```

Vagrant
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant version
Installed Version: 2.2.6
Latest Version: 2.2.7
 
To upgrade to the latest version, visit the downloads page and
download and install the latest version of Vagrant from the URL
below:

  https://www.vagrantup.com/downloads.html

If you're curious what changed in the latest release, view the
CHANGELOG below:

  https://github.com/hashicorp/vagrant/blob/v2.2.7/CHANGELOG.md
```

vagrant-vbguest plugin https://github.com/dotless-de/vagrant-vbguest

vagrant-disksize plugin https://github.com/sprotheroe/vagrant-disksize
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant plugin install vagrant-disksize
Installing the 'vagrant-disksize' plugin. This can take a few minutes...
Fetching: vagrant-disksize-0.1.3.gem (100%)
Installed the plugin 'vagrant-disksize (0.1.3)'!
```

## CentOS 7

http://cloud.centos.org/centos/7/vagrant/x86_64/images/CentOS-7-x86_64-Vagrant-2001_01.VirtualBox.box

Up
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Box 'centos7-cloud' could not be found. Attempting to find and install...
    default: Box Provider: virtualbox
    default: Box Version: >= 0
==> default: Box file was not detected as metadata. Adding it directly...
==> default: Adding box 'centos7-cloud' (v0) for provider: virtualbox
    default: Downloading: http://cloud.centos.org/centos/7/vagrant/x86_64/images/CentOS-7-x86_64-Vagrant-2001_01.VirtualBox.box
==> default: Successfully added box 'centos7-cloud' (v0) for 'virtualbox'!
==> default: Importing base box 'centos7-cloud'...
==> default: Matching MAC address for NAT networking...
==> default: Setting the name of the VM: go-to-kubernetes_default_1582673575531_74251
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Disk cannot be decreased in size. 20480 MB requested but disk is already 40960 MB.
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
==> default: Machine booted and ready!
[default] No Virtualbox Guest Additions installation found.
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirror.bit.edu.cn
 * extras: mirrors.huaweicloud.com
 * updates: mirror.bit.edu.cn
Package 1:make-3.82-24.el7.x86_64 already installed and latest version
Package bzip2-1.0.6-13.el7.x86_64 already installed and latest version
Resolving Dependencies
--> Running transaction check
---> Package binutils.x86_64 0:2.27-41.base.el7_7.1 will be updated
---> Package binutils.x86_64 0:2.27-41.base.el7_7.2 will be an update
---> Package elfutils-libelf-devel.x86_64 0:0.176-2.el7 will be installed
--> Processing Dependency: pkgconfig(zlib) for package: elfutils-libelf-devel-0.176-2.el7.x86_64
---> Package gcc.x86_64 0:4.8.5-39.el7 will be installed
--> Processing Dependency: cpp = 4.8.5-39.el7 for package: gcc-4.8.5-39.el7.x86_64
--> Processing Dependency: glibc-devel >= 2.2.90-12 for package: gcc-4.8.5-39.el7.x86_64
--> Processing Dependency: libmpfr.so.4()(64bit) for package: gcc-4.8.5-39.el7.x86_64
--> Processing Dependency: libmpc.so.3()(64bit) for package: gcc-4.8.5-39.el7.x86_64
---> Package kernel-devel.x86_64 0:3.10.0-1062.9.1.el7 will be installed
---> Package kernel-devel.x86_64 0:3.10.0-1062.12.1.el7 will be installed
---> Package perl.x86_64 4:5.16.3-294.el7_6 will be installed
--> Processing Dependency: perl-libs = 4:5.16.3-294.el7_6 for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Socket) >= 1.3 for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Scalar::Util) >= 1.10 for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl-macros for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl-libs for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(threads::shared) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(threads) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(constant) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Time::Local) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Time::HiRes) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Storable) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Socket) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Scalar::Util) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Pod::Simple::XHTML) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Pod::Simple::Search) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Getopt::Long) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Filter::Util::Call) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(File::Temp) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(File::Spec::Unix) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(File::Spec::Functions) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(File::Spec) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(File::Path) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Exporter) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Cwd) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: perl(Carp) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Processing Dependency: libperl.so()(64bit) for package: 4:perl-5.16.3-294.el7_6.x86_64
--> Running transaction check
---> Package cpp.x86_64 0:4.8.5-39.el7 will be installed
---> Package glibc-devel.x86_64 0:2.17-292.el7 will be installed
--> Processing Dependency: glibc-headers = 2.17-292.el7 for package: glibc-devel-2.17-292.el7.x86_64
--> Processing Dependency: glibc-headers for package: glibc-devel-2.17-292.el7.x86_64
---> Package libmpc.x86_64 0:1.0.1-3.el7 will be installed
---> Package mpfr.x86_64 0:3.1.1-4.el7 will be installed
---> Package perl-Carp.noarch 0:1.26-244.el7 will be installed
---> Package perl-Exporter.noarch 0:5.68-3.el7 will be installed
---> Package perl-File-Path.noarch 0:2.09-2.el7 will be installed
---> Package perl-File-Temp.noarch 0:0.23.01-3.el7 will be installed
---> Package perl-Filter.x86_64 0:1.49-3.el7 will be installed
---> Package perl-Getopt-Long.noarch 0:2.40-3.el7 will be installed
--> Processing Dependency: perl(Pod::Usage) >= 1.14 for package: perl-Getopt-Long-2.40-3.el7.noarch
--> Processing Dependency: perl(Text::ParseWords) for package: perl-Getopt-Long-2.40-3.el7.noarch
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
---> Package perl-libs.x86_64 4:5.16.3-294.el7_6 will be installed
---> Package perl-macros.x86_64 4:5.16.3-294.el7_6 will be installed
---> Package perl-threads.x86_64 0:1.87-4.el7 will be installed
---> Package perl-threads-shared.x86_64 0:1.43-6.el7 will be installed
---> Package zlib-devel.x86_64 0:1.2.7-18.el7 will be installed
--> Running transaction check
---> Package glibc-headers.x86_64 0:2.17-292.el7 will be installed
--> Processing Dependency: kernel-headers >= 2.2.1 for package: glibc-headers-2.17-292.el7.x86_64
--> Processing Dependency: kernel-headers for package: glibc-headers-2.17-292.el7.x86_64
---> Package perl-Encode.x86_64 0:2.51-7.el7 will be installed
---> Package perl-Pod-Escapes.noarch 1:1.04-294.el7_6 will be installed
---> Package perl-Pod-Usage.noarch 0:1.63-3.el7 will be installed
--> Processing Dependency: perl(Pod::Text) >= 3.15 for package: perl-Pod-Usage-1.63-3.el7.noarch
--> Processing Dependency: perl-Pod-Perldoc for package: perl-Pod-Usage-1.63-3.el7.noarch
---> Package perl-Text-ParseWords.noarch 0:3.29-4.el7 will be installed
--> Running transaction check
---> Package kernel-headers.x86_64 0:3.10.0-1062.12.1.el7 will be installed
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
 Package                   Arch      Version                   Repository  Size
================================================================================
Installing:
 elfutils-libelf-devel     x86_64    0.176-2.el7               base        39 k
 gcc                       x86_64    4.8.5-39.el7              base        16 M
 kernel-devel              x86_64    3.10.0-1062.9.1.el7       updates     18 M
 kernel-devel              x86_64    3.10.0-1062.12.1.el7      updates     18 M
 perl                      x86_64    4:5.16.3-294.el7_6        base       8.0 M
Updating:
 binutils                  x86_64    2.27-41.base.el7_7.2      updates    5.9 M
Installing for dependencies:
 cpp                       x86_64    4.8.5-39.el7              base       5.9 M
 glibc-devel               x86_64    2.17-292.el7              base       1.1 M
 glibc-headers             x86_64    2.17-292.el7              base       687 k
 kernel-headers            x86_64    3.10.0-1062.12.1.el7      updates    8.7 M
 libmpc                    x86_64    1.0.1-3.el7               base        51 k
 mpfr                      x86_64    3.1.1-4.el7               base       203 k
 perl-Carp                 noarch    1.26-244.el7              base        19 k
 perl-Encode               x86_64    2.51-7.el7                base       1.5 M
 perl-Exporter             noarch    5.68-3.el7                base        28 k
 perl-File-Path            noarch    2.09-2.el7                base        26 k
 perl-File-Temp            noarch    0.23.01-3.el7             base        56 k
 perl-Filter               x86_64    1.49-3.el7                base        76 k
 perl-Getopt-Long          noarch    2.40-3.el7                base        56 k
 perl-HTTP-Tiny            noarch    0.033-3.el7               base        38 k
 perl-PathTools            x86_64    3.40-5.el7                base        82 k
 perl-Pod-Escapes          noarch    1:1.04-294.el7_6          base        51 k
 perl-Pod-Perldoc          noarch    3.20-4.el7                base        87 k
 perl-Pod-Simple           noarch    1:3.28-4.el7              base       216 k
 perl-Pod-Usage            noarch    1.63-3.el7                base        27 k
 perl-Scalar-List-Utils    x86_64    1.27-248.el7              base        36 k
 perl-Socket               x86_64    2.010-4.el7               base        49 k
 perl-Storable             x86_64    2.45-3.el7                base        77 k
 perl-Text-ParseWords      noarch    3.29-4.el7                base        14 k
 perl-Time-HiRes           x86_64    4:1.9725-3.el7            base        45 k
 perl-Time-Local           noarch    1.2300-2.el7              base        24 k
 perl-constant             noarch    1.27-2.el7                base        19 k
 perl-libs                 x86_64    4:5.16.3-294.el7_6        base       688 k
 perl-macros               x86_64    4:5.16.3-294.el7_6        base        44 k
 perl-parent               noarch    1:0.225-244.el7           base        12 k
 perl-podlators            noarch    2.5.1-3.el7               base       112 k
 perl-threads              x86_64    1.87-4.el7                base        49 k
 perl-threads-shared       x86_64    1.43-6.el7                base        39 k
 zlib-devel                x86_64    1.2.7-18.el7              base        50 k

Transaction Summary
================================================================================
Install  5 Packages (+33 Dependent packages)
Upgrade  1 Package

Total download size: 85 M
Downloading packages:
No Presto metadata available for updates
Public key for elfutils-libelf-devel-0.176-2.el7.x86_64.rpm is not installed
warning: /var/cache/yum/x86_64/7/base/packages/elfutils-libelf-devel-0.176-2.el7.x86_64.rpm: Header V3 RSA/SHA256 Signature, key ID f4a80eb5: NOKEY
Public key for binutils-2.27-41.base.el7_7.2.x86_64.rpm is not installed
--------------------------------------------------------------------------------
Total                                              6.6 MB/s |  85 MB  00:13     
Retrieving key from file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Importing GPG key 0xF4A80EB5:
 Userid     : "CentOS-7 Key (CentOS 7 Official Signing Key) <security@centos.org>"
 Fingerprint: 6341 ab27 53d7 8a78 a7c2 7bb1 24c6 a8a7 f4a8 0eb5
 Package    : centos-release-7-7.1908.0.el7.centos.x86_64 (@anaconda)
 From       : /etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : mpfr-3.1.1-4.el7.x86_64                                     1/40 
  Installing : libmpc-1.0.1-3.el7.x86_64                                   2/40 
  Installing : cpp-4.8.5-39.el7.x86_64                                     3/40 
  Installing : 1:perl-parent-0.225-244.el7.noarch                          4/40 
  Installing : perl-HTTP-Tiny-0.033-3.el7.noarch                           5/40 
  Installing : perl-podlators-2.5.1-3.el7.noarch                           6/40 
  Installing : perl-Pod-Perldoc-3.20-4.el7.noarch                          7/40 
  Installing : 1:perl-Pod-Escapes-1.04-294.el7_6.noarch                    8/40 
  Installing : perl-Encode-2.51-7.el7.x86_64                               9/40 
  Installing : perl-Text-ParseWords-3.29-4.el7.noarch                     10/40 
  Installing : perl-Pod-Usage-1.63-3.el7.noarch                           11/40 
  Installing : 4:perl-libs-5.16.3-294.el7_6.x86_64                        12/40 
  Installing : 4:perl-macros-5.16.3-294.el7_6.x86_64                      13/40 
  Installing : perl-threads-1.87-4.el7.x86_64                             14/40 
  Installing : 4:perl-Time-HiRes-1.9725-3.el7.x86_64                      15/40 
  Installing : perl-Exporter-5.68-3.el7.noarch                            16/40 
  Installing : perl-constant-1.27-2.el7.noarch                            17/40 
  Installing : perl-Time-Local-1.2300-2.el7.noarch                        18/40 
  Installing : perl-Socket-2.010-4.el7.x86_64                             19/40 
  Installing : perl-Carp-1.26-244.el7.noarch                              20/40 
  Installing : perl-Storable-2.45-3.el7.x86_64                            21/40 
  Installing : perl-threads-shared-1.43-6.el7.x86_64                      22/40 
  Installing : 1:perl-Pod-Simple-3.28-4.el7.noarch                        23/40 
  Installing : perl-PathTools-3.40-5.el7.x86_64                           24/40 
  Installing : perl-Scalar-List-Utils-1.27-248.el7.x86_64                 25/40 
  Installing : perl-File-Temp-0.23.01-3.el7.noarch                        26/40 
  Installing : perl-File-Path-2.09-2.el7.noarch                           27/40 
  Installing : perl-Filter-1.49-3.el7.x86_64                              28/40 
  Installing : perl-Getopt-Long-2.40-3.el7.noarch                         29/40 
  Installing : 4:perl-5.16.3-294.el7_6.x86_64                             30/40 
  Installing : kernel-headers-3.10.0-1062.12.1.el7.x86_64                 31/40 
  Installing : glibc-headers-2.17-292.el7.x86_64                          32/40 
  Installing : glibc-devel-2.17-292.el7.x86_64                            33/40 
  Installing : zlib-devel-1.2.7-18.el7.x86_64                             34/40 
  Updating   : binutils-2.27-41.base.el7_7.2.x86_64                       35/40 
  Installing : gcc-4.8.5-39.el7.x86_64                                    36/40 
  Installing : elfutils-libelf-devel-0.176-2.el7.x86_64                   37/40 
  Installing : kernel-devel-3.10.0-1062.9.1.el7.x86_64                    38/40 
  Installing : kernel-devel-3.10.0-1062.12.1.el7.x86_64                   39/40 
  Cleanup    : binutils-2.27-41.base.el7_7.1.x86_64                       40/40 
  Verifying  : perl-HTTP-Tiny-0.033-3.el7.noarch                           1/40 
  Verifying  : binutils-2.27-41.base.el7_7.2.x86_64                        2/40 
  Verifying  : 4:perl-libs-5.16.3-294.el7_6.x86_64                         3/40 
  Verifying  : perl-threads-shared-1.43-6.el7.x86_64                       4/40 
  Verifying  : 4:perl-Time-HiRes-1.9725-3.el7.x86_64                       5/40 
  Verifying  : 1:perl-Pod-Escapes-1.04-294.el7_6.noarch                    6/40 
  Verifying  : perl-threads-1.87-4.el7.x86_64                              7/40 
  Verifying  : perl-Exporter-5.68-3.el7.noarch                             8/40 
  Verifying  : perl-constant-1.27-2.el7.noarch                             9/40 
  Verifying  : perl-PathTools-3.40-5.el7.x86_64                           10/40 
  Verifying  : gcc-4.8.5-39.el7.x86_64                                    11/40 
  Verifying  : 1:perl-parent-0.225-244.el7.noarch                         12/40 
  Verifying  : zlib-devel-1.2.7-18.el7.x86_64                             13/40 
  Verifying  : kernel-devel-3.10.0-1062.9.1.el7.x86_64                    14/40 
  Verifying  : perl-File-Temp-0.23.01-3.el7.noarch                        15/40 
  Verifying  : 1:perl-Pod-Simple-3.28-4.el7.noarch                        16/40 
  Verifying  : perl-Time-Local-1.2300-2.el7.noarch                        17/40 
  Verifying  : 4:perl-macros-5.16.3-294.el7_6.x86_64                      18/40 
  Verifying  : perl-Socket-2.010-4.el7.x86_64                             19/40 
  Verifying  : perl-Encode-2.51-7.el7.x86_64                              20/40 
  Verifying  : perl-Carp-1.26-244.el7.noarch                              21/40 
  Verifying  : kernel-devel-3.10.0-1062.12.1.el7.x86_64                   22/40 
  Verifying  : elfutils-libelf-devel-0.176-2.el7.x86_64                   23/40 
  Verifying  : perl-Storable-2.45-3.el7.x86_64                            24/40 
  Verifying  : perl-Scalar-List-Utils-1.27-248.el7.x86_64                 25/40 
  Verifying  : libmpc-1.0.1-3.el7.x86_64                                  26/40 
  Verifying  : perl-Pod-Usage-1.63-3.el7.noarch                           27/40 
  Verifying  : kernel-headers-3.10.0-1062.12.1.el7.x86_64                 28/40 
  Verifying  : glibc-devel-2.17-292.el7.x86_64                            29/40 
  Verifying  : perl-Pod-Perldoc-3.20-4.el7.noarch                         30/40 
  Verifying  : perl-podlators-2.5.1-3.el7.noarch                          31/40 
  Verifying  : perl-File-Path-2.09-2.el7.noarch                           32/40 
  Verifying  : mpfr-3.1.1-4.el7.x86_64                                    33/40 
  Verifying  : perl-Filter-1.49-3.el7.x86_64                              34/40 
  Verifying  : perl-Getopt-Long-2.40-3.el7.noarch                         35/40 
  Verifying  : perl-Text-ParseWords-3.29-4.el7.noarch                     36/40 
  Verifying  : 4:perl-5.16.3-294.el7_6.x86_64                             37/40 
  Verifying  : cpp-4.8.5-39.el7.x86_64                                    38/40 
  Verifying  : glibc-headers-2.17-292.el7.x86_64                          39/40 
  Verifying  : binutils-2.27-41.base.el7_7.1.x86_64                       40/40 

Installed:
  elfutils-libelf-devel.x86_64 0:0.176-2.el7                                    
  gcc.x86_64 0:4.8.5-39.el7                                                     
  kernel-devel.x86_64 0:3.10.0-1062.9.1.el7                                     
  kernel-devel.x86_64 0:3.10.0-1062.12.1.el7                                    
  perl.x86_64 4:5.16.3-294.el7_6                                                

Dependency Installed:
  cpp.x86_64 0:4.8.5-39.el7                                                     
  glibc-devel.x86_64 0:2.17-292.el7                                             
  glibc-headers.x86_64 0:2.17-292.el7                                           
  kernel-headers.x86_64 0:3.10.0-1062.12.1.el7                                  
  libmpc.x86_64 0:1.0.1-3.el7                                                   
  mpfr.x86_64 0:3.1.1-4.el7                                                     
  perl-Carp.noarch 0:1.26-244.el7                                               
  perl-Encode.x86_64 0:2.51-7.el7                                               
  perl-Exporter.noarch 0:5.68-3.el7                                             
  perl-File-Path.noarch 0:2.09-2.el7                                            
  perl-File-Temp.noarch 0:0.23.01-3.el7                                         
  perl-Filter.x86_64 0:1.49-3.el7                                               
  perl-Getopt-Long.noarch 0:2.40-3.el7                                          
  perl-HTTP-Tiny.noarch 0:0.033-3.el7                                           
  perl-PathTools.x86_64 0:3.40-5.el7                                            
  perl-Pod-Escapes.noarch 1:1.04-294.el7_6                                      
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
  perl-libs.x86_64 4:5.16.3-294.el7_6                                           
  perl-macros.x86_64 4:5.16.3-294.el7_6                                         
  perl-parent.noarch 1:0.225-244.el7                                            
  perl-podlators.noarch 0:2.5.1-3.el7                                           
  perl-threads.x86_64 0:1.87-4.el7                                              
  perl-threads-shared.x86_64 0:1.43-6.el7                                       
  zlib-devel.x86_64 0:1.2.7-18.el7                                              

Updated:
  binutils.x86_64 0:2.27-41.base.el7_7.2                                        

Complete!
Copy iso file /Applications/VirtualBox.app/Contents/MacOS/VBoxGuestAdditions.iso into the box /tmp/VBoxGuestAdditions.iso
Mounting Virtualbox Guest Additions ISO to: /mnt
mount: /dev/loop0 is write-protected, mounting read-only
Installing Virtualbox Guest Additions 5.2.36 - guest version is unknown
Verifying archive integrity... All good.
Uncompressing VirtualBox 5.2.36 Guest Additions for Linux........
VirtualBox Guest Additions installer
Copying additional installer modules ...
Installing additional modules ...
VirtualBox Guest Additions: Building the VirtualBox Guest Additions kernel 
modules.  This may take a while.
VirtualBox Guest Additions: To build modules for other installed kernels, run
VirtualBox Guest Additions:   /sbin/rcvboxadd quicksetup <version>
VirtualBox Guest Additions: Building the modules for kernel 
3.10.0-1062.9.1.el7.x86_64.
VirtualBox Guest Additions: Starting.
Redirecting to /bin/systemctl start vboxadd.service
Redirecting to /bin/systemctl start vboxadd-service.service
Unmounting Virtualbox Guest Additions ISO from: /mnt
==> default: Checking for guest additions in VM...
==> default: Configuring and enabling network interfaces...
==> default: Mounting shared folders...
    default: /Users/fanhongling/go => /Users/fanhongling/go
    default: /Users/fanhongling/Downloads => /Users/fanhongling/Downloads
```

ssh
```
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant ssh
[vagrant@localhost ~]$ 
```

centos
```
[vagrant@localhost ~]$ cat /etc/centos-release
CentOS Linux release 7.7.1908 (Core)
```

### docker

search
```
[vagrant@localhost ~]$ sudo yum search docker
Failed to set locale, defaulting to C
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirror.bit.edu.cn
 * extras: mirrors.huaweicloud.com
 * updates: mirror.bit.edu.cn
================================================================= N/S matched: docker =================================================================
cockpit-docker.x86_64 : Cockpit user interface for Docker containers
docker-client.x86_64 : Client side files for Docker
docker-client-latest.x86_64 : Client side files for Docker
docker-common.x86_64 : Common files for docker and docker-latest
docker-distribution.x86_64 : Docker toolset to pack, ship, store, and deliver content
docker-latest-logrotate.x86_64 : cron job to run logrotate on Docker containers
docker-latest-v1.10-migrator.x86_64 : Calculates SHA256 checksums for docker layer content
docker-logrotate.x86_64 : cron job to run logrotate on Docker containers
docker-lvm-plugin.x86_64 : Docker volume driver for lvm volumes
docker-registry.x86_64 : Registry server for Docker
docker-v1.10-migrator.x86_64 : Calculates SHA256 checksums for docker layer content
pcp-pmda-docker.x86_64 : Performance Co-Pilot (PCP) metrics from the Docker daemon
podman-docker.noarch : Emulate Docker CLI using podman
python-docker-py.noarch : An API client for docker written in Python
python-docker-pycreds.noarch : Python bindings for the docker credentials store API
docker.x86_64 : Automates deployment of containerized applications
docker-latest.x86_64 : Automates deployment of containerized applications
docker-novolume-plugin.x86_64 : Block container starts with local volumes defined
oci-systemd-hook.x86_64 : OCI systemd hook for docker
oci-umount.x86_64 : OCI umount hook for docker

  Name and summary matches only, use "search all" for everything.
```

Version
```
[vagrant@localhost ~]$ sudo yum list | grep -e '^docker.*'
Failed to set locale, defaulting to C
docker.x86_64                               2:1.13.1-108.git4ef4b30.el7.centos
docker-client.x86_64                        2:1.13.1-108.git4ef4b30.el7.centos
docker-client-latest.x86_64                 1.13.1-58.git87f2fab.el7.centos
docker-common.x86_64                        2:1.13.1-108.git4ef4b30.el7.centos
docker-distribution.x86_64                  2.6.2-2.git48294d9.el7     extras   
docker-latest.x86_64                        1.13.1-58.git87f2fab.el7.centos
docker-latest-logrotate.x86_64              1.13.1-58.git87f2fab.el7.centos
docker-latest-v1.10-migrator.x86_64         1.13.1-58.git87f2fab.el7.centos
docker-logrotate.x86_64                     2:1.13.1-108.git4ef4b30.el7.centos
docker-lvm-plugin.x86_64                    2:1.13.1-108.git4ef4b30.el7.centos
docker-novolume-plugin.x86_64               2:1.13.1-108.git4ef4b30.el7.centos
docker-registry.x86_64                      0.9.1-7.el7                extras   
docker-v1.10-migrator.x86_64                2:1.13.1-108.git4ef4b30.el7.centos
```

__docker-ce distribution__

https://docs.docker.com/install/linux/docker-ce/centos/

before
```
[vagrant@localhost ~]$ sudo yum install -y yum-utils device-mapper-persistent-data lvm2
Failed to set locale, defaulting to C
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirror.bit.edu.cn
 * extras: mirrors.huaweicloud.com
 * updates: mirror.bit.edu.cn
Package yum-utils-1.1.31-52.el7.noarch already installed and latest version
Resolving Dependencies
--> Running transaction check
---> Package device-mapper-persistent-data.x86_64 0:0.8.5-1.el7 will be installed
--> Processing Dependency: libaio.so.1(LIBAIO_0.4)(64bit) for package: device-mapper-persistent-data-0.8.5-1.el7.x86_64
--> Processing Dependency: libaio.so.1(LIBAIO_0.1)(64bit) for package: device-mapper-persistent-data-0.8.5-1.el7.x86_64
--> Processing Dependency: libaio.so.1()(64bit) for package: device-mapper-persistent-data-0.8.5-1.el7.x86_64
---> Package lvm2.x86_64 7:2.02.185-2.el7_7.2 will be installed
--> Processing Dependency: lvm2-libs = 7:2.02.185-2.el7_7.2 for package: 7:lvm2-2.02.185-2.el7_7.2.x86_64
--> Processing Dependency: liblvm2app.so.2.2(Base)(64bit) for package: 7:lvm2-2.02.185-2.el7_7.2.x86_64
--> Processing Dependency: libdevmapper-event.so.1.02(Base)(64bit) for package: 7:lvm2-2.02.185-2.el7_7.2.x86_64
--> Processing Dependency: liblvm2app.so.2.2()(64bit) for package: 7:lvm2-2.02.185-2.el7_7.2.x86_64
--> Processing Dependency: libdevmapper-event.so.1.02()(64bit) for package: 7:lvm2-2.02.185-2.el7_7.2.x86_64
--> Running transaction check
---> Package device-mapper-event-libs.x86_64 7:1.02.158-2.el7_7.2 will be installed
---> Package libaio.x86_64 0:0.3.109-13.el7 will be installed
---> Package lvm2-libs.x86_64 7:2.02.185-2.el7_7.2 will be installed
--> Processing Dependency: device-mapper-event = 7:1.02.158-2.el7_7.2 for package: 7:lvm2-libs-2.02.185-2.el7_7.2.x86_64
--> Running transaction check
---> Package device-mapper-event.x86_64 7:1.02.158-2.el7_7.2 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

=======================================================================================================================================================
 Package                                          Arch                      Version                                   Repository                  Size
=======================================================================================================================================================
Installing:
 device-mapper-persistent-data                    x86_64                    0.8.5-1.el7                               base                       423 k
 lvm2                                             x86_64                    7:2.02.185-2.el7_7.2                      updates                    1.3 M
Installing for dependencies:
 device-mapper-event                              x86_64                    7:1.02.158-2.el7_7.2                      updates                    190 k
 device-mapper-event-libs                         x86_64                    7:1.02.158-2.el7_7.2                      updates                    189 k
 libaio                                           x86_64                    0.3.109-13.el7                            base                        24 k
 lvm2-libs                                        x86_64                    7:2.02.185-2.el7_7.2                      updates                    1.1 M

Transaction Summary
=======================================================================================================================================================
Install  2 Packages (+4 Dependent packages)

Total download size: 3.2 M
Installed size: 8.0 M
Downloading packages:
(1/6): device-mapper-event-1.02.158-2.el7_7.2.x86_64.rpm                                                                        | 190 kB  00:00:00     
(2/6): libaio-0.3.109-13.el7.x86_64.rpm                                                                                         |  24 kB  00:00:00     
(3/6): lvm2-libs-2.02.185-2.el7_7.2.x86_64.rpm                                                                                  | 1.1 MB  00:00:00     
(4/6): device-mapper-event-libs-1.02.158-2.el7_7.2.x86_64.rpm                                                                   | 189 kB  00:00:00     
(5/6): device-mapper-persistent-data-0.8.5-1.el7.x86_64.rpm                                                                     | 423 kB  00:00:00     
(6/6): lvm2-2.02.185-2.el7_7.2.x86_64.rpm                                                                                       | 1.3 MB  00:00:01     
-------------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                                  1.9 MB/s | 3.2 MB  00:00:01     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : 7:device-mapper-event-libs-1.02.158-2.el7_7.2.x86_64                                                                                1/6 
  Installing : libaio-0.3.109-13.el7.x86_64                                                                                                        2/6 
  Installing : device-mapper-persistent-data-0.8.5-1.el7.x86_64                                                                                    3/6 
  Installing : 7:device-mapper-event-1.02.158-2.el7_7.2.x86_64                                                                                     4/6 
  Installing : 7:lvm2-libs-2.02.185-2.el7_7.2.x86_64                                                                                               5/6 
  Installing : 7:lvm2-2.02.185-2.el7_7.2.x86_64                                                                                                    6/6 
  Verifying  : device-mapper-persistent-data-0.8.5-1.el7.x86_64                                                                                    1/6 
  Verifying  : 7:lvm2-2.02.185-2.el7_7.2.x86_64                                                                                                    2/6 
  Verifying  : 7:lvm2-libs-2.02.185-2.el7_7.2.x86_64                                                                                               3/6 
  Verifying  : 7:device-mapper-event-1.02.158-2.el7_7.2.x86_64                                                                                     4/6 
  Verifying  : libaio-0.3.109-13.el7.x86_64                                                                                                        5/6 
  Verifying  : 7:device-mapper-event-libs-1.02.158-2.el7_7.2.x86_64                                                                                6/6 

Installed:
  device-mapper-persistent-data.x86_64 0:0.8.5-1.el7                                  lvm2.x86_64 7:2.02.185-2.el7_7.2                                 

Dependency Installed:
  device-mapper-event.x86_64 7:1.02.158-2.el7_7.2       device-mapper-event-libs.x86_64 7:1.02.158-2.el7_7.2       libaio.x86_64 0:0.3.109-13.el7      
  lvm2-libs.x86_64 7:2.02.185-2.el7_7.2                

Complete!
```

docker-ce Repository
```
[vagrant@localhost ~]$ sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
Failed to set locale, defaulting to C
Loaded plugins: fastestmirror
adding repo from: https://download.docker.com/linux/centos/docker-ce.repo
grabbing file https://download.docker.com/linux/centos/docker-ce.repo to /etc/yum.repos.d/docker-ce.repo
repo saved to /etc/yum.repos.d/docker-ce.repo
```

configuration background
```
[vagrant@localhost ~]$ head -7 /etc/yum.repos.d/docker-ce.repo 
[docker-ce-stable]
name=Docker CE Stable - $basearch
baseurl=https://download.docker.com/linux/centos/7/$basearch/stable
enabled=1
gpgcheck=1
gpgkey=https://download.docker.com/linux/centos/gpg

```

docker-ce version
```
[vagrant@localhost ~]$ sudo yum list | grep docker-ce
Failed to set locale, defaulting to C
containerd.io.x86_64                        1.2.10-3.2.el7             docker-ce-stable
docker-ce.x86_64                            3:19.03.6-3.el7            docker-ce-stable
docker-ce-cli.x86_64                        1:19.03.6-3.el7            docker-ce-stable
docker-ce-selinux.noarch                    17.03.3.ce-1.el7           docker-ce-stable
```

install
```
[vagrant@localhost ~]$ sudo yum install -y docker-ce docker-ce-cli containerd.io
Failed to set locale, defaulting to C
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirror.bit.edu.cn
 * extras: mirrors.huaweicloud.com
 * updates: mirror.bit.edu.cn
Resolving Dependencies
--> Running transaction check
---> Package containerd.io.x86_64 0:1.2.10-3.2.el7 will be installed
--> Processing Dependency: container-selinux >= 2:2.74 for package: containerd.io-1.2.10-3.2.el7.x86_64
---> Package docker-ce.x86_64 3:19.03.6-3.el7 will be installed
--> Processing Dependency: libcgroup for package: 3:docker-ce-19.03.6-3.el7.x86_64
---> Package docker-ce-cli.x86_64 1:19.03.6-3.el7 will be installed
--> Running transaction check
---> Package container-selinux.noarch 2:2.107-3.el7 will be installed
--> Processing Dependency: policycoreutils-python for package: 2:container-selinux-2.107-3.el7.noarch
---> Package libcgroup.x86_64 0:0.41-21.el7 will be installed
--> Running transaction check
---> Package policycoreutils-python.x86_64 0:2.5-33.el7 will be installed
--> Processing Dependency: setools-libs >= 3.3.8-4 for package: policycoreutils-python-2.5-33.el7.x86_64
--> Processing Dependency: libsemanage-python >= 2.5-14 for package: policycoreutils-python-2.5-33.el7.x86_64
--> Processing Dependency: audit-libs-python >= 2.1.3-4 for package: policycoreutils-python-2.5-33.el7.x86_64
--> Processing Dependency: python-IPy for package: policycoreutils-python-2.5-33.el7.x86_64
--> Processing Dependency: libqpol.so.1(VERS_1.4)(64bit) for package: policycoreutils-python-2.5-33.el7.x86_64
--> Processing Dependency: libqpol.so.1(VERS_1.2)(64bit) for package: policycoreutils-python-2.5-33.el7.x86_64
--> Processing Dependency: libapol.so.4(VERS_4.0)(64bit) for package: policycoreutils-python-2.5-33.el7.x86_64
--> Processing Dependency: checkpolicy for package: policycoreutils-python-2.5-33.el7.x86_64
--> Processing Dependency: libqpol.so.1()(64bit) for package: policycoreutils-python-2.5-33.el7.x86_64
--> Processing Dependency: libapol.so.4()(64bit) for package: policycoreutils-python-2.5-33.el7.x86_64
--> Running transaction check
---> Package audit-libs-python.x86_64 0:2.8.5-4.el7 will be installed
---> Package checkpolicy.x86_64 0:2.5-8.el7 will be installed
---> Package libsemanage-python.x86_64 0:2.5-14.el7 will be installed
---> Package python-IPy.noarch 0:0.75-6.el7 will be installed
---> Package setools-libs.x86_64 0:3.3.8-4.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

=======================================================================================================================================================
 Package                                    Arch                       Version                              Repository                            Size
=======================================================================================================================================================
Installing:
 containerd.io                              x86_64                     1.2.10-3.2.el7                       docker-ce-stable                      23 M
 docker-ce                                  x86_64                     3:19.03.6-3.el7                      docker-ce-stable                      24 M
 docker-ce-cli                              x86_64                     1:19.03.6-3.el7                      docker-ce-stable                      40 M
Installing for dependencies:
 audit-libs-python                          x86_64                     2.8.5-4.el7                          base                                  76 k
 checkpolicy                                x86_64                     2.5-8.el7                            base                                 295 k
 container-selinux                          noarch                     2:2.107-3.el7                        extras                                39 k
 libcgroup                                  x86_64                     0.41-21.el7                          base                                  66 k
 libsemanage-python                         x86_64                     2.5-14.el7                           base                                 113 k
 policycoreutils-python                     x86_64                     2.5-33.el7                           base                                 457 k
 python-IPy                                 noarch                     0.75-6.el7                           base                                  32 k
 setools-libs                               x86_64                     3.3.8-4.el7                          base                                 620 k

Transaction Summary
=======================================================================================================================================================
Install  3 Packages (+8 Dependent packages)

Total download size: 89 M
Installed size: 368 M
Downloading packages:
(1/11): audit-libs-python-2.8.5-4.el7.x86_64.rpm                                                                                |  76 kB  00:00:00     
(2/11): container-selinux-2.107-3.el7.noarch.rpm                                                                                |  39 kB  00:00:00     
(3/11): checkpolicy-2.5-8.el7.x86_64.rpm                                                                                        | 295 kB  00:00:00     
warning: /var/cache/yum/x86_64/7/docker-ce-stable/packages/containerd.io-1.2.10-3.2.el7.x86_64.rpm: Header V4 RSA/SHA512 Signature, key ID 621e9f35: NOKEY
Public key for containerd.io-1.2.10-3.2.el7.x86_64.rpm is not installed
(4/11): containerd.io-1.2.10-3.2.el7.x86_64.rpm                                                                                 |  23 MB  00:00:13     
(5/11): libsemanage-python-2.5-14.el7.x86_64.rpm                                                                                | 113 kB  00:00:00     
(6/11): libcgroup-0.41-21.el7.x86_64.rpm                                                                                        |  66 kB  00:00:00     
(7/11): python-IPy-0.75-6.el7.noarch.rpm                                                                                        |  32 kB  00:00:00     
(8/11): setools-libs-3.3.8-4.el7.x86_64.rpm                                                                                     | 620 kB  00:00:00     
(9/11): policycoreutils-python-2.5-33.el7.x86_64.rpm                                                                            | 457 kB  00:00:00     
(10/11): docker-ce-19.03.6-3.el7.x86_64.rpm                                                                                     |  24 MB  00:00:13     
(11/11): docker-ce-cli-19.03.6-3.el7.x86_64.rpm                                                                                 |  40 MB  00:00:10     
-------------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                                  3.7 MB/s |  89 MB  00:00:23     
Retrieving key from https://download.docker.com/linux/centos/gpg
Importing GPG key 0x621E9F35:
 Userid     : "Docker Release (CE rpm) <docker@docker.com>"
 Fingerprint: 060a 61c5 1b55 8a7f 742b 77aa c52f eb6b 621e 9f35
 From       : https://download.docker.com/linux/centos/gpg
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : libcgroup-0.41-21.el7.x86_64                                                                                                       1/11 
  Installing : setools-libs-3.3.8-4.el7.x86_64                                                                                                    2/11 
  Installing : audit-libs-python-2.8.5-4.el7.x86_64                                                                                               3/11 
  Installing : checkpolicy-2.5-8.el7.x86_64                                                                                                       4/11 
  Installing : python-IPy-0.75-6.el7.noarch                                                                                                       5/11 
  Installing : libsemanage-python-2.5-14.el7.x86_64                                                                                               6/11 
  Installing : policycoreutils-python-2.5-33.el7.x86_64                                                                                           7/11 
  Installing : 2:container-selinux-2.107-3.el7.noarch                                                                                             8/11 
  Installing : containerd.io-1.2.10-3.2.el7.x86_64                                                                                                9/11 
  Installing : 1:docker-ce-cli-19.03.6-3.el7.x86_64                                                                                              10/11 
  Installing : 3:docker-ce-19.03.6-3.el7.x86_64                                                                                                  11/11 
  Verifying  : 1:docker-ce-cli-19.03.6-3.el7.x86_64                                                                                               1/11 
  Verifying  : policycoreutils-python-2.5-33.el7.x86_64                                                                                           2/11 
  Verifying  : libsemanage-python-2.5-14.el7.x86_64                                                                                               3/11 
  Verifying  : 2:container-selinux-2.107-3.el7.noarch                                                                                             4/11 
  Verifying  : 3:docker-ce-19.03.6-3.el7.x86_64                                                                                                   5/11 
  Verifying  : python-IPy-0.75-6.el7.noarch                                                                                                       6/11 
  Verifying  : checkpolicy-2.5-8.el7.x86_64                                                                                                       7/11 
  Verifying  : audit-libs-python-2.8.5-4.el7.x86_64                                                                                               8/11 
  Verifying  : containerd.io-1.2.10-3.2.el7.x86_64                                                                                                9/11 
  Verifying  : setools-libs-3.3.8-4.el7.x86_64                                                                                                   10/11 
  Verifying  : libcgroup-0.41-21.el7.x86_64                                                                                                      11/11 

Installed:
  containerd.io.x86_64 0:1.2.10-3.2.el7               docker-ce.x86_64 3:19.03.6-3.el7               docker-ce-cli.x86_64 1:19.03.6-3.el7              

Dependency Installed:
  audit-libs-python.x86_64 0:2.8.5-4.el7          checkpolicy.x86_64 0:2.5-8.el7                  container-selinux.noarch 2:2.107-3.el7             
  libcgroup.x86_64 0:0.41-21.el7                  libsemanage-python.x86_64 0:2.5-14.el7          policycoreutils-python.x86_64 0:2.5-33.el7         
  python-IPy.noarch 0:0.75-6.el7                  setools-libs.x86_64 0:3.3.8-4.el7              

Complete!
```

start
```
[vagrant@localhost ~]$ sudo systemctl start docker
```

status
```
[vagrant@localhost ~]$ sudo systemctl status docker
● docker.service - Docker Application Container Engine
   Loaded: loaded (/usr/lib/systemd/system/docker.service; disabled; vendor preset: disabled)
   Active: active (running) since Wed 2020-02-26 00:00:41 UTC; 7s ago
     Docs: https://docs.docker.com
 Main PID: 7421 (dockerd)
    Tasks: 8
   Memory: 42.2M
   CGroup: /system.slice/docker.service
           └─7421 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock

Feb 26 00:00:41 localhost.localdomain dockerd[7421]: time="2020-02-26T00:00:41.229221283Z" level=info msg="scheme \"unix\" not registered, fa...le=grpc
Feb 26 00:00:41 localhost.localdomain dockerd[7421]: time="2020-02-26T00:00:41.229239698Z" level=info msg="ccResolverWrapper: sending update ...le=grpc
Feb 26 00:00:41 localhost.localdomain dockerd[7421]: time="2020-02-26T00:00:41.229250373Z" level=info msg="ClientConn switching balancer to \...le=grpc
Feb 26 00:00:41 localhost.localdomain dockerd[7421]: time="2020-02-26T00:00:41.262626965Z" level=info msg="Loading containers: start."
Feb 26 00:00:41 localhost.localdomain dockerd[7421]: time="2020-02-26T00:00:41.408754670Z" level=info msg="Default bridge (docker0) is assign...ddress"
Feb 26 00:00:41 localhost.localdomain dockerd[7421]: time="2020-02-26T00:00:41.471283625Z" level=info msg="Loading containers: done."
Feb 26 00:00:41 localhost.localdomain dockerd[7421]: time="2020-02-26T00:00:41.488930255Z" level=info msg="Docker daemon" commit=369ce74a3c g...19.03.6
Feb 26 00:00:41 localhost.localdomain dockerd[7421]: time="2020-02-26T00:00:41.489490453Z" level=info msg="Daemon has completed initialization"
Feb 26 00:00:41 localhost.localdomain dockerd[7421]: time="2020-02-26T00:00:41.518108044Z" level=info msg="API listen on /var/run/docker.sock"
Feb 26 00:00:41 localhost.localdomain systemd[1]: Started Docker Application Container Engine.
Hint: Some lines were ellipsized, use -l to show in full.
```

enabled at reboot
```
[vagrant@localhost ~]$ sudo systemctl enable docker
Created symlink from /etc/systemd/system/multi-user.target.wants/docker.service to /usr/lib/systemd/system/docker.service.
```

__non root user__

group
```
[vagrant@localhost ~]$ sudo usermod -aG docker vagrant
```

logout and login
```
[vagrant@localhost ~]$ exit
logout
Connection to 127.0.0.1 closed.
fanhonglingdeMacBook-Pro:go-to-kubernetes fanhongling$ vagrant ssh
Last login: Tue Feb 25 23:36:49 2020 from 10.0.2.2
```

__docker info__

```
[vagrant@localhost ~]$ docker info       
Client:
 Debug Mode: false

Server:
 Containers: 0
  Running: 0
  Paused: 0
  Stopped: 0
 Images: 0
 Server Version: 19.03.6
 Storage Driver: overlay2
  Backing Filesystem: xfs
  Supports d_type: true
  Native Overlay Diff: true
 Logging Driver: json-file
 Cgroup Driver: cgroupfs
 Plugins:
  Volume: local
  Network: bridge host ipvlan macvlan null overlay
  Log: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog
 Swarm: inactive
 Runtimes: runc
 Default Runtime: runc
 Init Binary: docker-init
 containerd version: b34a5c8af56e510852c35414db4c1f4fa6172339
 runc version: 3e425f80a8c931f88e6d94a8c831b9d5aa481657
 init version: fec3683
 Security Options:
  seccomp
   Profile: default
 Kernel Version: 3.10.0-1062.9.1.el7.x86_64
 Operating System: CentOS Linux 7 (Core)
 OSType: linux
 Architecture: x86_64
 CPUs: 1
 Total Memory: 1.795GiB
 Name: localhost.localdomain
 ID: XKOA:Z5HV:EDRR:FBWD:PVBW:5P2M:BB5N:NGN4:V7X6:BTCR:PLXE:UBIU
 Docker Root Dir: /var/lib/docker
 Debug Mode: false
 Registry: https://index.docker.io/v1/
 Labels:
 Experimental: false
 Insecure Registries:
  127.0.0.0/8
 Live Restore Enabled: false

WARNING: bridge-nf-call-iptables is disabled
WARNING: bridge-nf-call-ip6tables is disabled
```

refer
- https://bugzilla.redhat.com/show_bug.cgi?id=512206
- https://wiki.libvirt.org/page/Net.bridge.bridge-nf-call_and_sysctl.conf
- https://serverfault.com/questions/162366/iptables-bridge-and-forward-chain