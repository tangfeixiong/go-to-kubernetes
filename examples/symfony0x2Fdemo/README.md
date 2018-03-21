## Development

### Dev tools
```
[vagrant@kubedev-172-17-4-59 github.com]$ sudo dnf search php
Failed to set locale, defaulting to C
Last metadata expiration check: 1:30:03 ago on Sun Mar 18 15:35:57 2018.
============================================================ Name & Summary Matched: php =============================================================
php.x86_64 : PHP scripting language for creating dynamic web sites
<<< snippets >>>
```

```
[vagrant@kubedev-172-17-4-59 github.com]$ sudo dnf install php
Failed to set locale, defaulting to C
Last metadata expiration check: 1:37:00 ago on Sun Mar 18 15:35:57 2018.
Dependencies resolved.
======================================================================================================================================================
 Package                                   Arch                          Version                                 Repository                      Size
======================================================================================================================================================
Installing:
 php                                       x86_64                        7.1.15-1.fc26                           updates                        2.8 M
Installing dependencies:
 apr                                       x86_64                        1.6.3-1.fc26                            updates                        118 k
 apr-util                                  x86_64                        1.5.4-6.fc26                            updates                         95 k
 fedora-logos-httpd                        noarch                        26.0.1-1.fc26                           fedora                          33 k
 httpd                                     x86_64                        2.4.29-1.fc26                           updates                        1.3 M
 httpd-filesystem                          noarch                        2.4.29-1.fc26                           updates                         28 k
 httpd-tools                               x86_64                        2.4.29-1.fc26                           updates                         88 k
 mailcap                                   noarch                        2.1.48-1.fc26                           fedora                          37 k
 mod_http2                                 x86_64                        1.10.13-1.fc26                          updates                        148 k
 php-cli                                   x86_64                        7.1.15-1.fc26                           updates                        4.2 M
 php-common                                x86_64                        7.1.15-1.fc26                           updates                        1.0 M
 php-json                                  x86_64                        7.1.15-1.fc26                           updates                         70 k

Transaction Summary
======================================================================================================================================================
Install  12 Packages

Total download size: 9.9 M
Installed size: 35 M
Is this ok [y/N]: y
Downloading Packages:
(1/12): php-common-7.1.15-1.fc26.x86_64.rpm                                                                           837 kB/s | 1.0 MB     00:01    
(2/12): php-json-7.1.15-1.fc26.x86_64.rpm                                                                             376 kB/s |  70 kB     00:00    
(3/12): php-7.1.15-1.fc26.x86_64.rpm                                                                                  1.7 MB/s | 2.8 MB     00:01    
(4/12): httpd-filesystem-2.4.29-1.fc26.noarch.rpm                                                                     310 kB/s |  28 kB     00:00    
(5/12): httpd-tools-2.4.29-1.fc26.x86_64.rpm                                                                          615 kB/s |  88 kB     00:00    
(6/12): mailcap-2.1.48-1.fc26.noarch.rpm                                                                              270 kB/s |  37 kB     00:00    
(7/12): apr-1.6.3-1.fc26.x86_64.rpm                                                                                   661 kB/s | 118 kB     00:00    
(8/12): apr-util-1.5.4-6.fc26.x86_64.rpm                                                                              444 kB/s |  95 kB     00:00    
(9/12): httpd-2.4.29-1.fc26.x86_64.rpm                                                                                1.2 MB/s | 1.3 MB     00:01    
(10/12): mod_http2-1.10.13-1.fc26.x86_64.rpm                                                                          755 kB/s | 148 kB     00:00    
(11/12): fedora-logos-httpd-26.0.1-1.fc26.noarch.rpm                                                                  300 kB/s |  33 kB     00:00    
(12/12): php-cli-7.1.15-1.fc26.x86_64.rpm                                                                             1.5 MB/s | 4.2 MB     00:02    
------------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                                 1.5 MB/s | 9.9 MB     00:06     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                              1/1 
  Installing       : php-json-7.1.15-1.fc26.x86_64                                                                                               1/12 
  Installing       : php-common-7.1.15-1.fc26.x86_64                                                                                             2/12 
  Installing       : apr-1.6.3-1.fc26.x86_64                                                                                                     3/12 
  Running scriptlet: apr-1.6.3-1.fc26.x86_64                                                                                                     3/12 
  Installing       : apr-util-1.5.4-6.fc26.x86_64                                                                                                4/12 
  Running scriptlet: apr-util-1.5.4-6.fc26.x86_64                                                                                                4/12 
  Running scriptlet: httpd-filesystem-2.4.29-1.fc26.noarch                                                                                       5/12 
  Installing       : httpd-filesystem-2.4.29-1.fc26.noarch                                                                                       5/12 
  Installing       : httpd-tools-2.4.29-1.fc26.x86_64                                                                                            6/12 
  Installing       : php-cli-7.1.15-1.fc26.x86_64                                                                                                7/12 
  Installing       : fedora-logos-httpd-26.0.1-1.fc26.noarch                                                                                     8/12 
  Installing       : mailcap-2.1.48-1.fc26.noarch                                                                                                9/12 
  Installing       : mod_http2-1.10.13-1.fc26.x86_64                                                                                            10/12 
  Installing       : httpd-2.4.29-1.fc26.x86_64                                                                                                 11/12 
  Running scriptlet: httpd-2.4.29-1.fc26.x86_64                                                                                                 11/12 
  Installing       : php-7.1.15-1.fc26.x86_64                                                                                                   12/12 
  Running scriptlet: httpd-2.4.29-1.fc26.x86_64                                                                                                 12/12 
  Running scriptlet: php-7.1.15-1.fc26.x86_64                                                                                                   12/12 
  Verifying        : php-7.1.15-1.fc26.x86_64                                                                                                    1/12 
  Verifying        : php-cli-7.1.15-1.fc26.x86_64                                                                                                2/12 
  Verifying        : php-common-7.1.15-1.fc26.x86_64                                                                                             3/12 
  Verifying        : php-json-7.1.15-1.fc26.x86_64                                                                                               4/12 
  Verifying        : httpd-2.4.29-1.fc26.x86_64                                                                                                  5/12 
  Verifying        : httpd-filesystem-2.4.29-1.fc26.noarch                                                                                       6/12 
  Verifying        : httpd-tools-2.4.29-1.fc26.x86_64                                                                                            7/12 
  Verifying        : mailcap-2.1.48-1.fc26.noarch                                                                                                8/12 
  Verifying        : apr-1.6.3-1.fc26.x86_64                                                                                                     9/12 
  Verifying        : apr-util-1.5.4-6.fc26.x86_64                                                                                               10/12 
  Verifying        : mod_http2-1.10.13-1.fc26.x86_64                                                                                            11/12 
  Verifying        : fedora-logos-httpd-26.0.1-1.fc26.noarch                                                                                    12/12 

Installed:
  php.x86_64 7.1.15-1.fc26         apr.x86_64 1.6.3-1.fc26                apr-util.x86_64 1.5.4-6.fc26      fedora-logos-httpd.noarch 26.0.1-1.fc26 
  httpd.x86_64 2.4.29-1.fc26       httpd-filesystem.noarch 2.4.29-1.fc26  httpd-tools.x86_64 2.4.29-1.fc26  mailcap.noarch 2.1.48-1.fc26            
  mod_http2.x86_64 1.10.13-1.fc26  php-cli.x86_64 7.1.15-1.fc26           php-common.x86_64 7.1.15-1.fc26   php-json.x86_64 7.1.15-1.fc26           

Complete!
```

