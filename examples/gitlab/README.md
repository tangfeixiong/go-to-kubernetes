# 操作手册

## Table of content

* Install under Docker-Compose
* Push, Fork, and Pull
* 敏捷开发

## 安装

For `docker-compose`, refer to [docker-compose.md](./docker-compose.md)

For `gitlab`, refer to https://about.gitlab.com/downloads/

如：

* 清华大学镜像站
* 文档站

## Git Example

详细请参考 https://git-scm.com/book/en/v2/Getting-Started-The-Command-Line

### 示例将github的项目push到本地的gitlab

示例使用Spring Boot项目 https://github.com/tangfeixiong/osev3-examples
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

## Agile developement relations

### 项目组织

- upstream
  - master
  - milestone branch 1
  - release tag 1
  - ...
- fork
  - master
  - patch...

__基于RBAC的开发管理__

* Upstream由project manager和CI/CD团队负责管理
* 每个sub project由project manager与子项目team leader规划，如merge
* 每个开发者从子项目leader处获取fork的位置

__Programmer的行为__

每日push

* 开发者必须每日push开发内容，以能进行工作量评估

每sprint的pull request

* 一个冲刺中完成的features，经本地的unit/mock后，可进行it或staging，push到其个人的fork后，请求project manager或team leader合并到upstream
* 其leader必须严格负责review所有的pull request，提高CI/CD效率

__Leader的素质__

* 高级软件开发技能
* 精通项目管理系统，如git，maven（mvn，nexus仓库，Eclipse插件等，pom.xml），Gradle（gradle，Eclipse插件等，gradle.build）
* 熟悉Agile
* 了解CI/CD

### Gitlab的Issue管理

__Priority__

![屏幕快照 2017-06-26 下午9.06.29.png](./屏幕快照%202017-06-26%20下午9.06.29.png)

__列出Issue__

project manager，team leader，team的QA根据每个spring的CI/CD生成或其它测试结果，将问题列举到Issue

* project manager可以评估team的状态，人员的水平
* team leader答复问题，指派assignee解决问题等

__看板__

公告每个sprint的任务与结果

__Milestone__

公告每个里程碑的计划与结果

### 文档管理

项目目录结构

- docs
  - design
  - proposal
  - ...
  project1
  - src
  - ...
  project2
  ...

__使用gitlab wiki__

做到少而实用，减少各种无用的工具