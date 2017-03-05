

[root@localhost tmp]# mkdir brick1
[root@localhost tmp]# ls -d brick1/
brick1/

http://gluster.readthedocs.io/en/latest/Install-Guide/Install/#for-red-hatcentos

https://wiki.centos.org/SpecialInterestGroup/Storage/gluster-Quickstart

[root@localhost tmp]# yum --disablerepo=\* --enablerepo=offline-gluster38 --enablerepo=offline --enablerepo=c7-media list | grep glusterfs
glusterfs.x86_64                           3.7.9-12.el7.centos         @c7-media
glusterfs-client-xlators.x86_64            3.7.9-12.el7.centos         @c7-media
glusterfs-fuse.x86_64                      3.7.9-12.el7.centos         @c7-media
glusterfs-libs.x86_64                      3.7.9-12.el7.centos         @c7-media
glusterfs.x86_64                           3.8.8-1.el7                 offline-gluster38
glusterfs-api.x86_64                       3.8.8-1.el7                 offline-gluster38
glusterfs-api-devel.x86_64                 3.8.8-1.el7                 offline-gluster38
glusterfs-cli.x86_64                       3.8.8-1.el7                 offline-gluster38
glusterfs-client-xlators.x86_64            3.8.8-1.el7                 offline-gluster38
glusterfs-coreutils.x86_64                 0.2.0-1.el7                 offline-gluster38
glusterfs-devel.x86_64                     3.8.8-1.el7                 offline-gluster38
glusterfs-extra-xlators.x86_64             3.8.8-1.el7                 offline-gluster38
glusterfs-fuse.x86_64                      3.8.8-1.el7                 offline-gluster38
glusterfs-ganesha.x86_64                   3.8.8-1.el7                 offline-gluster38
glusterfs-geo-replication.x86_64           3.8.8-1.el7                 offline-gluster38
glusterfs-libs.x86_64                      3.8.8-1.el7                 offline-gluster38
glusterfs-rdma.x86_64                      3.8.8-1.el7                 offline-gluster38
glusterfs-resource-agents.noarch           3.8.8-1.el7                 offline-gluster38
glusterfs-server.x86_64                    3.8.8-1.el7                 offline-gluster38
samba-vfs-glusterfs.x86_64                 4.4.4-9.el7                 c7-media 


[root@localhost opt]# yum --disablerep=\* --enablerepo=c7-media --enablerepo=offline install -y centos-release-gluster38
已加载插件：fastestmirror
file:///home/ecp/centos-repo-mirror/cloud/x86_64/repodata/repomd.xml: [Errno 14] curl#37 - "Couldn't open file /home/ecp/centos-repo-mirror/cloud/x86_64/repodata/repomd.xml"
正在尝试其它镜像。
offline                                                                                            | 3.4 kB  00:00:00     
Loading mirror speeds from cached hostfile
 * c7-media: 
 * offline: 
正在解决依赖关系
--> 正在检查事务
---> 软件包 centos-release-gluster38.noarch.0.1.0-1.el7.centos 将被 安装
--> 正在处理依赖关系 centos-release-storage-common，它被软件包 centos-release-gluster38-1.0-1.el7.centos.noarch 需要
--> 正在检查事务
---> 软件包 centos-release-storage-common.noarch.0.1-2.el7.centos 将被 安装
--> 解决依赖关系完成

依赖关系解决

==========================================================================================================================
 Package                                    架构                版本                           源                    大小
==========================================================================================================================
正在安装:
 centos-release-gluster38                   noarch              1.0-1.el7.centos               offline              4.1 k
为依赖而安装:
 centos-release-storage-common              noarch              1-2.el7.centos                 offline              4.5 k

事务概要
==========================================================================================================================
安装  1 软件包 (+1 依赖软件包)

总下载量：8.6 k
安装大小：1.6 k
Downloading packages:
--------------------------------------------------------------------------------------------------------------------------
总计                                                                                      329 kB/s | 8.6 kB  00:00:00     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  正在安装    : centos-release-storage-common-1-2.el7.centos.noarch                                                   1/2 
  正在安装    : centos-release-gluster38-1.0-1.el7.centos.noarch                                                      2/2 
  验证中      : centos-release-storage-common-1-2.el7.centos.noarch                                                   1/2 
  验证中      : centos-release-gluster38-1.0-1.el7.centos.noarch                                                      2/2 

已安装:
  centos-release-gluster38.noarch 0:1.0-1.el7.centos                                                                      

作为依赖被安装:
  centos-release-storage-common.noarch 0:1-2.el7.centos                                                                   

完毕！


[root@localhost opt]# yum --disablerep=\* --enablerepo=c7-media --enablerepo=offline install -y glusterfs glusterfs-fuse
已加载插件：fastestmirror
Loading mirror speeds from cached hostfile
 * c7-media: 
 * offline: 
正在解决依赖关系
--> 正在检查事务
---> 软件包 glusterfs.x86_64.0.3.7.9-12.el7.centos 将被 安装
--> 正在处理依赖关系 glusterfs-libs(x86-64) = 3.7.9-12.el7.centos，它被软件包 glusterfs-3.7.9-12.el7.centos.x86_64 需要
--> 正在处理依赖关系 libglusterfs.so.0()(64bit)，它被软件包 glusterfs-3.7.9-12.el7.centos.x86_64 需要
--> 正在处理依赖关系 libgfxdr.so.0()(64bit)，它被软件包 glusterfs-3.7.9-12.el7.centos.x86_64 需要
--> 正在处理依赖关系 libgfrpc.so.0()(64bit)，它被软件包 glusterfs-3.7.9-12.el7.centos.x86_64 需要
---> 软件包 glusterfs-fuse.x86_64.0.3.7.9-12.el7.centos 将被 安装
--> 正在处理依赖关系 glusterfs-client-xlators(x86-64) = 3.7.9-12.el7.centos，它被软件包 glusterfs-fuse-3.7.9-12.el7.centos.x86_64 需要
--> 正在处理依赖关系 attr，它被软件包 glusterfs-fuse-3.7.9-12.el7.centos.x86_64 需要
--> 正在检查事务
---> 软件包 attr.x86_64.0.2.4.46-12.el7 将被 安装
---> 软件包 glusterfs-client-xlators.x86_64.0.3.7.9-12.el7.centos 将被 安装
---> 软件包 glusterfs-libs.x86_64.0.3.7.9-12.el7.centos 将被 安装
--> 解决依赖关系完成

依赖关系解决

==========================================================================================================================
 Package                               架构                版本                               源                     大小
==========================================================================================================================
正在安装:
 glusterfs                             x86_64              3.7.9-12.el7.centos                c7-media              462 k
 glusterfs-fuse                        x86_64              3.7.9-12.el7.centos                c7-media              109 k
为依赖而安装:
 attr                                  x86_64              2.4.46-12.el7                      c7-media               66 k
 glusterfs-client-xlators              x86_64              3.7.9-12.el7.centos                c7-media              837 k
 glusterfs-libs                        x86_64              3.7.9-12.el7.centos                c7-media              331 k

事务概要
==========================================================================================================================
安装  2 软件包 (+3 依赖软件包)

总下载量：1.8 M
安装大小：7.1 M
Downloading packages:
--------------------------------------------------------------------------------------------------------------------------
总计                                                                                      8.8 MB/s | 1.8 MB  00:00:00     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  正在安装    : glusterfs-libs-3.7.9-12.el7.centos.x86_64                                                             1/5 
  正在安装    : glusterfs-client-xlators-3.7.9-12.el7.centos.x86_64                                                   2/5 
  正在安装    : glusterfs-3.7.9-12.el7.centos.x86_64                                                                  3/5 
  正在安装    : attr-2.4.46-12.el7.x86_64                                                                             4/5 
  正在安装    : glusterfs-fuse-3.7.9-12.el7.centos.x86_64                                                             5/5 
  验证中      : glusterfs-client-xlators-3.7.9-12.el7.centos.x86_64                                                   1/5 
  验证中      : glusterfs-libs-3.7.9-12.el7.centos.x86_64                                                             2/5 
  验证中      : glusterfs-fuse-3.7.9-12.el7.centos.x86_64                                                             3/5 
  验证中      : attr-2.4.46-12.el7.x86_64                                                                             4/5 
  验证中      : glusterfs-3.7.9-12.el7.centos.x86_64                                                                  5/5 