```
[vagrant@kubedev-172-17-4-59 demo]$ php --version
PHP 7.1.15 (cli) (built: Feb 28 2018 11:19:18) ( NTS )
Copyright (c) 1997-2018 The PHP Group
Zend Engine v3.1.0, Copyright (c) 1998-2018 Zend Technologies
```

Manual Download (for example version 1.6.3)
```
[vagrant@kubedev-172-17-4-59 demo]$ php /Users/fanhongling/Downloads/99-mirror/php/composer.phar --version
Composer version 1.6.3 2018-01-31 16:28:17
```

## Project demo

`git clone https://github.com/symfony/demo`

```
[vagrant@kubedev-172-17-4-59 github.com]$ cd symfony/demo/
```

```
[vagrant@kubedev-172-17-4-59 demo]$ ls
CONTRIBUTING.md  app.json      bin            config            public        templates     var
LICENSE          appveyor.yml  composer.json  package.json      src           tests         webpack.config.js
README.md        assets        composer.lock  phpunit.xml.dist  symfony.lock  translations  yarn.lock
```

Issue
```
[vagrant@kubedev-172-17-4-59 demo]$ php /Users/fanhongling/Downloads/99-mirror/php/composer.phar install
Loading composer repositories with package information
Installing dependencies (including require-dev) from lock file
Your requirements could not be resolved to an installable set of packages.

  Problem 1
    - The requested PHP extension ext-pdo_sqlite * is missing from your system. Install or enable PHP's pdo_sqlite extension.
  Problem 2
    - Installation request for doctrine/dbal v2.6.3 -> satisfiable by doctrine/dbal[v2.6.3].
    - doctrine/dbal v2.6.3 requires ext-pdo * -> the requested PHP extension pdo is missing from your system.
  Problem 3
    - Installation request for doctrine/orm v2.6.1 -> satisfiable by doctrine/orm[v2.6.1].
    - doctrine/orm v2.6.1 requires ext-pdo * -> the requested PHP extension pdo is missing from your system.
  Problem 4
    - Installation request for erusev/parsedown 1.7.1 -> satisfiable by erusev/parsedown[1.7.1].
    - erusev/parsedown 1.7.1 requires ext-mbstring * -> the requested PHP extension mbstring is missing from your system.
  Problem 5
    - Installation request for symfony/framework-bundle v4.0.6 -> satisfiable by symfony/framework-bundle[v4.0.6].
    - symfony/framework-bundle v4.0.6 requires ext-xml * -> the requested PHP extension xml is missing from your system.
  Problem 6
    - Installation request for symfony/security-bundle v4.0.6 -> satisfiable by symfony/security-bundle[v4.0.6].
    - symfony/security-bundle v4.0.6 requires ext-xml * -> the requested PHP extension xml is missing from your system.
  Problem 7
    - Installation request for symfony/debug-bundle v4.0.6 -> satisfiable by symfony/debug-bundle[v4.0.6].
    - symfony/debug-bundle v4.0.6 requires ext-xml * -> the requested PHP extension xml is missing from your system.
  Problem 8
    - doctrine/dbal v2.6.3 requires ext-pdo * -> the requested PHP extension pdo is missing from your system.
    - dama/doctrine-test-bundle v5.0.1 requires doctrine/dbal ~2.5 -> satisfiable by doctrine/dbal[v2.6.3].
    - Installation request for dama/doctrine-test-bundle v5.0.1 -> satisfiable by dama/doctrine-test-bundle[v5.0.1].

  To enable extensions, verify that they are enabled in your .ini files:
    - /etc/php.ini
    - /etc/php.d/20-bz2.ini
    - /etc/php.d/20-calendar.ini
    - /etc/php.d/20-ctype.ini
    - /etc/php.d/20-curl.ini
    - /etc/php.d/20-exif.ini
    - /etc/php.d/20-fileinfo.ini
    - /etc/php.d/20-ftp.ini
    - /etc/php.d/20-gettext.ini
    - /etc/php.d/20-iconv.ini
    - /etc/php.d/20-json.ini
    - /etc/php.d/20-phar.ini
    - /etc/php.d/20-sockets.ini
    - /etc/php.d/20-tokenizer.ini
  You can also run `php --ini` inside terminal to see which files are used by PHP in CLI mode.
```

