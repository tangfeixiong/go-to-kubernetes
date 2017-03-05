

### Install _kubelet_ Daemon

Host

```
[root@localhost 1.3.10]# hostname -I | awk '{print $1}'
192.168.2.20

```
[root@localhost releases%2Fdownload%2Fv1.3.10]# tar -C /opt -xzf kubernetes.tar.gz kubernetes/server/kubernetes-server-linux-amd64.tar.gz

[root@localhost server]# tar -C /opt -zxf kubernetes-server-linux-amd64.tar.gz 

[root@localhost bin]# docker load -i kube-apiserver.tar 
[root@localhost bin]# docker images
REPOSITORY                                TAG                                IMAGE ID            CREATED             SIZE
tangfeixiong/netcat-hello-http            gitrev-7257e93                     bbadd9f66ea9        3 months ago        13.46 MB
gcr.io/google_containers/kube-apiserver   dc8bf1d543f71874b3150cdb1bf60205   febba0bcdbe6        3 months ago        111.1 MB
[root@localhost bin]# docker load -i kube-controller-manager.tar 
[root@localhost bin]# docker images
REPOSITORY                                         TAG                                IMAGE ID            CREATED             SIZE
tangfeixiong/netcat-hello-http                     gitrev-7257e93                     bbadd9f66ea9        3 months ago        13.46 MB
gcr.io/google_containers/kube-apiserver            dc8bf1d543f71874b3150cdb1bf60205   febba0bcdbe6        3 months ago        111.1 MB
gcr.io/google_containers/kube-controller-manager   6322fd6f4db015a29561a4881ff9c372   48e123fa9aa3        3 months ago        101.1 MB
[root@localhost bin]# docker load -i kube-scheduler.tar 
[root@localhost bin]# docker images
REPOSITORY                                         TAG                                IMAGE ID            CREATED             SIZE
tangfeixiong/netcat-hello-http                     gitrev-7257e93                     bbadd9f66ea9        3 months ago        13.46 MB
gcr.io/google_containers/kube-apiserver            dc8bf1d543f71874b3150cdb1bf60205   febba0bcdbe6        3 months ago        111.1 MB
gcr.io/google_containers/kube-scheduler            ec4cd4e378c840237c447cfa147f98a9   1dc259d4b9fb        3 months ago        60.08 MB
gcr.io/google_containers/kube-controller-manager   6322fd6f4db015a29561a4881ff9c372   48e123fa9aa3        3 months ago        101.1 MB
[root@localhost bin]# docker load -i kube-proxy.tar 
[root@localhost docker-images]# docker images
REPOSITORY                                         TAG                                IMAGE ID            CREATED             SIZE
tangfeixiong/netcat-hello-http                     gitrev-7257e93                     bbadd9f66ea9        3 months ago        13.46 MB
gcr.io/google_containers/kube-proxy                4a7d8f39438d3ec6a677f51dcd6b1820   52eb4cdb2eaa        3 months ago        180.1 MB
gcr.io/google_containers/kube-apiserver            dc8bf1d543f71874b3150cdb1bf60205   febba0bcdbe6        3 months ago        111.1 MB
gcr.io/google_containers/kube-scheduler            ec4cd4e378c840237c447cfa147f98a9   1dc259d4b9fb        3 months ago        60.08 MB
gcr.io/google_containers/kube-controller-manager   6322fd6f4db015a29561a4881ff9c372   48e123fa9aa3        3 months ago        101.1 MB
gcr.io/google_containers/pause-amd64               3.0                                99e59f495ffa        9 months ago        746.9 kB


[ecp@localhost ~]$ tar -C /opt -zxf https%3A%2F%2Fgithub.com%2Fkubernetes%2Fkubernetes/releases%2Fdownload%2Fv1.3.10/kubernetes.tar.gz kubernetes/server/kubernetes-manifests.tar.gz


[root@localhost releases%2Fdownload%2Fv1.3.10]# tar -C /opt -zxf kubernetes.tar.gz kubernetes/server/kubernetes-salt.tar.gz

[ecp@localhost ~]$ kubernetes-devops/salt-make-ca-cert.sh 

[ecp@localhost ~]$ sudo cp /opt/kubernetes/server/bin/kubectl ./bin
[ecp@localhost ~]$ sudo chown ecp: ./bin/kubectl 