已安装:
  glusterfs.x86_64 0:3.7.9-12.el7.centos                    glusterfs-fuse.x86_64 0:3.7.9-12.el7.centos                   

作为依赖被安装:
  attr.x86_64 0:2.4.46-12.el7                            glusterfs-client-xlators.x86_64 0:3.7.9-12.el7.centos           
  glusterfs-libs.x86_64 0:3.7.9-12.el7.centos           

完毕！


[root@localhost yum.repos.d]# yum --disablerepo=\* --enablerepo=offline-gluster38 --enablerepo=offline --enablerepo=c7-media install -y glusterfs glusterfs-server glusterfs-fuse
已加载插件：fastestmirror
file:///home/ecp/centos-repo-mirror/cloud/x86_64/repodata/repomd.xml: [Errno 14] curl#37 - "Couldn't open file /home/ecp/centos-repo-mirror/cloud/x86_64/repodata/repomd.xml"
正在尝试其它镜像。
offline                                                                                            | 3.4 kB  00:00:00     
offline-gluster38                                                                                  | 2.9 kB  00:00:00     
Loading mirror speeds from cached hostfile
 * c7-media: 
 * offline: 
正在解决依赖关系
--> 正在检查事务
---> 软件包 glusterfs.x86_64.0.3.7.9-12.el7.centos 将被 升级
---> 软件包 glusterfs.x86_64.0.3.8.8-1.el7 将被 更新
--> 正在处理依赖关系 glusterfs-libs(x86-64) = 3.8.8-1.el7，它被软件包 glusterfs-3.8.8-1.el7.x86_64 需要
---> 软件包 glusterfs-fuse.x86_64.0.3.7.9-12.el7.centos 将被 升级
---> 软件包 glusterfs-fuse.x86_64.0.3.8.8-1.el7 将被 更新
--> 正在处理依赖关系 glusterfs-client-xlators(x86-64) = 3.8.8-1.el7，它被软件包 glusterfs-fuse-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 psmisc，它被软件包 glusterfs-fuse-3.8.8-1.el7.x86_64 需要
---> 软件包 glusterfs-server.x86_64.0.3.8.8-1.el7 将被 安装
--> 正在处理依赖关系 glusterfs-cli(x86-64) = 3.8.8-1.el7，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 glusterfs-api(x86-64) = 3.8.8-1.el7，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 rpcbind，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 nfs-utils，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 libgfapi.so.0(GFAPI_PRIVATE_3.7.0)(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 libgfapi.so.0(GFAPI_PRIVATE_3.4.0)(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 libgfapi.so.0(GFAPI_3.7.4)(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 libgfapi.so.0(GFAPI_3.7.0)(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 libgfapi.so.0(GFAPI_3.6.0)(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 libgfapi.so.0(GFAPI_3.5.1)(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 libgfapi.so.0(GFAPI_3.4.2)(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 libgfapi.so.0(GFAPI_3.4.0)(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 liburcu-cds.so.1()(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 liburcu-bp.so.1()(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 libgfapi.so.0()(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在检查事务
---> 软件包 glusterfs-api.x86_64.0.3.8.8-1.el7 将被 安装
---> 软件包 glusterfs-cli.x86_64.0.3.8.8-1.el7 将被 安装
---> 软件包 glusterfs-client-xlators.x86_64.0.3.7.9-12.el7.centos 将被 升级
---> 软件包 glusterfs-client-xlators.x86_64.0.3.8.8-1.el7 将被 更新
---> 软件包 glusterfs-libs.x86_64.0.3.7.9-12.el7.centos 将被 升级
---> 软件包 glusterfs-libs.x86_64.0.3.8.8-1.el7 将被 更新
---> 软件包 nfs-utils.x86_64.1.1.3.0-0.33.el7 将被 安装
--> 正在处理依赖关系 libtirpc >= 0.2.4-0.7，它被软件包 1:nfs-utils-1.3.0-0.33.el7.x86_64 需要
--> 正在处理依赖关系 gssproxy >= 0.3.0-0，它被软件包 1:nfs-utils-1.3.0-0.33.el7.x86_64 需要
--> 正在处理依赖关系 quota，它被软件包 1:nfs-utils-1.3.0-0.33.el7.x86_64 需要
--> 正在处理依赖关系 libnfsidmap，它被软件包 1:nfs-utils-1.3.0-0.33.el7.x86_64 需要
--> 正在处理依赖关系 libevent，它被软件包 1:nfs-utils-1.3.0-0.33.el7.x86_64 需要
--> 正在处理依赖关系 keyutils，它被软件包 1:nfs-utils-1.3.0-0.33.el7.x86_64 需要
--> 正在处理依赖关系 libtirpc.so.1()(64bit)，它被软件包 1:nfs-utils-1.3.0-0.33.el7.x86_64 需要
--> 正在处理依赖关系 libnfsidmap.so.0()(64bit)，它被软件包 1:nfs-utils-1.3.0-0.33.el7.x86_64 需要
--> 正在处理依赖关系 libevent-2.0.so.5()(64bit)，它被软件包 1:nfs-utils-1.3.0-0.33.el7.x86_64 需要
---> 软件包 psmisc.x86_64.0.22.20-11.el7 将被 安装
---> 软件包 rpcbind.x86_64.0.0.2.0-38.el7 将被 安装
---> 软件包 userspace-rcu.x86_64.0.0.7.16-1.el7 将被 安装
--> 正在检查事务
---> 软件包 gssproxy.x86_64.0.0.4.1-13.el7 将被 安装
--> 正在处理依赖关系 libverto-tevent，它被软件包 gssproxy-0.4.1-13.el7.x86_64 需要
--> 正在处理依赖关系 libini_config.so.3(INI_CONFIG_1.1.0)(64bit)，它被软件包 gssproxy-0.4.1-13.el7.x86_64 需要
--> 正在处理依赖关系 libref_array.so.1()(64bit)，它被软件包 gssproxy-0.4.1-13.el7.x86_64 需要
--> 正在处理依赖关系 libini_config.so.3()(64bit)，它被软件包 gssproxy-0.4.1-13.el7.x86_64 需要
--> 正在处理依赖关系 libcollection.so.2()(64bit)，它被软件包 gssproxy-0.4.1-13.el7.x86_64 需要
--> 正在处理依赖关系 libbasicobjects.so.0()(64bit)，它被软件包 gssproxy-0.4.1-13.el7.x86_64 需要
---> 软件包 keyutils.x86_64.0.1.5.8-3.el7 将被 安装
---> 软件包 libevent.x86_64.0.2.0.21-4.el7 将被 安装
---> 软件包 libnfsidmap.x86_64.0.0.25-15.el7 将被 安装
---> 软件包 libtirpc.x86_64.0.0.2.4-0.8.el7 将被 安装
---> 软件包 quota.x86_64.1.4.01-14.el7 将被 安装
--> 正在处理依赖关系 quota-nls = 1:4.01-14.el7，它被软件包 1:quota-4.01-14.el7.x86_64 需要
--> 正在处理依赖关系 tcp_wrappers，它被软件包 1:quota-4.01-14.el7.x86_64 需要
--> 正在检查事务
---> 软件包 libbasicobjects.x86_64.0.0.1.1-27.el7 将被 安装
---> 软件包 libcollection.x86_64.0.0.6.2-27.el7 将被 安装
---> 软件包 libini_config.x86_64.0.1.3.0-27.el7 将被 安装
--> 正在处理依赖关系 libpath_utils.so.1(PATH_UTILS_0.2.1)(64bit)，它被软件包 libini_config-1.3.0-27.el7.x86_64 需要
--> 正在处理依赖关系 libpath_utils.so.1()(64bit)，它被软件包 libini_config-1.3.0-27.el7.x86_64 需要
---> 软件包 libref_array.x86_64.0.0.1.5-27.el7 将被 安装
---> 软件包 libverto-tevent.x86_64.0.0.2.5-4.el7 将被 安装
--> 正在处理依赖关系 libtevent.so.0(TEVENT_0.9.9)(64bit)，它被软件包 libverto-tevent-0.2.5-4.el7.x86_64 需要
--> 正在处理依赖关系 libtalloc.so.2(TALLOC_2.0.2)(64bit)，它被软件包 libverto-tevent-0.2.5-4.el7.x86_64 需要
--> 正在处理依赖关系 libtevent.so.0()(64bit)，它被软件包 libverto-tevent-0.2.5-4.el7.x86_64 需要
--> 正在处理依赖关系 libtalloc.so.2()(64bit)，它被软件包 libverto-tevent-0.2.5-4.el7.x86_64 需要
---> 软件包 quota-nls.noarch.1.4.01-14.el7 将被 安装
---> 软件包 tcp_wrappers.x86_64.0.7.6-77.el7 将被 安装
--> 正在检查事务
---> 软件包 libpath_utils.x86_64.0.0.2.1-27.el7 将被 安装
---> 软件包 libtalloc.x86_64.0.2.1.6-1.el7 将被 安装
---> 软件包 libtevent.x86_64.0.0.9.28-1.el7 将被 安装
--> 解决依赖关系完成

依赖关系解决

==========================================================================================================================
 Package                              架构               版本                         源                             大小
==========================================================================================================================
正在安装:
 glusterfs-server                     x86_64             3.8.8-1.el7                  offline-gluster38             1.4 M
正在更新:
 glusterfs                            x86_64             3.8.8-1.el7                  offline-gluster38             510 k
 glusterfs-fuse                       x86_64             3.8.8-1.el7                  offline-gluster38             133 k
为依赖而安装:
 glusterfs-api                        x86_64             3.8.8-1.el7                  offline-gluster38              90 k
 glusterfs-cli                        x86_64             3.8.8-1.el7                  offline-gluster38             183 k
 gssproxy                             x86_64             0.4.1-13.el7                 c7-media                       87 k
 keyutils                             x86_64             1.5.8-3.el7                  c7-media                       54 k
 libbasicobjects                      x86_64             0.1.1-27.el7                 c7-media                       25 k
 libcollection                        x86_64             0.6.2-27.el7                 c7-media                       41 k
 libevent                             x86_64             2.0.21-4.el7                 c7-media                      214 k
 libini_config                        x86_64             1.3.0-27.el7                 c7-media                       63 k
 libnfsidmap                          x86_64             0.25-15.el7                  c7-media                       47 k
 libpath_utils                        x86_64             0.2.1-27.el7                 c7-media                       27 k
 libref_array                         x86_64             0.1.5-27.el7                 c7-media                       26 k
 libtalloc                            x86_64             2.1.6-1.el7                  c7-media                       34 k
 libtevent                            x86_64             0.9.28-1.el7                 c7-media                       34 k
 libtirpc                             x86_64             0.2.4-0.8.el7                c7-media                       88 k
 libverto-tevent                      x86_64             0.2.5-4.el7                  c7-media                      9.0 k
 nfs-utils                            x86_64             1:1.3.0-0.33.el7             c7-media                      377 k
 psmisc                               x86_64             22.20-11.el7                 c7-media                      141 k
 quota                                x86_64             1:4.01-14.el7                c7-media                      179 k
 quota-nls                            noarch             1:4.01-14.el7                c7-media                       90 k
 rpcbind                              x86_64             0.2.0-38.el7                 c7-media                       59 k
 tcp_wrappers                         x86_64             7.6-77.el7                   c7-media                       78 k
 userspace-rcu                        x86_64             0.7.16-1.el7                 offline-gluster38              72 k
为依赖而更新:
 glusterfs-client-xlators             x86_64             3.8.8-1.el7                  offline-gluster38             781 k
 glusterfs-libs                       x86_64             3.8.8-1.el7                  offline-gluster38             379 k

事务概要
==========================================================================================================================
安装  1 软件包 (+22 依赖软件包)
升级  2 软件包 (+ 2 依赖软件包)

总下载量：5.1 M
Downloading packages:
警告：/home/ecp/centos-repo-mirror/storage/x86_64/gluster-3.8/glusterfs-3.8.8-1.el7.x86_64.rpm: 头V4 RSA/SHA1 Signature, 密钥 ID e451e5b5: NOKEY
glusterfs-3.8.8-1.el7.x86_64.rpm 的公钥尚未安装
--------------------------------------------------------------------------------------------------------------------------
总计                                                                                      8.0 MB/s | 5.1 MB  00:00:00     
从 file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-SIG-Storage 检索密钥
导入 GPG key 0xE451E5B5:
 用户ID     : "CentOS Storage SIG (http://wiki.centos.org/SpecialInterestGroup/Storage) <security@centos.org>"
 指纹       : 7412 9c0b 173b 071a 3775 951a d4a2 e50b e451 e5b5
 软件包     : centos-release-storage-common-1-2.el7.centos.noarch (@offline)
 来自       : /etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-SIG-Storage
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  正在更新    : glusterfs-libs-3.8.8-1.el7.x86_64                                                                    1/31 
  正在更新    : glusterfs-client-xlators-3.8.8-1.el7.x86_64                                                          2/31 
  正在更新    : glusterfs-3.8.8-1.el7.x86_64                                                                         3/31 
  正在安装    : libtirpc-0.2.4-0.8.el7.x86_64                                                                        4/31 
  正在安装    : rpcbind-0.2.0-38.el7.x86_64                                                                          5/31 
  正在安装    : psmisc-22.20-11.el7.x86_64                                                                           6/31 
  正在安装    : libcollection-0.6.2-27.el7.x86_64                                                                    7/31 
  正在安装    : libref_array-0.1.5-27.el7.x86_64                                                                     8/31 
  正在安装    : libbasicobjects-0.1.1-27.el7.x86_64                                                                  9/31 
  正在安装    : libtalloc-2.1.6-1.el7.x86_64                                                                        10/31 
  正在安装    : libtevent-0.9.28-1.el7.x86_64                                                                       11/31 
  正在安装    : libverto-tevent-0.2.5-4.el7.x86_64                                                                  12/31 
  正在更新    : glusterfs-fuse-3.8.8-1.el7.x86_64                                                                   13/31 
  正在安装    : glusterfs-api-3.8.8-1.el7.x86_64                                                                    14/31 
  正在安装    : glusterfs-cli-3.8.8-1.el7.x86_64                                                                    15/31 
  正在安装    : libpath_utils-0.2.1-27.el7.x86_64                                                                   16/31 
  正在安装    : libini_config-1.3.0-27.el7.x86_64                                                                   17/31 
  正在安装    : gssproxy-0.4.1-13.el7.x86_64                                                                        18/31 
  正在安装    : tcp_wrappers-7.6-77.el7.x86_64                                                                      19/31 
  正在安装    : keyutils-1.5.8-3.el7.x86_64                                                                         20/31 
  正在安装    : libevent-2.0.21-4.el7.x86_64                                                                        21/31 
  正在安装    : userspace-rcu-0.7.16-1.el7.x86_64                                                                   22/31 
  正在安装    : 1:quota-nls-4.01-14.el7.noarch                                                                      23/31 
  正在安装    : 1:quota-4.01-14.el7.x86_64                                                                          24/31 
  正在安装    : libnfsidmap-0.25-15.el7.x86_64                                                                      25/31 
  正在安装    : 1:nfs-utils-1.3.0-0.33.el7.x86_64                                                                   26/31 
  正在安装    : glusterfs-server-3.8.8-1.el7.x86_64                                                                 27/31 
  清理        : glusterfs-fuse-3.7.9-12.el7.centos.x86_64                                                           28/31 
  清理        : glusterfs-3.7.9-12.el7.centos.x86_64                                                                29/31 
  清理        : glusterfs-client-xlators-3.7.9-12.el7.centos.x86_64                                                 30/31 
  清理        : glusterfs-libs-3.7.9-12.el7.centos.x86_64                                                           31/31 
  验证中      : libtalloc-2.1.6-1.el7.x86_64                                                                         1/31 
  验证中      : 1:quota-4.01-14.el7.x86_64                                                                           2/31 
  验证中      : libnfsidmap-0.25-15.el7.x86_64                                                                       3/31 
  验证中      : libverto-tevent-0.2.5-4.el7.x86_64                                                                   4/31 
  验证中      : libini_config-1.3.0-27.el7.x86_64                                                                    5/31 
  验证中      : glusterfs-server-3.8.8-1.el7.x86_64                                                                  6/31 
  验证中      : 1:quota-nls-4.01-14.el7.noarch                                                                       7/31 
  验证中      : glusterfs-api-3.8.8-1.el7.x86_64                                                                     8/31 
  验证中      : libbasicobjects-0.1.1-27.el7.x86_64                                                                  9/31 
  验证中      : userspace-rcu-0.7.16-1.el7.x86_64                                                                   10/31 
  验证中      : glusterfs-libs-3.8.8-1.el7.x86_64                                                                   11/31 
  验证中      : libref_array-0.1.5-27.el7.x86_64                                                                    12/31 
  验证中      : gssproxy-0.4.1-13.el7.x86_64                                                                        13/31 
  验证中      : glusterfs-client-xlators-3.8.8-1.el7.x86_64                                                         14/31 
  验证中      : libevent-2.0.21-4.el7.x86_64                                                                        15/31 
  验证中      : glusterfs-cli-3.8.8-1.el7.x86_64                                                                    16/31 
  验证中      : libtevent-0.9.28-1.el7.x86_64                                                                       17/31 
  验证中      : glusterfs-3.8.8-1.el7.x86_64                                                                        18/31 
  验证中      : rpcbind-0.2.0-38.el7.x86_64                                                                         19/31 
  验证中      : libcollection-0.6.2-27.el7.x86_64                                                                   20/31 
  验证中      : glusterfs-fuse-3.8.8-1.el7.x86_64                                                                   21/31 
  验证中      : psmisc-22.20-11.el7.x86_64                                                                          22/31 
  验证中      : keyutils-1.5.8-3.el7.x86_64                                                                         23/31 
  验证中      : tcp_wrappers-7.6-77.el7.x86_64                                                                      24/31 
  验证中      : libpath_utils-0.2.1-27.el7.x86_64                                                                   25/31 
  验证中      : libtirpc-0.2.4-0.8.el7.x86_64                                                                       26/31 
  验证中      : 1:nfs-utils-1.3.0-0.33.el7.x86_64                                                                   27/31 
  验证中      : glusterfs-3.7.9-12.el7.centos.x86_64                                                                28/31 
  验证中      : glusterfs-client-xlators-3.7.9-12.el7.centos.x86_64                                                 29/31 
  验证中      : glusterfs-fuse-3.7.9-12.el7.centos.x86_64                                                           30/31 
  验证中      : glusterfs-libs-3.7.9-12.el7.centos.x86_64                                                           31/31 

已安装:
  glusterfs-server.x86_64 0:3.8.8-1.el7                                                                                   

作为依赖被安装:
  glusterfs-api.x86_64 0:3.8.8-1.el7     glusterfs-cli.x86_64 0:3.8.8-1.el7       gssproxy.x86_64 0:0.4.1-13.el7         
  keyutils.x86_64 0:1.5.8-3.el7          libbasicobjects.x86_64 0:0.1.1-27.el7    libcollection.x86_64 0:0.6.2-27.el7    
  libevent.x86_64 0:2.0.21-4.el7         libini_config.x86_64 0:1.3.0-27.el7      libnfsidmap.x86_64 0:0.25-15.el7       
  libpath_utils.x86_64 0:0.2.1-27.el7    libref_array.x86_64 0:0.1.5-27.el7       libtalloc.x86_64 0:2.1.6-1.el7         
  libtevent.x86_64 0:0.9.28-1.el7        libtirpc.x86_64 0:0.2.4-0.8.el7          libverto-tevent.x86_64 0:0.2.5-4.el7   
  nfs-utils.x86_64 1:1.3.0-0.33.el7      psmisc.x86_64 0:22.20-11.el7             quota.x86_64 1:4.01-14.el7             
  quota-nls.noarch 1:4.01-14.el7         rpcbind.x86_64 0:0.2.0-38.el7            tcp_wrappers.x86_64 0:7.6-77.el7       
  userspace-rcu.x86_64 0:0.7.16-1.el7   

更新完毕:
  glusterfs.x86_64 0:3.8.8-1.el7                            glusterfs-fuse.x86_64 0:3.8.8-1.el7                           

作为依赖被升级:
  glusterfs-client-xlators.x86_64 0:3.8.8-1.el7                    glusterfs-libs.x86_64 0:3.8.8-1.el7                   

完毕！


[root@localhost yum.repos.d]# systemctl list-unit-files glusterd.service
UNIT FILE        STATE   
glusterd.service disabled

1 unit files listed.


[root@localhost tmp]# systemctl enable glusterd.service
Created symlink from /etc/systemd/system/multi-user.target.wants/glusterd.service to /usr/lib/systemd/system/glusterd.service.
[root@localhost tmp]# systemctl start glusterd.service
[root@localhost tmp]# systemctl -l status glusterd.service
● glusterd.service - GlusterFS, a clustered file-system server
   Loaded: loaded (/usr/lib/systemd/system/glusterd.service; enabled; vendor preset: disabled)
   Active: active (running) since 一 2017-02-27 20:19:53 CST; 7s ago
  Process: 30089 ExecStart=/usr/sbin/glusterd -p /var/run/glusterd.pid --log-level $LOG_LEVEL $GLUSTERD_OPTIONS (code=exited, status=0/SUCCESS)
 Main PID: 30090 (glusterd)
   Memory: 12.9M
   CGroup: /system.slice/glusterd.service
           └─30090 /usr/sbin/glusterd -p /var/run/glusterd.pid --log-level INFO

2月 27 20:19:53 localhost.localdomain systemd[1]: Starting GlusterFS, a clustered file-system server...
2月 27 20:19:53 localhost.localdomain systemd[1]: Started GlusterFS, a clustered file-system server.






[root@zookeeper-3 yum.repos.d]# yum --disablerepo=\* --enablerepo=offline-base,offline-updates,offline-extras list | grep openshift
centos-release-openshift-origin.noarch  1-1.el7.centos           offline-extras 


[root@zookeeper-3 yum.repos.d]# yum --disablerepo=\* --enablerepo=offline-extras install -y centos-release-openshift-origin
已加载插件：fastestmirror, langpacks
Loading mirror speeds from cached hostfile
正在解决依赖关系
--> 正在检查事务
---> 软件包 centos-release-openshift-origin.noarch.0.1-1.el7.centos 将被 安装
--> 正在处理依赖关系 centos-release-paas-common，它被软件包 centos-release-openshift-origin-1-1.el7.centos.noarch 需要
--> 正在检查事务
---> 软件包 centos-release-paas-common.noarch.0.1-1.el7.centos 将被 安装
--> 解决依赖关系完成

依赖关系解决

============================================================================================================================
 Package                                     架构               版本                       源                          大小
============================================================================================================================
正在安装:
 centos-release-openshift-origin             noarch             1-1.el7.centos             offline-extras              11 k
为依赖而安装:
 centos-release-paas-common                  noarch             1-1.el7.centos             offline-extras              11 k

事务概要
============================================================================================================================
安装  1 软件包 (+1 依赖软件包)

总下载量：22 k
安装大小：37 k
Downloading packages:
(1/2): centos-release-paas-common-1-1.el7.centos.noarch.rpm                                          |  11 kB  00:00:00     
(2/2): centos-release-openshift-origin-1-1.el7.centos.noarch.rpm                                     |  11 kB  00:00:00     
----------------------------------------------------------------------------------------------------------------------------
总计                                                                                        214 kB/s |  22 kB  00:00:00     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  正在安装    : centos-release-paas-common-1-1.el7.centos.noarch                                                        1/2 
  正在安装    : centos-release-openshift-origin-1-1.el7.centos.noarch                                                   2/2 
  验证中      : centos-release-paas-common-1-1.el7.centos.noarch                                                        1/2 
  验证中      : centos-release-openshift-origin-1-1.el7.centos.noarch                                                   2/2 

已安装:
  centos-release-openshift-origin.noarch 0:1-1.el7.centos                                                                   

作为依赖被安装:
  centos-release-paas-common.noarch 0:1-1.el7.centos                                                                        

完毕！

[root@zookeeper-3 yum.repos.d]# ls /etc/pki/rpm-gpg/
RPM-GPG-KEY-CentOS-7        RPM-GPG-KEY-CentOS-SIG-PaaS     RPM-GPG-KEY-CentOS-Testing-7
RPM-GPG-KEY-CentOS-Debug-7  RPM-GPG-KEY-CentOS-SIG-Storage


[root@zookeeper-3 yum.repos.d]# yum --disablerepo=\* --enablerepo=offline-base,offline-updates,offline-extras list | grep openstack
centos-release-openstack.noarch         kilo-2.el7               offline-extras 
centos-release-openstack-kilo.noarch    1-2.el7                  offline-extras 
centos-release-openstack-liberty.noarch 1-4.el7                  offline-extras 
centos-release-openstack-mitaka.noarch  1-5.el7                  offline-extras 
centos-release-openstack-newton.noarch  1-1.el7                  offline-extras 

[root@zookeeper-3 yum.repos.d]# yum --disablerepo=\* --enablerepo=offline-extras install -y centos-release-openstack-newton
已加载插件：fastestmirror, langpacks
offline-extras                                                                                       | 3.4 kB  00:00:00     
Loading mirror speeds from cached hostfile
正在解决依赖关系
--> 正在检查事务
---> 软件包 centos-release-openstack-newton.noarch.0.1-1.el7 将被 安装
--> 正在处理依赖关系 centos-release-qemu-ev，它被软件包 centos-release-openstack-newton-1-1.el7.noarch 需要
--> 正在处理依赖关系 centos-release-ceph-jewel，它被软件包 centos-release-openstack-newton-1-1.el7.noarch 需要
--> 正在检查事务
---> 软件包 centos-release-ceph-jewel.noarch.0.1.0-1.el7.centos 将被 安装
--> 正在处理依赖关系 centos-release-storage-common，它被软件包 centos-release-ceph-jewel-1.0-1.el7.centos.noarch 需要
---> 软件包 centos-release-qemu-ev.noarch.0.1.0-1.el7 将被 安装
--> 正在处理依赖关系 centos-release-virt-common，它被软件包 centos-release-qemu-ev-1.0-1.el7.noarch 需要
--> 正在检查事务
---> 软件包 centos-release-storage-common.noarch.0.1-2.el7.centos 将被 安装
---> 软件包 centos-release-virt-common.noarch.0.1-1.el7.centos 将被 安装
--> 解决依赖关系完成

依赖关系解决

============================================================================================================================
 Package                                    架构              版本                          源                         大小
============================================================================================================================
正在安装:
 centos-release-openstack-newton            noarch            1-1.el7                       offline-extras            5.1 k
为依赖而安装:
 centos-release-ceph-jewel                  noarch            1.0-1.el7.centos              offline-extras            4.1 k
 centos-release-qemu-ev                     noarch            1.0-1.el7                     offline-extras             11 k
 centos-release-storage-common              noarch            1-2.el7.centos                offline-extras            4.5 k
 centos-release-virt-common                 noarch            1-1.el7.centos                offline-extras            4.5 k

事务概要
============================================================================================================================
安装  1 软件包 (+4 依赖软件包)

总下载量：29 k
安装大小：23 k
Downloading packages:
(1/5): centos-release-ceph-jewel-1.0-1.el7.centos.noarch.rpm                                         | 4.1 kB  00:00:00     
(2/5): centos-release-qemu-ev-1.0-1.el7.noarch.rpm                                                   |  11 kB  00:00:00     
(3/5): centos-release-openstack-newton-1-1.el7.noarch.rpm                                            | 5.1 kB  00:00:00     
(4/5): centos-release-storage-common-1-2.el7.centos.noarch.rpm                                       | 4.5 kB  00:00:00     
(5/5): centos-release-virt-common-1-1.el7.centos.noarch.rpm                                          | 4.5 kB  00:00:00     
----------------------------------------------------------------------------------------------------------------------------
总计                                                                                        265 kB/s |  29 kB  00:00:00     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  正在安装    : centos-release-virt-common-1-1.el7.centos.noarch                                                        1/5 
  正在安装    : centos-release-qemu-ev-1.0-1.el7.noarch                                                                 2/5 
  正在安装    : centos-release-storage-common-1-2.el7.centos.noarch                                                     3/5 
  正在安装    : centos-release-ceph-jewel-1.0-1.el7.centos.noarch                                                       4/5 
  正在安装    : centos-release-openstack-newton-1-1.el7.noarch                                                          5/5 
  验证中      : centos-release-storage-common-1-2.el7.centos.noarch                                                     1/5 
  验证中      : centos-release-openstack-newton-1-1.el7.noarch                                                          2/5 
  验证中      : centos-release-ceph-jewel-1.0-1.el7.centos.noarch                                                       3/5 
  验证中      : centos-release-virt-common-1-1.el7.centos.noarch                                                        4/5 
  验证中      : centos-release-qemu-ev-1.0-1.el7.noarch                                                                 5/5 

已安装:
  centos-release-openstack-newton.noarch 0:1-1.el7                                                                          

作为依赖被安装:
  centos-release-ceph-jewel.noarch 0:1.0-1.el7.centos            centos-release-qemu-ev.noarch 0:1.0-1.el7                  
  centos-release-storage-common.noarch 0:1-2.el7.centos          centos-release-virt-common.noarch 0:1-1.el7.centos         

完毕！


[root@zookeeper-3 yum.repos.d]# ls /etc/pki/rpm-gpg/
RPM-GPG-KEY-CentOS-7          RPM-GPG-KEY-CentOS-SIG-PaaS            RPM-GPG-KEY-CentOS-Testing-7
RPM-GPG-KEY-CentOS-Debug-7    RPM-GPG-KEY-CentOS-SIG-Storage
RPM-GPG-KEY-CentOS-SIG-Cloud  RPM-GPG-KEY-CentOS-SIG-Virtualization

[root@zookeeper-3 tmp]# yum --disablerepo=\* --enablerepo=offline-base,offline-updates,offline-extras list | grep "release-gluster"
centos-release-gluster36.noarch         1.0-3.el7.centos         offline-extras 
centos-release-gluster37.noarch         1.0-4.el7.centos         offline-extras 
centos-release-gluster38.noarch         1.0-1.el7.centos         offline-extras 

[root@zookeeper-3 tmp]# yum --disablerepo=\* --enablerepo=offline-base,offline-updates,offline-extras install -y centos-release-gluster38
已加载插件：fastestmirror, langpacks
Loading mirror speeds from cached hostfile
正在解决依赖关系
--> 正在检查事务
---> 软件包 centos-release-gluster38.noarch.0.1.0-1.el7.centos 将被 安装
--> 解决依赖关系完成

依赖关系解决

============================================================================================================================
 Package                               架构                版本                           源                           大小
============================================================================================================================
正在安装:
 centos-release-gluster38              noarch              1.0-1.el7.centos               offline-extras              4.1 k

事务概要
============================================================================================================================
安装  1 软件包

总下载量：4.1 k
安装大小：581  
Downloading packages:
centos-release-gluster38-1.0-1.el7.centos.noarch.rpm                                                 | 4.1 kB  00:00:00     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  正在安装    : centos-release-gluster38-1.0-1.el7.centos.noarch                                                        1/1 
  验证中      : centos-release-gluster38-1.0-1.el7.centos.noarch                                                        1/1 

已安装:
  centos-release-gluster38.noarch 0:1.0-1.el7.centos                                                                        

完毕！


[root@zookeeper-3 yum.repos.d]# ls
CentOS-Base.repo        CentOS-Debuginfo.repo    CentOS-OpenShift-Origin.repo  CentOS-Vault.repo
CentOS-Base.repo.bak    CentOS-fasttrack.repo    CentOS-OpenStack-newton.repo  offline.repo
CentOS-Ceph-Jewel.repo  CentOS-Gluster-3.8.repo  CentOS-QEMU-EV.repo
CentOS-CR.repo          CentOS-Media.repo        CentOS-Sources.repo


[root@zookeeper-3 yum.repos.d]# yum --disablerepo=\* --enablerepo=offline-base,offline-updates,offline-extras,offline-gluster38 search glusterfs
已加载插件：fastestmirror, langpacks
Loading mirror speeds from cached hostfile
================================================== N/S matched: glusterfs ==================================================
glusterfs-api.x86_64 : GlusterFS api library
glusterfs-cli.x86_64 : GlusterFS CLI
glusterfs-client-xlators.x86_64 : GlusterFS client-side translators
glusterfs-geo-replication.x86_64 : GlusterFS Geo-replication
glusterfs-libs.x86_64 : GlusterFS common libraries
glusterfs-rdma.x86_64 : GlusterFS rdma support for ib-verbs
glusterfs-resource-agents.noarch : OCF Resource Agents for GlusterFS
samba-vfs-glusterfs.x86_64 : Samba VFS module for GlusterFS
centos-release-gluster36.noarch : GlusterFS 3.6 packages from the CentOS Storage SIG repository
centos-release-gluster37.noarch : GlusterFS 3.7 packages from the CentOS Storage SIG repository
centos-release-gluster38.noarch : GlusterFS 3.8 packages from the CentOS Storage SIG repository
glusterfs.x86_64 : Distributed File System
glusterfs-api-devel.x86_64 : Development Libraries
glusterfs-coreutils.x86_64 : Core Utilities for the Gluster Distributed File System
glusterfs-devel.x86_64 : Development Libraries
glusterfs-extra-xlators.x86_64 : Extra Gluster filesystem Translators
glusterfs-fuse.x86_64 : Fuse client
glusterfs-ganesha.x86_64 : NFS-Ganesha configuration
glusterfs-server.x86_64 : Clustered file-system server
heketi.x86_64 : RESTful based volume management framework for GlusterFS
heketi-devel.noarch : RESTful based volume management framework for GlusterFS
python-gluster.noarch : GlusterFS python library

  名称和简介匹配 only，使用“search all”试试。

[root@zookeeper-3 tmp]# yum --disablerepo=\* --enablerepo=offline-base,offline-updates,offline-extras,offline-gluster38 list | grep glusterfs
glusterfs.x86_64                        3.7.9-12.el7.centos      @anaconda      
glusterfs-api.x86_64                    3.7.9-12.el7.centos      @anaconda      
glusterfs-client-xlators.x86_64         3.7.9-12.el7.centos      @anaconda      
glusterfs-fuse.x86_64                   3.7.9-12.el7.centos      @anaconda      
glusterfs-libs.x86_64                   3.7.9-12.el7.centos      @anaconda      
glusterfs.x86_64                        3.8.8-1.el7              offline-gluster38
glusterfs-api.x86_64                    3.8.8-1.el7              offline-gluster38
glusterfs-api-devel.x86_64              3.8.8-1.el7              offline-gluster38
glusterfs-cli.x86_64                    3.8.8-1.el7              offline-gluster38
glusterfs-client-xlators.x86_64         3.8.8-1.el7              offline-gluster38
glusterfs-coreutils.x86_64              0.2.0-1.el7              offline-gluster38
glusterfs-devel.x86_64                  3.8.8-1.el7              offline-gluster38
glusterfs-extra-xlators.x86_64          3.8.8-1.el7              offline-gluster38
glusterfs-fuse.x86_64                   3.8.8-1.el7              offline-gluster38
glusterfs-ganesha.x86_64                3.8.8-1.el7              offline-gluster38
glusterfs-geo-replication.x86_64        3.8.8-1.el7              offline-gluster38
glusterfs-libs.x86_64                   3.8.8-1.el7              offline-gluster38
glusterfs-rdma.x86_64                   3.8.8-1.el7              offline-gluster38
glusterfs-resource-agents.noarch        3.8.8-1.el7              offline-gluster38
glusterfs-server.x86_64                 3.8.8-1.el7              offline-gluster38
samba-vfs-glusterfs.x86_64              4.4.4-12.el7_3           offline-updates


[root@zookeeper-3 tmp]# mkdir -p {brick1,brick2}
[root@zookeeper-3 tmp]# ls
brick1
brick2


[root@zookeeper-3 tmp]# yum --disablerepo=\* --enablerepo=offline-base,offline-updates,offline-extras,offline-gluster38 install -y glusterfs gluster-server gluster-fuse
已加载插件：fastestmirror, langpacks
Loading mirror speeds from cached hostfile
没有可用软件包 gluster-server。
没有可用软件包 gluster-fuse。
正在解决依赖关系
--> 正在检查事务
---> 软件包 glusterfs.x86_64.0.3.7.9-12.el7.centos 将被 升级
--> 正在处理依赖关系 glusterfs(x86-64) = 3.7.9-12.el7.centos，它被软件包 glusterfs-api-3.7.9-12.el7.centos.x86_64 需要
--> 正在处理依赖关系 glusterfs(x86-64) = 3.7.9-12.el7.centos，它被软件包 glusterfs-fuse-3.7.9-12.el7.centos.x86_64 需要
---> 软件包 glusterfs.x86_64.0.3.8.8-1.el7 将被 更新
--> 正在处理依赖关系 glusterfs-libs(x86-64) = 3.8.8-1.el7，它被软件包 glusterfs-3.8.8-1.el7.x86_64 需要
--> 正在检查事务
---> 软件包 glusterfs-api.x86_64.0.3.7.9-12.el7.centos 将被 升级
---> 软件包 glusterfs-api.x86_64.0.3.8.8-1.el7 将被 更新
--> 正在处理依赖关系 glusterfs-client-xlators(x86-64) = 3.8.8-1.el7，它被软件包 glusterfs-api-3.8.8-1.el7.x86_64 需要
---> 软件包 glusterfs-fuse.x86_64.0.3.7.9-12.el7.centos 将被 升级
---> 软件包 glusterfs-fuse.x86_64.0.3.8.8-1.el7 将被 更新
---> 软件包 glusterfs-libs.x86_64.0.3.7.9-12.el7.centos 将被 升级
---> 软件包 glusterfs-libs.x86_64.0.3.8.8-1.el7 将被 更新
--> 正在检查事务
---> 软件包 glusterfs-client-xlators.x86_64.0.3.7.9-12.el7.centos 将被 升级
---> 软件包 glusterfs-client-xlators.x86_64.0.3.8.8-1.el7 将被 更新
--> 解决依赖关系完成

依赖关系解决

============================================================================================================================
 Package                               架构                版本                        源                              大小
============================================================================================================================
正在更新:
 glusterfs                             x86_64              3.8.8-1.el7                 offline-gluster38              510 k
为依赖而更新:
 glusterfs-api                         x86_64              3.8.8-1.el7                 offline-gluster38               90 k
 glusterfs-client-xlators              x86_64              3.8.8-1.el7                 offline-gluster38              781 k
 glusterfs-fuse                        x86_64              3.8.8-1.el7                 offline-gluster38              133 k
 glusterfs-libs                        x86_64              3.8.8-1.el7                 offline-gluster38              379 k

事务概要
============================================================================================================================
升级  1 软件包 (+4 依赖软件包)

总下载量：1.8 M
Downloading packages:
No Presto metadata available for offline-gluster38
警告：/var/cache/yum/x86_64/7/offline-gluster38/packages/glusterfs-3.8.8-1.el7.x86_64.rpm: 头V4 RSA/SHA1 Signature, 密钥 ID e451e5b5: NOKEY
glusterfs-3.8.8-1.el7.x86_64.rpm 的公钥尚未安装
(1/5): glusterfs-3.8.8-1.el7.x86_64.rpm                                                              | 510 kB  00:00:00     
(2/5): glusterfs-api-3.8.8-1.el7.x86_64.rpm                                                          |  90 kB  00:00:00     
(3/5): glusterfs-client-xlators-3.8.8-1.el7.x86_64.rpm                                               | 781 kB  00:00:00     
(4/5): glusterfs-fuse-3.8.8-1.el7.x86_64.rpm                                                         | 133 kB  00:00:00     
(5/5): glusterfs-libs-3.8.8-1.el7.x86_64.rpm                                                         | 379 kB  00:00:00     
----------------------------------------------------------------------------------------------------------------------------
总计                                                                                         14 MB/s | 1.8 MB  00:00:00     
从 file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-SIG-Storage 检索密钥
导入 GPG key 0xE451E5B5:
 用户ID     : "CentOS Storage SIG (http://wiki.centos.org/SpecialInterestGroup/Storage) <security@centos.org>"
 指纹       : 7412 9c0b 173b 071a 3775 951a d4a2 e50b e451 e5b5
 软件包     : centos-release-storage-common-1-2.el7.centos.noarch (@offline-extras)
 来自       : /etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-SIG-Storage
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  正在更新    : glusterfs-libs-3.8.8-1.el7.x86_64                                                                      1/10 
  正在更新    : glusterfs-client-xlators-3.8.8-1.el7.x86_64                                                            2/10 
  正在更新    : glusterfs-3.8.8-1.el7.x86_64                                                                           3/10 
  正在更新    : glusterfs-api-3.8.8-1.el7.x86_64                                                                       4/10 
  正在更新    : glusterfs-fuse-3.8.8-1.el7.x86_64                                                                      5/10 
  清理        : glusterfs-fuse-3.7.9-12.el7.centos.x86_64                                                              6/10 
  清理        : glusterfs-api-3.7.9-12.el7.centos.x86_64                                                               7/10 
  清理        : glusterfs-3.7.9-12.el7.centos.x86_64                                                                   8/10 
  清理        : glusterfs-client-xlators-3.7.9-12.el7.centos.x86_64                                                    9/10 
  清理        : glusterfs-libs-3.7.9-12.el7.centos.x86_64                                                             10/10 
  验证中      : glusterfs-api-3.8.8-1.el7.x86_64                                                                       1/10 
  验证中      : glusterfs-fuse-3.8.8-1.el7.x86_64                                                                      2/10 
  验证中      : glusterfs-client-xlators-3.8.8-1.el7.x86_64                                                            3/10 
  验证中      : glusterfs-libs-3.8.8-1.el7.x86_64                                                                      4/10 
  验证中      : glusterfs-3.8.8-1.el7.x86_64                                                                           5/10 
  验证中      : glusterfs-client-xlators-3.7.9-12.el7.centos.x86_64                                                    6/10 
  验证中      : glusterfs-3.7.9-12.el7.centos.x86_64                                                                   7/10 
  验证中      : glusterfs-api-3.7.9-12.el7.centos.x86_64                                                               8/10 
  验证中      : glusterfs-libs-3.7.9-12.el7.centos.x86_64                                                              9/10 
  验证中      : glusterfs-fuse-3.7.9-12.el7.centos.x86_64                                                             10/10 

更新完毕:
  glusterfs.x86_64 0:3.8.8-1.el7                                                                                            

作为依赖被升级:
  glusterfs-api.x86_64 0:3.8.8-1.el7   glusterfs-client-xlators.x86_64 0:3.8.8-1.el7  glusterfs-fuse.x86_64 0:3.8.8-1.el7 
  glusterfs-libs.x86_64 0:3.8.8-1.el7 

完毕！


[root@zookeeper-3 yum.repos.d]# yum --disablerepo=\* --enablerepo=offline-base,offline-updates,offline-extras,offline-gluster38 install -y glusterfs glusterfs-server glusterfs-fuse
已加载插件：fastestmirror, langpacks
Loading mirror speeds from cached hostfile
软件包 glusterfs-3.8.8-1.el7.x86_64 已安装并且是最新版本
软件包 glusterfs-fuse-3.8.8-1.el7.x86_64 已安装并且是最新版本
正在解决依赖关系
--> 正在检查事务
---> 软件包 glusterfs-server.x86_64.0.3.8.8-1.el7 将被 安装
--> 正在处理依赖关系 glusterfs-cli(x86-64) = 3.8.8-1.el7，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 liburcu-cds.so.1()(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在处理依赖关系 liburcu-bp.so.1()(64bit)，它被软件包 glusterfs-server-3.8.8-1.el7.x86_64 需要
--> 正在检查事务
---> 软件包 glusterfs-cli.x86_64.0.3.8.8-1.el7 将被 安装
---> 软件包 userspace-rcu.x86_64.0.0.7.16-1.el7 将被 安装
--> 解决依赖关系完成

依赖关系解决

============================================================================================================================
 Package                         架构                  版本                          源                                大小
============================================================================================================================
正在安装:
 glusterfs-server                x86_64                3.8.8-1.el7                   offline-gluster38                1.4 M
为依赖而安装:
 glusterfs-cli                   x86_64                3.8.8-1.el7                   offline-gluster38                183 k
 userspace-rcu                   x86_64                0.7.16-1.el7                  offline-gluster38                 72 k

事务概要
============================================================================================================================
安装  1 软件包 (+2 依赖软件包)

总下载量：1.7 M
安装大小：5.3 M
Downloading packages:
(1/3): glusterfs-cli-3.8.8-1.el7.x86_64.rpm                                                          | 183 kB  00:00:00     
(2/3): userspace-rcu-0.7.16-1.el7.x86_64.rpm                                                         |  72 kB  00:00:00     
(3/3): glusterfs-server-3.8.8-1.el7.x86_64.rpm                                                       | 1.4 MB  00:00:00     
----------------------------------------------------------------------------------------------------------------------------
总计                                                                                         13 MB/s | 1.7 MB  00:00:00     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  正在安装    : userspace-rcu-0.7.16-1.el7.x86_64                                                                       1/3 
  正在安装    : glusterfs-cli-3.8.8-1.el7.x86_64                                                                        2/3 
  正在安装    : glusterfs-server-3.8.8-1.el7.x86_64                                                                     3/3 
  验证中      : glusterfs-cli-3.8.8-1.el7.x86_64                                                                        1/3 
  验证中      : userspace-rcu-0.7.16-1.el7.x86_64                                                                       2/3 
  验证中      : glusterfs-server-3.8.8-1.el7.x86_64                                                                     3/3 

已安装:
  glusterfs-server.x86_64 0:3.8.8-1.el7                                                                                     

作为依赖被安装:
  glusterfs-cli.x86_64 0:3.8.8-1.el7                           userspace-rcu.x86_64 0:0.7.16-1.el7                          

完毕！


[root@zookeeper-3 yum.repos.d]# systemctl list-unit-files --type=service glusterd.service
UNIT FILE        STATE   
glusterd.service disabled

1 unit files listed.


[root@zookeeper-3 yum.repos.d]# systemctl enable glusterd.service
Created symlink from /etc/systemd/system/multi-user.target.wants/glusterd.service to /usr/lib/systemd/system/glusterd.service.
[root@zookeeper-3 yum.repos.d]# systemctl start glusterd.service



[root@zookeeper-3 yum.repos.d]# getenforce 
Enforcing
[root@zookeeper-3 yum.repos.d]# setenforce 0
[root@zookeeper-3 yum.repos.d]# getenforce 
Permissive
[root@zookeeper-3 yum.repos.d]# systemctl list-units firewalld.service
0 loaded units listed. Pass --all to see loaded but inactive units, too.
To show all installed unit files use 'systemctl list-unit-files'.


[root@zookeeper-3 yum.repos.d]# gluster peer probe 192.168.2.20
peer probe: success. 
[root@zookeeper-3 yum.repos.d]# gluster peer status
Number of Peers: 1

Hostname: 192.168.2.20
Uuid: 21a1e975-4c2c-4259-8938-d40a5cdf5f91
State: Peer in Cluster (Connected)


[root@zookeeper-3 yum.repos.d]# systemctl -l status glusterd.service
● glusterd.service - GlusterFS, a clustered file-system server
   Loaded: loaded (/usr/lib/systemd/system/glusterd.service; enabled; vendor preset: disabled)
   Active: active (running) since 二 2017-02-28 12:51:08 CST; 8min ago
  Process: 13780 ExecStart=/usr/sbin/glusterd -p /var/run/glusterd.pid --log-level $LOG_LEVEL $GLUSTERD_OPTIONS (code=exited, status=0/SUCCESS)
 Main PID: 13781 (glusterd)
   CGroup: /system.slice/glusterd.service
           └─13781 /usr/sbin/glusterd -p /var/run/glusterd.pid --log-level INFO

2月 28 12:51:08 zookeeper-3 systemd[1]: Starting GlusterFS, a clustered file-system server...
2月 28 12:51:08 zookeeper-3 systemd[1]: Started GlusterFS, a clustered file-system server.


[root@localhost tmp]# vgs
  VG #PV #LV #SN Attr   VSize  VFree
  cl   1   2   0 wz--n- 99.00g 4.00m
[root@localhost tmp]# lvs
  LV   VG Attr       LSize  Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
  root cl -wi-ao---- 91.12g                                                    
  swap cl -wi-ao----  7.88g                                                    
[root@localhost tmp]# lvreduce --size=-10g cl/root
  WARNING: Reducing active and open logical volume to 81.12 GiB.
  THIS MAY DESTROY YOUR DATA (filesystem etc.)
Do you really want to reduce cl/root? [y/n]: y
  Size of logical volume cl/root changed from 91.12 GiB (23326 extents) to 81.12 GiB (20766 extents).
  Logical volume cl/root successfully resized.

[root@localhost tmp]# lvcreate -L 5G -n brick1 cl
  Logical volume "brick1" created.
[root@localhost tmp]# lvcreate -L 5G -n brick2 cl
  Logical volume "brick2" created.

[root@localhost tmp]# mkfs.xfs /dev/mapper/cl-brick1
meta-data=/dev/mapper/cl-brick1  isize=512    agcount=4, agsize=327680 blks
         =                       sectsz=512   attr=2, projid32bit=1
         =                       crc=1        finobt=0, sparse=0
data     =                       bsize=4096   blocks=1310720, imaxpct=25
         =                       sunit=0      swidth=0 blks
naming   =version 2              bsize=4096   ascii-ci=0 ftype=1
log      =internal log           bsize=4096   blocks=2560, version=2
         =                       sectsz=512   sunit=0 blks, lazy-count=1
realtime =none                   extsz=4096   blocks=0, rtextents=0
[root@localhost tmp]# mkfs.xfs /dev/mapper/cl-brick2
meta-data=/dev/mapper/cl-brick2  isize=512    agcount=4, agsize=327680 blks
         =                       sectsz=512   attr=2, projid32bit=1
         =                       crc=1        finobt=0, sparse=0
data     =                       bsize=4096   blocks=1310720, imaxpct=25
         =                       sunit=0      swidth=0 blks
naming   =version 2              bsize=4096   ascii-ci=0 ftype=1
log      =internal log           bsize=4096   blocks=2560, version=2
         =                       sectsz=512   sunit=0 blks, lazy-count=1
realtime =none                   extsz=4096   blocks=0, rtextents=0
[root@localhost tmp]# mount /dev/mapper/cl-brick1 /tmp/brick1
[root@localhost tmp]# mount /dev/mapper/cl-brick2 /tmp/brick2



[root@zookeeper-3 tmp]# lvcreate -L 5G -n brick1 cl
  Logical volume "brick1" created.
[root@zookeeper-3 tmp]# lvcreate -L 5G -n brick2 cl
  Logical volume "brick2" created.
[root@zookeeper-3 tmp]# mkfs.xfs /dev/mapper/cl-brick1
meta-data=/dev/mapper/cl-brick1  isize=512    agcount=4, agsize=327680 blks
         =                       sectsz=512   attr=2, projid32bit=1
         =                       crc=1        finobt=0, sparse=0
data     =                       bsize=4096   blocks=1310720, imaxpct=25
         =                       sunit=0      swidth=0 blks
naming   =version 2              bsize=4096   ascii-ci=0 ftype=1
log      =internal log           bsize=4096   blocks=2560, version=2
         =                       sectsz=512   sunit=0 blks, lazy-count=1
realtime =none                   extsz=4096   blocks=0, rtextents=0
[root@zookeeper-3 tmp]# mkfs.xfs /dev/mapper/cl-brick2
meta-data=/dev/mapper/cl-brick2  isize=512    agcount=4, agsize=327680 blks
         =                       sectsz=512   attr=2, projid32bit=1
         =                       crc=1        finobt=0, sparse=0
data     =                       bsize=4096   blocks=1310720, imaxpct=25
         =                       sunit=0      swidth=0 blks
naming   =version 2              bsize=4096   ascii-ci=0 ftype=1
log      =internal log           bsize=4096   blocks=2560, version=2
         =                       sectsz=512   sunit=0 blks, lazy-count=1
realtime =none                   extsz=4096   blocks=0, rtextents=0


[root@zookeeper-3 tmp]# mount /dev/mapper/cl-brick1 /tmp/brick1
[root@zookeeper-3 tmp]# mount /dev/mapper/cl-brick2 /tmp/brick2


[root@zookeeper-3 tmp]# mkdir -p /tmp/brick1/brick /tmp/brick2/brick


[root@zookeeper-3 tmp]# gluster volume create glustervol1 replica 2 transport tcp 192.168.2.20:/tmp/brick1/brick 192.168.2.19:/tmp/brick1/brick
volume create: glustervol1: success: please start the volume to access data

[root@zookeeper-3 tmp]# gluster volume start glustervol1
volume start: glustervol1: success
[root@zookeeper-3 tmp]# ls /tmp/brick1/brick
[root@zookeeper-3 tmp]# gluster volume info all
 
Volume Name: glustervol1
Type: Replicate
Volume ID: feea39e5-434b-450b-9783-4fee9996b670
Status: Started
Snapshot Count: 0
Number of Bricks: 1 x 2 = 2
Transport-type: tcp
Bricks:
Brick1: 192.168.2.20:/tmp/brick1/brick
Brick2: 192.168.2.19:/tmp/brick1/brick
Options Reconfigured:
transport.address-family: inet
performance.readdir-ahead: on
nfs.disable: on


[root@zookeeper-3 tmp]# gluster volume info all
 
Volume Name: glustervol1
Type: Replicate
Volume ID: feea39e5-434b-450b-9783-4fee9996b670
Status: Created
Snapshot Count: 0
Number of Bricks: 1 x 2 = 2
Transport-type: tcp
Bricks:
Brick1: 192.168.2.20:/tmp/brick1/brick
Brick2: 192.168.2.19:/tmp/brick1/brick
Options Reconfigured:
transport.address-family: inet
performance.readdir-ahead: on
nfs.disable: on



[root@localhost /]# mount -t glusterfs 192.168.2.19:/glustervol1 /mnt/

[root@localhost ecp]# gluster peer probe 192.168.2.19
peer probe: success. Host 192.168.2.19 port 24007 already in peer list
[root@localhost ecp]# gluster peer status
Number of Peers: 1

Hostname: 192.168.2.19
Uuid: 804a1ac1-bc16-4f38-bd62-a55a953ea5d0
State: Peer in Cluster (Connected)

[root@zookeeper-3 brick]# gluster peer probe 192.168.2.20
peer probe: success. Host 192.168.2.20 port 24007 already in peer list
[root@zookeeper-3 brick]# gluster peer status
Number of Peers: 1

Hostname: 192.168.2.20
Uuid: 21a1e975-4c2c-4259-8938-d40a5cdf5f91
State: Peer in Cluster (Connected)


