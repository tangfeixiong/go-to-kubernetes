# DevOps

Note: about kubernetes repository, refer to https://kubernetes.io/docs/setup/independent/install-kubeadm

## Table of Contents

1. [Mirroring Kubernetes RPM repository for CentOS/Fedora](#rpm)
1. [Mirroring Kubernetes DEB repository for Ubuntu/Debian](#deb)

## RPM

### GPG

Current download is on releasing V1.8.4
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

### Repos

Now download again to V1.10.2
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ ./yum-repos-el7-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 71412  100 71412    0     0  37345      0  0:00:01  0:00:01 --:--:-- 37329
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 10899  100 10899    0     0  10007      0  0:00:01  0:00:01 --:--:-- 10008
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  330k  100  330k    0     0   172k      0  0:00:01  0:00:01 --:--:--  172k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 20382  100 20382    0     0  16370      0  0:00:01  0:00:01 --:--:-- 16384
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  359k  100  359k    0     0   104k      0  0:00:03  0:00:03 --:--:--  104k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 29584  100 29584    0     0  20809      0  0:00:01  0:00:01 --:--:-- 20804
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1418  100  1418    0     0   1374      0  0:00:01  0:00:01 --:--:--  1375
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   454  100   454    0     0    317      0  0:00:01  0:00:01 --:--:--   317
```

### Pool

Find required package from _repos_
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ egrep '^  <location href=".*\.rpm"></location>' https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/primary.xml | cut -d '"' -f 2
../../pool/5116fa4b73c700823cfc76b3cedff6622f2fbd0a3d2fa09bce6d93329771e291-kubeadm-1.6.0-0.x86_64.rpm
../../pool/23961d0f7dca1ed118b948195f2fb5dd7a07503d69d7d8ab4433219ea98d033e-kubeadm-1.6.1-0.x86_64.rpm
../../pool/415bc9f86ecfc3b195f22d25fb599e19080d301358c145539d154be95f1985f5-kubeadm-1.6.2-0.x86_64.rpm
../../pool/81158f40764a08356242a53ef4bf7e3c219f48f364c55260db571cae51ce6eec-kubeadm-1.6.3-0.x86_64.rpm
../../pool/0c37012e387fae5fd0d66b5d4664f495ab19fbb709df655533b2b2ecb1a05cd8-kubeadm-1.6.4-0.x86_64.rpm
../../pool/5e62d624af011598f2df59c7759bf6ff4a6bba5675f0c07a2fb18e3c468ea41c-kubeadm-1.6.5-0.x86_64.rpm
../../pool/03ea47d2b0f1912c3a721c2b7353ead2e28a154b6c883200492db0558f6b09f2-kubeadm-1.6.6-0.x86_64.rpm
../../pool/6b36b8d9783ee59480bda0caabcb9ce95a6bf69c62ea38e7bb0cee65174fe479-kubeadm-1.6.7-0.x86_64.rpm
../../pool/67cb3635a4c3927b7009f0163589dfc0d261726a055893c52531c77b6e6bb0f8-kubeadm-1.6.8-0.x86_64.rpm
../../pool/dce6dfbff7a60ac39f9288b211d8e4cf788a03ee64f2456ee6e0860fc549e1a9-kubeadm-1.6.9-0.x86_64.rpm
../../pool/2bb1e670317f17fde3142c3e3549c91e15fc629ddfd0f274c290918c4289caf4-kubeadm-1.6.10-0.x86_64.rpm
../../pool/c10b6d67b13101afd0270873972f10b14e9719ad8564797cc6609f7c58ae033c-kubeadm-1.6.11-0.x86_64.rpm
../../pool/fda8179e8174e43c17662d3a019f54e862d270fbfc07c3f7a88a58a033198153-kubeadm-1.6.12-0.x86_64.rpm
../../pool/ed8ab0500a5aa2d443653aa144c65fc246e3d047bee7cb1ed6499d0c926d0690-kubeadm-1.6.13-0.x86_64.rpm
../../pool/cb7034dff51af01c237e03ea342dc74269b46cce1e43bb0ab9ef1d6f006ebda9-kubeadm-1.7.0-0.x86_64.rpm
../../pool/77bea7f33178ab4feb6afaf31c6c779511498ef47dec433eb5464909f7a26dc3-kubeadm-1.7.1-0.x86_64.rpm
../../pool/1a6f5f73f43077a50d877df505481e5a3d765c979b89fda16b8b9622b9ebd9a4-kubeadm-1.7.2-0.x86_64.rpm
../../pool/f7ec56b0f36a81c0f91bcf26e05f23088082b468b77dac576dc505444dd8cd48-kubeadm-1.7.3-1.x86_64.rpm
../../pool/f0a51fcde5e3b329050d7a6cf70f04a6cdf09eacfbad55f4324bfa2ea4312d0e-kubeadm-1.7.4-0.x86_64.rpm
../../pool/02f3a7ff6e04943bd288ff302f449b600e8db3d19868dfe4308d0d902c0ba927-kubeadm-1.7.5-0.x86_64.rpm
../../pool/4ff875dc8857b85c490b42b750527ba20a154a49a8dacd256d16cbbf5e708dfd-kubeadm-1.7.6-1.x86_64.rpm
../../pool/2bb9ddc5197dec31ac73a549067ab9b6a529cc31275f9223ff04efbfb5451602-kubeadm-1.7.7-1.x86_64.rpm
../../pool/8af3d230e7c1c775124b518aa0d69fa911e761e7b22fd081bfae8dece2381970-kubeadm-1.7.8-1.x86_64.rpm
../../pool/e9e277f88f2747a493dd6b22360c58701c30216849e6ae74b190f4916b58125b-kubeadm-1.7.9-0.x86_64.rpm
../../pool/c0638ef93a73ad4343fe3c2b8105e3826605e2fe000017324b3328975aa36a82-kubeadm-1.7.10-0.x86_64.rpm
../../pool/b10ac8cf7ee52d4f4144b523b1f33061b7429bea3fcf9e863261423de090804a-kubeadm-1.7.11-0.x86_64.rpm
../../pool/0ce3d048f9832d29000f6c69e05eb7a1fbc013810c7d1f66ad43e17f8578bd05-kubeadm-1.7.14-0.x86_64.rpm
../../pool/7de188358ee8ee1babf424ab9d8d83ab2eb82325cba7133d97d8d5de3a365918-kubeadm-1.7.15-0.x86_64.rpm
../../pool/9dda7d73d6c4cce878eb2fedaf2dfed78278202b0b5ee73dcd841b4edd643ae5-kubeadm-1.7.16-0.x86_64.rpm
../../pool/ee93b5249497dba6488262bd2b56e32438a69f78993eb973bafa52d72f9b883a-kubeadm-1.8.0-0.x86_64.rpm
../../pool/7fe94dd0f828ef64d4ccdfef160417457309b904758ddf10c040612d7e5ef441-kubeadm-1.8.0-1.x86_64.rpm
../../pool/0f7d8ea10144399f3d60446fab5469395afb809c175bdc0eae4d12c1fcc3cb62-kubeadm-1.8.1-0.x86_64.rpm
../../pool/d64bc1d0ca27196030c6f574252a0872b998b29be90c6cb7e97b91cd0bc78fed-kubeadm-1.8.2-0.x86_64.rpm
../../pool/cab6b288e91e613d81c63101c7d293059a4a9f2c0795228042c880f770a9ec60-kubeadm-1.8.3-0.x86_64.rpm
../../pool/aeaad1e283c54876b759a089f152228d7cd4c049f271125c23623995b8e76f96-kubeadm-1.8.4-0.x86_64.rpm
../../pool/f9a3e9f13f7dacb8a39b90a2331007bebbed4f84643448e83e5c18b3efe3d45b-kubeadm-1.8.5-0.x86_64.rpm
../../pool/919d83307b30c808a9bf17f08ab6d72612d08860a8923366e666ed072012f62a-kubeadm-1.8.6-0.x86_64.rpm
../../pool/005fee933b99ae653baf508af37144c194d330b6324a62b32ebc08557a838937-kubeadm-1.8.7-0.x86_64.rpm
../../pool/2e001cae2c8933ef43f6d5607a8bd57200957685bfb932dfb753b23abb9f8074-kubeadm-1.8.8-0.x86_64.rpm
../../pool/e1cf1b76ab1c7d92a24d0ccccd2f3cdaae5b02bd063797938d27760531b76049-kubeadm-1.8.9-0.x86_64.rpm
../../pool/1e84e0b4101a6ded3d395f77ff5736e6d15d9730b5b9fb00ddf625878ef585f4-kubeadm-1.8.10-0.x86_64.rpm
../../pool/280eb1e1232c7dc99ec6e40bcd1ae3b7f72031a54bf33c7adb9abca64b0b7411-kubeadm-1.8.11-0.x86_64.rpm
../../pool/59570c2320c2e8ead952893f25df6e4337c5ab3056ece69694520a05cadd0396-kubeadm-1.8.12-0.x86_64.rpm
../../pool/bf953a68a7f2b897d88ad14922562387142334ae322057e5cb7c733e7fbb4ca1-kubeadm-1.8.13-0.x86_64.rpm
../../pool/aa9948f82e7af317c97a242f3890985159c09c183b46ac8aab19d2ad307e6970-kubeadm-1.9.0-0.x86_64.rpm
../../pool/ccc2b7d8ac99c8ead43087e7bb5cc7fb3fe684bfd78241c8240feb945303c40e-kubeadm-1.9.1-0.x86_64.rpm
../../pool/eb3f5ec9f28b4e6548acc87211b74534c6eae3cabfe870addeff84ac3243a77e-kubeadm-1.9.2-0.x86_64.rpm
../../pool/df7f026762a109d02bf92ec559bbeb3f490e8e35885ffc2631042ffcd9b945ab-kubeadm-1.9.3-0.x86_64.rpm
../../pool/12f95589196fb1331fb136084c16dcc647ef9db391307a64f6a23391a7afc6de-kubeadm-1.9.4-0.x86_64.rpm
../../pool/4c45f9cada16d14ffd030aade7728b8ac5d148be70206cd3fbeb4188ee790f00-kubeadm-1.9.5-0.x86_64.rpm
../../pool/f56f3294d633ecfa7f2aac506f7267c00547d4c529b134bc4698a563402897c3-kubeadm-1.9.6-0.x86_64.rpm
../../pool/e569b64b2da8847696c665c44c0f0d5ba6fad79525230898219c514d79f1b1c3-kubeadm-1.9.7-0.x86_64.rpm
../../pool/8b0cb6654a2f6d014555c6c85148a5adc5918de937f608a30b0c0ae955d8abce-kubeadm-1.10.0-0.x86_64.rpm
../../pool/35ea034b2efccf2529cca8ed44f1bdcc0c3b26f0139694d8cbea315077a1bf6e-kubeadm-1.10.1-0.x86_64.rpm
../../pool/b754a6990af7d7012189610b0dc69e6e950c13a8c415b9ebea8d56352e9719fd-kubeadm-1.10.2-0.x86_64.rpm
../../pool/fd9b1e2215cab6da7adc6e87e2710b91ecfda3b617edfe7e71c92277301a63ab-kubectl-1.5.4-0.x86_64.rpm
../../pool/e6aef7b2b7d9e5bd4db1e5747ebbc9f1f97bbfb8c7817ad68028565ca263a672-kubectl-1.6.0-0.x86_64.rpm
../../pool/9d1ccf0877dfc4228a923a9614661b27d663694680e2bc0a1c184aa937fbf7f2-kubectl-1.6.1-0.x86_64.rpm
../../pool/ff72cbf0dfa986c36097a5c3ac2301ecb73ed28df8db86e13641a79e9df9b3ea-kubectl-1.6.2-0.x86_64.rpm
../../pool/d5cc6bb2e439298eb1b3e45c3ac58010580c1d3c3a2fa71040a7c73b1dc22881-kubectl-1.6.3-0.x86_64.rpm
../../pool/9fcf36d1acf5e17c4aec783e956b8b70ab0e91693d394e6268c16533fe8a71a9-kubectl-1.6.4-0.x86_64.rpm
../../pool/504e6824515b1b91702d1cd9aa8701fbcb2fabeb61e8395204f4e2cd07f7dfb7-kubectl-1.6.5-0.x86_64.rpm
../../pool/b9584919848f1cc3096ec8cd0bbeceb95ea0786df339eb4883556a4a89f51010-kubectl-1.6.6-0.x86_64.rpm
../../pool/b542021d5f32457f8c1f6f726aaa077eec66b0906440a020cfada28275df50f7-kubectl-1.6.7-0.x86_64.rpm
../../pool/396d4c944b93310b86afad34627390e31a3f839ca1f7f91df95919ac4c85065c-kubectl-1.6.8-0.x86_64.rpm
../../pool/27a094522739c8fa954133ca6b9d28b1e39537d7a9cd33a1b6e2e586e9e6811d-kubectl-1.6.9-0.x86_64.rpm
../../pool/ad4326997f2a06c4802ce8ae170648f6a07522f5de629be89d5324ecfe843780-kubectl-1.6.10-0.x86_64.rpm
../../pool/9b3b2d1cbe36086a7008d261594b31bb3e085e8c170aa319e1befebe604a99a7-kubectl-1.6.11-0.x86_64.rpm
../../pool/712b311a1853d2b73be6729f57fc448ebc8818f0a9236cbbf1f11c6db3c3d8dc-kubectl-1.6.12-0.x86_64.rpm
../../pool/0e7baf0e754f3c62dcbb50d850167f0d8eee0819aab40cb3932ddc72f098f4ea-kubectl-1.6.13-0.x86_64.rpm
../../pool/888aff6bf68f988e7480efd06f38852eef8a44eed1baaa62f3ec4fbb07c35f7d-kubectl-1.7.0-0.x86_64.rpm
../../pool/b8a64634ad1555e15a61b84685fd04959435ed6374b25e369f5bda89c8f03a6b-kubectl-1.7.1-0.x86_64.rpm
../../pool/dc8329515fc3245404fea51839241b58774e577d7736f99f21276e764c309db5-kubectl-1.7.2-0.x86_64.rpm
../../pool/c8a50a1ce9c948e7a238b02d3967d220e71e13e1ac8916a961176726eabe8aa1-kubectl-1.7.3-1.x86_64.rpm
../../pool/041d5a6813dab590b160865fea7259bc2db762a9667379d03aca8d4596a3cccd-kubectl-1.7.4-0.x86_64.rpm
../../pool/c2d4b7c1f39ca9532a2965ea513fcd900fdcdd8785639c8fbf69f30780cae764-kubectl-1.7.5-0.x86_64.rpm
../../pool/71aa78fc7472de3664511c88f9d58d9a9c6922f26d67323869b5a661b106e0d0-kubectl-1.7.6-1.x86_64.rpm
../../pool/4af2eb4114017f12fb36b92a85c8149de6c54020a73061f3666d033bfde8f3e3-kubectl-1.7.7-1.x86_64.rpm
../../pool/761837737577fe362b16ff2b1fc34c383d86b0f9f746a9901c62587fd5f4e0f6-kubectl-1.7.8-1.x86_64.rpm
../../pool/2ebbb529212cb34b3d3a973fc60c6d4d1039b34ef2b11c79cbea2e1b6b6f5bfc-kubectl-1.7.9-0.x86_64.rpm
../../pool/ef4f06f44a3e9412735722d5625f814505054bd5f672ccad3e631470e5da9cd0-kubectl-1.7.10-0.x86_64.rpm
../../pool/7da91dc73456ad61f773fdd893a96b22b9537da395658b29c2cbda376f701f50-kubectl-1.7.11-0.x86_64.rpm
../../pool/4919fb3d1a5908461418f8a901c1c89bfd37e3418d72bbd4a8a215104015dae6-kubectl-1.7.14-0.x86_64.rpm
../../pool/5e509f4b46906b21ec01844d1566374b2699ea2e27f82c7902df6e94381c516e-kubectl-1.7.15-0.x86_64.rpm
../../pool/2d60f27c4698dcf7575809b91202f546caeb59bbe775ca202c5f69625a863a51-kubectl-1.7.16-0.x86_64.rpm
../../pool/582709dfb1c59f85c78f7db2c7609a1f3bf7308b73255bdbc910695b31b8660f-kubectl-1.8.0-0.x86_64.rpm
../../pool/0ee48e6b4033fdf520f5893759b0665090ffb83eefdbe3f0b41edf54f2247ee4-kubectl-1.8.1-0.x86_64.rpm
../../pool/3cc05eb1b1565779e8033743f1a2b9c8fb4e3b432421719ec56a3024d33dfccc-kubectl-1.8.2-0.x86_64.rpm
../../pool/ae43f92f96e828779f9744b3660e207199d97dda22eb44c054d2f3150da76b94-kubectl-1.8.3-0.x86_64.rpm
../../pool/a9db28728641ddbf7f025b8b496804d82a396d0ccb178fffd124623fb2f999ea-kubectl-1.8.4-0.x86_64.rpm
../../pool/e5c2fec310104577d9746c20b61d078ecf842d34785f615d8a89550a48acc1a8-kubectl-1.8.5-0.x86_64.rpm
../../pool/432a40bb7dbe35e3c85cd926b015269b8de54809ad0f14ceecfd5a05acbe44a4-kubectl-1.8.6-0.x86_64.rpm
../../pool/d8cbd7dd54e61c27d437cc32f2af5168ae693b888a7fa94f86583be94e9e8fbb-kubectl-1.8.7-0.x86_64.rpm
../../pool/bb633ec4b637d8dbfdba4b50f427dffd1347ae8bb44c49ffc75809d6d052eab9-kubectl-1.8.8-0.x86_64.rpm
../../pool/d8701448da60b3b00bc02930f5418ad19e8e9da398442e6d2b806e6d9e2cdc74-kubectl-1.8.9-0.x86_64.rpm
../../pool/83b1b027b99b38cb67713bd4c60e7fc3f08d2a60f53b28fe7e6ccc64a47bb1f2-kubectl-1.8.10-0.x86_64.rpm
../../pool/a5fdacbb2676408f4fa8dc779ef442d54f8af32878741c9d3e18d7d278ab241b-kubectl-1.8.11-0.x86_64.rpm
../../pool/3182954b1bf7afabe1d3ad39888800c54c80f4b6a03bbef1df84732e5c9f3a1f-kubectl-1.8.12-0.x86_64.rpm
../../pool/ced5e4a04e8fe35ec61962717fb996e5bc712c3c523451cf9d8882937ba22a73-kubectl-1.8.13-0.x86_64.rpm
../../pool/bc390a3d43256791bfb844696e7215fd7ad8a09f70a42667dab4997415a6ba75-kubectl-1.9.0-0.x86_64.rpm
../../pool/331270b6ba931a571640b0552d8737631cf30cc9cde14eaa45a5acd2afb8f304-kubectl-1.9.1-0.x86_64.rpm
../../pool/2dd849a46b6ca33d527c707136f6f66dfc12887640171333f3bb8fab9f95faac-kubectl-1.9.2-0.x86_64.rpm
../../pool/eb54809406a7bc73f1996ce669323b65f0846d1830ef04a1e69e66f633a627a7-kubectl-1.9.3-0.x86_64.rpm
../../pool/9566533de10dab5d80d8e7a8799ffea78e2b0f43e838e3b9d08f5e51c4dbc7ba-kubectl-1.9.4-0.x86_64.rpm
../../pool/b8fb4f730cd9b661657b8baf0119e291add74e753d50e313843cfee6e0e65577-kubectl-1.9.5-0.x86_64.rpm
../../pool/c9a30a9b3cd4f8b83a3ffcbfe5a23b32c0c78ec90a8e67505ba4ae31ed1d7a69-kubectl-1.9.6-0.x86_64.rpm
../../pool/051265818f6c9d68fa4d5b582c996facf3292b67256ce8519f658c058fe8a85c-kubectl-1.9.7-0.x86_64.rpm
../../pool/9ff2e28300e012e2b692e1d4445786f0bed0fd5c13ef650d61369097bfdd0519-kubectl-1.10.0-0.x86_64.rpm
../../pool/b8f0cc3bc85e8614f0340547b14ca5377410afd087f51077410945a594f0b71b-kubectl-1.10.1-0.x86_64.rpm
../../pool/32e8bd812a3944ccf07750d52088a118fa11493d34e009e2873317e0f0b0dfd2-kubectl-1.10.2-0.x86_64.rpm
../../pool/2e63c1f9436c6413a4ea0f45145b97c6dbf55e9bb2a251adc38db1292bbd6ed1-kubelet-1.5.4-0.x86_64.rpm
../../pool/454dc2629a55437b2297742eadad6429117afd3c48965e7ef6bd830cd71b4bf6-kubelet-1.5.4-1.x86_64.rpm
../../pool/af8567f1ba6f8dc1d43b60702d45c02aca88607b0e721d76897e70f6a6e53115-kubelet-1.6.0-0.x86_64.rpm
../../pool/756d35ea4c7e7fde7343154bf073f4056521c8bdf4d9562ce5c325e9bcf79bcf-kubelet-1.6.0-1.x86_64.rpm
../../pool/cde0b9092d421800f7207df677d6a1f321c82179e89a24e4b3c009a875e62c91-kubelet-1.6.1-0.x86_64.rpm
../../pool/0e7171c32ec6c65100eb284f8ebc3ed06bfa62a3e2ac752994da93d2261d8de9-kubelet-1.6.1-1.x86_64.rpm
../../pool/a8a2b37431da90798331a67b8b452572a72173414b1543368786e356f23bc4ce-kubelet-1.6.2-0.x86_64.rpm
../../pool/9e4c4e215f62ccd63ac6327b1d64faa9a66dbb24975f781a9c9c9a0a5c1e441b-kubelet-1.6.2-1.x86_64.rpm
../../pool/b610c239bd57c3ca0dcae3f4d4ae60a0a1eab99c7b059cf1b0fe8dd7267c3f73-kubelet-1.6.3-0.x86_64.rpm
../../pool/02f43f121f2fdff8217e62500bc2ef060650795e87c61000cd9a70842d668fb1-kubelet-1.6.3-1.x86_64.rpm
../../pool/08706c8f6fbce961cd8994043cd27204f2bd3b47afd54cbd8755dae4b6cc4279-kubelet-1.6.4-0.x86_64.rpm
../../pool/c90b57c8768e5d9992afcb5543bb20bdaeba5d0fcc581b34a00220fcf54fbb1e-kubelet-1.6.4-1.x86_64.rpm
../../pool/7d43767d519f9f76605408517c999631e3835621afa8f5e80b9b8fd0925842ca-kubelet-1.6.5-0.x86_64.rpm
../../pool/475f4d3ace5d98b41269c7734d27be722deb97862196439f9d764211463e0ccb-kubelet-1.6.5-1.x86_64.rpm
../../pool/f2c5f556143a820ed5bbb6a0ac7601f68dea8b28c8edc5db0a0d5d0ad4e94fd0-kubelet-1.6.6-0.x86_64.rpm
../../pool/2c3367bcc8239a3c681569a95f7837c66d6eabc6a534cd4bf9689a8fe8cb5243-kubelet-1.6.6-1.x86_64.rpm
../../pool/b58a3f13d494458fbe23dbf22abc0213dc2c9ffb1e30f76dad3d7321d0715444-kubelet-1.6.7-0.x86_64.rpm
../../pool/882d43aa9fb8414d5e1eef331f4f5bdbb73d5934938250f42d1f98a21506f5fd-kubelet-1.6.7-1.x86_64.rpm
../../pool/441742abb2a3725d39a38b418f7bb63fcc24181a55fd350f2bc47521df3a528e-kubelet-1.6.8-0.x86_64.rpm
../../pool/233c075a115b32eb8b53212a0b8e21414c944cc0d701d8c6ef334084105bc1be-kubelet-1.6.8-1.x86_64.rpm
../../pool/d6b67999f4591bc4b3eaec94446458568d49345338c281e429b32e7b41c2c9ae-kubelet-1.6.9-0.x86_64.rpm
../../pool/697f565212c13f7d0312b52a89b154ec66670385f4f0db1685704c75cdfd3836-kubelet-1.6.9-1.x86_64.rpm
../../pool/8695191ce64e8742766e44fd08beb61a640522db81cd1a628e2a578c8cc5df0a-kubelet-1.6.10-0.x86_64.rpm
../../pool/0554e280fb1cab3a572a0a4b22817dc01fb4d206d7022e39a29761903579abdd-kubelet-1.6.10-1.x86_64.rpm
../../pool/254343aac568684ff6bdf75e69c015127b45cd8a6e7f9b0eff335b0a082b3117-kubelet-1.6.11-0.x86_64.rpm
../../pool/31d0505434a4a9856f594daa6e478eab83c815f9d1b570d02ce25f83d9649f09-kubelet-1.6.11-1.x86_64.rpm
../../pool/f469e89265a9a215358f6d2e541c0d5dd4f6b522fcbbe401117ce656b3aba931-kubelet-1.6.12-0.x86_64.rpm
../../pool/3d7054b9bbb673a8dcfab836dc9fed4791a398199a7c08778cd025f4a1dd4df9-kubelet-1.6.12-1.x86_64.rpm
../../pool/01fad4462c4e6a1b27be8aec6a5ee730df26ea7ed9e9e965211eaf031e022df0-kubelet-1.6.13-0.x86_64.rpm
../../pool/14cef290532e68668924850f5f7002ffea058113917399265d8beb6648641cd9-kubelet-1.6.13-1.x86_64.rpm
../../pool/8f94c56214c25f72f999f8a9c099275c2e1b99110a155679c1eda293192c26a7-kubelet-1.7.0-0.x86_64.rpm
../../pool/6b748a98eaf71f33d3b37b911c8ba77650b25ff944a66743e703aa4bf1af0918-kubelet-1.7.0-1.x86_64.rpm
../../pool/30136924ea242fee92df3aaf0297365dc58e8d1260b7bde66c94bda5a42a6f04-kubelet-1.7.1-0.x86_64.rpm
../../pool/05be0e2d097c68b6355de205fe11064d6593593058d93fbe8031d7cc7fa74f5c-kubelet-1.7.1-1.x86_64.rpm
../../pool/1e508e26f2b02971a7ff5f034b48a6077d613e0b222e0ec973351117b4ff45ea-kubelet-1.7.2-0.x86_64.rpm
../../pool/3334e8cb551b23cbccd2e01d09c1647b83a247579b0e37c6bde2ce4fbd54edba-kubelet-1.7.2-1.x86_64.rpm
../../pool/28b76e6e1c2ec397a9b6111045316a0943da73dd5602ee8e53752cdca62409e6-kubelet-1.7.3-1.x86_64.rpm
../../pool/74106830515df731030ad7b3eac96e5cc44470dd62e8724fb91bb016f0dfc102-kubelet-1.7.3-2.x86_64.rpm
../../pool/4f60c17a926175fb9abcfdd487cebafbbbce0e2248d2b99c189ae0877376b88d-kubelet-1.7.4-0.x86_64.rpm
../../pool/737d15479c1dec0037df731db6ebbc7339c5205c24a9ce0fe669bccf18fe2c9b-kubelet-1.7.4-1.x86_64.rpm
../../pool/c87d884d28952332d952d5f7f294ead1838a8e1219d2fc75923a5be06903afaf-kubelet-1.7.5-0.x86_64.rpm
../../pool/e50c6e3313325838d5d1c27b48763cf2e5c53f1bade0502611c15655e86370e4-kubelet-1.7.5-1.x86_64.rpm
../../pool/f049d9a0a9172b00aa2725c86a0de6f4ee51571105344b31b0b2523be9fda635-kubelet-1.7.6-1.x86_64.rpm
../../pool/71b0d0d00f423a1cc7499d1e1544de54319653d1f2118bb96d35b3eb732e24e7-kubelet-1.7.6-2.x86_64.rpm
../../pool/45bd874fbf333c3e2595da358ce7ad3220af46b1ff69d69f57f882dee0db52a8-kubelet-1.7.7-1.x86_64.rpm
../../pool/04dc3c71a396b2c166fffc5dba87df260919f6905774e314e3405b11019f3dd3-kubelet-1.7.7-2.x86_64.rpm
../../pool/305fd1a89e9da295f56ac2df41dd73e521e29c768c57074a365096d713bfe928-kubelet-1.7.8-1.x86_64.rpm
../../pool/11bff1343306ea07d59022267f27b9008bd0101bd738bc49d6957ed7be68f025-kubelet-1.7.8-2.x86_64.rpm
../../pool/c5f3b22a90539fc5880728f80a4faabe92d0101ccd333ba864ea9898dce754d2-kubelet-1.7.9-0.x86_64.rpm
../../pool/68b9b7b4921a3b1287cc518b51b040f8c6c0586f2175c186b3e52d36402bbd83-kubelet-1.7.9-1.x86_64.rpm
../../pool/b7127de06c5bec3e197d279ac910b7a119fa2f799fb18ca3ec11499506396957-kubelet-1.7.10-0.x86_64.rpm
../../pool/a68572067e8ed51183782689c9e98da8d4753f86305c45a1fdce7d202b8b66e5-kubelet-1.7.10-1.x86_64.rpm
../../pool/57fc3fb190b0f538f1f6b109f0b23f3456bc48aae3e2ceac5041c6438aeb6a50-kubelet-1.7.11-0.x86_64.rpm
../../pool/02180e8b3266343d1f3ddebda362f6e199f1149528ab4b51479a5ff082de0739-kubelet-1.7.11-1.x86_64.rpm
../../pool/5c0e31c65e992e0f20cdb3b783d5f00293f0690cafa045a97f4af86cd861c39f-kubelet-1.7.14-0.x86_64.rpm
../../pool/8ebb830fda83dc4d4d0d8703818d997174c2be7b2f370da7aac768f0e46eaab6-kubelet-1.7.15-0.x86_64.rpm
../../pool/5310f9f1b7961b4e5b11a5ae3d692d0526e24f7e68bbacc286e5ce150fdb8a13-kubelet-1.7.16-0.x86_64.rpm
../../pool/d51b547da5980e2f6bfd33a382779194a0eacecac999ca726f85ef76183a9033-kubelet-1.8.0-0.x86_64.rpm
../../pool/9699684a155be5500822a37a01bc76e2ad5605b8541262365488b4530665a4e7-kubelet-1.8.0-1.x86_64.rpm
../../pool/a35571037b554243d386436ff729c90a3cb270f5797b7cd254ef0afbd4e706bf-kubelet-1.8.1-0.x86_64.rpm
../../pool/c8e037e025d56a6fba51fd30c9a4001d442732bb5b122a476aa9ab98c9ebbb68-kubelet-1.8.1-1.x86_64.rpm
../../pool/7d976686554e598267d8c5eb030f2a8f4e575f47015ecf94459913b80c9e5bb8-kubelet-1.8.2-0.x86_64.rpm
../../pool/5a7fa12561d34b95da6c1913086316a5f08169138a68a4d3874b76a64fd95114-kubelet-1.8.2-1.x86_64.rpm
../../pool/a53acfe63a475bf61661036c12890217f4921a6d6d6c3e6ecb4c598fc11cac19-kubelet-1.8.3-0.x86_64.rpm
../../pool/9de65bc04416de79a9c045f0191b1dda187151f9b1c665aa03b66f482092cde0-kubelet-1.8.3-1.x86_64.rpm
../../pool/1acca81eb5cf99453f30466876ff03146112b7f12c625cb48f12508684e02665-kubelet-1.8.4-0.x86_64.rpm
../../pool/d353208063910d0d9a8ff9a3e9bfb0cbee43be501bc4ca16d69223560c4a5cc0-kubelet-1.8.4-1.x86_64.rpm
../../pool/7330c09dd7dccc300c9c5f696fc4a8a76327e5068e51a60e4517b59a1c4b3dbc-kubelet-1.8.5-0.x86_64.rpm
../../pool/26944be0d909a90d8ef7559b17a3ce4381821d623314a35eb0c74e62d055d944-kubelet-1.8.5-1.x86_64.rpm
../../pool/f4011419193577161ae891a1b26986cae5f5e588941340db3abb771e2a677de7-kubelet-1.8.6-0.x86_64.rpm
../../pool/e2792eaf4f3088dd639feea4b4bbfb0b2f96ac498d7a67c31b47488758ff7633-kubelet-1.8.7-0.x86_64.rpm
../../pool/0e9672058439598e1d6d57ae5bd36007730225f94139cc0f03ddc81ddd798787-kubelet-1.8.8-0.x86_64.rpm
../../pool/6330c00f0fcefa1ad4e5541abe6ff2682668ec1c34b9116a26c54191203c4d24-kubelet-1.8.9-0.x86_64.rpm
../../pool/bb3cbbc83f073688b4d4b005f1608cbf4b2d81d209e9a7bd16e825678a49d048-kubelet-1.8.10-0.x86_64.rpm
../../pool/16c56093ac1b63f9e03ddd31c763ea6779754799add991e1468db9e27f04be83-kubelet-1.8.11-0.x86_64.rpm
../../pool/9fbc91e22fb4c35e8d1f7cdc2ee10938153cc2cba66d647c9448ff373494b213-kubelet-1.8.12-0.x86_64.rpm
../../pool/cc2d0955725febc15f54f2696197ac0d0d0a42af6bea61266f9da0327bbd22e8-kubelet-1.8.13-0.x86_64.rpm
../../pool/8f507de9e1cc26e5b0043e334e26d62001c171d8e54d839128e9bade25ecda95-kubelet-1.9.0-0.x86_64.rpm
../../pool/cec192f6a1a3a90321f0458d336dd56ccbe78f2a47b33bfd6e8fd78151fa3326-kubelet-1.9.1-0.x86_64.rpm
../../pool/0e1c33997496242d47dd85fffa18e364ea000dc423afdb65bb91f8a53ff98a6f-kubelet-1.9.2-0.x86_64.rpm
../../pool/b59d38d81913c58b9b382e4c39a00303d68ece464f5402046f0900e70ea106f8-kubelet-1.9.3-0.x86_64.rpm
../../pool/d69020c41cbd116872424bef4b6d6c5c7b29cbcc8fca3ca949aca421899c27fd-kubelet-1.9.4-0.x86_64.rpm
../../pool/aea4066130f6381c7145630edddd65df6293728b2fef3370e040920367f8e33d-kubelet-1.9.5-0.x86_64.rpm
../../pool/fff4e7133c41ce6aaca95adb598f47967d180c4c5e8e6c67d8bfe58345bfde27-kubelet-1.9.6-0.x86_64.rpm
../../pool/4a7b197df9b4f478f6e26863e450811c89c28126993eb0cc777fb39b819cedc9-kubelet-1.9.7-0.x86_64.rpm
../../pool/5844c6be68e95a741f1ad403e5d4f6962044be060bc6df9614c2547fdbf91ae5-kubelet-1.10.0-0.x86_64.rpm
../../pool/4cad7573e2617d903a12f5d318f597ec47bc0e67792a04f67ffee3e06c0ad373-kubelet-1.10.1-0.x86_64.rpm
../../pool/bdee083331998c4631bf6653454c584fb796944fe97271906acbaacbf340e1d5-kubelet-1.10.2-0.x86_64.rpm
../../pool/0c923b191423caccc65c550ef07ce61b97f991ad54785dab70dc07a5763f4222-kubernetes-cni-0.3.0.1-0.07a8a2.x86_64.rpm
../../pool/e7a4403227dd24036f3b0615663a371c4e07a95be5fee53505e647fd8ae58aa6-kubernetes-cni-0.5.1-0.x86_64.rpm
../../pool/79f9ba89dbe7000e7dfeda9b119f711bb626fe2c2d56abeb35141142cda00342-kubernetes-cni-0.5.1-1.x86_64.rpm
../../pool/fe33057ffe95bfae65e2f269e1b05e99308853176e24a4d027bc082b471a07c0-kubernetes-cni-0.6.0-0.x86_64.rpm
../../pool/fd5f5da2d1a262fa687404d34e72813520274364557e648bc64a8136f1a02630-rkt-1.25.0-1.x86_64.rpm
../../pool/7a382e59dc2c39a66083e03ec061f33771e4a7130c98cd0ef61492b2196c0378-rkt-1.26.0-1.x86_64.rpm
../../pool/01b97b0ddb967d0ed9fd78327a784efbfea8cd0d9789f5bab8b9bbfe94477c60-rkt-1.27.0-1.x86_64.rpm
```

The example is earlier content of less and equal than V1.8.4
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

Then modify scripts for latest packages to download
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ ./yum-pool-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 16.9M  100 16.9M    0     0   557k      0  0:00:31  0:00:31 --:--:--  510k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 7759k  100 7759k    0     0   460k      0  0:00:16  0:00:16 --:--:--  661k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 17.3M  100 17.3M    0     0   503k      0  0:00:35  0:00:35 --:--:--  631k
```

For example, the latest is for V1.10.2
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ find https0x3A0x2F0x2Fpackages.cloud.google.com/yum/pool -regex ".*1\.10\.2.*"
https0x3A0x2F0x2Fpackages.cloud.google.com/yum/pool/32e8bd812a3944ccf07750d52088a118fa11493d34e009e2873317e0f0b0dfd2-kubectl-1.10.2-0.x86_64.rpm
https0x3A0x2F0x2Fpackages.cloud.google.com/yum/pool/b754a6990af7d7012189610b0dc69e6e950c13a8c415b9ebea8d56352e9719fd-kubeadm-1.10.2-0.x86_64.rpm
https0x3A0x2F0x2Fpackages.cloud.google.com/yum/pool/bdee083331998c4631bf6653454c584fb796944fe97271906acbaacbf340e1d5-kubelet-1.10.2-0.x86_64.rpm
```

Or
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ ls https0x3A0x2F0x2Fpackages.cloud.google.com/yum/pool/*1.10.2*
https0x3A0x2F0x2Fpackages.cloud.google.com/yum/pool/32e8bd812a3944ccf07750d52088a118fa11493d34e009e2873317e0f0b0dfd2-kubectl-1.10.2-0.x86_64.rpm
https0x3A0x2F0x2Fpackages.cloud.google.com/yum/pool/b754a6990af7d7012189610b0dc69e6e950c13a8c415b9ebea8d56352e9719fd-kubeadm-1.10.2-0.x86_64.rpm
https0x3A0x2F0x2Fpackages.cloud.google.com/yum/pool/bdee083331998c4631bf6653454c584fb796944fe97271906acbaacbf340e1d5-kubelet-1.10.2-0.x86_64.rpm
```

Just download CNI
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

just download V1.8.2
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

### Repository mirror configuration

#### Fedora 23

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

### GPG

Download
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

### Dists

Currently V1.10.2 is released
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ ./apt-dists-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  8993  100  8993    0     0   3210      0  0:00:02  0:00:02 --:--:--  3209
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  8490  100  8490    0     0   6778      0  0:00:01  0:00:01 --:--:--  6781
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   454  100   454    0     0    579      0 --:--:-- --:--:-- --:--:--   579
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  127k  100  127k    0     0   101k      0  0:00:01  0:00:01 --:--:--  101k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 16113  100 16113    0     0   9626      0  0:00:01  0:00:01 --:--:--  9625
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   228  100   228    0     0    281      0 --:--:-- --:--:-- --:--:--   281
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    23  100    23    0     0     28      0 --:--:-- --:--:-- --:--:--    28
```

### Pool

List packages
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ egrep '^Filename: pool/\S+deb' https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages | awk '{print $2}'
pool/docker-engine_1.11.2-0~xenial_amd64_5554a8bc383e65fb10d556239c72b26457cc5d97f49e3a353c3382f8434e7d21.deb
pool/kubeadm_1.5.7-00_amd64_2759fc99e5b23e44c92b44c506ed9cc1c2087780786bfa97c715da02da84c55d.deb
pool/kubeadm_1.6.1-00_amd64_46af294c4bcb0f923b5593c37a01c27393ba89e1219b10b81f0e6e029c8c1754.deb
pool/kubeadm_1.6.2-00_amd64_5916370a48f2a60174fce83443179f6a0caf79212420c1b5cce73c30f5b79c50.deb
pool/kubeadm_1.6.3-00_amd64_ff5e882c88a5db71803aab900c0b341eb63038558da73c51c3d351070b0c62af.deb
pool/kubeadm_1.6.4-00_amd64_2cc86f287c52e2f63eb9341d7141c593090efcb5164a5ad6004b763e62be51bd.deb
pool/kubeadm_1.6.5-00_amd64_dcdd54c3b0dd4dded1e8fa3e8e04b6427e97a0bd8495b2f7712f1da2a8210872.deb
pool/kubeadm_1.6.6-00_amd64_60fc9618a55876c066f2ccf9adfa95e8baf40499519a8ac5da82640027bcbe51.deb
pool/kubeadm_1.6.7-00_amd64_ece025894ced2577ed0b3bdafa1eb7f615fd35386dbe4e7716fcf6bfdbedbf02.deb
pool/kubeadm_1.6.8-00_amd64_6905ef6a397bf00173583b4005509626832741fc03874ed45cb6370fe624f796.deb
pool/kubeadm_1.6.9-00_amd64_8e4eb0029f66fa86548e63a7acf5b29b8d1f4009d8ba99d567a0a3a6db45a0c1.deb
pool/kubeadm_1.6.10-00_amd64_e7abe0f9e0fdae3061ffc9543382b3a373dc937afb0b08a8febd0ca0fae1bf86.deb
pool/kubeadm_1.6.11-01_amd64_7a00a0c4e91715b1a6bc0439923069efae4f128779e7b2e93c6479db7cc1f385.deb
pool/kubeadm_1.6.12-00_amd64_da9679c05a321ea98ea38c44795233713db86fb612c1d675e3b358ae5f1984dd.deb
pool/kubeadm_1.6.13-00_amd64_5c6da07212819f6c1b62272a18a82101b9e5a28ded130043c7ae0b275c84b47e.deb
pool/kubeadm_1.7.0-00_amd64_6a9f69bdbc29c9f5c26729329225b9199e6927529516de2437ab8df1b92f9919.deb
pool/kubeadm_1.7.1-00_amd64_e6241f9203632e76335fd17ddf7ee8bd403792fdf0e5c97834b43ec3a7eb4b68.deb
pool/kubeadm_1.7.2-00_amd64_47f0e01203a2777cc49ebbf389959c44886c071efd26c1b80c40b0321c640d53.deb
pool/kubeadm_1.7.3-01_amd64_c4bac3fcbc1a820523a3617495707aff0adab0db52ae02f3a5ee5001ff0a1e74.deb
pool/kubeadm_1.7.4-00_amd64_4a999131c64664ccd0ed1792cf517974e03ef531ded8c624cfef58a20321cd0e.deb
pool/kubeadm_1.7.5-00_amd64_5683ceb924f3cefd14a3307ddef30af20d567ad90692ef6a92f229ac9fa1b3a0.deb
pool/kubeadm_1.7.6-00_amd64_68f5bae22936d6c609364154bdd4ffffb91ffb502da6cf54647de04026bf142c.deb
pool/kubeadm_1.7.7-00_amd64_b8317e83d75cff7d4a01566b17fe7b24440c07526108ad8794b0d07e5b55d3a7.deb
pool/kubeadm_1.7.8-00_amd64_8e826d3ef835ca1c1abd1f4be11ef19f7b0ad5d6b85d52e9b26b6ed19acf7d95.deb
pool/kubeadm_1.7.9-00_amd64_f92ec1a030fd840fd17bbfa8b052cff945b92962bef8d61fee51a10cae525067.deb
pool/kubeadm_1.7.10-00_amd64_987b54cca425ba69b0b799f0a232a6252211fefdf768a45f2bac16e83a81c8a6.deb
pool/kubeadm_1.7.11-00_amd64_eb2d595ce17ad92341d7e08742c4bbd72c7ebb3a20bd35ac8bbf5f1855da6e17.deb
pool/kubeadm_1.7.14-00_amd64_e7bfddf64bea954e1ce8115f0f8803b4b95e1d79fef147780f4e19978f952f4c.deb
pool/kubeadm_1.7.15-00_amd64_c6f376c65346540f6806c7288bd772745744a6b472dee080c8b9cbedef7a4957.deb
pool/kubeadm_1.7.16-00_amd64_f2d39cff3636ccea71e838df9fd675fd8e96ee7f00f66b56739dc5b368960d7b.deb
pool/kubeadm_1.8.0-00_amd64_22842ddc6d1ffabc04718f384ac001ffa56324cc61e6c3a7c991337bf3e39e06.deb
pool/kubeadm_1.8.0-01_amd64_15e6b68827964187a6c3a60711fab285d06d19953f7473154d798ff52a418185.deb
pool/kubeadm_1.8.1-01_amd64_5863420bee513756daf8cd08af801c5420397409191df19535d8c618a100989c.deb
pool/kubeadm_1.8.2-00_amd64_e743a9538b855d08ddaa68e7910af2dfc2bb9c1a0938d79089b0a9d3f1c19dde.deb
pool/kubeadm_1.8.3-00_amd64_a2cf7650f63eebb1b95887e2989c19d7f61b24f1da379d705bcb3e17123bbd75.deb
pool/kubeadm_1.8.4-00_amd64_0088836fbb451bc49ece82f34c035f50f2e1dd4dea78f6d585574d585da11e8e.deb
pool/kubeadm_1.8.5-00_amd64_e6138f65d23ae9f18d3ecbb69d27cc40c30feff200d1652ce7814e365848ad88.deb
pool/kubeadm_1.8.6-00_amd64_66ce5093b3e309db9ae0d79be4bdf370251ae7d705a8a361130d82a7eac155c4.deb
pool/kubeadm_1.8.7-00_amd64_4a0952ba784b8bdb10c488a76fc89c4878ab86cb1f858ea38da52bc7909d6b0c.deb
pool/kubeadm_1.8.8-00_amd64_f9d534729452a1e97820390c845e67c00b953e0cea806f0d8ac1d494ffecb45c.deb
pool/kubeadm_1.8.9-00_amd64_170dac07edd6fbb7888e55f8f84e3135eadcb326ee41a7ffb91411fd62a917c4.deb
pool/kubeadm_1.8.10-00_amd64_bddec18c35b3f0d004fc6951ed3bf177b4719f5f9d3e7138a5bd9b31a294d26a.deb
pool/kubeadm_1.8.11-00_amd64_c85712972eea993f59edd68cead2017451165d279cec3bc0d2c8b62ae6f8a62c.deb
pool/kubeadm_1.8.12-00_amd64_bac0585cbd950a79cc06dc1dc5e0d74fdf386efbfeaad1bb3a7c9aea4e119e9d.deb
pool/kubeadm_1.8.13-00_amd64_1e80966640bc0b8a769ef708d6af032b698e98eb2888859519d998f8478ebbdf.deb
pool/kubeadm_1.9.0-00_amd64_191bd1d8a63d263cdb318f77b03fad6c8387a912cf16a21b9b24d7e9108b4911.deb
pool/kubeadm_1.9.1-00_amd64_31ec5709ec1eb37b194563cb28eaafafd1f0c9d009239cdfb694686066b35c2e.deb
pool/kubeadm_1.9.2-00_amd64_869c826d406035aad1617118f88d5c627e16797388a5e2b3735d7b825d62e7ad.deb
pool/kubeadm_1.9.3-00_amd64_df3d3c15601640c5cd5db3902d52fb2e33263b79c62d299700971083b7bec9f5.deb
pool/kubeadm_1.9.4-00_amd64_7c267288a1a9cebc480778b50ef00ab5dfe477f545bc917a50adb322ea2b9c24.deb
pool/kubeadm_1.9.5-00_amd64_3d2a1245ecfd2ef57d1d6f0f3897658d5e506ed467d8197f0669be445e9db259.deb
pool/kubeadm_1.9.6-00_amd64_9763ef588d8820e343c180f1e2c16d1c51be4d6a686f2d5ebd87a64d8b4fd661.deb
pool/kubeadm_1.9.7-00_amd64_af2c1e96e8339736360470580f6bb642f5b8702bf969c5acdeeb23626198d11d.deb
pool/kubeadm_1.10.0-00_amd64_ea6b408af5de27ae3df9a1f96724067104998608b64869fee2c680ea9f9c000d.deb
pool/kubeadm_1.10.1-00_amd64_d5942f2d1a7b85136b82dc79c6f6204373c731b6ef7fe204a96f1c56af4859f3.deb
pool/kubeadm_1.10.2-00_amd64_4bdf79bcda2a820210f5c0f7e84f9f57d6cc196ba9b59943f296544624e6d743.deb
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
pool/kubectl_1.6.8-00_amd64_0aa203bad55d821a0729009bb1436618b70eafe4549f3178f02b1cd11b2fca83.deb
pool/kubectl_1.6.9-00_amd64_c6ea792bb752621b2af1c00f5bb56686df831bcd9b2ef4a5e090383fc4004656.deb
pool/kubectl_1.6.10-00_amd64_2ed0f68b0a1e8257fd363bb35de28b29be3ec086dacea066790fa00fdf4a0e74.deb
pool/kubectl_1.6.11-00_amd64_0049f491da3c6d5826d5cc0862a84da361daf3c356fddb507746891f456e4bac.deb
pool/kubectl_1.6.12-00_amd64_54410fc4e429df54df34b406b2515cdb3043b11f14ab87ab522f54a6a4e10e92.deb
pool/kubectl_1.6.13-00_amd64_a2c7e0339dccaecf28740a2e8ac3f953c45f770db4a7f68003f99fcd7fdceb42.deb
pool/kubectl_1.7.0-00_amd64_6bd51a1ed15eb74ad2d72823a301c71e769dcc12ec10294eac3026cb77e03860.deb
pool/kubectl_1.7.1-00_amd64_7f6d0471be6524d0d0aec8e9f515e0261a8e364d4fd2f25f22cba4895634bab0.deb
pool/kubectl_1.7.2-00_amd64_cd4dec0c4d75f1553712543c03720c3fede37ff8a46b97f287d7f4bea58f36c6.deb
pool/kubectl_1.7.3-01_amd64_84173e3f104760b832ddf732bf02f1834cbf719032746518a52efcc942dba4c6.deb
pool/kubectl_1.7.4-00_amd64_6ee4015eabaa1e0cd4e2895c102c1c341f1c0037b28cb7700424d0f6f04a2416.deb
pool/kubectl_1.7.5-00_amd64_363445c3be161e924d1209c60e24dd2d0657d68bcf7c17e8ec2a6dc603c5d363.deb
pool/kubectl_1.7.6-00_amd64_a8c59bb8105f119061ed368451923465351488efa537a6cf9f88e3ce0f0422d1.deb
pool/kubectl_1.7.7-00_amd64_22e3fbc7a55ca7973eac5de600896827a15de0f750c6b7e37e46a8cb99a971fb.deb
pool/kubectl_1.7.8-00_amd64_0b904de238294283cbfdfd317664f552766dd974947fa5e0c2f3bd0077331522.deb
pool/kubectl_1.7.9-00_amd64_101ff6eb079ebee128ed10f3ff1dcef4d21b094e6d81150c649064ef1454b405.deb
pool/kubectl_1.7.10-00_amd64_b725d080205ddf237a4df726dba6c81c9f56d1bc089a530e9f2df0bcd659af3e.deb
pool/kubectl_1.7.11-00_amd64_ef865d283165d40f85ee77c0ced027fc416f26f8f646d489b5265f33b5a85fd4.deb
pool/kubectl_1.7.14-00_amd64_dd862ac20277caa52836dfc306dc8119e3141776af1b905afd9e36a81fed46c0.deb
pool/kubectl_1.7.15-00_amd64_29a6ff4fad741a2516e4255402230c3393b98dde6088ebadc74b974e23b142b5.deb
pool/kubectl_1.8.0-00_amd64_9bd409d2a0300d4b550cd2a7fd2eca6437ea0767805f6fd807912b245ec949ed.deb
pool/kubectl_1.8.1-00_amd64_c87f3e1e54ad045af07001028282f9b76e1c26425f6ccaeba52bb0a5fa93a6b9.deb
pool/kubectl_1.8.2-00_amd64_b01f6fa567e98752181a3ad057275851e00aa5fa1a8db6eb5a6d81c0f499e1ec.deb
pool/kubectl_1.8.3-00_amd64_4ea956fb6586472bfec46f099ad34fa83585969bcfaf2ff995dd66570f824a8b.deb
pool/kubectl_1.8.4-00_amd64_b48511a481ddcfbf935ad935bc6c3992c1c4315fcd8f3f794e367b9b26b775be.deb
pool/kubectl_1.8.5-00_amd64_b9f161f02571cb379cfc93ed5a95e84b68b259cbc92d2f8291a5be3a6970ae37.deb
pool/kubectl_1.8.6-00_amd64_bc0fe25db6363a762821ec2460be64c6fe66e857210eafc45289f25ea92b1408.deb
pool/kubectl_1.8.7-00_amd64_7c6fa8b9de95ec7585b6af78997fa427c19e224a66b0da08cc4e76954c0bf700.deb
pool/kubectl_1.8.8-00_amd64_1a766866712f71c62ad30e6a247d13fab7eefe6eb3335a683b7697733bf36652.deb
pool/kubectl_1.8.9-00_amd64_5f045afdadb2b6253721343febd5cddf7bb98f40935439e427f9a2df0e96cf21.deb
pool/kubectl_1.8.10-00_amd64_4ece6cd7ef9bff0e1d2b6bb9c30b40a500133e9b1cb4dae778edca1d4a5bec44.deb
pool/kubectl_1.8.11-00_amd64_86ab3ffd7dece91a755136b67d40eb4d9ea2ec1404d91377d079833c31d647e1.deb
pool/kubectl_1.8.12-00_amd64_894345a819f1c6a86d411be6240f63d5c9ea69eea2505fae05e642272a9c6a19.deb
pool/kubectl_1.8.13-00_amd64_828ddd59b708e767ee49b866d231af2932412a64f5be69ed8602e2ae341c8a82.deb
pool/kubectl_1.9.0-00_amd64_8ea712c18d89d56090c8753a9630d22fd8ae5cb03d4ec79a1fe6d262c8b4eb36.deb
pool/kubectl_1.9.1-00_amd64_bdfb1ad90e0f02a7ae614502079a87ed99209679bdedf0c62873564c186c9f99.deb
pool/kubectl_1.9.2-00_amd64_a9d5a970c989cb1f77c30c5d27f611cab58659751005e7052541540dca19855c.deb
pool/kubectl_1.9.3-00_amd64_bf9b3256304edcd2f4bd85b74857dacfe6796bec548e1411ed6f93350dd01316.deb
pool/kubectl_1.9.4-00_amd64_0f814c1c7d852379aa61744ed119518558a0f0cfa1a6e36e8fc2037bd023dad4.deb
pool/kubectl_1.9.5-00_amd64_fc4313ac6797482c4ece18c14ba2ad2c46bb6bed4b4e9e5c1ee30f64c31ff13d.deb
pool/kubectl_1.9.6-00_amd64_8f3805e25732ca6ecc12e2751932130a5131adb605691b85e06c6204a0b6326f.deb
pool/kubectl_1.9.7-00_amd64_c532c2abf3344a081d16fc69253c470f1e22c39350b6778dec2f73bb88eaa941.deb
pool/kubectl_1.10.0-00_amd64_e391c19fa377b84587676c5577222ceb5d8fcbde442c79a9cd55d1f344293834.deb
pool/kubectl_1.10.1-00_amd64_a0d873581d1de51178e88938df4fc968fab57d8e0ecf3b9c211829e673f98ace.deb
pool/kubectl_1.10.2-00_amd64_a3ca21b1875a20567024584f18ec6aca390e8be8a8d92e20b5b7258452e419c9.deb
pool/kubelet_1.5.1-00_amd64_bb82dd4bcf0c9bc813c599f62afa48832bf34302d723c5a38347c2754f3735e2.deb
pool/kubelet_1.5.1-01_amd64_e61a09a7906a5c4e3293d73dd6890fc20a80338e2850445bd147c9dba97a815f.deb
pool/kubelet_1.5.2-00_amd64_ac51b2e458539f4803b2d2b31b3fad0b6d503ca8ddc1994e55b4150f33909163.deb
pool/kubelet_1.5.2-01_amd64_d39273bf2d91b194035d9c924c865393a9271c35a6db5b106c9cf78699ed7573.deb
pool/kubelet_1.5.3-00_amd64_6df7dd70c4db4afb77bdf309569d5e9ad92e7319ab4d3b9bcd0bd56f457a8755.deb
pool/kubelet_1.5.3-01_amd64_6e218165b6d7a024866fa430bce5940a2247ca469056c00b6ff7310e36d044c4.deb
pool/kubelet_1.5.6-00_amd64_a9a67135a020676e3879cca1971447f78bd4b365a243b94a42123cac54d03d9b.deb
pool/kubelet_1.5.6-01_amd64_1d17f35764899052ae936a509327a60eb753f06befdaef9345f0f2cca966decb.deb
pool/kubelet_1.5.7-00_amd64_f872da8a98d992681d49218c24703762fc1a9699f9a6af46de24e9401d6fb0aa.deb
pool/kubelet_1.5.7-01_amd64_b4fdcd12d31717c829362d7ee07c850ec91cb7981afdce68c362f87ea25a8405.deb
pool/kubelet_1.6.0-00_amd64_850ba773992934d5eb46b064274d3449bb5c07e2f1a1019b10e068e0a0efc78d.deb
pool/kubelet_1.6.0-01_amd64_9bb4f26a5ea2154b8b75cb2c4e8cfa69d25b276b1727f5004fa96a92596a59b6.deb
pool/kubelet_1.6.1-00_amd64_8f4433bbf40b5c2fdf75089f520544969cca62ee53b7ea87eb4f1123fa2f4163.deb
pool/kubelet_1.6.1-01_amd64_d1d4c6c5b3686288e5b8fdad38eaf3ef9311f2e441bc9da5f2595f1d11cd36e1.deb
pool/kubelet_1.6.2-00_amd64_8537952290ccb565bb9e6d7beaeeaa2b037eb5966df097d5ca95a7cb811c51d5.deb
pool/kubelet_1.6.2-01_amd64_c75d1faa0625ba7da33942809a796af6edda0eda7ab21402b6a5435f86f4b2e9.deb
pool/kubelet_1.6.3-00_amd64_a94b0cfa5b26939d87097dfd45260474c1effcad879e35940eb6d36e7d953d6c.deb
pool/kubelet_1.6.3-01_amd64_2a99513cfb75e929ffbf11e20b9311f8a298f130716f26339ecce517d42b3297.deb
pool/kubelet_1.6.4-00_amd64_2affcdcf43b6cc5a89237d0ceb1b5a3a90c605d4da73b97c5842f1a53368231d.deb
pool/kubelet_1.6.4-01_amd64_1315a3eca22fe54f50840a172bd5cc481876ab47cb53f216705bfd83b5ac3654.deb
pool/kubelet_1.6.5-00_amd64_2d1e2ea0fdd3f85594ed9d57b4bc19e3652f9bd66ad9099f0abfc268e6d9c379.deb
pool/kubelet_1.6.5-01_amd64_5c8209689a75443ac7cd918c9353ea14cd01f4e90a83ff6ac74a83ad7fc38d9f.deb
pool/kubelet_1.6.6-00_amd64_7e64ef3101aff3c7e59b7e38c3ff4acf7659730aeda47fd43bb7a12384069f11.deb
pool/kubelet_1.6.6-01_amd64_a16dc33161ee16dcfaabb0658a97fc3f49372db66af2bfe00f2249cef375fb47.deb
pool/kubelet_1.6.7-00_amd64_b3c7f8b1743cbb05af4fa62f159952572a4520a3289a742a0544fcd8068703e8.deb
pool/kubelet_1.6.7-01_amd64_307b05dfb47ccdcd50a6fdac9623e3f388f557b87424ca36d1fd6d41ee2b4ca1.deb
pool/kubelet_1.6.8-00_amd64_7bbb4de452b304799d6550d13e1065c439e798ccc9274d7e7ee2b30926b4dcd7.deb
pool/kubelet_1.6.8-01_amd64_8491be969b337db5213095b7e1d64e25ffb94cd834198936cd8f17692ae35405.deb
pool/kubelet_1.6.9-00_amd64_7ceae6daffdfbb50b2333063b8f06c5c8e704c82f4fcbf273f4d34a729ef82ef.deb
pool/kubelet_1.6.9-01_amd64_226d44d8f970337d3f07205583a00ddc05fed705a936d4a623186c20fe7e6604.deb
pool/kubelet_1.6.10-00_amd64_b7ff1d3504ad670137f83f081de7d41768e9204c8f6ddd69c55e0032d8011e2a.deb
pool/kubelet_1.6.10-01_amd64_79b4282c2e51af49fd6fb72eed00d44ccfc99b87a7f52dea9bf3c70afdef6b05.deb
pool/kubelet_1.6.11-00_amd64_2dd18ba275fd1bf41acd6280975fd7a45467ca2f2842e61703c2787856f4e2a7.deb
pool/kubelet_1.6.11-01_amd64_b61d8439af0c31b8a0e10e18e2f4f161f1f651ab4ab7407077fca91d670981a9.deb
pool/kubelet_1.6.12-00_amd64_c302b5075f37fc8bdc251b2170b1b7b358c0faaf993ee1f0cf20204e971009ef.deb
pool/kubelet_1.6.12-01_amd64_f61d249652954ea98d139332c81d1fb99db2b56d40265bc4aa88b7e8a1210740.deb
pool/kubelet_1.6.13-00_amd64_7355ee8511bff4cb8103e18f2e6f86bfa691c54bd6556da7950c18c8de6b4ca9.deb
pool/kubelet_1.6.13-01_amd64_71548a4d6d62a90932f286f3235b758b9f42572dec6fe859e405af394794ec11.deb
pool/kubelet_1.7.0-00_amd64_b7c9be22a6379d8cc47ba1183bbc9de0acf49ad1418ffa8ae57628ad1d7c051f.deb
pool/kubelet_1.7.0-01_amd64_12c44bfb8e7be653a22b2c2b9a988f7643bb0df9e00daa13cd54099a54453d6c.deb
pool/kubelet_1.7.1-00_amd64_8e6aeef38ca042b85765751652f8115960892ca4cbd1a906b63197931351cb8d.deb
pool/kubelet_1.7.1-01_amd64_100fffc507151ead00e2c44cb2a116be934080fe6deac29e72cf0bf3712f6bc8.deb
pool/kubelet_1.7.2-00_amd64_d685940c01bd21e18caa8528cf793b09f9d81f9c431173356461a8ecc0704f28.deb
pool/kubelet_1.7.2-01_amd64_44cf5b68e3e7f013e080e18138dd0fe3cae7fb3046dc461840e8ef21b82c91b3.deb
pool/kubelet_1.7.3-01_amd64_73541f99b3465ea2388e3f48d2207862a41c6e980782e82678d3a1bcab39630b.deb
pool/kubelet_1.7.3-02_amd64_5a07c5b0d7bb64d0ac96a48a3d30d36a154744ff4e393b1ffbf98ffc0e7be247.deb
pool/kubelet_1.7.4-00_amd64_8e9b2a6d7b233240f816591b8180d9e92b08a73afbe93ed0c783537831bbb4c6.deb
pool/kubelet_1.7.4-01_amd64_e82906fddbd5d5ccfd58c261610a6cc88f6423e5a9f908f5df77f2e5833304fa.deb
pool/kubelet_1.7.5-00_amd64_0db3259738d749acca6ffff03ae4ccb87a4e6a64c3ac9884b34df7da936a3052.deb
pool/kubelet_1.7.5-01_amd64_a8e6374be2df5c4f1abf69de707e8d40ac0220d2afd078ecae52942f51ddcd14.deb
pool/kubelet_1.7.6-00_amd64_65531a88ce1ba75817f0d6d32e10efcb1ea5aabaf70e9e9aa53f83180ca41834.deb
pool/kubelet_1.7.6-01_amd64_73fdc5da1321ad5f03116b2ff95bdd2fb3e86e21d08fbd3cc4b522ed0e28079e.deb
pool/kubelet_1.7.7-00_amd64_c66f68c9a638cc3acae96c86d6207f91a4e0efb27f43fadb02a78f72c71cc894.deb
pool/kubelet_1.7.7-01_amd64_ed7701ff112d36936da2f20c596dbb7c23e13b9df21f6bfb0d51d355b579cad2.deb
pool/kubelet_1.7.8-00_amd64_429a02bce3df7d943bbec28b2b35a9c9e1ba78a01484bfe1e4e7471a362c38bd.deb
pool/kubelet_1.7.8-01_amd64_8fff4fce7da3e0240e15818dafbd6f1de10c8412523ab5127e4d6997733c8491.deb
pool/kubelet_1.7.9-00_amd64_d4ba8476a776e3200c9376d352eb6abce0d573ecb6dd0961ea97d9db927b5dac.deb
pool/kubelet_1.7.9-01_amd64_1d7a73cb3de7907e8bc784325a08bbc7ddc3f7cf70f12a441f817f604a79f0ff.deb
pool/kubelet_1.7.10-00_amd64_4322e4474136b891848d5d67658b9cd22a41a834c88ecea382d58da092156556.deb
pool/kubelet_1.7.10-01_amd64_fe1de868c2cb929962e03802f96f9066c463271caa2d6bd991bbe96585d46b2b.deb
pool/kubelet_1.7.11-00_amd64_e21f114f3eee04b13b64218805ae6e594171c5a8d3d28a35d88bc926ddf9d094.deb
pool/kubelet_1.7.11-01_amd64_b0b19ff353777b3bf8b08ebffa9e2f2a140e9575a77d84e5621e27bae7a3f9be.deb
pool/kubelet_1.7.14-00_amd64_d197387f749b3807e0b6498e4783303936b92087e286851fb86465c18b7506ce.deb
pool/kubelet_1.7.15-00_amd64_d92cc5ccec7b452c457c47c08cc6df68efab96b9b2f1406627d3a48e2f35b65e.deb
pool/kubelet_1.7.16-00_amd64_d73c96d4f76b0b2816580e414ed8db4ea753e5e969edcf4a2424bc92f83bce10.deb
pool/kubelet_1.8.0-00_amd64_00b7c77c924d654c7def52c83cfeb9a3d1836c1e7b40683c3fe8207b0bd299d1.deb
pool/kubelet_1.8.0-01_amd64_c85cfd1d6ea9f1d2f5141692bdd828f49a71bf48b64ad4e667292800693cf46f.deb
pool/kubelet_1.8.1-00_amd64_c0f1902d204f6741eff08560120bfe193b992596068158a2042988203132cd51.deb
pool/kubelet_1.8.1-01_amd64_890853010cc549b3909a5a6c271bea87eccfd70c48747e2cdcc14dba975dec7b.deb
pool/kubelet_1.8.2-00_amd64_9646b1262ea4c99a89e94801b150e355a4e984c462bf8fb8a2fe9a33dacd74e0.deb
pool/kubelet_1.8.2-01_amd64_d740f0a12bba7d05f91e04223280967214dce453ca125ced549696d03f84e60d.deb
pool/kubelet_1.8.3-00_amd64_10bf853bba6b7713f8f3001f92612637ccee6893e699f7c41d91d2b4c3e6e25f.deb
pool/kubelet_1.8.3-01_amd64_7ddb604efa4a846fc72f03e7b5e03a3644a40a4a3570ca78269fc6d919616317.deb
pool/kubelet_1.8.4-00_amd64_601882506070723b643552ae98325c849840b09b1fc1666de74c7b69a07f06df.deb
pool/kubelet_1.8.4-01_amd64_e478234f2876b3170efac14ad2eb58893e8dd72fe1a7962575356d2445d8d1b4.deb
pool/kubelet_1.8.5-00_amd64_19c18818c0a5efc2fdbc86c87407db05d1de403b01a35e3a145b23cf3995ce5c.deb
pool/kubelet_1.8.5-01_amd64_af944a975e45ee2889cfb6117de955a68f7e8233ab831d073a179eccde1a3298.deb
pool/kubelet_1.8.6-00_amd64_39a5f313074cbc2dda28f97cc7739c352f0037895211c4f894845dc74063711b.deb
pool/kubelet_1.8.7-00_amd64_24344d1fccddb84f462d4458dc7ef60516b7b34c8f7e53050b8cebd904f1a8b0.deb
pool/kubelet_1.8.8-00_amd64_64d4ec64faa4ed93bf67900985ab8b87b0c96449c0a5b23dad8fe0138fef99de.deb
pool/kubelet_1.8.9-00_amd64_9319e46c4964ea2dbb3da011b491953ad95b5b038861ce94d59db714c74dd036.deb
pool/kubelet_1.8.10-00_amd64_f02917067d77dbcb01ab6b881f68314d1e6b4ff178560e2d596d9c3226726c79.deb
pool/kubelet_1.8.11-00_amd64_e38ba6a15aaeda29fe2866b8b1365702d4b4dea577d646df50da2f8c5ee762cf.deb
pool/kubelet_1.8.12-00_amd64_7dc5d77b1b85ce6eef0bb128ab5b1790affc0542822facf4515ffcea051b7781.deb
pool/kubelet_1.8.13-00_amd64_95894fc832a74a6be8461c516595b4e7ea12edf1e5f4ff20599d23fca50f491f.deb
pool/kubelet_1.9.0-00_amd64_100df9788226fe76ce828503cf24b8c4c6f9705f15504093844c9138e0b2a97f.deb
pool/kubelet_1.9.1-00_amd64_e1111bdd0cb45976216b7498a0b1d68d15ddaacc462a067d831b0f20ae8c276b.deb
pool/kubelet_1.9.2-00_amd64_cdb2335fec48cd8cb3ad2bdd874fb29828ceb893d200ebee44657ae51f0851b4.deb
pool/kubelet_1.9.3-00_amd64_50317ce119d81a3a69ba891c5c77b48b153b97a175598dc62ee811e081a063ee.deb
pool/kubelet_1.9.4-00_amd64_c8c2af29e421adf188c94b7049e30ed1b6368dad49cfe5e8c865d43a16be6b76.deb
pool/kubelet_1.9.5-00_amd64_5400807b4c667bfeaa5074245fb1edb20cf4b1b7f3ac88cb23df437f8558cd13.deb
pool/kubelet_1.9.6-00_amd64_abc92a438e6bcade2e879b789fa0634474f448df7428157703206448805a948f.deb
pool/kubelet_1.9.7-00_amd64_6a510a9b3d8b54473e8237cda857be4487dbc8cbd2b6b84ad292f5a3c3ec01c5.deb
pool/kubelet_1.10.0-00_amd64_211eced3f0715db1dbbb9fa0973299a8f6b764ef43e25bd94fece3947b13091f.deb
pool/kubelet_1.10.1-00_amd64_7002d1aa12fea2f0c81efdc5ea577607e387d21177c6fa57e1f897037325b824.deb
pool/kubelet_1.10.2-00_amd64_3ff69468521886de7b64ba9b4932bab7e6ef71f7344ace36ac0855df7bd2427a.deb
pool/kubernetes-cni_0.3.0.1-07a8a2-00_amd64_9e41a275b2afeb51dcde86b922c056c7b6dc656b54dd66fa2f1a0bb8266e9c22.deb
pool/kubernetes-cni_0.5.1-00_amd64_08cbe5c42366ec3184cc91a4353f6e066f2d7325b4db1cb4f87c7dcc8c0eb620.deb
pool/kubernetes-cni_0.6.0-00_amd64_43460dd3c97073851f84b32f5e8eebdc84fadedb5d5a00d1fc6872f30a4dd42c.deb
pool/rkt_1.25.0-1_amd64_65f4d768116520be6ee5d56ebcdce87b8b39fa2e41345b7c2960a01f91b7c081.deb
pool/rkt_1.26.0-1_amd64_af62ca3979b90f2fcbe7c48ba782dc1b0c4266832e7544ef4c69ce0a8a7dfa87.deb
pool/rkt_1.27.0-1_amd64_b4f1d67b835d1f1b0263ddfc3f8df80cd7c9e36ac78d983e12aca5575835e122.deb
pool/rkt_1.28.1-1_amd64_66ea8b0ef5724aa364a1273b794b2aaf5c302c774e076115adee8f7c04a0e0f9.deb
pool/rkt_1.29.0-1_amd64_ea87d719359030f33fd48890875c934135c62eccda72c37d79ff604307b905b5.deb
```

Update name into script and do execution
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ ./apt-pool-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 19.7M  100 19.7M    0     0   517k      0  0:00:39  0:00:39 --:--:--  971k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 8700k  100 8700k    0     0   433k      0  0:00:20  0:00:20 --:--:--  558k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 20.1M  100 20.1M    0     0   730k      0  0:00:28  0:00:28 --:--:--  833k
```

Just for V1.8.4
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ ./apt-pool-v1_8_4-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 17.2M  100 17.2M    0     0  38971      0  0:07:45  0:07:45 --:--:-- 44992
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 8409k  100 8409k    0     0  59412      0  0:02:24  0:02:24 --:--:--  409k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 18.3M  100 18.3M    0     0  40909      0  0:07:49  0:07:49 --:--:-- 47686
```

```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ egrep '^Filename: pool/\S+deb' https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64/Packages | awk '{print $2}'
pool/kubectl_1.5.6-00_amd64_620d1f16fcf779e72e778fd000b247612eabb351e054573b1153df8f6aad1342.deb
pool/kubectl_1.6.0-00_amd64_7021195dccb229242f160a98c069ccfe3097b3af1549b81940d201313238b519.deb
```

## v1.9.0 (Deprecated)

### Update RPM repository mirror

update repodata
```
[vagrant@kubedev-172-17-4-59 https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ ./yum-repos-el7-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 54792  100 54792    0     0  52435      0  0:00:01  0:00:01 --:--:-- 52482
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  8135  100  8135    0     0   8940      0 --:--:-- --:--:-- --:--:--  8939
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  230k  100  230k    0     0   195k      0  0:00:01  0:00:01 --:--:--  195k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 15287  100 15287    0     0  14282      0  0:00:01  0:00:01 --:--:-- 14286
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  254k  100  254k    0     0   209k      0  0:00:01  0:00:01 --:--:--  209k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 21173  100 21173    0     0  18724      0  0:00:01  0:00:01 --:--:-- 18737
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1417  100  1417    0     0   1477      0 --:--:-- --:--:-- --:--:--  1477
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   454  100   454    0     0    502      0 --:--:-- --:--:-- --:--:--   502
```

filter
```
[vagrant@kubedev-172-17-4-59 https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ samples=9; egrep '^  <location href=".*\.rpm"></location>' https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/primary.xml | cut -d '"' -f 2 | awk -v s=$samples '/1\.8\.4/{getline; for (i = 0; i <= s && $1 !~ /1\.5/; i++) {print; getline;}}'
../../pool/f9a3e9f13f7dacb8a39b90a2331007bebbed4f84643448e83e5c18b3efe3d45b-kubeadm-1.8.5-0.x86_64.rpm
../../pool/aa9948f82e7af317c97a242f3890985159c09c183b46ac8aab19d2ad307e6970-kubeadm-1.9.0-0.x86_64.rpm
../../pool/e5c2fec310104577d9746c20b61d078ecf842d34785f615d8a89550a48acc1a8-kubectl-1.8.5-0.x86_64.rpm
../../pool/bc390a3d43256791bfb844696e7215fd7ad8a09f70a42667dab4997415a6ba75-kubectl-1.9.0-0.x86_64.rpm
../../pool/d353208063910d0d9a8ff9a3e9bfb0cbee43be501bc4ca16d69223560c4a5cc0-kubelet-1.8.4-1.x86_64.rpm
../../pool/7330c09dd7dccc300c9c5f696fc4a8a76327e5068e51a60e4517b59a1c4b3dbc-kubelet-1.8.5-0.x86_64.rpm
../../pool/26944be0d909a90d8ef7559b17a3ce4381821d623314a35eb0c74e62d055d944-kubelet-1.8.5-1.x86_64.rpm
../../pool/8f507de9e1cc26e5b0043e334e26d62001c171d8e54d839128e9bade25ecda95-kubelet-1.9.0-0.x86_64.rpm
../../pool/0c923b191423caccc65c550ef07ce61b97f991ad54785dab70dc07a5763f4222-kubernetes-cni-0.3.0.1-0.07a8a2.x86_64.rpm
../../pool/e7a4403227dd24036f3b0615663a371c4e07a95be5fee53505e647fd8ae58aa6-kubernetes-cni-0.5.1-0.x86_64.rpm
../../pool/79f9ba89dbe7000e7dfeda9b119f711bb626fe2c2d56abeb35141142cda00342-kubernetes-cni-0.5.1-1.x86_64.rpm
../../pool/fe33057ffe95bfae65e2f269e1b05e99308853176e24a4d027bc082b471a07c0-kubernetes-cni-0.6.0-0.x86_64.rpm
../../pool/fd5f5da2d1a262fa687404d34e72813520274364557e648bc64a8136f1a02630-rkt-1.25.0-1.x86_64.rpm
../../pool/7a382e59dc2c39a66083e03ec061f33771e4a7130c98cd0ef61492b2196c0378-rkt-1.26.0-1.x86_64.rpm
```

edit yum-pool-curl.sh then download
```
[vagrant@kubedev-172-17-4-59 https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ ./yum-pool-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 16.4M  100 16.4M    0     0  1679k      0  0:00:10  0:00:10 --:--:-- 2002k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 9092k  100 9092k    0     0  1042k      0  0:00:08  0:00:08 --:--:-- 1454k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 16.7M  100 16.7M    0     0  1535k      0  0:00:11  0:00:11 --:--:-- 2694k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 8797k  100 8797k    0     0  1126k      0  0:00:07  0:00:07 --:--:-- 1981k
```

### Update DEB repository mirror

update dists
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ ./apt-dists-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  8951  100  8951    0     0   6155      0  0:00:01  0:00:01 --:--:--  6156
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  8448  100  8448    0     0  11219      0 --:--:-- --:--:-- --:--:-- 11219
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   454  100   454    0     0    561      0 --:--:-- --:--:-- --:--:--   561
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 94412  100 94412    0     0  94965      0 --:--:-- --:--:-- --:--:-- 94886
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 12051  100 12051    0     0  14840      0 --:--:-- --:--:-- --:--:-- 14841
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   207  100   207    0     0    262      0 --:--:-- --:--:-- --:--:--   262
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    23  100    23    0     0     27      0 --:--:-- --:--:-- --:--:--    27
```

Filter
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ samples=9; egrep '^Filename: pool/\S+deb' https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages | awk '{print $2}' | awk -v s=$samples '/1\.8\.4/{getline; for (i = 0; i <= s && $1 !~ /1\.5/; i++) {print; getline;}}'
pool/kubeadm_1.8.5-00_amd64_e6138f65d23ae9f18d3ecbb69d27cc40c30feff200d1652ce7814e365848ad88.deb
pool/kubeadm_1.9.0-00_amd64_191bd1d8a63d263cdb318f77b03fad6c8387a912cf16a21b9b24d7e9108b4911.deb
pool/kubectl_1.8.5-00_amd64_b9f161f02571cb379cfc93ed5a95e84b68b259cbc92d2f8291a5be3a6970ae37.deb
pool/kubectl_1.9.0-00_amd64_8ea712c18d89d56090c8753a9630d22fd8ae5cb03d4ec79a1fe6d262c8b4eb36.deb
pool/kubelet_1.8.4-01_amd64_e478234f2876b3170efac14ad2eb58893e8dd72fe1a7962575356d2445d8d1b4.deb
pool/kubelet_1.8.5-00_amd64_19c18818c0a5efc2fdbc86c87407db05d1de403b01a35e3a145b23cf3995ce5c.deb
pool/kubelet_1.8.5-01_amd64_af944a975e45ee2889cfb6117de955a68f7e8233ab831d073a179eccde1a3298.deb
pool/kubelet_1.9.0-00_amd64_100df9788226fe76ce828503cf24b8c4c6f9705f15504093844c9138e0b2a97f.deb
pool/kubernetes-cni_0.3.0.1-07a8a2-00_amd64_9e41a275b2afeb51dcde86b922c056c7b6dc656b54dd66fa2f1a0bb8266e9c22.deb
pool/kubernetes-cni_0.5.1-00_amd64_08cbe5c42366ec3184cc91a4353f6e066f2d7325b4db1cb4f87c7dcc8c0eb620.deb
pool/kubernetes-cni_0.6.0-00_amd64_43460dd3c97073851f84b32f5e8eebdc84fadedb5d5a00d1fc6872f30a4dd42c.deb
pool/rkt_1.25.0-1_amd64_65f4d768116520be6ee5d56ebcdce87b8b39fa2e41345b7c2960a01f91b7c081.deb
pool/rkt_1.26.0-1_amd64_af62ca3979b90f2fcbe7c48ba782dc1b0c4266832e7544ef4c69ce0a8a7dfa87.deb
pool/rkt_1.27.0-1_amd64_b4f1d67b835d1f1b0263ddfc3f8df80cd7c9e36ac78d983e12aca5575835e122.deb
```

edit apt-pool-curl.sh, then download
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ ./apt-pool-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 19.1M  100 19.1M    0     0   564k      0  0:00:34  0:00:34 --:--:--  956k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 10.0M  100 10.0M    0     0   787k      0  0:00:13  0:00:13 --:--:-- 1952k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 19.5M  100 19.5M    0     0   903k      0  0:00:22  0:00:22 --:--:-- 1543k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 5771k  100 5771k    0     0   587k      0  0:00:09  0:00:09 --:--:--  820k
```