[ecp@localhost ~]$ kubectl config set-cluster kube --server=https://192.168.2.20:6443
cluster "kube" set.
[ecp@localhost ~]$ encoded=$(base64 -w0 kubernetes-devops/cacerts/ca.crt); kubectl config set clusters.kube.certificate-authority-data $encoded
property "clusters.kube.certificate-authority-data" set.
[ecp@localhost ~]$ kubectl config set-credentials admin
user "admin" set.
[ecp@localhost ~]$ encoded=$(base64 -w0 kubernetes-devops/cacerts/kubecfg.crt); kubectl config set users.admin.client-certificate-data $encoded
property "users.admin.client-certificate-data" set.
[ecp@localhost ~]$ encoded=$(base64 -w0 kubernetes-devops/cacerts/kubecfg.key); kubectl config set users.admin.client-key-data $encoded
property "users.admin.client-key-data" set.
[ecp@localhost ~]$ kubectl config set-context kube-admin --cluster=kube --user=admin
context "kube-admin" set.
[ecp@localhost ~]$ kubectl config use-context kube-admin
switched to context "kube-admin".
[ecp@localhost ~]$ cat .kube/config 
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURXVENDQWtHZ0F3SUJBZ0lKQUo4TmVrOC85NFV6TUEwR0NTcUdTSWIzRFFFQkN3VUFNQ0l4SURBZUJnTlYKQkFNVUZ6RTVNaTR4TmpndU1pNHlNRUF4TkRnM056VXlNalU0TUI0WERURTNNREl5TWpBNE16QTFPRm9YRFRJMwpNREl5TURBNE16QTFPRm93SWpFZ01CNEdBMVVFQXhRWE1Ua3lMakUyT0M0eUxqSXdRREUwT0RjM05USXlOVGd3CmdnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUUNjTmtyQXhvYlVleVFNNWdndHBsaW8KYWlBdEpVdXJpWS82Sm5WMGdLQ0pGMU5iTHI0djRjRFd3bklsb2V2NVBIWTJ1UEVNNU9UT3Y3RENGNm15S1JNWApvVmVPMDJJY2JKRlIrUDNOMkp1eVZzSm5lZzRoWjcyZ1J3Y28vWkpVT3NtM0Zpa2lKRCszSmIyS2NhSVhUZGo5CndmVUJ2a3Nya2c5cjMxTzR6dkJkcFRCY1hzQ1hQaWpkQ1djZWFkTHRCeFYzSGZHclZYNXE1N0pWVXRzM2ZGbUcKU1E0YkVJS3FkaUluSitkTHR2dEFQTENONmlNUWlFYzU1b3VOemdDY0ZNVE9sSnBVa0VJaXg2bDJ6a3NuQUpUNwpXWFJVb1B5Yi9qMEh2Q0gzQ0orOGFqQlVOVFE3dzlkelQyLzkwUkhmVVJuMWlJU1lJVTBsMVhwS1htS2xnL01KCkFnTUJBQUdqZ1pFd2dZNHdIUVlEVlIwT0JCWUVGTTJqamxCNkxGeitQTnArcitVNEtmL3hLbmFYTUZJR0ExVWQKSXdSTE1FbUFGTTJqamxCNkxGeitQTnArcitVNEtmL3hLbmFYb1Nha0pEQWlNU0F3SGdZRFZRUURGQmN4T1RJdQpNVFk0TGpJdU1qQkFNVFE0TnpjMU1qSTFPSUlKQUo4TmVrOC85NFV6TUF3R0ExVWRFd1FGTUFNQkFmOHdDd1lEClZSMFBCQVFEQWdFR01BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQnlUZDFIRnBZVGdoZ29QL1FPR0NvSVNxVk4KZ1dZZVBZam91MUU2RmlhL2lNcTl4TkZYWE8wbEtJUk5qeFFSZFZrQjZ0eUJuUjlsYmhQUzl4dm5KUmxNMFlXKwpnc0ZwbFFjcVBaQ2dIN3YvTWw5WE1MbE9EVHgyRzlVMXltVUxHeFVtK0JGMHIrTjE3ekhyVnZFc1IwZ2xhOXBKCk5tTXY0MitZQlFmTDRaL0RNa1dFNSsrVlR2emgvckI1djlJa2doNk9vL3VxVFM1a2VrbkxIRWV0eTJ1ckNUZzgKU0VDdXlSL042KzQ1YjBndU05Um9WWlRtVlIwMzJNSVdJM2dua0lnMm5Oa1pzUzZVZ1liNHNCSHdBMjNlaUdTYwpqd1FUc05FOE82Z004cnhnWFVseVhUT3Y5RUt4YVdBb2x1bVFROGI3a0RPS1BWRk05bjRmSEZjWUFGNDUKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    server: https://192.168.2.20:6443
  name: kube
contexts:
- context:
    cluster: kube
    user: admin
  name: kube-admin
