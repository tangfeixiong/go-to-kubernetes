# Instruction

## Table of Contents

[Mirroring Kubernetes repository for CentOS/Fedora](#rpm)

[Mirroring Kubernetes repository for Ubuntu/Debian](#deb)

## RPM

gpg
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ ./yum-gpg-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   663  100   663    0     0     59      0  0:00:11  0:00:11 --:--:--   197
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   975  100   975    0     0    312      0  0:00:03  0:00:03 --:--:--   312
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   991  100   991    0     0    108      0  0:00:09  0:00:09 --:--:--   249
```
repos
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ ./yum-repos-el7-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 36002  100 36002    0     0   7422      0  0:00:04  0:00:04 --:--:--  7638
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  5209  100  5209    0     0   1330      0  0:00:03  0:00:03 --:--:--  1330
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  127k  100  127k    0     0  13103      0  0:00:09  0:00:09 --:--:-- 18815
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  9834  100  9834    0     0    952      0  0:00:10  0:00:10 --:--:--  2291
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  129k  100  129k    0     0  20026      0  0:00:06  0:00:06 --:--:-- 30061
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 12526  100 12526    0     0   1962      0  0:00:06  0:00:06 --:--:--  3262
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1416  100  1416    0     0    202      0  0:00:07  0:00:06  0:00:01   409
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   454  100   454    0     0    470      0 --:--:-- --:--:-- --:--:--   469
```

packages (note: just less and equal than v1.8.4)
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ egrep '^  <location href=".*\.rpm"></location>' https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/primary.xml
  <location href="../../pool/5116fa4b73c700823cfc76b3cedff6622f2fbd0a3d2fa09bce6d93329771e291-kubeadm-1.6.0-0.x86_64.rpm"></location>
  <location href="../../pool/23961d0f7dca1ed118b948195f2fb5dd7a07503d69d7d8ab4433219ea98d033e-kubeadm-1.6.1-0.x86_64.rpm"></location>
  <location href="../../pool/415bc9f86ecfc3b195f22d25fb599e19080d301358c145539d154be95f1985f5-kubeadm-1.6.2-0.x86_64.rpm"></location>
  <location href="../../pool/81158f40764a08356242a53ef4bf7e3c219f48f364c55260db571cae51ce6eec-kubeadm-1.6.3-0.x86_64.rpm"></location>
  <location href="../../pool/0c37012e387fae5fd0d66b5d4664f495ab19fbb709df655533b2b2ecb1a05cd8-kubeadm-1.6.4-0.x86_64.rpm"></location>
  <location href="../../pool/5e62d624af011598f2df59c7759bf6ff4a6bba5675f0c07a2fb18e3c468ea41c-kubeadm-1.6.5-0.x86_64.rpm"></location>
  <location href="../../pool/03ea47d2b0f1912c3a721c2b7353ead2e28a154b6c883200492db0558f6b09f2-kubeadm-1.6.6-0.x86_64.rpm"></location>
  <location href="../../pool/6b36b8d9783ee59480bda0caabcb9ce95a6bf69c62ea38e7bb0cee65174fe479-kubeadm-1.6.7-0.x86_64.rpm"></location>
  <location href="../../pool/c10b6d67b13101afd0270873972f10b14e9719ad8564797cc6609f7c58ae033c-kubeadm-1.6.11-0.x86_64.rpm"></location>
  <location href="../../pool/fda8179e8174e43c17662d3a019f54e862d270fbfc07c3f7a88a58a033198153-kubeadm-1.6.12-0.x86_64.rpm"></location>
  <location href="../../pool/cb7034dff51af01c237e03ea342dc74269b46cce1e43bb0ab9ef1d6f006ebda9-kubeadm-1.7.0-0.x86_64.rpm"></location>
  <location href="../../pool/77bea7f33178ab4feb6afaf31c6c779511498ef47dec433eb5464909f7a26dc3-kubeadm-1.7.1-0.x86_64.rpm"></location>
  <location href="../../pool/1a6f5f73f43077a50d877df505481e5a3d765c979b89fda16b8b9622b9ebd9a4-kubeadm-1.7.2-0.x86_64.rpm"></location>
  <location href="../../pool/f7ec56b0f36a81c0f91bcf26e05f23088082b468b77dac576dc505444dd8cd48-kubeadm-1.7.3-1.x86_64.rpm"></location>
  <location href="../../pool/f0a51fcde5e3b329050d7a6cf70f04a6cdf09eacfbad55f4324bfa2ea4312d0e-kubeadm-1.7.4-0.x86_64.rpm"></location>
  <location href="../../pool/02f3a7ff6e04943bd288ff302f449b600e8db3d19868dfe4308d0d902c0ba927-kubeadm-1.7.5-0.x86_64.rpm"></location>
  <location href="../../pool/4ff875dc8857b85c490b42b750527ba20a154a49a8dacd256d16cbbf5e708dfd-kubeadm-1.7.6-1.x86_64.rpm"></location>
  <location href="../../pool/2bb9ddc5197dec31ac73a549067ab9b6a529cc31275f9223ff04efbfb5451602-kubeadm-1.7.7-1.x86_64.rpm"></location>
  <location href="../../pool/8af3d230e7c1c775124b518aa0d69fa911e761e7b22fd081bfae8dece2381970-kubeadm-1.7.8-1.x86_64.rpm"></location>
  <location href="../../pool/c0638ef93a73ad4343fe3c2b8105e3826605e2fe000017324b3328975aa36a82-kubeadm-1.7.10-0.x86_64.rpm"></location>
  <location href="../../pool/ee93b5249497dba6488262bd2b56e32438a69f78993eb973bafa52d72f9b883a-kubeadm-1.8.0-0.x86_64.rpm"></location>
  <location href="../../pool/7fe94dd0f828ef64d4ccdfef160417457309b904758ddf10c040612d7e5ef441-kubeadm-1.8.0-1.x86_64.rpm"></location>
  <location href="../../pool/0f7d8ea10144399f3d60446fab5469395afb809c175bdc0eae4d12c1fcc3cb62-kubeadm-1.8.1-0.x86_64.rpm"></location>
  <location href="../../pool/d64bc1d0ca27196030c6f574252a0872b998b29be90c6cb7e97b91cd0bc78fed-kubeadm-1.8.2-0.x86_64.rpm"></location>
  <location href="../../pool/cab6b288e91e613d81c63101c7d293059a4a9f2c0795228042c880f770a9ec60-kubeadm-1.8.3-0.x86_64.rpm"></location>
  <location href="../../pool/aeaad1e283c54876b759a089f152228d7cd4c049f271125c23623995b8e76f96-kubeadm-1.8.4-0.x86_64.rpm"></location>
  <location href="../../pool/fd9b1e2215cab6da7adc6e87e2710b91ecfda3b617edfe7e71c92277301a63ab-kubectl-1.5.4-0.x86_64.rpm"></location>
  <location href="../../pool/e6aef7b2b7d9e5bd4db1e5747ebbc9f1f97bbfb8c7817ad68028565ca263a672-kubectl-1.6.0-0.x86_64.rpm"></location>
  <location href="../../pool/9d1ccf0877dfc4228a923a9614661b27d663694680e2bc0a1c184aa937fbf7f2-kubectl-1.6.1-0.x86_64.rpm"></location>
  <location href="../../pool/ff72cbf0dfa986c36097a5c3ac2301ecb73ed28df8db86e13641a79e9df9b3ea-kubectl-1.6.2-0.x86_64.rpm"></location>
  <location href="../../pool/d5cc6bb2e439298eb1b3e45c3ac58010580c1d3c3a2fa71040a7c73b1dc22881-kubectl-1.6.3-0.x86_64.rpm"></location>
  <location href="../../pool/9fcf36d1acf5e17c4aec783e956b8b70ab0e91693d394e6268c16533fe8a71a9-kubectl-1.6.4-0.x86_64.rpm"></location>
  <location href="../../pool/504e6824515b1b91702d1cd9aa8701fbcb2fabeb61e8395204f4e2cd07f7dfb7-kubectl-1.6.5-0.x86_64.rpm"></location>
  <location href="../../pool/b9584919848f1cc3096ec8cd0bbeceb95ea0786df339eb4883556a4a89f51010-kubectl-1.6.6-0.x86_64.rpm"></location>
  <location href="../../pool/b542021d5f32457f8c1f6f726aaa077eec66b0906440a020cfada28275df50f7-kubectl-1.6.7-0.x86_64.rpm"></location>
  <location href="../../pool/9b3b2d1cbe36086a7008d261594b31bb3e085e8c170aa319e1befebe604a99a7-kubectl-1.6.11-0.x86_64.rpm"></location>
  <location href="../../pool/712b311a1853d2b73be6729f57fc448ebc8818f0a9236cbbf1f11c6db3c3d8dc-kubectl-1.6.12-0.x86_64.rpm"></location>
  <location href="../../pool/888aff6bf68f988e7480efd06f38852eef8a44eed1baaa62f3ec4fbb07c35f7d-kubectl-1.7.0-0.x86_64.rpm"></location>
  <location href="../../pool/b8a64634ad1555e15a61b84685fd04959435ed6374b25e369f5bda89c8f03a6b-kubectl-1.7.1-0.x86_64.rpm"></location>
  <location href="../../pool/dc8329515fc3245404fea51839241b58774e577d7736f99f21276e764c309db5-kubectl-1.7.2-0.x86_64.rpm"></location>
  <location href="../../pool/c8a50a1ce9c948e7a238b02d3967d220e71e13e1ac8916a961176726eabe8aa1-kubectl-1.7.3-1.x86_64.rpm"></location>
  <location href="../../pool/041d5a6813dab590b160865fea7259bc2db762a9667379d03aca8d4596a3cccd-kubectl-1.7.4-0.x86_64.rpm"></location>
  <location href="../../pool/c2d4b7c1f39ca9532a2965ea513fcd900fdcdd8785639c8fbf69f30780cae764-kubectl-1.7.5-0.x86_64.rpm"></location>
  <location href="../../pool/71aa78fc7472de3664511c88f9d58d9a9c6922f26d67323869b5a661b106e0d0-kubectl-1.7.6-1.x86_64.rpm"></location>
  <location href="../../pool/4af2eb4114017f12fb36b92a85c8149de6c54020a73061f3666d033bfde8f3e3-kubectl-1.7.7-1.x86_64.rpm"></location>
  <location href="../../pool/761837737577fe362b16ff2b1fc34c383d86b0f9f746a9901c62587fd5f4e0f6-kubectl-1.7.8-1.x86_64.rpm"></location>
  <location href="../../pool/ef4f06f44a3e9412735722d5625f814505054bd5f672ccad3e631470e5da9cd0-kubectl-1.7.10-0.x86_64.rpm"></location>
  <location href="../../pool/582709dfb1c59f85c78f7db2c7609a1f3bf7308b73255bdbc910695b31b8660f-kubectl-1.8.0-0.x86_64.rpm"></location>
  <location href="../../pool/0ee48e6b4033fdf520f5893759b0665090ffb83eefdbe3f0b41edf54f2247ee4-kubectl-1.8.1-0.x86_64.rpm"></location>
  <location href="../../pool/3cc05eb1b1565779e8033743f1a2b9c8fb4e3b432421719ec56a3024d33dfccc-kubectl-1.8.2-0.x86_64.rpm"></location>
  <location href="../../pool/ae43f92f96e828779f9744b3660e207199d97dda22eb44c054d2f3150da76b94-kubectl-1.8.3-0.x86_64.rpm"></location>
  <location href="../../pool/a9db28728641ddbf7f025b8b496804d82a396d0ccb178fffd124623fb2f999ea-kubectl-1.8.4-0.x86_64.rpm"></location>
  <location href="../../pool/2e63c1f9436c6413a4ea0f45145b97c6dbf55e9bb2a251adc38db1292bbd6ed1-kubelet-1.5.4-0.x86_64.rpm"></location>
  <location href="../../pool/af8567f1ba6f8dc1d43b60702d45c02aca88607b0e721d76897e70f6a6e53115-kubelet-1.6.0-0.x86_64.rpm"></location>
  <location href="../../pool/cde0b9092d421800f7207df677d6a1f321c82179e89a24e4b3c009a875e62c91-kubelet-1.6.1-0.x86_64.rpm"></location>
  <location href="../../pool/a8a2b37431da90798331a67b8b452572a72173414b1543368786e356f23bc4ce-kubelet-1.6.2-0.x86_64.rpm"></location>
  <location href="../../pool/b610c239bd57c3ca0dcae3f4d4ae60a0a1eab99c7b059cf1b0fe8dd7267c3f73-kubelet-1.6.3-0.x86_64.rpm"></location>
  <location href="../../pool/08706c8f6fbce961cd8994043cd27204f2bd3b47afd54cbd8755dae4b6cc4279-kubelet-1.6.4-0.x86_64.rpm"></location>
  <location href="../../pool/7d43767d519f9f76605408517c999631e3835621afa8f5e80b9b8fd0925842ca-kubelet-1.6.5-0.x86_64.rpm"></location>
  <location href="../../pool/f2c5f556143a820ed5bbb6a0ac7601f68dea8b28c8edc5db0a0d5d0ad4e94fd0-kubelet-1.6.6-0.x86_64.rpm"></location>
  <location href="../../pool/b58a3f13d494458fbe23dbf22abc0213dc2c9ffb1e30f76dad3d7321d0715444-kubelet-1.6.7-0.x86_64.rpm"></location>
  <location href="../../pool/254343aac568684ff6bdf75e69c015127b45cd8a6e7f9b0eff335b0a082b3117-kubelet-1.6.11-0.x86_64.rpm"></location>
  <location href="../../pool/f469e89265a9a215358f6d2e541c0d5dd4f6b522fcbbe401117ce656b3aba931-kubelet-1.6.12-0.x86_64.rpm"></location>
  <location href="../../pool/8f94c56214c25f72f999f8a9c099275c2e1b99110a155679c1eda293192c26a7-kubelet-1.7.0-0.x86_64.rpm"></location>
  <location href="../../pool/30136924ea242fee92df3aaf0297365dc58e8d1260b7bde66c94bda5a42a6f04-kubelet-1.7.1-0.x86_64.rpm"></location>
  <location href="../../pool/1e508e26f2b02971a7ff5f034b48a6077d613e0b222e0ec973351117b4ff45ea-kubelet-1.7.2-0.x86_64.rpm"></location>
  <location href="../../pool/28b76e6e1c2ec397a9b6111045316a0943da73dd5602ee8e53752cdca62409e6-kubelet-1.7.3-1.x86_64.rpm"></location>
  <location href="../../pool/4f60c17a926175fb9abcfdd487cebafbbbce0e2248d2b99c189ae0877376b88d-kubelet-1.7.4-0.x86_64.rpm"></location>
  <location href="../../pool/c87d884d28952332d952d5f7f294ead1838a8e1219d2fc75923a5be06903afaf-kubelet-1.7.5-0.x86_64.rpm"></location>
  <location href="../../pool/f049d9a0a9172b00aa2725c86a0de6f4ee51571105344b31b0b2523be9fda635-kubelet-1.7.6-1.x86_64.rpm"></location>
  <location href="../../pool/45bd874fbf333c3e2595da358ce7ad3220af46b1ff69d69f57f882dee0db52a8-kubelet-1.7.7-1.x86_64.rpm"></location>
  <location href="../../pool/305fd1a89e9da295f56ac2df41dd73e521e29c768c57074a365096d713bfe928-kubelet-1.7.8-1.x86_64.rpm"></location>
  <location href="../../pool/b7127de06c5bec3e197d279ac910b7a119fa2f799fb18ca3ec11499506396957-kubelet-1.7.10-0.x86_64.rpm"></location>
  <location href="../../pool/d51b547da5980e2f6bfd33a382779194a0eacecac999ca726f85ef76183a9033-kubelet-1.8.0-0.x86_64.rpm"></location>
  <location href="../../pool/a35571037b554243d386436ff729c90a3cb270f5797b7cd254ef0afbd4e706bf-kubelet-1.8.1-0.x86_64.rpm"></location>
  <location href="../../pool/7d976686554e598267d8c5eb030f2a8f4e575f47015ecf94459913b80c9e5bb8-kubelet-1.8.2-0.x86_64.rpm"></location>
  <location href="../../pool/a53acfe63a475bf61661036c12890217f4921a6d6d6c3e6ecb4c598fc11cac19-kubelet-1.8.3-0.x86_64.rpm"></location>
  <location href="../../pool/1acca81eb5cf99453f30466876ff03146112b7f12c625cb48f12508684e02665-kubelet-1.8.4-0.x86_64.rpm"></location>
  <location href="../../pool/0c923b191423caccc65c550ef07ce61b97f991ad54785dab70dc07a5763f4222-kubernetes-cni-0.3.0.1-0.07a8a2.x86_64.rpm"></location>
  <location href="../../pool/e7a4403227dd24036f3b0615663a371c4e07a95be5fee53505e647fd8ae58aa6-kubernetes-cni-0.5.1-0.x86_64.rpm"></location>
  <location href="../../pool/79f9ba89dbe7000e7dfeda9b119f711bb626fe2c2d56abeb35141142cda00342-kubernetes-cni-0.5.1-1.x86_64.rpm"></location>
  <location href="../../pool/fd5f5da2d1a262fa687404d34e72813520274364557e648bc64a8136f1a02630-rkt-1.25.0-1.x86_64.rpm"></location>
  <location href="../../pool/7a382e59dc2c39a66083e03ec061f33771e4a7130c98cd0ef61492b2196c0378-rkt-1.26.0-1.x86_64.rpm"></location>
  <location href="../../pool/01b97b0ddb967d0ed9fd78327a784efbfea8cd0d9789f5bab8b9bbfe94477c60-rkt-1.27.0-1.x86_64.rpm"></location>
```

or
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ egrep '^  <location href=".*\.rpm"></location>' https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/primary.xml | cut -d '"' -f 2
../../pool/5116fa4b73c700823cfc76b3cedff6622f2fbd0a3d2fa09bce6d93329771e291-kubeadm-1.6.0-0.x86_64.rpm
../../pool/23961d0f7dca1ed118b948195f2fb5dd7a07503d69d7d8ab4433219ea98d033e-kubeadm-1.6.1-0.x86_64.rpm
../../pool/415bc9f86ecfc3b195f22d25fb599e19080d301358c145539d154be95f1985f5-kubeadm-1.6.2-0.x86_64.rpm
../../pool/81158f40764a08356242a53ef4bf7e3c219f48f364c55260db571cae51ce6eec-kubeadm-1.6.3-0.x86_64.rpm
../../pool/0c37012e387fae5fd0d66b5d4664f495ab19fbb709df655533b2b2ecb1a05cd8-kubeadm-1.6.4-0.x86_64.rpm
../../pool/5e62d624af011598f2df59c7759bf6ff4a6bba5675f0c07a2fb18e3c468ea41c-kubeadm-1.6.5-0.x86_64.rpm
../../pool/03ea47d2b0f1912c3a721c2b7353ead2e28a154b6c883200492db0558f6b09f2-kubeadm-1.6.6-0.x86_64.rpm
../../pool/6b36b8d9783ee59480bda0caabcb9ce95a6bf69c62ea38e7bb0cee65174fe479-kubeadm-1.6.7-0.x86_64.rpm
../../pool/c10b6d67b13101afd0270873972f10b14e9719ad8564797cc6609f7c58ae033c-kubeadm-1.6.11-0.x86_64.rpm
../../pool/fda8179e8174e43c17662d3a019f54e862d270fbfc07c3f7a88a58a033198153-kubeadm-1.6.12-0.x86_64.rpm
../../pool/cb7034dff51af01c237e03ea342dc74269b46cce1e43bb0ab9ef1d6f006ebda9-kubeadm-1.7.0-0.x86_64.rpm
../../pool/77bea7f33178ab4feb6afaf31c6c779511498ef47dec433eb5464909f7a26dc3-kubeadm-1.7.1-0.x86_64.rpm
../../pool/1a6f5f73f43077a50d877df505481e5a3d765c979b89fda16b8b9622b9ebd9a4-kubeadm-1.7.2-0.x86_64.rpm
../../pool/f7ec56b0f36a81c0f91bcf26e05f23088082b468b77dac576dc505444dd8cd48-kubeadm-1.7.3-1.x86_64.rpm
../../pool/f0a51fcde5e3b329050d7a6cf70f04a6cdf09eacfbad55f4324bfa2ea4312d0e-kubeadm-1.7.4-0.x86_64.rpm
../../pool/02f3a7ff6e04943bd288ff302f449b600e8db3d19868dfe4308d0d902c0ba927-kubeadm-1.7.5-0.x86_64.rpm
../../pool/4ff875dc8857b85c490b42b750527ba20a154a49a8dacd256d16cbbf5e708dfd-kubeadm-1.7.6-1.x86_64.rpm
../../pool/2bb9ddc5197dec31ac73a549067ab9b6a529cc31275f9223ff04efbfb5451602-kubeadm-1.7.7-1.x86_64.rpm
../../pool/8af3d230e7c1c775124b518aa0d69fa911e761e7b22fd081bfae8dece2381970-kubeadm-1.7.8-1.x86_64.rpm
../../pool/c0638ef93a73ad4343fe3c2b8105e3826605e2fe000017324b3328975aa36a82-kubeadm-1.7.10-0.x86_64.rpm
../../pool/ee93b5249497dba6488262bd2b56e32438a69f78993eb973bafa52d72f9b883a-kubeadm-1.8.0-0.x86_64.rpm
../../pool/7fe94dd0f828ef64d4ccdfef160417457309b904758ddf10c040612d7e5ef441-kubeadm-1.8.0-1.x86_64.rpm
../../pool/0f7d8ea10144399f3d60446fab5469395afb809c175bdc0eae4d12c1fcc3cb62-kubeadm-1.8.1-0.x86_64.rpm
../../pool/d64bc1d0ca27196030c6f574252a0872b998b29be90c6cb7e97b91cd0bc78fed-kubeadm-1.8.2-0.x86_64.rpm
../../pool/cab6b288e91e613d81c63101c7d293059a4a9f2c0795228042c880f770a9ec60-kubeadm-1.8.3-0.x86_64.rpm
../../pool/aeaad1e283c54876b759a089f152228d7cd4c049f271125c23623995b8e76f96-kubeadm-1.8.4-0.x86_64.rpm
../../pool/fd9b1e2215cab6da7adc6e87e2710b91ecfda3b617edfe7e71c92277301a63ab-kubectl-1.5.4-0.x86_64.rpm
../../pool/e6aef7b2b7d9e5bd4db1e5747ebbc9f1f97bbfb8c7817ad68028565ca263a672-kubectl-1.6.0-0.x86_64.rpm
../../pool/9d1ccf0877dfc4228a923a9614661b27d663694680e2bc0a1c184aa937fbf7f2-kubectl-1.6.1-0.x86_64.rpm
../../pool/ff72cbf0dfa986c36097a5c3ac2301ecb73ed28df8db86e13641a79e9df9b3ea-kubectl-1.6.2-0.x86_64.rpm
../../pool/d5cc6bb2e439298eb1b3e45c3ac58010580c1d3c3a2fa71040a7c73b1dc22881-kubectl-1.6.3-0.x86_64.rpm
../../pool/9fcf36d1acf5e17c4aec783e956b8b70ab0e91693d394e6268c16533fe8a71a9-kubectl-1.6.4-0.x86_64.rpm
../../pool/504e6824515b1b91702d1cd9aa8701fbcb2fabeb61e8395204f4e2cd07f7dfb7-kubectl-1.6.5-0.x86_64.rpm
../../pool/b9584919848f1cc3096ec8cd0bbeceb95ea0786df339eb4883556a4a89f51010-kubectl-1.6.6-0.x86_64.rpm
../../pool/b542021d5f32457f8c1f6f726aaa077eec66b0906440a020cfada28275df50f7-kubectl-1.6.7-0.x86_64.rpm
../../pool/9b3b2d1cbe36086a7008d261594b31bb3e085e8c170aa319e1befebe604a99a7-kubectl-1.6.11-0.x86_64.rpm
../../pool/712b311a1853d2b73be6729f57fc448ebc8818f0a9236cbbf1f11c6db3c3d8dc-kubectl-1.6.12-0.x86_64.rpm
../../pool/888aff6bf68f988e7480efd06f38852eef8a44eed1baaa62f3ec4fbb07c35f7d-kubectl-1.7.0-0.x86_64.rpm
../../pool/b8a64634ad1555e15a61b84685fd04959435ed6374b25e369f5bda89c8f03a6b-kubectl-1.7.1-0.x86_64.rpm
../../pool/dc8329515fc3245404fea51839241b58774e577d7736f99f21276e764c309db5-kubectl-1.7.2-0.x86_64.rpm
../../pool/c8a50a1ce9c948e7a238b02d3967d220e71e13e1ac8916a961176726eabe8aa1-kubectl-1.7.3-1.x86_64.rpm
../../pool/041d5a6813dab590b160865fea7259bc2db762a9667379d03aca8d4596a3cccd-kubectl-1.7.4-0.x86_64.rpm
../../pool/c2d4b7c1f39ca9532a2965ea513fcd900fdcdd8785639c8fbf69f30780cae764-kubectl-1.7.5-0.x86_64.rpm
../../pool/71aa78fc7472de3664511c88f9d58d9a9c6922f26d67323869b5a661b106e0d0-kubectl-1.7.6-1.x86_64.rpm
../../pool/4af2eb4114017f12fb36b92a85c8149de6c54020a73061f3666d033bfde8f3e3-kubectl-1.7.7-1.x86_64.rpm
../../pool/761837737577fe362b16ff2b1fc34c383d86b0f9f746a9901c62587fd5f4e0f6-kubectl-1.7.8-1.x86_64.rpm
../../pool/ef4f06f44a3e9412735722d5625f814505054bd5f672ccad3e631470e5da9cd0-kubectl-1.7.10-0.x86_64.rpm
../../pool/582709dfb1c59f85c78f7db2c7609a1f3bf7308b73255bdbc910695b31b8660f-kubectl-1.8.0-0.x86_64.rpm
../../pool/0ee48e6b4033fdf520f5893759b0665090ffb83eefdbe3f0b41edf54f2247ee4-kubectl-1.8.1-0.x86_64.rpm
../../pool/3cc05eb1b1565779e8033743f1a2b9c8fb4e3b432421719ec56a3024d33dfccc-kubectl-1.8.2-0.x86_64.rpm
../../pool/ae43f92f96e828779f9744b3660e207199d97dda22eb44c054d2f3150da76b94-kubectl-1.8.3-0.x86_64.rpm
../../pool/a9db28728641ddbf7f025b8b496804d82a396d0ccb178fffd124623fb2f999ea-kubectl-1.8.4-0.x86_64.rpm
../../pool/2e63c1f9436c6413a4ea0f45145b97c6dbf55e9bb2a251adc38db1292bbd6ed1-kubelet-1.5.4-0.x86_64.rpm
../../pool/af8567f1ba6f8dc1d43b60702d45c02aca88607b0e721d76897e70f6a6e53115-kubelet-1.6.0-0.x86_64.rpm
../../pool/cde0b9092d421800f7207df677d6a1f321c82179e89a24e4b3c009a875e62c91-kubelet-1.6.1-0.x86_64.rpm
../../pool/a8a2b37431da90798331a67b8b452572a72173414b1543368786e356f23bc4ce-kubelet-1.6.2-0.x86_64.rpm
../../pool/b610c239bd57c3ca0dcae3f4d4ae60a0a1eab99c7b059cf1b0fe8dd7267c3f73-kubelet-1.6.3-0.x86_64.rpm
../../pool/08706c8f6fbce961cd8994043cd27204f2bd3b47afd54cbd8755dae4b6cc4279-kubelet-1.6.4-0.x86_64.rpm
../../pool/7d43767d519f9f76605408517c999631e3835621afa8f5e80b9b8fd0925842ca-kubelet-1.6.5-0.x86_64.rpm
../../pool/f2c5f556143a820ed5bbb6a0ac7601f68dea8b28c8edc5db0a0d5d0ad4e94fd0-kubelet-1.6.6-0.x86_64.rpm
../../pool/b58a3f13d494458fbe23dbf22abc0213dc2c9ffb1e30f76dad3d7321d0715444-kubelet-1.6.7-0.x86_64.rpm
../../pool/254343aac568684ff6bdf75e69c015127b45cd8a6e7f9b0eff335b0a082b3117-kubelet-1.6.11-0.x86_64.rpm
../../pool/f469e89265a9a215358f6d2e541c0d5dd4f6b522fcbbe401117ce656b3aba931-kubelet-1.6.12-0.x86_64.rpm
../../pool/8f94c56214c25f72f999f8a9c099275c2e1b99110a155679c1eda293192c26a7-kubelet-1.7.0-0.x86_64.rpm
../../pool/30136924ea242fee92df3aaf0297365dc58e8d1260b7bde66c94bda5a42a6f04-kubelet-1.7.1-0.x86_64.rpm
../../pool/1e508e26f2b02971a7ff5f034b48a6077d613e0b222e0ec973351117b4ff45ea-kubelet-1.7.2-0.x86_64.rpm
../../pool/28b76e6e1c2ec397a9b6111045316a0943da73dd5602ee8e53752cdca62409e6-kubelet-1.7.3-1.x86_64.rpm
../../pool/4f60c17a926175fb9abcfdd487cebafbbbce0e2248d2b99c189ae0877376b88d-kubelet-1.7.4-0.x86_64.rpm
../../pool/c87d884d28952332d952d5f7f294ead1838a8e1219d2fc75923a5be06903afaf-kubelet-1.7.5-0.x86_64.rpm
../../pool/f049d9a0a9172b00aa2725c86a0de6f4ee51571105344b31b0b2523be9fda635-kubelet-1.7.6-1.x86_64.rpm
../../pool/45bd874fbf333c3e2595da358ce7ad3220af46b1ff69d69f57f882dee0db52a8-kubelet-1.7.7-1.x86_64.rpm
../../pool/305fd1a89e9da295f56ac2df41dd73e521e29c768c57074a365096d713bfe928-kubelet-1.7.8-1.x86_64.rpm
../../pool/b7127de06c5bec3e197d279ac910b7a119fa2f799fb18ca3ec11499506396957-kubelet-1.7.10-0.x86_64.rpm
../../pool/d51b547da5980e2f6bfd33a382779194a0eacecac999ca726f85ef76183a9033-kubelet-1.8.0-0.x86_64.rpm
../../pool/a35571037b554243d386436ff729c90a3cb270f5797b7cd254ef0afbd4e706bf-kubelet-1.8.1-0.x86_64.rpm
../../pool/7d976686554e598267d8c5eb030f2a8f4e575f47015ecf94459913b80c9e5bb8-kubelet-1.8.2-0.x86_64.rpm
../../pool/a53acfe63a475bf61661036c12890217f4921a6d6d6c3e6ecb4c598fc11cac19-kubelet-1.8.3-0.x86_64.rpm
../../pool/1acca81eb5cf99453f30466876ff03146112b7f12c625cb48f12508684e02665-kubelet-1.8.4-0.x86_64.rpm
../../pool/0c923b191423caccc65c550ef07ce61b97f991ad54785dab70dc07a5763f4222-kubernetes-cni-0.3.0.1-0.07a8a2.x86_64.rpm
../../pool/e7a4403227dd24036f3b0615663a371c4e07a95be5fee53505e647fd8ae58aa6-kubernetes-cni-0.5.1-0.x86_64.rpm
../../pool/79f9ba89dbe7000e7dfeda9b119f711bb626fe2c2d56abeb35141142cda00342-kubernetes-cni-0.5.1-1.x86_64.rpm
../../pool/fd5f5da2d1a262fa687404d34e72813520274364557e648bc64a8136f1a02630-rkt-1.25.0-1.x86_64.rpm
../../pool/7a382e59dc2c39a66083e03ec061f33771e4a7130c98cd0ef61492b2196c0378-rkt-1.26.0-1.x86_64.rpm
../../pool/01b97b0ddb967d0ed9fd78327a784efbfea8cd0d9789f5bab8b9bbfe94477c60-rkt-1.27.0-1.x86_64.rpm
```

pool, for example cni 0.3/0.5
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ ./yum-pool-cni-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  9.8M  100  9.8M    0     0  60589      0  0:02:49  0:02:49 --:--:--  118k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 7617k  100 7617k    0     0  79571      0  0:01:38  0:01:38 --:--:-- 95605
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 7619k  100 7619k    0     0  68644      0  0:01:53  0:01:53 --:--:-- 80336
```

pool, for example k8s v1.8.2
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ ./yum-pool-v1_8_2-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 14.8M  100 14.8M    0     0  25283      0  0:10:17  0:10:17 --:--:-- 10791
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 7440k  100 7440k    0     0  27799      0  0:04:34  0:04:34 --:--:-- 35923
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 15.7M  100 15.7M    0     0  20183      0  0:13:39  0:13:39 --:--:-- 11534
```

### Fedora 23

Online `dnf`
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ cat /etc/yum.repos.d/kubernetes.repo 
[kubernetes]
name=Kubernetes
#baseurl=http://yum.kubernetes.io/repos/kubernetes-el7-x86_64
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
    https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
```

```
[vagrant@localhost ~]$ dnf list --verbose | grep kube
导入 GPG 公钥 0xA7317B0F:
 Userid: "Google Cloud Packages Automatic Signing Key <gc-team@google.com>"
 指纹: D0BC 747F D8CA F711 7500 D6FA 3746 C208 A731 7B0F
 来自: https://packages.cloud.google.com/yum/doc/yum-key.gpg
^C^C已终止。
```

```
[vagrant@localhost ~]$ sudo dnf list --verbose | grep kube
导入 GPG 公钥 0xA7317B0F:
 Userid: "Google Cloud Packages Automatic Signing Key <gc-team@google.com>"
 指纹: D0BC 747F D8CA F711 7500 D6FA 3746 C208 A731 7B0F
 来自: https://packages.cloud.google.com/yum/doc/yum-key.gpg
确定吗？[y/N]： y
导入 GPG 公钥 0x3E1BA8D5:
 Userid: "Google Cloud Packages RPM Signing Key <gc-team@google.com>"
 指纹: 3749 E1BA 95A8 6CE0 5454 6ED2 F09C 394C 3E1B A8D5
 来自: https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
确定吗？[y/N]： y
kubernetes：正在使用截止于 Wed Oct  1 11:06:17 47844245 的元数据。
cockpit-kubernetes.x86_64                0.96-1.fc23                     updates
kubeadm.x86_64                           1.8.2-0                         kubernetes
kubectl.x86_64                           1.8.2-0                         kubernetes
kubelet.x86_64                           1.8.2-0                         kubernetes
kubernetes.x86_64                        1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-ansible.noarch                0.6.0-0.2.gitd65ebd5.fc23       updates
kubernetes-client.x86_64                 1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-cni.x86_64                    0.5.1-1                         kubernetes
kubernetes-devel.x86_64                  1.1.0-0.5.gite44c8e6.fc23       fedora 
kubernetes-devel.noarch                  1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-master.x86_64                 1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-node.x86_64                   1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-unit-test.x86_64              1.2.0-0.26.git4a3f9c5.fc23      updates
rkt.x86_64                               1.27.0-1                        kubernetes
```

Offline
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ echo -e "\n[kubernetes-offline]\nname=Kubernetes Offline\nbaseurl=file:///Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64\nenabled=0\ngpgcheck=0\n" | sudo tee -a /etc/yum.repos.d/kubernetes.repo 

[kubernetes-offline]
name=Kubernetes Offline
baseurl=file:///Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=0
gpgcheck=0
```

```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ sudo dnf --disablerepo kubernetes --enablerepo kubernetes-offline list | grep kube
cockpit-kubernetes.x86_64                0.96-1.fc23                     updates
kubeadm.x86_64                           1.8.4-0                         kubernetes-offline
kubectl.x86_64                           1.8.4-0                         kubernetes-offline
kubelet.x86_64                           1.8.4-0                         kubernetes-offline
kubernetes.x86_64                        1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-ansible.noarch                0.6.0-0.2.gitd65ebd5.fc23       updates
kubernetes-client.x86_64                 1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-cni.x86_64                    0.5.1-1                         kubernetes-offline
kubernetes-devel.x86_64                  1.1.0-0.5.gite44c8e6.fc23       fedora 
kubernetes-devel.noarch                  1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-master.x86_64                 1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-node.x86_64                   1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-unit-test.x86_64              1.2.0-0.26.git4a3f9c5.fc23      updates
rkt.x86_64                               1.27.0-1                        kubernetes-offline
```

Local HTTP
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ curl -L http://172.17.4.50:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
<pre>
<a href="repodata/">repodata/</a>
</pre>
```

```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ echo -e "\n[kubernetes-172-17-4-50]\nname=Kubernetes local http\nbaseurl=http://172.17.4.50:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64\nenabled=0\ngpgcheck=0\n" | sudo tee -a /etc/yum.repos.d/kubernetes.repo 

[kubernetes-172-17-4-50]
name=Kubernetes local http
baseurl=http://172.17.4.50:48080/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=0
gpgcheck=0
```

```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ sudo dnf --disablerepo kubernetes --enablerepo kubernetes-172-17-4-50 list | grep kube
cockpit-kubernetes.x86_64                0.96-1.fc23                     updates
kubeadm.x86_64                           1.8.4-0                         kubernetes-172-17-4-50
kubectl.x86_64                           1.8.4-0                         kubernetes-172-17-4-50
kubelet.x86_64                           1.8.4-0                         kubernetes-172-17-4-50
kubernetes.x86_64                        1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-ansible.noarch                0.6.0-0.2.gitd65ebd5.fc23       updates
kubernetes-client.x86_64                 1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-cni.x86_64                    0.5.1-1                         kubernetes-172-17-4-50
kubernetes-devel.x86_64                  1.1.0-0.5.gite44c8e6.fc23       fedora 
kubernetes-devel.noarch                  1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-master.x86_64                 1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-node.x86_64                   1.2.0-0.26.git4a3f9c5.fc23      updates
kubernetes-unit-test.x86_64              1.2.0-0.26.git4a3f9c5.fc23      updates
rkt.x86_64                               1.27.0-1                        kubernetes-172-17-4-50
```

commands
```
[vagrant@localhost ~]$ sudo dnf --help
usage: dnf [options] COMMAND

主命令列表

autoremove                
check-update              检查是否有软件包升级
clean                     删除缓存的数据
distro-sync               已同步软件包到最新可用版本
downgrade                 降级包
group                     显示或使用组信息
help                      显示用法信息
history                   显示或使用事务历史
info                      显示关于软件包或组的详细信息
install                   向系统中安装一个或多个软件包
list                      列出一个或一组软件包
makecache                 创建元数据缓存
mark                      标记或取消标记已安装的软件包中由用户安装的。
provides                  查找提供指定内容的软件包
reinstall                 重装一个包
remove                    从系统中移除一个或多个软件包
repolist                  显示已配置的软件仓库
repository-packages       对指定仓库中的所有软件包运行命令
search                    在软件包详细信息中搜索指定字符串
updateinfo                显示软件包的参考建议
upgrade                   升级系统中的一个或多个软件包
upgrade-to                升级系统中的一个软件包到指定版本

插件命令列表

builddep                  Install build dependencies for package or spec file
config-manager            管理 dnf 配置选项和软件仓库
copr                      与 Copr 仓库交互
debuginfo-install         安装除错信息软件包
download                  下载软件包至当前目录
needs-restarting          判断所升级的二进制文件是否需要重启
playground                与 Playground 仓库交互。
repoquery                 搜索匹配关键字的软件包
reposync                  下载远程仓库中的全部软件包

optional arguments:
  --allowerasing        允许解决依赖关系时删除已安装软件包
  -b, --best            在事务中尝试最佳软件包版本。
  -C, --cacheonly       完全从系统缓存运行，不升级缓存
  -c [config file], --config [config file]
                        配置文件位置
  -d [debug level], --debuglevel [debug level]
                        调试输出级别
  --debugsolver         转储详细解决结果至文件
  --showduplicates      在 list/search 命令下，显示仓库里重复的条目
  -e ERRORLEVEL, --errorlevel ERRORLEVEL
                        错误输出级别
  --rpmverbosity [debug level name]
                        rpm调试输出等级
  -q, --quiet           静默执行
  -v, --verbose         详尽执行
  -y, --assumeyes       全部问题回答为是
  --assumeno            全部问题回答为否
  --version             显示 DNF 版本信息并退出
  --installroot [path]  设置目标根目录
  --enablerepo [repo]
  --disablerepo [repo]
  -x [package], --exclude [package]
                        用全名或通配符排除软件包
  --disableexcludes [repo]
                        禁用排除
  --repofrompath [repo,path]
                        指向附加仓库的标记和路径，可以指定多次。
  --noplugins           禁用所有插件
  --nogpgcheck          禁用 GPG 签名检查
  --disableplugin [plugin]
                        禁用指定名称的插件
  --color COLOR         配置是否使用颜色
  --releasever RELEASEVER
                        覆盖在配置文件和仓库文件中 $releasever 的值
  --setopt SETOPTS      设置任意配置和仓库选项
  --refresh             在运行命令之前将元数据标记为过期。
  -4                    仅解析 IPv4 地址
  -6                    仅解析 IPv6 地址
  --downloadonly        only download packages
  -h, --help            show help
导入 GPG 公钥 0xA7317B0F:
 Userid: "Google Cloud Packages Automatic Signing Key <gc-team@google.com>"
 指纹: D0BC 747F D8CA F711 7500 D6FA 3746 C208 A731 7B0F
 来自: https://packages.cloud.google.com/yum/doc/yum-key.gpg
确定吗？[y/N]： y
导入 GPG 公钥 0x3E1BA8D5:
 Userid: "Google Cloud Packages RPM Signing Key <gc-team@google.com>"
 指纹: 3749 E1BA 95A8 6CE0 5454 6ED2 F09C 394C 3E1B A8D5
 来自: https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
确定吗？[y/N]： y
```

## DEB

gpg
```
fanhonglingdeMacBook-Pro:http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ ./download-apt-gpg.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   663  100   663    0     0    186      0  0:00:03  0:00:03 --:--:--   186
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   975  100   975    0     0    321      0  0:00:03  0:00:03 --:--:--   321
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   991  100   991    0     0    103      0  0:00:09  0:00:09 --:--:--   247
```

dists
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ ./download-apt-dists.sh 
--2017-11-08 18:47:36--  https://packages.cloud.google.com/apt/dists
正在解析主机 packages.cloud.google.com (packages.cloud.google.com)... 172.217.25.238, 2404:6800:4004:81b::200e
正在连接 packages.cloud.google.com (packages.cloud.google.com)|172.217.25.238|:443... 已连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：17757 (17K) [text/html]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/dists”

dists                                100%[====================================================================>]  17.34K  10.8KB/s    in 1.6s    

缺少“Last-modified”文件头 -- 关闭时间戳标记。
2017-11-08 18:47:40 (10.8 KB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/dists” [17757/17757])

正在载入 robots.txt；请忽略错误消息。
--2017-11-08 18:47:40--  https://packages.cloud.google.com/robots.txt
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 404 Not Found
2017-11-08 18:47:41 错误 404：Not Found。

正在删除 /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/dists 因为它应该被指定了拒绝下载。

--2017-11-08 18:47:41--  https://packages.cloud.google.com/apt/dists/kubernetes-trusty
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：403 [text/html]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty”

kubernetes-trusty                    100%[====================================================================>]     403  --.-KB/s    in 0s      

缺少“Last-modified”文件头 -- 关闭时间戳标记。
2017-11-08 18:47:41 (27.2 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty” [403/403])

--2017-11-08 18:47:41--  https://packages.cloud.google.com/apt/dists/kubernetes-xenial
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：403 [text/html]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial”

kubernetes-xenial                    100%[====================================================================>]     403  --.-KB/s    in 0s      

缺少“Last-modified”文件头 -- 关闭时间戳标记。
2017-11-08 18:47:42 (4.23 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial” [403/403])

--2017-11-08 18:47:42--  https://packages.cloud.google.com/apt/dists/kubernetes-trusty/InRelease
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：6296 (6.1K) [text/plain]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/InRelease”

kubernetes-trusty/InRelease          100%[====================================================================>]   6.15K  --.-KB/s    in 0.001s  

2017-11-08 18:47:43 (4.30 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/InRelease” [6296/6296])

--2017-11-08 18:47:43--  https://packages.cloud.google.com/apt/dists/kubernetes-trusty/Release
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：5793 (5.7K) [text/plain]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/Release”

kubernetes-trusty/Release            100%[====================================================================>]   5.66K  --.-KB/s    in 0.001s  

2017-11-08 18:47:44 (8.08 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/Release” [5793/5793])

--2017-11-08 18:47:44--  https://packages.cloud.google.com/apt/dists/kubernetes-trusty/Release.gpg
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：454 [application/pgp-signature]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/Release.gpg”

kubernetes-trusty/Release.gpg        100%[====================================================================>]     454  --.-KB/s    in 0s      

2017-11-08 18:47:45 (32.3 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/Release.gpg” [454/454])

--2017-11-08 18:47:45--  https://packages.cloud.google.com/apt/dists/kubernetes-trusty/main
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：546 [text/html]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/main”

kubernetes-trusty/main               100%[====================================================================>]     546  --.-KB/s    in 0s      

缺少“Last-modified”文件头 -- 关闭时间戳标记。
2017-11-08 18:47:49 (30.6 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/main” [546/546])

--2017-11-08 18:47:49--  https://packages.cloud.google.com/apt/dists/kubernetes-xenial/InRelease
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：8942 (8.7K) [text/plain]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/InRelease”

kubernetes-xenial/InRelease          100%[====================================================================>]   8.73K  --.-KB/s    in 0.001s  

2017-11-08 18:47:51 (8.24 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/InRelease” [8942/8942])

--2017-11-08 18:47:51--  https://packages.cloud.google.com/apt/dists/kubernetes-xenial/Release
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：8439 (8.2K) [text/plain]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/Release”

kubernetes-xenial/Release            100%[====================================================================>]   8.24K  --.-KB/s    in 0.03s   

2017-11-08 18:47:52 (239 KB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/Release” [8439/8439])

--2017-11-08 18:47:52--  https://packages.cloud.google.com/apt/dists/kubernetes-xenial/Release.gpg
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：454 [application/pgp-signature]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/Release.gpg”

kubernetes-xenial/Release.gpg        100%[====================================================================>]     454  --.-KB/s    in 0s      

2017-11-08 18:47:58 (38.1 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/Release.gpg” [454/454])

--2017-11-08 18:47:58--  https://packages.cloud.google.com/apt/dists/kubernetes-xenial/main
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：650 [text/html]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main”

kubernetes-xenial/main               100%[====================================================================>]     650  --.-KB/s    in 0s      

缺少“Last-modified”文件头 -- 关闭时间戳标记。
2017-11-08 18:47:59 (56.4 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main” [650/650])

--2017-11-08 18:47:59--  https://packages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：421 [text/html]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64”

kubernetes-trusty/main/binary-amd64  100%[====================================================================>]     421  --.-KB/s    in 0s      

缺少“Last-modified”文件头 -- 关闭时间戳标记。
2017-11-08 18:48:00 (30.0 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64” [421/421])

--2017-11-08 18:48:00--  https://packages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：421 [text/html]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64”

kubernetes-xenial/main/binary-amd64  100%[====================================================================>]     421  --.-KB/s    in 0s      

缺少“Last-modified”文件头 -- 关闭时间戳标记。
2017-11-08 18:48:01 (5.51 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64” [421/421])

--2017-11-08 18:48:01--  https://packages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64/Packages
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：1053 (1.0K) [text/plain]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64/Packages”

kubernetes-trusty/main/binary-amd64/ 100%[====================================================================>]   1.03K  --.-KB/s    in 0s      

2017-11-08 18:48:02 (4.00 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64/Packages” [1053/1053])

--2017-11-08 18:48:02--  https://packages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64/Packages.gz
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：458 [application/gzip]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64/Packages.gz”

kubernetes-trusty/main/binary-amd64/ 100%[====================================================================>]     458  --.-KB/s    in 0s      

2017-11-08 18:48:07 (32.1 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64/Packages.gz” [458/458])

--2017-11-08 18:48:07--  https://packages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64/Release
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：207 [text/plain]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64/Release”

kubernetes-trusty/main/binary-amd64/ 100%[====================================================================>]     207  --.-KB/s    in 0s      

2017-11-08 18:48:08 (13.4 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64/Release” [207/207])

--2017-11-08 18:48:08--  https://packages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：52080 (51K) [text/plain]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages”

kubernetes-xenial/main/binary-amd64/ 100%[====================================================================>]  50.86K  --.-KB/s    in 0.07s   

2017-11-08 18:48:09 (778 KB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages” [52080/52080])

--2017-11-08 18:48:09--  https://packages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages.gz
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：7460 (7.3K) [application/gzip]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages.gz”

kubernetes-xenial/main/binary-amd64/ 100%[====================================================================>]   7.29K  --.-KB/s    in 0.001s  

2017-11-08 18:48:10 (9.28 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages.gz” [7460/7460])

--2017-11-08 18:48:10--  https://packages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Release
再次使用存在的到 packages.cloud.google.com:443 的连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：207 [text/plain]
正在保存至: “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Release”

kubernetes-xenial/main/binary-amd64/ 100%[====================================================================>]     207  --.-KB/s    in 0s      

2017-11-08 18:48:11 (15.9 MB/s) - 已保存 “/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Release” [207/207])

下载完毕 --2017-11-08 18:48:11--
总用时：35s
下载了：19 个文件，1.7s (64.5 KB/s) 中的 110K
```

pool
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ ./apt-pool-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 13.7M  100 13.7M    0     0  23437      0  0:10:17  0:10:17 --:--:-- 29466
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 17.2M  100 17.2M    0     0  24528      0  0:12:18  0:12:18 --:--:-- 13524
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 8407k  100 8407k    0     0  30637      0  0:04:41  0:04:40  0:00:01 24967
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 18.3M  100 18.3M    0     0  24360      0  0:13:08  0:13:08 --:--:-- 32846
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 5429k  100 5429k    0     0  28114      0  0:03:17  0:03:17 --:--:-- 38607
```

Packages
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ egrep '^Filename: pool/\S+deb' https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages | awk '{print $2}'
pool/docker-engine_1.11.2-0~xenial_amd64_5554a8bc383e65fb10d556239c72b26457cc5d97f49e3a353c3382f8434e7d21.deb
pool/kubeadm_1.5.7-00_amd64_2759fc99e5b23e44c92b44c506ed9cc1c2087780786bfa97c715da02da84c55d.deb
pool/kubeadm_1.6.1-00_amd64_46af294c4bcb0f923b5593c37a01c27393ba89e1219b10b81f0e6e029c8c1754.deb
pool/kubeadm_1.6.2-00_amd64_5916370a48f2a60174fce83443179f6a0caf79212420c1b5cce73c30f5b79c50.deb
pool/kubeadm_1.6.3-00_amd64_ff5e882c88a5db71803aab900c0b341eb63038558da73c51c3d351070b0c62af.deb
pool/kubeadm_1.6.4-00_amd64_2cc86f287c52e2f63eb9341d7141c593090efcb5164a5ad6004b763e62be51bd.deb
pool/kubeadm_1.6.5-00_amd64_dcdd54c3b0dd4dded1e8fa3e8e04b6427e97a0bd8495b2f7712f1da2a8210872.deb
pool/kubeadm_1.6.6-00_amd64_60fc9618a55876c066f2ccf9adfa95e8baf40499519a8ac5da82640027bcbe51.deb
pool/kubeadm_1.6.7-00_amd64_ece025894ced2577ed0b3bdafa1eb7f615fd35386dbe4e7716fcf6bfdbedbf02.deb
pool/kubeadm_1.6.11-01_amd64_7a00a0c4e91715b1a6bc0439923069efae4f128779e7b2e93c6479db7cc1f385.deb
pool/kubeadm_1.6.12-00_amd64_da9679c05a321ea98ea38c44795233713db86fb612c1d675e3b358ae5f1984dd.deb
pool/kubeadm_1.7.0-00_amd64_6a9f69bdbc29c9f5c26729329225b9199e6927529516de2437ab8df1b92f9919.deb
pool/kubeadm_1.7.1-00_amd64_e6241f9203632e76335fd17ddf7ee8bd403792fdf0e5c97834b43ec3a7eb4b68.deb
pool/kubeadm_1.7.2-00_amd64_47f0e01203a2777cc49ebbf389959c44886c071efd26c1b80c40b0321c640d53.deb
pool/kubeadm_1.7.3-01_amd64_c4bac3fcbc1a820523a3617495707aff0adab0db52ae02f3a5ee5001ff0a1e74.deb
pool/kubeadm_1.7.4-00_amd64_4a999131c64664ccd0ed1792cf517974e03ef531ded8c624cfef58a20321cd0e.deb
pool/kubeadm_1.7.5-00_amd64_5683ceb924f3cefd14a3307ddef30af20d567ad90692ef6a92f229ac9fa1b3a0.deb
pool/kubeadm_1.7.6-00_amd64_68f5bae22936d6c609364154bdd4ffffb91ffb502da6cf54647de04026bf142c.deb
pool/kubeadm_1.7.7-00_amd64_b8317e83d75cff7d4a01566b17fe7b24440c07526108ad8794b0d07e5b55d3a7.deb
pool/kubeadm_1.7.8-00_amd64_8e826d3ef835ca1c1abd1f4be11ef19f7b0ad5d6b85d52e9b26b6ed19acf7d95.deb
pool/kubeadm_1.7.10-00_amd64_987b54cca425ba69b0b799f0a232a6252211fefdf768a45f2bac16e83a81c8a6.deb
pool/kubeadm_1.8.0-00_amd64_22842ddc6d1ffabc04718f384ac001ffa56324cc61e6c3a7c991337bf3e39e06.deb
pool/kubeadm_1.8.0-01_amd64_15e6b68827964187a6c3a60711fab285d06d19953f7473154d798ff52a418185.deb
pool/kubeadm_1.8.1-01_amd64_5863420bee513756daf8cd08af801c5420397409191df19535d8c618a100989c.deb
pool/kubeadm_1.8.2-00_amd64_e743a9538b855d08ddaa68e7910af2dfc2bb9c1a0938d79089b0a9d3f1c19dde.deb
pool/kubectl_1.5.1-00_amd64_c188aff0b89c6164bf75a0ad2d16080082222c7dd04943b4895ffdedee523ecf.deb
pool/kubectl_1.5.2-00_amd64_7ca7fea62ab4672bb09d49a0ec3a23d09386cb303f97713f6a4406459ef0a8e0.deb
pool/kubectl_1.5.3-00_amd64_10ffbdd74343d2b8b3ba31adc7eeb37b690badd4ce7f9bfb9a54e86ee05d433d.deb
pool/kubectl_1.5.6-00_amd64_aeaf2d825c8426b970c092fc8b0e12d701d6c87ec09027a58244538841ae0ad2.deb
pool/kubectl_1.5.7-00_amd64_0e7de187eed0ca04d6ef148f617c09c781c7aec28a15140691e60d5878c7b30c.deb
pool/kubectl_1.6.0-00_amd64_a068b847837acec84e1922277dc46f6fee2b0c2f930405ff232cf9ac7e3473e5.deb
pool/kubectl_1.6.1-00_amd64_b224b79174ff2d1a5ba8db17f67b7284e4e969512506b056cf7f7754fbdccbdc.deb
pool/kubectl_1.6.2-00_amd64_b9d37ee035847c965a21427aa6cc3f61f1a9a6d2c8eee8288d1f8657628930ae.deb
pool/kubectl_1.6.3-00_amd64_05d48fa118b6538ee2bc0b864aeb09f2cede83990fc8819875f698a5dece0c9b.deb
pool/kubectl_1.6.4-00_amd64_fd34da38db241a88a6e95b825139e0df6a627f141ae9ed39f948107a8437935e.deb
pool/kubectl_1.6.5-00_amd64_c87736ba4285ed607853a585faa0b6fe591454f0b0ba0a2d655ffc88e2c98ee6.deb
pool/kubectl_1.6.6-00_amd64_671ee06fc52fb9298990a3f5044033df14a6dc542cc1eee1349fb8b111e83992.deb
pool/kubectl_1.6.7-00_amd64_3137dc28aeb6a9dd364750d4c8d38581926fa45a2273ac5f6b6319aa5b1a4481.deb
pool/kubectl_1.6.11-00_amd64_0049f491da3c6d5826d5cc0862a84da361daf3c356fddb507746891f456e4bac.deb
pool/kubectl_1.6.12-00_amd64_54410fc4e429df54df34b406b2515cdb3043b11f14ab87ab522f54a6a4e10e92.deb
pool/kubectl_1.7.0-00_amd64_6bd51a1ed15eb74ad2d72823a301c71e769dcc12ec10294eac3026cb77e03860.deb
pool/kubectl_1.7.1-00_amd64_7f6d0471be6524d0d0aec8e9f515e0261a8e364d4fd2f25f22cba4895634bab0.deb
pool/kubectl_1.7.2-00_amd64_cd4dec0c4d75f1553712543c03720c3fede37ff8a46b97f287d7f4bea58f36c6.deb
pool/kubectl_1.7.3-01_amd64_84173e3f104760b832ddf732bf02f1834cbf719032746518a52efcc942dba4c6.deb
pool/kubectl_1.7.4-00_amd64_6ee4015eabaa1e0cd4e2895c102c1c341f1c0037b28cb7700424d0f6f04a2416.deb
pool/kubectl_1.7.5-00_amd64_363445c3be161e924d1209c60e24dd2d0657d68bcf7c17e8ec2a6dc603c5d363.deb
pool/kubectl_1.7.6-00_amd64_a8c59bb8105f119061ed368451923465351488efa537a6cf9f88e3ce0f0422d1.deb
pool/kubectl_1.7.7-00_amd64_22e3fbc7a55ca7973eac5de600896827a15de0f750c6b7e37e46a8cb99a971fb.deb
pool/kubectl_1.7.8-00_amd64_0b904de238294283cbfdfd317664f552766dd974947fa5e0c2f3bd0077331522.deb
pool/kubectl_1.7.10-00_amd64_b725d080205ddf237a4df726dba6c81c9f56d1bc089a530e9f2df0bcd659af3e.deb
pool/kubectl_1.8.0-00_amd64_9bd409d2a0300d4b550cd2a7fd2eca6437ea0767805f6fd807912b245ec949ed.deb
pool/kubectl_1.8.1-00_amd64_c87f3e1e54ad045af07001028282f9b76e1c26425f6ccaeba52bb0a5fa93a6b9.deb
pool/kubectl_1.8.2-00_amd64_b01f6fa567e98752181a3ad057275851e00aa5fa1a8db6eb5a6d81c0f499e1ec.deb
pool/kubelet_1.5.1-00_amd64_bb82dd4bcf0c9bc813c599f62afa48832bf34302d723c5a38347c2754f3735e2.deb
pool/kubelet_1.5.2-00_amd64_ac51b2e458539f4803b2d2b31b3fad0b6d503ca8ddc1994e55b4150f33909163.deb
pool/kubelet_1.5.3-00_amd64_6df7dd70c4db4afb77bdf309569d5e9ad92e7319ab4d3b9bcd0bd56f457a8755.deb
pool/kubelet_1.5.6-00_amd64_a9a67135a020676e3879cca1971447f78bd4b365a243b94a42123cac54d03d9b.deb
pool/kubelet_1.5.7-00_amd64_f872da8a98d992681d49218c24703762fc1a9699f9a6af46de24e9401d6fb0aa.deb
pool/kubelet_1.6.0-00_amd64_850ba773992934d5eb46b064274d3449bb5c07e2f1a1019b10e068e0a0efc78d.deb
pool/kubelet_1.6.1-00_amd64_8f4433bbf40b5c2fdf75089f520544969cca62ee53b7ea87eb4f1123fa2f4163.deb
pool/kubelet_1.6.2-00_amd64_8537952290ccb565bb9e6d7beaeeaa2b037eb5966df097d5ca95a7cb811c51d5.deb
pool/kubelet_1.6.3-00_amd64_a94b0cfa5b26939d87097dfd45260474c1effcad879e35940eb6d36e7d953d6c.deb
pool/kubelet_1.6.4-00_amd64_2affcdcf43b6cc5a89237d0ceb1b5a3a90c605d4da73b97c5842f1a53368231d.deb
pool/kubelet_1.6.5-00_amd64_2d1e2ea0fdd3f85594ed9d57b4bc19e3652f9bd66ad9099f0abfc268e6d9c379.deb
pool/kubelet_1.6.6-00_amd64_7e64ef3101aff3c7e59b7e38c3ff4acf7659730aeda47fd43bb7a12384069f11.deb
pool/kubelet_1.6.7-00_amd64_b3c7f8b1743cbb05af4fa62f159952572a4520a3289a742a0544fcd8068703e8.deb
pool/kubelet_1.6.11-00_amd64_2dd18ba275fd1bf41acd6280975fd7a45467ca2f2842e61703c2787856f4e2a7.deb
pool/kubelet_1.6.12-00_amd64_c302b5075f37fc8bdc251b2170b1b7b358c0faaf993ee1f0cf20204e971009ef.deb
pool/kubelet_1.7.0-00_amd64_b7c9be22a6379d8cc47ba1183bbc9de0acf49ad1418ffa8ae57628ad1d7c051f.deb
pool/kubelet_1.7.1-00_amd64_8e6aeef38ca042b85765751652f8115960892ca4cbd1a906b63197931351cb8d.deb
pool/kubelet_1.7.2-00_amd64_d685940c01bd21e18caa8528cf793b09f9d81f9c431173356461a8ecc0704f28.deb
pool/kubelet_1.7.3-01_amd64_73541f99b3465ea2388e3f48d2207862a41c6e980782e82678d3a1bcab39630b.deb
pool/kubelet_1.7.4-00_amd64_8e9b2a6d7b233240f816591b8180d9e92b08a73afbe93ed0c783537831bbb4c6.deb
pool/kubelet_1.7.5-00_amd64_0db3259738d749acca6ffff03ae4ccb87a4e6a64c3ac9884b34df7da936a3052.deb
pool/kubelet_1.7.6-00_amd64_65531a88ce1ba75817f0d6d32e10efcb1ea5aabaf70e9e9aa53f83180ca41834.deb
pool/kubelet_1.7.7-00_amd64_c66f68c9a638cc3acae96c86d6207f91a4e0efb27f43fadb02a78f72c71cc894.deb
pool/kubelet_1.7.8-00_amd64_429a02bce3df7d943bbec28b2b35a9c9e1ba78a01484bfe1e4e7471a362c38bd.deb
pool/kubelet_1.7.10-00_amd64_4322e4474136b891848d5d67658b9cd22a41a834c88ecea382d58da092156556.deb
pool/kubelet_1.8.0-00_amd64_00b7c77c924d654c7def52c83cfeb9a3d1836c1e7b40683c3fe8207b0bd299d1.deb
pool/kubelet_1.8.1-00_amd64_c0f1902d204f6741eff08560120bfe193b992596068158a2042988203132cd51.deb
pool/kubelet_1.8.2-00_amd64_9646b1262ea4c99a89e94801b150e355a4e984c462bf8fb8a2fe9a33dacd74e0.deb
pool/kubernetes-cni_0.3.0.1-07a8a2-00_amd64_9e41a275b2afeb51dcde86b922c056c7b6dc656b54dd66fa2f1a0bb8266e9c22.deb
pool/kubernetes-cni_0.5.1-00_amd64_08cbe5c42366ec3184cc91a4353f6e066f2d7325b4db1cb4f87c7dcc8c0eb620.deb
pool/rkt_1.25.0-1_amd64_65f4d768116520be6ee5d56ebcdce87b8b39fa2e41345b7c2960a01f91b7c081.deb
pool/rkt_1.26.0-1_amd64_af62ca3979b90f2fcbe7c48ba782dc1b0c4266832e7544ef4c69ce0a8a7dfa87.deb
pool/rkt_1.27.0-1_amd64_b4f1d67b835d1f1b0263ddfc3f8df80cd7c9e36ac78d983e12aca5575835e122.deb
pool/rkt_1.28.1-1_amd64_66ea8b0ef5724aa364a1273b794b2aaf5c302c774e076115adee8f7c04a0e0f9.deb
pool/rkt_1.29.0-1_amd64_ea87d719359030f33fd48890875c934135c62eccda72c37d79ff604307b905b5.deb
```

### 3rd-party

Refer to http://linoxide.com/containers/setup-kubernetes-kubeadm-centos/
```
[vagrant@localhost ~]$ sudo vi /etc/yum.repos.d/kubernetes.repo

[vagrant@localhost ~]$ sudo dnf repository-packages kubernetes list
Kubernetes                                                                                      1.9 kB/s | 3.4 kB     00:01    
上次元数据过期检查在 0:00:00 前执行于 Thu Dec 29 22:55:09 2016。
可安装的软件包
kubeadm.x86_64                                      1.6.0-0.alpha.0.2074.a092d8e0f95f52                               kubernetes
kubectl.x86_64                                      1.5.1-0                                                           kubernetes
kubelet.x86_64                                      1.5.1-0                                                           kubernetes
kubernetes-cni.x86_64                               0.3.0.1-0.07a8a2                                                  kubernetes
rkt.x86_64                                          1.21.0-1                                                          kubernetes

[vagrant@localhost http%3A%2F%2Fyum.kubernetes.io]$ ./download-el7-pkgs.sh 
Loaded plugins: fastestmirror, ovl
Loading mirror speeds from cached hostfile
 * base: mirror.0x.sg
 * extras: mirror.0x.sg
 * updates: mirror.0x.sg
Error: No package(s) available to install
[vagrant@localhost http%3A%2F%2Fyum.kubernetes.io]$ vi download-el7-pkgs.sh 
[vagrant@localhost http%3A%2F%2Fyum.kubernetes.io]$ ./download-el7-pkgs.sh 
Loaded plugins: fastestmirror, ovl
Loading mirror speeds from cached hostfile
 * base: mirror.0x.sg
 * extras: mirror.0x.sg
 * updates: mirror.0x.sg
No package # available.
Resolving Dependencies
--> Running transaction check
---> Package kubeadm.x86_64 0:1.6.0-0.alpha.0.2074.a092d8e0f95f52 will be installed
---> Package kubectl.x86_64 0:1.5.1-0 will be installed
---> Package kubelet.x86_64 0:1.5.1-0 will be installed
--> Processing Dependency: iptables >= 1.4.21 for package: kubelet-1.5.1-0.x86_64
--> Processing Dependency: socat for package: kubelet-1.5.1-0.x86_64
--> Processing Dependency: iproute for package: kubelet-1.5.1-0.x86_64
--> Processing Dependency: ethtool for package: kubelet-1.5.1-0.x86_64
--> Processing Dependency: ebtables for package: kubelet-1.5.1-0.x86_64
---> Package kubernetes-cni.x86_64 0:0.3.0.1-0.07a8a2 will be installed
---> Package rkt.x86_64 0:1.21.0-1 will be installed
--> Running transaction check
---> Package ebtables.x86_64 0:2.0.10-15.el7 will be installed
---> Package ethtool.x86_64 2:4.5-3.el7 will be installed
---> Package iproute.x86_64 0:3.10.0-74.el7 will be installed
--> Processing Dependency: libmnl.so.0(LIBMNL_1.0)(64bit) for package: iproute-3.10.0-74.el7.x86_64
--> Processing Dependency: libmnl.so.0()(64bit) for package: iproute-3.10.0-74.el7.x86_64
---> Package iptables.x86_64 0:1.4.21-17.el7 will be installed
--> Processing Dependency: libnfnetlink.so.0()(64bit) for package: iptables-1.4.21-17.el7.x86_64
--> Processing Dependency: libnetfilter_conntrack.so.3()(64bit) for package: iptables-1.4.21-17.el7.x86_64
---> Package socat.x86_64 0:1.7.2.2-5.el7 will be installed
--> Running transaction check
---> Package libmnl.x86_64 0:1.0.3-7.el7 will be installed
---> Package libnetfilter_conntrack.x86_64 0:1.0.4-2.el7 will be installed
---> Package libnfnetlink.x86_64 0:1.0.1-4.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================================================================
 Package                           Arch              Version                                        Repository             Size
================================================================================================================================
Installing:
 kubeadm                           x86_64            1.6.0-0.alpha.0.2074.a092d8e0f95f52            kubernetes            5.8 M
 kubectl                           x86_64            1.5.1-0                                        kubernetes            6.6 M
 kubelet                           x86_64            1.5.1-0                                        kubernetes             12 M
 kubernetes-cni                    x86_64            0.3.0.1-0.07a8a2                               kubernetes            9.8 M
 rkt                               x86_64            1.21.0-1                                       kubernetes             87 M
Installing for dependencies:
 ebtables                          x86_64            2.0.10-15.el7                                  base                  123 k
 ethtool                           x86_64            2:4.5-3.el7                                    base                  121 k
 iproute                           x86_64            3.10.0-74.el7                                  base                  618 k
 iptables                          x86_64            1.4.21-17.el7                                  base                  426 k
 libmnl                            x86_64            1.0.3-7.el7                                    base                   23 k
 libnetfilter_conntrack            x86_64            1.0.4-2.el7                                    base                   53 k
 libnfnetlink                      x86_64            1.0.1-4.el7                                    base                   26 k
 socat                             x86_64            1.7.2.2-5.el7                                  base                  255 k

Transaction Summary
================================================================================================================================
Install  5 Packages (+8 Dependent packages)

Total download size: 123 M
Installed size: 276 M
Background downloading packages, then exiting:
warning: /var/cache/yum/x86_64/7/base/packages/ebtables-2.0.10-15.el7.x86_64.rpm.1.tmp: Header V3 RSA/SHA256 Signature, key ID f4a80eb5: NOKEY
Public key for ebtables-2.0.10-15.el7.x86_64.rpm.1.tmp is not installed
(1/13): ebtables-2.0.10-15.el7.x86_64.rpm                                                                | 123 kB  00:00:01     
(2/13): ethtool-4.5-3.el7.x86_64.rpm                                                                     | 121 kB  00:00:01     
(3/13): iptables-1.4.21-17.el7.x86_64.rpm                                                                | 426 kB  00:00:02     
(4/13): iproute-3.10.0-74.el7.x86_64.rpm                                                                 | 618 kB  00:00:04     
(5/13): 93af9d0fbd67365fa5bf3f85e3d36060138a62ab77e133e35f6cadc1fdc15299-kubectl-1.5.1-0.x86_64.rpm      | 6.6 MB  00:00:08     
(6/13): 8a299eb1db946b2bdf01c5d5c58ef959e7a9d9a0dd706e570028ebb14d48c42e-kubelet-1.5.1-0.x86_64.rpm      |  12 MB  00:00:17     
(7/13): libmnl-1.0.3-7.el7.x86_64.rpm                                                                    |  23 kB  00:00:00     
(8/13): libnfnetlink-1.0.1-4.el7.x86_64.rpm                                                              |  26 kB  00:00:00     
(9/13): libnetfilter_conntrack-1.0.4-2.el7.x86_64.rpm                                                    |  53 kB  00:00:00     
(10/13): 5612db97409141d7fd839e734d9ad3864dcc16a630b2a91c312589a0a0d960d0-kubeadm-1.6.0-0.alpha.0.2074.a | 5.8 MB  00:00:37     
(11/13): socat-1.7.2.2-5.el7.x86_64.rpm                                                                  | 255 kB  00:00:03     
(12/13): 567600102f687e0f27bd1fd3d8211ec1cb12e71742221526bb4e14a412f4fdb5-kubernetes-cni-0.3.0.1-0.07a8a | 9.8 MB  00:00:40     
(13/13): efd51c756948693d0c7334edfd01e77cc875aa471e46d2d1800429427bfbbde1-rkt-1.21.0-1.x86_64.rpm        |  87 MB  00:01:45     
--------------------------------------------------------------------------------------------------------------------------------
Total                                                                                           882 kB/s | 123 MB  00:02:23     
exiting because "Download Only" specified

[vagrant@localhost http%3A%2F%2Fyum.kubernetes.io]$ ./download-el7-pkgs.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   663  100   663    0     0    891      0 --:--:-- --:--:-- --:--:--   891
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   975  100   975    0     0    752      0  0:00:01  0:00:01 --:--:--   752
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   991  100   991    0     0   1167      0 --:--:-- --:--:-- --:--:--  1167
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  6436  100  6436    0     0   4895      0  0:00:01  0:00:01 --:--:--  4898
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1442  100  1442    0     0   1265      0  0:00:01  0:00:01 --:--:--  1266
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  7677  100  7677    0     0   6421      0  0:00:01  0:00:01 --:--:--  6418
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1195  100  1195    0     0    925      0  0:00:01  0:00:01 --:--:--   926
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  8695  100  8695    0     0   7456      0  0:00:01  0:00:01 --:--:--  7463
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  2074  100  2074    0     0   1693      0  0:00:01  0:00:01 --:--:--  1694
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1410  100  1410    0     0   1109      0  0:00:01  0:00:01 --:--:--  1108
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1449  100  1449    0     0   1834      0 --:--:-- --:--:-- --:--:--  1836
Loaded plugins: fastestmirror, ovl
Loading mirror speeds from cached hostfile
 * base: mirror.0x.sg
 * extras: mirror.0x.sg
 * updates: mirror.0x.sg
No package ###rkt available.
Resolving Dependencies
--> Running transaction check
---> Package kubeadm.x86_64 0:1.6.0-0.alpha.0.2074.a092d8e0f95f52 will be installed
---> Package kubectl.x86_64 0:1.5.1-0 will be installed
---> Package kubelet.x86_64 0:1.5.1-0 will be installed
--> Processing Dependency: iptables >= 1.4.21 for package: kubelet-1.5.1-0.x86_64
--> Processing Dependency: socat for package: kubelet-1.5.1-0.x86_64
--> Processing Dependency: iproute for package: kubelet-1.5.1-0.x86_64
--> Processing Dependency: ethtool for package: kubelet-1.5.1-0.x86_64
--> Processing Dependency: ebtables for package: kubelet-1.5.1-0.x86_64
---> Package kubernetes-cni.x86_64 0:0.3.0.1-0.07a8a2 will be installed
--> Running transaction check
---> Package ebtables.x86_64 0:2.0.10-15.el7 will be installed
---> Package ethtool.x86_64 2:4.5-3.el7 will be installed
---> Package iproute.x86_64 0:3.10.0-74.el7 will be installed
--> Processing Dependency: libmnl.so.0(LIBMNL_1.0)(64bit) for package: iproute-3.10.0-74.el7.x86_64
--> Processing Dependency: libmnl.so.0()(64bit) for package: iproute-3.10.0-74.el7.x86_64
---> Package iptables.x86_64 0:1.4.21-17.el7 will be installed
--> Processing Dependency: libnfnetlink.so.0()(64bit) for package: iptables-1.4.21-17.el7.x86_64
--> Processing Dependency: libnetfilter_conntrack.so.3()(64bit) for package: iptables-1.4.21-17.el7.x86_64
---> Package socat.x86_64 0:1.7.2.2-5.el7 will be installed
--> Running transaction check
---> Package libmnl.x86_64 0:1.0.3-7.el7 will be installed
---> Package libnetfilter_conntrack.x86_64 0:1.0.4-2.el7 will be installed
---> Package libnfnetlink.x86_64 0:1.0.1-4.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================================================================
 Package                           Arch              Version                                        Repository             Size
================================================================================================================================
Installing:
 kubeadm                           x86_64            1.6.0-0.alpha.0.2074.a092d8e0f95f52            kubernetes            5.8 M
 kubectl                           x86_64            1.5.1-0                                        kubernetes            6.6 M
 kubelet                           x86_64            1.5.1-0                                        kubernetes             12 M
 kubernetes-cni                    x86_64            0.3.0.1-0.07a8a2                               kubernetes            9.8 M
Installing for dependencies:
 ebtables                          x86_64            2.0.10-15.el7                                  base                  123 k
 ethtool                           x86_64            2:4.5-3.el7                                    base                  121 k
 iproute                           x86_64            3.10.0-74.el7                                  base                  618 k
 iptables                          x86_64            1.4.21-17.el7                                  base                  426 k
 libmnl                            x86_64            1.0.3-7.el7                                    base                   23 k
 libnetfilter_conntrack            x86_64            1.0.4-2.el7                                    base                   53 k
 libnfnetlink                      x86_64            1.0.1-4.el7                                    base                   26 k
 socat                             x86_64            1.7.2.2-5.el7                                  base                  255 k

Transaction Summary
================================================================================================================================
Install  4 Packages (+8 Dependent packages)

Total size: 36 M
Installed size: 171 M
Background downloading packages, then exiting:
exiting because "Download Only" specified
[vagrant@localhost http%3A%2F%2Fyum.kubernetes.io]$ ls
download-el7-pkgs.sh  https%3A%2F%2Fpackages.cloud.google.com%2Fyum  kubernetes-el7-x86_64  kubernetes.repo  README.md


```

