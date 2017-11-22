# Instruction

Refer to https://github.com/kubernetes/kubernetes/issues/45555

## CentOS 7.2

version
```
[vagrant@openshiftdev ~]$ cat /etc/centos-release
CentOS Linux release 7.2.1511 (Core) 
```

Because google repository not arrival
```
[vagrant@openshiftdev ~]$ sudo yum --disablerepo=kubernetes --enablerepo=offline\*,kubernetes-local list | egrep  ^docker\.x86_64
docker.x86_64                           1.10.3-46.el7.centos.14  @extras        
docker.x86_64                           2:1.12.6-61.git85d7426.el7.centos
```

```
[vagrant@openshiftdev ~]$ sudo yum --disablerepo=kubernetes --enablerepo=offline\*,kubernetes-local info docker.x86_64
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: centos.ustc.edu.cn
 * epel: mirrors.ustc.edu.cn
 * extras: centos.ustc.edu.cn
 * updates: centos.ustc.edu.cn
Installed Packages
Name        : docker
Arch        : x86_64
Version     : 1.10.3
Release     : 46.el7.centos.14
Size        : 44 M
Repo        : installed
From repo   : extras
Summary     : Automates deployment of containerized applications
URL         : https://github.com/docker/docker
License     : ASL 2.0
Description : Docker is an open-source engine that automates the deployment of any
            : application as a lightweight, portable, self-sufficient container that will
            : run virtually anywhere.
            : 
            : Docker containers can encapsulate any payload, and will run consistently on
            : and between virtually any server. The same container that a developer builds
            : and tests on a laptop will run at scale, in production*, on VMs, bare-metal
            : servers, OpenStack clusters, public instances, or combinations of the above.

Available Packages
Name        : docker
Arch        : x86_64
Epoch       : 2
Version     : 1.12.6
Release     : 61.git85d7426.el7.centos
Size        : 15 M
Repo        : extras/7/x86_64
Summary     : Automates deployment of containerized applications
URL         : https://github.com/docker/docker
License     : ASL 2.0
Description : Docker is an open-source engine that automates the deployment of any
            : application as a lightweight, portable, self-sufficient container that will
            : run virtually anywhere.
            : 
            : Docker containers can encapsulate any payload, and will run consistently on
            : and between virtually any server. The same container that a developer builds
            : and tests on a laptop will run at scale, in production*, on VMs, bare-metal
            : servers, OpenStack clusters, public instances, or combinations of the above.
```

