# Instruction

kubeadm
```
[vagrant@openshiftdev ~]$ which kubeadm
/bin/kubeadm
```

```
[vagrant@openshiftdev ~]$ kubeadm version
kubeadm version: &version.Info{Major:"1", Minor:"8", GitVersion:"v1.8.2", GitCommit:"bdaeafa71f6c7c04636251031f93464384d54963", GitTreeState:"clean", BuildDate:"2017-10-24T19:38:10Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
```

```
[vagrant@openshiftdev ~]$ kubeadm init --apiserver-bind-port 443 --kubernetes-version v1.6.4
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
this version of kubeadm only supports deploying clusters with the control plane version >= 1.7.0. Current version: v1.6.4
```

## K8s v1.7.4

Extract docker images
```
[vagrant@openshiftdev ~]$ tar -C /home/vagrant -zxf /vagrant_f/99-mirror/https0x3A0x2F0x2Fdl.k8s.io/v1.7.4/kubernetes-server-linux-amd64.tar.gz 
```

```
[vagrant@openshiftdev ~]$ ls kubernetes/server/bin/
apiextensions-apiserver              kube-aggregator             kube-controller-manager             kube-proxy
cloud-controller-manager             kube-aggregator.docker_tag  kube-controller-manager.docker_tag  kube-proxy.docker_tag
cloud-controller-manager.docker_tag  kube-aggregator.tar         kube-controller-manager.tar         kube-proxy.tar
cloud-controller-manager.tar         kube-apiserver              kubectl                             kube-scheduler
hyperkube                            kube-apiserver.docker_tag   kubefed                             kube-scheduler.docker_tag
kubeadm                              kube-apiserver.tar          kubelet                             kube-scheduler.tar
```

```
[vagrant@openshiftdev ~]$ bindir=./kubernetes/server/bin; for i in $(ls $bindir/*.tar); do docker load -i $i; name=$(basename $i .tar); tag=$(cat $bindir/$name.docker_tag); repo=gcr.io/google_containers/$name; img=$(docker images -q $repo:$tag); [[ -n $img ]] && docker images $repo:$tag || (echo "...failed to load" && docker images $repo); done
REPOSITORY                                          TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/cloud-controller-manager   v1.7.4              a42a0fd85571        3 months ago        125.3 MB
REPOSITORY                                 TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-aggregator   v1.7.4              c1752dd09c50        3 months ago        59.16 MB
REPOSITORY                                TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-apiserver   v1.7.4              5260ecb5129c        3 months ago        186.1 MB
REPOSITORY                                         TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-controller-manager   v1.7.4              d2adddc4b1cb        3 months ago        138 MB
REPOSITORY                            TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-proxy   v1.7.4              0f3bf654ec61        3 months ago        114.7 MB
REPOSITORY                                TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/kube-scheduler   v1.7.4              b1cd468ba656        3 months ago        77.2 MB
```

```
[vagrant@openshiftdev ~]$ docker images gcr.io/google_containers/*
REPOSITORY                                          TAG                 IMAGE ID            CREATED             SIZE
gcr.io/google_containers/cloud-controller-manager   v1.7.4              a42a0fd85571        3 months ago        125.3 MB
gcr.io/google_containers/kube-controller-manager    v1.7.4              d2adddc4b1cb        3 months ago        138 MB
gcr.io/google_containers/kube-apiserver             v1.7.4              5260ecb5129c        3 months ago        186.1 MB
gcr.io/google_containers/kube-proxy                 v1.7.4              0f3bf654ec61        3 months ago        114.7 MB
gcr.io/google_containers/kube-aggregator            v1.7.4              c1752dd09c50        3 months ago        59.16 MB
gcr.io/google_containers/kube-scheduler             v1.7.4              b1cd468ba656        3 months ago        77.2 MB
```

### Offline install