Refer to https://www.if-not-true-then-false.com/2010/install-apache-php-on-fedora-centos-red-hat-rhel/
```
[vagrant@kubedev-172-17-4-59 demo]$ sudo dnf list | egrep ^php-pdo 
Failed to set locale, defaulting to C
php-pdo.x86_64                           7.1.15-1.fc26                   updates
php-pdo-dblib.x86_64                     7.1.15-1.fc26                   updates
```

```
[vagrant@kubedev-172-17-4-59 demo]$ sudo dnf list | egrep ^php-xml
Failed to set locale, defaulting to C
php-xml.x86_64                           7.1.15-1.fc26                   updates
php-xmlrpc.x86_64                        7.1.15-1.fc26                   updates
php-xmlseclibs.noarch                    1.3.1-6.fc26                    fedora 
```

```
[vagrant@kubedev-172-17-4-59 demo]$ sudo dnf list | egrep ^php-mbstring
Failed to set locale, defaulting to C
php-mbstring.x86_64                      7.1.15-1.fc26                   updates
```

Resolve
```
[vagrant@kubedev-172-17-4-59 demo]$ sudo dnf install php-pdo php-xml php-mbstring
Failed to set locale, defaulting to C
Last metadata expiration check: 2:17:39 ago on Sun Mar 18 15:35:57 2018.
Dependencies resolved.
======================================================================================================================================================
 Package                               Arch                            Version                                 Repository                        Size
======================================================================================================================================================
Installing:
 php-mbstring                          x86_64                          7.1.15-1.fc26                           updates                          591 k
 php-pdo                               x86_64                          7.1.15-1.fc26                           updates                          133 k
 php-xml                               x86_64                          7.1.15-1.fc26                           updates                          219 k
Installing dependencies:
 libxslt                               x86_64                          1.1.29-1.fc26                           fedora                           252 k

Transaction Summary
======================================================================================================================================================
Install  4 Packages

Total download size: 1.2 M
Installed size: 4.5 M
Is this ok [y/N]: y
Downloading Packages:
(1/4): php-pdo-7.1.15-1.fc26.x86_64.rpm                                                                               298 kB/s | 133 kB     00:00    
(2/4): php-xml-7.1.15-1.fc26.x86_64.rpm                                                                               384 kB/s | 219 kB     00:00    
(3/4): libxslt-1.1.29-1.fc26.x86_64.rpm                                                                               319 kB/s | 252 kB     00:00    
(4/4): php-mbstring-7.1.15-1.fc26.x86_64.rpm                                                                          842 kB/s | 591 kB     00:00    
------------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                                 264 kB/s | 1.2 MB     00:04     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                              1/1 
  Installing       : libxslt-1.1.29-1.fc26.x86_64                                                                                                 1/4 
  Running scriptlet: libxslt-1.1.29-1.fc26.x86_64                                                                                                 1/4 
  Installing       : php-xml-7.1.15-1.fc26.x86_64                                                                                                 2/4 
  Installing       : php-mbstring-7.1.15-1.fc26.x86_64                                                                                            3/4 
  Installing       : php-pdo-7.1.15-1.fc26.x86_64                                                                                                 4/4 
  Running scriptlet: php-pdo-7.1.15-1.fc26.x86_64                                                                                                 4/4 
  Verifying        : php-pdo-7.1.15-1.fc26.x86_64                                                                                                 1/4 
  Verifying        : php-xml-7.1.15-1.fc26.x86_64                                                                                                 2/4 
  Verifying        : libxslt-1.1.29-1.fc26.x86_64                                                                                                 3/4 
  Verifying        : php-mbstring-7.1.15-1.fc26.x86_64                                                                                            4/4 

Installed:
  php-mbstring.x86_64 7.1.15-1.fc26        php-pdo.x86_64 7.1.15-1.fc26        php-xml.x86_64 7.1.15-1.fc26        libxslt.x86_64 1.1.29-1.fc26       

Complete!
```

