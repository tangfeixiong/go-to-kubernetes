

pki
```
ubuntu@kubedev-10-64-33-200:~$ scp -r ubuntu@10.64.33.195:/home/ubuntu/etc0x2Fkubernetes .
The authenticity of host '10.64.33.195 (10.64.33.195)' can't be established.
ECDSA key fingerprint is SHA256:QQBkXG5+IgQTMn7rh2aTbJi+IQLE106NbIBXJqKIoPI.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '10.64.33.195' (ECDSA) to the list of known hosts.
kubelet.conf                                                                                                  100% 5539     5.4KB/s   00:00    
controller-manager.conf                                                                                       100% 5487     5.4KB/s   00:00    
sa.key                                                                                                        100% 1679     1.6KB/s   00:00    
apiserver.crt                                                                                                 100% 1387     1.4KB/s   00:00    
front-proxy-ca.crt                                                                                            100% 1025     1.0KB/s   00:00    
front-proxy-client.crt                                                                                        100% 1050     1.0KB/s   00:00    
apiserver.key                                                                                                 100% 1679     1.6KB/s   00:00    
ca.key                                                                                                        100% 1679     1.6KB/s   00:00    
ca.crt                                                                                                        100% 1025     1.0KB/s   00:00    
apiserver-kubelet-client.crt                                                                                  100% 1099     1.1KB/s   00:00    
apiserver-kubelet-client.key                                                                                  100% 1679     1.6KB/s   00:00    
front-proxy-client.key                                                                                        100% 1675     1.6KB/s   00:00    
front-proxy-ca.key                                                                                            100% 1675     1.6KB/s   00:00    
sa.pub                                                                                                        100%  451     0.4KB/s   00:00    
kube-scheduler.yaml                                                                                           100%  991     1.0KB/s   00:00    
kube-apiserver.yaml                                                                                           100% 2619     2.6KB/s   00:00    
kube-controller-manager.yaml                                                                                  100% 2232     2.2KB/s   00:00    
scheduler.conf                                                                                                100% 5435     5.3KB/s   00:00    
admin.conf                                                                                                    100% 5447     5.3KB/s   00:00    
```