Enalbe local repository
```
[vagrant@openshiftdev ~]$ sudo yum-config-manager --enablerepo=offline\*,kubernetes-local
Loaded plugins: fastestmirror
=============================================================== main ===============================================================
[main]
alwaysprompt = True
assumeno = False
assumeyes = False
autocheck_running_kernel = True
bandwidth = 0
bugtracker_url = https://bugzilla.redhat.com/enter_bug.cgi?product=Fedora&version=rawhide&component=yum
cache = 0
cachedir = /var/cache/yum/x86_64/7
check_config_file_age = True
clean_requirements_on_remove = False
color = auto
color_list_available_downgrade = dim,cyan
color_list_available_install = normal
color_list_available_reinstall = bold,underline,green
color_list_available_running_kernel = bold,underline
color_list_available_upgrade = bold,blue
color_list_installed_extra = bold,red
color_list_installed_newer = bold,yellow
color_list_installed_older = bold
color_list_installed_reinstall = normal
color_list_installed_running_kernel = bold,underline
color_search_match = bold
color_update_installed = normal
color_update_local = bold
color_update_remote = normal
commands = 
debuglevel = 2
deltarpm = 2
deltarpm_metadata_percentage = 100
deltarpm_percentage = 75
depsolve_loop_limit = 100
disable_includes = 
diskspacecheck = True
distroverpkg = system-release(releasever),
   redhat-release
downloaddir = 
downloadonly = 
enable_group_conditionals = True
enabled = True
enablegroups = True
errorlevel = 2
exactarch = True
exactarchlist = 
exclude = 
exit_on_lock = False
failovermethod = priority
fssnap_abort_on_errors = any
fssnap_automatic_keep = 1
fssnap_automatic_post = False
fssnap_automatic_pre = False
fssnap_devices = !*/swap,
   !*/lv_swap
fssnap_percentage = 100
gaftonmode = False
gpgcheck = False
group_command = objects
group_package_types = mandatory,
   default
groupremove_leaf_only = False
history_list_view = single-user-commands
history_record = True
history_record_packages = yum,
   rpm
http_caching = all
installonly_limit = 3
installonlypkgs = kernel,
   kernel-bigmem,
   installonlypkg(kernel-module),
   installonlypkg(vm),
   kernel-enterprise,
   kernel-smp,
   kernel-debug,
   kernel-unsupported,
   kernel-source,
   kernel-devel,
   kernel-PAE,
   kernel-PAE-debug
installroot = /
ip_resolve = 
keepalive = True
keepcache = False
kernelpkgnames = kernel,
   kernel-smp,
   kernel-enterprise,
   kernel-bigmem,
   kernel-BOOT,
   kernel-PAE,
   kernel-PAE-debug
loadts_ignoremissing = False
loadts_ignorenewrpm = False
loadts_ignorerpm = False
localpkg_gpgcheck = False
logfile = /var/log/yum.log
max_connections = 0
mddownloadpolicy = sqlite
mdpolicy = group:small
metadata_expire = 9999999
metadata_expire_filter = read-only:present
minrate = 0
mirrorlist_expire = 86400
multilib_policy = best
obsoletes = True
override_install_langs = 
overwrite_groups = False
password = 
persistdir = /var/lib/yum
pluginconfpath = /etc/yum/pluginconf.d
pluginpath = /usr/share/yum-plugins,
   /usr/lib/yum-plugins
plugins = True
progess_obj = 
protected_multilib = True
protected_packages = yum,
   systemd
proxy = False
proxy_password = 
proxy_username = 
query_install_excludes = True
recent = 7
recheck_installed_requires = True
remove_leaf_only = False
repo_gpgcheck = False
repopkgsremove_leaf_only = False
reposdir = /etc/yum/repos.d,
   /etc/yum.repos.d
reset_nice = True
retries = 20
rpm_check_debug = True
rpmverbosity = info
showdupesfromrepos = False
skip_broken = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
syslog_device = /dev/log
syslog_facility = LOG_USER
syslog_ident = 
throttle = 0
timeout = 120.0
tolerant = True
tsflags = 
ui_repoid_vars = releasever,
   basearch
upgrade_group_objects_upgrade = True
upgrade_requirements_on_install = False
username = 
usr_w_check = True

============================================================ repo: base ============================================================
[base]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = 
cache = 0
cachedir = /var/cache/yum/x86_64/7/base
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = True
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/base/gpgcadir
gpgcakey = 
gpgcheck = True
gpgdir = /var/lib/yum/repos/x86_64/7/base/gpgdir
gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
hdrdir = /var/cache/yum/x86_64/7/base/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = http://mirrorlist.centos.org/?release=7&arch=x86_64&repo=os&infra=stock
mirrorlist_expire = 86400
name = CentOS-7 - Base
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/base
pkgdir = /var/cache/yum/x86_64/7/base/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = base/7/x86_64
ui_repoid_vars = releasever,
   basearch
username = 

======================================== repo: dl.google.com_linux_chrome_rpm_stable_x86_64 ========================================
[dl.google.com_linux_chrome_rpm_stable_x86_64]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = http://dl.google.com/linux/chrome/rpm/stable/x86_64
cache = 0
cachedir = /var/cache/yum/x86_64/7/dl.google.com_linux_chrome_rpm_stable_x86_64
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = True
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/dl.google.com_linux_chrome_rpm_stable_x86_64/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/dl.google.com_linux_chrome_rpm_stable_x86_64/gpgdir
gpgkey = 
hdrdir = /var/cache/yum/x86_64/7/dl.google.com_linux_chrome_rpm_stable_x86_64/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = added from: http://dl.google.com/linux/chrome/rpm/stable/x86_64
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/dl.google.com_linux_chrome_rpm_stable_x86_64
pkgdir = /var/cache/yum/x86_64/7/dl.google.com_linux_chrome_rpm_stable_x86_64/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = dl.google.com_linux_chrome_rpm_stable_x86_64
ui_repoid_vars = releasever,
   basearch
username = 

============================================================ repo: epel ============================================================
[epel]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = 
cache = 0
cachedir = /var/cache/yum/x86_64/7/epel
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = True
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/epel/gpgcadir
gpgcakey = 
gpgcheck = True
gpgdir = /var/lib/yum/repos/x86_64/7/epel/gpgdir
gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-EPEL-7
hdrdir = /var/cache/yum/x86_64/7/epel/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = https://mirrors.fedoraproject.org/metalink?repo=epel-7&arch=x86_64
mirrorlist_expire = 86400
name = Extra Packages for Enterprise Linux 7 - x86_64
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/epel
pkgdir = /var/cache/yum/x86_64/7/epel/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = epel/x86_64
ui_repoid_vars = releasever,
   basearch
username = 

=========================================================== repo: extras ===========================================================
[extras]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = 
cache = 0
cachedir = /var/cache/yum/x86_64/7/extras
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = True
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/extras/gpgcadir
gpgcakey = 
gpgcheck = True
gpgdir = /var/lib/yum/repos/x86_64/7/extras/gpgdir
gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
hdrdir = /var/cache/yum/x86_64/7/extras/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = http://mirrorlist.centos.org/?release=7&arch=x86_64&repo=extras&infra=stock
mirrorlist_expire = 86400
name = CentOS-7 - Extras
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/extras
pkgdir = /var/cache/yum/x86_64/7/extras/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = extras/7/x86_64
ui_repoid_vars = releasever,
   basearch
username = 

======================================================= repo: google-chrome ========================================================
[google-chrome]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = http://dl.google.com/linux/chrome/rpm/stable/x86_64
cache = 0
cachedir = /var/cache/yum/x86_64/7/google-chrome
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = True
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/google-chrome/gpgcadir
gpgcakey = 
gpgcheck = True
gpgdir = /var/lib/yum/repos/x86_64/7/google-chrome/gpgdir
gpgkey = https://dl.google.com/linux/linux_signing_key.pub
hdrdir = /var/cache/yum/x86_64/7/google-chrome/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = google-chrome
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/google-chrome
pkgdir = /var/cache/yum/x86_64/7/google-chrome/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = google-chrome
ui_repoid_vars = releasever,
   basearch
username = 

========================================================= repo: kubernetes =========================================================
[kubernetes]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
cache = 0
cachedir = /var/cache/yum/x86_64/7/kubernetes
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = True
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/kubernetes/gpgcadir
gpgcakey = 
gpgcheck = True
gpgdir = /var/lib/yum/repos/x86_64/7/kubernetes/gpgdir
gpgkey = https://packages.cloud.google.com/yum/doc/yum-key.gpg,
   https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
hdrdir = /var/cache/yum/x86_64/7/kubernetes/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = Kubernetes
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/kubernetes
pkgdir = /var/cache/yum/x86_64/7/kubernetes/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = True
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = kubernetes
ui_repoid_vars = releasever,
   basearch
username = 

====================================================== repo: kubernetes-local ======================================================
[kubernetes-local]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/http0x3A0x2F0x2Fkubernetes.io0x2Fdocs0x2Fsetup0x2Findependent0x2Finstall-kubeadm/https0x3A0x2F0x2Fpackages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
cache = 0
cachedir = /var/cache/yum/x86_64/7/kubernetes-local
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/kubernetes-local/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/kubernetes-local/gpgdir
gpgkey = 
hdrdir = /var/cache/yum/x86_64/7/kubernetes-local/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = Kubernetes offline mirror
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/kubernetes-local
pkgdir = /var/cache/yum/x86_64/7/kubernetes-local/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = kubernetes-local
ui_repoid_vars = releasever,
   basearch
username = 

========================================================== repo: offline ===========================================================
[offline]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///media/cdrom
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline/gpgdir
gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
hdrdir = /var/cache/yum/x86_64/7/offline/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline
pkgdir = /var/cache/yum/x86_64/7/offline/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline
ui_repoid_vars = releasever,
   basearch
username = 

===================================================== repo: offline-centosplus =====================================================
[offline-centosplus]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/centosplus/x86_64/
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-centosplus
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-centosplus/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-centosplus/gpgdir
gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
hdrdir = /var/cache/yum/x86_64/7/offline-centosplus/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline centosplus
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-centosplus
pkgdir = /var/cache/yum/x86_64/7/offline-centosplus/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-centosplus
ui_repoid_vars = releasever,
   basearch
username = 

======================================================= repo: offline-cloud ========================================================
[offline-cloud]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/cloud/x86_64/openstack-pike
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-cloud
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-cloud/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-cloud/gpgdir
gpgkey = 
hdrdir = /var/cache/yum/x86_64/7/offline-cloud/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline cloud openstack-pike
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-cloud
pkgdir = /var/cache/yum/x86_64/7/offline-cloud/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-cloud
ui_repoid_vars = releasever,
   basearch
username = 

========================================================= repo: offline-cr =========================================================
[offline-cr]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/cr/x86_64/
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-cr
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-cr/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-cr/gpgdir
gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
hdrdir = /var/cache/yum/x86_64/7/offline-cr/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline cr
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-cr
pkgdir = /var/cache/yum/x86_64/7/offline-cr/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-cr
ui_repoid_vars = releasever,
   basearch
username = 

======================================================= repo: offline-extras =======================================================
[offline-extras]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/extras/x86_64/
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-extras
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-extras/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-extras/gpgdir
gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
hdrdir = /var/cache/yum/x86_64/7/offline-extras/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline extras
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-extras
pkgdir = /var/cache/yum/x86_64/7/offline-extras/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-extras
ui_repoid_vars = releasever,
   basearch
username = 

===================================================== repo: offline-fasttrack ======================================================
[offline-fasttrack]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/fasttrack/x86_64/
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-fasttrack
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-fasttrack/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-fasttrack/gpgdir
gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
hdrdir = /var/cache/yum/x86_64/7/offline-fasttrack/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline fasttrack
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-fasttrack
pkgdir = /var/cache/yum/x86_64/7/offline-fasttrack/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-fasttrack
ui_repoid_vars = releasever,
   basearch
username = 

======================================================== repo: offline-kvm =========================================================
[offline-kvm]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/virt/x86_64/kvm-common
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-kvm
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-kvm/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-kvm/gpgdir
gpgkey = 
hdrdir = /var/cache/yum/x86_64/7/offline-kvm/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline kvm
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-kvm
pkgdir = /var/cache/yum/x86_64/7/offline-kvm/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-kvm
ui_repoid_vars = releasever,
   basearch
username = 

====================================================== repo: offline-libvirt =======================================================
[offline-libvirt]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/virt/x86_64/libvirt-latest
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-libvirt
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-libvirt/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-libvirt/gpgdir
gpgkey = 
hdrdir = /var/cache/yum/x86_64/7/offline-libvirt/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline libvirt
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-libvirt
pkgdir = /var/cache/yum/x86_64/7/offline-libvirt/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-libvirt
ui_repoid_vars = releasever,
   basearch
username = 

========================================================= repo: offline-os =========================================================
[offline-os]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/os/x86_64/
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-os
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-os/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-os/gpgdir
gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
hdrdir = /var/cache/yum/x86_64/7/offline-os/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline os
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-os
pkgdir = /var/cache/yum/x86_64/7/offline-os/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-os
ui_repoid_vars = releasever,
   basearch
username = 

====================================================== repo: offline-ovirt-41 ======================================================
[offline-ovirt-41]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/virt/x86_64/ovirt-4.1
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-ovirt-41
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-ovirt-41/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-ovirt-41/gpgdir
gpgkey = 
hdrdir = /var/cache/yum/x86_64/7/offline-ovirt-41/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline ovirt 4.1
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-ovirt-41
pkgdir = /var/cache/yum/x86_64/7/offline-ovirt-41/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-ovirt-41
ui_repoid_vars = releasever,
   basearch
username = 

======================================================== repo: offline-paas ========================================================
[offline-paas]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/paas/x86_64/openshift-origin36
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-paas
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-paas/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-paas/gpgdir
gpgkey = 
hdrdir = /var/cache/yum/x86_64/7/offline-paas/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline paas 3.6
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-paas
pkgdir = /var/cache/yum/x86_64/7/offline-paas/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-paas
ui_repoid_vars = releasever,
   basearch
username = 

================================================= repo: offline-storage-ceph-jewel =================================================
[offline-storage-ceph-jewel]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/storage/x86_64/ceph-jewel
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-storage-ceph-jewel
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-storage-ceph-jewel/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-storage-ceph-jewel/gpgdir
gpgkey = 
hdrdir = /var/cache/yum/x86_64/7/offline-storage-ceph-jewel/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline storage ceph-jewel
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-storage-ceph-jewel
pkgdir = /var/cache/yum/x86_64/7/offline-storage-ceph-jewel/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-storage-ceph-jewel
ui_repoid_vars = releasever,
   basearch
username = 

================================================= repo: offline-storage-gluster-38 =================================================
[offline-storage-gluster-38]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/storage/x86_64/gluster-3.8
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-storage-gluster-38
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-storage-gluster-38/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-storage-gluster-38/gpgdir
gpgkey = 
hdrdir = /var/cache/yum/x86_64/7/offline-storage-gluster-38/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline storage gluster-38
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-storage-gluster-38
pkgdir = /var/cache/yum/x86_64/7/offline-storage-gluster-38/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-storage-gluster-38
ui_repoid_vars = releasever,
   basearch
username = 

====================================================== repo: offline-updates =======================================================
[offline-updates]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = file:///vagrant_f/16-mirror/centos/rsync0x3A0x2F0x2Fmirror0x2Ecentos0x2Eorg0x2Fcentos/7/updates/x86_64/
cache = 0
cachedir = /var/cache/yum/x86_64/7/offline-updates
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = 1
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/offline-updates/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/offline-updates/gpgdir
gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
hdrdir = /var/cache/yum/x86_64/7/offline-updates/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = CentOS-7 - Offline updates
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/offline-updates
pkgdir = /var/cache/yum/x86_64/7/offline-updates/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = offline-updates
ui_repoid_vars = releasever,
   basearch
username = 

======================================================= repo: openshift-deps =======================================================
[openshift-deps]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = https://mirror.openshift.com/pub/openshift-v3/dependencies/centos7/x86_64/
cache = 0
cachedir = /var/cache/yum/x86_64/7/openshift-deps
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = True
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/openshift-deps/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/openshift-deps/gpgdir
gpgkey = 
hdrdir = /var/cache/yum/x86_64/7/openshift-deps/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = openshift-deps
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/openshift-deps
pkgdir = /var/cache/yum/x86_64/7/openshift-deps/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = openshift-deps
ui_repoid_vars = releasever,
   basearch
username = 

===================================================== repo: origin-deps-rhel7 ======================================================
[origin-deps-rhel7]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = http://mirror.openshift.com/pub/openshift-origin/nightly/rhel-7/dependencies/x86_64/
cache = 0
cachedir = /var/cache/yum/x86_64/7/origin-deps-rhel7
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = True
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/origin-deps-rhel7/gpgcadir
gpgcakey = 
gpgcheck = False
gpgdir = /var/lib/yum/repos/x86_64/7/origin-deps-rhel7/gpgdir
gpgkey = 
hdrdir = /var/cache/yum/x86_64/7/origin-deps-rhel7/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = 
mirrorlist_expire = 86400
name = OpenShift Origin Dependency repo for RHEL 7
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/origin-deps-rhel7
pkgdir = /var/cache/yum/x86_64/7/origin-deps-rhel7/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = origin-deps-rhel7
ui_repoid_vars = releasever,
   basearch
username = 

========================================================== repo: updates ===========================================================
[updates]
async = True
bandwidth = 0
base_persistdir = /var/lib/yum/repos/x86_64/7
baseurl = 
cache = 0
cachedir = /var/cache/yum/x86_64/7/updates
check_config_file_age = True
cost = 1000
deltarpm_metadata_percentage = 100
deltarpm_percentage = 
enabled = True
enablegroups = True
exclude = 
failovermethod = priority
gpgcadir = /var/lib/yum/repos/x86_64/7/updates/gpgcadir
gpgcakey = 
gpgcheck = True
gpgdir = /var/lib/yum/repos/x86_64/7/updates/gpgdir
gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
hdrdir = /var/cache/yum/x86_64/7/updates/headers
http_caching = all
includepkgs = 
ip_resolve = 
keepalive = True
keepcache = False
mddownloadpolicy = sqlite
mdpolicy = group:small
mediaid = 
metadata_expire = 9999999
metadata_expire_filter = read-only:present
metalink = 
minrate = 0
mirrorlist = http://mirrorlist.centos.org/?release=7&arch=x86_64&repo=updates&infra=stock
mirrorlist_expire = 86400
name = CentOS-7 - Updates
old_base_cache_dir = 
password = 
persistdir = /var/lib/yum/repos/x86_64/7/updates
pkgdir = /var/cache/yum/x86_64/7/updates/packages
proxy = False
proxy_dict = 
proxy_password = 
proxy_username = 
repo_gpgcheck = False
retries = 20
skip_if_unavailable = False
ssl_check_cert_permissions = True
sslcacert = 
sslclientcert = 
sslclientkey = 
sslverify = True
throttle = 0
timeout = 120.0
ui_id = updates/7/x86_64
ui_repoid_vars = releasever,
   basearch
username = 

```