current-context: kube-admin
kind: Config
preferences: {}
users:
- name: admin
  user:
    client-certificate-data: Q2VydGlmaWNhdGU6CiAgICBEYXRhOgogICAgICAgIFZlcnNpb246IDMgKDB4MikKICAgICAgICBTZXJpYWwgTnVtYmVyOiAyICgweDIpCiAgICBTaWduYXR1cmUgQWxnb3JpdGhtOiBzaGEyNTZXaXRoUlNBRW5jcnlwdGlvbgogICAgICAgIElzc3VlcjogQ049MTkyLjE2OC4yLjIwQDE0ODc3NTIyNTgKICAgICAgICBWYWxpZGl0eQogICAgICAgICAgICBOb3QgQmVmb3JlOiBGZWIgMjIgMDg6MzA6NTggMjAxNyBHTVQKICAgICAgICAgICAgTm90IEFmdGVyIDogRmViIDIwIDA4OjMwOjU4IDIwMjcgR01UCiAgICAgICAgU3ViamVjdDogQ049a3ViZWNmZwogICAgICAgIFN1YmplY3QgUHVibGljIEtleSBJbmZvOgogICAgICAgICAgICBQdWJsaWMgS2V5IEFsZ29yaXRobTogcnNhRW5jcnlwdGlvbgogICAgICAgICAgICAgICAgUHVibGljLUtleTogKDIwNDggYml0KQogICAgICAgICAgICAgICAgTW9kdWx1czoKICAgICAgICAgICAgICAgICAgICAwMDpjNzo0ZTo4Mzo1ODpiZDpkZjpmZDo3MjplNDozNjo2NzowNDo3NjowYToKICAgICAgICAgICAgICAgICAgICBiNTo0NzozYzpjYzoyYTpmMzpmMjphNDpmYzo3Mzo3NTo5MzozZDozNToyMDoKICAgICAgICAgICAgICAgICAgICA2NTo5MjpkYjo0NToxMzoyNDoxNjo3MTpkNjpkODo2Njo4OTo2ZDpmNTowMToKICAgICAgICAgICAgICAgICAgICA3ODpjNDoxMzo2MDo3MTpjMjo0OTo2ODo2Yzo2ODo5Mzo3OTpiMToxNDpkYjoKICAgICAgICAgICAgICAgICAgICBhODpjMTo4ZjpmNTphZDo1ZDowNzpkYTo2YTo2YTpkMjo2MTpmMDpmOToxODoKICAgICAgICAgICAgICAgICAgICAzOToxYjozNjoyNDoyOTowOTphMjo4OTo4OTplZTo3YjowYjoxMjpkNDo1MzoKICAgICAgICAgICAgICAgICAgICBhNDowNDo3ZTo1NTo2OTplYzphOTplNDo2MzpiOTo5MDowNDpiZDplMzpmNzoKICAgICAgICAgICAgICAgICAgICBiZjoxOTplODozNDpmNzo3NDo3MTpmMzo2OToxYTo5Nzo5NDo0ODpkYzo5MToKICAgICAgICAgICAgICAgICAgICAyMzo2ODpmZDplZDo5MDphNjphODpjNDpjYzoxYjpmZTo3MzpmODpiZTpjMzoKICAgICAgICAgICAgICAgICAgICBhZTo4MDpmYzo1NzpkODo5YjphYjoxMzo1YTplMzo4ZjpiYjo5YzpkYTozMjoKICAgICAgICAgICAgICAgICAgICBjYjpiZDplMzo4NjozZDpkNzo0YTo1ZDo2NTpmMTo5YzplYTo0ODo4NjowNToKICAgICAgICAgICAgICAgICAgICA0MDphOTo4MDo0MzplMzoxNTplYzozMzo0Njo4ZjplOTo3MDo0Njo1ZjpjYToKICAgICAgICAgICAgICAgICAgICAxYToyZDo3NTplMjpkMzo0YzplNzoyYjpiYTpiYzphMjowODoyNzo0Zjo3OToKICAgICAgICAgICAgICAgICAgICA5ODo5YTpiYTo3MzpiMDowMDoyNTpkNzo3NTozMjowNDpjOTowZDo3OTozMjoKICAgICAgICAgICAgICAgICAgICA2Mjo1Zjo4ZTozYTozMTo0ZTpiMzpmYTplZTozNjo5ZTo0ODo4MzpmNDoxMDoKICAgICAgICAgICAgICAgICAgICA3Yjo1OTozMzo1YjplNjo4YTo2Mzo4MTpiOTo4MjpiYzphYTo4YTo5MTphYToKICAgICAgICAgICAgICAgICAgICBjNjo3Mzo5NDo5ZjplOTplMTpiYjpiNTowNjo5MDo1MzplZjo0MDowOTo3YzoKICAgICAgICAgICAgICAgICAgICA3ZToyNQogICAgICAgICAgICAgICAgRXhwb25lbnQ6IDY1NTM3ICgweDEwMDAxKQogICAgICAgIFg1MDl2MyBleHRlbnNpb25zOgogICAgICAgICAgICBYNTA5djMgQmFzaWMgQ29uc3RyYWludHM6IAogICAgICAgICAgICAgICAgQ0E6RkFMU0UKICAgICAgICAgICAgWDUwOXYzIFN1YmplY3QgS2V5IElkZW50aWZpZXI6IAogICAgICAgICAgICAgICAgMkM6RTM6NDY6RUI6OUU6MzQ6QzA6QzI6NzM6RTE6RDA6RDY6RDM6NDY6QTM6MUM6OTU6Njc6RDE6NjQKICAgICAgICAgICAgWDUwOXYzIEF1dGhvcml0eSBLZXkgSWRlbnRpZmllcjogCiAgICAgICAgICAgICAgICBrZXlpZDpDRDpBMzo4RTo1MDo3QToyQzo1QzpGRTozQzpEQTo3RTpBRjpFNTozODoyOTpGRjpGMToyQTo3Njo5NwogICAgICAgICAgICAgICAgRGlyTmFtZTovQ049MTkyLjE2OC4yLjIwQDE0ODc3NTIyNTgKICAgICAgICAgICAgICAgIHNlcmlhbDo5RjowRDo3QTo0RjozRjpGNzo4NTozMwoKICAgICAgICAgICAgWDUwOXYzIEV4dGVuZGVkIEtleSBVc2FnZTogCiAgICAgICAgICAgICAgICBUTFMgV2ViIENsaWVudCBBdXRoZW50aWNhdGlvbgogICAgICAgICAgICBYNTA5djMgS2V5IFVzYWdlOiAKICAgICAgICAgICAgICAgIERpZ2l0YWwgU2lnbmF0dXJlCiAgICBTaWduYXR1cmUgQWxnb3JpdGhtOiBzaGEyNTZXaXRoUlNBRW5jcnlwdGlvbgogICAgICAgICAwNDoyYjo1OTplNzpjZjoyNjo5NjowMTplZDozNzowZTozMTo0Yzo4MDpjOToyNzphYzo2NToKICAgICAgICAgN2U6MDU6Mzc6NjA6Nzg6MzQ6ZmI6Njk6ODQ6M2Y6Zjk6YzE6ZmE6M2E6MmE6MWE6YzE6Mjg6CiAgICAgICAgIDM3OjhkOmVmOjk2OjgyOjlmOmEyOjc5OjQ2OjcyOjAxOjg2OmQ5OmFlOjM1OjBiOjgwOmQwOgogICAgICAgICBkNDpkMzo4OToxYjpiZjpkMDphMzo4NTpkMDo1YjphNzpiMjoxZjpiMjo5NTpiMToxYTo1NzoKICAgICAgICAgN2U6NTU6Y2Y6MDI6Yjk6NjE6NTM6NDQ6MzE6YzE6ODM6MTg6M2E6ODY6ZWM6Nzg6NGU6NTI6CiAgICAgICAgIGFhOmI5OjUwOjg5Ojg4OmM5OjM3OmJlOmExOjJlOmY2OjIzOjgwOjFlOmZjOmE3OjRhOjk4OgogICAgICAgICA5MDpkOTpiMTo1MjozNDo5ZjpiODozYzo2NTpjZTpjZTpmYzpkMzpiODoyMDo1NDo4ODo2ZjoKICAgICAgICAgNmU6MTc6NDM6NmI6ZTE6MmU6ZWE6ZDc6ZjM6N2U6ZTQ6YWI6NWI6YzY6NzU6ZWE6ZTE6MmY6CiAgICAgICAgIGY1OjE2OjhhOmVlOmRlOjIwOjMxOmE3OjVhOjBlOjZkOjFiOjdiOjQ2OjUzOmUyOjg5OjVjOgogICAgICAgICBiZjoxMDphZDpjNzpiOTpkYjo5NzoyMzo0Njo2ZDoxYzphZTo1ZDo5Yzo2ZTo0YzpkYzozNDoKICAgICAgICAgM2E6Zjc6ZmE6ZjI6ZTU6YjE6N2Q6YzI6MmQ6MmQ6ZDM6Njk6ODQ6MDA6ZmQ6ZmE6ZTA6ZmI6CiAgICAgICAgIDczOjQ0OmU3OjdlOjI4OjJkOmNlOjczOjJlOmFjOjljOjhhOjJhOmRmOjA1OjNhOmVhOmRhOgogICAgICAgICA5Njo1MDo3NzoxMTo0ODpmODoyZDo0ZDoyYjphNTo5MTo3ZjpiYjoxZTphNToxNDpjNTpiNToKICAgICAgICAgZWI6YTk6ZDk6ZDk6YWY6NDY6OTg6NGM6Mjg6ZGU6YTY6YWI6ZjE6Yzk6OTU6OGI6MDM6ODA6CiAgICAgICAgIDdkOmE2OjI4OmIwCi0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLQpNSUlEVXpDQ0FqdWdBd0lCQWdJQkFqQU5CZ2txaGtpRzl3MEJBUXNGQURBaU1TQXdIZ1lEVlFRREZCY3hPVEl1Ck1UWTRMakl1TWpCQU1UUTROemMxTWpJMU9EQWVGdzB4TnpBeU1qSXdPRE13TlRoYUZ3MHlOekF5TWpBd09ETXcKTlRoYU1CSXhFREFPQmdOVkJBTVRCMnQxWW1WalptY3dnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBdwpnZ0VLQW9JQkFRREhUb05ZdmQvOWN1UTJad1IyQ3JWSFBNd3E4L0trL0hOMWt6MDFJR1dTMjBVVEpCWngxdGhtCmlXMzFBWGpFRTJCeHdrbG9iR2lUZWJFVTI2akJqL1d0WFFmYWFtclNZZkQ1R0RrYk5pUXBDYUtKaWU1N0N4TFUKVTZRRWZsVnA3S25rWTdtUUJMM2o5NzhaNkRUM2RISHphUnFYbEVqY2tTTm8vZTJRcHFqRXpCditjL2krdzY2QQovRmZZbTZzVFd1T1B1NXphTXN1OTQ0WTkxMHBkWmZHYzZraUdCVUNwZ0VQakZld3pSby9wY0VaZnlob3RkZUxUClRPY3J1cnlpQ0NkUGVaaWF1bk93QUNYWGRUSUV5UTE1TW1KZmpqb3hUclA2N2phZVNJUDBFSHRaTTF2bWltT0IKdVlLOHFvcVJxc1p6bEovcDRidTFCcEJUNzBBSmZINGxBZ01CQUFHamdhTXdnYUF3Q1FZRFZSMFRCQUl3QURBZApCZ05WSFE0RUZnUVVMT05HNjU0MHdNSno0ZERXMDBhakhKVm4wV1F3VWdZRFZSMGpCRXN3U1lBVXphT09VSG9zClhQNDgybjZ2NVRncC8vRXFkcGVoSnFRa01DSXhJREFlQmdOVkJBTVVGekU1TWk0eE5qZ3VNaTR5TUVBeE5EZzMKTnpVeU1qVTRnZ2tBbncxNlR6LzNoVE13RXdZRFZSMGxCQXd3Q2dZSUt3WUJCUVVIQXdJd0N3WURWUjBQQkFRRApBZ2VBTUEwR0NTcUdTSWIzRFFFQkN3VUFBNElCQVFBRUsxbm56eWFXQWUwM0RqRk1nTWtuckdWK0JUZGdlRFQ3CmFZUS8rY0g2T2lvYXdTZzNqZStXZ3AraWVVWnlBWWJacmpVTGdORFUwNGtidjlDamhkQmJwN0lmc3BXeEdsZCsKVmM4Q3VXRlRSREhCZ3hnNmh1eDRUbEtxdVZDSmlNazN2cUV1OWlPQUh2eW5TcGlRMmJGU05KKzRQR1hPenZ6VAp1Q0JVaUc5dUYwTnI0UzdxMS9OKzVLdGJ4blhxNFMvMUZvcnUzaUF4cDFvT2JSdDdSbFBpaVZ5L0VLM0h1ZHVYCkkwWnRISzVkbkc1TTNEUTY5L3J5NWJGOXdpMHQwMm1FQVAzNjRQdHpST2QrS0MzT2N5NnNuSW9xM3dVNjZ0cVcKVUhjUlNQZ3RUU3Vsa1grN0hxVVV4YlhycWRuWnIwYVlUQ2plcHF2eHlaV0xBNEI5cGlpdwotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    client-key-data: LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2QUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktZd2dnU2lBZ0VBQW9JQkFRREhUb05ZdmQvOWN1UTIKWndSMkNyVkhQTXdxOC9Lay9ITjFrejAxSUdXUzIwVVRKQlp4MXRobWlXMzFBWGpFRTJCeHdrbG9iR2lUZWJFVQoyNmpCai9XdFhRZmFhbXJTWWZENUdEa2JOaVFwQ2FLSmllNTdDeExVVTZRRWZsVnA3S25rWTdtUUJMM2o5NzhaCjZEVDNkSEh6YVJxWGxFamNrU05vL2UyUXBxakV6QnYrYy9pK3c2NkEvRmZZbTZzVFd1T1B1NXphTXN1OTQ0WTkKMTBwZFpmR2M2a2lHQlVDcGdFUGpGZXd6Um8vcGNFWmZ5aG90ZGVMVFRPY3J1cnlpQ0NkUGVaaWF1bk93QUNYWApkVElFeVExNU1tSmZqam94VHJQNjdqYWVTSVAwRUh0Wk0xdm1pbU9CdVlLOHFvcVJxc1p6bEovcDRidTFCcEJUCjcwQUpmSDRsQWdNQkFBRUNnZ0VBV2tMRVdUd2taTFBUaG8vQkJYUjBCSDhVTjUrakUvVUFsUTdKZVhYaWlrK2oKU0xmZk1rOEtVZVdMVnBvRUIwaC9OUVc0R1FKL21jcFMyQTlpcmNaMGpMN0UvN1dTQ3hVdi9KOXJ2VmNYcVQ0agp6Z1laYXIvcEZ4K0orRDcrajZvT2h0QmpDVVljZVZHSnlrSFBPd0lSV1JzUVgvVlh5Z3d3N1N1U214L2U3c0IwClhoTmZDZ2lZYVRud3BBM3RjRkxWZ2ZMcnpHZVcyZHFhcFZhSlVkVnQyZFNHdzNBWm1WSzd4YkhjODl6S1FEZzUKbDJYSzh6ekxaNGtOSWZGdW9OTGQrNUVGelZRdzNMOU0xV0ZZRU5pVHhYZzhJUnpuejdvd3lzTlkrQ25SbUo2QgpTYUdHOHp4eFBSOHA4NUpsMzBpTVArYUJkVkRBSXBkblFJdmtEQzRLQVFLQmdRRHNjZkpKOGhBVDlCWkcybXdjCitxVUxVZlRML2g5R2pTWkVPZTFrNytjVU9RU3NsczRjOVVXREo4eVlkNkduL0NhVS9nTExNTlBXeXBPMEcvR0UKR0VEYi9qNGtUeVpHUjlpQlo5MUQ3RmpiTlNCSHdXU1NtcTFCNXgxTWcxaDI4UURPS1QwR3h0Vm5VckJzZHFsSQpqMlhENytjT0x1TWNLSWhGSXVIaVNTNkc2UUtCZ1FEWHlrUW9SRWxSaDZKYnp5TmwwR1Nsc0ttVUFVNzQ5dWNJCjRhMlpjOXJIZENrNkJSbGJHZFZHWXRocWRGbHU3QVpDK3h1Y0RkRWVXZmhkbkVteTdmVWZUcDVrY0ZjT2ZFMWYKaCtYczFhTmJSVFdPdk53b0diZnZucDBoV0pvZ2xuTTdIKzNvdjZBKzVWWlBHNnN5ODZlVmppU0VwKzVLcFl1WQpaL09GQTF0djNRS0JnQTQ2clVJVjl6RzhOcnI4MmhURVlMcFZyRTI2ZFZkaGx6UHo2SDF5a28xOUFQTUNBRlZ0CnpVa2ZWQWk0NmxtWnF4aThVSjY1eVlYYm14Znp0bzBraUlLMHdqaWRLc2gvK2wrUFNKbVYxNk1aaDEvS0c2dFcKODZaTHFrKzRkOUp3bm9QSnA0cUkzTDdqRXdyeitTd010U3NkZk9JMW9DVkpxeW05OTZvT2pXWkJBb0dBWVpXSgpydi82ZkJFdlBmOTFUWVNwRHpqTjV5MlVUZ2d4d0pWMkVSQVFYdVJNQkgzcjBvUGpLOXBsYmZiQnZ1U2pqRVRNClhvenRCYUhBTEcwUjh3V2pOUTR6bU00b3dGYzhFamg5cE5XVDh5RmdML1YrUmZBamV3d3FtTHJkc0dENUtVS3UKTHZRQmZvL3RzUWFkTEFSOXc2Y3RJb1Jpd1lVTmxOVmxqY0JQNkUwQ2dZQVd4S2VCaW1BWS8yc1dkMWtxNElUTApoN3l0QU5Vb3liWGc2YW5Zb0gyM0s3U3hVTExRQ0gwNjhYY2J5MnFUQXpSb3dJcDB1NmVMMmJiR3M0eGtBdnprCkszd2lYYmpIRmw3Q0t3WHgwQkpabnFGRUxrbGNGeTVRYmprcWw5NDB4MThjSmVkMGNSM0pYYUFSb3AzLzVQN3AKNENVSmI1THFwbHB6Nm5KU2FFL1pjdz09Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K


