# Nightly Build rust-lang

## Git Clone

For example, in VM (Fedora 23)

    [vagrant@localhost github.com]$ git clone --depth=1 https://github.com/rust-lang/rust.git rust-lang/rust
    正克隆到 'rust-lang/rust'...
    remote: Counting objects: 8353, done.
    remote: Compressing objects: 100% (7400/7400), done.
    remote: Total 8353 (delta 4484), reused 2255 (delta 912), pack-reused 0
    接收对象中: 100% (8353/8353), 7.72 MiB | 47.00 KiB/s, 完成.
    处理 delta 中: 100% (4484/4484), 完成.
    检查连接... 完成。
    [vagrant@localhost github.com]$ cd rust-lang/rust/
    [vagrant@localhost rust]$ ls
    COMPILER_TESTS.md  configure  CONTRIBUTING.md  COPYRIGHT  LICENSE-APACHE  LICENSE-MIT  Makefile.in  man  mk  README.md  RELEASES.md  src

## Linux Build

### Prerequisites

Must install GCC or CLang build tool chain

* Centos/Fedora/Redhat series

[Centos guide](http://www.cyberciti.biz/faq/centos-linux-install-gcc-c-c-compiler/) 

    sudo yum update
    sudo yum groupinstall "Development Tools" "Legacy Software Development"

* Debian / Ubuntu / Debian derivatives

[For example](https://www.garron.me/en/go2linux/gnu-gcc-development-tools-linux-fedora-arch-debian.html)

    sudo apt-get update
    sudo apt-get install build-essential

* Arch Linux

    sudo pacman -Sy base-devel

### Hands-on

Mac Book + Virtual Box (VM memory: 4096M)

Fedora 23

    sudo dnf groupinstall "Development Tools"

Configure Makefile

    [vagrant@localhost rust]$ ./configure 
    configure: looking for configure programs
    configure: found program 'cmp'
    configure: found program 'mkdir'
    configure: found program 'printf'
    configure: found program 'cut'
    configure: found program 'head'
    configure: found program 'grep'
    configure: found program 'xargs'
    configure: found program 'cp'
    configure: found program 'find'
    configure: found program 'uname'
    configure: found program 'date'
    configure: found program 'tr'
    configure: found program 'sed'
    configure: found program 'file'
    configure: found program 'make'
    configure: inspecting environment
    configure: recreating config.tmp
    configure: 
    configure: processing ./configure args
    configure: 
    configure: CFG_LOCALSTATEDIR    := /var/lib 
    configure: CFG_SYSCONFDIR       := /etc 
    configure: CFG_DATADIR          := /share 
    configure: CFG_INFODIR          := /share/info 
    configure: CFG_LLVM_ROOT        :=  
    configure: CFG_PYTHON           :=  
    configure: CFG_JEMALLOC_ROOT    :=  
    configure: CFG_BUILD            := x86_64-unknown-linux-gnu 
    configure: CFG_ANDROID_CROSS_PATH :=  
    configure: CFG_I686_LINUX_ANDROID_NDK :=  
    configure: CFG_ARM_LINUX_ANDROIDEABI_NDK :=  
    configure: CFG_ARMV7_LINUX_ANDROIDEABI_NDK :=  
    configure: CFG_AARCH64_LINUX_ANDROID_NDK :=  
    configure: CFG_NACL_CROSS_PATH  :=  
    configure: CFG_RELEASE_CHANNEL  := dev 
    configure: CFG_MUSL_ROOT        := /usr/local 
    configure: CFG_EXTRA_FILENAME   :=  
    configure: CFG_DEFAULT_LINKER   := cc 
    configure: CFG_DEFAULT_AR       := ar 
    configure: CFG_LIBDIR           := /usr/local/lib 
    configure: 
    configure: validating ./configure args
    configure: 
    configure: 
    configure: looking for build programs
    configure: 
    configure: CFG_CURL             := /bin/curl (7.43.0)
    configure: CFG_PYTHON           := /bin/python2.7 
    configure: CFG_GIT              := /bin/git (2.5.0)
    configure: CFG_MD5              :=  
    configure: CFG_MD5SUM           := /bin/md5sum (8.24)
    configure: CFG_HASH_COMMAND     := /bin/md5sum | cut -c 1-8 
    configure: CFG_CLANG            :=  
    configure: CFG_CCACHE           :=  
    configure: CFG_GCC              := /bin/gcc (5.3.1)
    configure: CFG_LD               := /bin/ld (2.25-15.fc23)
    configure: CFG_VALGRIND         :=  
    configure: CFG_PERF             :=  
    configure: CFG_ISCC             :=  
    configure: CFG_ANTLR4           :=  
    configure: CFG_GRUN             :=  
    configure: CFG_FLEX             :=  
    configure: CFG_BISON            :=  
    configure: CFG_GDB              :=  
    configure: CFG_LLDB             :=  
    configure: CFG_DISABLE_VALGRIND_RPASS := 1 
    configure: CFG_LLDB_PYTHON      := /bin/python2.7 
    configure: 
    configure: looking for target specific programs
    configure: 
    configure: CFG_ADB              :=  
    configure: CFG_CC               := gcc 
    configure: CFG_STDCPP_NAME      := stdc++ 
    configure: 
    configure: making directories
    configure: 
    configure: 
    configure: configuring submodules
    configure: 
    configure: git: submodule sync
    为 'src/compiler-rt' 同步子模组 url
    为 'src/jemalloc' 同步子模组 url
    为 'src/liblibc' 同步子模组 url
    为 'src/llvm' 同步子模组 url
    为 'src/rt/hoedown' 同步子模组 url
    为 'src/rust-installer' 同步子模组 url
    configure: git: submodule init
    configure: git: submodule update
    正克隆到 'src/llvm'...
    remote: Counting objects: 1109268, done.
    remote: Total 1109268 (delta 0), reused 0 (delta 0), pack-reused 1109268
    接收对象中: 100% (1109268/1109268), 439.39 MiB | 259.00 KiB/s, 完成.
    处理 delta 中: 100% (909146/909146), 完成.
    检查连接... 完成。
    子模组路径 'src/llvm'：检出 '80ad955b60b3ac02d0462a4a65fcea597d0ebfb1'
    正克隆到 'src/rt/hoedown'...
    remote: Counting objects: 2475, done.
    remote: Total 2475 (delta 0), reused 0 (delta 0), pack-reused 2475
    接收对象中: 100% (2475/2475), 635.34 KiB | 129.00 KiB/s, 完成.
    处理 delta 中: 100% (1468/1468), 完成.
    检查连接... 完成。
    子模组路径 'src/rt/hoedown'：检出 'a3736a0a1907cbc8bf619708738815a5fd789c80'
    正克隆到 'src/rust-installer'...
    remote: Counting objects: 354, done.
    接收对象中: 100% (354/354), 112.35 KiB | 74.00 KiB/s, 完成.
    remote: Total 354 (delta 0), reused 0 (delta 0), pack-reused 354
    处理 delta 中: 100% (215/215), 完成.
    检查连接... 完成。
    子模组路径 'src/rust-installer'：检出 'c37d3747da75c280237dc2d6b925078e69555499'
    configure: git: submodule foreach sync
    正在进入 'src/compiler-rt'
    正在进入 'src/jemalloc'
    正在进入 'src/liblibc'
    正在进入 'src/llvm'
    正在进入 'src/rt/hoedown'
    正在进入 'src/rust-installer'
    configure: git: submodule foreach update
    configure: git: submodule status
     57315f7e07d09b6f0341ebbcd50dded6c20d782f src/compiler-rt (remotes/origin/rust-2015-11-06)
     aab1c0a0e0b39825b16673128729ef46310a5da8 src/jemalloc (3.6.0-531-gaab1c0a)
     45d85899e99d33e291b2bf3259881b46cc5365d7 src/liblibc (0.2.11-57-g45d8589)
     80ad955b60b3ac02d0462a4a65fcea597d0ebfb1 src/llvm (remotes/origin/rust-llvm-2016-03-13)
     a3736a0a1907cbc8bf619708738815a5fd789c80 src/rt/hoedown (2.0.0-211-ga3736a0)
     c37d3747da75c280237dc2d6b925078e69555499 src/rust-installer (remotes/origin/next-27-gc37d374)
    -aed73472416064642911af790b25d57c9390b6c7 src/rust-installer/test/rust-installer-v1
    -e577c97b494be2815b215e3042207d6d4b7c5516 src/rust-installer/test/rust-installer-v2
    configure: git: submodule clobber
    正在进入 'src/compiler-rt'
    正在进入 'src/jemalloc'
    正在进入 'src/liblibc'
    正在进入 'src/llvm'
    正在进入 'src/rt/hoedown'
    正在进入 'src/rust-installer'
    正在进入 'src/compiler-rt'
    正在进入 'src/jemalloc'
    正在进入 'src/liblibc'
    正在进入 'src/llvm'
    正在进入 'src/rt/hoedown'
    正在进入 'src/rust-installer'
    configure: 
    configure: looking at LLVM
    configure: 
    configure: configuring LLVM for x86_64-unknown-linux-gnu
    configure: configuring LLVM with:
    configure: --enable-targets=x86,x86_64,arm,aarch64,mips,powerpc --enable-optimized --disable-assertions --disable-docs --enable-bindings=none --disable-terminfo --disable-zlib --disable-libffi --build=x86_64-unknown-linux-gnu                         --host=x86_64-unknown-linux-gnu --target=x86_64-unknown-linux-gnu --with-python=/bin/python2.7
    checking for x86_64-unknown-linux-gnu-clang... gcc
    checking for C compiler default output file name... a.out
    checking whether the C compiler works... yes
    checking whether we are cross compiling... no
    checking for suffix of executables... 
    checking for suffix of object files... o
    checking whether we are using the GNU C compiler... yes
    checking whether gcc accepts -g... yes
    checking for gcc option to accept ISO C89... none needed
    checking whether we are using the GNU C++ compiler... yes
    checking whether g++ accepts -g... yes
    checking how to run the C preprocessor... gcc -E
    checking whether GCC or Clang is our host compiler... gcc
    checking build system type... x86_64-unknown-linux-gnu
    checking host system type... x86_64-unknown-linux-gnu
    checking target system type... x86_64-unknown-linux-gnu
    checking type of operating system we're going to host on... Linux
    checking type of operating system we're going to target... Linux
    checking target architecture... x86_64
    checking whether GCC is new enough... yes
    checking optimization flags... -O3
    checking for GNU make... make
    checking whether ln -s works... yes
    checking for nm... /bin/nm
    checking for cmp... /bin/cmp
    checking for cp... /bin/cp
    checking for date... /bin/date
    checking for find... /bin/find
    checking for grep... /bin/grep
    checking for mkdir... /bin/mkdir
    checking for mv... /bin/mv
    checking for x86_64-unknown-linux-gnu-ranlib... no
    checking for ranlib... ranlib
    checking for x86_64-unknown-linux-gnu-ar... no
    checking for ar... ar
    checking for rm... /bin/rm
    checking for sed... /bin/sed
    checking for tar... /bin/tar
    checking for pwd... /bin/pwd
    checking for dot... echo dot
    checking for a BSD-compatible install... /bin/install -c
    checking for bzip2... /bin/bzip2
    checking for cat... /bin/cat
    checking for doxygen... /bin/doxygen
    checking for groff... /bin/groff
    checking for gzip... /bin/gzip
    checking for pdfroff... no
    checking for zip... /bin/zip
    checking for go... /opt/go/bin/go
    checking for ocamlfind... no
    checking for gas... no
    checking for as... /bin/as
    checking for linker version... 2.25
    checking for compiler -Wl,-R<path> option... yes
    checking for compiler -rdynamic option... yes
    checking for compiler -Wl,--version-script option... yes
    checking for grep that handles long lines and -e... (cached) /bin/grep
    checking for egrep... /bin/grep -E
    checking for ANSI C header files... yes
    checking for sys/types.h... yes
    checking for sys/stat.h... yes
    checking for stdlib.h... yes
    checking for string.h... yes
    checking for memory.h... yes
    checking for strings.h... yes
    checking for inttypes.h... yes
    checking for stdint.h... yes
    checking for unistd.h... yes
    checking errno.h usability... yes
    checking errno.h presence... yes
    checking for errno.h... yes
    checking tool compatibility... ok
    checking optional compiler flags... -Wno-variadic-macros -Wno-missing-field-initializers   -Wno-maybe-uninitialized -Wno-comment
    checking for python... user defined: /bin/python2.7
    checking for python >= 2.7... /bin/python2.7 (2.7.11)
    checking for sin in -lm... yes
    checking for library containing dlopen... -ldl
    checking for library containing clock_gettime... none required
    checking for library containing el_init... no
    checking for library containing mallinfo... none required
    checking for pthread_mutex_init in -lpthread... yes
    checking for library containing pthread_mutex_lock... none required
    checking for library containing pthread_rwlock_init... none required
    checking for library containing pthread_getspecific... none required
    checking for xml2-config... no
    checking for libxml2 includes... xml2-config not found
    checking for dirent.h that defines DIR... yes
    checking for library containing opendir... none required
    checking for MAP_ANONYMOUS vs. MAP_ANON... yes
    checking whether stat file-mode macros are broken... no
    checking whether time.h and sys/time.h may both be included... yes
    checking for cxxabi.h... yes
    checking dlfcn.h usability... yes
    checking dlfcn.h presence... yes
    checking for dlfcn.h... yes
    checking execinfo.h usability... yes
    checking execinfo.h presence... yes
    checking for execinfo.h... yes
    checking fcntl.h usability... yes
    checking fcntl.h presence... yes
    checking for fcntl.h... yes
    checking for inttypes.h... (cached) yes
    checking link.h usability... yes
    checking link.h presence... yes
    checking for link.h... yes
    checking malloc.h usability... yes
    checking malloc.h presence... yes
    checking for malloc.h... yes
    checking setjmp.h usability... yes
    checking setjmp.h presence... yes
    checking for setjmp.h... yes
    checking signal.h usability... yes
    checking signal.h presence... yes
    checking for signal.h... yes
    checking for stdint.h... (cached) yes
    checking termios.h usability... yes
    checking termios.h presence... yes
    checking for termios.h... yes
    checking for unistd.h... (cached) yes
    checking sys/mman.h usability... yes
    checking sys/mman.h presence... yes
    checking for sys/mman.h... yes
    checking sys/param.h usability... yes
    checking sys/param.h presence... yes
    checking for sys/param.h... yes
    checking sys/resource.h usability... yes
    checking sys/resource.h presence... yes
    checking for sys/resource.h... yes
    checking sys/time.h usability... yes
    checking sys/time.h presence... yes
    checking for sys/time.h... yes
    checking sys/uio.h usability... yes
    checking sys/uio.h presence... yes
    checking for sys/uio.h... yes
    checking sys/ioctl.h usability... yes
    checking sys/ioctl.h presence... yes
    checking for sys/ioctl.h... yes
    checking malloc/malloc.h usability... no
    checking malloc/malloc.h presence... no
    checking for malloc/malloc.h... no
    checking mach/mach.h usability... no
    checking mach/mach.h presence... no
    checking for mach/mach.h... no
    checking valgrind/valgrind.h usability... no
    checking valgrind/valgrind.h presence... no
    checking for valgrind/valgrind.h... no
    checking fenv.h usability... yes
    checking fenv.h presence... yes
    checking for fenv.h... yes
    checking whether FE_ALL_EXCEPT is declared... yes
    checking whether FE_INEXACT is declared... yes
    checking pthread.h usability... yes
    checking pthread.h presence... yes
    checking for pthread.h... yes
    checking CrashReporterClient.h usability... no
    checking CrashReporterClient.h presence... no
    checking for CrashReporterClient.h... no
    checking __crashreporter_info__... no
    checking for HUGE_VAL sanity... yes
    checking for pid_t... yes
    checking for size_t... yes
    checking whether struct tm is in sys/time.h or time.h... time.h
    checking for int64_t... yes
    checking for uint64_t... yes
    checking for backtrace... yes
    checking for getcwd... yes
    checking for getpagesize... yes
    checking for getrusage... yes
    checking for getrlimit... yes
    checking for setrlimit... yes
    checking for gettimeofday... yes
    checking for isatty... yes
    checking for mkdtemp... yes
    checking for mkstemp... yes
    checking for mktemp... yes
    checking for posix_spawn... yes
    checking for pread... yes
    checking for realpath... yes
    checking for sbrk... yes
    checking for setrlimit... (cached) yes
    checking for strerror... yes
    checking for strerror_r... yes
    checking for setenv... yes
    checking for strtoll... yes
    checking for strtoq... yes
    checking for sysconf... yes
    checking for malloc_zone_statistics... no
    checking for setjmp... yes
    checking for longjmp... yes
    checking for writev... yes
    checking for futimes... yes
    checking for futimens... yes
    checking if printf has the %a format character... yes
    checking whether arc4random is declared... no
    checking whether strerror_s is declared... no
    checking for stdlib.h... (cached) yes
    checking for unistd.h... (cached) yes
    checking for getpagesize... (cached) yes
    checking for working mmap... yes
    checking for mmap of files... yes
    checking if /dev/zero is needed for mmap... no
    checking for GCC atomic builtins... yes
    checking for 32-bit userspace on 64-bit system... no
    checking for __dso_handle... yes
    checking for compiler -fvisibility-inlines-hidden option... yes
    configure: creating ./config.status
    config.status: creating include/llvm/Config/Targets.def
    config.status: creating include/llvm/Config/AsmPrinters.def
    config.status: creating include/llvm/Config/AsmParsers.def
    config.status: creating include/llvm/Config/Disassemblers.def
    config.status: creating Makefile.config
    config.status: creating llvm.spec
    config.status: creating docs/doxygen.cfg
    config.status: creating bindings/ocaml/llvm/META.llvm
    config.status: creating include/llvm/Config/config.h
    config.status: creating include/llvm/Config/llvm-config.h
    config.status: creating include/llvm/Support/DataTypes.h
    config.status: executing setup commands
    config.status: executing Makefile commands
    config.status: executing Makefile.common commands
    config.status: executing examples/Makefile commands
    config.status: executing lib/Makefile commands
    config.status: executing test/Makefile commands
    config.status: executing test/Makefile.tests commands
    config.status: executing unittests/Makefile commands
    config.status: executing tools/Makefile commands
    config.status: executing utils/Makefile commands
    config.status: executing projects/Makefile commands
    config.status: executing bindings/Makefile commands
    config.status: executing bindings/ocaml/Makefile.ocaml commands
    
    
    ################################################################################
    ################################################################################
    The LLVM project has deprecated building with configure & make.
    The autoconf-based makefile build system will be removed in the 3.9 release.
    
    Please migrate to the CMake-based build system.
    For more information see: http://llvm.org/docs/CMake.html
    ################################################################################
    ################################################################################
    configure: 
    configure: writing configuration
    configure: 
    configure: CFG_SRC_DIR          := /data/src/github.com/rust-lang/rust ...
    configure: CFG_SRC_DIR_RELATIVE := ./ 
    configure: CFG_BUILD_DIR        := /data/src/github.com/rust-lang/rust ...
    configure: CFG_OSTYPE           := unknown-linux-gnu 
    configure: CFG_CPUTYPE          := x86_64 
    configure: CFG_CONFIGURE_ARGS   :=  
    configure: CFG_PREFIX           := /usr/local 
    configure: CFG_HOST             := x86_64-unknown-linux-gnu 
    configure: CFG_TARGET           := x86_64-unknown-linux-gnu 
    configure: CFG_LIBDIR_RELATIVE  := lib 
    configure: CFG_DISABLE_MANAGE_SUBMODULES :=  
    configure: CFG_AARCH64_LINUX_ANDROID_NDK :=  
    configure: CFG_ARM_LINUX_ANDROIDEABI_NDK :=  
    configure: CFG_ARMV7_LINUX_ANDROIDEABI_NDK :=  
    configure: CFG_I686_LINUX_ANDROID_NDK :=  
    configure: CFG_NACL_CROSS_PATH  :=  
    configure: CFG_MANDIR           := /usr/local/share/man 
    configure: CFG_USING_LIBCPP     := 0 
    configure: CFG_LLVM_SRC_DIR     := /data/src/github.com/rust-lang/rust ...
    configure: CFG_LLVM_BUILD_DIR_x86_64_unknown_linux_gnu := /data/src/github.com/rust-lang/rust ...
    configure: CFG_LLVM_INST_DIR_x86_64_unknown_linux_gnu := /data/src/github.com/rust-lang/rust ...
    configure: 
    configure: cp -f /data/src/github.com/rust-lang/rust/Makefile.in ./Makefile
    configure: mv -f config.tmp config.mk
    configure: 
    configure: configured in release mode. for development consider --enable-debug
    configure: 
    configure: run `make help`
    configure: 

Make bin

    [vagrant@localhost rust]$ make
    cfg: version 1.11.0-dev (25f349db3 2016-06-19)
    cfg: build triple x86_64-unknown-linux-gnu
    cfg: host triples x86_64-unknown-linux-gnu
    cfg: target triples x86_64-unknown-linux-gnu
    cfg: host for x86_64-unknown-linux-gnu is x86_64
    cfg: os for x86_64-unknown-linux-gnu is unknown-linux-gnu
    cfg: have good valgrind for x86_64-unknown-linux-gnu
    cfg: using CC=gcc (CFG_CC)
    cfg: disabling valgrind run-pass tests
    make: cleaning llvm
    make[1]: Entering directory '/data/src/github.com/rust-lang/rust'
    cfg: version 1.11.0-dev (25f349db3 2016-06-19)
    cfg: build triple x86_64-unknown-linux-gnu
    cfg: host triples x86_64-unknown-linux-gnu
    cfg: target triples x86_64-unknown-linux-gnu
    cfg: host for x86_64-unknown-linux-gnu is x86_64
    cfg: os for x86_64-unknown-linux-gnu is unknown-linux-gnu
    cfg: have good valgrind for x86_64-unknown-linux-gnu
    cfg: using CC=gcc (CFG_CC)
    cfg: disabling valgrind run-pass tests
    cfg: including dist rules
    cfg: including clean rules
    make[2]: Entering directory '/data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/llvm'
    llvm[2]: Constructing LLVMBuild project information.
    make[3]: Entering directory '/data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/llvm/lib/Support'
    llvm[3]: Constructing LLVMBuild project information.
    make[3]: Leaving directory '/data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/llvm/lib/Support'

...

    make: done cleaning llvm
    touch -r /data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/llvm/llvm-auto-clean-stamp.start_time /data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/llvm/llvm-auto-clean-stamp && rm /data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/llvm/llvm-auto-clean-stamp.start_time
    make: llvm
    make[1]: Entering directory '/data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/llvm'
    make[2]: Entering directory '/data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/llvm/lib/Support'
    llvm[2]: Compiling APFloat.cpp for Release build
    llvm[2]: Compiling APInt.cpp for Release build
    llvm[2]: Compiling APSInt.cpp for Release build

...

    llvm[4]: Compiling AArch64LoadStoreOptimizer.cpp for Release build
    /data/src/github.com/rust-lang/rust/src/llvm/lib/Target/AArch64/AArch64LoadStoreOptimizer.cpp: 在成员函数‘llvm::MachineBasicBlock::iterator {匿名}::AArch64LoadStoreOpt::mergePairedInsns(llvm::MachineBasicBlock::iterator, llvm::MachineBasicBlock::iterator, const LdStPairFlags&)’中:
    /data/src/github.com/rust-lang/rust/src/llvm/lib/Target/AArch64/AArch64LoadStoreOptimizer.cpp:675:19: 警告：variable ‘NewMemMI’ set but not used [-Wunused-but-set-variable]
         MachineInstr *NewMemMI, *BitExtMI1, *BitExtMI2;
                       ^
    /data/src/github.com/rust-lang/rust/src/llvm/lib/Target/AArch64/AArch64LoadStoreOptimizer.cpp:675:30: 警告：variable ‘BitExtMI1’ set but not used [-Wunused-but-set-variable]
         MachineInstr *NewMemMI, *BitExtMI1, *BitExtMI2;
                                  ^
    /data/src/github.com/rust-lang/rust/src/llvm/lib/Target/AArch64/AArch64LoadStoreOptimizer.cpp:675:42: 警告：variable ‘BitExtMI2’ set but not used [-Wunused-but-set-variable]
         MachineInstr *NewMemMI, *BitExtMI1, *BitExtMI2;
                                              ^
    /data/src/github.com/rust-lang/rust/src/llvm/lib/Target/AArch64/AArch64LoadStoreOptimizer.cpp: 在成员函数‘llvm::MachineBasicBlock::iterator {匿名}::AArch64LoadStoreOpt::promoteLoadFromStore(llvm::MachineBasicBlock::iterator, llvm::MachineBasicBlock::iterator)’中:
    /data/src/github.com/rust-lang/rust/src/llvm/lib/Target/AArch64/AArch64LoadStoreOptimizer.cpp:861:17: 警告：variable ‘BitExtMI’ set but not used [-Wunused-but-set-variable]
       MachineInstr *BitExtMI;
                     ^
    llvm[4]: Compiling AArch64MCInstLower.cpp for Release build
    llvm[4]: Compiling AArch64PBQPRegAlloc.cpp for Release build
    llvm[4]: Compiling AArch64PromoteConstant.cpp for Release build

...

    llvm[1]: ***** Completed Release Build
    make[1]: Leaving directory '/data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/llvm'
    fetch: x86_64-unknown-linux-gnu/stage0/bin/rustc
    downloading https://static.rust-lang.org/dist/2016-05-24/rustc-beta-x86_64-unknown-linux-gnu.tar.gz.sha256 to /tmp/tmpKJN0XW.sha256
      % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                     Dload  Upload   Total   Spent    Left  Speed
    100   109  100   109    0     0     17      0  0:00:06  0:00:06 --:--:--    28
    downloading https://static.rust-lang.org/dist/2016-05-24/rustc-beta-x86_64-unknown-linux-gnu.tar.gz to /tmp/tmpFwKtKD
      % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                     Dload  Upload   Total   Spent    Left  Speed
    100 44.3M  100 44.3M    0     0   901k      0  0:00:50  0:00:50 --:--:-- 1089k
    verifying /tmp/tmpFwKtKD
    moving /tmp/tmpFwKtKD to dl/rustc-beta-x86_64-unknown-linux-gnu.tar.gz
    removing /tmp/tmpKJN0XW.sha256
    extracting dl/rustc-beta-x86_64-unknown-linux-gnu.tar.gz
      extracting rustc-beta-x86_64-unknown-linux-gnu/rustc
      extracting rustc-beta-x86_64-unknown-linux-gnu/rustc/manifest.in
      extracting rustc-beta-x86_64-unknown-linux-gnu/rustc/bin
      extracting rustc-beta-x86_64-unknown-linux-gnu/rustc/bin/rustdoc
      extracting rustc-beta-x86_64-unknown-linux-gnu/rustc/bin/rust-gdb
      extracting rustc-beta-x86_64-unknown-linux-gnu/rustc/bin/rustc
      extracting rustc-beta-x86_64-unknown-linux-gnu/rustc/lib
      extracting rustc-beta-x86_64-unknown-linux-gnu/rustc/lib/libterm-a4922355.so
      extracting rustc-beta-x86_64-unknown-linux-gnu/rustc/lib/libarena-a4922355.so
      extracting rustc-beta-x86_64-unknown-linux-gnu/rustc/lib/librustc_driver-a4922355.so

...

    make: compiler-rt
    make[1]: Entering directory '/data/src/github.com/rust-lang/rust/src/compiler-rt'
      MKDIR:     /data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/rt/compiler-rt/triple/builtins/x86_64/SubDir.lib__builtins
      COMPILE:   triple/builtins/x86_64: /data/src/github.com/rust-lang/rust/src/compiler-rt/lib/builtins/absvdi2.c
      COMPILE:   triple/builtins/x86_64: /data/src/github.com/rust-lang/rust/src/compiler-rt/lib/builtins/absvsi2.c
      COMPILE:   triple/builtins/x86_64: /data/src/github.com/rust-lang/rust/src/compiler-rt/lib/builtins/absvti2.c

...

    cp: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcompiler-rt.a
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcore
    src/libcore/num/mod.rs:1036:9: 1036:41 warning: unused attribute, #[warn(unused_attributes)] on by default
    src/libcore/num/mod.rs:1036         #[rustc_inherit_overflow_checks]
                                        ^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

...

    ===============================================================================
    jemalloc version   : 0.0.0-0-g0000000000000000000000000000000000000000
    library revision   : 2
    
    CONFIG             : --build=x86_64-unknown-linux-gnu --host=x86_64-unknown-linux-gnu 'CC=gcc -m64' AR=ar 'RANLIB=ar s' 'CPPFLAGS=-I /data/src/github.com/rust-lang/rust/src/rt/' 'EXTRA_CFLAGS=-g1 -ffunction-sections -fdata-sections' build_alias=x86_64-unknown-linux-gnu host_alias=x86_64-unknown-linux-gnu
    CC                 : gcc -m64
    CFLAGS             : -std=gnu99 -Wall -Werror=declaration-after-statement -pipe -g3 -g1 -ffunction-sections -fdata-sections -fvisibility=hidden -O3 -funroll-loops
    CPPFLAGS           : -I /data/src/github.com/rust-lang/rust/src/rt/ -D_GNU_SOURCE -D_REENTRANT
    LDFLAGS            : 
    EXTRA_LDFLAGS      : 
    LIBS               :  -lpthread
    TESTLIBS           : 
    RPATH_EXTRA        : 
    
    XSLTPROC           : /bin/xsltproc
    XSLROOT            : 
    
    PREFIX             : /usr/local
    BINDIR             : /usr/local/bin
    DATADIR            : /usr/local/share
    INCLUDEDIR         : /usr/local/include
    LIBDIR             : /usr/local/lib
    MANDIR             : /usr/local/share/man
    
    srcroot            : /data/src/github.com/rust-lang/rust/src/jemalloc/
    abs_srcroot        : /data/src/github.com/rust-lang/rust/src/jemalloc/
    objroot            : 
    abs_objroot        : /data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/rt/jemalloc/
    
    JEMALLOC_PREFIX    : 
    JEMALLOC_PRIVATE_NAMESPACE
                       : je_
    install_suffix     : 
    malloc_conf        : 
    autogen            : 0
    cc-silence         : 1
    debug              : 0
    code-coverage      : 0
    stats              : 1
    prof               : 0
    prof-libunwind     : 0
    prof-libgcc        : 0
    prof-gcc           : 0
    tcache             : 1
    fill               : 1
    utrace             : 0
    valgrind           : 0
    xmalloc            : 0
    munmap             : 0
    lazy_lock          : 0
    tls                : 1
    cache-oblivious    : 1
    ===============================================================================

...

    libtool: link: ar cru .libs/libbacktrace.a .libs/atomic.o .libs/dwarf.o .libs/fileline.o .libs/posix.o .libs/print.o .libs/sort.o .libs/state.o .libs/backtrace.o .libs/simple.o .libs/elf.o .libs/mmapio.o .libs/mmap.o 
    ar: `u' modifier ignored since `D' is the default (see `U')
    libtool: link: ar s .libs/libbacktrace.a
    libtool: link: ( cd ".libs" && rm -f "libbacktrace.la" && ln -s "../libbacktrace.la" "libbacktrace.la" )
    true  DO=all multi-do # make
    make[2]: Leaving directory '/data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/rt/libbacktrace'
    make[1]: Leaving directory '/data/src/github.com/rust-lang/rust/x86_64-unknown-linux-gnu/rt/libbacktrace'
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/libstd
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/libarena
    compile: x86_64-unknown-linux-gnu/rt/miniz.o
    link: x86_64-unknown-linux-gnu/rt/libminiz.a
    ar: `u' modifier ignored since `D' is the default (see `U')
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/libflate
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/libgetopts
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/libgraphviz
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/libterm
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/liblog
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/libserialize
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_bitflags
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/libsyntax
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/libfmt_macros
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librbml
    compile: x86_64-unknown-linux-gnu/rustllvm/ExecutionEngineWrapper.o
    compile: x86_64-unknown-linux-gnu/rustllvm/RustWrapper.o
    compile: x86_64-unknown-linux-gnu/rustllvm/PassWrapper.o
    compile: x86_64-unknown-linux-gnu/rustllvm/ArchiveWrapper.o
    link: x86_64-unknown-linux-gnu/rt/librustllvm.a
    ar: `u' modifier ignored since `D' is the default (see `U')
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_llvm
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_back
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_data_structures
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_const_math
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_const_eval
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_mir
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_borrowck
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_platform_intrinsics
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_typeck
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_resolve
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_incremental
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_trans
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_privacy
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_lint
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_metadata
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_plugin
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/libsyntax_ext
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_passes
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_save_analysis
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_driver
    mkdir -p x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/bin/
    rustc: x86_64-unknown-linux-gnu/stage0/lib/rustlib/x86_64-unknown-linux-gnu/bin/rustc
    cp: x86_64-unknown-linux-gnu/stage1/lib/libstd
    cp: x86_64-unknown-linux-gnu/stage1/lib/libarena
    cp: x86_64-unknown-linux-gnu/stage1/lib/libflate
    cp: x86_64-unknown-linux-gnu/stage1/lib/libgetopts
    cp: x86_64-unknown-linux-gnu/stage1/lib/libgraphviz
    cp: x86_64-unknown-linux-gnu/stage1/lib/libterm
    cp: x86_64-unknown-linux-gnu/stage1/lib/liblog
    cp: x86_64-unknown-linux-gnu/stage1/lib/libserialize
    cp: x86_64-unknown-linux-gnu/stage1/lib/libsyntax
    cp: x86_64-unknown-linux-gnu/stage1/lib/libfmt_macros
    cp: x86_64-unknown-linux-gnu/stage1/lib/librbml
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_llvm
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_back
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_data_structures
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_const_math
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_const_eval
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_mir
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_borrowck
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_platform_intrinsics
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_typeck
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_resolve
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_incremental
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_trans
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_privacy
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_lint
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_metadata
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_plugin
    cp: x86_64-unknown-linux-gnu/stage1/lib/libsyntax_ext
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_passes
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_save_analysis
    cp: x86_64-unknown-linux-gnu/stage1/lib/librustc_driver
    cp: x86_64-unknown-linux-gnu/stage1/bin/rustc
    mkdir -p x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/
    cp: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcompiler-rt.a
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcore
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/liblibc
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librand
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc_system
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_unicode
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcollections
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libpanic_abort
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libunwind
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libpanic_unwind
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc_jemalloc
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libstd
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libarena
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libflate
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libgetopts
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libgraphviz
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libterm
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/liblog
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libserialize
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_bitflags
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libsyntax
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libfmt_macros
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librbml
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_llvm
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_back
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_data_structures
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_const_math
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_const_eval
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_mir
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_borrowck
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_platform_intrinsics
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_typeck
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_resolve
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_incremental
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_trans
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_privacy
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_lint
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_metadata
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_plugin
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libsyntax_ext
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_passes
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_save_analysis
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_driver
    mkdir -p x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/bin/
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/bin/rustc
    cp: x86_64-unknown-linux-gnu/stage2/lib/libstd
    cp: x86_64-unknown-linux-gnu/stage2/lib/libarena
    cp: x86_64-unknown-linux-gnu/stage2/lib/libflate
    cp: x86_64-unknown-linux-gnu/stage2/lib/libgetopts
    cp: x86_64-unknown-linux-gnu/stage2/lib/libgraphviz
    cp: x86_64-unknown-linux-gnu/stage2/lib/libterm
    cp: x86_64-unknown-linux-gnu/stage2/lib/liblog
    cp: x86_64-unknown-linux-gnu/stage2/lib/libserialize
    cp: x86_64-unknown-linux-gnu/stage2/lib/libsyntax
    cp: x86_64-unknown-linux-gnu/stage2/lib/libfmt_macros
    cp: x86_64-unknown-linux-gnu/stage2/lib/librbml
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_llvm
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_back
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_data_structures
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_const_math
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_const_eval
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_mir
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_borrowck
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_platform_intrinsics
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_typeck
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_resolve
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_incremental
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_trans
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_privacy
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_lint
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_metadata
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_plugin
    cp: x86_64-unknown-linux-gnu/stage2/lib/libsyntax_ext
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_passes
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_save_analysis
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustc_driver
    cp: x86_64-unknown-linux-gnu/stage2/bin/rustc
    mkdir -p x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/
    cp: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcompiler-rt.a
    compile: x86_64-unknown-linux-gnu/rt/rust_test_helpers.o
    link: x86_64-unknown-linux-gnu/rt/librust_test_helpers.a
    ar: `u' modifier ignored since `D' is the default (see `U')
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/libtest
    compile: x86_64-unknown-linux-gnu/rt/hoedown/src/autolink.o
    compile: x86_64-unknown-linux-gnu/rt/hoedown/src/buffer.o
    compile: x86_64-unknown-linux-gnu/rt/hoedown/src/document.o
    compile: x86_64-unknown-linux-gnu/rt/hoedown/src/escape.o
    compile: x86_64-unknown-linux-gnu/rt/hoedown/src/html.o
    compile: x86_64-unknown-linux-gnu/rt/hoedown/src/html_blocks.o
    compile: x86_64-unknown-linux-gnu/rt/hoedown/src/html_smartypants.o
    compile: x86_64-unknown-linux-gnu/rt/hoedown/src/stack.o
    compile: x86_64-unknown-linux-gnu/rt/hoedown/src/version.o
    link: x86_64-unknown-linux-gnu/rt/libhoedown.a
    ar: `u' modifier ignored since `D' is the default (see `U')
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustdoc
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/bin/rustdoc
    cp: x86_64-unknown-linux-gnu/stage2/lib/libtest
    cp: x86_64-unknown-linux-gnu/stage2/lib/librustdoc
    cp: x86_64-unknown-linux-gnu/stage2/bin/rustdoc
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcore
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/liblibc
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librand
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc_system
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_unicode
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcollections
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libpanic_abort
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libunwind
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libpanic_unwind
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc_jemalloc
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libstd
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libterm
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/liblog
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libserialize
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libarena
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_bitflags
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libsyntax
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libfmt_macros
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libsyntax_ext
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libflate
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libgetopts
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librbml
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libgraphviz
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_llvm
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_back
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_data_structures
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_const_math
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_platform_intrinsics
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_const_eval
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_typeck
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_mir
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_borrowck
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_resolve
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_incremental
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_trans
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_privacy
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_lint
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_metadata
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_plugin
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_passes
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_save_analysis
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_driver
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/libtest
    rustc: x86_64-unknown-linux-gnu/stage2/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustdoc
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/bin/rustbook
    cp: x86_64-unknown-linux-gnu/stage2/bin/rustbook
    rustbook: doc/book/index.html
    rustbook: doc/nomicon/index.html
    rustbook: doc/style/index.html
    rustc: x86_64-unknown-linux-gnu/stage1/lib/rustlib/x86_64-unknown-linux-gnu/bin/error_index_generator
    cp: x86_64-unknown-linux-gnu/stage2/bin/error_index_generator
    error_index_generator: doc/error-index.html
    version-info: doc/version_info.html
    cp: doc/rust.css
    cp: doc/favicon.inc
    cp: doc/footer.inc
    rustdoc: doc/not_found.html
    rustdoc: doc/index.html
    rustdoc: doc/complement-lang-faq.html
    rustdoc: doc/complement-design-faq.html
    rustdoc: doc/complement-project-faq.html
    rustdoc: doc/rustdoc.html
    cp: doc/full-toc.inc
    rustdoc: doc/reference.html
    rustdoc: doc/grammar.html
    rustdoc: doc/guide-crates.html
    rustdoc: doc/guide-error-handling.html
    rustdoc: doc/guide-ffi.html
    rustdoc: doc/guide-macros.html
    rustdoc: doc/guide.html
    rustdoc: doc/guide-ownership.html
    rustdoc: doc/guide-plugins.html
    rustdoc: doc/guide-pointers.html
    rustdoc: doc/guide-strings.html
    rustdoc: doc/guide-tasks.html
    rustdoc: doc/guide-testing.html
    rustdoc: doc/tutorial.html
    rustdoc: doc/intro.html
    rustdoc: doc/std/index.html
    rustdoc: doc/alloc/index.html
    rustdoc: doc/collections/index.html
    rustdoc: doc/core/index.html
    rustdoc: doc/libc/index.html
    rustdoc: doc/rustc_unicode/index.html