### Selinux

It is not enforcing
```
[vagrant@openshiftdev ~]$ sudo setenforce 0
[vagrant@openshiftdev ~]$ sudo getenforce
Permissive
```

### Issue

如果docker版本时1.10.x, 需使用命令选项 --skip-preflight-checks
```
[vagrant@openshiftdev ~]$ sudo kubeadm init --apiserver-bind-port 443 --kubernetes-version v1.7.4
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[init] Using Kubernetes version: v1.7.4
[init] Using Authorization modes: [Node RBAC]
[preflight] Running pre-flight checks
[preflight] The system verification failed. Printing the output from the verification:
OS: Linux
KERNEL_VERSION: 3.10.0-229.el7.x86_64
CONFIG_NAMESPACES: enabled
CONFIG_NET_NS: enabled
CONFIG_PID_NS: enabled
CONFIG_IPC_NS: enabled
CONFIG_UTS_NS: enabled
CONFIG_CGROUPS: enabled
CONFIG_CGROUP_CPUACCT: enabled
CONFIG_CGROUP_DEVICE: enabled
CONFIG_CGROUP_FREEZER: enabled
CONFIG_CGROUP_SCHED: enabled
CONFIG_CPUSETS: enabled
CONFIG_MEMCG: enabled
CONFIG_INET: enabled
CONFIG_EXT4_FS: enabled (as module)
CONFIG_PROC_FS: enabled
CONFIG_NETFILTER_XT_TARGET_REDIRECT: enabled (as module)
CONFIG_NETFILTER_XT_MATCH_COMMENT: enabled (as module)
CONFIG_OVERLAY_FS: enabled (as module)
CONFIG_AUFS_FS: not set - Required for aufs.
CONFIG_BLK_DEV_DM: enabled (as module)
CGROUPS_CPU: enabled
CGROUPS_CPUACCT: enabled
CGROUPS_CPUSET: enabled
CGROUPS_DEVICES: enabled
CGROUPS_FREEZER: enabled
CGROUPS_MEMORY: enabled
DOCKER_VERSION: 1.10.3
[preflight] WARNING: kubelet service is not enabled, please run 'systemctl enable kubelet.service'
[preflight] WARNING: Running with swap on is not supported. Please disable swap or set kubelet's --fail-swap-on flag to false.
[preflight] Some fatal errors occurred:
	unsupported docker version: 1.10.3
	/proc/sys/net/bridge/bridge-nf-call-iptables contents are not set to 1
[preflight] If you know what you are doing, you can skip pre-flight checks with `--skip-preflight-checks`
```

