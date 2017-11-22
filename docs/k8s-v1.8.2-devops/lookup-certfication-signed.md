For example
```
[vagrant@openshiftdev ~]$ openssl x509 -in /etc/kubernetes/pki/apiserver.crt -text -noout
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 6538979717300771866 (0x5abf1f686b325c1a)
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: CN=kubernetes
        Validity
            Not Before: Nov 21 21:30:07 2017 GMT
            Not After : Nov 21 21:30:08 2018 GMT
        Subject: CN=kube-apiserver
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:d6:c0:6d:b6:f6:9a:ed:da:fb:89:c5:90:98:fa:
                    d4:62:39:fc:a4:dd:55:fa:27:11:0e:f0:ca:27:ec:
                    90:56:11:fc:b9:a7:47:8e:33:6c:40:c8:04:93:01:
                    29:c4:a3:54:f4:86:7b:23:ac:ab:eb:c8:57:06:7c:
                    20:86:84:65:2a:c3:94:7d:30:f2:44:ef:11:d9:b5:
                    a4:2a:c4:8c:b8:95:c6:48:67:1c:e5:11:32:8c:20:
                    04:db:b6:b2:07:33:5e:8d:09:6b:5d:74:c7:43:bc:
                    b4:20:3e:f1:d9:c3:ed:ad:93:34:2e:99:b2:b2:27:
                    7e:81:19:22:37:c2:53:ad:7e:4e:e5:75:bf:58:19:
                    bf:ae:8f:76:8c:07:5d:1b:78:42:b8:bd:f9:9f:4f:
                    a3:50:40:38:36:35:1d:d0:b2:59:e0:df:0a:b7:fb:
                    06:b6:2a:f4:b1:08:d8:84:9f:1c:91:ba:18:29:24:
                    de:71:3b:f3:f5:4d:54:80:9a:97:47:57:60:5c:a8:
                    4c:58:50:ef:2f:62:b1:e3:46:43:83:1b:b7:34:c2:
                    5b:eb:6e:88:f3:65:18:53:16:16:43:23:9f:cb:14:
                    0e:7b:46:b8:aa:b7:00:b3:6d:0a:e3:ab:23:44:b3:
                    04:56:60:1b:ea:1a:67:4a:ac:8c:00:87:2d:39:f4:
                    3f:65
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment
            X509v3 Extended Key Usage: 
                TLS Web Server Authentication
            X509v3 Subject Alternative Name: 
                DNS:openshiftdev.local, DNS:kubernetes, DNS:kubernetes.default, DNS:kubernetes.default.svc, DNS:kubernetes.default.svc.cluster.local, IP Address:10.96.0.1, IP Address:10.64.33.82
    Signature Algorithm: sha256WithRSAEncryption
         c3:d2:cd:48:5c:99:f0:e8:06:47:9a:86:4c:90:60:fe:6d:cc:
         41:a5:98:43:b8:f4:20:8d:a1:e7:71:0f:19:d2:2b:3b:fd:fa:
         55:1b:a7:fe:72:df:97:20:ac:ba:1d:a5:54:24:f2:f7:57:16:
         11:fe:b4:87:48:48:6a:8e:07:cb:d0:88:ad:4a:7e:b8:a9:92:
         e6:c8:1d:14:78:25:e7:fc:ae:63:c2:a2:d6:ef:1b:99:58:eb:
         5c:6e:93:4c:46:a0:a2:fc:3d:45:fb:68:40:70:46:05:6c:fc:
         04:d5:06:a0:50:1b:a4:be:ef:30:7c:16:af:d5:49:97:e6:57:
         10:fa:32:06:3f:df:40:e5:13:94:31:52:4f:a8:09:a5:38:f7:
         65:3a:69:24:bc:bd:68:a4:e2:e1:5e:a4:04:9c:07:25:d3:9f:
         ce:1c:cd:10:4e:ac:77:12:45:88:6f:0c:27:da:4f:84:0e:a5:
         83:dc:0f:37:63:e9:30:4e:1f:4c:9e:96:5f:93:83:06:bd:10:
         a5:5f:46:dd:fa:f8:13:3f:13:da:7d:a5:ac:e7:e7:19:24:d0:
         2f:b4:da:c2:39:0a:d0:66:35:00:95:cd:58:ee:75:c5:ef:9b:
         cb:ac:43:0d:de:d1:dd:97:d2:29:b6:7e:62:d4:11:be:2c:dd:
         5b:63:d7:4f
```