[root@localhost 1.3.10]# vi etc%2Fsysconfig/kubelet

Service file

    [vagrant@localhost ~]$ vi kubernetes/1.3.10/centos/systemd/system/kubelet.service
    [Unit]
    Description=Kubernetes container mgmt daemon - kubelet
    Documentation=https://github.com/kubernetes/kubernetes
    After=docker.service
    Requires=docker.service
    # Wants=
    Conflicts=openshift-node.service

    [Service]
    # Type=notify
    # NotifyAccess=all
    Type=simple
    User=root

    Environment=KUBERNETES_BUILD_VERSION=1.3.10
    EnvironmentFile=/etc/sysconfig/kubelet
    WorkingDirectory=/opt/kubernetes

    ExecStartPre=/usr/bin/mkdir -p /etc/kubernetes/manifests
    ExecStartPre=/usr/bin/mkdir -p /var/log/containers

    ExecStart=/opt/kubernetes/server/bin/kubelet $KUBELET_OPTS
    # ExecStart=/bin/bash -c "/opt/kubernetes/${KUBERNETES_BUILD_VERSION}/centos/bin/kubelet ${KUBELET_OPTS}"

    # Restart=always
    Restart=on-failure
    RestartSec=10
    KillMode=process

    [Install]
    WantedBy=multi-user.target