Swap is not supported
```
[vagrant@openshiftdev ~]$ sudo kubeadm init --apiserver-bind-port 443 --kubernetes-version v1.7.4 --skip-preflight-checks
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[init] Using Kubernetes version: v1.7.4
[init] Using Authorization modes: [Node RBAC]
[preflight] Skipping pre-flight checks
[kubeadm] WARNING: starting in 1.8, tokens expire after 24 hours by default (if you require a non-expiring token use --token-ttl 0)
[certificates] Generated ca certificate and key.
[certificates] Generated apiserver certificate and key.
[certificates] apiserver serving cert is signed for DNS names [openshiftdev.local kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.0.2.15]
[certificates] Generated apiserver-kubelet-client certificate and key.
[certificates] Generated sa key and public key.
[certificates] Generated front-proxy-ca certificate and key.
[certificates] Generated front-proxy-client certificate and key.
[certificates] Valid certificates and keys now exist in "/etc/kubernetes/pki"
[kubeconfig] Wrote KubeConfig file to disk: "admin.conf"
[kubeconfig] Wrote KubeConfig file to disk: "kubelet.conf"
[kubeconfig] Wrote KubeConfig file to disk: "controller-manager.conf"
[kubeconfig] Wrote KubeConfig file to disk: "scheduler.conf"
[controlplane] Wrote Static Pod manifest for component kube-apiserver to "/etc/kubernetes/manifests/kube-apiserver.yaml"
[controlplane] Wrote Static Pod manifest for component kube-controller-manager to "/etc/kubernetes/manifests/kube-controller-manager.yaml"
[controlplane] Wrote Static Pod manifest for component kube-scheduler to "/etc/kubernetes/manifests/kube-scheduler.yaml"
[etcd] Wrote Static Pod manifest for a local etcd instance to "/etc/kubernetes/manifests/etcd.yaml"
[init] Waiting for the kubelet to boot up the control plane as Static Pods from directory "/etc/kubernetes/manifests"
[init] This often takes around a minute; or longer if the control plane images have to be pulled.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz' failed with error: Get http://localhost:10255/healthz: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz' failed with error: Get http://localhost:10255/healthz: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz' failed with error: Get http://localhost:10255/healthz: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz/syncloop' failed with error: Get http://localhost:10255/healthz/syncloop: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz/syncloop' failed with error: Get http://localhost:10255/healthz/syncloop: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz/syncloop' failed with error: Get http://localhost:10255/healthz/syncloop: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz' failed with error: Get http://localhost:10255/healthz: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz/syncloop' failed with error: Get http://localhost:10255/healthz/syncloop: dial tcp [::1]:10255: getsockopt: connection refused.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10255/healthz' failed with error: Get http://localhost:10255/healthz: dial tcp [::1]:10255: getsockopt: connection refused.

Unfortunately, an error has occurred:
	timed out waiting for the condition

This error is likely caused by that:
	- The kubelet is not running
	- The kubelet is unhealthy due to a misconfiguration of the node in some way (required cgroups disabled)
	- There is no internet connection; so the kubelet can't pull the following control plane images:
		- gcr.io/google_containers/kube-apiserver-amd64:v1.7.4
		- gcr.io/google_containers/kube-controller-manager-amd64:v1.7.4
		- gcr.io/google_containers/kube-scheduler-amd64:v1.7.4

You can troubleshoot this for example with the following commands if you're on a systemd-powered system:
	- 'systemctl status kubelet'
	- 'journalctl -xeu kubelet'
couldn't initialize a Kubernetes cluster
```

