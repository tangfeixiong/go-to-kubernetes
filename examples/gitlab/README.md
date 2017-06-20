
### Example push

https://github.com/tangfeixiong/osev3-examples

```
[vagrant@bogon osev3-examples]$ git remote -v
origin	https://github.com/tangfeixiong/osev3-examples (fetch)
origin	https://github.com/tangfeixiong/osev3-examples (push)
[vagrant@bogon osev3-examples]$ git remote add gitlab-local http://root@172.17.4.50:10080/root/osev3-examples.git
[vagrant@bogon osev3-examples]$ git remote -v
gitlab-local	http://root@172.17.4.50:10080/root/osev3-examples.git (fetch)
gitlab-local	http://root@172.17.4.50:10080/root/osev3-examples.git (push)
origin	https://github.com/tangfeixiong/osev3-examples (fetch)
origin	https://github.com/tangfeixiong/osev3-examples (push)
[vagrant@bogon osev3-examples]$ git push gitlab-local master
Password for 'http://root@172.17.4.50:10080': 
对象计数中: 342, 完成.
Delta compression using up to 2 threads.
压缩对象中: 100% (173/173), 完成.
写入对象中: 100% (342/342), 13.93 MiB | 172.00 KiB/s, 完成.
Total 342 (delta 108), reused 314 (delta 96)
remote: Resolving deltas: 100% (108/108), done.
To http://root@172.17.4.50:10080/root/osev3-examples.git
 * [new branch]      master -> master
```