Opts file

    [vagrant@localhost ~]$ vi kubernetes/1.3.10/centos/systemd/conf/kubelet
    # KUBELET_VERSION=1.3.10

    # --api-servers=[]: Comma separated list of Kubernetes API servers (ip:port).
    # API_SERVERS=https://10.64.33.81:6443

    # --hostname-override="": As identification instad of the actual hostname
    # HOSTNAME_OVERRIDE=10.64.33.81

    # --cert-dir="/var/run/kubernetes"
    # CERT_DIR=/srv/kubernetes

    # CONFIG_PATH=/etc/kubernetes/manifests

    # --kubeconfig="/var/lib/kubelet/kubeconfig": Path to a kubeconfig file
    # KUBECONFIG_PATH=/srv/kubernetes/kubeconfig

    # --master-service-namespace="default": The master services should be injected into pods

    # --service-cluster-ip-range=<nil>: A CIDR notation IP range
    # Same as GKE, cluster CIDR: 10.120.0.0/14, service CIDR: 10.123.240.0/20
    # CLUSTER_CIDR=10.120.0.0/14
    # SERVICE_CIDR=10.123.240.0/20
    # CLUSTER_MASTER=10.123.240.1
    # CLUSTER_DNS=10.123.240.10

    KUBELET_OPTS="--address=10.64.33.81 \
      --allow-privileged=true \
      --api-servers=https://10.64.33.81:6443 \
      --cadvisor-port=4194 \
      --cert-dir=/srv/kubernetes \
      --cluster-dns=10.123.240.10 \
      --cluster-domain=cluster.local \
      --config=/etc/kubernetes/manifests \
      --configure-cbr0=false \
      --container-runtime=docker \
      --docker=unix:///var/run/docker.sock \
      --healthz-bind-address=127.0.0.1 \
      --healthz-port=10248 \
      --hostname-override=10.64.33.81 \
      --kubeconfig=/srv/kubernetes/kubeconfig \
      --master-service-namespace=default \
      --max-open-files=1000000 \
      --max-pods=110 \
      --node-ip=10.64.33.81 \
      --non-masquerade-cidr=10.0.0.0/8 \
      --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 \
      --pods-per-core=0 \
      --port=10250 \
      --read-only-port=10255 \
      --root-dir=/var/lib/kubelet \
      --tls-cert-file=/srv/kubernetes/server.cert \
      --tls-private-key-file=/srv/kubernetes/server.key \
      --v=2"

