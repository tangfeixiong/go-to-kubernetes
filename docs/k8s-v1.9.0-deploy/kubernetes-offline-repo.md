## DevOps

### Fedora 26

gpg, if not imported ever
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo rpm --import http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/doc/yum-key.gpg
```

repo
```
[vagrant@kubedev-172-17-4-59 ~]$ curl -L http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/etc0x2Fyum.repos.d/kubernetes.repo | sudo tee /etc/yum.repos.d/kubernetes.repo
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1011  100  1011    0     0  31991      0 --:--:-- --:--:-- --:--:-- 37444
[kubernetes]
name=Kubernetes
#baseurl=http://yum.kubernetes.io/repos/kubernetes-el7-x86_64
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
        https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
		
### Modify URL first ###		
[kubernetes-local-http]
name=Kubernetes offline - file server
baseurl=http://127.0.0.1:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/
enabled=0
gpgcheck=0

### Modify path first ###
[kubernetes-local-file]
name=Kubernetes offline - downloads
baseurl=file:////Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/
enabled=0
gpgcheck=0

```

Issue
```
[vagrant@kubedev-172-17-4-59 ~]$ sudo dnf --disablerepo kubernetes --enablerepo kubernetes-local\* repolist
Failed to set locale, defaulting to C
Failed to synchronize cache for repo 'kubernetes-local-http', disabling.
Failed to synchronize cache for repo 'kubernetes-local-file', disabling.
Last metadata expiration check: 0:12:50 ago on Mon Dec 18 05:53:18 2017.
repo id                                                    repo name                                                                      status
*fedora                                                    Fedora 26 - x86_64                                                             53912
*updates                                                   Fedora 26 - x86_64 - Updates                                                   11821
```

https://unix.stackexchange.com/questions/390805/repos-not-working-on-fedora-error-failed-to-synchronize-cache-for-repo-update