images
```
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-apiserver.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  185M  100  185M    0     0  82.1M      0  0:00:02  0:00:02 --:--:-- 82.1M
0271b8eebde3: Loading layer [==================================================>] 1.338 MB/1.338 MB
f72e92775dd7: Loading layer [==================================================>] 193.2 MB/193.2 MB
Loaded image: gcr.io/google_containers/kube-apiserver:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-controller-manager.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  123M  100  123M    0     0  94.1M      0  0:00:01  0:00:01 --:--:-- 94.1M
4f92df6d8677: Loading layer [==================================================>] 128.2 MB/128.2 MB
Loaded image: gcr.io/google_containers/kube-controller-manager:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-scheduler.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 52.6M  100 52.6M    0     0  94.1M      0 --:--:-- --:--:-- --:--:-- 94.0M
da6851a1e488: Loading layer [==================================================>] 53.85 MB/53.85 MB
Loaded image: gcr.io/google_containers/kube-scheduler:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/linux-bin/k8s-v1.8.4/kubernetes/server/bin/kube-proxy.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 90.6M  100 90.6M    0     0  73.2M      0  0:00:01  0:00:01 --:--:-- 73.3M
08ae86c4c4c9: Loading layer [==================================================>] 42.05 MB/42.05 MB
48a108f012f8: Loading layer [==================================================>] 5.045 MB/5.045 MB
9fc6ccb688b9: Loading layer [==================================================>] 47.93 MB/47.93 MB
Loaded image: gcr.io/google_containers/kube-proxy:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/docker.io0x2Fweaveworks0x2Fweave-kube0x3A2.1.1.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 90.8M  100 90.8M    0     0  10.0M      0  0:00:09  0:00:09 --:--:-- 6349k
5bef08742407: Loading layer [==================================================>] 4.221 MB/4.221 MB
a740a6d19f48: Loading layer [==================================================>] 18.88 MB/18.88 MB
d86a8ab219a9: Loading layer [==================================================>]    27 MB/27 MB
f680fc950fce: Loading layer [==================================================>] 10.26 MB/10.26 MB
4f89cf228c9b: Loading layer [==================================================>] 2.048 kB/2.048 kB
802751b045ac: Loading layer [==================================================>]   277 kB/277 kB
d56c780427c7: Loading layer [==================================================>] 34.62 MB/34.62 MB
Loaded image: weaveworks/weave-kube:2.1.1
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/docker.io0x2Fweaveworks0x2Fweave-npc0x3A2.1.1.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 44.7M  100 44.7M    0     0  41.1M      0  0:00:01  0:00:01 --:--:-- 41.2M
8dccfe2dec8c: Loading layer [==================================================>] 2.811 MB/2.811 MB
11565a7c0730: Loading layer [==================================================>] 39.91 MB/39.91 MB
f56503d36fe6: Loading layer [==================================================>]  2.56 kB/2.56 kB
Loaded image: weaveworks/weave-npc:2.1.1
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fetcd0x2E3.0.17.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  161M  100  161M    0     0  32.5M      0  0:00:04  0:00:04 --:--:-- 33.3M
38ac8d0f5bb3: Loading layer [==================================================>] 1.312 MB/1.312 MB
c872b2c2ac77: Loading layer [==================================================>] 136.7 MB/136.7 MB
be71b2a80bbd: Loading layer [==================================================>] 31.16 MB/31.16 MB
Loaded image: gcr.io/google_containers/etcd:3.0.17
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-dnsmasq-nanny-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 39.7M  100 39.7M    0     0  47.2M      0 --:--:-- --:--:-- --:--:-- 47.1M
e620d0ac6539: Loading layer [==================================================>]  2.56 kB/2.56 kB
9f4f5a30979e: Loading layer [==================================================>]   362 kB/362 kB
fd7319ac31de: Loading layer [==================================================>] 3.072 kB/3.072 kB
b23d166217e1: Loading layer [==================================================>]  37.1 MB/37.1 MB
Loaded image: gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-kube-dns-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 47.3M  100 47.3M    0     0  47.8M      0 --:--:-- --:--:-- --:--:-- 47.8M
a1a7a797fc0e: Loading layer [==================================================>] 45.42 MB/45.42 MB
Loaded image: gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.5
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fk8s-dns-sidecar-amd640x3A1.14.5.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 40.1M  100 40.1M    0     0  43.9M      0 --:--:-- --:--:-- --:--:-- 43.9M
acfc450a47fa: Loading layer [==================================================>] 37.86 MB/37.86 MB
Loaded image: gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.5
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/gcr.io0x2Fgoogle_containers0x2Fpause-amd640x3A3.0.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  747k  100  747k    0     0  7623k      0 --:--:-- --:--:-- --:--:-- 7706k
5f70bf18a086: Loading layer [==================================================>] 1.024 kB/1.024 kB
41ff149e94f2: Loading layer [==================================================>] 748.5 kB/748.5 kB
Loaded image: gcr.io/google_containers/pause-amd64:3.0
ubuntu@kubedev-10-64-33-200:~$ curl http://10.64.33.195:48080/vagrant_f/99-mirror/docker-images/quay.io0x2Fcoreos0x2Fflannel0x3Av0.9.0-amd64.tar | docker load
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 49.7M  100 49.7M    0     0  29.3M      0  0:00:01  0:00:01 --:--:-- 29.4M
f439636ab0f0: Loading layer [==================================================>] 6.797 MB/6.797 MB
91b6f6ead101: Loading layer [==================================================>] 4.414 MB/4.414 MB
39f837629582: Loading layer [==================================================>] 34.49 MB/34.49 MB
d3e99a0118c5: Loading layer [==================================================>] 2.225 MB/2.225 MB
32adca76eade: Loading layer [==================================================>]  5.12 kB/5.12 kB
Loaded image: quay.io/coreos/flannel:v0.9.0-amd64
ubuntu@kubedev-10-64-33-200:~$ docker images
REPOSITORY                                             TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-apiserver                v1.8.4              10a052dccbc5        5 days ago          194 MB
gcr.io/google_containers/kube-controller-manager       v1.8.4              7058ac4d4af5        5 days ago          129 MB
gcr.io/google_containers/kube-proxy                    v1.8.4              65a61c14e8c2        5 days ago          93.2 MB
gcr.io/google_containers/kube-scheduler                v1.8.4              0d985fed7f95        5 days ago          55 MB
weaveworks/weave-npc                                   2.1.1               70a3a81a2ad5        8 days ago          46.6 MB
weaveworks/weave-kube                                  2.1.1               bc65281cfd26        8 days ago          92.6 MB
gcr.io/google_containers/k8s-dns-sidecar-amd64         1.14.5              fed89e8b4248        8 weeks ago         41.8 MB
gcr.io/google_containers/k8s-dns-kube-dns-amd64        1.14.5              512cd7425a73        8 weeks ago         49.4 MB
gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64   1.14.5              459944ce8cc4        8 weeks ago         41.4 MB
quay.io/coreos/flannel                                 v0.9.0-amd64        4c600a64a18a        2 months ago        51.3 MB
gcr.io/google_containers/etcd                          3.0.17              243830dae7dd        9 months ago        169 MB
gcr.io/google_containers/pause-amd64                   3.0                 99e59f495ffa        19 months ago       747 kB
ubuntu@kubedev-10-64-33-200:~$ docker tag gcr.io/google_containers/kube-apiserver:v1.8.4 gcr.io/google_containers/kube-apiserver-amd64:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ docker tag gcr.io/google_containers/kube-controller-manager:v1.8.4 gcr.io/google_containers/kube-controller-manager-amd64:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ docker tag gcr.io/google_containers/kube-scheduler:v1.8.4 gcr.io/google_containers/kube-scheduler-amd64:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ docker tag gcr.io/google_containers/kube-proxy:v1.8.4 gcr.io/google_containers/kube-proxy-amd64:v1.8.4
ubuntu@kubedev-10-64-33-200:~$ docker tag gcr.io/google_containers/etcd:3.0.17 gcr.io/google_containers/etcd-amd64:3.0.17
```