Manually install

[root@localhost kubernetes]# systemctl -l status firewalld.service
● firewalld.service - firewalld - dynamic firewall daemon
   Loaded: loaded (/usr/lib/systemd/system/firewalld.service; enabled; vendor preset: enabled)
   Active: active (running) since 二 2017-02-21 11:23:54 CST; 1 day 6h ago
     Docs: man:firewalld(1)
 Main PID: 658 (firewalld)
   Memory: 0B
   CGroup: /system.slice/firewalld.service
           └─658 /usr/bin/python -Es /usr/sbin/firewalld --nofork --nopid

2月 21 17:40:06 localhost.localdomain firewalld[658]: WARNING: COMMAND_FAILED: '/usr/sbin/iptables -w2 -t nat -C POSTROUTING -s 172.17.0.0/16 ! -o docker0 -j MASQUERADE' failed:
2月 21 17:40:06 localhost.localdomain firewalld[658]: WARNING: COMMAND_FAILED: '/usr/sbin/iptables -w2 -t nat -C DOCKER -i docker0 -j RETURN' failed:
2月 21 17:40:06 localhost.localdomain firewalld[658]: WARNING: COMMAND_FAILED: '/usr/sbin/iptables -w2 -D FORWARD -i docker0 -o docker0 -j DROP' failed:
2月 21 17:40:06 localhost.localdomain firewalld[658]: WARNING: COMMAND_FAILED: '/usr/sbin/iptables -w2 -t filter -C FORWARD -i docker0 -o docker0 -j ACCEPT' failed:
2月 21 17:40:06 localhost.localdomain firewalld[658]: WARNING: COMMAND_FAILED: '/usr/sbin/iptables -w2 -t filter -C FORWARD -i docker0 ! -o docker0 -j ACCEPT' failed:
2月 21 17:40:06 localhost.localdomain firewalld[658]: WARNING: COMMAND_FAILED: '/usr/sbin/iptables -w2 -t filter -C FORWARD -o docker0 -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT' failed:
2月 21 17:40:06 localhost.localdomain firewalld[658]: WARNING: COMMAND_FAILED: '/usr/sbin/iptables -w2 -t nat -C PREROUTING -m addrtype --dst-type LOCAL -j DOCKER' failed:
2月 21 17:40:06 localhost.localdomain firewalld[658]: WARNING: COMMAND_FAILED: '/usr/sbin/iptables -w2 -t nat -C OUTPUT -m addrtype --dst-type LOCAL -j DOCKER ! --dst 127.0.0.0/8' failed:
2月 21 17:40:06 localhost.localdomain firewalld[658]: WARNING: COMMAND_FAILED: '/usr/sbin/iptables -w2 -t filter -C FORWARD -o docker0 -j DOCKER' failed:
2月 21 17:40:06 localhost.localdomain firewalld[658]: WARNING: COMMAND_FAILED: '/usr/sbin/iptables -w2 -t filter -C FORWARD -j DOCKER-ISOLATION' failed:

[root@localhost kubernetes]# systemctl stop firewalld.service
[root@localhost kubernetes]# systemctl disable firewalld.service
Removed symlink /etc/systemd/system/dbus-org.fedoraproject.FirewallD1.service.
Removed symlink /etc/systemd/system/basic.target.wants/firewalld.service.
[root@localhost kubernetes]# setenforce 0
[root@localhost kubernetes]# vi /etc/sysconfig/selinux 


    [vagrant@localhost ~]$ sudo cp kubernetes/1.3.10/centos/systemd/conf/kubelet /etc/sysconfig/

    [vagrant@localhost ~]$ sudo cp kubernetes/1.3.10/centos/systemd/system/kubelet.service /etc/systemd/system

    [vagrant@localhost ~]$ sudo systemctl enable kubelet.service
    Created symlink from /etc/systemd/system/multi-user.target.wants/kubelet.service to /etc/systemd/system/kubelet.service.