```
[vagrant@openshiftdev ~]$ cat /etc/systemd/system/kubelet.service
[Unit]
Description=kubelet: The Kubernetes Node Agent
Documentation=http://kubernetes.io/docs/

[Service]
ExecStart=/usr/bin/kubelet
Restart=always
StartLimitInterval=0
RestartSec=10

[Install]
WantedBy=multi-user.target
```

```
[vagrant@openshiftdev ~]$ cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf 
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_SYSTEM_PODS_ARGS=--pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true"
Environment="KUBELET_NETWORK_ARGS=--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
Environment="KUBELET_DNS_ARGS=--cluster-dns=10.96.0.10 --cluster-domain=cluster.local"
Environment="KUBELET_AUTHZ_ARGS=--authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt"
Environment="KUBELET_CADVISOR_ARGS=--cadvisor-port=0"
Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=systemd"
Environment="KUBELET_CERTIFICATE_ARGS=--rotate-certificates=true --cert-dir=/var/lib/kubelet/pki"
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS
```

```
[vagrant@openshiftdev ~]$ ls /etc/kubernetes/
admin.conf  controller-manager.conf  kubelet.conf  manifests  pki  scheduler.conf
```

```
[vagrant@openshiftdev ~]$ sudo systemctl start kubelet.service
[vagrant@openshiftdev ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: activating (auto-restart) (Result: exit-code) since Sun 2017-11-19 20:44:47 UTC; 6s ago
     Docs: http://kubernetes.io/docs/
  Process: 7102 ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS (code=exited, status=1/FAILURE)
 Main PID: 7102 (code=exited, status=1/FAILURE)

Nov 19 20:44:47 openshiftdev.local systemd[1]: kubelet.service: main process exited, code=exited, status=1/FAILURE
Nov 19 20:44:47 openshiftdev.local systemd[1]: Unit kubelet.service entered failed state.
Nov 19 20:44:47 openshiftdev.local systemd[1]: kubelet.service failed.
```

```
[vagrant@openshiftdev ~]$ sudo journalctl --no-pager --no-tail --pager-end --unit kubelet.service
-- Logs begin at Sat 2017-11-18 19:57:25 UTC, end at Sun 2017-11-19 20:53:17 UTC. --
Nov 19 20:46:13 openshiftdev.local systemd[1]: kubelet.service failed.
Nov 19 20:46:24 openshiftdev.local systemd[1]: kubelet.service holdoff time over, scheduling restart.
Nov 19 20:46:24 openshiftdev.local systemd[1]: Started kubelet: The Kubernetes Node Agent.
Nov 19 20:46:24 openshiftdev.local systemd[1]: Starting kubelet: The Kubernetes Node Agent...
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.284075    7326 feature_gate.go:156] feature gates: map[]
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.284130    7326 controller.go:114] kubelet config controller: starting controller
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.284133    7326 controller.go:118] kubelet config controller: validating combination of defaults and flags
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.289648    7326 client.go:75] Connecting to docker on unix:///var/run/docker.sock
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.289665    7326 client.go:95] Start docker client with request timeout=2m0s
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: W1119 20:46:24.311745    7326 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.315707    7326 feature_gate.go:156] feature gates: map[]
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: W1119 20:46:24.315796    7326 server.go:289] --cloud-provider=auto-detect is deprecated. The desired cloud provider should be set explicitly
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.332361    7326 certificate_manager.go:361] Requesting new certificate.
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: E1119 20:46:24.332909    7326 certificate_manager.go:284] Failed while requesting a signed certificate from the master: cannot create certificate signing request: Post https://10.0.2.15:443/apis/certificates.k8s.io/v1beta1/certificatesigningrequests: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.333284    7326 manager.go:149] cAdvisor running in container: "/sys/fs/cgroup/cpu,cpuacct/system.slice/kubelet.service"
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: W1119 20:46:24.484286    7326 manager.go:157] unable to connect to Rkt api service: rkt: cannot tcp Dial rkt api service: dial tcp [::1]:15441: getsockopt: connection refused
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: W1119 20:46:24.484497    7326 manager.go:166] unable to connect to CRI-O api service: Get http://%2Fvar%2Frun%2Fcrio.sock/info: dial unix /var/run/crio.sock: connect: no such file or directory
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.635854    7326 fs.go:139] Filesystem UUIDs: map[2016-12-05-13-52-39-00:/dev/loop0 429bdb70-bbae-41a0-816f-951c5be317db:/dev/dm-4 4b70357b-d057-4a9a-8df7-0736a701697c:/dev/dm-0 ac4ce984-dbef-430b-9660-53c263197aa4:/dev/dm-1 cd4bd4de-b837-42a2-8daf-81b00b03d824:/dev/sda1 ea7569e9-a5ed-42af-bace-7ff166280bde:/dev/dm-6]
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.635876    7326 fs.go:140] Filesystem partitions: map[docker-253:0-203725749-pool:{mountpoint: major:253 minor:2 fsType:devicemapper blockSize:128} tmpfs:{mountpoint:/dev/shm major:0 minor:17 fsType:tmpfs blockSize:0} /dev/mapper/centos-root:{mountpoint:/ major:253 minor:0 fsType:xfs blockSize:0} /dev/mapper/centos-openshift--xfs--vol--dir:{mountpoint:/mnt/openshift-xfs-vol-dir major:253 minor:4 fsType:xfs blockSize:0} /dev/sda1:{mountpoint:/boot major:8 minor:1 fsType:xfs blockSize:0}]
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.639632    7326 manager.go:216] Machine: {NumCores:1 CpuFrequency:3392494 MemoryCapacity:3976323072 HugePages:[{PageSize:2048 NumPages:0}] MachineID:50e9e11ee2d84dc98dcfc8d70a476ae7 SystemUUID:A6CAFB25-3104-4470-9098-847053BB0273 BootID:b607fb3b-986b-406e-98ad-71592e602be8 Filesystems:[{Device:tmpfs DeviceMajor:0 DeviceMinor:17 Capacity:1988161536 Type:vfs Inodes:485391 HasInodes:true} {Device:/dev/mapper/centos-root DeviceMajor:253 DeviceMinor:0 Capacity:41281146880 Type:vfs Inodes:40333312 HasInodes:true} {Device:/dev/mapper/centos-openshift--xfs--vol--dir DeviceMajor:253 DeviceMinor:4 Capacity:10701766656 Type:vfs Inodes:10461184 HasInodes:true} {Device:/dev/sda1 DeviceMajor:8 DeviceMinor:1 Capacity:520794112 Type:vfs Inodes:512000 HasInodes:true} {Device:docker-253:0-203725749-pool DeviceMajor:253 DeviceMinor:2 Capacity:30094131200 Type:devicemapper Inodes:0 HasInodes:false}] DiskMap:map[253:5:{Name:dm-5 Major:253 Minor:5 Size:30094131200 Scheduler:none} 253:6:{Name:dm-6 Major:253 Minor:6 Size:10737418240 Scheduler:none} 8:0:{Name:sda Major:8 Minor:0 Size:85899345920 Scheduler:cfq} 253:0:{Name:dm-0 Major:253 Minor:0 Size:41301311488 Scheduler:none} 253:1:{Name:dm-1 Major:253 Minor:1 Size:1073741824 Scheduler:none} 253:2:{Name:dm-2 Major:253 Minor:2 Size:30094131200 Scheduler:none} 253:3:{Name:dm-3 Major:253 Minor:3 Size:2189426688 Scheduler:none} 253:4:{Name:dm-4 Major:253 Minor:4 Size:10712252416 Scheduler:none}] NetworkDevices:[{Name:enp0s3 MacAddress:08:00:27:53:92:84 Speed:1000 Mtu:1500} {Name:enp0s8 MacAddress:08:00:27:33:cc:33 Speed:1000 Mtu:1500}] Topology:[{Id:0 Memory:4294500352 Cores:[{Id:0 Threads:[0] Caches:[{Size:32768 Type:Data Level:1} {Size:32768 Type:Instruction Level:1} {Size:262144 Type:Unified Level:2}]}] Caches:[{Size:6291456 Type:Unified Level:3}]}] CloudProvider:Unknown InstanceType:Unknown InstanceID:None}
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.682685    7326 manager.go:222] Version: {KernelVersion:3.10.0-229.el7.x86_64 ContainerOsVersion:CentOS Linux 7 (Core) DockerVersion:1.10.3 DockerAPIVersion:1.22 CadvisorVersion: CadvisorRevision:}
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: I1119 20:46:24.683025    7326 server.go:422] --cgroups-per-qos enabled, but --cgroup-root was not specified.  defaulting to /
Nov 19 20:46:24 openshiftdev.local kubelet[7326]: error: failed to run Kubelet: Running with swap on is not supported, please disable swap! or set --fail-swap-on flag to false. /proc/swaps contained: [Filename                                Type                Size        Used        Priority /dev/dm-1                               partition        1048572        1988        -1]
Nov 19 20:46:24 openshiftdev.local systemd[1]: kubelet.service: main process exited, code=exited, status=1/FAILURE
Nov 19 20:46:24 openshiftdev.local systemd[1]: Unit kubelet.service entered failed state.
### snippets ###
```