deploy
```
ubuntu@kubedev-10-64-33-200:~$ sudo kubeadm join --token 90d3e2.adc22e3d43b3faa2 10.64.33.82:443 --discovery-token-ca-cert-hash sha256:92202ed5b88a724e1206b95dff886ddce6cc67628886b2eeafb95be4a5d888b6 --skip-preflight-checks 
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[preflight] Skipping pre-flight checks
[discovery] Trying to connect to API Server "10.64.33.82:443"
[discovery] Created cluster-info discovery client, requesting info from "https://10.64.33.82:443"
[discovery] Requesting info from "https://10.64.33.82:443" again to validate TLS against the pinned public key
[discovery] Cluster info signature and contents are valid and TLS certificate validates against pinned roots, will use API Server "10.64.33.82:443"
[discovery] Successfully established connection with API Server "10.64.33.82:443"
[bootstrap] Detected server version: v1.8.4
[bootstrap] The server supports the Certificates API (certificates.k8s.io/v1beta1)

Node join complete:
* Certificate signing request sent to master and response
  received.
* Kubelet informed of new secure connection details.

Run 'kubectl get nodes' on the master to see this machine join.
```

```
ubuntu@kubedev-10-64-33-200:~$ export KUBECONFIG=/home/ubuntu/etc0x2Fkubernetes/admin.conf 
ubuntu@kubedev-10-64-33-200:~$ kubectl get nodes
NAME                   STATUS    ROLES     AGE       VERSION
kubedev-10-64-33-195   Ready     master    7h        v1.8.4
kubedev-10-64-33-199   Ready     master    8h        v1.8.4
kubedev-10-64-33-200   Ready     <none>    57s       v1.8.4
kubedev-10-64-33-82    Ready     master    1h        v1.8.4
```

```
ubuntu@kubedev-10-64-33-200:~$ kubectl get pods -o wide -n kube-system | grep kubedev-10-64-33-200
kube-flannel-ds-r74c6                          0/1       CrashLoopBackOff   5          5m        10.64.33.200   kubedev-10-64-33-200
kube-proxy-nz8dz                               1/1       Running            0          5m        10.64.33.200   kubedev-10-64-33-200
```