[root@localhost kubernetes]# touch /var/log/containers/kube-controller-manager.log



Experimental run

    [vagrant@localhost ~]$ KUBELET_OPTS="--address=0.0.0.0 \
    >       --allow-privileged=true \
    >       --api-servers=https://10.64.33.81:6443 \
    >       --cadvisor-port=4194 \
    >       --cert-dir=/srv/kubernetes \
    >       --cluster-dns=10.123.240.10 \
    >       --cluster-domain=cluster.local \
    >       --config=/opt/kubernetes/manifests \
    >       --configure-cbr0=false \
    >       --container-runtime=docker \
    >       --docker=unix:///var/run/docker.sock \
    >       --experimental-nvidia-gpus=0 \
    >       --healthz-bind-address=127.0.0.1 \
    >       --healthz-port=10248 \
    >       --hostname-override=10.64.33.81 \
    >       --kubeconfig=/srv/kubernetes/kubeconfig \
    >       --master-service-namespace=default \
    >       --max-open-files=1000000 \
    >       --max-pods=110 \
    >       --node-ip=10.64.33.81 \
    >       --non-masquerade-cidr=10.0.0.0/8 \
    >       --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 \
    >       --pods-per-core=0 \
    >       --port=10250 \
    >       --root-dir=/var/lib/kubelet \
    >       --tls-cert-file=/srv/kubernetes/server.cert \
    >       --tls-private-key-file=/srv/kubernetes/server.key \
    >       --v=2"; sudo /opt/kubernetes/server/bin/kubelet $KUBELET_OPTS

First start

    [vagrant@localhost ~]$ sudo systemctl daemon-reload                             

    [vagrant@localhost ~]$ sudo systemctl start kubelet.service

Validation

    [vagrant@localhost ~]$ sudo systemctl -l status kubelet.service
    鈼▒ kubelet.service - Kubernetes container mgmt daemon - kubelet
       Loaded: loaded (/etc/systemd/system/kubelet.service; enabled; vendor preset: disabled)
       Active: active (running) since Sat 2016-11-12 15:11:08 UTC; 40min ago
         Docs: https://github.com/kubernetes/kubernetes
     Main PID: 17657 (kubelet)
       CGroup: /system.slice/kubelet.service
               鈹溾攢17657 /opt/kubernetes/server/bin/kubelet --address=10.64.33.81 --allow-privileged=true --api-servers=https://10.64.33.81:6443 --cadvisor-port=4194 --cert-dir=/srv/kubernetes --cluster-dns=10.123.240.10 --cluster-domain=cluster.local --config=/etc/kubernetes/manifests --configure-cbr0=false --container-runtime=docker --docker=unix:///var/run/docker.sock --healthz-bind-address=127.0.0.1 --healthz-port=10248 --hostname-override=10.64.33.81 --kubeconfig=/srv/kubernetes/kubeconfig --master-service-namespace=default --max-open-files=1000000 --max-pods=0 --node-ip=10.64.33.81 --non-masquerade-cidr=10.0.0.0/8 --pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.0 --pods-per-core=0 --port=10250 --read-only-port=10255 --root-dir=/var/lib/kubelet --tls-cert-file=/srv/kubernetes/server.cert --tls-private-key-file=/srv/kubernetes/server.key --v=2

    Nov 12 15:46:09 localhost.localdomain kubelet[17657]: W1112 15:46:09.340116   17657 container_manager_linux.go:589] CPUAccounting not enabled for pid: 1076
    Nov 12 15:46:09 localhost.localdomain kubelet[17657]: W1112 15:46:09.340156   17657 container_manager_linux.go:592] MemoryAccounting not enabled for pid: 1076
    Nov 12 15:46:09 localhost.localdomain kubelet[17657]: I1112 15:46:09.340171   17657 container_manager_linux.go:290] Discovered runtime cgroups name: /system.slice/docker.service
    Nov 12 15:46:09 localhost.localdomain kubelet[17657]: W1112 15:46:09.340318   17657 container_manager_linux.go:589] CPUAccounting not enabled for pid: 17657
    Nov 12 15:46:09 localhost.localdomain kubelet[17657]: W1112 15:46:09.340334   17657 container_manager_linux.go:592] MemoryAccounting not enabled for pid: 17657
    Nov 12 15:51:09 localhost.localdomain kubelet[17657]: W1112 15:51:09.347243   17657 container_manager_linux.go:589] CPUAccounting not enabled for pid: 1076
    Nov 12 15:51:09 localhost.localdomain kubelet[17657]: W1112 15:51:09.347277   17657 container_manager_linux.go:592] MemoryAccounting not enabled for pid: 1076
    Nov 12 15:51:09 localhost.localdomain kubelet[17657]: I1112 15:51:09.347292   17657 container_manager_linux.go:290] Discovered runtime cgroups name: /system.slice/docker.service
    Nov 12 15:51:09 localhost.localdomain kubelet[17657]: W1112 15:51:09.354968   17657 container_manager_linux.go:589] CPUAccounting not enabled for pid: 17657
    Nov 12 15:51:09 localhost.localdomain kubelet[17657]: W1112 15:51:09.355003   17657 container_manager_linux.go:592] MemoryAccounting not enabled for pid: 17657

    [vagrant@localhost ~]$ sudo journalctl --follow --no-tail --no-pager --pager-end -u kubelet.service

    [vagrant@localhost ~]$ sudo tail -100 /var/log/messages | grep kubelet

    [vagrant@localhost ~]$ sudo journalctl -k -f


[root@localhost manifests]# cat kube-apiserver.yaml 
apiVersion: v1
kind: Pod
metadata:
  name: kube-apiserver
  namespace: kube-system