```
[vagrant@openshiftdev ~]$ sudo yum --disablerepo=kubernetes --enablerepo=offline\*,kubernetes-local upgrade -y docker
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: centos.ustc.edu.cn
 * epel: mirrors.ustc.edu.cn
 * extras: centos.ustc.edu.cn
 * updates: centos.ustc.edu.cn
Resolving Dependencies
--> Running transaction check
---> Package docker.x86_64 0:1.10.3-46.el7.centos.14 will be updated
---> Package docker.x86_64 2:1.12.6-61.git85d7426.el7.centos will be an update
--> Processing Dependency: docker-common = 2:1.12.6-61.git85d7426.el7.centos for package: 2:docker-1.12.6-61.git85d7426.el7.centos.x86_64
--> Processing Dependency: docker-client = 2:1.12.6-61.git85d7426.el7.centos for package: 2:docker-1.12.6-61.git85d7426.el7.centos.x86_64
--> Running transaction check
---> Package docker-client.x86_64 2:1.12.6-61.git85d7426.el7.centos will be installed
---> Package docker-common.x86_64 0:1.10.3-46.el7.centos.14 will be updated
---> Package docker-common.x86_64 2:1.12.6-61.git85d7426.el7.centos will be an update
--> Processing Dependency: oci-umount >= 2:2.0.0-1 for package: 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64
--> Processing Dependency: oci-systemd-hook >= 1:0.1.4-9 for package: 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64
--> Processing Dependency: oci-register-machine >= 1:0-3.10 for package: 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64
--> Processing Dependency: container-storage-setup >= 0.7.0-1 for package: 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64
--> Processing Dependency: container-selinux >= 2:2.21-2 for package: 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64
--> Processing Dependency: skopeo-containers for package: 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64
--> Running transaction check
---> Package container-selinux.noarch 2:2.28-1.git85ce147.el7 will be obsoleting
--> Processing Dependency: policycoreutils >= 2.5-11 for package: 2:container-selinux-2.28-1.git85ce147.el7.noarch
---> Package container-storage-setup.noarch 0:0.7.0-1.git4ca59c5.el7 will be installed
---> Package docker-selinux.x86_64 0:1.10.3-46.el7.centos.14 will be obsoleted
---> Package oci-register-machine.x86_64 1:0-1.8.gitaf6c129.el7 will be updated
---> Package oci-register-machine.x86_64 1:0-3.13.gitcd1e331.el7 will be an update
---> Package oci-systemd-hook.x86_64 1:0.1.4-4.git41491a3.el7 will be updated
---> Package oci-systemd-hook.x86_64 1:0.1.14-1.git1ba44c6.el7 will be an update
---> Package oci-umount.x86_64 2:2.0.0-1.git299e781.el7 will be installed
---> Package skopeo-containers.x86_64 1:0.1.24-1.dev.git28d4e08.el7 will be installed
--> Running transaction check
---> Package policycoreutils.x86_64 0:2.2.5-20.el7 will be updated
--> Processing Dependency: policycoreutils = 2.2.5-20.el7 for package: policycoreutils-python-2.2.5-20.el7.x86_64
---> Package policycoreutils.x86_64 0:2.5-17.1.el7 will be an update
--> Processing Dependency: libsepol >= 2.5-6 for package: policycoreutils-2.5-17.1.el7.x86_64
--> Processing Dependency: libselinux-utils >= 2.5-6 for package: policycoreutils-2.5-17.1.el7.x86_64
--> Processing Dependency: libsepol.so.1(LIBSEPOL_1.1)(64bit) for package: policycoreutils-2.5-17.1.el7.x86_64
--> Processing Dependency: libsepol.so.1(LIBSEPOL_1.0)(64bit) for package: policycoreutils-2.5-17.1.el7.x86_64
--> Processing Dependency: libsemanage.so.1(LIBSEMANAGE_1.1)(64bit) for package: policycoreutils-2.5-17.1.el7.x86_64
--> Running transaction check
---> Package libselinux-utils.x86_64 0:2.2.2-6.el7 will be updated
---> Package libselinux-utils.x86_64 0:2.5-11.el7 will be an update
--> Processing Dependency: libselinux(x86-64) = 2.5-11.el7 for package: libselinux-utils-2.5-11.el7.x86_64
---> Package libsemanage.x86_64 0:2.1.10-18.el7 will be updated
--> Processing Dependency: libsemanage = 2.1.10-18.el7 for package: libsemanage-python-2.1.10-18.el7.x86_64
---> Package libsemanage.x86_64 0:2.5-8.el7 will be an update
---> Package libsepol.x86_64 0:2.1.9-3.el7 will be updated
--> Processing Dependency: libsepol = 2.1.9-3.el7 for package: libsepol-devel-2.1.9-3.el7.x86_64
---> Package libsepol.x86_64 0:2.5-6.el7 will be an update
---> Package policycoreutils-python.x86_64 0:2.2.5-20.el7 will be updated
---> Package policycoreutils-python.x86_64 0:2.5-17.1.el7 will be an update
--> Processing Dependency: setools-libs >= 3.3.8-1 for package: policycoreutils-python-2.5-17.1.el7.x86_64
--> Running transaction check
---> Package libselinux.x86_64 0:2.2.2-6.el7 will be updated
--> Processing Dependency: libselinux = 2.2.2-6.el7 for package: libselinux-python-2.2.2-6.el7.x86_64
--> Processing Dependency: libselinux = 2.2.2-6.el7 for package: libselinux-devel-2.2.2-6.el7.x86_64
---> Package libselinux.x86_64 0:2.5-11.el7 will be an update
---> Package libsemanage-python.x86_64 0:2.1.10-18.el7 will be updated
---> Package libsemanage-python.x86_64 0:2.5-8.el7 will be an update
---> Package libsepol-devel.x86_64 0:2.1.9-3.el7 will be updated
---> Package libsepol-devel.x86_64 0:2.5-6.el7 will be an update
---> Package setools-libs.x86_64 0:3.3.7-46.el7 will be updated
---> Package setools-libs.x86_64 0:3.3.8-1.1.el7 will be an update
--> Running transaction check
---> Package libselinux-devel.x86_64 0:2.2.2-6.el7 will be updated
---> Package libselinux-devel.x86_64 0:2.5-11.el7 will be an update
---> Package libselinux-python.x86_64 0:2.2.2-6.el7 will be updated
---> Package libselinux-python.x86_64 0:2.5-11.el7 will be an update
--> Processing Conflict: libselinux-2.5-11.el7.x86_64 conflicts systemd < 219-20
--> Restarting Dependency Resolution with new changes.
--> Running transaction check
---> Package systemd.x86_64 0:219-19.el7_2.13 will be updated
--> Processing Dependency: systemd = 219-19.el7_2.13 for package: systemd-sysv-219-19.el7_2.13.x86_64
--> Processing Dependency: systemd = 219-19.el7_2.13 for package: systemd-devel-219-19.el7_2.13.x86_64
---> Package systemd.x86_64 0:219-42.el7_4.4 will be an update
--> Processing Dependency: systemd-libs = 219-42.el7_4.4 for package: systemd-219-42.el7_4.4.x86_64
--> Running transaction check
---> Package systemd-devel.x86_64 0:219-19.el7_2.13 will be updated
---> Package systemd-devel.x86_64 0:219-42.el7_4.4 will be an update
---> Package systemd-libs.x86_64 0:219-19.el7_2.13 will be updated
--> Processing Dependency: systemd-libs = 219-19.el7_2.13 for package: libgudev1-219-19.el7_2.13.x86_64
---> Package systemd-libs.x86_64 0:219-42.el7_4.4 will be an update
---> Package systemd-sysv.x86_64 0:219-19.el7_2.13 will be updated
---> Package systemd-sysv.x86_64 0:219-42.el7_4.4 will be an update
--> Running transaction check
---> Package libgudev1.x86_64 0:219-19.el7_2.13 will be updated
---> Package libgudev1.x86_64 0:219-42.el7_4.4 will be an update
--> Processing Conflict: systemd-219-42.el7_4.4.x86_64 conflicts dracut < 033-499
--> Restarting Dependency Resolution with new changes.
--> Running transaction check
---> Package dracut.x86_64 0:033-360.el7_2.1 will be updated
--> Processing Dependency: dracut = 033-360.el7_2.1 for package: dracut-config-rescue-033-360.el7_2.1.x86_64
--> Processing Dependency: dracut = 033-360.el7_2.1 for package: dracut-network-033-360.el7_2.1.x86_64
---> Package dracut.x86_64 0:033-502.el7 will be an update
--> Running transaction check
---> Package dracut-config-rescue.x86_64 0:033-360.el7_2.1 will be updated
---> Package dracut-config-rescue.x86_64 0:033-502.el7 will be an update
---> Package dracut-network.x86_64 0:033-360.el7_2.1 will be updated
---> Package dracut-network.x86_64 0:033-502.el7 will be an update
--> Processing Conflict: libsemanage-2.5-8.el7.x86_64 conflicts selinux-policy-base < 3.13.1-66
--> Restarting Dependency Resolution with new changes.
--> Running transaction check
---> Package selinux-policy-targeted.noarch 0:3.13.1-60.el7_2.9 will be updated
---> Package selinux-policy-targeted.noarch 0:3.13.1-166.el7_4.5 will be an update
--> Processing Dependency: selinux-policy = 3.13.1-166.el7_4.5 for package: selinux-policy-targeted-3.13.1-166.el7_4.5.noarch
--> Processing Dependency: selinux-policy = 3.13.1-166.el7_4.5 for package: selinux-policy-targeted-3.13.1-166.el7_4.5.noarch
--> Running transaction check
---> Package selinux-policy.noarch 0:3.13.1-60.el7_2.9 will be updated
---> Package selinux-policy.noarch 0:3.13.1-166.el7_4.5 will be an update
--> Finished Dependency Resolution

Dependencies Resolved

====================================================================================================================================
 Package                            Arch              Version                                      Repository                  Size
====================================================================================================================================
Installing:
 container-selinux                  noarch            2:2.28-1.git85ce147.el7                      extras                      29 k
     replacing  docker-selinux.x86_64 1.10.3-46.el7.centos.14
Updating:
 docker                             x86_64            2:1.12.6-61.git85d7426.el7.centos            extras                      15 M
 dracut                             x86_64            033-502.el7                                  base                       321 k
 selinux-policy-targeted            noarch            3.13.1-166.el7_4.5                           offline-updates            6.5 M
 systemd                            x86_64            219-42.el7_4.4                               offline-updates            5.2 M
Installing for dependencies:
 container-storage-setup            noarch            0.7.0-1.git4ca59c5.el7                       extras                      32 k
 docker-client                      x86_64            2:1.12.6-61.git85d7426.el7.centos            extras                     3.4 M
 oci-umount                         x86_64            2:2.0.0-1.git299e781.el7                     extras                      27 k
 skopeo-containers                  x86_64            1:0.1.24-1.dev.git28d4e08.el7                extras                     8.5 k
Updating for dependencies:
 docker-common                      x86_64            2:1.12.6-61.git85d7426.el7.centos            extras                      80 k
 dracut-config-rescue               x86_64            033-502.el7                                  base                        55 k
 dracut-network                     x86_64            033-502.el7                                  base                        97 k
 libgudev1                          x86_64            219-42.el7_4.4                               offline-updates             83 k
 libselinux                         x86_64            2.5-11.el7                                   base                       162 k
 libselinux-devel                   x86_64            2.5-11.el7                                   base                       186 k
 libselinux-python                  x86_64            2.5-11.el7                                   base                       234 k
 libselinux-utils                   x86_64            2.5-11.el7                                   base                       151 k
 libsemanage                        x86_64            2.5-8.el7                                    base                       145 k
 libsemanage-python                 x86_64            2.5-8.el7                                    base                       104 k
 libsepol                           x86_64            2.5-6.el7                                    base                       288 k
 libsepol-devel                     x86_64            2.5-6.el7                                    base                        74 k
 oci-register-machine               x86_64            1:0-3.13.gitcd1e331.el7                      extras                     1.1 M
 oci-systemd-hook                   x86_64            1:0.1.14-1.git1ba44c6.el7                    extras                      32 k
 policycoreutils                    x86_64            2.5-17.1.el7                                 base                       858 k
 policycoreutils-python             x86_64            2.5-17.1.el7                                 base                       446 k
 selinux-policy                     noarch            3.13.1-166.el7_4.5                           offline-updates            436 k
 setools-libs                       x86_64            3.3.8-1.1.el7                                base                       612 k
 systemd-devel                      x86_64            219-42.el7_4.4                               offline-updates            186 k
 systemd-libs                       x86_64            219-42.el7_4.4                               offline-updates            376 k
 systemd-sysv                       x86_64            219-42.el7_4.4                               offline-updates             70 k

Transaction Summary
====================================================================================================================================
Install  1 Package  (+ 4 Dependent packages)
Upgrade  4 Packages (+21 Dependent packages)

Total download size: 36 M
Downloading packages:
No Presto metadata available for base
extras/7/x86_64/prestodelta                                                                                  |  51 kB  00:00:00     
(1/23): container-selinux-2.28-1.git85ce147.el7.noarch.rpm                                                   |  29 kB  00:00:00     
(2/23): container-storage-setup-0.7.0-1.git4ca59c5.el7.noarch.rpm                                            |  32 kB  00:00:00     
(3/23): docker-common-1.12.6-61.git85d7426.el7.centos.x86_64.rpm                                             |  80 kB  00:00:00     
(4/23): dracut-033-502.el7.x86_64.rpm                                                                        | 321 kB  00:00:01     
(5/23): docker-client-1.12.6-61.git85d7426.el7.centos.x86_64.rpm                                             | 3.4 MB  00:00:01     
(6/23): libselinux-2.5-11.el7.x86_64.rpm                                                                     | 162 kB  00:00:00     
(7/23): dracut-config-rescue-033-502.el7.x86_64.rpm                                                          |  55 kB  00:00:02     
(8/23): dracut-network-033-502.el7.x86_64.rpm                                                                |  97 kB  00:00:02     
(9/23): libselinux-python-2.5-11.el7.x86_64.rpm                                                              | 234 kB  00:00:01     
(10/23): libselinux-devel-2.5-11.el7.x86_64.rpm                                                              | 186 kB  00:00:01     
(11/23): libsemanage-2.5-8.el7.x86_64.rpm                                                                    | 145 kB  00:00:00     
(12/23): libsemanage-python-2.5-8.el7.x86_64.rpm                                                             | 104 kB  00:00:00     
(13/23): libsepol-devel-2.5-6.el7.x86_64.rpm                                                                 |  74 kB  00:00:00     
(14/23): libselinux-utils-2.5-11.el7.x86_64.rpm                                                              | 151 kB  00:00:02     
(15/23): libsepol-2.5-6.el7.x86_64.rpm                                                                       | 288 kB  00:00:00     
(16/23): oci-register-machine-0-3.13.gitcd1e331.el7.x86_64.rpm                                               | 1.1 MB  00:00:01     
(17/23): oci-systemd-hook-0.1.14-1.git1ba44c6.el7.x86_64.rpm                                                 |  32 kB  00:00:00     
(18/23): oci-umount-2.0.0-1.git299e781.el7.x86_64.rpm                                                        |  27 kB  00:00:00     
(19/23): skopeo-containers-0.1.24-1.dev.git28d4e08.el7.x86_64.rpm                                            | 8.5 kB  00:00:00     
(20/23): policycoreutils-python-2.5-17.1.el7.x86_64.rpm                                                      | 446 kB  00:00:00     
(21/23): policycoreutils-2.5-17.1.el7.x86_64.rpm                                                             | 858 kB  00:00:02     
(22/23): setools-libs-3.3.8-1.1.el7.x86_64.rpm                                                               | 612 kB  00:00:04     
(23/23): docker-1.12.6-61.git85d7426.el7.centos.x86_64.rpm                                                   |  15 MB  00:02:13     
------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                               274 kB/s |  36 MB  00:02:14     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Updating   : libsepol-2.5-6.el7.x86_64                                                                                       1/56 
  Updating   : libselinux-2.5-11.el7.x86_64                                                                                    2/56 
  Updating   : systemd-libs-219-42.el7_4.4.x86_64                                                                              3/56 
  Updating   : systemd-219-42.el7_4.4.x86_64                                                                                   4/56 
warning: /etc/systemd/system.conf created as /etc/systemd/system.conf.rpmnew
  Updating   : libsemanage-2.5-8.el7.x86_64                                                                                    5/56 
  Updating   : dracut-033-502.el7.x86_64                                                                                       6/56 
  Updating   : libselinux-utils-2.5-11.el7.x86_64                                                                              7/56 
  Updating   : policycoreutils-2.5-17.1.el7.x86_64                                                                             8/56 
  Updating   : selinux-policy-3.13.1-166.el7_4.5.noarch                                                                        9/56 
  Updating   : selinux-policy-targeted-3.13.1-166.el7_4.5.noarch                                                              10/56 
warning: /etc/selinux/targeted/seusers created as /etc/selinux/targeted/seusers.rpmnew
‘/etc/selinux/targeted/modules/active/seusers’ -> ‘/etc/selinux/targeted/active/seusers.local’
  Updating   : libsemanage-python-2.5-8.el7.x86_64                                                                            11/56 
  Updating   : 1:oci-register-machine-0-3.13.gitcd1e331.el7.x86_64                                                            12/56 
  Installing : 2:oci-umount-2.0.0-1.git299e781.el7.x86_64                                                                     13/56 
  Updating   : 1:oci-systemd-hook-0.1.14-1.git1ba44c6.el7.x86_64                                                              14/56 
  Updating   : libselinux-python-2.5-11.el7.x86_64                                                                            15/56 
  Updating   : setools-libs-3.3.8-1.1.el7.x86_64                                                                              16/56 
  Updating   : policycoreutils-python-2.5-17.1.el7.x86_64                                                                     17/56 
  Installing : 2:container-selinux-2.28-1.git85ce147.el7.noarch                                                               18/56 
  Updating   : libsepol-devel-2.5-6.el7.x86_64                                                                                19/56 
  Installing : 1:skopeo-containers-0.1.24-1.dev.git28d4e08.el7.x86_64                                                         20/56 
  Installing : container-storage-setup-0.7.0-1.git4ca59c5.el7.noarch                                                          21/56 
  Updating   : 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64                                                         22/56 
warning: /etc/sysconfig/docker created as /etc/sysconfig/docker.rpmnew
  Installing : 2:docker-client-1.12.6-61.git85d7426.el7.centos.x86_64                                                         23/56 
  Updating   : 2:docker-1.12.6-61.git85d7426.el7.centos.x86_64                                                                24/56 
  Updating   : libselinux-devel-2.5-11.el7.x86_64                                                                             25/56 
  Updating   : dracut-config-rescue-033-502.el7.x86_64                                                                        26/56 
  Updating   : dracut-network-033-502.el7.x86_64                                                                              27/56 
  Updating   : systemd-devel-219-42.el7_4.4.x86_64                                                                            28/56 
  Updating   : systemd-sysv-219-42.el7_4.4.x86_64                                                                             29/56 
  Updating   : libgudev1-219-42.el7_4.4.x86_64                                                                                30/56 
  Cleanup    : libselinux-devel-2.2.2-6.el7.x86_64                                                                            31/56 
  Cleanup    : systemd-devel-219-19.el7_2.13.x86_64                                                                           32/56 
  Cleanup    : docker-1.10.3-46.el7.centos.14.x86_64                                                                          33/56 
  Erasing    : docker-selinux-1.10.3-46.el7.centos.14.x86_64                                                                  34/56 
  Cleanup    : selinux-policy-targeted-3.13.1-60.el7_2.9.noarch                                                               35/56 
warning: /etc/selinux/targeted/modules/active/seusers.final saved as /etc/selinux/targeted/modules/active/seusers.final.rpmsave
  Cleanup    : selinux-policy-3.13.1-60.el7_2.9.noarch                                                                        36/56 
  Cleanup    : libsepol-devel-2.1.9-3.el7.x86_64                                                                              37/56 
  Cleanup    : dracut-network-033-360.el7_2.1.x86_64                                                                          38/56 
  Cleanup    : systemd-sysv-219-19.el7_2.13.x86_64                                                                            39/56 
  Cleanup    : dracut-config-rescue-033-360.el7_2.1.x86_64                                                                    40/56 
  Cleanup    : policycoreutils-python-2.2.5-20.el7.x86_64                                                                     41/56 
  Cleanup    : policycoreutils-2.2.5-20.el7.x86_64                                                                            42/56 
  Cleanup    : setools-libs-3.3.7-46.el7.x86_64                                                                               43/56 
  Cleanup    : libselinux-utils-2.2.2-6.el7.x86_64                                                                            44/56 
  Cleanup    : libselinux-python-2.2.2-6.el7.x86_64                                                                           45/56 
  Cleanup    : libsemanage-python-2.1.10-18.el7.x86_64                                                                        46/56 
  Cleanup    : libsemanage-2.1.10-18.el7.x86_64                                                                               47/56 
  Cleanup    : dracut-033-360.el7_2.1.x86_64                                                                                  48/56 
  Cleanup    : 1:oci-register-machine-0-1.8.gitaf6c129.el7.x86_64                                                             49/56 
  Cleanup    : systemd-219-19.el7_2.13.x86_64                                                                                 50/56 
  Cleanup    : 1:oci-systemd-hook-0.1.4-4.git41491a3.el7.x86_64                                                               51/56 
  Cleanup    : libgudev1-219-19.el7_2.13.x86_64                                                                               52/56 
  Cleanup    : docker-common-1.10.3-46.el7.centos.14.x86_64                                                                   53/56 
  Cleanup    : systemd-libs-219-19.el7_2.13.x86_64                                                                            54/56 
  Cleanup    : libselinux-2.2.2-6.el7.x86_64                                                                                  55/56 
  Cleanup    : libsepol-2.1.9-3.el7.x86_64                                                                                    56/56 
  Verifying  : libselinux-utils-2.5-11.el7.x86_64                                                                              1/56 
  Verifying  : systemd-devel-219-42.el7_4.4.x86_64                                                                             2/56 
  Verifying  : selinux-policy-3.13.1-166.el7_4.5.noarch                                                                        3/56 
  Verifying  : libselinux-2.5-11.el7.x86_64                                                                                    4/56 
  Verifying  : libsepol-2.5-6.el7.x86_64                                                                                       5/56 
  Verifying  : libsepol-devel-2.5-6.el7.x86_64                                                                                 6/56 
  Verifying  : policycoreutils-python-2.5-17.1.el7.x86_64                                                                      7/56 
  Verifying  : 2:container-selinux-2.28-1.git85ce147.el7.noarch                                                                8/56 
  Verifying  : 1:oci-register-machine-0-3.13.gitcd1e331.el7.x86_64                                                             9/56 
  Verifying  : 2:docker-1.12.6-61.git85d7426.el7.centos.x86_64                                                                10/56 
  Verifying  : 2:docker-common-1.12.6-61.git85d7426.el7.centos.x86_64                                                         11/56 
  Verifying  : 2:oci-umount-2.0.0-1.git299e781.el7.x86_64                                                                     12/56 
  Verifying  : selinux-policy-targeted-3.13.1-166.el7_4.5.noarch                                                              13/56 
  Verifying  : dracut-config-rescue-033-502.el7.x86_64                                                                        14/56 
  Verifying  : policycoreutils-2.5-17.1.el7.x86_64                                                                            15/56 
  Verifying  : libsemanage-python-2.5-8.el7.x86_64                                                                            16/56 
  Verifying  : 1:oci-systemd-hook-0.1.14-1.git1ba44c6.el7.x86_64                                                              17/56 
  Verifying  : libselinux-devel-2.5-11.el7.x86_64                                                                             18/56 
  Verifying  : systemd-219-42.el7_4.4.x86_64                                                                                  19/56 
  Verifying  : container-storage-setup-0.7.0-1.git4ca59c5.el7.noarch                                                          20/56 
  Verifying  : libgudev1-219-42.el7_4.4.x86_64                                                                                21/56 
  Verifying  : systemd-sysv-219-42.el7_4.4.x86_64                                                                             22/56 
  Verifying  : systemd-libs-219-42.el7_4.4.x86_64                                                                             23/56 
  Verifying  : libselinux-python-2.5-11.el7.x86_64                                                                            24/56 
  Verifying  : 1:skopeo-containers-0.1.24-1.dev.git28d4e08.el7.x86_64                                                         25/56 
  Verifying  : dracut-network-033-502.el7.x86_64                                                                              26/56 
  Verifying  : dracut-033-502.el7.x86_64                                                                                      27/56 
  Verifying  : setools-libs-3.3.8-1.1.el7.x86_64                                                                              28/56 
  Verifying  : 2:docker-client-1.12.6-61.git85d7426.el7.centos.x86_64                                                         29/56 
  Verifying  : libsemanage-2.5-8.el7.x86_64                                                                                   30/56 
  Verifying  : dracut-config-rescue-033-360.el7_2.1.x86_64                                                                    31/56 
  Verifying  : docker-1.10.3-46.el7.centos.14.x86_64                                                                          32/56 
  Verifying  : libsepol-devel-2.1.9-3.el7.x86_64                                                                              33/56 
  Verifying  : selinux-policy-3.13.1-60.el7_2.9.noarch                                                                        34/56 
  Verifying  : systemd-devel-219-19.el7_2.13.x86_64                                                                           35/56 
  Verifying  : libselinux-devel-2.2.2-6.el7.x86_64                                                                            36/56 
  Verifying  : setools-libs-3.3.7-46.el7.x86_64                                                                               37/56 
  Verifying  : dracut-network-033-360.el7_2.1.x86_64                                                                          38/56 
  Verifying  : libgudev1-219-19.el7_2.13.x86_64                                                                               39/56 
  Verifying  : selinux-policy-targeted-3.13.1-60.el7_2.9.noarch                                                               40/56 
  Verifying  : dracut-033-360.el7_2.1.x86_64                                                                                  41/56 
  Verifying  : libsemanage-python-2.1.10-18.el7.x86_64                                                                        42/56 
  Verifying  : libselinux-utils-2.2.2-6.el7.x86_64                                                                            43/56 
  Verifying  : 1:oci-systemd-hook-0.1.4-4.git41491a3.el7.x86_64                                                               44/56 
  Verifying  : systemd-libs-219-19.el7_2.13.x86_64                                                                            45/56 
  Verifying  : systemd-sysv-219-19.el7_2.13.x86_64                                                                            46/56 
  Verifying  : libsepol-2.1.9-3.el7.x86_64                                                                                    47/56 
  Verifying  : libsemanage-2.1.10-18.el7.x86_64                                                                               48/56 
  Verifying  : 1:oci-register-machine-0-1.8.gitaf6c129.el7.x86_64                                                             49/56 
  Verifying  : docker-selinux-1.10.3-46.el7.centos.14.x86_64                                                                  50/56 
  Verifying  : libselinux-python-2.2.2-6.el7.x86_64                                                                           51/56 
  Verifying  : systemd-219-19.el7_2.13.x86_64                                                                                 52/56 
  Verifying  : policycoreutils-2.2.5-20.el7.x86_64                                                                            53/56 
  Verifying  : docker-common-1.10.3-46.el7.centos.14.x86_64                                                                   54/56 
  Verifying  : libselinux-2.2.2-6.el7.x86_64                                                                                  55/56 
  Verifying  : policycoreutils-python-2.2.5-20.el7.x86_64                                                                     56/56 

Installed:
  container-selinux.noarch 2:2.28-1.git85ce147.el7                                                                                  

Dependency Installed:
  container-storage-setup.noarch 0:0.7.0-1.git4ca59c5.el7           docker-client.x86_64 2:1.12.6-61.git85d7426.el7.centos          
  oci-umount.x86_64 2:2.0.0-1.git299e781.el7                        skopeo-containers.x86_64 1:0.1.24-1.dev.git28d4e08.el7          

Updated:
  docker.x86_64 2:1.12.6-61.git85d7426.el7.centos  dracut.x86_64 0:033-502.el7  selinux-policy-targeted.noarch 0:3.13.1-166.el7_4.5 
  systemd.x86_64 0:219-42.el7_4.4                 

Dependency Updated:
  docker-common.x86_64 2:1.12.6-61.git85d7426.el7.centos              dracut-config-rescue.x86_64 0:033-502.el7                     
  dracut-network.x86_64 0:033-502.el7                                 libgudev1.x86_64 0:219-42.el7_4.4                             
  libselinux.x86_64 0:2.5-11.el7                                      libselinux-devel.x86_64 0:2.5-11.el7                          
  libselinux-python.x86_64 0:2.5-11.el7                               libselinux-utils.x86_64 0:2.5-11.el7                          
  libsemanage.x86_64 0:2.5-8.el7                                      libsemanage-python.x86_64 0:2.5-8.el7                         
  libsepol.x86_64 0:2.5-6.el7                                         libsepol-devel.x86_64 0:2.5-6.el7                             
  oci-register-machine.x86_64 1:0-3.13.gitcd1e331.el7                 oci-systemd-hook.x86_64 1:0.1.14-1.git1ba44c6.el7             
  policycoreutils.x86_64 0:2.5-17.1.el7                               policycoreutils-python.x86_64 0:2.5-17.1.el7                  
  selinux-policy.noarch 0:3.13.1-166.el7_4.5                          setools-libs.x86_64 0:3.3.8-1.1.el7                           
  systemd-devel.x86_64 0:219-42.el7_4.4                               systemd-libs.x86_64 0:219-42.el7_4.4                          
  systemd-sysv.x86_64 0:219-42.el7_4.4                               

Replaced:
  docker-selinux.x86_64 0:1.10.3-46.el7.centos.14                                                                                   

Complete!
```