Install rust-lang compiler

    [vagrant@localhost rust]$ sudo make install
    cfg: version 1.11.0-dev (25f349db3 2016-06-19)
    cfg: build triple x86_64-unknown-linux-gnu
    cfg: host triples x86_64-unknown-linux-gnu
    cfg: target triples x86_64-unknown-linux-gnu
    cfg: host for x86_64-unknown-linux-gnu is x86_64
    cfg: os for x86_64-unknown-linux-gnu is unknown-linux-gnu
    cfg: have good valgrind for x86_64-unknown-linux-gnu
    cfg: using CC=gcc (CFG_CC)
    cfg: disabling valgrind run-pass tests
    cfg: including prepare rules
    cfg: including dist rules
    cfg: including install rules
    cfg: version 1.11.0-dev (25f349db3 2016-06-19)
    cfg: build triple x86_64-unknown-linux-gnu
    cfg: host triples x86_64-unknown-linux-gnu
    cfg: target triples x86_64-unknown-linux-gnu
    cfg: host for x86_64-unknown-linux-gnu is x86_64
    cfg: os for x86_64-unknown-linux-gnu is unknown-linux-gnu
    cfg: have good valgrind for x86_64-unknown-linux-gnu
    cfg: using CC=gcc (CFG_CC)
    cfg: disabling valgrind run-pass tests
    cfg: including prepare rules
    cfg: including dist rules
    cfg: including install rules
    cleaning destination tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/bin
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/etc
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/share/man/man1
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/libstd-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/libterm-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/liblog-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/libserialize-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/libarena-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/libsyntax-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/libfmt_macros-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/libflate-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/libgetopts-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librbml-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/libgraphviz-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_llvm-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_back-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_data_structures-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_const_math-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_const_eval-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_mir-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_borrowck-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_platform_intrinsics-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_typeck-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_resolve-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_incremental-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_trans-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_privacy-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_lint-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_metadata-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_plugin-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/libsyntax_ext-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_passes-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_save_analysis-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustc_driver-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/libtest-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/librustdoc-*.so
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/bin/rustdoc
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/share/man/man1/rustdoc.1
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/bin/rustc
    prepare: tmp/dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu-image/share/man/man1/rustc.1
    rustdoc: doc/std/index.html
    rustdoc: doc/alloc/index.html
    rustdoc: doc/collections/index.html
    rustdoc: doc/core/index.html
    rustdoc: doc/libc/index.html
    rustdoc: doc/rustc_unicode/index.html
    build: dist/rustc-1.11.0-dev-x86_64-unknown-linux-gnu.tar.gz
    gen-installer: looking for programs
    gen-installer: 
    gen-installer: found tar
    gen-installer: found cp
    gen-installer: found rm
    gen-installer: found mkdir
    gen-installer: found echo
    gen-installer: found tr
    gen-installer: found awk
    gen-installer: 
    gen-installer: processing arguments
    gen-installer: 
    gen-installer: CFG_PRODUCT_NAME     := Rust 
    gen-installer: CFG_COMPONENT_NAME   := rustc 
    gen-installer: CFG_PACKAGE_NAME     := rustc-1.11.0-dev-x86_64-unknown-lin ...
    gen-installer: CFG_REL_MANIFEST_DIR := rustlib 
    gen-installer: CFG_SUCCESS_MESSAGE  := Rust-is-ready-to-roll. 
    gen-installer: CFG_LEGACY_MANIFEST_DIRS := rustlib,cargo 
    gen-installer: CFG_NON_INSTALLED_OVERLAY := tmp/dist/rustc-1.11.0-dev-x86_64-un ...
    gen-installer: CFG_BULK_DIRS        :=  
    gen-installer: CFG_IMAGE_DIR        := tmp/dist/rustc-1.11.0-dev-x86_64-un ...
    gen-installer: CFG_WORK_DIR         := tmp/dist 
    gen-installer: CFG_OUTPUT_DIR       := dist 
    gen-installer: 
    gen-installer: validating arguments
    gen-installer: 
    gen-install-script: looking for install programs
    gen-install-script: 
    gen-install-script: found sed
    gen-install-script: found chmod
    gen-install-script: found cat
    gen-install-script: 
    gen-install-script: processing arguments
    gen-install-script: 
    gen-install-script: CFG_PRODUCT_NAME     := Rust 
    gen-install-script: CFG_REL_MANIFEST_DIR := rustlib 
    gen-install-script: CFG_SUCCESS_MESSAGE  := Rust-is-ready-to-roll. 
    gen-install-script: CFG_OUTPUT_SCRIPT    := tmp/dist/rustc-1.11.0-dev-x86_64-un ...
    gen-install-script: CFG_LEGACY_MANIFEST_DIRS := rustlib,cargo 
    gen-install-script: 
    gen-install-script: validating arguments
    gen-install-script: 
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/bin
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/liblibc-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libstd-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libstd-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libterm-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libterm-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libgetopts-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libgetopts-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcollections-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libtest-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libtest-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librand-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcore-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_unicode-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_bitflags-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc_system-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libpanic_abort-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libpanic_unwind-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libunwind-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc_jemalloc-*.rlib
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libsyntax-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libsyntax_ext-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_typeck-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_mir-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_borrowck-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_resolve-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_driver-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_trans-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_back-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_llvm-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_privacy-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_lint-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_data_structures-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_platform_intrinsics-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_plugin-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_metadata-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_passes-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_save_analysis-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_const_eval-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_const_math-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_incremental-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustdoc-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libfmt_macros-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libflate-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libarena-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libgraphviz-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/librbml-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/liblog-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libserialize-*.so
    prepare: tmp/dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu-image/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcompiler-rt.a
    build: dist/rust-std-1.11.0-dev-x86_64-unknown-linux-gnu.tar.gz
    gen-installer: looking for programs
    gen-installer: 
    gen-installer: found tar
    gen-installer: found cp
    gen-installer: found rm
    gen-installer: found mkdir
    gen-installer: found echo
    gen-installer: found tr
    gen-installer: found awk
    gen-installer: 
    gen-installer: processing arguments
    gen-installer: 
    gen-installer: CFG_PRODUCT_NAME     := Rust 
    gen-installer: CFG_COMPONENT_NAME   := rust-std-x86_64-unknown-linux-gnu 
    gen-installer: CFG_PACKAGE_NAME     := rust-std-1.11.0-dev-x86_64-unknown- ...
    gen-installer: CFG_REL_MANIFEST_DIR := rustlib 
    gen-installer: CFG_SUCCESS_MESSAGE  := std-is-standing-at-the-ready. 
    gen-installer: CFG_LEGACY_MANIFEST_DIRS := rustlib,cargo 
    gen-installer: CFG_NON_INSTALLED_OVERLAY :=  
    gen-installer: CFG_BULK_DIRS        :=  
    gen-installer: CFG_IMAGE_DIR        := tmp/dist/rust-std-1.11.0-dev-x86_64 ...
    gen-installer: CFG_WORK_DIR         := tmp/dist 
    gen-installer: CFG_OUTPUT_DIR       := dist 
    gen-installer: 
    gen-installer: validating arguments
    gen-installer: 
    gen-install-script: looking for install programs
    gen-install-script: 
    gen-install-script: found sed
    gen-install-script: found chmod
    gen-install-script: found cat
    gen-install-script: 
    gen-install-script: processing arguments
    gen-install-script: 
    gen-install-script: CFG_PRODUCT_NAME     := Rust 
    gen-install-script: CFG_REL_MANIFEST_DIR := rustlib 
    gen-install-script: CFG_SUCCESS_MESSAGE  := std-is-standing-at-the-ready. 
    gen-install-script: CFG_OUTPUT_SCRIPT    := tmp/dist/rust-std-1.11.0-dev-x86_64 ...
    gen-install-script: CFG_LEGACY_MANIFEST_DIRS := rustlib,cargo 
    gen-install-script: 
    gen-install-script: validating arguments
    gen-install-script: 
    build: dist/rust-docs-1.11.0-dev-x86_64-unknown-linux-gnu.tar.gz
    gen-installer: looking for programs
    gen-installer: 
    gen-installer: found tar
    gen-installer: found cp
    gen-installer: found rm
    gen-installer: found mkdir
    gen-installer: found echo
    gen-installer: found tr
    gen-installer: found awk
    gen-installer: 
    gen-installer: processing arguments
    gen-installer: 
    gen-installer: CFG_PRODUCT_NAME     := Rust-Documentation 
    gen-installer: CFG_COMPONENT_NAME   := rust-docs 
    gen-installer: CFG_PACKAGE_NAME     := rust-docs-1.11.0-dev-x86_64-unknown ...
    gen-installer: CFG_REL_MANIFEST_DIR := rustlib 
    gen-installer: CFG_SUCCESS_MESSAGE  := Rust-documentation-is-installed. 
    gen-installer: CFG_LEGACY_MANIFEST_DIRS := rustlib,cargo 
    gen-installer: CFG_NON_INSTALLED_OVERLAY :=  
    gen-installer: CFG_BULK_DIRS        := share/doc/rust/html 
    gen-installer: CFG_IMAGE_DIR        := tmp/dist/rust-docs-1.11.0-dev-x86_6 ...
    gen-installer: CFG_WORK_DIR         := tmp/dist 
    gen-installer: CFG_OUTPUT_DIR       := dist 
    gen-installer: 
    gen-installer: validating arguments
    gen-installer: 
    gen-install-script: looking for install programs
    gen-install-script: 
    gen-install-script: found sed
    gen-install-script: found chmod
    gen-install-script: found cat
    gen-install-script: 
    gen-install-script: processing arguments
    gen-install-script: 
    gen-install-script: CFG_PRODUCT_NAME     := Rust-Documentation 
    gen-install-script: CFG_REL_MANIFEST_DIR := rustlib 
    gen-install-script: CFG_SUCCESS_MESSAGE  := Rust-documentation-is-installed. 
    gen-install-script: CFG_OUTPUT_SCRIPT    := tmp/dist/rust-docs-1.11.0-dev-x86_6 ...
    gen-install-script: CFG_LEGACY_MANIFEST_DIRS := rustlib,cargo 
    gen-install-script: 
    gen-install-script: validating arguments
    gen-install-script: 
    mkdir -p tmp/empty_dir
    install: creating uninstall script at /usr/local/lib/rustlib/uninstall.sh
    install: installing component 'rust-docs'
    
        Rust documentation is installed.
    
    install: creating uninstall script at /usr/local/lib/rustlib/uninstall.sh
    install: installing component 'rust-std-x86_64-unknown-linux-gnu'
    
        std is standing at the ready.
    
    install: creating uninstall script at /usr/local/lib/rustlib/uninstall.sh
    install: installing component 'rustc'
    
        Rust is ready to roll.
    