```
[vagrant@openshiftdev ~]$ sudo sed -i 's/\(ExecStart=.*ARGS\)/\1 --fail-swap-on=false/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
```

```
[vagrant@openshiftdev ~]$ cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf 
[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf"
Environment="KUBELET_SYSTEM_PODS_ARGS=--pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true"
Environment="KUBELET_NETWORK_ARGS=--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
Environment="KUBELET_DNS_ARGS=--cluster-dns=10.96.0.10 --cluster-domain=cluster.local"
Environment="KUBELET_AUTHZ_ARGS=--authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt"
Environment="KUBELET_CADVISOR_ARGS=--cadvisor-port=0"
Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=systemd"
Environment="KUBELET_CERTIFICATE_ARGS=--rotate-certificates=true --cert-dir=/var/lib/kubelet/pki"
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS --fail-swap-on=false
```

CentOS Cgroup settings other than Ubuntu
```
[vagrant@openshiftdev ~]$ sudo systemctl daemon-reload 
[vagrant@openshiftdev ~]$ sudo systemctl restart kubelet.service
[vagrant@openshiftdev ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: activating (auto-restart) (Result: exit-code) since Sun 2017-11-19 21:08:57 UTC; 6s ago
     Docs: http://kubernetes.io/docs/
  Process: 10596 ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_CERTIFICATE_ARGS $KUBELET_EXTRA_ARGS --fail-swap-on=false (code=exited, status=1/FAILURE)
 Main PID: 10596 (code=exited, status=1/FAILURE)

Nov 19 21:08:57 openshiftdev.local systemd[1]: kubelet.service: main process exited, code=exited, status=1/FAILURE
Nov 19 21:08:57 openshiftdev.local systemd[1]: Unit kubelet.service entered failed state.
Nov 19 21:08:57 openshiftdev.local systemd[1]: kubelet.service failed.
```

