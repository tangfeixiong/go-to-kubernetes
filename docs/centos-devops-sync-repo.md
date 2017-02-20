Sync CentOS packages for mirror YUM Repositories - 镜像YUM安装仓库
===================================================================

Tables of Content
-------------------

* Sync CentOS packages

* Search sample software packages into meida repo

* Install software from media repo

Using Everything ISO as OS-Media Repo - 下载CentOS Everything ISO作为离线包仓库
-----------------------------------------------------------------------------

Execute __Linux__ command _wget_ to download - 从[阿里云镜像站](http://mirrors.aliyun.com)下载

```
tangf@DESKTOP-H68OQDV /cygdrive/f/16-mirror/centos
$ ls
centos-rsync.filter
dockerproject.exclude
extras.exclude
extras.include
https%3A%2F%2Fpackages.cloud.google.com%2Fyum
kubernetes.repo
kubernetes.wgetrc
kubernetes-yum-key.gpg
kubernetes-yum-package-key.gpg
mirror-by-rsync.sh
mirror-by-wget.sh
paas.exclude
paas.include
rsync%3A%2F%2Fmirrors.yun-idc.com%2Fcentos%2F7
storage.exclude
storage.include

tangf@DESKTOP-H68OQDV /cygdrive/f/16-mirror/centos
$ ./mirror-by-rsync.sh

receiving incremental file list
extras/x86_64/Packages/
extras/x86_64/drpms/
extras/x86_64/repodata/

sent 1,122 bytes  received 12,107 bytes  8,819.33 bytes/sec
total size is 464,833,789  speedup is 35,137.48

```

The mirror repos except __isos__ and __os__ (same with Everything ISO), and size is a bit larger

```
tangf@DESKTOP-H68OQDV /cygdrive/f/16-mirror/centos
$ ls -1 /cygdrive/f/16-mirror/centos/rsync%3A%2F%2Fmirrors.yun-idc.com%2Fcentos%2F7/
centosplus
cr
extras
fasttrack
paas
sclo
storage
updates
virt

tangf@DESKTOP-H68OQDV /cygdrive/f/16-mirror/centos
$ du -sh rsync%3A%2F%2Fmirrors.yun-idc.com%2Fcentos%2F7/
7.5G    rsync%3A%2F%2Fmirrors.yun-idc.com%2Fcentos%2F7/

localhost:puhua-yanyong fanhongling$ ls /Volumes/TOURO\ Mobile/99-mirror/centos/rsync%3A%2F%2Fmirrors.yun-idc.com%2Fcentos%2F7/extras/x86_64/Packages/ | wc -l
     414

```

mirror repo

```
[vagrant@openshiftdev yum.repos.d]$ mkdir -p /home/vagrant/offline-repo
[vagrant@openshiftdev yum.repos.d]$ cp -r /vagrant_data/16-mirror/centos/rsync%3A%2F%2Fmirrors.yun-idc.com%2Fcentos%2F7/* /home/vagrant/offline-repo/
[vagrant@openshiftdev yum.repos.d]$ sudo vi offline.repo                        

```

For example, list packages

```
[vagrant@openshiftdev yum.repos.d]$ sudo yum --disablerepo=\* --enablerepo=c7-media --enablerepo=offline list | grep docker
docker.x86_64                             1.10.3-46.el7.centos.14      @extras
docker-common.x86_64                      1.10.3-46.el7.centos.14      @extras
docker-selinux.x86_64                     1.10.3-46.el7.centos.14      @extras
cockpit-docker.x86_64                     122-3.el7.centos             offline
docker.x86_64                             2:1.10.3-59.el7.centos       offline
docker-common.x86_64                      2:1.10.3-59.el7.centos       offline
docker-devel.x86_64                       1.3.2-4.el7.centos           offline
docker-distribution.x86_64                2.5.1-1.el7                  offline
docker-forward-journald.x86_64            1.10.3-44.el7.centos         offline
docker-latest.x86_64                      1.12.3-10.el7.centos         offline
docker-latest-logrotate.x86_64            1.12.3-10.el7.centos         offline
docker-latest-v1.10-migrator.x86_64       1.12.3-10.el7.centos         offline
docker-logrotate.x86_64                   2:1.10.3-59.el7.centos       offline
docker-lvm-plugin.x86_64                  2:1.10.3-59.el7.centos       offline
docker-novolume-plugin.x86_64             2:1.10.3-59.el7.centos       offline
docker-python.x86_64                      1.4.0-115.el7                offline
docker-registry.noarch                    0.6.8-8.el7                  offline
docker-registry.x86_64                    0.9.1-7.el7                  offline
docker-unit-test.x86_64                   2:1.10.3-59.el7.centos       offline
docker-v1.10-migrator.x86_64              2:1.10.3-59.el7.centos       offline
python-docker-py.noarch                   1.9.0-1.el7                  offline

[vagrant@openshiftdev yum.repos.d]$ sudo yum --disablerepo=\* --enablerepo=c7-media --enablerepo=offline list | grep gluster
centos-release-gluster36.noarch           1.0-3.el7.centos             offline
centos-release-gluster37.noarch           1.0-4.el7.centos             offline
centos-release-gluster38.noarch           1.0-1.el7.centos             offline
glusterfs.x86_64                          3.7.9-12.el7.centos          c7-media
glusterfs-api.x86_64                      3.7.9-12.el7.centos          c7-media
glusterfs-api-devel.x86_64                3.7.9-12.el7.centos          c7-media
glusterfs-cli.x86_64                      3.7.9-12.el7.centos          c7-media
glusterfs-client-xlators.x86_64           3.7.9-12.el7.centos          c7-media
glusterfs-devel.x86_64                    3.7.9-12.el7.centos          c7-media
glusterfs-fuse.x86_64                     3.7.9-12.el7.centos          c7-media
glusterfs-libs.x86_64                     3.7.9-12.el7.centos          c7-media
glusterfs-rdma.x86_64                     3.7.9-12.el7.centos          c7-media
pcp-pmda-gluster.x86_64                   3.11.3-4.el7                 c7-media
python-gluster.noarch                     3.7.9-12.el7.centos          c7-media
samba-vfs-glusterfs.x86_64                4.4.4-9.el7                  c7-media

[vagrant@openshiftdev yum.repos.d]$ sudo yum --disablerepo=\* --enablerepo=c7-media --enablerepo=offline list | grep ceph
centos-release-ceph-hammer.noarch         1.0-5.el7.centos             offline
centos-release-ceph-jewel.noarch          1.0-1.el7.centos             offline
ceph-common.x86_64                        1:0.94.5-1.el7               c7-media

[vagrant@openshiftdev yum.repos.d]$ sudo yum --disablerepo=\* --enablerepo=c7-media --enablerepo=offline list | grep etcd
etcd.x86_64                               2.3.7-4.el7                  offline
etcd3.x86_64                              3.0.14-2.el7                 offline

[vagrant@openshiftdev yum.repos.d]$ sudo yum --disablerepo=\* --enablerepo=c7-media --enablerepo=offline list | grep flannel
flannel.x86_64                            0.5.5-1.el7                  offline

[vagrant@openshiftdev yum.repos.d]$ sudo yum --disablerepo=\* --enablerepo=c7-media --enablerepo=offline list | grep kubernetes
cockpit-kubernetes.x86_64                 122-3.el7.centos             offline
kubernetes.x86_64                         1.3.0-0.3.git86dc49a.el7     offline
kubernetes-client.x86_64                  1.3.0-0.3.git86dc49a.el7     offline
kubernetes-master.x86_64                  1.3.0-0.3.git86dc49a.el7     offline
kubernetes-node.x86_64                    1.3.0-0.3.git86dc49a.el7     offline
kubernetes-unit-test.x86_64               1.3.0-0.3.git86dc49a.el7     offline

```