### Test

Test

    [vagrant@localhost rust]$ which rustc
    /usr/local/bin/rustc
    [vagrant@localhost rust]$ rustc --version
    rustc 1.11.0-dev (25f349db3 2016-06-19)

## cargo

### Clone

Clone by git

    [vagrant@localhost rust-lang]$ git clone --recursive https://github.com/rust-lang/cargo cargo
    正克隆到 'cargo'...
    remote: Counting objects: 22052, done.
    remote: Compressing objects: 100% (238/238), done.
    remote: Total 22052 (delta 217), reused 33 (delta 33), pack-reused 21778
    接收对象中: 100% (22052/22052), 10.63 MiB | 462.00 KiB/s, 完成.
    处理 delta 中: 100% (15989/15989), 完成.
    检查连接... 完成。
    子模组 'src/rust-installer' (https://github.com/rust-lang/rust-installer.git) 未对路径 'src/rust-installer' 注册
    正克隆到 'src/rust-installer'...
    remote: Counting objects: 354, done.
    remote: Total 354 (delta 0), reused 0 (delta 0), pack-reused 354
    接收对象中: 100% (354/354), 112.35 KiB | 190.00 KiB/s, 完成.
    处理 delta 中: 100% (215/215), 完成.
    检查连接... 完成。
    子模组路径 'src/rust-installer'：检出 'c37d3747da75c280237dc2d6b925078e69555499'
    子模组 'test/rust-installer-v1' (https://github.com/rust-lang/rust-installer) 未对路径 'test/rust-installer-v1' 注册
    子模组 'test/rust-installer-v2' (https://github.com/rust-lang/rust-installer) 未对路径 'test/rust-installer-v2' 注册
    正克隆到 'test/rust-installer-v1'...
    remote: Counting objects: 354, done.
    remote: Total 354 (delta 0), reused 0 (delta 0), pack-reused 354
    接收对象中: 100% (354/354), 112.35 KiB | 175.00 KiB/s, 完成.
    处理 delta 中: 100% (215/215), 完成.
    检查连接... 完成。
    子模组路径 'src/rust-installer/test/rust-installer-v1'：检出 'aed73472416064642911af790b25d57c9390b6c7'
    正克隆到 'test/rust-installer-v2'...
    remote: Counting objects: 354, done.
    remote: Total 354 (delta 0), reused 0 (delta 0), pack-reused 354
    接收对象中: 100% (354/354), 112.35 KiB | 128.00 KiB/s, 完成.
    处理 delta 中: 100% (215/215), 完成.
    检查连接... 完成。
    子模组路径 'src/rust-installer/test/rust-installer-v2'：检出 'e577c97b494be2815b215e3042207d6d4b7c5516'
    子模组 'test/rust-installer-v1' (https://github.com/rust-lang/rust-installer) 未对路径 'test/rust-installer-v1' 注册
    正克隆到 'test/rust-installer-v1'...
    remote: Counting objects: 354, done.
    remote: Total 354 (delta 0), reused 0 (delta 0), pack-reused 354
    接收对象中: 100% (354/354), 112.35 KiB | 95.00 KiB/s, 完成.
    处理 delta 中: 100% (215/215), 完成.
    检查连接... 完成。
    子模组路径 'src/rust-installer/test/rust-installer-v2/test/rust-installer-v1'：检出 'aed73472416064642911af790b25d57c9390b6c7'
    
    [vagrant@localhost cargo]$ ./configure 
    configure: looking for configure programs
    configure: found cmp
    configure: found mkdir
    configure: found printf
    configure: found cut
    configure: found head
    configure: found grep
    configure: found xargs
    configure: found cp
    configure: found find
    configure: found uname
    configure: found date
    configure: found tr
    configure: found sed
    configure: error: need cmake
    [vagrant@localhost cargo]$ sudo dnf install cmake
    上次元数据过期检查在 0:00:01 前执行于 Mon Jun 20 18:28:20 2016。
    依赖关系解决。
    ===============================================================================================================================================
     Package                        架构                          版本                                        仓库                            大小
    ===============================================================================================================================================
    安装:
     cmake                          x86_64                        3.4.3-1.fc23                                updates                        7.5 M
     jsoncpp                        x86_64                        0.6.0-0.18.rc2.fc23                         fedora                          63 k
    
    事务概要
    ===============================================================================================================================================
    安装  2 Packages
    
    总下载：7.6 M
    安装大小：23 M
    确定吗？[y/N]： y
    下载软件包：
    (1/2): jsoncpp-0.6.0-0.18.rc2.fc23.x86_64.rpm                                                                   87 kB/s |  63 kB     00:00    
    (2/2): cmake-3.4.3-1.fc23.x86_64.rpm                                                                           1.4 MB/s | 7.5 MB     00:05    
    -----------------------------------------------------------------------------------------------------------------------------------------------
    总计                                                                                                           691 kB/s | 7.6 MB     00:11     
    运行事务检查
    事务检查成功。
    运行事务测试
    事务测试成功。
    运行事务
      安装: jsoncpp-0.6.0-0.18.rc2.fc23.x86_64                                                                                                 1/2 
      安装: cmake-3.4.3-1.fc23.x86_64                                                                                                          2/2 
      验证: cmake-3.4.3-1.fc23.x86_64                                                                                                          1/2 
      验证: jsoncpp-0.6.0-0.18.rc2.fc23.x86_64                                                                                                 2/2 
    
    已安装:
      cmake.x86_64 3.4.3-1.fc23                                         jsoncpp.x86_64 0.6.0-0.18.rc2.fc23                                        
    
    完毕！

    [vagrant@localhost cargo]$ ./configure 
    configure: looking for configure programs
    configure: found cmp
    configure: found mkdir
    configure: found printf
    configure: found cut
    configure: found head
    configure: found grep
    configure: found xargs
    configure: found cp
    configure: found find
    configure: found uname
    configure: found date
    configure: found tr
    configure: found sed
    configure: found cmake
    configure: found curl
    configure: recreating config.tmp
    configure: 
    configure: processing ./configure args
    configure: 
    configure: CFG_PREFIX           := /usr/local 
    configure: CFG_LOCAL_RUST_ROOT  :=  
    configure: CFG_RUSTC            := /usr/local/bin/rustc (1.11.0-dev)
    configure: CFG_BUILD            := x86_64-unknown-linux-gnu 
    configure: CFG_HOST             := x86_64-unknown-linux-gnu 
    configure: CFG_TARGET           := x86_64-unknown-linux-gnu 
    configure: CFG_LOCALSTATEDIR    := /var/lib 
    configure: CFG_SYSCONFDIR       := /etc 
    configure: CFG_DATADIR          := /usr/local/share 
    configure: CFG_INFODIR          := /usr/local/share/info 
    configure: CFG_MANDIR           := /usr/local/share/man 
    configure: CFG_LIBDIR           := /usr/local/lib 
    configure: CFG_LOCAL_CARGO      :=  
    configure: 
    configure: validating ./configure args
    configure: 
    configure: 
    configure: looking for build programs
    configure: 
    configure: CFG_CURLORWGET       := /bin/curl (7.43.0)
    configure: CFG_PYTHON           := /bin/python 
    configure: CFG_CC               := /bin/cc (5.3.1)
    configure: CFG_RUSTDOC          := /usr/local/bin/rustdoc (1.11.0-dev)
    configure: 
    configure: writing configuration
    configure: 
    configure: CFG_SRC_DIR          := /data/src/github.com/rust-lang/carg ...
    configure: CFG_BUILD_DIR        := /data/src/github.com/rust-lang/carg ...
    configure: CFG_CONFIGURE_ARGS   :=  
    configure: CFG_PREFIX           := /usr/local 
    configure: CFG_BUILD            := x86_64-unknown-linux-gnu 
    configure: CFG_HOST             := x86_64-unknown-linux-gnu 
    configure: CFG_TARGET           := x86_64-unknown-linux-gnu 
    configure: CFG_LIBDIR           := /usr/local/lib 
    configure: CFG_MANDIR           := /usr/local/share/man 
    configure: CFG_RUSTC            := /usr/local/bin/rustc 
    configure: CFG_RUSTDOC          := /usr/local/bin/rustdoc 
    configure: 
    configure: cp /data/src/github.com/rust-lang/cargo/Makefile.in ./Makefile
    configure: mv config.tmp config.mk
    configure: 
    configure: complete
    configure: 
    configure: 

    [vagrant@localhost cargo]$ make 
    /bin/python /data/src/github.com/rust-lang/cargo//src/etc/dl-snapshot.py x86_64-unknown-linux-gnu
    running: curl -o target/dl/cargo-nightly-x86_64-unknown-linux-gnu.tar.gz https://static.rust-lang.org/cargo-dist/2016-03-21/cargo-nightly-x86_64-unknown-linux-gnu.tar.gz
    extracting cargo-nightly-x86_64-unknown-linux-gnu/cargo/manifest.in
    extracting cargo-nightly-x86_64-unknown-linux-gnu/cargo/bin/cargo
    extracting cargo-nightly-x86_64-unknown-linux-gnu/cargo/share/zsh/site-functions/_cargo
    extracting cargo-nightly-x86_64-unknown-linux-gnu/cargo/share/man/man1/cargo.1
    extracting cargo-nightly-x86_64-unknown-linux-gnu/cargo/share/doc/cargo/LICENSE-THIRD-PARTY
    extracting cargo-nightly-x86_64-unknown-linux-gnu/cargo/share/doc/cargo/LICENSE-APACHE
    extracting cargo-nightly-x86_64-unknown-linux-gnu/cargo/share/doc/cargo/LICENSE-MIT
    extracting cargo-nightly-x86_64-unknown-linux-gnu/cargo/share/doc/cargo/README.md
    extracting cargo-nightly-x86_64-unknown-linux-gnu/cargo/etc/bash_completion.d/cargo
    touch target/snapshot/bin/cargo
    /usr/local/bin/rustc -V
    rustc 1.11.0-dev (25f349db3 2016-06-19)
    target/snapshot/bin/cargo --version
    cargo 0.10.0-nightly (132b82d 2016-03-19)
    target/snapshot/bin/cargo build --target x86_64-unknown-linux-gnu --manifest-path /data/src/github.com/rust-lang/cargo//Cargo.toml --release   
        Updating registry `https://github.com/rust-lang/crates.io-index`
     Downloading advapi32-sys v0.1.2
     Downloading tar v0.4.5
     Downloading filetime v0.1.10
     Downloading git2 v0.4.3
     Downloading env_logger v0.3.2
     Downloading url v1.1.0
     Downloading term v0.4.4
     Downloading libgit2-sys v0.4.3
     Downloading crossbeam v0.2.8
     Downloading log v0.3.5
     Downloading tempdir v0.3.4
     Downloading num_cpus v0.2.11
     Downloading rustc-serialize v0.3.18
     Downloading semver v0.2.3
     Downloading fs2 v0.2.3
     Downloading git2-curl v0.5.0
     Downloading kernel32-sys v0.2.1
     Downloading winapi v0.2.6
     Downloading miow v0.1.2
     Downloading toml v0.1.30
     Downloading regex v0.1.58
     Downloading flate2 v0.2.13
     Downloading curl v0.3.0
     Downloading libc v0.2.8
     Downloading glob v0.2.11
     Downloading docopt v0.6.78
     Downloading winapi-build v0.1.1
     Downloading bitflags v0.1.1
     Downloading idna v0.1.0
     Downloading matches v0.1.2
     Downloading unicode-bidi v0.2.3
     Downloading unicode-normalization v0.1.2
     Downloading openssl-sys v0.7.8
     Downloading libssh2-sys v0.1.37
     Downloading libz-sys v1.0.2
     Downloading pkg-config v0.3.8
     Downloading gcc v0.3.26
     Downloading cmake v0.1.16
     Downloading utf8-ranges v0.1.3
     Downloading aho-corasick v0.5.1
     Downloading memchr v0.1.10
     Downloading regex-syntax v0.3.0
     Downloading rand v0.3.14
     Downloading curl-sys v0.2.0
     Downloading nom v1.2.2
     Downloading net2 v0.2.24
     Downloading ws2_32-sys v0.2.1
     Downloading cfg-if v0.1.0
     Downloading miniz-sys v0.1.7
     Downloading strsim v0.3.0
       Compiling cfg-if v0.1.0
       Compiling strsim v0.3.0
       Compiling matches v0.1.2
       Compiling unicode-bidi v0.2.3
       Compiling crossbeam v0.2.8
       Compiling pkg-config v0.3.8
       Compiling libc v0.2.8
       Compiling memchr v0.1.10
       Compiling aho-corasick v0.5.1
       Compiling rand v0.3.14
       Compiling net2 v0.2.24
       Compiling tempdir v0.3.4
       Compiling winapi v0.2.6
       Compiling unicode-normalization v0.1.2
       Compiling gcc v0.3.26
       Compiling miniz-sys v0.1.7
       Compiling curl-sys v0.2.0
       Compiling winapi-build v0.1.1
       Compiling ws2_32-sys v0.2.1
       Compiling glob v0.2.11
       Compiling advapi32-sys v0.1.2
       Compiling num_cpus v0.2.11
       Compiling regex-syntax v0.3.0
       Compiling nom v1.2.2
       Compiling semver v0.2.3
       Compiling openssl-sys v0.7.8
       Compiling cmake v0.1.16
       Compiling libgit2-sys v0.4.3
       Compiling libssh2-sys v0.1.37
       Compiling libz-sys v1.0.2
       Compiling filetime v0.1.10
       Compiling bitflags v0.1.1
       Compiling flate2 v0.2.13
       Compiling log v0.3.5
       Compiling idna v0.1.0
       Compiling url v1.1.0
       Compiling kernel32-sys v0.2.1
       Compiling fs2 v0.2.3
       Compiling term v0.4.4
       Compiling utf8-ranges v0.1.3
       Compiling regex v0.1.58
       Compiling env_logger v0.3.2
       Compiling git2 v0.4.3
       Compiling curl v0.3.0
       Compiling git2-curl v0.5.0
       Compiling rustc-serialize v0.3.18
       Compiling toml v0.1.30
       Compiling docopt v0.6.78
       Compiling crates-io v0.4.0 (file:///data/src/github.com/rust-lang/cargo/src/crates-io)
       Compiling tar v0.4.5
       Compiling miow v0.1.2
       Compiling cargo v0.12.0 (file:///data/src/github.com/rust-lang/cargo)

install 

    [vagrant@localhost cargo]$ sudo make install
    rm -rf target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image
    mkdir -p target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image/bin target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image/lib/cargo target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image/share/man/man1 target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image/share/doc/cargo target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image/share/zsh/site-functions target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image/etc/bash_completion.d
    cp target/x86_64-unknown-linux-gnu/release/cargo target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image/bin
    cp /data/src/github.com/rust-lang/cargo//src/etc/man/*.1 target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image/share/man/man1
    cp /data/src/github.com/rust-lang/cargo//src/etc/_cargo target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image/share/zsh/site-functions/_cargo
    cp /data/src/github.com/rust-lang/cargo//src/etc/cargo.bashcomp.sh target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image/etc/bash_completion.d/cargo
    cp /data/src/github.com/rust-lang/cargo//README.md /data/src/github.com/rust-lang/cargo//LICENSE-MIT /data/src/github.com/rust-lang/cargo//LICENSE-APACHE /data/src/github.com/rust-lang/cargo//LICENSE-THIRD-PARTY target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image/share/doc/cargo
    rm -Rf target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-overlay
    mkdir -p target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-overlay
    cp /data/src/github.com/rust-lang/cargo//README.md /data/src/github.com/rust-lang/cargo//LICENSE-MIT /data/src/github.com/rust-lang/cargo//LICENSE-APACHE /data/src/github.com/rust-lang/cargo//LICENSE-THIRD-PARTY target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-overlay
    echo "0.12.0 (4ba8264 2016-06-19)" > target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-overlay/version
    sh /data/src/github.com/rust-lang/cargo//src/rust-installer/gen-installer.sh --product-name=Rust --rel-manifest-dir=rustlib --success-message=Rust-is-ready-to-roll. --image-dir=target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image --work-dir=./target/x86_64-unknown-linux-gnu/release/dist --output-dir=./target/x86_64-unknown-linux-gnu/release/dist --non-installed-overlay=target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-overlay --package-name=cargo-0.12.0-x86_64-unknown-linux-gnu --component-name=cargo --legacy-manifest-dirs=rustlib,cargo
    gen-installer: looking for programs
    gen-installer: 
    gen-installer: found tar
    gen-installer: found cp
    gen-installer: found rm
    gen-installer: found mkdir
    gen-installer: found echo
    gen-installer: found tr
    gen-installer: found awk
    gen-installer: 
    gen-installer: processing arguments
    gen-installer: 
    gen-installer: CFG_PRODUCT_NAME     := Rust 
    gen-installer: CFG_COMPONENT_NAME   := cargo 
    gen-installer: CFG_PACKAGE_NAME     := cargo-0.12.0-x86_64-unknown-linux-g ...
    gen-installer: CFG_REL_MANIFEST_DIR := rustlib 
    gen-installer: CFG_SUCCESS_MESSAGE  := Rust-is-ready-to-roll. 
    gen-installer: CFG_LEGACY_MANIFEST_DIRS := rustlib,cargo 
    gen-installer: CFG_NON_INSTALLED_OVERLAY := target/x86_64-unknown-linux-gnu/rel ...
    gen-installer: CFG_BULK_DIRS        :=  
    gen-installer: CFG_IMAGE_DIR        := target/x86_64-unknown-linux-gnu/rel ...
    gen-installer: CFG_WORK_DIR         := ./target/x86_64-unknown-linux-gnu/r ...
    gen-installer: CFG_OUTPUT_DIR       := ./target/x86_64-unknown-linux-gnu/r ...
    gen-installer: 
    gen-installer: validating arguments
    gen-installer: 
    gen-install-script: looking for install programs
    gen-install-script: 
    gen-install-script: found sed
    gen-install-script: found chmod
    gen-install-script: found cat
    gen-install-script: 
    gen-install-script: processing arguments
    gen-install-script: 
    gen-install-script: CFG_PRODUCT_NAME     := Rust 
    gen-install-script: CFG_REL_MANIFEST_DIR := rustlib 
    gen-install-script: CFG_SUCCESS_MESSAGE  := Rust-is-ready-to-roll. 
    gen-install-script: CFG_OUTPUT_SCRIPT    := ./target/x86_64-unknown-linux-gnu/r ...
    gen-install-script: CFG_LEGACY_MANIFEST_DIRS := rustlib,cargo 
    gen-install-script: 
    gen-install-script: validating arguments
    gen-install-script: 
    rm -Rf target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu-image
    target/x86_64-unknown-linux-gnu/release/dist/cargo-0.12.0-x86_64-unknown-linux-gnu/install.sh --prefix="/usr/local" --destdir="" 
    install: creating uninstall script at /usr/local/lib/rustlib/uninstall.sh
    install: installing component 'cargo'
    
        Rust is ready to roll.

version

    [vagrant@localhost cargo]$ cargo --version
    cargo 0.12.0 (4ba8264 2016-06-19)
    [vagrant@localhost cargo]$ which cargo
    /usr/local/bin/cargo