```
[vagrant@kubedev-172-17-4-59 demo]$ php /Users/fanhongling/Downloads/99-mirror/php/composer.phar install
Loading composer repositories with package information
Installing dependencies (including require-dev) from lock file
Package operations: 88 installs, 0 updates, 0 removals
  - Installing ocramius/package-versions (1.3.0): Downloading (100%)         
  - Installing symfony/flex (v1.0.71): Downloading (100%)         

Prefetching 86 packages 🎶
  - Downloading (100%)

  - Installing symfony/polyfill-mbstring (v1.7.0): Loading from cache
  - Installing doctrine/lexer (v1.0.1): Loading from cache
  - Installing doctrine/inflector (v1.3.0): Loading from cache
  - Installing doctrine/collections (v1.5.0): Loading from cache
  - Installing doctrine/cache (v1.7.1): Loading from cache
  - Installing doctrine/annotations (v1.6.0): Loading from cache
  - Installing doctrine/common (v2.8.1): Loading from cache
  - Installing symfony/doctrine-bridge (v4.0.6): Loading from cache
  - Installing doctrine/doctrine-cache-bundle (1.3.2): Loading from cache
  - Installing symfony/routing (v4.0.6): Loading from cache
  - Installing symfony/http-foundation (v4.0.6): Loading from cache
  - Installing symfony/event-dispatcher (v4.0.6): Loading from cache
  - Installing psr/log (1.0.2): Loading from cache
  - Installing symfony/debug (v4.0.6): Loading from cache
  - Installing symfony/http-kernel (v4.0.6): Loading from cache
  - Installing symfony/finder (v4.0.6): Loading from cache
  - Installing symfony/filesystem (v4.0.6): Loading from cache
  - Installing psr/container (1.0.0): Loading from cache
  - Installing symfony/dependency-injection (v4.0.6): Loading from cache
  - Installing symfony/config (v4.0.6): Loading from cache
  - Installing psr/simple-cache (1.0.1): Loading from cache
  - Installing psr/cache (1.0.1): Loading from cache
  - Installing symfony/cache (v4.0.6): Loading from cache
  - Installing symfony/framework-bundle (v4.0.6): Loading from cache
  - Installing symfony/console (v4.0.6): Loading from cache
  - Installing jdorn/sql-formatter (v1.2.17): Loading from cache
  - Installing doctrine/dbal (v2.6.3): Loading from cache
  - Installing doctrine/doctrine-bundle (1.8.1): Loading from cache
  - Installing doctrine/data-fixtures (v1.3.0): Loading from cache
  - Installing doctrine/doctrine-fixtures-bundle (3.0.2): Loading from cache
  - Installing symfony/yaml (v4.0.6): Loading from cache
  - Installing zendframework/zend-eventmanager (3.2.0): Loading from cache
  - Installing zendframework/zend-code (3.3.0): Loading from cache
  - Installing ocramius/proxy-manager (2.1.1): Loading from cache
  - Installing doctrine/migrations (v1.6.2): Loading from cache
  - Installing doctrine/doctrine-migrations-bundle (v1.3.1): Loading from cache
  - Installing doctrine/instantiator (1.1.0): Loading from cache
  - Installing doctrine/orm (v2.6.1): Loading from cache
  - Installing egulias/email-validator (2.1.3): Loading from cache
  - Installing erusev/parsedown (1.7.1): Loading from cache
  - Installing ezyang/htmlpurifier (v4.10.0): Loading from cache
  - Installing sensio/framework-extra-bundle (v5.1.6): Loading from cache
  - Installing composer/ca-bundle (1.1.0): Loading from cache
  - Installing sensiolabs/security-checker (v4.1.8): Loading from cache
  - Installing symfony/asset (v4.0.6): Loading from cache
  - Installing symfony/expression-language (v4.0.6): Loading from cache
  - Installing symfony/inflector (v4.0.6): Loading from cache
  - Installing symfony/property-access (v4.0.6): Loading from cache
  - Installing symfony/options-resolver (v4.0.6): Loading from cache
  - Installing symfony/intl (v4.0.6): Loading from cache
  - Installing symfony/polyfill-intl-icu (v1.7.0): Loading from cache
  - Installing symfony/form (v4.0.6): Loading from cache
  - Installing monolog/monolog (1.23.0): Loading from cache
  - Installing symfony/monolog-bridge (v4.0.6): Loading from cache
  - Installing symfony/monolog-bundle (v3.2.0): Loading from cache
  - Installing symfony/polyfill-apcu (v1.7.0): Loading from cache
  - Installing symfony/security (v4.0.6): Loading from cache
  - Installing symfony/security-bundle (v4.0.6): Loading from cache
  - Installing swiftmailer/swiftmailer (v6.0.2): Loading from cache
  - Installing symfony/swiftmailer-bundle (v3.2.1): Loading from cache
  - Installing symfony/translation (v4.0.6): Loading from cache
  - Installing symfony/validator (v4.0.6): Loading from cache
  - Installing twig/twig (v2.4.6): Loading from cache
  - Installing twig/extensions (v1.5.1): Loading from cache
  - Installing symfony/twig-bridge (v4.0.6): Loading from cache
  - Installing symfony/twig-bundle (v4.0.6): Loading from cache
  - Installing pagerfanta/pagerfanta (v1.0.5): Loading from cache
  - Installing white-october/pagerfanta-bundle (v1.1.2): Loading from cache
  - Installing dama/doctrine-test-bundle (v5.0.1): Loading from cache
  - Installing symfony/stopwatch (v4.0.6): Loading from cache
  - Installing symfony/process (v4.0.6): Loading from cache
  - Installing symfony/polyfill-php72 (v1.7.0): Loading from cache
  - Installing paragonie/random_compat (v2.0.11): Loading from cache
  - Installing symfony/polyfill-php70 (v1.7.0): Loading from cache
  - Installing php-cs-fixer/diff (v1.3.0): Loading from cache
  - Installing composer/semver (1.4.2): Loading from cache
  - Installing friendsofphp/php-cs-fixer (v2.10.4): Loading from cache
  - Installing symfony/dom-crawler (v4.0.6): Loading from cache
  - Installing symfony/browser-kit (v4.0.6): Loading from cache
  - Installing symfony/css-selector (v4.0.6): Loading from cache
  - Installing symfony/var-dumper (v4.0.6): Loading from cache
  - Installing symfony/debug-bundle (v4.0.6): Loading from cache
  - Installing symfony/dotenv (v4.0.6): Loading from cache
  - Installing symfony/phpunit-bridge (v4.0.6): Loading from cache
  - Installing symfony/web-profiler-bundle (v4.0.6): Loading from cache
  - Installing symfony/web-server-bundle (v4.0.6): Loading from cache
Generating autoload files
ocramius/package-versions:  Generating version class...
ocramius/package-versions: ...done generating version class
```

