# DevOps

Note: about kubernetes repository, refer to https://kubernetes.io/docs/setup/independent/install-kubeadm

## Table of Contents

1. [Mirroring Kubernetes RPM repository for CentOS/Fedora](#rpm)
1. [Mirroring Kubernetes DEB repository for Ubuntu/Debian](#deb)
1. CentOS yum (Fedora dnf) package mirror configuration 
1. Ubuntu apt package mirror configuration

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

Last execution (DEC-17-2019)
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm$ ./yum-repos-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  133k  100  133k    0     0  32437      0  0:00:04  0:00:04 --:--:-- 32576
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 21373  100 21373    0     0  10729      0  0:00:01  0:00:01 --:--:-- 10724
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  698k  100  698k    0     0   261k      0  0:00:02  0:00:02 --:--:--  261k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 40712  100 40712    0     0  20214      0  0:00:02  0:00:02 --:--:-- 20224
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  765k  100  765k    0     0   138k      0  0:00:05  0:00:05 --:--:--  176k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 62142  100 62142    0     0  31258      0  0:00:01  0:00:01 --:--:-- 31242
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1419  100  1419    0     0   1000      0  0:00:01  0:00:01 --:--:--  1000
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   454  100   454    0     0    156      0  0:00:02  0:00:02 --:--:--   156
```

### Pool

Find required package from _repos_
```
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm$ egrep '^  <location href=".*\.rpm"></location>' https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64/repodata/primary.xml | cut -d '"' -f 2
Packages/29e7806d1d54cc0eea2963f4ab276778526538816c88c963ece7d1a05fd80792-cri-tools-1.0.0_beta.1-0.x86_64.rpm
Packages/e253c692a017b164ebb9ad1b6537ff8afd93c35e9ebc340a52c5bd42425c0760-cri-tools-1.11.0-0.x86_64.rpm
Packages/f70d8cb23c7b91c0509292f4862060367edabce8788b855c38a7c174f014b9e2-cri-tools-1.11.1-0.x86_64.rpm
Packages/53edc739a0e51a4c17794de26b13ee5df939bd3161b37f503fe2af8980b41a89-cri-tools-1.12.0-0.x86_64.rpm
Packages/14bfe6e75a9efc8eca3f638eb22c7e2ce759c67f95b43b16fae4ebabde1549f3-cri-tools-1.13.0-0.x86_64.rpm
Packages/5116fa4b73c700823cfc76b3cedff6622f2fbd0a3d2fa09bce6d93329771e291-kubeadm-1.6.0-0.x86_64.rpm
Packages/23961d0f7dca1ed118b948195f2fb5dd7a07503d69d7d8ab4433219ea98d033e-kubeadm-1.6.1-0.x86_64.rpm
Packages/415bc9f86ecfc3b195f22d25fb599e19080d301358c145539d154be95f1985f5-kubeadm-1.6.2-0.x86_64.rpm
Packages/81158f40764a08356242a53ef4bf7e3c219f48f364c55260db571cae51ce6eec-kubeadm-1.6.3-0.x86_64.rpm
Packages/0c37012e387fae5fd0d66b5d4664f495ab19fbb709df655533b2b2ecb1a05cd8-kubeadm-1.6.4-0.x86_64.rpm
Packages/5e62d624af011598f2df59c7759bf6ff4a6bba5675f0c07a2fb18e3c468ea41c-kubeadm-1.6.5-0.x86_64.rpm
Packages/03ea47d2b0f1912c3a721c2b7353ead2e28a154b6c883200492db0558f6b09f2-kubeadm-1.6.6-0.x86_64.rpm
Packages/6b36b8d9783ee59480bda0caabcb9ce95a6bf69c62ea38e7bb0cee65174fe479-kubeadm-1.6.7-0.x86_64.rpm
Packages/67cb3635a4c3927b7009f0163589dfc0d261726a055893c52531c77b6e6bb0f8-kubeadm-1.6.8-0.x86_64.rpm
Packages/dce6dfbff7a60ac39f9288b211d8e4cf788a03ee64f2456ee6e0860fc549e1a9-kubeadm-1.6.9-0.x86_64.rpm
Packages/2bb1e670317f17fde3142c3e3549c91e15fc629ddfd0f274c290918c4289caf4-kubeadm-1.6.10-0.x86_64.rpm
Packages/c10b6d67b13101afd0270873972f10b14e9719ad8564797cc6609f7c58ae033c-kubeadm-1.6.11-0.x86_64.rpm
Packages/fda8179e8174e43c17662d3a019f54e862d270fbfc07c3f7a88a58a033198153-kubeadm-1.6.12-0.x86_64.rpm
Packages/ed8ab0500a5aa2d443653aa144c65fc246e3d047bee7cb1ed6499d0c926d0690-kubeadm-1.6.13-0.x86_64.rpm
Packages/cb7034dff51af01c237e03ea342dc74269b46cce1e43bb0ab9ef1d6f006ebda9-kubeadm-1.7.0-0.x86_64.rpm
Packages/77bea7f33178ab4feb6afaf31c6c779511498ef47dec433eb5464909f7a26dc3-kubeadm-1.7.1-0.x86_64.rpm
Packages/1a6f5f73f43077a50d877df505481e5a3d765c979b89fda16b8b9622b9ebd9a4-kubeadm-1.7.2-0.x86_64.rpm
Packages/f7ec56b0f36a81c0f91bcf26e05f23088082b468b77dac576dc505444dd8cd48-kubeadm-1.7.3-1.x86_64.rpm
Packages/f0a51fcde5e3b329050d7a6cf70f04a6cdf09eacfbad55f4324bfa2ea4312d0e-kubeadm-1.7.4-0.x86_64.rpm
Packages/02f3a7ff6e04943bd288ff302f449b600e8db3d19868dfe4308d0d902c0ba927-kubeadm-1.7.5-0.x86_64.rpm
Packages/4ff875dc8857b85c490b42b750527ba20a154a49a8dacd256d16cbbf5e708dfd-kubeadm-1.7.6-1.x86_64.rpm
Packages/2bb9ddc5197dec31ac73a549067ab9b6a529cc31275f9223ff04efbfb5451602-kubeadm-1.7.7-1.x86_64.rpm
Packages/8af3d230e7c1c775124b518aa0d69fa911e761e7b22fd081bfae8dece2381970-kubeadm-1.7.8-1.x86_64.rpm
Packages/e9e277f88f2747a493dd6b22360c58701c30216849e6ae74b190f4916b58125b-kubeadm-1.7.9-0.x86_64.rpm
Packages/c0638ef93a73ad4343fe3c2b8105e3826605e2fe000017324b3328975aa36a82-kubeadm-1.7.10-0.x86_64.rpm
Packages/b10ac8cf7ee52d4f4144b523b1f33061b7429bea3fcf9e863261423de090804a-kubeadm-1.7.11-0.x86_64.rpm
Packages/0ce3d048f9832d29000f6c69e05eb7a1fbc013810c7d1f66ad43e17f8578bd05-kubeadm-1.7.14-0.x86_64.rpm
Packages/7de188358ee8ee1babf424ab9d8d83ab2eb82325cba7133d97d8d5de3a365918-kubeadm-1.7.15-0.x86_64.rpm
Packages/9dda7d73d6c4cce878eb2fedaf2dfed78278202b0b5ee73dcd841b4edd643ae5-kubeadm-1.7.16-0.x86_64.rpm
Packages/ee93b5249497dba6488262bd2b56e32438a69f78993eb973bafa52d72f9b883a-kubeadm-1.8.0-0.x86_64.rpm
Packages/7fe94dd0f828ef64d4ccdfef160417457309b904758ddf10c040612d7e5ef441-kubeadm-1.8.0-1.x86_64.rpm
Packages/0f7d8ea10144399f3d60446fab5469395afb809c175bdc0eae4d12c1fcc3cb62-kubeadm-1.8.1-0.x86_64.rpm
Packages/d64bc1d0ca27196030c6f574252a0872b998b29be90c6cb7e97b91cd0bc78fed-kubeadm-1.8.2-0.x86_64.rpm
Packages/cab6b288e91e613d81c63101c7d293059a4a9f2c0795228042c880f770a9ec60-kubeadm-1.8.3-0.x86_64.rpm
Packages/aeaad1e283c54876b759a089f152228d7cd4c049f271125c23623995b8e76f96-kubeadm-1.8.4-0.x86_64.rpm
Packages/f9a3e9f13f7dacb8a39b90a2331007bebbed4f84643448e83e5c18b3efe3d45b-kubeadm-1.8.5-0.x86_64.rpm
Packages/919d83307b30c808a9bf17f08ab6d72612d08860a8923366e666ed072012f62a-kubeadm-1.8.6-0.x86_64.rpm
Packages/005fee933b99ae653baf508af37144c194d330b6324a62b32ebc08557a838937-kubeadm-1.8.7-0.x86_64.rpm
Packages/2e001cae2c8933ef43f6d5607a8bd57200957685bfb932dfb753b23abb9f8074-kubeadm-1.8.8-0.x86_64.rpm
Packages/e1cf1b76ab1c7d92a24d0ccccd2f3cdaae5b02bd063797938d27760531b76049-kubeadm-1.8.9-0.x86_64.rpm
Packages/1e84e0b4101a6ded3d395f77ff5736e6d15d9730b5b9fb00ddf625878ef585f4-kubeadm-1.8.10-0.x86_64.rpm
Packages/280eb1e1232c7dc99ec6e40bcd1ae3b7f72031a54bf33c7adb9abca64b0b7411-kubeadm-1.8.11-0.x86_64.rpm
Packages/59570c2320c2e8ead952893f25df6e4337c5ab3056ece69694520a05cadd0396-kubeadm-1.8.12-0.x86_64.rpm
Packages/bf953a68a7f2b897d88ad14922562387142334ae322057e5cb7c733e7fbb4ca1-kubeadm-1.8.13-0.x86_64.rpm
Packages/39976c316261b6cb7db6f994b22d2251ff5dd60cf58519de761946c18ef7d0aa-kubeadm-1.8.14-0.x86_64.rpm
Packages/0746100f7f8f3d05fc69a9f1c1015abac1f9ddc6964ef3ec3afdac33c6fb0cda-kubeadm-1.8.15-0.x86_64.rpm
Packages/aa9948f82e7af317c97a242f3890985159c09c183b46ac8aab19d2ad307e6970-kubeadm-1.9.0-0.x86_64.rpm
Packages/ccc2b7d8ac99c8ead43087e7bb5cc7fb3fe684bfd78241c8240feb945303c40e-kubeadm-1.9.1-0.x86_64.rpm
Packages/eb3f5ec9f28b4e6548acc87211b74534c6eae3cabfe870addeff84ac3243a77e-kubeadm-1.9.2-0.x86_64.rpm
Packages/df7f026762a109d02bf92ec559bbeb3f490e8e35885ffc2631042ffcd9b945ab-kubeadm-1.9.3-0.x86_64.rpm
Packages/12f95589196fb1331fb136084c16dcc647ef9db391307a64f6a23391a7afc6de-kubeadm-1.9.4-0.x86_64.rpm
Packages/4c45f9cada16d14ffd030aade7728b8ac5d148be70206cd3fbeb4188ee790f00-kubeadm-1.9.5-0.x86_64.rpm
Packages/f56f3294d633ecfa7f2aac506f7267c00547d4c529b134bc4698a563402897c3-kubeadm-1.9.6-0.x86_64.rpm
Packages/e569b64b2da8847696c665c44c0f0d5ba6fad79525230898219c514d79f1b1c3-kubeadm-1.9.7-0.x86_64.rpm
Packages/de3443400bd67115edb3969cc782ecce4839993a01ec4f45a821128e319ca116-kubeadm-1.9.8-0.x86_64.rpm
Packages/c43d6eb082ff68cc9efc4091f5905530017de90d10aa0dfaa2ca9fa84138ea4b-kubeadm-1.9.9-0.x86_64.rpm
Packages/dbba81c5ad57ef7d9f0af3130951b87cf9fc97429b5b2e6d04184dcccd6cf8ab-kubeadm-1.9.10-0.x86_64.rpm
Packages/efc0607a317f5338ef87b97f34a053e973e9011847bcb58d01aab464a257175b-kubeadm-1.9.11-0.x86_64.rpm
Packages/8b0cb6654a2f6d014555c6c85148a5adc5918de937f608a30b0c0ae955d8abce-kubeadm-1.10.0-0.x86_64.rpm
Packages/35ea034b2efccf2529cca8ed44f1bdcc0c3b26f0139694d8cbea315077a1bf6e-kubeadm-1.10.1-0.x86_64.rpm
Packages/b754a6990af7d7012189610b0dc69e6e950c13a8c415b9ebea8d56352e9719fd-kubeadm-1.10.2-0.x86_64.rpm
Packages/fbb4fa88ad2e4dcd32aa454bf3f334ca18be719dc32088cb8644eff208ef4567-kubeadm-1.10.3-0.x86_64.rpm
Packages/56f73d32965af49f582da3d790a145c121843c20e99fdcc6dc05a23f6afff338-kubeadm-1.10.4-0.x86_64.rpm
Packages/3ea9c50d098c50a7e968c35915d3d8af7f54c58c0cedb0f9603674720743de4e-kubeadm-1.10.5-0.x86_64.rpm
Packages/d335a12bd98045bce734300c976b330190291124b02a14092514604d4fda07cc-kubeadm-1.10.6-0.x86_64.rpm
Packages/644a5ba4e2c167588f424681234fe554bbc1177ac948b2292851d240c3698336-kubeadm-1.10.7-0.x86_64.rpm
Packages/5d1013f714965d28e1aa73de7fbd3dc91bb6c02c60e899603a51c6eadd17ef73-kubeadm-1.10.8-0.x86_64.rpm
Packages/d9530d3c2a50af3fdde899bb6bbc0e0f5a40390ba751a3ba319a017109f36ff3-kubeadm-1.10.9-0.x86_64.rpm
Packages/0f86642931be4b9fc96ebed1a7fb425dc90b1eb59d76886266245c07e251baf2-kubeadm-1.10.10-0.x86_64.rpm
Packages/35123b3f59f4c2fbdeac90b05ac8aab1afd5f8554491b796a94eb46ee2ebd1cc-kubeadm-1.10.11-0.x86_64.rpm
Packages/6cf7336247f0cee728548c92956baeb267d544740c7e80e7b3c0aa37209367b4-kubeadm-1.10.12-0.x86_64.rpm
Packages/5bc6edbce5a26d13aad44dbecde701caf1cc40b2747a5255dae1afdbf879443e-kubeadm-1.10.13-0.x86_64.rpm
Packages/c276be7b12fbca50df96fdc0fc07ecbb758060a31b4ef3c1e8908b7bd9ebddcf-kubeadm-1.11.0-0.x86_64.rpm
Packages/40ff4cf56f1b01f7415f0a4708e190cff5fbf037319c38583c52ae0469c8fcf3-kubeadm-1.11.1-0.x86_64.rpm
Packages/2cb76e20a5cf7c7006b5e13c2f53838eda7a3c1283773646d7b2ec4dc63d417a-kubeadm-1.11.2-0.x86_64.rpm
Packages/0576a2ac6e3ff9e96d807b33ed6caabcfc3cec94e333fb86431d99a398318251-kubeadm-1.11.3-0.x86_64.rpm
Packages/9d463dcfbedec747775fa979d45a2bea915b3326f92b4b079c1ee4a94a64cc5e-kubeadm-1.11.4-0.x86_64.rpm
Packages/44c6c5f096761d55cce0200a32eecc2d7c000d2a0492724acbf0cb10de04a008-kubeadm-1.11.5-0.x86_64.rpm
Packages/bdebac39334ab4450a7f8d502c6ed56f0d65cfe9afd6fc172ddebbf5bfb15bd7-kubeadm-1.11.6-0.x86_64.rpm
Packages/e2c8f2b74ee6f4205fa490d396379cb01698e90642bf8c711771dded6eada4a8-kubeadm-1.11.7-0.x86_64.rpm
Packages/b5383ee71cb1a335216ff1da8ad41614ab57fa3156ce07cc696096269fc4d8db-kubeadm-1.11.8-0.x86_64.rpm
Packages/2d3fea397e599d94ec3d67582c522719ca3e9bd946c0f89a33caa07a49e1b883-kubeadm-1.11.9-0.x86_64.rpm
Packages/fa3dfb16bbf2e91381debe50eeef2040d4af723fa799c12107b413e1f0ca2491-kubeadm-1.11.10-0.x86_64.rpm
Packages/c07649bdb3bd346d55cdb4b786f3e6076f1511834bbb41d94e5f380a0c41a9af-kubeadm-1.12.0-0.x86_64.rpm
Packages/9c31cf74973740c100242b0cfc8d97abe2a95a3c126b1c4391c9f7915bdfd22b-kubeadm-1.12.1-0.x86_64.rpm
Packages/6bd058ff754287c0b6b7431ee8f08bd35af40b0b1d098d94acf1448de0c8053b-kubeadm-1.12.2-0.x86_64.rpm
Packages/3bc525742ac6b36379c3ef1cca788bb55525842c634c91da2389ebb7d57e2ac5-kubeadm-1.12.3-0.x86_64.rpm
Packages/ef8739a3a637246743ee5238ac929308d7c322c0ad1a8806699af04434e545b3-kubeadm-1.12.4-0.x86_64.rpm
Packages/ecd982369a9c7ad728faa96d891d47f8f1cc54b1d2dac72f7953863d771aecb1-kubeadm-1.12.5-0.x86_64.rpm
Packages/5af46e521696f654883099229af8f2b5f6c208ab5996b8385f017f01b24f13d0-kubeadm-1.12.6-0.x86_64.rpm
Packages/a7aaf0551faafa396a9b69a961486486430b91be00fe0034cfd5aa9a4f56298c-kubeadm-1.12.7-0.x86_64.rpm
Packages/b8985d1f6cbda39c20f52e45bc2fa9aac877d2ced4e482b1506c5b0cf1f712f5-kubeadm-1.12.8-0.x86_64.rpm
Packages/45586bda53ff222d2b6757d2d5a2e3f650ac3fe54b1c39f789bee2b4ecb242ba-kubeadm-1.12.9-0.x86_64.rpm
Packages/4943c168a879e38d88e2d9eb4be4ac0d462ec88ead8f6eb2e25fa44cb137415b-kubeadm-1.12.10-0.x86_64.rpm
Packages/4c66c2a8c4ed256bc5e41f2cb08d6a32d704b510e0f74f737c5bdad69b49774f-kubeadm-1.13.0-0.x86_64.rpm
Packages/5af5ecd0bc46fca6c51cc23280f0c0b1522719c282e23a2b1c39b8e720195763-kubeadm-1.13.1-0.x86_64.rpm
Packages/416b2856f8dbb6f07a50a46018fee8596479ebc0eaeec069c26bedfa29033315-kubeadm-1.13.2-0.x86_64.rpm
Packages/9cc01fa0dc1866086d4bcb7096be0c3def7b529da708e7e3793588a95a2cbc41-kubeadm-1.13.3-0.x86_64.rpm
Packages/ca5618c4268261f3deb7cc4cb12fedf8735db81f6a327436c2b8418b61295256-kubeadm-1.13.4-0.x86_64.rpm
Packages/16f46a4eb371ba77ea3cafab76399aba50f223201534422182698f891fd31cab-kubeadm-1.13.5-0.x86_64.rpm
Packages/2a9062c7b4eeacc363321b6e17dbdc69d20443a7f099b71dd2add8dedcbaf75c-kubeadm-1.13.6-0.x86_64.rpm
Packages/f1c6a4234f90f7bf9970442aa48bb8f96f1990a6b60db373048c6696b0ec92b8-kubeadm-1.13.7-0.x86_64.rpm
Packages/a85b40da362eb1ab08c8b1a06882644ecba27227514a2842fcea9aba97728629-kubeadm-1.13.8-0.x86_64.rpm
Packages/4089ed790752f8ee4567b92af5d2ab64b42192b105d3e21f59512bfc9161ff61-kubeadm-1.13.9-0.x86_64.rpm
Packages/5fda481d5850ab6085797ed3d7aeee6394fba4117b6c9c759263466d358b7630-kubeadm-1.13.10-0.x86_64.rpm
Packages/b2e88a9cd451ce177c15c2aad451035dab58556dab66115e94564a3328370c29-kubeadm-1.13.11-0.x86_64.rpm
Packages/e8a8c1b525d40f9f2bbf1ec1c53e8545665e23692a00f805198439dac3fd3f7b-kubeadm-1.13.12-0.x86_64.rpm
Packages/fea2c041b42bef6e4de4ee45eee4456236f2feb3d66572ac310f857676fe9598-kubeadm-1.14.0-0.x86_64.rpm
Packages/9e1af74c18311f2f6f8460dbd1aa3e02911105bfd455416381e995d8172a0a01-kubeadm-1.14.1-0.x86_64.rpm
Packages/de639995840837d724cc5a4816733d5aef5a6bf384eaff22c786def53fb4e1d5-kubeadm-1.14.2-0.x86_64.rpm
Packages/0ac0a679e4fe90bc49165841f8cdd7527940b05a5bef641ded55fd828c21ec0d-kubeadm-1.14.3-0.x86_64.rpm
Packages/18fecdcafd7a4ab490285c4c96e69ebbce1a627a6e1928d6d1d3a691c3640e64-kubeadm-1.14.4-0.x86_64.rpm
Packages/fa4b18cda7e2d930baf62fdde559d62198b7f0b7a9f5d64bb141cb596514049d-kubeadm-1.14.5-0.x86_64.rpm
Packages/cc14b325a3549f2a7b1c032573e5a1de27b221dc272569377728b00b7e82ab06-kubeadm-1.14.6-0.x86_64.rpm
Packages/a2186f03a518583d1a78b96ac08fa2493053f7646943c9eff55ed0fd4756abe4-kubeadm-1.14.7-0.x86_64.rpm
Packages/c1101e7903201b851394502c28830132a130290a9d496c89172a471c2f2f5a28-kubeadm-1.14.8-0.x86_64.rpm
Packages/b65516938c9d9096bd1d772f446375df1b5ef83f46d7457501253e285a134c59-kubeadm-1.14.9-0.x86_64.rpm
Packages/e2de8688286f5f4516813d1346654a50d64c4edf68a7bbeae98946cd99eb41a7-kubeadm-1.14.10-0.x86_64.rpm
Packages/7143f62ad72a1eb1849d5c1e9490567d405870d2c00ab2b577f1f3bdf9f547ba-kubeadm-1.15.0-0.x86_64.rpm
Packages/aa386b8f2cac67415283227ccb01dc043d718aec142e32e1a2ba6dbd5173317b-kubeadm-1.15.1-0.x86_64.rpm
Packages/b3e3099a44905838335bf405565659ba6113699d8457ccf0e3462dd47691de65-kubeadm-1.15.2-0.x86_64.rpm
Packages/d87e7e3e61cc561f18376d2caa53207ba96b932ca2011be9e8e370bdc281d859-kubeadm-1.15.3-0.x86_64.rpm
Packages/2813cf105c52ad1240f2cc6cba9a3f779bf2f5c4940c731a27df6e5d9557a5b1-kubeadm-1.15.4-0.x86_64.rpm
Packages/a1ae562a4bcac2ccc85a4a7947cd062ecab691011ec59657bc705318e7477143-kubeadm-1.15.5-0.x86_64.rpm
Packages/62cd53776f5e5d531971b8ba4aac5c9524ca95d2bb87e83996cf3f54873211e5-kubeadm-1.15.6-0.x86_64.rpm
Packages/82c26855d0f055599563d1b46866a7b3f1b009b18f60f91dfbbeb5e078c1852c-kubeadm-1.15.7-0.x86_64.rpm
Packages/47a816a6535a89c89622fb9c0c7f0031c3d5426cb7df91a952f9a98b118e21b5-kubeadm-1.15.8-0.x86_64.rpm
Packages/b7ad3944a7ae65db3b8854d16e845dce02b4bc7226eb41385c4fe96265424828-kubeadm-1.15.9-0.x86_64.rpm
Packages/697ad1a31f01e90f44ad3f0c8fe06f32d7bdd3b227fcf705275d5ad241d52eb6-kubeadm-1.16.0-0.x86_64.rpm
Packages/a137f796adacb1e9e52b1a46faaf5406bfca203314708067c7d0f18076d5f949-kubeadm-1.16.1-0.x86_64.rpm
Packages/bd3de06f520c4a8c0017b653e2673cd6cd1b1386213b600f018fb67d93ffd60b-kubeadm-1.16.2-0.x86_64.rpm
Packages/b45a63e77d36fc7e1ef84f1cd2f7b84bccf650c8248191a37d20c69564d8b8df-kubeadm-1.16.3-0.x86_64.rpm
Packages/937e48b992617a0eeae40c93ad2105005ebf8491b4cbc1a43d2bcb35e8a9e865-kubeadm-1.16.4-0.x86_64.rpm
Packages/ed05d02b026ce1785b2709553190d24818555600bbd085e4311692f29df10218-kubeadm-1.16.5-0.x86_64.rpm
Packages/0eeb459890b1c8f07c91a9771ce5f4df6c2b318ff2e8902ed9228bf01944cfd7-kubeadm-1.16.6-0.x86_64.rpm
Packages/2c6d2fa074d044b3c58ce931349e74c25427f173242c6a5624f0f789e329bc75-kubeadm-1.17.0-0.x86_64.rpm
Packages/0ec02f105d2a5d0059711f73187d45cf5a03d366fe6e895cfc5c3d3d73e88b4e-kubeadm-1.17.1-0.x86_64.rpm
Packages/105d89f0607c7baf91305ba352e78000bd20aad5cdf706bffff3b31cd546dbf3-kubeadm-1.17.2-0.x86_64.rpm
Packages/fd9b1e2215cab6da7adc6e87e2710b91ecfda3b617edfe7e71c92277301a63ab-kubectl-1.5.4-0.x86_64.rpm
Packages/e6aef7b2b7d9e5bd4db1e5747ebbc9f1f97bbfb8c7817ad68028565ca263a672-kubectl-1.6.0-0.x86_64.rpm
Packages/9d1ccf0877dfc4228a923a9614661b27d663694680e2bc0a1c184aa937fbf7f2-kubectl-1.6.1-0.x86_64.rpm
Packages/ff72cbf0dfa986c36097a5c3ac2301ecb73ed28df8db86e13641a79e9df9b3ea-kubectl-1.6.2-0.x86_64.rpm
Packages/d5cc6bb2e439298eb1b3e45c3ac58010580c1d3c3a2fa71040a7c73b1dc22881-kubectl-1.6.3-0.x86_64.rpm
Packages/9fcf36d1acf5e17c4aec783e956b8b70ab0e91693d394e6268c16533fe8a71a9-kubectl-1.6.4-0.x86_64.rpm
Packages/504e6824515b1b91702d1cd9aa8701fbcb2fabeb61e8395204f4e2cd07f7dfb7-kubectl-1.6.5-0.x86_64.rpm
Packages/b9584919848f1cc3096ec8cd0bbeceb95ea0786df339eb4883556a4a89f51010-kubectl-1.6.6-0.x86_64.rpm
Packages/b542021d5f32457f8c1f6f726aaa077eec66b0906440a020cfada28275df50f7-kubectl-1.6.7-0.x86_64.rpm
Packages/396d4c944b93310b86afad34627390e31a3f839ca1f7f91df95919ac4c85065c-kubectl-1.6.8-0.x86_64.rpm
Packages/27a094522739c8fa954133ca6b9d28b1e39537d7a9cd33a1b6e2e586e9e6811d-kubectl-1.6.9-0.x86_64.rpm
Packages/ad4326997f2a06c4802ce8ae170648f6a07522f5de629be89d5324ecfe843780-kubectl-1.6.10-0.x86_64.rpm
Packages/9b3b2d1cbe36086a7008d261594b31bb3e085e8c170aa319e1befebe604a99a7-kubectl-1.6.11-0.x86_64.rpm
Packages/712b311a1853d2b73be6729f57fc448ebc8818f0a9236cbbf1f11c6db3c3d8dc-kubectl-1.6.12-0.x86_64.rpm
Packages/0e7baf0e754f3c62dcbb50d850167f0d8eee0819aab40cb3932ddc72f098f4ea-kubectl-1.6.13-0.x86_64.rpm
Packages/888aff6bf68f988e7480efd06f38852eef8a44eed1baaa62f3ec4fbb07c35f7d-kubectl-1.7.0-0.x86_64.rpm
Packages/b8a64634ad1555e15a61b84685fd04959435ed6374b25e369f5bda89c8f03a6b-kubectl-1.7.1-0.x86_64.rpm
Packages/dc8329515fc3245404fea51839241b58774e577d7736f99f21276e764c309db5-kubectl-1.7.2-0.x86_64.rpm
Packages/c8a50a1ce9c948e7a238b02d3967d220e71e13e1ac8916a961176726eabe8aa1-kubectl-1.7.3-1.x86_64.rpm
Packages/041d5a6813dab590b160865fea7259bc2db762a9667379d03aca8d4596a3cccd-kubectl-1.7.4-0.x86_64.rpm
Packages/c2d4b7c1f39ca9532a2965ea513fcd900fdcdd8785639c8fbf69f30780cae764-kubectl-1.7.5-0.x86_64.rpm
Packages/71aa78fc7472de3664511c88f9d58d9a9c6922f26d67323869b5a661b106e0d0-kubectl-1.7.6-1.x86_64.rpm
Packages/4af2eb4114017f12fb36b92a85c8149de6c54020a73061f3666d033bfde8f3e3-kubectl-1.7.7-1.x86_64.rpm
Packages/761837737577fe362b16ff2b1fc34c383d86b0f9f746a9901c62587fd5f4e0f6-kubectl-1.7.8-1.x86_64.rpm
Packages/2ebbb529212cb34b3d3a973fc60c6d4d1039b34ef2b11c79cbea2e1b6b6f5bfc-kubectl-1.7.9-0.x86_64.rpm
Packages/ef4f06f44a3e9412735722d5625f814505054bd5f672ccad3e631470e5da9cd0-kubectl-1.7.10-0.x86_64.rpm
Packages/7da91dc73456ad61f773fdd893a96b22b9537da395658b29c2cbda376f701f50-kubectl-1.7.11-0.x86_64.rpm
Packages/4919fb3d1a5908461418f8a901c1c89bfd37e3418d72bbd4a8a215104015dae6-kubectl-1.7.14-0.x86_64.rpm
Packages/5e509f4b46906b21ec01844d1566374b2699ea2e27f82c7902df6e94381c516e-kubectl-1.7.15-0.x86_64.rpm
Packages/2d60f27c4698dcf7575809b91202f546caeb59bbe775ca202c5f69625a863a51-kubectl-1.7.16-0.x86_64.rpm
Packages/582709dfb1c59f85c78f7db2c7609a1f3bf7308b73255bdbc910695b31b8660f-kubectl-1.8.0-0.x86_64.rpm
Packages/0ee48e6b4033fdf520f5893759b0665090ffb83eefdbe3f0b41edf54f2247ee4-kubectl-1.8.1-0.x86_64.rpm
Packages/3cc05eb1b1565779e8033743f1a2b9c8fb4e3b432421719ec56a3024d33dfccc-kubectl-1.8.2-0.x86_64.rpm
Packages/ae43f92f96e828779f9744b3660e207199d97dda22eb44c054d2f3150da76b94-kubectl-1.8.3-0.x86_64.rpm
Packages/a9db28728641ddbf7f025b8b496804d82a396d0ccb178fffd124623fb2f999ea-kubectl-1.8.4-0.x86_64.rpm
Packages/e5c2fec310104577d9746c20b61d078ecf842d34785f615d8a89550a48acc1a8-kubectl-1.8.5-0.x86_64.rpm
Packages/432a40bb7dbe35e3c85cd926b015269b8de54809ad0f14ceecfd5a05acbe44a4-kubectl-1.8.6-0.x86_64.rpm
Packages/d8cbd7dd54e61c27d437cc32f2af5168ae693b888a7fa94f86583be94e9e8fbb-kubectl-1.8.7-0.x86_64.rpm
Packages/bb633ec4b637d8dbfdba4b50f427dffd1347ae8bb44c49ffc75809d6d052eab9-kubectl-1.8.8-0.x86_64.rpm
Packages/d8701448da60b3b00bc02930f5418ad19e8e9da398442e6d2b806e6d9e2cdc74-kubectl-1.8.9-0.x86_64.rpm
Packages/83b1b027b99b38cb67713bd4c60e7fc3f08d2a60f53b28fe7e6ccc64a47bb1f2-kubectl-1.8.10-0.x86_64.rpm
Packages/a5fdacbb2676408f4fa8dc779ef442d54f8af32878741c9d3e18d7d278ab241b-kubectl-1.8.11-0.x86_64.rpm
Packages/3182954b1bf7afabe1d3ad39888800c54c80f4b6a03bbef1df84732e5c9f3a1f-kubectl-1.8.12-0.x86_64.rpm
Packages/ced5e4a04e8fe35ec61962717fb996e5bc712c3c523451cf9d8882937ba22a73-kubectl-1.8.13-0.x86_64.rpm
Packages/6eb894d76ca3990cf837467237556f047b5d2e6f048a5474dd7e91ac138f84c8-kubectl-1.8.14-0.x86_64.rpm
Packages/c190976cdadf204b3774d28de9574c4e9b0ef8ff90415841961439621dcb7041-kubectl-1.8.15-0.x86_64.rpm
Packages/bc390a3d43256791bfb844696e7215fd7ad8a09f70a42667dab4997415a6ba75-kubectl-1.9.0-0.x86_64.rpm
Packages/331270b6ba931a571640b0552d8737631cf30cc9cde14eaa45a5acd2afb8f304-kubectl-1.9.1-0.x86_64.rpm
Packages/2dd849a46b6ca33d527c707136f6f66dfc12887640171333f3bb8fab9f95faac-kubectl-1.9.2-0.x86_64.rpm
Packages/eb54809406a7bc73f1996ce669323b65f0846d1830ef04a1e69e66f633a627a7-kubectl-1.9.3-0.x86_64.rpm
Packages/9566533de10dab5d80d8e7a8799ffea78e2b0f43e838e3b9d08f5e51c4dbc7ba-kubectl-1.9.4-0.x86_64.rpm
Packages/b8fb4f730cd9b661657b8baf0119e291add74e753d50e313843cfee6e0e65577-kubectl-1.9.5-0.x86_64.rpm
Packages/c9a30a9b3cd4f8b83a3ffcbfe5a23b32c0c78ec90a8e67505ba4ae31ed1d7a69-kubectl-1.9.6-0.x86_64.rpm
Packages/051265818f6c9d68fa4d5b582c996facf3292b67256ce8519f658c058fe8a85c-kubectl-1.9.7-0.x86_64.rpm
Packages/a67c69b918bb9b017599e392c5a81528218e44bdb9771d42f82d23c37b4d5805-kubectl-1.9.8-0.x86_64.rpm
Packages/00c0c484243f992e70c799383fcbbca2239fa7545a9d50e61f706f03de369196-kubectl-1.9.9-0.x86_64.rpm
Packages/f43e84212658037104239e44173df6d71d8359922eca4cac7d00546e7e24370c-kubectl-1.9.10-0.x86_64.rpm
Packages/f5d3ae18af92fa61ff4e38fd388a5688bb6af9a49ab7a26863b6a36b4b754766-kubectl-1.9.11-0.x86_64.rpm
Packages/9ff2e28300e012e2b692e1d4445786f0bed0fd5c13ef650d61369097bfdd0519-kubectl-1.10.0-0.x86_64.rpm
Packages/b8f0cc3bc85e8614f0340547b14ca5377410afd087f51077410945a594f0b71b-kubectl-1.10.1-0.x86_64.rpm
Packages/32e8bd812a3944ccf07750d52088a118fa11493d34e009e2873317e0f0b0dfd2-kubectl-1.10.2-0.x86_64.rpm
Packages/571c54a5e4049647541a24d77337898fb4243f6b39c7f3df5d92ab180055bd87-kubectl-1.10.3-0.x86_64.rpm
Packages/3bfc22d32a20afc46175242700455a62bb88eda8edf0867295ec83098de09aaf-kubectl-1.10.4-0.x86_64.rpm
Packages/1eed768852fa3e497e1b7bdf4e93afbe3b4b0fdcb59fda801d817736578b9838-kubectl-1.10.5-0.x86_64.rpm
Packages/25f64ebf153c034919d0343e45e71066a03b85db7257a40ddaa5ce360cc40959-kubectl-1.10.6-0.x86_64.rpm
Packages/c26a21174ab52d9c553c8ba5dd54e55c29d838a3bf1982873c30fe218e873e13-kubectl-1.10.7-0.x86_64.rpm
Packages/6782ef80bd50ecbb5326d8e5aaff04eba0332f43df1a52ef301b87e5570d57aa-kubectl-1.10.8-0.x86_64.rpm
Packages/cfff5f0efc01e895f65726c0b05e21ed8235a79cbf71c7dfb410aa503e4e68df-kubectl-1.10.9-0.x86_64.rpm
Packages/806d94bbb8d96f8837662a306167ab81b9d94577f00c2ac525ca39039fc4de8f-kubectl-1.10.10-0.x86_64.rpm
Packages/6498f6a56309d21b454f50424d7739bb01309235ac1e7a465129324fdbb8f8bc-kubectl-1.10.11-0.x86_64.rpm
Packages/3d82ae28cd39aa5f80b5b68b74d63b4a4c84b31ba7973c08b8ae72c1e4eaeade-kubectl-1.10.12-0.x86_64.rpm
Packages/4815089dde3fc771734b222016132caa8c8a2848048aa2a2169d319436e8d032-kubectl-1.10.13-0.x86_64.rpm
Packages/5736d31b18c9a00419105394d462ecec4847e6a9bbc7b5a4e41790a67e29c817-kubectl-1.11.0-0.x86_64.rpm
Packages/ac4c5b25cbe9d3eab5e6e83b5e7b43545645496d7dc2b37137f486477574e7c5-kubectl-1.11.1-0.x86_64.rpm
Packages/a554c1728ecf79871b4d3e0fc797568e53149f4ed7ec7e437c949a02f197a1ab-kubectl-1.11.2-0.x86_64.rpm
Packages/a36ef4d0907a8d32359931252b65516b46bb72e19f997574ac72fe3e88c0b03c-kubectl-1.11.3-0.x86_64.rpm
Packages/4fc73ec9258ada89e60ce0f905bf9468b8fbf9a9b4030302396b3230e264a566-kubectl-1.11.4-0.x86_64.rpm
Packages/ae867e22960a1f2646b7c05943d1f6aaa854605a02ac8b6c4d66b7c22b461c94-kubectl-1.11.5-0.x86_64.rpm
Packages/76a32ce9b4195b5078fac9bd0285e8dfc87a7f8dcf7c0615891b05fd66dbab52-kubectl-1.11.6-0.x86_64.rpm
Packages/8349803e982cc9d1d533feef43e7156310ec6d1cc605f82dee10bae8e964c1b6-kubectl-1.11.7-0.x86_64.rpm
Packages/f76cd4b1ad135a9b22b4cbb580a718ce9e87a7403d09c91e60015895e3df6941-kubectl-1.11.8-0.x86_64.rpm
Packages/542cbe1b886808a7f025595b2bfe776fd08de5d2dab0a80f22947cf400b89a2c-kubectl-1.11.9-0.x86_64.rpm
Packages/5ce34f8aadc7cf77445262e9e8ddc530753c01c1ad79a87022da10bbc7607193-kubectl-1.11.10-0.x86_64.rpm
Packages/c49bc7ad8a3cc689bc3c722c8d7530f7dc1fa4bc2e3682db15e4c9de321936ac-kubectl-1.12.0-0.x86_64.rpm
Packages/ed7d25314d0fc930c9d0bae114016bf49ee852b3c4f243184630cf2c6cd62d43-kubectl-1.12.1-0.x86_64.rpm
Packages/ead06eb2dc5ff86ca39823fa5f9e944b2f197e681b76ea0ba4a72f5ca6c51f32-kubectl-1.12.2-0.x86_64.rpm
Packages/6b5372132635c9ebca23e1adafbc8352bab9a1c6ba027f2e718817cfe558fcdd-kubectl-1.12.3-0.x86_64.rpm
Packages/e1ef6b554386a454ddc0d711b7950e5a95ecd5ea750e9c3aa6eb235e15bc7a80-kubectl-1.12.4-0.x86_64.rpm
Packages/60a0ade4faf3cbda6546c931f39009cbb709c9ce93a71f2f6f138f1577bf1d4f-kubectl-1.12.5-0.x86_64.rpm
Packages/29d4de58bb91d8705e2cb4f1227d05887dc9f2a1313b4c6631db6b522f78b774-kubectl-1.12.6-0.x86_64.rpm
Packages/a68b4fbdee8689c7f78dab3fcbcfbc13eac7d7aa1eb708c313ab9a80512b6a0a-kubectl-1.12.7-0.x86_64.rpm
Packages/958aed8f592cf5b47d9a3ddd7468a1e65d1b966346cd7f11e72987dbbbbd062d-kubectl-1.12.8-0.x86_64.rpm
Packages/3f74d9a31a2fccc6f3e58d24bc166bdbaf7290fc2a61b3d2cb76e3b2aeb249e5-kubectl-1.12.9-0.x86_64.rpm
Packages/1ac30b454049ba3e7896a648e20dab138de45ea7a162d7d1d918e676cfe29d63-kubectl-1.12.10-0.x86_64.rpm
Packages/61676fc8f5ea7e3b6ea765d38162e961727a802304bb9bf7feec24549a184f36-kubectl-1.13.0-0.x86_64.rpm
Packages/7855313ff2b42ebcf499bc195f51d56b8372abee1a19bbf15bb4165941c0229d-kubectl-1.13.1-0.x86_64.rpm
Packages/5ed19f5427225e4d0afcc4b2387b87e0a6b85fb57e31904d80e9f638a0e66837-kubectl-1.13.2-0.x86_64.rpm
Packages/e3438a5f740b3a907758799c3be2512a4b5c64dbe30352b2428788775c6b359e-kubectl-1.13.3-0.x86_64.rpm
Packages/7d8997e89f52340aecaa7fc2b1586665d2d9b568a8c887a8c84abc651b87c9ff-kubectl-1.13.4-0.x86_64.rpm
Packages/e93cef1270a595348aceceee32275c4242d6ea0b1c6a0489ec3f7548c80be5f1-kubectl-1.13.5-0.x86_64.rpm
Packages/423b37843aa713a14771ef55bf9bd3549ae0593a4bc8e31c07ecbd878294eb87-kubectl-1.13.6-0.x86_64.rpm
Packages/d03962249df8e3dfce1f893c78baee0c542201dce2345c0e0faef049bd6361f8-kubectl-1.13.7-0.x86_64.rpm
Packages/d11337018fc03b8c579444fe6e4d23fc63dae855d5f74407d34d1a406f0b33b2-kubectl-1.13.8-0.x86_64.rpm
Packages/8cbe9e26be5b90257310818b22b22f417572b464fe35bb627b9975f72ba2c497-kubectl-1.13.9-0.x86_64.rpm
Packages/a4530744eb83edd8a7395517844ce7becddd71ec2611ea5f9b59390b0de59d20-kubectl-1.13.10-0.x86_64.rpm
Packages/3f3ef061d8af6d85234d846328f7eb88e3ac5f4c1366368ff5acaba6064eebff-kubectl-1.13.11-0.x86_64.rpm
Packages/fbc121a458bbd7db9a547282e5b9fb8a15117d86d1df21e7e034c17f3fda839b-kubectl-1.13.12-0.x86_64.rpm
Packages/2b52e839216dfc620bd1429cdb87d08d00516eaa75597ad4491a9c1e7db3c392-kubectl-1.14.0-0.x86_64.rpm
Packages/5c6cb3beb5142fa21020e2116824ba77a2d1389a3321601ea53af5ceefe70ad1-kubectl-1.14.1-0.x86_64.rpm
Packages/7adc7890a14396a4ae88e7b8ed44c855c7d44dc3eefb98e4c729b99c2df6fa03-kubectl-1.14.2-0.x86_64.rpm
Packages/8425067aee8e57b64d3eff0d78817c30e486d58a6b749fa55a46366bb9addb7e-kubectl-1.14.3-0.x86_64.rpm
Packages/9cebc278102d9801308e0145ca44be1d5f74ec01db653c937f0d259ac714c081-kubectl-1.14.4-0.x86_64.rpm
Packages/4b25ffaf3846e4f3ec885814d80c863f19c795ce1a8dda67f472b6bcca2e91e5-kubectl-1.14.5-0.x86_64.rpm
Packages/e3993ee3349729053a2c38398473b888e0461de177fb86e559b57413ef7beb9d-kubectl-1.14.6-0.x86_64.rpm
Packages/c0b279e7d3fd4eb1e31524173b6beb9458b3106e18ff0275eeca12ba8b581ccc-kubectl-1.14.7-0.x86_64.rpm
Packages/2de91de7e1a27c52533426a65bb9dba3aa255a92844eaa22fb9801e8d5585a42-kubectl-1.14.8-0.x86_64.rpm
Packages/a26435dba43300a07536d53f03e6e7bd1ec91fe4878c6f527bcfce8ef3196656-kubectl-1.14.9-0.x86_64.rpm
Packages/0c130e4f7834bdc9e6cf088f30331520d424f667e6598c3e7a8c1b403c9a6e74-kubectl-1.14.10-0.x86_64.rpm
Packages/3d5dd3e6a783afcd660f9954dec3999efa7e498cac2c14d63725fafa1b264f14-kubectl-1.15.0-0.x86_64.rpm
Packages/f27b0d7e1770ae83c9fce9ab30a5a7eba4453727cdc53ee96dc4542c8577a464-kubectl-1.15.1-0.x86_64.rpm
Packages/bfc2b9cfa0bcfa8dd9e90edb0f0754e6256ba466de7afbd75da014e412ecedfa-kubectl-1.15.2-0.x86_64.rpm
Packages/bfef0ebaf6721dd74cafd3303de20286d6fd78cf4fdc53cd2f348681e434a493-kubectl-1.15.3-0.x86_64.rpm
Packages/de6039d16a6e77e0f38ce47cfaff9d450545757b3d09d34a10daf5667cd95ef6-kubectl-1.15.4-0.x86_64.rpm
Packages/90dbd280f7fa38882e359cb83ca415f8d3596d9e4ff3e8e8fc11013042a0c192-kubectl-1.15.5-0.x86_64.rpm
Packages/5181c2b7eee876b8ce205f0eca87db2b3d00ffd46d541882620cb05b738d7a80-kubectl-1.15.6-0.x86_64.rpm
Packages/8922dc33ce7258b1dc00786108640d17739523794ced9cee27f12898bfeb6f80-kubectl-1.15.7-0.x86_64.rpm
Packages/abc03613f3358722b6488cf145fc4a2f63f4b700a8e71e175bad93472fb82a19-kubectl-1.15.8-0.x86_64.rpm
Packages/08ced8941ffc8be01091ef602d034f9c8adb2aad9650a545586ce021eb04939a-kubectl-1.15.9-0.x86_64.rpm
Packages/be6a085d0285f83d74429540b878a137a9341ba34dbd85dfbae997eedcb046fe-kubectl-1.16.0-0.x86_64.rpm
Packages/9c569e36042bf917f770648feb2c515c1ef90cee2de2d599652166593373150c-kubectl-1.16.1-0.x86_64.rpm
Packages/26d3e29e517cb0fd27fca12c02bd75ffa306bc5ce78c587d83a0242ba20588f0-kubectl-1.16.2-0.x86_64.rpm
Packages/fd6465355a85b8ddbc0b2e7cb073e3a40160c7c359576b86e9b8eca0a2d7805b-kubectl-1.16.3-0.x86_64.rpm
Packages/70b9657194f6d201244df9e111e9270d162241c014a196125e0c399ad45613e7-kubectl-1.16.4-0.x86_64.rpm
Packages/58b74cf8ceade30204ccdbda85887626810a345ea60be011c0dea476f62d88bb-kubectl-1.16.5-0.x86_64.rpm
Packages/0bfd3f23e53d4663860cd89b9304fba5a7853d7b02a8985e4d3c240d10bf6587-kubectl-1.16.6-0.x86_64.rpm
Packages/bf67b612b185159556555b03e1e3a1ac5b10096afe48e4a7b7f5f9c4542238eb-kubectl-1.17.0-0.x86_64.rpm
Packages/120c3cc5366afa7767e6a627578699244c514bc84d4c5727b8afc09a547b5ea6-kubectl-1.17.1-0.x86_64.rpm
Packages/b44630896c69cd411db53be1d5cb5ae899a40aba7c0766317ea904390fcfc45b-kubectl-1.17.2-0.x86_64.rpm
Packages/2e63c1f9436c6413a4ea0f45145b97c6dbf55e9bb2a251adc38db1292bbd6ed1-kubelet-1.5.4-0.x86_64.rpm
Packages/454dc2629a55437b2297742eadad6429117afd3c48965e7ef6bd830cd71b4bf6-kubelet-1.5.4-1.x86_64.rpm
Packages/af8567f1ba6f8dc1d43b60702d45c02aca88607b0e721d76897e70f6a6e53115-kubelet-1.6.0-0.x86_64.rpm
Packages/756d35ea4c7e7fde7343154bf073f4056521c8bdf4d9562ce5c325e9bcf79bcf-kubelet-1.6.0-1.x86_64.rpm
Packages/cde0b9092d421800f7207df677d6a1f321c82179e89a24e4b3c009a875e62c91-kubelet-1.6.1-0.x86_64.rpm
Packages/0e7171c32ec6c65100eb284f8ebc3ed06bfa62a3e2ac752994da93d2261d8de9-kubelet-1.6.1-1.x86_64.rpm
Packages/a8a2b37431da90798331a67b8b452572a72173414b1543368786e356f23bc4ce-kubelet-1.6.2-0.x86_64.rpm
Packages/9e4c4e215f62ccd63ac6327b1d64faa9a66dbb24975f781a9c9c9a0a5c1e441b-kubelet-1.6.2-1.x86_64.rpm
Packages/b610c239bd57c3ca0dcae3f4d4ae60a0a1eab99c7b059cf1b0fe8dd7267c3f73-kubelet-1.6.3-0.x86_64.rpm
Packages/02f43f121f2fdff8217e62500bc2ef060650795e87c61000cd9a70842d668fb1-kubelet-1.6.3-1.x86_64.rpm
Packages/08706c8f6fbce961cd8994043cd27204f2bd3b47afd54cbd8755dae4b6cc4279-kubelet-1.6.4-0.x86_64.rpm
Packages/c90b57c8768e5d9992afcb5543bb20bdaeba5d0fcc581b34a00220fcf54fbb1e-kubelet-1.6.4-1.x86_64.rpm
Packages/7d43767d519f9f76605408517c999631e3835621afa8f5e80b9b8fd0925842ca-kubelet-1.6.5-0.x86_64.rpm
Packages/475f4d3ace5d98b41269c7734d27be722deb97862196439f9d764211463e0ccb-kubelet-1.6.5-1.x86_64.rpm
Packages/f2c5f556143a820ed5bbb6a0ac7601f68dea8b28c8edc5db0a0d5d0ad4e94fd0-kubelet-1.6.6-0.x86_64.rpm
Packages/2c3367bcc8239a3c681569a95f7837c66d6eabc6a534cd4bf9689a8fe8cb5243-kubelet-1.6.6-1.x86_64.rpm
Packages/b58a3f13d494458fbe23dbf22abc0213dc2c9ffb1e30f76dad3d7321d0715444-kubelet-1.6.7-0.x86_64.rpm
Packages/882d43aa9fb8414d5e1eef331f4f5bdbb73d5934938250f42d1f98a21506f5fd-kubelet-1.6.7-1.x86_64.rpm
Packages/441742abb2a3725d39a38b418f7bb63fcc24181a55fd350f2bc47521df3a528e-kubelet-1.6.8-0.x86_64.rpm
Packages/233c075a115b32eb8b53212a0b8e21414c944cc0d701d8c6ef334084105bc1be-kubelet-1.6.8-1.x86_64.rpm
Packages/d6b67999f4591bc4b3eaec94446458568d49345338c281e429b32e7b41c2c9ae-kubelet-1.6.9-0.x86_64.rpm
Packages/697f565212c13f7d0312b52a89b154ec66670385f4f0db1685704c75cdfd3836-kubelet-1.6.9-1.x86_64.rpm
Packages/8695191ce64e8742766e44fd08beb61a640522db81cd1a628e2a578c8cc5df0a-kubelet-1.6.10-0.x86_64.rpm
Packages/0554e280fb1cab3a572a0a4b22817dc01fb4d206d7022e39a29761903579abdd-kubelet-1.6.10-1.x86_64.rpm
Packages/254343aac568684ff6bdf75e69c015127b45cd8a6e7f9b0eff335b0a082b3117-kubelet-1.6.11-0.x86_64.rpm
Packages/31d0505434a4a9856f594daa6e478eab83c815f9d1b570d02ce25f83d9649f09-kubelet-1.6.11-1.x86_64.rpm
Packages/f469e89265a9a215358f6d2e541c0d5dd4f6b522fcbbe401117ce656b3aba931-kubelet-1.6.12-0.x86_64.rpm
Packages/3d7054b9bbb673a8dcfab836dc9fed4791a398199a7c08778cd025f4a1dd4df9-kubelet-1.6.12-1.x86_64.rpm
Packages/01fad4462c4e6a1b27be8aec6a5ee730df26ea7ed9e9e965211eaf031e022df0-kubelet-1.6.13-0.x86_64.rpm
Packages/14cef290532e68668924850f5f7002ffea058113917399265d8beb6648641cd9-kubelet-1.6.13-1.x86_64.rpm
Packages/8f94c56214c25f72f999f8a9c099275c2e1b99110a155679c1eda293192c26a7-kubelet-1.7.0-0.x86_64.rpm
Packages/6b748a98eaf71f33d3b37b911c8ba77650b25ff944a66743e703aa4bf1af0918-kubelet-1.7.0-1.x86_64.rpm
Packages/30136924ea242fee92df3aaf0297365dc58e8d1260b7bde66c94bda5a42a6f04-kubelet-1.7.1-0.x86_64.rpm
Packages/05be0e2d097c68b6355de205fe11064d6593593058d93fbe8031d7cc7fa74f5c-kubelet-1.7.1-1.x86_64.rpm
Packages/1e508e26f2b02971a7ff5f034b48a6077d613e0b222e0ec973351117b4ff45ea-kubelet-1.7.2-0.x86_64.rpm
Packages/3334e8cb551b23cbccd2e01d09c1647b83a247579b0e37c6bde2ce4fbd54edba-kubelet-1.7.2-1.x86_64.rpm
Packages/28b76e6e1c2ec397a9b6111045316a0943da73dd5602ee8e53752cdca62409e6-kubelet-1.7.3-1.x86_64.rpm
Packages/74106830515df731030ad7b3eac96e5cc44470dd62e8724fb91bb016f0dfc102-kubelet-1.7.3-2.x86_64.rpm
Packages/4f60c17a926175fb9abcfdd487cebafbbbce0e2248d2b99c189ae0877376b88d-kubelet-1.7.4-0.x86_64.rpm
Packages/737d15479c1dec0037df731db6ebbc7339c5205c24a9ce0fe669bccf18fe2c9b-kubelet-1.7.4-1.x86_64.rpm
Packages/c87d884d28952332d952d5f7f294ead1838a8e1219d2fc75923a5be06903afaf-kubelet-1.7.5-0.x86_64.rpm
Packages/e50c6e3313325838d5d1c27b48763cf2e5c53f1bade0502611c15655e86370e4-kubelet-1.7.5-1.x86_64.rpm
Packages/f049d9a0a9172b00aa2725c86a0de6f4ee51571105344b31b0b2523be9fda635-kubelet-1.7.6-1.x86_64.rpm
Packages/71b0d0d00f423a1cc7499d1e1544de54319653d1f2118bb96d35b3eb732e24e7-kubelet-1.7.6-2.x86_64.rpm
Packages/45bd874fbf333c3e2595da358ce7ad3220af46b1ff69d69f57f882dee0db52a8-kubelet-1.7.7-1.x86_64.rpm
Packages/04dc3c71a396b2c166fffc5dba87df260919f6905774e314e3405b11019f3dd3-kubelet-1.7.7-2.x86_64.rpm
Packages/305fd1a89e9da295f56ac2df41dd73e521e29c768c57074a365096d713bfe928-kubelet-1.7.8-1.x86_64.rpm
Packages/11bff1343306ea07d59022267f27b9008bd0101bd738bc49d6957ed7be68f025-kubelet-1.7.8-2.x86_64.rpm
Packages/c5f3b22a90539fc5880728f80a4faabe92d0101ccd333ba864ea9898dce754d2-kubelet-1.7.9-0.x86_64.rpm
Packages/68b9b7b4921a3b1287cc518b51b040f8c6c0586f2175c186b3e52d36402bbd83-kubelet-1.7.9-1.x86_64.rpm
Packages/b7127de06c5bec3e197d279ac910b7a119fa2f799fb18ca3ec11499506396957-kubelet-1.7.10-0.x86_64.rpm
Packages/a68572067e8ed51183782689c9e98da8d4753f86305c45a1fdce7d202b8b66e5-kubelet-1.7.10-1.x86_64.rpm
Packages/57fc3fb190b0f538f1f6b109f0b23f3456bc48aae3e2ceac5041c6438aeb6a50-kubelet-1.7.11-0.x86_64.rpm
Packages/02180e8b3266343d1f3ddebda362f6e199f1149528ab4b51479a5ff082de0739-kubelet-1.7.11-1.x86_64.rpm
Packages/5c0e31c65e992e0f20cdb3b783d5f00293f0690cafa045a97f4af86cd861c39f-kubelet-1.7.14-0.x86_64.rpm
Packages/8ebb830fda83dc4d4d0d8703818d997174c2be7b2f370da7aac768f0e46eaab6-kubelet-1.7.15-0.x86_64.rpm
Packages/5310f9f1b7961b4e5b11a5ae3d692d0526e24f7e68bbacc286e5ce150fdb8a13-kubelet-1.7.16-0.x86_64.rpm
Packages/d51b547da5980e2f6bfd33a382779194a0eacecac999ca726f85ef76183a9033-kubelet-1.8.0-0.x86_64.rpm
Packages/9699684a155be5500822a37a01bc76e2ad5605b8541262365488b4530665a4e7-kubelet-1.8.0-1.x86_64.rpm
Packages/a35571037b554243d386436ff729c90a3cb270f5797b7cd254ef0afbd4e706bf-kubelet-1.8.1-0.x86_64.rpm
Packages/c8e037e025d56a6fba51fd30c9a4001d442732bb5b122a476aa9ab98c9ebbb68-kubelet-1.8.1-1.x86_64.rpm
Packages/7d976686554e598267d8c5eb030f2a8f4e575f47015ecf94459913b80c9e5bb8-kubelet-1.8.2-0.x86_64.rpm
Packages/5a7fa12561d34b95da6c1913086316a5f08169138a68a4d3874b76a64fd95114-kubelet-1.8.2-1.x86_64.rpm
Packages/a53acfe63a475bf61661036c12890217f4921a6d6d6c3e6ecb4c598fc11cac19-kubelet-1.8.3-0.x86_64.rpm
Packages/9de65bc04416de79a9c045f0191b1dda187151f9b1c665aa03b66f482092cde0-kubelet-1.8.3-1.x86_64.rpm
Packages/1acca81eb5cf99453f30466876ff03146112b7f12c625cb48f12508684e02665-kubelet-1.8.4-0.x86_64.rpm
Packages/d353208063910d0d9a8ff9a3e9bfb0cbee43be501bc4ca16d69223560c4a5cc0-kubelet-1.8.4-1.x86_64.rpm
Packages/7330c09dd7dccc300c9c5f696fc4a8a76327e5068e51a60e4517b59a1c4b3dbc-kubelet-1.8.5-0.x86_64.rpm
Packages/26944be0d909a90d8ef7559b17a3ce4381821d623314a35eb0c74e62d055d944-kubelet-1.8.5-1.x86_64.rpm
Packages/f4011419193577161ae891a1b26986cae5f5e588941340db3abb771e2a677de7-kubelet-1.8.6-0.x86_64.rpm
Packages/e2792eaf4f3088dd639feea4b4bbfb0b2f96ac498d7a67c31b47488758ff7633-kubelet-1.8.7-0.x86_64.rpm
Packages/0e9672058439598e1d6d57ae5bd36007730225f94139cc0f03ddc81ddd798787-kubelet-1.8.8-0.x86_64.rpm
Packages/6330c00f0fcefa1ad4e5541abe6ff2682668ec1c34b9116a26c54191203c4d24-kubelet-1.8.9-0.x86_64.rpm
Packages/bb3cbbc83f073688b4d4b005f1608cbf4b2d81d209e9a7bd16e825678a49d048-kubelet-1.8.10-0.x86_64.rpm
Packages/16c56093ac1b63f9e03ddd31c763ea6779754799add991e1468db9e27f04be83-kubelet-1.8.11-0.x86_64.rpm
Packages/9fbc91e22fb4c35e8d1f7cdc2ee10938153cc2cba66d647c9448ff373494b213-kubelet-1.8.12-0.x86_64.rpm
Packages/cc2d0955725febc15f54f2696197ac0d0d0a42af6bea61266f9da0327bbd22e8-kubelet-1.8.13-0.x86_64.rpm
Packages/d4b843350544a966650eecfa941e29c886fee3cfc4303f8180e6ca3d67cb7558-kubelet-1.8.14-0.x86_64.rpm
Packages/affa4694a663dc9c94c88f40e8c4dd704d220f72d38274482840307145fa868b-kubelet-1.8.15-0.x86_64.rpm
Packages/8f507de9e1cc26e5b0043e334e26d62001c171d8e54d839128e9bade25ecda95-kubelet-1.9.0-0.x86_64.rpm
Packages/cec192f6a1a3a90321f0458d336dd56ccbe78f2a47b33bfd6e8fd78151fa3326-kubelet-1.9.1-0.x86_64.rpm
Packages/0e1c33997496242d47dd85fffa18e364ea000dc423afdb65bb91f8a53ff98a6f-kubelet-1.9.2-0.x86_64.rpm
Packages/b59d38d81913c58b9b382e4c39a00303d68ece464f5402046f0900e70ea106f8-kubelet-1.9.3-0.x86_64.rpm
Packages/d69020c41cbd116872424bef4b6d6c5c7b29cbcc8fca3ca949aca421899c27fd-kubelet-1.9.4-0.x86_64.rpm
Packages/aea4066130f6381c7145630edddd65df6293728b2fef3370e040920367f8e33d-kubelet-1.9.5-0.x86_64.rpm
Packages/fff4e7133c41ce6aaca95adb598f47967d180c4c5e8e6c67d8bfe58345bfde27-kubelet-1.9.6-0.x86_64.rpm
Packages/4a7b197df9b4f478f6e26863e450811c89c28126993eb0cc777fb39b819cedc9-kubelet-1.9.7-0.x86_64.rpm
Packages/0ec1a0934164c6ca5a943f9f94b6243c7e9f951845c973d5e7272c4efe224a61-kubelet-1.9.8-0.x86_64.rpm
Packages/6d60a69c40d3ff19fc8ac5b42dbc2f9affd83d4652723189bf93f5395b9dc1e7-kubelet-1.9.9-0.x86_64.rpm
Packages/389f8c674531635f3851697aef1fdd892db5ec8cffb5725b840d7185b4886b90-kubelet-1.9.10-0.x86_64.rpm
Packages/8433088087912f0df2101cf71db075c327b72e292c83788ccb1a8139d5db4f3c-kubelet-1.9.11-0.x86_64.rpm
Packages/5844c6be68e95a741f1ad403e5d4f6962044be060bc6df9614c2547fdbf91ae5-kubelet-1.10.0-0.x86_64.rpm
Packages/4cad7573e2617d903a12f5d318f597ec47bc0e67792a04f67ffee3e06c0ad373-kubelet-1.10.1-0.x86_64.rpm
Packages/bdee083331998c4631bf6653454c584fb796944fe97271906acbaacbf340e1d5-kubelet-1.10.2-0.x86_64.rpm
Packages/07b0095a6a3400334a66719a92161dcc4c8f73fb3652b47cbffe02ee5537936a-kubelet-1.10.3-0.x86_64.rpm
Packages/2d41bb10c7142326719c9052e9f0301f0875c3f25761fa4740ba862eaae0888e-kubelet-1.10.4-0.x86_64.rpm
Packages/94d062f2d86b8f4f55f4d23a3610af25931da9168b7f651967c269273955a5a2-kubelet-1.10.5-0.x86_64.rpm
Packages/1f90ed0086a1403b91d1b5d3b0d4e1b4a4506755bdff4f6b600d7be80f38f6a7-kubelet-1.10.6-0.x86_64.rpm
Packages/703df0b4a0909147f46a26f179add02065191da20a49f2e75074179eea6bcf00-kubelet-1.10.7-0.x86_64.rpm
Packages/8868d4fb999d79bce430333c5267f6c490834aa1ab87e0d7f386377189abe0e9-kubelet-1.10.8-0.x86_64.rpm
Packages/ef22149d1225058f14d073c9f1d7f05c36d6c499ba7244fd50c897ed8f1e5b13-kubelet-1.10.9-0.x86_64.rpm
Packages/aa12eec663bd19a74af7bd2980e34a591f317fe943dc52632c04359112613c0f-kubelet-1.10.10-0.x86_64.rpm
Packages/fd9075e0e1dcceb43de6d75859d703c7be4e58353e9465c57fc610d7c4e54e91-kubelet-1.10.11-0.x86_64.rpm
Packages/2f6f9dbc0cefa96e1282ec67151f4975781be2cc95c0575b8e8a3fd7ba2e19ad-kubelet-1.10.12-0.x86_64.rpm
Packages/7b77927bb6aeecca5a16db120ec131820d8dba6f468cfde31d98263032d93c5c-kubelet-1.10.13-0.x86_64.rpm
Packages/d19addf9b24468fcd6bb09ffa509eb00259659018d7381867c4249b493dfdba8-kubelet-1.11.0-0.x86_64.rpm
Packages/ecaed8cbeff649711591e180d286b37d4b71bc6539ef0b598004e0bef7260ed7-kubelet-1.11.1-0.x86_64.rpm
Packages/518f6ba7535c69a4fa3fc7b40e370c4c4821ba9231fc25dafa556f5b4554b001-kubelet-1.11.2-0.x86_64.rpm
Packages/f309f5183598b920355bc18ff53f3816f55c27ac1c6e8a4369344141d9d4f127-kubelet-1.11.3-0.x86_64.rpm
Packages/aeee56adca1c0722f5cb67eff7a99462005032554131006ceaf9b75313af9707-kubelet-1.11.4-0.x86_64.rpm
Packages/e0b14f554af69ce3bb32bf07b94d0c8a19c63837e8f17932c8e09e86276af9d0-kubelet-1.11.5-0.x86_64.rpm
Packages/8057c5fe56cc8dcdfd4a056578a29b2932e13e9fee79fee3e8b66ce8ac61892a-kubelet-1.11.6-0.x86_64.rpm
Packages/c7db7da05f052a332aa9eb30fe1fed9b09d4151ccafbed450a9ea4d730f4d6b5-kubelet-1.11.7-0.x86_64.rpm
Packages/2fc5e2d281761814436bad836eecec92f1c522301e7c6eaeb5e1497673755b3f-kubelet-1.11.8-0.x86_64.rpm
Packages/b82b358f2951dc396af6e2e56f9b1ca6ab5b494208996a5788682d9e49312d7b-kubelet-1.11.9-0.x86_64.rpm
Packages/8c80b05ea419461398ae0fc9faaad0a4215a6313c4a7804f08fb25393c9aed49-kubelet-1.11.10-0.x86_64.rpm
Packages/d8a3f7394ecf1cba6d7ca2c13775a0b64ba927fcc7b4d741b03a9d3c23efd484-kubelet-1.12.0-0.x86_64.rpm
Packages/c4ebaa2e1ce38cda719cbe51274c4871b7ccb30371870525a217f6a430e60e3a-kubelet-1.12.1-0.x86_64.rpm
Packages/c2c59cf03fa14b8d91026c94d63c877091d54ae8492ee3f49faa7aeccd5cac3f-kubelet-1.12.2-0.x86_64.rpm
Packages/154b25bec8341d22baa95c504c20bdebcee73162715adf92198c3f316c6e96b4-kubelet-1.12.3-0.x86_64.rpm
Packages/10ab933fd7be536c378f623ccf154a2ce0e9dc32c9776686f76ac836d9d47267-kubelet-1.12.4-0.x86_64.rpm
Packages/60c6045ab59dc3ffb9f28bb60aa30e1dffca14dc8044fadd15afb82b8a496273-kubelet-1.12.5-0.x86_64.rpm
Packages/186e510a6ee10ddea8f01ba8f51231bbdfc0006bea27e484248c7fdb633e0507-kubelet-1.12.6-0.x86_64.rpm
Packages/583cdf9a7d7fe76b4b05ac00ff57d902b4c276fe98b3cbe0e696afcbdfd0e456-kubelet-1.12.7-0.x86_64.rpm
Packages/ba4ec631df6ab11e169fd7747020dfbee7d0388477a9ddcd8b1a4576248451cb-kubelet-1.12.8-0.x86_64.rpm
Packages/9be86d0b5cba4464d0c4094ba414613776a0813673a5eea3dd6cfbfce4946b8f-kubelet-1.12.9-0.x86_64.rpm
Packages/8cfb58f4ca56a955a65d28a99f7e0e4178fc0ac889936ce1cb6e53eda58437ed-kubelet-1.12.10-0.x86_64.rpm
Packages/26d0db769cb8f64b417c3bbdcb24b7214f63d211fb889de55fe2fecc440e8fa7-kubelet-1.13.0-0.x86_64.rpm
Packages/25cd948f63fea40e81e43fbe2e5b635227cc5bbda6d5e15d42ab52decf09a5ac-kubelet-1.13.1-0.x86_64.rpm
Packages/bb7224bef9afe2684b74e4d4494f6c68c79b3a3cce30faf76025bc1a9f31c140-kubelet-1.13.2-0.x86_64.rpm
Packages/b5f5e57f3ef573de50ea00fd14447d6f0911c283a5fc685236ed744123ebb1c4-kubelet-1.13.3-0.x86_64.rpm
Packages/344d06740c6e21a13712258f16654eb1695ccab66a4149ecd2d383024bbfe5a6-kubelet-1.13.4-0.x86_64.rpm
Packages/bbd7f31400ca115134068ae2d77fd0ebd33652b42db4a45e7e3ad329d8cef92a-kubelet-1.13.5-0.x86_64.rpm
Packages/4d547dcac2ede40184462367cf513079704b377ca6290d2b115db71863a1dd56-kubelet-1.13.6-0.x86_64.rpm
Packages/3b3d42dd4695a1ae1a34d8e0910e66bdb89927324f88871236099d5dcb29101c-kubelet-1.13.7-0.x86_64.rpm
Packages/d81081f5c034fd076ee20b6a73a5a2c657f0db2c9df7045dde8a5893f2e1d2a6-kubelet-1.13.8-0.x86_64.rpm
Packages/36fc1c8d69199a30cce2d2d2081fcb5a3d6a2d6c51b87bbe7e41bed329c0ba76-kubelet-1.13.9-0.x86_64.rpm
Packages/1eaad53882a5124d9eccd267b56a036149175ab9176a8c6e9f1df21ee663254b-kubelet-1.13.10-0.x86_64.rpm
Packages/466796027abd41f27e5299bfd25bf38314b79fd0089aa3b75a17acc465cac116-kubelet-1.13.11-0.x86_64.rpm
Packages/fe8b2abe201059a84229ba14798d41a82fd547dfe02eea65e19444a571ecce36-kubelet-1.13.12-0.x86_64.rpm
Packages/6089961a11403e579c547532462e16b1bb1f97ec539e4671c4c15f377c427c18-kubelet-1.14.0-0.x86_64.rpm
Packages/e1e8f430609698d7ec87642179ab57605925cb9aa48d406da97dedfb629bebf2-kubelet-1.14.1-0.x86_64.rpm
Packages/1a181064b472261b78b534b5a233a4c73d505673c02acbe01d95db819940006e-kubelet-1.14.2-0.x86_64.rpm
Packages/77f40346fedfa1917524bd3bcf9f7cdfb0ec234b965609b451d3c8ea7ffc6a92-kubelet-1.14.3-0.x86_64.rpm
Packages/219048ddb1cad26549ab59ec9e89b91fc670dce5895a2299714e23d217346c75-kubelet-1.14.4-0.x86_64.rpm
Packages/82846e5d62bfa57c20fb5c596bc7b6648e1e21bdb93ab57cdbb5b1b75eb8afff-kubelet-1.14.5-0.x86_64.rpm
Packages/d1aa7bd03ad0054a15c536bc7dc0016249ce37844202e57916be43abd84c2bf9-kubelet-1.14.6-0.x86_64.rpm
Packages/8521499cb634d24ba9720885b86a10fcedac5bd7cecc6a34fa266152e195f0a9-kubelet-1.14.7-0.x86_64.rpm
Packages/651677e32820964b6d269b50f519a369b1efb13bd143f64324a3c61b1179b34a-kubelet-1.14.8-0.x86_64.rpm
Packages/14e2a7c0f19b5f4700033a16bf0557c13c585d794001636afc044cf1355f986a-kubelet-1.14.9-0.x86_64.rpm
Packages/7081b7854fee4f93b44d3fc523fc946ceb4f69b29bedcdc1a4194638acca7598-kubelet-1.14.10-0.x86_64.rpm
Packages/557c2f4e11a3ab262c72a52d240f2f440c63f539911ff5e05237904893fc36bb-kubelet-1.15.0-0.x86_64.rpm
Packages/f5edc025972c2d092ac41b05877c89b50cedaa7177978d9e5e49b5a2979dbc85-kubelet-1.15.1-0.x86_64.rpm
Packages/ade2907a0bab0715751dc037a4d0582bf9d933d3f15b588876b48465d8d9169b-kubelet-1.15.2-0.x86_64.rpm
Packages/a869eccfd0558795d1e001250e1128eb3507905ca5cbf98e5dbfa58088666a74-kubelet-1.15.3-0.x86_64.rpm
Packages/dd6f87ffb5e04121b39c9b8301573225aff0a135ebf73046dcd3e21c4ab7a6cb-kubelet-1.15.4-0.x86_64.rpm
Packages/4e3e9edc797eed91c0d5ab63b3dd464e82d877e355cae5f35e8f31c9e203658a-kubelet-1.15.5-0.x86_64.rpm
Packages/e9e7cc53edd19d0ceb654d1bde95ec79f89d26de91d33af425ffe8464582b36e-kubelet-1.15.6-0.x86_64.rpm
Packages/b3220d8f199d4e5fc1b12433db3ea3766ce6e5b827c352c95efcd3824125fa20-kubelet-1.15.7-0.x86_64.rpm
Packages/757e11c379aa2dde8bc617e98f3e80cbe3e6fe1bb3b5a51c49fa1c7a5f86becc-kubelet-1.15.8-0.x86_64.rpm
Packages/d4a105272b51e260203dbe22836d947d942b8a0991b6025a4a590eafbd70b59a-kubelet-1.15.9-0.x86_64.rpm
Packages/d58731a80db4214f3f1021f5676fe7f4da7239fc4a48c60f1541a5c904769313-kubelet-1.16.0-0.x86_64.rpm
Packages/5c2055ea4f698f2cfe8ee777e3fdb27be95161175cea21cd914bd7d28c299e09-kubelet-1.16.1-0.x86_64.rpm
Packages/0939db1dc940fa6800429c7ebef9d51fd9d46ff540589817cdb1927b8fae7aaa-kubelet-1.16.2-0.x86_64.rpm
Packages/8a0e2b605c7a616d7cb72c25c9058b2327e41d869046c7c6cb3930f10a3dc012-kubelet-1.16.3-0.x86_64.rpm
Packages/d13f20c41037ec4923cc02f0aa48a67061284f3f77c3a05c2d19df72b89bbc98-kubelet-1.16.4-0.x86_64.rpm
Packages/a419b3a5c4104e6f2ecb1e6ef202610db07e6bf0ed4d1f31e7121760718a52a7-kubelet-1.16.5-0.x86_64.rpm
Packages/6f0d57f3271c856b9790f6628d0fa2f2d51e5e5c33faf2d826f3fc07a1907cde-kubelet-1.16.6-0.x86_64.rpm
Packages/7d9e0a47eb6eaf5322bd45f05a2360a033c29845543a4e76821ba06becdca6fd-kubelet-1.17.0-0.x86_64.rpm
Packages/2c0df9a4c416331d9447d7e3da9fbe108e7598b2a0cfd02cdcd0429e35a0895c-kubelet-1.17.1-0.x86_64.rpm
Packages/3ee7f2dff78e6fbb3ac3af8acb1a907f4bec1b1ef4cf627cbe02fa553707f2e9-kubelet-1.17.2-0.x86_64.rpm
Packages/0c923b191423caccc65c550ef07ce61b97f991ad54785dab70dc07a5763f4222-kubernetes-cni-0.3.0.1-0.07a8a2.x86_64.rpm
Packages/e7a4403227dd24036f3b0615663a371c4e07a95be5fee53505e647fd8ae58aa6-kubernetes-cni-0.5.1-0.x86_64.rpm
Packages/79f9ba89dbe7000e7dfeda9b119f711bb626fe2c2d56abeb35141142cda00342-kubernetes-cni-0.5.1-1.x86_64.rpm
Packages/fe33057ffe95bfae65e2f269e1b05e99308853176e24a4d027bc082b471a07c0-kubernetes-cni-0.6.0-0.x86_64.rpm
Packages/548a0dcd865c16a50980420ddfa5fbccb8b59621179798e6dc905c9bf8af3b34-kubernetes-cni-0.7.5-0.x86_64.rpm
Packages/fd5f5da2d1a262fa687404d34e72813520274364557e648bc64a8136f1a02630-rkt-1.25.0-1.x86_64.rpm
Packages/7a382e59dc2c39a66083e03ec061f33771e4a7130c98cd0ef61492b2196c0378-rkt-1.26.0-1.x86_64.rpm
Packages/01b97b0ddb967d0ed9fd78327a784efbfea8cd0d9789f5bab8b9bbfe94477c60-rkt-1.27.0-1.x86_64.rpm
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
100 8881k  100 8881k    0     0  2823k      0  0:00:03  0:00:03 --:--:-- 2823k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 9630k  100 9630k    0     0  1396k      0  0:00:06  0:00:06 --:--:-- 1878k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 20.4M  100 20.4M    0     0  3719k      0  0:00:05  0:00:05 --:--:-- 4526k
```

For example, the latest is for V1.10.2
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ find -regex ".*1\.10\.2.*"  https0x3A0x2F0x2Fpackages.cloud.google.com/yum/pool
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

The *yum-Packages-curl_el_all.sh* contains all packages to avaiable download

## DEB

Ubuntu release cycle
+ https://ubuntu.com/about/release-cycle

Debian release cycle
+ https://wiki.debian.org/DebianReleases
+ https://www.debian.org/releases/
+ https://wiki.debian.org/LTS

### GPG

Download
```
fanhonglingdeMacBook-Pro:http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ ./apt-gpg-curl.sh 
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

Last updated (Dec 8 2019)
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ ./apt-dists-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  8993  100  8993    0     0   3444      0  0:00:02  0:00:02 --:--:--  3445
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  8490  100  8490    0     0   7335      0  0:00:01  0:00:01 --:--:--  7331
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   454  100   454    0     0    459      0 --:--:-- --:--:-- --:--:--   459
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  260k  100  260k    0     0   185k      0  0:00:01  0:00:01 --:--:--  185k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 31290  100 31290    0     0  27009      0  0:00:01  0:00:01 --:--:-- 27020
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   228  100   228    0     0    214      0  0:00:01  0:00:01 --:--:--   215
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:--  0:00:01 --:--:--     0
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    23  100    23    0     0     23      0  0:00:01 --:--:--  0:00:01    23
```

### Pool

List packages
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ egrep '^Filename: pool/\S+deb' https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-xenial/main/binary-amd64/Packages | awk '{print $2}'
pool/cri-tools_1.0.0-beta.1-00_amd64_68880f674e9bf5959744d2cde7389a1a57b857e3fb769051ae0c506b19269ff0.deb
pool/cri-tools_1.11.0-00_amd64_768e5551f9badfde12b10c42c88afb45c412c1bf307a5985a4b29f4499d341bd.deb
pool/cri-tools_1.11.1-00_amd64_e6300f0f4ec2fb5d6967533416580e1a22be90277944370ceb2551b04d7bc1a3.deb
pool/cri-tools_1.12.0-00_amd64_2d9f048a50a9dfeceebd84635f1322955aca6381d9c05b4d60b3da1edb7d856c.deb
pool/cri-tools_1.13.0-00_amd64_6930e446a683884314deef354fbd8a7c5fc2be5c69c58903ad83b69b42529da4.deb
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
pool/kubeadm_1.8.14-00_amd64_8d7826f7c47edd4e1b2b535aa839ff9294ddbdbbc26aaae6439d3ebe1fcd7de5.deb
pool/kubeadm_1.8.15-00_amd64_9e4387b817f841a0a5ce067484e8464482336c43c14d673883791d2b73dd1401.deb
pool/kubeadm_1.9.0-00_amd64_191bd1d8a63d263cdb318f77b03fad6c8387a912cf16a21b9b24d7e9108b4911.deb
pool/kubeadm_1.9.1-00_amd64_31ec5709ec1eb37b194563cb28eaafafd1f0c9d009239cdfb694686066b35c2e.deb
pool/kubeadm_1.9.2-00_amd64_869c826d406035aad1617118f88d5c627e16797388a5e2b3735d7b825d62e7ad.deb
pool/kubeadm_1.9.3-00_amd64_df3d3c15601640c5cd5db3902d52fb2e33263b79c62d299700971083b7bec9f5.deb
pool/kubeadm_1.9.4-00_amd64_7c267288a1a9cebc480778b50ef00ab5dfe477f545bc917a50adb322ea2b9c24.deb
pool/kubeadm_1.9.5-00_amd64_3d2a1245ecfd2ef57d1d6f0f3897658d5e506ed467d8197f0669be445e9db259.deb
pool/kubeadm_1.9.6-00_amd64_9763ef588d8820e343c180f1e2c16d1c51be4d6a686f2d5ebd87a64d8b4fd661.deb
pool/kubeadm_1.9.7-00_amd64_af2c1e96e8339736360470580f6bb642f5b8702bf969c5acdeeb23626198d11d.deb
pool/kubeadm_1.9.8-00_amd64_298b1c2e4d85c758eb2e7e1a0d0e8fc3d126e76eeae3ed01bda68ed7f23b3708.deb
pool/kubeadm_1.9.9-00_amd64_507be0d47033a48879fef0b1af2667a75f41f6aa160b178f162265c451e162d1.deb
pool/kubeadm_1.9.10-00_amd64_dc408ee1cf892ff364e9a2d358cc320aa8b72ba9d55c7a2c89ef3e637089591c.deb
pool/kubeadm_1.9.11-00_amd64_5e44c0f0a14ca244230304201336b64f47022bed21c464d2df5647a7d37b4971.deb
pool/kubeadm_1.10.0-00_amd64_ea6b408af5de27ae3df9a1f96724067104998608b64869fee2c680ea9f9c000d.deb
pool/kubeadm_1.10.1-00_amd64_d5942f2d1a7b85136b82dc79c6f6204373c731b6ef7fe204a96f1c56af4859f3.deb
pool/kubeadm_1.10.2-00_amd64_4bdf79bcda2a820210f5c0f7e84f9f57d6cc196ba9b59943f296544624e6d743.deb
pool/kubeadm_1.10.3-00_amd64_094e7960fd50b98a2a34abd9807db7dea4cc25d7c9f0ecb45d6968ccf99b54b4.deb
pool/kubeadm_1.10.4-00_amd64_3b879e35929f633d45dc47081d615ef763e1f46b090330f54f57b0c7a67b8a52.deb
pool/kubeadm_1.10.5-00_amd64_599778e9ebcdf6340b3068dd19bbf3674a9d7fe37d47455879c259b2e780b62f.deb
pool/kubeadm_1.10.6-00_amd64_ac88dba881a1928edc8972b2081c32e5a34096a69398aab11953e18c8dc5e6f8.deb
pool/kubeadm_1.10.7-00_amd64_fa8e085bffbf1f1a17ccae13759469a2e70ba0258c8f46b4c55e6367f289fbf7.deb
pool/kubeadm_1.10.8-00_amd64_9353dba1fefc8062661695ed58fbac2e0826a2dd6a027feae84d12a164fdb5eb.deb
pool/kubeadm_1.10.9-00_amd64_05bf1f5ad996b79b5cc1a6611b8cab43042dbd52e6f14bcefab2a3d6d4656196.deb
pool/kubeadm_1.10.10-00_amd64_9da7c1b7f14e32e1304d58e9427aa32cd21efa517370c00cb8253fea751b14ec.deb
pool/kubeadm_1.10.11-00_amd64_38c964b551f7827ea1c831ae97be9ee6cab7c6d66e4e01ae7eca0b2319a919fa.deb
pool/kubeadm_1.10.12-00_amd64_24befdcb526409eeb19967f268df416f29145682424932fea937f79dba92064a.deb
pool/kubeadm_1.10.13-00_amd64_4594e00a6f68d2c7f8c7990cf6408ca6f1582b9df05de22d3d68fd84cf05344c.deb
pool/kubeadm_1.11.0-00_amd64_81a49d243289b9d052175800e06cba49276588aa2bcc7f40273061994ca55793.deb
pool/kubeadm_1.11.1-00_amd64_46c19917292eebc540eb1e8e7b1bbe12a0d927c9eab6109f2d48876691eac39e.deb
pool/kubeadm_1.11.2-00_amd64_7602f5c4362b9c17aba83e8424830a98ca66074e36dead31d239f2beda91f1ff.deb
pool/kubeadm_1.11.3-00_amd64_872cfed42ad845cd058216f3c8ddcddb725926251e5ba1816a0f24d20c96572a.deb
pool/kubeadm_1.11.4-00_amd64_7c548a81640edf27c7ab0e7606495a32d0141b42729e575e98937f24aafca6fd.deb
pool/kubeadm_1.11.5-00_amd64_0901d56f40f4c9948116eb5333afc314ff4034f1f74e415163154d0786342198.deb
pool/kubeadm_1.11.6-00_amd64_2da945126bf1338a9055ac59e85531fd6a6b45a886dbccf439332b78851cf0ab.deb
pool/kubeadm_1.11.7-00_amd64_f28f01fbd17107809dca19fb7a03a4e3c930a61b37045ad123eccab009cb5b29.deb
pool/kubeadm_1.11.8-00_amd64_f462e6c6b137accb4364e8cd4d02e4e68c1973a5408c200c907ef423ea62eab7.deb
pool/kubeadm_1.11.9-00_amd64_90e71bfe9856b9ec3ceced0cb9ceff1331187233e8adae457507d5ffa92e193f.deb
pool/kubeadm_1.11.10-00_amd64_a518251af86868412223ce6b12d0212d5783278419bf7cf403b77cce1a31cfb9.deb
pool/kubeadm_1.12.0-00_amd64_db6058f8d51e8d4f06998ec6adbdf635c394ab9153b35a0b1872e8981e67154a.deb
pool/kubeadm_1.12.1-00_amd64_3caa0392444658082aa459418728abf81e440f439960963f54f08c6e6fdefbd3.deb
pool/kubeadm_1.12.2-00_amd64_777b8af5ad6ffb59d0ad4e6ea08061ae9714c3c86f11b31c945058d48ec98a44.deb
pool/kubeadm_1.12.3-00_amd64_1e5e9c6b1ad5de0ad9e7225a988a55fd9578f548750707825e2bd97659b626c9.deb
pool/kubeadm_1.12.4-00_amd64_d834e365c16fefdd069bae0f6c4f845528bd5bfa417688a0f09b75779d542e9d.deb
pool/kubeadm_1.12.5-00_amd64_187bcc23b5f65ba944ec6fce83d6d17f5017fd0f94f286ce1bfa8292dc3bc137.deb
pool/kubeadm_1.12.6-00_amd64_9e54189ca1061169060b4160dba125d2d33431e513c595e34875860972782bb5.deb
pool/kubeadm_1.12.7-00_amd64_b85ff1e7d24ebebacbbb6b35113a058f7aa27353d8247159b4e86a3a3b0f7d98.deb
pool/kubeadm_1.12.8-00_amd64_be31c04079363054e2e112a38a87074b8468edac182781d1d5d9477149131ecb.deb
pool/kubeadm_1.12.9-00_amd64_d37876641e54752e43bdf1b50a3b76d56dd245fb2040042bb211ef2923c35da6.deb
pool/kubeadm_1.12.10-00_amd64_600513b5bdc3ff15caab99660a4e0e8d5b7d31bc907ce63f8a4fcd18e0c9c6f6.deb
pool/kubeadm_1.13.0-00_amd64_c1f5d00d11a393159bcacbc9182922a0d6f5a16385d2f947c38b5200dc5a6f07.deb
pool/kubeadm_1.13.1-00_amd64_8fa37b4b7cf615e44ea2b16865449cf9106407996e9da1d01ae06422dc627b4c.deb
pool/kubeadm_1.13.2-00_amd64_1f6de98e354d2b7b8cd6c0c32ae56edd9ca694c64c9a25062d4baee735b78bec.deb
pool/kubeadm_1.13.3-00_amd64_afb1cffd17cfd3557e2ca9fe885fd75df110e011afc81278a5cbd623a4e49e38.deb
pool/kubeadm_1.13.4-00_amd64_1094a7c75b7764a1b43d1009e76532f5a30e940e5a8154ccb82ea0abf8792685.deb
pool/kubeadm_1.13.5-00_amd64_3a8177f070bce398f62a14a7ddbe81d064b378c085d7e2c4391571dfb87c259a.deb
pool/kubeadm_1.13.6-00_amd64_3db73983bbcfc47a9db490a140cf8da9861e10ec4a6cddbdd0f4dcae88cc6762.deb
pool/kubeadm_1.13.7-00_amd64_909d8a77b8693bbf669b0ade781fc46013a6ae7d9c794ef58cce7811c0f3acf9.deb
pool/kubeadm_1.13.8-00_amd64_88873dde740ebaa0092840ae78e88c30ccc27b078ed249fff085b50b24193e4c.deb
pool/kubeadm_1.13.9-00_amd64_a14b36feddfff15c4bb70babd53793dba0161e6bcc60bb1edcfccb28b810c78a.deb
pool/kubeadm_1.13.10-00_amd64_1a1545c3d60027dce7f53210764486404dc59467ee1c6ac6b01846125c5556e3.deb
pool/kubeadm_1.13.11-00_amd64_cc9795db439dbfa8920df6109abd9ff8bfb6cebd660f4fcaf0065e9b34cc6ad9.deb
pool/kubeadm_1.13.12-00_amd64_5791923bba94b850ca898fdfaa75916bdd3b4f5d86a50ea738d96f14230029b9.deb
pool/kubeadm_1.14.0-00_amd64_37d071fc4060e54bd626faae826e8ab2750a7cb6bf90d189cea52453677df80a.deb
pool/kubeadm_1.14.1-00_amd64_eb746d66f6ef69232e09a6800c570e1c86d9a242f3491843efc6f6c0e3803803.deb
pool/kubeadm_1.14.2-00_amd64_7258f899e6bd237d9fbdd029617ab2396606c7ecbd94ee4a6098e44861942415.deb
pool/kubeadm_1.14.3-00_amd64_bc84f915d52d0ce17b5a160b9145caaaf0e56b677a9af26db6aa3d07149a307c.deb
pool/kubeadm_1.14.4-00_amd64_364296f65fb4adc59cd310d14de3f59bd5c0792395643bc9a4ac444466e26196.deb
pool/kubeadm_1.14.5-00_amd64_10662653b2cbdad2a1f30de5eb1c78ccf9740121f04730d698241dd6dc4d45ce.deb
pool/kubeadm_1.14.6-00_amd64_d838f990ae337a3810902056c6e9c2d51d36b82baacf1bd9f6ed472aacb191c5.deb
pool/kubeadm_1.14.7-00_amd64_18cf2992172646c1e75b59b26938a6d965f87b60fb083cfd52763d087c2e4bc5.deb
pool/kubeadm_1.14.8-00_amd64_4ced76557ba80440cd6218591bb6d700415eb14c3259d7d2abb711165c20829a.deb
pool/kubeadm_1.14.9-00_amd64_5c8c4bba3ba94e1349acc2148b57844ffad756bbb20e2ed76e655e0a074ee244.deb
pool/kubeadm_1.14.10-00_amd64_5da105d01ccee9bc6dcf675b10a351b3d35b52d2b558937156c5dbbe28f59355.deb
pool/kubeadm_1.15.0-00_amd64_e016ea058d1cb8bc82b0d6795eb65f1fa5a637d9706b5c8967a7f8e71927c86d.deb
pool/kubeadm_1.15.1-00_amd64_5c678c8c18cbe3f18c83456e4e34b53db208203bb1d4aeeae68a8608e8b136b2.deb
pool/kubeadm_1.15.2-00_amd64_b795bc600ec750d4f07a152eda7b65c5d8b25d9781b41143a4f8bbf86e3a7de6.deb
pool/kubeadm_1.15.3-00_amd64_f1971bc935de7b82b8ddd6d28fbcfbdf34df74f4f14eb2f0e7d6419a4b6d5b00.deb
pool/kubeadm_1.15.4-00_amd64_9eecce619de47a2c19ff52d1ba595b48d736d2e8867a698f608d7a80f050b6d7.deb
pool/kubeadm_1.15.5-00_amd64_cffe0070e6279c8cdca599202eabeab1774b3265d0c590933d5e1115e739668b.deb
pool/kubeadm_1.15.6-00_amd64_2c1d026f59d6e012a8fc7658dddcde404f73c9e1159cdeb000aea0d4de3ed3b3.deb
pool/kubeadm_1.15.7-00_amd64_3bb705bc4f067a915507fbc9689fea3fd3f844a9a581e2830fad9a284ca866ac.deb
pool/kubeadm_1.15.8-00_amd64_94cb1f98765df28e0c072fe0d4a1b4487f39e2e51b2decb76ba17b3ab963f865.deb
pool/kubeadm_1.15.9-00_amd64_9926841da54b433b42fc999d6de239d8d2fa9ccd7672d973bfc5861bf57f37d6.deb
pool/kubeadm_1.16.0-00_amd64_b0fa26a7ac8cd90e9c3e388282828f320766264acc307b5bf45ffa79b5abed0c.deb
pool/kubeadm_1.16.1-00_amd64_857e08a99b0f02aed302881c69e5a0f68a8c412f7f44cb380905455e7fae487a.deb
pool/kubeadm_1.16.2-00_amd64_52c08f02d316c643d02cdef59e90caca2dffce9b7525eff98cd8e1ac21360bed.deb
pool/kubeadm_1.16.3-00_amd64_8ad50e8252c65f968e35b1a38240b0ee66ef3b72c5d498830788e123238aa1a6.deb
pool/kubeadm_1.16.4-00_amd64_f50fb707e1d94be390b31b4642ddf1dfc02707b5efb32e75a8eaf97bb085d832.deb
pool/kubeadm_1.16.5-00_amd64_fc6290b91af09da31679bf2ec2af287539a993bb50b48f59666a6add0aeb5270.deb
pool/kubeadm_1.16.6-00_amd64_a8145f49526c6e97abb5f75e5f49bd414b942cf4fc0f56f1f12315282085850f.deb
pool/kubeadm_1.17.0-00_amd64_86626f0c89bdb57f0536336f684b29a6c4c02129703ef6de529dd07f487135cc.deb
pool/kubeadm_1.17.1-00_amd64_dc58ea341070ce8643ebd1267d33757eb08ecfb590ea81028cfdd6d722b5621c.deb
pool/kubeadm_1.17.2-00_amd64_041e0e6cc71212fff2cdaadfa7fd52d992fad23efd15157d360df86a22cf7649.deb
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
pool/kubectl_1.8.14-00_amd64_42f4fd1ba6f389672c6a388aab0617d097a92eb41e578c4eadf59b2a755eeb1e.deb
pool/kubectl_1.8.15-00_amd64_28d815300cfca34fcfe509ce5afcfbb76ba939405182a01524ea851b01344ea3.deb
pool/kubectl_1.9.0-00_amd64_8ea712c18d89d56090c8753a9630d22fd8ae5cb03d4ec79a1fe6d262c8b4eb36.deb
pool/kubectl_1.9.1-00_amd64_bdfb1ad90e0f02a7ae614502079a87ed99209679bdedf0c62873564c186c9f99.deb
pool/kubectl_1.9.2-00_amd64_a9d5a970c989cb1f77c30c5d27f611cab58659751005e7052541540dca19855c.deb
pool/kubectl_1.9.3-00_amd64_bf9b3256304edcd2f4bd85b74857dacfe6796bec548e1411ed6f93350dd01316.deb
pool/kubectl_1.9.4-00_amd64_0f814c1c7d852379aa61744ed119518558a0f0cfa1a6e36e8fc2037bd023dad4.deb
pool/kubectl_1.9.5-00_amd64_fc4313ac6797482c4ece18c14ba2ad2c46bb6bed4b4e9e5c1ee30f64c31ff13d.deb
pool/kubectl_1.9.6-00_amd64_8f3805e25732ca6ecc12e2751932130a5131adb605691b85e06c6204a0b6326f.deb
pool/kubectl_1.9.7-00_amd64_c532c2abf3344a081d16fc69253c470f1e22c39350b6778dec2f73bb88eaa941.deb
pool/kubectl_1.9.8-00_amd64_46aefae6d33ea6e9ac78c858e150b4f92e3c22f831ce0f0614b405a277548da3.deb
pool/kubectl_1.9.9-00_amd64_92a8bc50dd5047989a98fdf268fd9a08e7590361658e93a28c30ab03b54bf7d6.deb
pool/kubectl_1.9.10-00_amd64_6f885fb1eb44245ed655a867d8590db0e448e9085a292fbd0a2fea505c7b18ba.deb
pool/kubectl_1.9.11-00_amd64_51ff00ecaba11907484ac957a443c4d78e35fc42d476e905e999e96b7f1f422f.deb
pool/kubectl_1.10.0-00_amd64_e391c19fa377b84587676c5577222ceb5d8fcbde442c79a9cd55d1f344293834.deb
pool/kubectl_1.10.1-00_amd64_a0d873581d1de51178e88938df4fc968fab57d8e0ecf3b9c211829e673f98ace.deb
pool/kubectl_1.10.2-00_amd64_a3ca21b1875a20567024584f18ec6aca390e8be8a8d92e20b5b7258452e419c9.deb
pool/kubectl_1.10.3-00_amd64_ad1b331006db395c716364b3049739b0f65caa531ba24585dc15f4009eb48bef.deb
pool/kubectl_1.10.4-00_amd64_5984dc4ad9713ff8a548d8292717b8f7060896e887a96f2d800a144fd70267f5.deb
pool/kubectl_1.10.5-00_amd64_76a98e003a46be66d802b5c79ba11fdc6f10463e524e67db02db1684fd9dad8e.deb
pool/kubectl_1.10.6-00_amd64_d258552fb68beb8896ebcb854b31464fd0f80819bdc7943c9ad4e88e88adc7ee.deb
pool/kubectl_1.10.7-00_amd64_52542710044cb4ad30def799dc4392d18a56cd6f2dc0d675825a085d0b2882f1.deb
pool/kubectl_1.10.8-00_amd64_ef17dfbb5de2b0a96d2988381a3c049a450ecf1ced1664c5bf6785375a6910ff.deb
pool/kubectl_1.10.9-00_amd64_c10c2a2409e1d46047476977b0b5bd95716189873c2a6e77901e46a8d0b2d5da.deb
pool/kubectl_1.10.10-00_amd64_5db60542dc46646b26dc798487fdc1513137695b9403fbf6aa708c4e267e8919.deb
pool/kubectl_1.10.11-00_amd64_bf2f2f36be03f7e551641e90746f7ad26e72601e01dfee26fca34a1fe8efb02e.deb
pool/kubectl_1.10.12-00_amd64_b2b75dcd72f73f84cefee2788a0df786a6531b0ea462c753c7f4aaf8dbcced59.deb
pool/kubectl_1.10.13-00_amd64_46d296e9a00d96a4f7a2f64e19e1b8aa846a0fc187bbab7aaf8309478a988b25.deb
pool/kubectl_1.11.0-00_amd64_2d3db54de02a65983e028ef5f9391230b8661e31a15f563787463df781b3aa31.deb
pool/kubectl_1.11.1-00_amd64_d5cdd588ae155f60f59a77af17afabdc871b558da506fbe7a4bac894e5ddc023.deb
pool/kubectl_1.11.2-00_amd64_49e2a857e4852da0c27e3e92bc92fef4d33db7c93c2a4628cb9374e3a486bc92.deb
pool/kubectl_1.11.3-00_amd64_36eb4192b3726ce626aca3d8f116487508481f53152191a1e0abe0b91ab8ccb1.deb
pool/kubectl_1.11.4-00_amd64_fa5c1cf4fed170d958d421486c036ee6bfea14c542a3b4b59f3573050ba1d317.deb
pool/kubectl_1.11.5-00_amd64_d91fb68903b1d778a3c3d9aefd965a476bc50c24ecd6e12c64e24b85b3510eaa.deb
pool/kubectl_1.11.6-00_amd64_edf916789809222ad40b5ead1d23a94522833c5792569cf76e30f1805196a712.deb
pool/kubectl_1.11.7-00_amd64_c690be55439c3d8820c73394796babbe4d47ecb20b1d945d57c3ce08f460a190.deb
pool/kubectl_1.11.8-00_amd64_464b4108119827d9eb3e8a628a16cbfab033637a4212650190342f24ff260fb0.deb
pool/kubectl_1.11.9-00_amd64_767549b796f6d09158b5b417475d3f5111e90295f7e41512baae3a30dc78a16b.deb
pool/kubectl_1.11.10-00_amd64_d66a02f63226306fcbb193013174821ea63999dd91a7048fc0613bf4da161efc.deb
pool/kubectl_1.12.0-00_amd64_242af67011dc074d0683131d96c22eb586f44553d0626767f0313d1eb8dc2b9f.deb
pool/kubectl_1.12.1-00_amd64_f02235e5b8633b589e6f7c39a581b29bef4ea723f52812c152a609302073e5fe.deb
pool/kubectl_1.12.2-00_amd64_ec905f7eb4a579661afdda2122ca8fa63fc46f561afd905e0c6f9eb5ded3671c.deb
pool/kubectl_1.12.3-00_amd64_4acbe6dfb3760bb2376f686c5bc9afa99f601329895df6f7cc9712fc81c81723.deb
pool/kubectl_1.12.4-00_amd64_523056d573555f0a60b4031ff6efd59598eb75210bd35c5dbaad35835e378de8.deb
pool/kubectl_1.12.5-00_amd64_63227207caf7d58e3bf1583bd4080c81931f7b48b6b6aa24b309a3c2fa2fa3ef.deb
pool/kubectl_1.12.6-00_amd64_800332fc8d3a91be2cff5137e63b4e15565f8f980584adf34f00c27f7a2b89cd.deb
pool/kubectl_1.12.7-00_amd64_00dbe2b6f91b1b8beffb1731ca77d5b37efbaebd20f8ef2eb1837bbd4370092c.deb
pool/kubectl_1.12.8-00_amd64_a6cbdf22d5cbe7a77784a9dc02797b4365fe9ad05473f96cc9f82cc8f10c2ba1.deb
pool/kubectl_1.12.9-00_amd64_b3b61921a3e8bbe955b5c882f3d910fcda092a44a78db4eacc139a1869b57678.deb
pool/kubectl_1.12.10-00_amd64_896bcfcacec129e50e5dcf6121ae5c32ff6b53c3fc4c5a082226d74b5fbfb0c5.deb
pool/kubectl_1.13.0-00_amd64_6e20e10266a9e0698489555a822c13af97e3a5d815fdea17a36995152550ab74.deb
pool/kubectl_1.13.1-00_amd64_85ea781848ea3e9aff07cec72aaa3a0fdddafeaf2a969e8f3f29c0778b442faa.deb
pool/kubectl_1.13.2-00_amd64_fedd83bfb3726db0ad367bc0183fa5d32ecbf30bb806963cd0c027bf2b794b0a.deb
pool/kubectl_1.13.3-00_amd64_f8ac822c2fef60a2bc13fcb68ecd786f735c1e76e1fa02cd932f339ccb5cecd1.deb
pool/kubectl_1.13.4-00_amd64_40eeb57411431e9c317002357210d2231268bd9b5bd886ad1246459481d549e5.deb
pool/kubectl_1.13.5-00_amd64_c1285f382497a45307c184aba77afe41e3200d6109ccd9cd9852800302f5495f.deb
pool/kubectl_1.13.6-00_amd64_b70158dbf1083f5f630e48a7ce84ee4ac9df3305ef4c5b56357ef134bbf85b51.deb
pool/kubectl_1.13.7-00_amd64_a93591488a26f891e03f3d083f064f531e27975418147b0d8e739d6cdc803b6c.deb
pool/kubectl_1.13.8-00_amd64_b51dcc633407547c4d92e53ca7685b8890b9b8a69a73b3cf363f59764f302767.deb
pool/kubectl_1.13.9-00_amd64_c5fe37dbe847b81f365884c81b5c87b0ef711471fe8eedb4c642497cfce2a5d4.deb
pool/kubectl_1.13.10-00_amd64_b3db8c0736c29afebfc3258fa3e377b53bbe265b9bc06a17058a7f1743786535.deb
pool/kubectl_1.13.11-00_amd64_48469e504f2f9643817f064084d9c28618e2475d7503842b60b5628246a98e37.deb
pool/kubectl_1.13.12-00_amd64_b32ad754b6b4910ca709d76ec909bf28e227082a38867870d80205af4631467e.deb
pool/kubectl_1.14.0-00_amd64_f6aa3b36f039aa9ce3614c08978ecbdff0900c7fc966bde6647c5d5b969dfabc.deb
pool/kubectl_1.14.1-00_amd64_790b6cb2747741d766f595d924bfe25c4f83fdf24f8d64cf184b790d75f3d007.deb
pool/kubectl_1.14.2-00_amd64_98c512ffae2ab4b42b86ba3ae51c9a3621828867c395f6729931d977b5c51ddb.deb
pool/kubectl_1.14.3-00_amd64_86aa11d7fe030c6dec0eadd5eac5d1faa06714652358254511fc8b00c2eda9ff.deb
pool/kubectl_1.14.4-00_amd64_caed738cd96163a28b9f9d685d7c9eabebcc9843037d3e7638c23cb347621dbf.deb
pool/kubectl_1.14.5-00_amd64_87a17024aef018b54e62ca08394a75a9508ff244c99784740637cfbb22d978bd.deb
pool/kubectl_1.14.6-00_amd64_7d93d4db5348bde6fa67f2fa1d44b3bc4448b9620b252eea0c7cf1ce763d5baf.deb
pool/kubectl_1.14.7-00_amd64_2eb0e7867fa10a47d02abc04efa990803ffc44014d84aa29c44415b0f225b2f3.deb
pool/kubectl_1.14.8-00_amd64_e5ecc61a3ffb77af088033b31510afbabd4d6526b1639cef0f17f58b5fd1a647.deb
pool/kubectl_1.14.9-00_amd64_0f0bb925d3a662e699a135af91406a0e1e7fda6fc8bcee6addaf8cc8b7018e13.deb
pool/kubectl_1.14.10-00_amd64_d99ce89e6d2c461fd16dbb17fb83facaf94a7de5cc41cc6a69b35fccf410b115.deb
pool/kubectl_1.15.0-00_amd64_c7d9a92902c178bea81001e2d753ecc00d3575f25ea2cf0737038368946c7f69.deb
pool/kubectl_1.15.1-00_amd64_782942e1a713d28d5f38075deb19b298f896756aa4cc889f441dcae69fd81e07.deb
pool/kubectl_1.15.2-00_amd64_45bcc3f1dfa3ad9d6fe2c5d9e6509cfc38a5b65e40b7bf2924f6974853be421d.deb
pool/kubectl_1.15.3-00_amd64_f798287e734bbf94e438651311c4ca16866559871c1dfb34a33628257249cf7a.deb
pool/kubectl_1.15.4-00_amd64_f4ffdce61580112eecc65832d51b46f5347c4c93e37ae1714dc3a20afb563854.deb
pool/kubectl_1.15.5-00_amd64_bc99b7c6736e0d254263f270a4fec7e303fd6cb77d5ee97209ea7b34e539e4bc.deb
pool/kubectl_1.15.6-00_amd64_a82e18c2533dc96cda21e03c605f325f352ac6fabbdd09fef43cca97254b436e.deb
pool/kubectl_1.15.7-00_amd64_83a4bd7636fc1c94ddd08f2a0cd1afad814df716e084bfe9c4f1aaf003aec994.deb
pool/kubectl_1.15.8-00_amd64_1549c2499b976043b6a50c85dfa4d47d2fe08b5e0dd59b175e3bc3b12f175adc.deb
pool/kubectl_1.15.9-00_amd64_2726b6f779246359046fd12487a8ccde944bd65e0a7f6e65ad8efb2bc94d2167.deb
pool/kubectl_1.16.0-00_amd64_679986772c12ed40781ae02317f3211f8615427c033194618ba2fdecc1cee43f.deb
pool/kubectl_1.16.1-00_amd64_73509770c675609a9be110b60b3c948110b4f06f1de60eb6796f73cb601c9493.deb
pool/kubectl_1.16.2-00_amd64_3fbcf105864121e234a95cef203b08212e3eccefa96e1e37b10e20b37968ab93.deb
pool/kubectl_1.16.3-00_amd64_bbca932fd543846177182fcba147e8b5375875038c7d0508f520b3f5a31bf4d9.deb
pool/kubectl_1.16.4-00_amd64_0615b160cb88a40272a212ad9ed4929feb3aaf4b4f665bbc295c68613169d267.deb
pool/kubectl_1.16.5-00_amd64_338f26d79fc0cb395461f4ff7e2d55a132cc27da7ffa7a61848106406e934df3.deb
pool/kubectl_1.16.6-00_amd64_a3c80436c55a7cf3122e4a176e24ebb5c26abb122933e7854a5ca429b081c1ae.deb
pool/kubectl_1.17.0-00_amd64_fba222a6fd264d10f11a1a5fdf676ad4b3d41aebe9c23dae5155ef20d0a964dd.deb
pool/kubectl_1.17.1-00_amd64_0dc19318c9114db2931552bb8bf650a14227a9603cb73fe0917ac7868ec7fcf0.deb
pool/kubectl_1.17.2-00_amd64_4cf57f0a5445409293963fddc279fbba9f48c6c280fa27479651d1908848b96a.deb
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
pool/kubelet_1.8.14-00_amd64_b5f96759cc1c51138474a66738a94482428d90a0b3e553505d7dafe251503993.deb
pool/kubelet_1.8.15-00_amd64_dd4034732b629067c07cf989c642a7a612afd727fdb15b3a4fc9b3bb94c8122d.deb
pool/kubelet_1.9.0-00_amd64_100df9788226fe76ce828503cf24b8c4c6f9705f15504093844c9138e0b2a97f.deb
pool/kubelet_1.9.1-00_amd64_e1111bdd0cb45976216b7498a0b1d68d15ddaacc462a067d831b0f20ae8c276b.deb
pool/kubelet_1.9.2-00_amd64_cdb2335fec48cd8cb3ad2bdd874fb29828ceb893d200ebee44657ae51f0851b4.deb
pool/kubelet_1.9.3-00_amd64_50317ce119d81a3a69ba891c5c77b48b153b97a175598dc62ee811e081a063ee.deb
pool/kubelet_1.9.4-00_amd64_c8c2af29e421adf188c94b7049e30ed1b6368dad49cfe5e8c865d43a16be6b76.deb
pool/kubelet_1.9.5-00_amd64_5400807b4c667bfeaa5074245fb1edb20cf4b1b7f3ac88cb23df437f8558cd13.deb
pool/kubelet_1.9.6-00_amd64_abc92a438e6bcade2e879b789fa0634474f448df7428157703206448805a948f.deb
pool/kubelet_1.9.7-00_amd64_6a510a9b3d8b54473e8237cda857be4487dbc8cbd2b6b84ad292f5a3c3ec01c5.deb
pool/kubelet_1.9.8-00_amd64_1119302b483e26aa69963e683002056df72acee69e06514e7773b3786b258053.deb
pool/kubelet_1.9.9-00_amd64_5218659dc22bcd6b18fe18ac110427ec8be12e175b290ca5ea5e60d8a3e531a6.deb
pool/kubelet_1.9.10-00_amd64_c10641e06a3c2f5ff2cc6c97686c29fb40f076b5c7f2d37754c84f4fb4fed2a9.deb
pool/kubelet_1.9.11-00_amd64_bcbc6de418e0e97c3210bab33716b8ecb0ad151f0948021437800aa3b4d9bbab.deb
pool/kubelet_1.10.0-00_amd64_211eced3f0715db1dbbb9fa0973299a8f6b764ef43e25bd94fece3947b13091f.deb
pool/kubelet_1.10.1-00_amd64_7002d1aa12fea2f0c81efdc5ea577607e387d21177c6fa57e1f897037325b824.deb
pool/kubelet_1.10.2-00_amd64_3ff69468521886de7b64ba9b4932bab7e6ef71f7344ace36ac0855df7bd2427a.deb
pool/kubelet_1.10.3-00_amd64_cc71162f227c16487999c2f8190bf8714127410633b62383c8b733c6fce23fbb.deb
pool/kubelet_1.10.4-00_amd64_999ab6b2379ea1496d0e8f2099ceffb39a23043039484707639cade63bb07753.deb
pool/kubelet_1.10.5-00_amd64_10aa030ab7d406ca639c57d16ab154b32e75f55925fb24065e33359c918722af.deb
pool/kubelet_1.10.6-00_amd64_3430c8f965aa03419b96fb43ec3ec9ff3471a0413e0e107540343aee33e20a66.deb
pool/kubelet_1.10.7-00_amd64_fd815d1c628fc90f5a0a3b3150d3314e6639a91c940fc212e13114a7086f1b10.deb
pool/kubelet_1.10.8-00_amd64_443f3fabf3e6a71f685c3d422c4a50ceddce46ab3ac2778c9ee5adf8879f8d6b.deb
pool/kubelet_1.10.9-00_amd64_dcd129388b2fe925712da9f69725a7ed91bf261bf401773cf1617a103f366a4c.deb
pool/kubelet_1.10.10-00_amd64_f0b95158b6b5823d1b379a26f3ae5bb58cb5e07a3910de4581d9755abc2d62cb.deb
pool/kubelet_1.10.11-00_amd64_66d5acc214366f2a01a36038266b50d14d794d635a05d760d5bce25573b18c6a.deb
pool/kubelet_1.10.12-00_amd64_c827b0e83b82824601b1894ce8528e71f3ed11467a742b1e097b5f5b1daf989c.deb
pool/kubelet_1.10.13-00_amd64_aa7069a9b903f263a16bf54059363e75b697b6ef9529c1042f8fe2e238eaf686.deb
pool/kubelet_1.11.0-00_amd64_464fcc0e54df9ec18935f8b36f33cb048a3a6c638ec80205900bff48dc61b2a6.deb
pool/kubelet_1.11.1-00_amd64_28c71ce044870c518d0cc12602e331ffae5729d26b4873c99ec587682ea88781.deb
pool/kubelet_1.11.2-00_amd64_7537d39713573280e1cc245915fc7565ac49d041fbd0e0515daa1ea2ac659dbb.deb
pool/kubelet_1.11.3-00_amd64_f992d76af73d547fcc1d8e3c1751ca05f801ac3682868786b7fdb528cba48ac0.deb
pool/kubelet_1.11.4-00_amd64_81d43d83a7a87f7742b2cc779e90a014f7e72a850766dc652292c02fb19b0062.deb
pool/kubelet_1.11.5-00_amd64_d0a93e8fe821f989d1226b2feb6c62c9de6eb92cadc103da3f680ec7dbc5bd78.deb
pool/kubelet_1.11.6-00_amd64_da813320d69fa2909137da42dbe41adbfe0b0a1023c5f2f1d9edc514e29d15d4.deb
pool/kubelet_1.11.7-00_amd64_73ee422e4d381d6226b93c6fdb5c5e6dcc875d9a91c6683a054b45ce71423e08.deb
pool/kubelet_1.11.8-00_amd64_4becff872b43a37a95b156282ac170d434deb0729e6d50de4334b2905ecd5420.deb
pool/kubelet_1.11.9-00_amd64_807679ed430d50e623e9c56345a6c6c2be1f9a9579d4c0cc8d80c440aa87954d.deb
pool/kubelet_1.11.10-00_amd64_b9d8e8599b9d606930e6d31159843cd60e1c89d978ef1d2123f02985b2922564.deb
pool/kubelet_1.12.0-00_amd64_24f580dcc7cb8cb8439da40c2e1488f1851dac9dfaac0b64b94492538e59f948.deb
pool/kubelet_1.12.1-00_amd64_ae75f76207874c43f7e20e8184265e6a9fe38ef0c93e61115362280227bed23f.deb
pool/kubelet_1.12.2-00_amd64_62015ea9af675de7f785081d518110d4df18580ddf9d382d2616e389a196f3da.deb
pool/kubelet_1.12.3-00_amd64_23cddff0a0db034778e7eba3405173fe4635d4b806ee77d943011c5e395e6eb7.deb
pool/kubelet_1.12.4-00_amd64_c8197c9e1105b81a08f31657645ed1a968fb54fd822cc0ca84b880da6db52636.deb
pool/kubelet_1.12.5-00_amd64_d0086a819074bd850859f213645401107a5573f84b9363060c65a9c337f70d91.deb
pool/kubelet_1.12.6-00_amd64_23adc40edbd6b23c416fe4a615e6abf3fb6560b14bfb547b8b34df53224f8067.deb
pool/kubelet_1.12.7-00_amd64_a6da4d228370714066a00ad5ae786cf4e0aff84a4438a4133a8782a16f428e4c.deb
pool/kubelet_1.12.8-00_amd64_ce9409c85a8d50a2315163aa1d944936e61f716af557cb0e173aec423030ec14.deb
pool/kubelet_1.12.9-00_amd64_964d7de4aabdc09a5e807a2e3d16f17d66fbb9b5cacc64ca6c0255cd5699332c.deb
pool/kubelet_1.12.10-00_amd64_cfdef95a909ed3b1de9e53a4499c7bfeb2e47eed3b6e9d13c0adc8ad60d2076d.deb
pool/kubelet_1.13.0-00_amd64_f428332552df40f60aa9e1be21632406279dd3d7f64528c4f9058e8fe78eb119.deb
pool/kubelet_1.13.1-00_amd64_158e81cf03ce9a1e2c9cc0dc609f864ba4959ceade5227d5b63708a2c2e3cedf.deb
pool/kubelet_1.13.2-00_amd64_82d092e6177b68cfdde60b9c58de4091ab466e61a4a0328d897275811bc62500.deb
pool/kubelet_1.13.3-00_amd64_48d32f5aa84b897343500c9d2a4d283e46e230e7916902dffa2cdad34efd9b1b.deb
pool/kubelet_1.13.4-00_amd64_4fbf1a8278a7d36f9fc3a2435bf33450a450699eb09a542208824bacbc07f948.deb
pool/kubelet_1.13.5-00_amd64_d64da1273816d9f634d52e499cf36868a50743da5390f72ef9039927d43110e5.deb
pool/kubelet_1.13.6-00_amd64_ba374740ab4ab6ef647aeff12431777ed64cf7a9e19267dfe1b99d6d59dfe14e.deb
pool/kubelet_1.13.7-00_amd64_a109c038d37d5e6cd08122ee7d09c15e117ceb82a90d909a2d1d5d81a9d48831.deb
pool/kubelet_1.13.8-00_amd64_9f654da8188e1fa9783d5145973c9147aa461a27aeff3c3f2a5bba3a28882ef1.deb
pool/kubelet_1.13.9-00_amd64_d824078900acb099e53f80ff5f3f6f00a61392d5d907988080b60215fc529748.deb
pool/kubelet_1.13.10-00_amd64_5c094e22dd6ccc126e54c01db3f31d0cf033402417494ecceeba9c809dbace7e.deb
pool/kubelet_1.13.11-00_amd64_67ffa1960546d342d4959deabbe20f546e107297e785cb0a797a15eb605e41d6.deb
pool/kubelet_1.13.12-00_amd64_bbb3a620583678b250fa4b64a208a20cc5592bf7ff926ccbc9f8bd92e8515f94.deb
pool/kubelet_1.14.0-00_amd64_a0b145a6601fb6c03c8f401d65c5c94cc75bf17bc7a19a00fdff44a17b6e230f.deb
pool/kubelet_1.14.1-00_amd64_60ce697defdd990aee5e4b40abc7274f71e3aef6d9cc4feab240c8dc5681ed05.deb
pool/kubelet_1.14.2-00_amd64_6f10676ec9e3ae926be6eca3a2dbebe36b13b394a6d989932419e94ef3264694.deb
pool/kubelet_1.14.3-00_amd64_7367f139f822617b472475f8b7db6db0716a8513df9eeb9844a33e4b1c3d844f.deb
pool/kubelet_1.14.4-00_amd64_ad981158d637a54a56dbc567a5d7bed9de2237e778b97918b9e92bb7430715da.deb
pool/kubelet_1.14.5-00_amd64_f4125f9d0b91e6a32b2fdac30b3d1203a8f41e12678db3241889ab3dd01c4975.deb
pool/kubelet_1.14.6-00_amd64_b31764889eae515e862eb7646b9ac1e7719f31736366cca75b370872ecfed0bd.deb
pool/kubelet_1.14.7-00_amd64_9a85e300c2efbd29fd7bc78d314f7e03577c6068227ec1571f8cca4229e50ced.deb
pool/kubelet_1.14.8-00_amd64_fd90fe7eb59436d416103f7ea2bc905a73b8fbafead7074fd33f1e2ecec774c3.deb
pool/kubelet_1.14.9-00_amd64_0997197a365ce053050b65d0d0b88a76059b9ef3b5d0e1ef7d3c4a8d9a29e9e3.deb
pool/kubelet_1.14.10-00_amd64_c9612950af755428347c4596ad9949123af3c564aa6cdc724c5cc54df50fd822.deb
pool/kubelet_1.15.0-00_amd64_2e2c6e716b5effa1358beba20129baa5651d6ab085a094346645f416c038389b.deb
pool/kubelet_1.15.1-00_amd64_ac23776caec693946c79c349d76bade5bd10e822d0b6e33581d8fbdb7420d3f3.deb
pool/kubelet_1.15.2-00_amd64_7d5b5aeb52177ee84169cdf4a1d624ce27e79a1d3cee9f500dc360ffe3f5adb6.deb
pool/kubelet_1.15.3-00_amd64_04030c9f58fdb521f5587b47d15a7382419f2600195450079380cf34b02abc75.deb
pool/kubelet_1.15.4-00_amd64_be85fb9796d3b433a4eef58c95d953836c39917578f6cd4f3328133cf865fc04.deb
pool/kubelet_1.15.5-00_amd64_feba4d4831a02a994a71708885f0fd043b983ae2787a6d2eb1f1ae80b0f199f0.deb
pool/kubelet_1.15.6-00_amd64_ef88a8ae8a9474be1aee0b508ae51ab990503086061f1a9c83677b67466f15bb.deb
pool/kubelet_1.15.7-00_amd64_0d2f457797d92038f9ba7f0c73aa550759d865185219a9ec84fcee1fac4664ba.deb
pool/kubelet_1.15.8-00_amd64_1bce461a8ece0f4f3ca3d45f7b2a837d679768d2ca34a15239417deea5d8ede0.deb
pool/kubelet_1.15.9-00_amd64_6a8a4a1643d5e2c84b2c99c478eb94c008bc3ac3ee7a1bdf9d4dd0b911acf04e.deb
pool/kubelet_1.16.0-00_amd64_e919939f5dad4bc7b0047fe2cf2870ab1946e87f948ab7f75bc1d63305436664.deb
pool/kubelet_1.16.1-00_amd64_ae1d16ae8e2b73a3224a41c0730be2d812870b7cb7307a8f02d13038d572d05d.deb
pool/kubelet_1.16.2-00_amd64_f36a759a47ab98ec48e32503c4934c9cb83c323108b2c0a7aaf4c964cf62e5c1.deb
pool/kubelet_1.16.3-00_amd64_77c3d05510bc1f753834c30de6a1a7fc032ecfb906527bc270663428b371c477.deb
pool/kubelet_1.16.4-00_amd64_a2fc1b783de61bd4e8baaa8b8549f3b9c7c22a6d2922dbd24f693c5d241d7f0b.deb
pool/kubelet_1.16.5-00_amd64_bccca7c4203a55af958e073d3868c3df210052976e156aa38c5b5c20eb529bfd.deb
pool/kubelet_1.16.6-00_amd64_37e04fde1b70d0da78ce5ee8b4ca7ed9e40350a9bd90930466a2382c13bdc587.deb
pool/kubelet_1.17.0-00_amd64_ee141d9d5fe9c4db5c88022d837c9a587d52223c9dac5171c60b596a8381365c.deb
pool/kubelet_1.17.1-00_amd64_c9133d2dedcc6364fe525cd8ea15a3024b7bf866f5465121d561032006e20be6.deb
pool/kubelet_1.17.2-00_amd64_f22c86b79e7e98d18aaca1690d96b977bca11e9541170716bab0cb664f9e583c.deb
pool/kubernetes-cni_0.3.0.1-07a8a2-00_amd64_9e41a275b2afeb51dcde86b922c056c7b6dc656b54dd66fa2f1a0bb8266e9c22.deb
pool/kubernetes-cni_0.5.1-00_amd64_08cbe5c42366ec3184cc91a4353f6e066f2d7325b4db1cb4f87c7dcc8c0eb620.deb
pool/kubernetes-cni_0.6.0-00_amd64_43460dd3c97073851f84b32f5e8eebdc84fadedb5d5a00d1fc6872f30a4dd42c.deb
pool/kubernetes-cni_0.7.5-00_amd64_b38a324bb34f923d353203adf0e048f3b911f49fa32f1d82051a71ecfe2cd184.deb
pool/rkt_1.25.0-1_amd64_65f4d768116520be6ee5d56ebcdce87b8b39fa2e41345b7c2960a01f91b7c081.deb
pool/rkt_1.26.0-1_amd64_af62ca3979b90f2fcbe7c48ba782dc1b0c4266832e7544ef4c69ce0a8a7dfa87.deb
pool/rkt_1.27.0-1_amd64_b4f1d67b835d1f1b0263ddfc3f8df80cd7c9e36ac78d983e12aca5575835e122.deb
pool/rkt_1.28.1-1_amd64_66ea8b0ef5724aa364a1273b794b2aaf5c302c774e076115adee8f7c04a0e0f9.deb
pool/rkt_1.29.0-1_amd64_ea87d719359030f33fd48890875c934135c62eccda72c37d79ff604307b905b5.deb
```

Update packages into script and do execution
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ ./apt-pool-curl.sh 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 7872k  100 7872k    0     0   823k      0  0:00:09  0:00:09 --:--:-- 1479k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 8533k  100 8533k    0     0  2261k      0  0:00:03  0:00:03 --:--:-- 2261k
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 18.3M  100 18.3M    0     0  2336k      0  0:00:08  0:00:08 --:--:-- 2171k
```

```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ ls https0x3A0x2F0x2Fpackages.cloud.google.com/apt/pool/
cri-tools_1.11.1-00_amd64_e6300f0f4ec2fb5d6967533416580e1a22be90277944370ceb2551b04d7bc1a3.deb
cri-tools_1.12.0-00_amd64_2d9f048a50a9dfeceebd84635f1322955aca6381d9c05b4d60b3da1edb7d856c.deb
cri-tools_1.13.0-00_amd64_6930e446a683884314deef354fbd8a7c5fc2be5c69c58903ad83b69b42529da4.deb
docker-engine_1.11.2-0~xenial_amd64_5554a8bc383e65fb10d556239c72b26457cc5d97f49e3a353c3382f8434e7d21.deb
kubeadm_1.10.13-00_amd64_4594e00a6f68d2c7f8c7990cf6408ca6f1582b9df05de22d3d68fd84cf05344c.deb
kubeadm_1.11.10-00_amd64_a518251af86868412223ce6b12d0212d5783278419bf7cf403b77cce1a31cfb9.deb
kubeadm_1.12.10-00_amd64_600513b5bdc3ff15caab99660a4e0e8d5b7d31bc907ce63f8a4fcd18e0c9c6f6.deb
kubeadm_1.13.12-00_amd64_5791923bba94b850ca898fdfaa75916bdd3b4f5d86a50ea738d96f14230029b9.deb
kubeadm_1.14.9-00_amd64_5c8c4bba3ba94e1349acc2148b57844ffad756bbb20e2ed76e655e0a074ee244.deb
kubeadm_1.15.6-00_amd64_2c1d026f59d6e012a8fc7658dddcde404f73c9e1159cdeb000aea0d4de3ed3b3.deb
kubeadm_1.16.3-00_amd64_8ad50e8252c65f968e35b1a38240b0ee66ef3b72c5d498830788e123238aa1a6.deb
kubeadm_1.17.2-00_amd64_041e0e6cc71212fff2cdaadfa7fd52d992fad23efd15157d360df86a22cf7649.deb
kubeadm_1.8.4-00_amd64_0088836fbb451bc49ece82f34c035f50f2e1dd4dea78f6d585574d585da11e8e.deb
kubeadm_1.9.0-00_amd64_191bd1d8a63d263cdb318f77b03fad6c8387a912cf16a21b9b24d7e9108b4911.deb
kubectl_1.10.13-00_amd64_46d296e9a00d96a4f7a2f64e19e1b8aa846a0fc187bbab7aaf8309478a988b25.deb
kubectl_1.11.10-00_amd64_d66a02f63226306fcbb193013174821ea63999dd91a7048fc0613bf4da161efc.deb
kubectl_1.12.10-00_amd64_896bcfcacec129e50e5dcf6121ae5c32ff6b53c3fc4c5a082226d74b5fbfb0c5.deb
kubectl_1.13.12-00_amd64_b32ad754b6b4910ca709d76ec909bf28e227082a38867870d80205af4631467e.deb
kubectl_1.14.9-00_amd64_0f0bb925d3a662e699a135af91406a0e1e7fda6fc8bcee6addaf8cc8b7018e13.deb
kubectl_1.15.6-00_amd64_a82e18c2533dc96cda21e03c605f325f352ac6fabbdd09fef43cca97254b436e.deb
kubectl_1.16.3-00_amd64_bbca932fd543846177182fcba147e8b5375875038c7d0508f520b3f5a31bf4d9.deb
kubectl_1.17.2-00_amd64_4cf57f0a5445409293963fddc279fbba9f48c6c280fa27479651d1908848b96a.deb
kubectl_1.8.4-00_amd64_b48511a481ddcfbf935ad935bc6c3992c1c4315fcd8f3f794e367b9b26b775be.deb
kubectl_1.9.0-00_amd64_8ea712c18d89d56090c8753a9630d22fd8ae5cb03d4ec79a1fe6d262c8b4eb36.deb
kubelet_1.10.13-00_amd64_aa7069a9b903f263a16bf54059363e75b697b6ef9529c1042f8fe2e238eaf686.deb
kubelet_1.11.10-00_amd64_b9d8e8599b9d606930e6d31159843cd60e1c89d978ef1d2123f02985b2922564.deb
kubelet_1.12.10-00_amd64_cfdef95a909ed3b1de9e53a4499c7bfeb2e47eed3b6e9d13c0adc8ad60d2076d.deb
kubelet_1.13.12-00_amd64_bbb3a620583678b250fa4b64a208a20cc5592bf7ff926ccbc9f8bd92e8515f94.deb
kubelet_1.14.9-00_amd64_0997197a365ce053050b65d0d0b88a76059b9ef3b5d0e1ef7d3c4a8d9a29e9e3.deb
kubelet_1.15.6-00_amd64_ef88a8ae8a9474be1aee0b508ae51ab990503086061f1a9c83677b67466f15bb.deb
kubelet_1.16.3-00_amd64_77c3d05510bc1f753834c30de6a1a7fc032ecfb906527bc270663428b371c477.deb
kubelet_1.17.2-00_amd64_f22c86b79e7e98d18aaca1690d96b977bca11e9541170716bab0cb664f9e583c.deb
kubelet_1.8.4-00_amd64_601882506070723b643552ae98325c849840b09b1fc1666de74c7b69a07f06df.deb
kubelet_1.9.0-00_amd64_100df9788226fe76ce828503cf24b8c4c6f9705f15504093844c9138e0b2a97f.deb
kubernetes-cni_0.3.0.1-07a8a2-00_amd64_9e41a275b2afeb51dcde86b922c056c7b6dc656b54dd66fa2f1a0bb8266e9c22.deb
kubernetes-cni_0.5.1-00_amd64_08cbe5c42366ec3184cc91a4353f6e066f2d7325b4db1cb4f87c7dcc8c0eb620.deb
kubernetes-cni_0.6.0-00_amd64_43460dd3c97073851f84b32f5e8eebdc84fadedb5d5a00d1fc6872f30a4dd42c.deb
kubernetes-cni_0.7.5-00_amd64_b38a324bb34f923d353203adf0e048f3b911f49fa32f1d82051a71ecfe2cd184.deb
```

The `apt-pool-curl_xenial_all.sh` has all packages to download, but in Mac
```
fanhonglingdeMacBook-Pro:https0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm fanhongling$ sed -i '' 's/^pool.*deb$/& \\/g' apt-pool-curl_xenial_all.sh 
```

### For Ubuntu Trusty (14.04)

Ubuntu Trusty (14.04) can only install kubectl
```
[vagrant@localhost http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm]$ egrep '^Filename: pool/\S+deb' https0x3A0x2F0x2Fpackages.cloud.google.com/apt/dists/kubernetes-trusty/main/binary-amd64/Packages | awk '{print $2}'
pool/kubectl_1.5.6-00_amd64_620d1f16fcf779e72e778fd000b247612eabb351e054573b1153df8f6aad1342.deb
pool/kubectl_1.6.0-00_amd64_7021195dccb229242f160a98c069ccfe3097b3af1549b81940d201313238b519.deb
```


## Repository mirror configuration

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


## v1.9.0

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