[vagrant@openshiftdev ~]$ sudo journalctl --no-pager --no-tail --pager-end --unit kubelet.service
```
### snippets ###
Nov 19 21:09:29 openshiftdev.local systemd[1]: kubelet.service holdoff time over, scheduling restart.
Nov 19 21:09:29 openshiftdev.local systemd[1]: Started kubelet: The Kubernetes Node Agent.
Nov 19 21:09:29 openshiftdev.local systemd[1]: Starting kubelet: The Kubernetes Node Agent...
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.042035   10689 feature_gate.go:156] feature gates: map[]
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.042036   10689 controller.go:114] kubelet config controller: starting controller
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.042036   10689 controller.go:118] kubelet config controller: validating combination of defaults and flags
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.047665   10689 client.go:75] Connecting to docker on unix:///var/run/docker.sock
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.047683   10689 client.go:95] Start docker client with request timeout=2m0s
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: W1119 21:09:30.070241   10689 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.073591   10689 feature_gate.go:156] feature gates: map[]
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: W1119 21:09:30.073677   10689 server.go:289] --cloud-provider=auto-detect is deprecated. The desired cloud provider should be set explicitly
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.090751   10689 certificate_manager.go:361] Requesting new certificate.
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: E1119 21:09:30.091624   10689 certificate_manager.go:284] Failed while requesting a signed certificate from the master: cannot create certificate signing request: Post https://10.0.2.15:443/apis/certificates.k8s.io/v1beta1/certificatesigningrequests: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.091889   10689 manager.go:149] cAdvisor running in container: "/sys/fs/cgroup/cpu,cpuacct/system.slice/kubelet.service"
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: W1119 21:09:30.241561   10689 manager.go:157] unable to connect to Rkt api service: rkt: cannot tcp Dial rkt api service: dial tcp [::1]:15441: getsockopt: connection refused
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: W1119 21:09:30.241634   10689 manager.go:166] unable to connect to CRI-O api service: Get http://%2Fvar%2Frun%2Fcrio.sock/info: dial unix /var/run/crio.sock: connect: no such file or directory
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.399801   10689 fs.go:139] Filesystem UUIDs: map[cd4bd4de-b837-42a2-8daf-81b00b03d824:/dev/sda1 ea7569e9-a5ed-42af-bace-7ff166280bde:/dev/dm-6 2016-12-05-13-52-39-00:/dev/loop0 429bdb70-bbae-41a0-816f-951c5be317db:/dev/dm-4 4b70357b-d057-4a9a-8df7-0736a701697c:/dev/dm-0 ac4ce984-dbef-430b-9660-53c263197aa4:/dev/dm-1]
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.399828   10689 fs.go:140] Filesystem partitions: map[tmpfs:{mountpoint:/dev/shm major:0 minor:17 fsType:tmpfs blockSize:0} /dev/mapper/centos-root:{mountpoint:/ major:253 minor:0 fsType:xfs blockSize:0} /dev/mapper/centos-openshift--xfs--vol--dir:{mountpoint:/mnt/openshift-xfs-vol-dir major:253 minor:4 fsType:xfs blockSize:0} /dev/sda1:{mountpoint:/boot major:8 minor:1 fsType:xfs blockSize:0} docker-253:0-203725749-pool:{mountpoint: major:253 minor:2 fsType:devicemapper blockSize:128}]
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.404009   10689 manager.go:216] Machine: {NumCores:1 CpuFrequency:3392494 MemoryCapacity:3976323072 HugePages:[{PageSize:2048 NumPages:0}] MachineID:50e9e11ee2d84dc98dcfc8d70a476ae7 SystemUUID:A6CAFB25-3104-4470-9098-847053BB0273 BootID:b607fb3b-986b-406e-98ad-71592e602be8 Filesystems:[{Device:docker-253:0-203725749-pool DeviceMajor:253 DeviceMinor:2 Capacity:30094131200 Type:devicemapper Inodes:0 HasInodes:false} {Device:tmpfs DeviceMajor:0 DeviceMinor:17 Capacity:1988161536 Type:vfs Inodes:485391 HasInodes:true} {Device:/dev/mapper/centos-root DeviceMajor:253 DeviceMinor:0 Capacity:41281146880 Type:vfs Inodes:40333312 HasInodes:true} {Device:/dev/mapper/centos-openshift--xfs--vol--dir DeviceMajor:253 DeviceMinor:4 Capacity:10701766656 Type:vfs Inodes:10461184 HasInodes:true} {Device:/dev/sda1 DeviceMajor:8 DeviceMinor:1 Capacity:520794112 Type:vfs Inodes:512000 HasInodes:true}] DiskMap:map[253:5:{Name:dm-5 Major:253 Minor:5 Size:30094131200 Scheduler:none} 253:6:{Name:dm-6 Major:253 Minor:6 Size:10737418240 Scheduler:none} 8:0:{Name:sda Major:8 Minor:0 Size:85899345920 Scheduler:cfq} 253:0:{Name:dm-0 Major:253 Minor:0 Size:41301311488 Scheduler:none} 253:1:{Name:dm-1 Major:253 Minor:1 Size:1073741824 Scheduler:none} 253:2:{Name:dm-2 Major:253 Minor:2 Size:30094131200 Scheduler:none} 253:3:{Name:dm-3 Major:253 Minor:3 Size:2189426688 Scheduler:none} 253:4:{Name:dm-4 Major:253 Minor:4 Size:10712252416 Scheduler:none}] NetworkDevices:[{Name:enp0s3 MacAddress:08:00:27:53:92:84 Speed:1000 Mtu:1500} {Name:enp0s8 MacAddress:08:00:27:33:cc:33 Speed:1000 Mtu:1500}] Topology:[{Id:0 Memory:4294500352 Cores:[{Id:0 Threads:[0] Caches:[{Size:32768 Type:Data Level:1} {Size:32768 Type:Instruction Level:1} {Size:262144 Type:Unified Level:2}]}] Caches:[{Size:6291456 Type:Unified Level:3}]}] CloudProvider:Unknown InstanceType:Unknown InstanceID:None}
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.464857   10689 manager.go:222] Version: {KernelVersion:3.10.0-229.el7.x86_64 ContainerOsVersion:CentOS Linux 7 (Core) DockerVersion:1.10.3 DockerAPIVersion:1.22 CadvisorVersion: CadvisorRevision:}
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.465187   10689 server.go:422] --cgroups-per-qos enabled, but --cgroup-root was not specified.  defaulting to /
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.466574   10689 container_manager_linux.go:252] container manager verified user specified cgroup-root exists: /
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.466629   10689 container_manager_linux.go:257] Creating Container Manager object based on Node Config: {RuntimeCgroupsName: SystemCgroupsName: KubeletCgroupsName: ContainerRuntime:docker CgroupsPerQOS:true CgroupRoot:/ CgroupDriver:systemd ProtectKernelDefaults:false NodeAllocatableConfig:{KubeReservedCgroupName: SystemReservedCgroupName: EnforceNodeAllocatable:map[pods:{}] KubeReserved:map[] SystemReserved:map[] HardEvictionThresholds:[{Signal:memory.available Operator:LessThan Value:{Quantity:100Mi Percentage:0} GracePeriod:0s MinReclaim:<nil>} {Signal:nodefs.available Operator:LessThan Value:{Quantity:<nil> Percentage:0.1} GracePeriod:0s MinReclaim:<nil>} {Signal:nodefs.inodesFree Operator:LessThan Value:{Quantity:<nil> Percentage:0.05} GracePeriod:0s MinReclaim:<nil>}]} ExperimentalQOSReserved:map[] ExperimentalCPUManagerPolicy:none ExperimentalCPUManagerReconcilePeriod:10s}
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.466719   10689 container_manager_linux.go:288] Creating device plugin handler: false
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.466807   10689 kubelet.go:273] Adding manifest file: /etc/kubernetes/manifests
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.466837   10689 kubelet.go:283] Watching apiserver
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: E1119 21:09:30.511907   10689 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.0.2.15:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: E1119 21:09:30.511964   10689 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.0.2.15:443/api/v1/services?resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: E1119 21:09:30.511991   10689 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.0.2.15:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: W1119 21:09:30.514488   10689 kubelet_network.go:69] Hairpin mode set to "promiscuous-bridge" but kubenet is not enabled, falling back to "hairpin-veth"
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.514501   10689 kubelet.go:517] Hairpin mode set to "hairpin-veth"
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: W1119 21:09:30.514530   10689 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: W1119 21:09:30.537820   10689 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: W1119 21:09:30.540494   10689 cni.go:196] Unable to update cni config: No networks found in /etc/cni/net.d
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: I1119 21:09:30.540507   10689 docker_service.go:207] Docker cri networking managed by cni
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: W1119 21:09:30.650250   10689 docker_service.go:216] No cgroup driver is set in Docker
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: W1119 21:09:30.650322   10689 docker_service.go:217] Falling back to use the default driver: "cgroupfs"
Nov 19 21:09:30 openshiftdev.local kubelet[10689]: error: failed to run Kubelet: failed to create kubelet: misconfiguration: kubelet cgroup driver: "systemd" is different from docker cgroup driver: "cgroupfs"
Nov 19 21:09:30 openshiftdev.local systemd[1]: kubelet.service: main process exited, code=exited, status=1/FAILURE
Nov 19 21:09:30 openshiftdev.local systemd[1]: Unit kubelet.service entered failed state.
Nov 19 21:09:30 openshiftdev.local systemd[1]: kubelet.service failed.
```

