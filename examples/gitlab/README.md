


fanhonglingdeMacBook-Pro:gs-rest-service fanhongling$ git remote --verbose
origin	https://github.com/spring-guides/gs-rest-service (fetch)
origin	https://github.com/spring-guides/gs-rest-service (push)

fanhonglingdeMacBook-Pro:gs-rest-service fanhongling$ git remote add gitlab-local http://172.17.4.200:10080/root/gs-rest-service

fanhonglingdeMacBook-Pro:gs-rest-service fanhongling$ git remote --verbose
gitlab-local	http://172.17.4.200:10080/root/gs-rest-service (fetch)
gitlab-local	http://172.17.4.200:10080/root/gs-rest-service (push)
origin	https://github.com/spring-guides/gs-rest-service (fetch)
origin	https://github.com/spring-guides/gs-rest-service (push)


fanhonglingdeMacBook-Pro:gs-rest-service fanhongling$ git pull gitlab-local
Username for 'http://172.17.4.200:10080': admin@example.com
Password for 'http://admin@example.com@172.17.4.200:10080': 
You asked to pull from the remote 'gitlab-local', but did not specify
a branch. Because this is not the default configured remote
for your current branch, you must specify a branch on the command line.


fanhonglingdeMacBook-Pro:gs-rest-service fanhongling$ git checkout -b gitlab-local-demo
Switched to a new branch 'gitlab-local-demo'


fanhonglingdeMacBook-Pro:gs-rest-service fanhongling$ git push gitlab-local
warning: push.default is unset; its implicit value has changed in
Git 2.0 from 'matching' to 'simple'. To squelch this message
and maintain the traditional behavior, use:

  git config --global push.default matching

To squelch this message and adopt the new behavior now, use:

  git config --global push.default simple

When push.default is set to 'matching', git will push local branches
to the remote branches that already exist with the same name.

Since Git 2.0, Git defaults to the more conservative 'simple'
behavior, which only pushes the current branch to the corresponding
remote branch that 'git pull' uses to update the current branch.

See 'git help config' and search for 'push.default' for further information.
(the 'simple' mode was introduced in Git 1.7.11. Use the similar mode
'current' instead of 'simple' if you sometimes use older versions of Git)

Counting objects: 1532, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (591/591), done.
Writing objects: 100% (1532/1532), 819.19 KiB | 0 bytes/s, done.
Total 1532 (delta 830), reused 1532 (delta 830)
remote: Resolving deltas: 100% (830/830), done.
To http://172.17.4.200:10080/root/gs-rest-service
 * [new branch]      gitlab-local-demo -> gitlab-local-demo



vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/gitlab$ docker network ls
NETWORK ID          NAME                DRIVER
0178d75935e1        gitlab_default      bridge              
352882a33218        bridge              bridge              
590c2ddbaedd        none                null                
14a2869a71b0        host                host                

vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/gitlab$ docker-compose -p gitlab -f https0x3A0x2F0x2Fraw.githubusercontent.com0x2Fstackdocker0x2Fdocker-gitlab0x2Fmaster/docker-compose.yml down
Stopping gitlab_gitlab_1 ... done
Stopping gitlab_redis_1 ... done
Stopping gitlab_postgresql_1 ... done
Removing gitlab_gitlab_1 ... done
Removing gitlab_redis_1 ... done
Removing gitlab_postgresql_1 ... done
Removing network gitlab_default

vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-kubernetes/examples/gitlab$ docker network ls
NETWORK ID          NAME                DRIVER
590c2ddbaedd        none                null                
14a2869a71b0        host                host                
352882a33218        bridge              bridge              