```
[vagrant@kubedev-172-17-4-59 demo]$ ls vendor/
autoload.php  composer  doctrine  erusev  friendsofphp  monolog   pagerfanta  php-cs-fixer  sensio      swiftmailer  twig           zendframework
bin           dama      egulias   ezyang  jdorn         ocramius  paragonie   psr           sensiolabs  symfony      white-october
```

### CI/CD

Run
```
[vagrant@kubedev-172-17-4-59 demo]$ php bin/console server:run

 [OK] Server listening on http://127.0.0.1:8000                                                                         

 // Quit the server with CONTROL-C.                                                                                     

PHP 7.1.15 Development Server started at Sun Mar 18 18:01:54 2018
Listening on http://127.0.0.1:8000
Document root is /Users/fanhongling/Downloads/workspace/src/github.com/symfony/demo/public
Press Ctrl-C to quit.
^C
```

Refer to https://symfony.com/doc/current/setup.html

Run in VM and show in Chrome
```
[vagrant@kubedev-172-17-4-59 demo]$ php bin/console server:run 0.0.0.0:8080

 [OK] Server listening on http://0.0.0.0:8080                                                                           

 // Quit the server with CONTROL-C.                                                                                     

PHP 7.1.15 Development Server started at Sun Mar 18 18:05:56 2018
Listening on http://0.0.0.0:8080
Document root is /Users/fanhongling/Downloads/workspace/src/github.com/symfony/demo/public
Press Ctrl-C to quit.
[Sun Mar 18 18:06:47 2018] 172.17.4.1:59846 [302]: /
[Sun Mar 18 18:06:50 2018] 172.17.4.1:59845 [200]: /zh_CN
[Sun Mar 18 18:06:50 2018] 172.17.4.1:59847 [200]: /build/css/app.css
[Sun Mar 18 18:06:50 2018] 172.17.4.1:59848 [200]: /build/manifest.js
[Sun Mar 18 18:06:50 2018] 172.17.4.1:59849 [200]: /build/js/common.js
[Sun Mar 18 18:06:50 2018] 172.17.4.1:59850 [200]: /build/js/app.js
[Sun Mar 18 18:06:50 2018] 172.17.4.1:59851 [200]: /build/fonts/lato-normal.woff2
[Sun Mar 18 18:06:50 2018] 172.17.4.1:59852 [200]: /build/fonts/lato-bold.woff2
[Sun Mar 18 18:06:50 2018] 172.17.4.1:59853 [200]: /build/fonts/fontawesome-webfont.woff2
[Sun Mar 18 18:06:53 2018] 172.17.4.1:59854 [200]: /_wdt/dcc5dc
```

![屏幕快照 2018-03-18 上午11.07.16.png](./屏幕快照%202018-03-18%20上午11.07.16.png)