```
[vagrant@openshiftdev ~]$ sudo sed -i 's/\(Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=systemd"\)/#\1\nEnvironment="KUBELET_CGROUP_ARGS=--cgroup-driver=cgroupfs"/' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf
```

```
[vagrant@openshiftdev ~]$ sudo systemctl daemon-reload 
[vagrant@openshiftdev ~]$ sudo systemctl restart kubelet.service
[vagrant@openshiftdev ~]$ sudo systemctl -l status kubelet.service
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; disabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Sun 2017-11-19 21:16:27 UTC; 3s ago
     Docs: http://kubernetes.io/docs/
 Main PID: 11932 (kubelet)
   Memory: 25.0M
   CGroup: /system.slice/kubelet.service
           └─11932 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true --network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin --cluster-dns=10.96.0.10 --cluster-domain=cluster.local --authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt --cadvisor-port=0 --cgroup-driver=cgroupfs --rotate-certificates=true --cert-dir=/var/lib/kubelet/pki --fail-swap-on=false

Nov 19 21:16:29 openshiftdev.local kubelet[11932]: E1119 21:16:29.910453   11932 kubelet_node_status.go:107] Unable to register node "openshiftdev.local" with API server: Post https://10.0.2.15:443/api/v1/nodes: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:30 openshiftdev.local kubelet[11932]: E1119 21:16:30.382097   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.0.2.15:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:30 openshiftdev.local kubelet[11932]: E1119 21:16:30.405038   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.0.2.15:443/api/v1/services?resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:30 openshiftdev.local kubelet[11932]: E1119 21:16:30.417405   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.0.2.15:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:30 openshiftdev.local kubelet[11932]: I1119 21:16:30.710880   11932 kubelet_node_status.go:280] Setting node annotation to enable volume controller attach/detach
Nov 19 21:16:30 openshiftdev.local kubelet[11932]: I1119 21:16:30.811518   11932 kubelet_node_status.go:83] Attempting to register node openshiftdev.local
Nov 19 21:16:30 openshiftdev.local kubelet[11932]: E1119 21:16:30.811775   11932 kubelet_node_status.go:107] Unable to register node "openshiftdev.local" with API server: Post https://10.0.2.15:443/api/v1/nodes: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:31 openshiftdev.local kubelet[11932]: E1119 21:16:31.383195   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.0.2.15:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:31 openshiftdev.local kubelet[11932]: E1119 21:16:31.405884   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.0.2.15:443/api/v1/services?resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:16:31 openshiftdev.local kubelet[11932]: E1119 21:16:31.418781   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.0.2.15:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
```

```
[vagrant@openshiftdev ~]$ sudo journalctl --no-pager --no-tail --pager-end --unit kubelet.service | grep E1119 | tail
Nov 19 21:30:00 openshiftdev.local kubelet[11932]: E1119 21:30:00.616706   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.0.2.15:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:00 openshiftdev.local kubelet[11932]: E1119 21:30:00.619308   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.0.2.15:443/api/v1/services?resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:01 openshiftdev.local kubelet[11932]: E1119 21:30:01.616746   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.0.2.15:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:01 openshiftdev.local kubelet[11932]: E1119 21:30:01.623751   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.0.2.15:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:01 openshiftdev.local kubelet[11932]: E1119 21:30:01.625465   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.0.2.15:443/api/v1/services?resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:02 openshiftdev.local kubelet[11932]: E1119 21:30:02.624423   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/config/apiserver.go:47: Failed to list *v1.Pod: Get https://10.0.2.15:443/api/v1/pods?fieldSelector=spec.nodeName%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:02 openshiftdev.local kubelet[11932]: E1119 21:30:02.625036   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:422: Failed to list *v1.Node: Get https://10.0.2.15:443/api/v1/nodes?fieldSelector=metadata.name%3Dopenshiftdev.local&resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:02 openshiftdev.local kubelet[11932]: E1119 21:30:02.628671   11932 reflector.go:205] k8s.io/kubernetes/pkg/kubelet/kubelet.go:413: Failed to list *v1.Service: Get https://10.0.2.15:443/api/v1/services?resourceVersion=0: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:02 openshiftdev.local kubelet[11932]: E1119 21:30:02.951025   11932 kubelet.go:1612] Failed creating a mirror pod for "kube-controller-manager-openshiftdev.local_kube-system(68a0ca8ea0a05b769d4c2c3355e47ae5)": Post https://10.0.2.15:443/api/v1/namespaces/kube-system/pods: dial tcp 10.0.2.15:443: getsockopt: connection refused
Nov 19 21:30:03 openshiftdev.local kubelet[11932]: E1119 21:30:03.146769   11932 kubelet.go:2095] Container runtime network not ready: NetworkReady=false reason:NetworkPluginNotReady message:docker: network plugin is not ready: cni config uninitialized
```

```
[vagrant@openshiftdev ~]$ sudo kubeadm init --apiserver-advertise-address 10.64.33.82 --apiserver-bind-port 443 --kubernetes-version v1.7.4 --skip-preflight-checks 
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[init] Using Kubernetes version: v1.7.4
[init] Using Authorization modes: [Node RBAC]
[preflight] Skipping pre-flight checks
[kubeadm] WARNING: starting in 1.8, tokens expire after 24 hours by default (if you require a non-expiring token use --token-ttl 0)
[certificates] Using the existing ca certificate and key.
[certificates] Using the existing apiserver certificate and key.
[certificates] Using the existing apiserver-kubelet-client certificate and key.
[certificates] Using the existing sa key.
[certificates] Using the existing front-proxy-ca certificate and key.
[certificates] Using the existing front-proxy-client certificate and key.
[certificates] Valid certificates and keys now exist in "/etc/kubernetes/pki"
a kubeconfig file "/etc/kubernetes/admin.conf" exists already but has got the wrong API Server URL
```

### Client

Cli
```
[vagrant@openshiftdev ~]$ kubectl version --client
Client Version: version.Info{Major:"1", Minor:"8", GitVersion:"v1.8.2", GitCommit:"bdaeafa71f6c7c04636251031f93464384d54963", GitTreeState:"clean", BuildDate:"2017-10-24T19:48:57Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"linux/amd64"}
```