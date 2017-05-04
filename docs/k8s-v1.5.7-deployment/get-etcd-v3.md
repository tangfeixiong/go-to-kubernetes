Getting ETCD v3
---------------

For example (https://github.com/coreos/etcd/releases/tag/v3.0.17)

```
vagrant@vagrant-ubuntu-trusty-64:~$ ETCD_VER=v3.0.17
vagrant@vagrant-ubuntu-trusty-64:~$ DOWNLOAD_URL=https://github.com/coreos/etcd/releases/download
vagrant@vagrant-ubuntu-trusty-64:~$ curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -O
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   607    0   607    0     0    117      0 --:--:--  0:00:05 --:--:--   166
100  9.8M  100  9.8M    0     0  68845      0  0:02:29  0:02:29 --:--:-- 50495
vagrant@vagrant-ubuntu-trusty-64:~$ tar -C /vagrant/workspace/bin -zxf etcd-v3.0.17-linux-amd64.tar.gz --strip-components=1 etcd-v3.0.17-linux-amd64/etcd
vagrant@vagrant-ubuntu-trusty-64:~$ tar -C /vagrant/workspace/bin -zxf etcd-v3.0.17-linux-amd64.tar.gz --strip-components=1 etcd-v3.0.17-linux-amd64/etcdctl
vagrant@vagrant-ubuntu-trusty-64:~$ /vagrant/workspace/bin/etcd --version
etcd Version: 3.0.17
Git SHA: cc198e2
Go Version: go1.6.4
Go OS/Arch: linux/amd64
vagrant@vagrant-ubuntu-trusty-64:~$ /vagrant/workspace/bin/etcdctl --version
etcdctl version: 3.0.17
API version: 2

```