spec:
  containers:
  - command:
    - /bin/sh
    - -c
    - /usr/local/bin/kube-apiserver
      --address=0.0.0.0
      --etcd-servers=http://192.168.2.20:2379
      --etcd-servers-overrides=/events#http://192.168.2.20:2379
      --admission-control=NamespaceLifecycle,NamespaceExists,LimitRanger,ServiceAccount,SecurityContextDeny,ResourceQuota
      --service-cluster-ip-range=10.123.240.0/20
      --client-ca-file=/srv/kubernetes/ca.crt
      --basic-auth-file=/srv/kubernetes/basic_auth.csv
      --tls-cert-file=/srv/kubernetes/server.cert
      --tls-private-key-file=/srv/kubernetes/server.key
      --secure-port=6443
      --token-auth-file=/srv/kubernetes/known_tokens.csv
      --v=2
      --allow-privileged=True
      1>>/var/log/kube-apiserver.log 2>&1
    image: gcr.io/google_containers/kube-apiserver:dc8bf1d543f71874b3150cdb1bf60205
    imagePullPolicy: Never # Always IfNotPresent
    livenessProbe:
      httpGet:
        host: 127.0.0.1
        path: /healthz
        port: 8080
        scheme: HTTP
      initialDelaySeconds: 15
      timeoutSeconds: 15
    name: kube-apiserver
    ports:
    - containerPort: 6443
      hostPort: 6443
      name: https
      protocol: TCP
    - containerPort: 8080
      hostPort: 8080
      name: local
      protocol: TCP
    resources:
      requests:
        cpu: 250m
    volumeMounts:
    - mountPath: /srv/kubernetes
      name: srvkube
      readOnly: true
    - mountPath: /var/log/kube-apiserver.log
      name: logfile
    - mountPath: /etc/ssl
      name: etcssl
      readOnly: true
    - mountPath: /usr/share/ca-certificates
      name: usrsharecacerts
      readOnly: true
    - mountPath: /srv/sshproxy
      name: srvsshproxy
  dnsPolicy: ClusterFirst
  hostNetwork: true
  restartPolicy: Always
  terminationGracePeriodSeconds: 30
  volumes:
  - hostPath:
      path: /etc/kubernetes/cacerts
    name: srvkube
  - hostPath:
      path: /var/log/containers/kube-apiserver.log
    name: logfile
  - hostPath:
      path: /etc/ssl
    name: etcssl
  - hostPath:
      path: /usr/share/ca-certificates
    name: usrsharecacerts
  - hostPath:
      path: /srv/sshproxy
    name: srvsshproxy


[root@localhost manifests]# cat kube-controller-manager.yaml 
apiVersion: v1
kind: Pod
metadata:
  name: kube-controller-manager
  namespace: kube-system
spec:
  containers:
  - command:
    - /bin/sh
    - -c
    - /usr/local/bin/kube-controller-manager
      --master=127.0.0.1:8080
      --cluster-name=kubernetes
      --cluster-cidr=10.120.0.0/14
      --service-cluster-ip-range=10.123.240.0/20
      --allocate-node-cidrs=true
      --service-account-private-key-file=/srv/kubernetes/server.key
      --v=2
      --root-ca-file=/srv/kubernetes/ca.crt
      1>>/var/log/kube-controller-manager.log 2>&1
    image: gcr.io/google_containers/kube-controller-manager:6322fd6f4db015a29561a4881ff9c372
    imagePullPolicy: Never # Always IfNotPresent
    livenessProbe:
      httpGet:
        host: 127.0.0.1
        path: /healthz
        port: 10252
        scheme: HTTP
      initialDelaySeconds: 15
      timeoutSeconds: 15
    name: kube-controller-manager
    resources:
      limits:
        cpu: 200m
      requests:
        cpu: 200m
    volumeMounts:
    - mountPath: /srv/kubernetes
      name: srvkube
      readOnly: true
    - mountPath: /var/log/kube-controller-manager.log
      name: logfile
    - mountPath: /etc/ssl
      name: etcssl
      readOnly: true
    - mountPath: /usr/share/ca-certificates
      name: usrsharecacerts
      readOnly: true
  dnsPolicy: ClusterFirst
  hostNetwork: true
  restartPolicy: Always
  terminationGracePeriodSeconds: 30
  volumes:
  - hostPath:
      path: /etc/kubernetes/cacerts
    name: srvkube
  - hostPath:
      path: /var/log/containers/kube-controller-manager.log
    name: logfile
  - hostPath:
      path: /etc/ssl
    name: etcssl
  - hostPath:
      path: /usr/share/ca-certificates
    name: usrsharecacerts


[root@localhost manifests]# cat kube-scheduler.yaml 
apiVersion: v1
kind: Pod
metadata:
  name: kube-scheduler
  namespace: kube-system
spec:
  containers:
  - command:
    - /bin/sh
    - -c
    - /usr/local/bin/kube-scheduler
      --master=127.0.0.1:8080
      --v=2
      1>>/var/log/kube-scheduler.log 2>&1
    image: gcr.io/google_containers/kube-scheduler:ec4cd4e378c840237c447cfa147f98a9
    imagePullPolicy: Never # Always IfNotPresent
    livenessProbe:
      httpGet:
        host: 127.0.0.1
        path: /healthz
        port: 10251
        scheme: HTTP
      initialDelaySeconds: 15
      timeoutSeconds: 15
    name: kube-scheduler
    resources:
      requests:
        cpu: 100m
    volumeMounts:
    - mountPath: /var/log/kube-scheduler.log
      name: logfile
  dnsPolicy: ClusterFirst
  hostNetwork: true
  nodeName: ecp2-lab-master
  restartPolicy: Always
  terminationGracePeriodSeconds: 30
  volumes:
  - hostPath:
      path: /var/log/containers/kube-scheduler.log
    name: logfile
