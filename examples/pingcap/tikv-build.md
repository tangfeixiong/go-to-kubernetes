

## pingcap/tikv

fork and clone

    [vagrant@localhost github.com]$ git clone https://github.com/tangfeixiong/tikv pingcap/tikv
    正克隆到 'pingcap/tikv'...
    remote: Counting objects: 10582, done.
    remote: Compressing objects: 100% (7/7), done.
    remote: Total 10582 (delta 0), reused 0 (delta 0), pack-reused 10575
    接收对象中: 100% (10582/10582), 3.04 MiB | 531.00 KiB/s, 完成.
    处理 delta 中: 100% (7640/7640), 完成.
    检查连接... 完成。
    [vagrant@localhost github.com]$ cd pingcap/tikv/
    [vagrant@localhost tikv]$ git remote add upstream https://github.com/pingcap/tikv

build

    [vagrant@localhost tikv]$ cargo build --verbose
        Updating registry `https://github.com/rust-lang/crates.io-index`
        Updating git repository `https://github.com/ngaut/rust-rocksdb.git`
        Updating git repository `https://github.com/carllerche/bytes`
        Updating git repository `https://github.com/stepancheg/rust-protobuf.git`
        Updating git repository `https://github.com/pingcap/tipb.git`
        Updating git repository `https://github.com/pingcap/kvproto`
        Updating git repository `https://github.com/rust-lang-nursery/getopts`
        Updating git repository `https://github.com/carllerche/mio.git`
        Updating git repository `https://github.com/carllerche/nix-rust`
     Downloading crc v1.2.0
     Downloading log v0.3.6
     Downloading libc v0.2.11
     Downloading toml v0.1.28
     Downloading num v0.1.32
     Downloading rand v0.3.14
     Downloading quick-error v0.2.2
     Downloading tempdir v0.3.4
     Downloading cadence v0.5.1
     Downloading uuid v0.1.18
     Downloading byteorder v0.5.1
     Downloading threadpool v1.3.0
     Downloading fs2 v0.2.4
     Downloading time v0.1.35
     Downloading lazy_static v0.1.16
     Downloading rustc-serialize v0.3.19
     Downloading libc v0.1.12
     Downloading num-integer v0.1.32
     Downloading num-iter v0.1.32
     Downloading num-traits v0.1.32
     Downloading num-bigint v0.1.32
     Downloading num-rational v0.1.32
     Downloading num-complex v0.1.32
     Downloading winapi v0.2.6
     Downloading slab v0.1.3
     Downloading miow v0.1.2
     Downloading bytes v0.3.0
     Downloading net2 v0.2.23
     Downloading ws2_32-sys v0.2.1
     Downloading kernel32-sys v0.2.2
     Downloading winapi-build v0.1.1
     Downloading cfg-if v0.1.0
     Downloading bitflags v0.3.3
       Compiling libc v0.2.11
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/libc-0.2.11/src/lib.rs --crate-name libc --crate-type lib -g --cfg feature=\"use_std\" --cfg feature=\"default\" -C metadata=6b483f9a7097e9a4 -C extra-filename=-6b483f9a7097e9a4 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling bytes v0.3.0 (https://github.com/carllerche/bytes#12dfc417)
         Running `rustc /home/vagrant/.cargo/git/checkouts/bytes-72cf269819849e12/master/src/lib.rs --crate-name bytes --crate-type lib -g -C metadata=794ac4308adce037 -C extra-filename=-794ac4308adce037 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling rand v0.3.14
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/rand-0.3.14/src/lib.rs --crate-name rand --crate-type lib -g -C metadata=49a08859d086fffe -C extra-filename=-49a08859d086fffe --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --cap-lints allow`
       Compiling protobuf v1.0.20 (https://github.com/stepancheg/rust-protobuf.git#668cbcd4)
         Running `rustc /home/vagrant/.cargo/git/checkouts/rust-protobuf-c5cebfff504ea45b/master/src/lib/protobuf.rs --crate-name protobuf --crate-type lib -g -C metadata=6a49bfd14e8dbeda -C extra-filename=-6a49bfd14e8dbeda --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling tipb v0.0.1 (https://github.com/pingcap/tipb.git#4140a6e6)
         Running `rustc /home/vagrant/.cargo/git/checkouts/tipb-ef7de851bf2183d8/master/src/lib.rs --crate-name tipb --crate-type lib -g -C metadata=6ea80d2d576da521 -C extra-filename=-6ea80d2d576da521 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --cap-lints allow`
       Compiling winapi-build v0.1.1
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/winapi-build-0.1.1/src/lib.rs --crate-name build --crate-type lib -g -C metadata=493a7b0628804707 -C extra-filename=-493a7b0628804707 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling kernel32-sys v0.2.2
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/kernel32-sys-0.2.2/build.rs --crate-name build_script_build --crate-type bin -g --out-dir /data/src/github.com/pingcap/tikv/target/debug/build/kernel32-sys-d6afa5bd3d7cfaef --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern build=/data/src/github.com/pingcap/tikv/target/debug/deps/libbuild-493a7b0628804707.rlib --cap-lints allow`
       Compiling cfg-if v0.1.0
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/cfg-if-0.1.0/src/lib.rs --crate-name cfg_if --crate-type lib -g -C metadata=72c1f992b13d5087 -C extra-filename=-72c1f992b13d5087 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling quick-error v0.2.2
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/quick-error-0.2.2/src/lib.rs --crate-name quick_error --crate-type lib -g -C metadata=92f249c4e55991b4 -C extra-filename=-92f249c4e55991b4 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling libc v0.1.12
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/libc-0.1.12/rust/src/liblibc/lib.rs --crate-name libc --crate-type lib -g --cfg feature=\"default\" --cfg feature=\"cargo-build\" -C metadata=0c57a2126eb1e7e8 -C extra-filename=-0c57a2126eb1e7e8 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling tempdir v0.3.4
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/tempdir-0.3.4/src/lib.rs --crate-name tempdir --crate-type lib -g -C metadata=a136a238b008f647 -C extra-filename=-a136a238b008f647 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --cap-lints allow`
       Compiling lazy_static v0.1.16
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/lazy_static-0.1.16/src/lib.rs --crate-name lazy_static --crate-type lib -g -C metadata=3a04918be71c80ee -C extra-filename=-3a04918be71c80ee --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling bitflags v0.3.3
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/bitflags-0.3.3/src/lib.rs --crate-name bitflags --crate-type lib -g -C metadata=e61ad67c3301e77d -C extra-filename=-e61ad67c3301e77d --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling nix v0.5.0-pre (https://github.com/carllerche/nix-rust?rev=c4257f8a76#c4257f8a)
         Running `rustc /home/vagrant/.cargo/git/checkouts/nix-rust-d606302ceff0dca5/c4257f8a76/src/lib.rs --crate-name nix --crate-type lib -g -C metadata=f8d033ce6354dacb -C extra-filename=-f8d033ce6354dacb --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern bitflags=/data/src/github.com/pingcap/tikv/target/debug/deps/libbitflags-e61ad67c3301e77d.rlib --cap-lints allow`
       Compiling log v0.3.6
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/log-0.3.6/src/lib.rs --crate-name log --crate-type lib -g --cfg feature=\"use_std\" --cfg feature=\"default\" -C metadata=bf16bb9a4912b11d -C extra-filename=-bf16bb9a4912b11d --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling cadence v0.5.1
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/cadence-0.5.1/src/lib.rs --crate-name cadence --crate-type lib -g -C metadata=9576ac57842bcf06 -C extra-filename=-9576ac57842bcf06 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --cap-lints allow`
       Compiling getopts v0.2.14 (https://github.com/rust-lang-nursery/getopts#4dfc197a)
         Running `rustc /home/vagrant/.cargo/git/checkouts/getopts-984dbba6d7c8c8ca/master/src/lib.rs --crate-name getopts --crate-type lib -g -C metadata=82790631581b1d43 -C extra-filename=-82790631581b1d43 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling slab v0.1.3
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/slab-0.1.3/src/lib.rs --crate-name slab --crate-type lib -g -C metadata=4f8b9e9b6e35857c -C extra-filename=-4f8b9e9b6e35857c --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
         Running `/data/src/github.com/pingcap/tikv/target/debug/build/kernel32-sys-d6afa5bd3d7cfaef/build-script-build`
       Compiling num-traits v0.1.32
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/num-traits-0.1.32/src/lib.rs --crate-name num_traits --crate-type lib -g -C metadata=555e3a6260c26680 -C extra-filename=-555e3a6260c26680 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling num-integer v0.1.32
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/num-integer-0.1.32/src/lib.rs --crate-name num_integer --crate-type lib -g -C metadata=52fdddf28cd8e924 -C extra-filename=-52fdddf28cd8e924 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern num_traits=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_traits-555e3a6260c26680.rlib --cap-lints allow`
       Compiling num-iter v0.1.32
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/num-iter-0.1.32/src/lib.rs --crate-name num_iter --crate-type lib -g -C metadata=50df698bc905252c -C extra-filename=-50df698bc905252c --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern num_integer=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_integer-52fdddf28cd8e924.rlib --extern num_traits=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_traits-555e3a6260c26680.rlib --cap-lints allow`
       Compiling ws2_32-sys v0.2.1
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/ws2_32-sys-0.2.1/build.rs --crate-name build_script_build --crate-type bin -g --out-dir /data/src/github.com/pingcap/tikv/target/debug/build/ws2_32-sys-fe7373db9ed2332c --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern build=/data/src/github.com/pingcap/tikv/target/debug/deps/libbuild-493a7b0628804707.rlib --cap-lints allow`
         Running `/data/src/github.com/pingcap/tikv/target/debug/build/ws2_32-sys-fe7373db9ed2332c/build-script-build`
       Compiling winapi v0.2.6
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/winapi-0.2.6/src/lib.rs --crate-name winapi --crate-type lib -g -C metadata=0f5f77a5f89d9e92 -C extra-filename=-0f5f77a5f89d9e92 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/ws2_32-sys-0.2.1/src/lib.rs --crate-name ws2_32 --crate-type lib -g -C metadata=c73408aa15cb2b6f -C extra-filename=-c73408aa15cb2b6f --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern winapi=/data/src/github.com/pingcap/tikv/target/debug/deps/libwinapi-0f5f77a5f89d9e92.rlib --cap-lints allow`
       Compiling kvproto v0.0.1 (https://github.com/pingcap/kvproto#99be4b9d)
         Running `rustc /home/vagrant/.cargo/git/checkouts/kvproto-e694950bee90bc86/master/src/lib.rs --crate-name kvproto --crate-type lib -g -C metadata=b8a85b01f8205945 -C extra-filename=-b8a85b01f8205945 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --cap-lints allow`
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/kernel32-sys-0.2.2/src/lib.rs --crate-name kernel32 --crate-type lib -g -C metadata=df86a08647459244 -C extra-filename=-df86a08647459244 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern winapi=/data/src/github.com/pingcap/tikv/target/debug/deps/libwinapi-0f5f77a5f89d9e92.rlib --cap-lints allow`
       Compiling time v0.1.35
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/time-0.1.35/src/lib.rs --crate-name time --crate-type lib -g -C metadata=750bfdd52feafcb7 -C extra-filename=-750bfdd52feafcb7 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern winapi=/data/src/github.com/pingcap/tikv/target/debug/deps/libwinapi-0f5f77a5f89d9e92.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern kernel32=/data/src/github.com/pingcap/tikv/target/debug/deps/libkernel32-df86a08647459244.rlib --cap-lints allow`
       Compiling net2 v0.2.23
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/net2-0.2.23/src/lib.rs --crate-name net2 --crate-type lib -g -C metadata=3c7950c09adfba48 -C extra-filename=-3c7950c09adfba48 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern winapi=/data/src/github.com/pingcap/tikv/target/debug/deps/libwinapi-0f5f77a5f89d9e92.rlib --extern ws2_32=/data/src/github.com/pingcap/tikv/target/debug/deps/libws2_32-c73408aa15cb2b6f.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern kernel32=/data/src/github.com/pingcap/tikv/target/debug/deps/libkernel32-df86a08647459244.rlib --extern cfg_if=/data/src/github.com/pingcap/tikv/target/debug/deps/libcfg_if-72c1f992b13d5087.rlib --cap-lints allow`
       Compiling fs2 v0.2.4
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/fs2-0.2.4/src/lib.rs --crate-name fs2 --crate-type lib -g -C metadata=78ccc7fbf1e1ae18 -C extra-filename=-78ccc7fbf1e1ae18 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern winapi=/data/src/github.com/pingcap/tikv/target/debug/deps/libwinapi-0f5f77a5f89d9e92.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern kernel32=/data/src/github.com/pingcap/tikv/target/debug/deps/libkernel32-df86a08647459244.rlib --cap-lints allow`
       Compiling miow v0.1.2
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/miow-0.1.2/src/lib.rs --crate-name miow --crate-type lib -g -C metadata=b827bf6129dcd844 -C extra-filename=-b827bf6129dcd844 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern winapi=/data/src/github.com/pingcap/tikv/target/debug/deps/libwinapi-0f5f77a5f89d9e92.rlib --extern ws2_32=/data/src/github.com/pingcap/tikv/target/debug/deps/libws2_32-c73408aa15cb2b6f.rlib --extern kernel32=/data/src/github.com/pingcap/tikv/target/debug/deps/libkernel32-df86a08647459244.rlib --extern net2=/data/src/github.com/pingcap/tikv/target/debug/deps/libnet2-3c7950c09adfba48.rlib --cap-lints allow`
       Compiling bytes v0.3.0
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/bytes-0.3.0/src/lib.rs --crate-name bytes --crate-type lib -g -C metadata=f9fbfc75eb2416e0 -C extra-filename=-f9fbfc75eb2416e0 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling mio v0.5.0 (https://github.com/carllerche/mio.git#7d4778f2)
         Running `rustc /home/vagrant/.cargo/git/checkouts/mio-e79f285e2377d154/master/src/lib.rs --crate-name mio --crate-type lib -g -C metadata=2d9ad9925bada084 -C extra-filename=-2d9ad9925bada084 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern winapi=/data/src/github.com/pingcap/tikv/target/debug/deps/libwinapi-0f5f77a5f89d9e92.rlib --extern slab=/data/src/github.com/pingcap/tikv/target/debug/deps/libslab-4f8b9e9b6e35857c.rlib --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern miow=/data/src/github.com/pingcap/tikv/target/debug/deps/libmiow-b827bf6129dcd844.rlib --extern bytes=/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-f9fbfc75eb2416e0.rlib --extern net2=/data/src/github.com/pingcap/tikv/target/debug/deps/libnet2-3c7950c09adfba48.rlib --extern nix=/data/src/github.com/pingcap/tikv/target/debug/deps/libnix-f8d033ce6354dacb.rlib --extern time=/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib --cap-lints allow`
       Compiling crc v1.2.0
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/crc-1.2.0/src/lib.rs --crate-name crc --crate-type lib -g -C metadata=1fceeef3c778b3f2 -C extra-filename=-1fceeef3c778b3f2 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern lazy_static=/data/src/github.com/pingcap/tikv/target/debug/deps/liblazy_static-3a04918be71c80ee.rlib --cap-lints allow`
       Compiling librocksdb_sys v0.1.0 (https://github.com/ngaut/rust-rocksdb.git#c78df78c)
         Running `rustc /home/vagrant/.cargo/git/checkouts/rust-rocksdb-40347f12054c8f1a/master/librocksdb_sys/src/lib.rs --crate-name librocksdb_sys --crate-type lib -g -C metadata=f3a3d7f876a016af -C extra-filename=-f3a3d7f876a016af --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-0c57a2126eb1e7e8.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --cap-lints allow`
       Compiling rocksdb v0.3.0 (https://github.com/ngaut/rust-rocksdb.git#c78df78c)
         Running `rustc /home/vagrant/.cargo/git/checkouts/rust-rocksdb-40347f12054c8f1a/master/src/lib.rs --crate-name rocksdb --crate-type lib -g --cfg feature=\"default\" -C metadata=7a44dca17654e942 -C extra-filename=-7a44dca17654e942 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern librocksdb_sys=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibrocksdb_sys-f3a3d7f876a016af.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-0c57a2126eb1e7e8.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --cap-lints allow`
       Compiling threadpool v1.3.0
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/threadpool-1.3.0/src/lib.rs --crate-name threadpool --crate-type lib -g -C metadata=fad72996148b2f07 -C extra-filename=-fad72996148b2f07 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling rustc-serialize v0.3.19
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/rustc-serialize-0.3.19/src/lib.rs --crate-name rustc_serialize --crate-type lib -g -C metadata=3561541d79c18212 -C extra-filename=-3561541d79c18212 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling num-complex v0.1.32
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/num-complex-0.1.32/src/lib.rs --crate-name num_complex --crate-type lib -g --cfg feature=\"rustc-serialize\" --cfg feature=\"default\" -C metadata=6a8fb72fb99d524f -C extra-filename=-6a8fb72fb99d524f --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern rustc_serialize=/data/src/github.com/pingcap/tikv/target/debug/deps/librustc_serialize-3561541d79c18212.rlib --extern num_traits=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_traits-555e3a6260c26680.rlib --cap-lints allow`
       Compiling toml v0.1.28
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/toml-0.1.28/src/lib.rs --crate-name toml --crate-type lib -g --cfg feature=\"rustc-serialize\" --cfg feature=\"default\" -C metadata=88c2092368f85c4f -C extra-filename=-88c2092368f85c4f --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern rustc_serialize=/data/src/github.com/pingcap/tikv/target/debug/deps/librustc_serialize-3561541d79c18212.rlib --cap-lints allow`
       Compiling num-bigint v0.1.32
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/num-bigint-0.1.32/src/lib.rs --crate-name num_bigint --crate-type lib -g --cfg feature=\"rustc-serialize\" --cfg feature=\"rand\" --cfg feature=\"default\" -C metadata=0f2c1e47b214fd65 -C extra-filename=-0f2c1e47b214fd65 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern num_integer=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_integer-52fdddf28cd8e924.rlib --extern rustc_serialize=/data/src/github.com/pingcap/tikv/target/debug/deps/librustc_serialize-3561541d79c18212.rlib --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --extern num_traits=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_traits-555e3a6260c26680.rlib --cap-lints allow`
       Compiling uuid v0.1.18
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/uuid-0.1.18/src/lib.rs --crate-name uuid --crate-type lib -g -C metadata=13c4decae1c5ee4b -C extra-filename=-13c4decae1c5ee4b --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern rustc_serialize=/data/src/github.com/pingcap/tikv/target/debug/deps/librustc_serialize-3561541d79c18212.rlib --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --cap-lints allow`
       Compiling byteorder v0.5.1
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/byteorder-0.5.1/src/lib.rs --crate-name byteorder --crate-type lib -g -C metadata=c882ee50a4e0aa4f -C extra-filename=-c882ee50a4e0aa4f --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --cap-lints allow`
       Compiling num-rational v0.1.32
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/num-rational-0.1.32/src/lib.rs --crate-name num_rational --crate-type lib -g --cfg feature=\"rustc-serialize\" --cfg feature=\"num-bigint\" --cfg feature=\"default\" --cfg feature=\"bigint\" -C metadata=576f79d570f1b9fd -C extra-filename=-576f79d570f1b9fd --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern num_integer=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_integer-52fdddf28cd8e924.rlib --extern rustc_serialize=/data/src/github.com/pingcap/tikv/target/debug/deps/librustc_serialize-3561541d79c18212.rlib --extern num_traits=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_traits-555e3a6260c26680.rlib --extern num_bigint=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_bigint-0f2c1e47b214fd65.rlib --cap-lints allow`
       Compiling num v0.1.32
         Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/num-0.1.32/src/lib.rs --crate-name num --crate-type lib -g --cfg feature=\"num-rational\" --cfg feature=\"num-complex\" --cfg feature=\"rustc-serialize\" --cfg feature=\"num-bigint\" --cfg feature=\"default\" --cfg feature=\"complex\" --cfg feature=\"bigint\" --cfg feature=\"rational\" -C metadata=f6961afa9fa1a02b -C extra-filename=-f6961afa9fa1a02b --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern num_integer=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_integer-52fdddf28cd8e924.rlib --extern num_iter=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_iter-50df698bc905252c.rlib --extern num_traits=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_traits-555e3a6260c26680.rlib --extern num_bigint=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_bigint-0f2c1e47b214fd65.rlib --extern num_rational=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_rational-576f79d570f1b9fd.rlib --extern num_complex=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_complex-6a8fb72fb99d524f.rlib --cap-lints allow`
       Compiling tikv v0.0.1 (file:///data/src/github.com/pingcap/tikv)
         Running `rustc src/lib.rs --crate-name tikv --crate-type lib -g --cfg feature=\"default\" --out-dir /data/src/github.com/pingcap/tikv/target/debug --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern getopts=/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib --extern crc=/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern rocksdb=/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib --extern num=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib --extern tipb=/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --extern kvproto=/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --extern quick_error=/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib --extern mio=/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --extern cadence=/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib --extern uuid=/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib --extern byteorder=/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib --extern bytes=/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib --extern threadpool=/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib --extern fs2=/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib --extern time=/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib`
         Running `rustc src/bin/tikv-server.rs --crate-name tikv_server --crate-type bin -g --cfg feature=\"default\" --out-dir /data/src/github.com/pingcap/tikv/target/debug --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern getopts=/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib --extern crc=/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern rocksdb=/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib --extern num=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib --extern tipb=/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --extern kvproto=/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --extern quick_error=/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib --extern mio=/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --extern cadence=/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib --extern uuid=/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib --extern byteorder=/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib --extern bytes=/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib --extern threadpool=/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib --extern fs2=/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib --extern time=/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib --extern tikv=/data/src/github.com/pingcap/tikv/target/debug/libtikv.rlib`
    error: linking with `cc` failed: exit code: 1
    note: "cc" "-Wl,--as-needed" "-Wl,-z,noexecstack" "-m64" "-L" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib" "/data/src/github.com/pingcap/tikv/target/debug/tikv_server.0.o" "-o" "/data/src/github.com/pingcap/tikv/target/debug/tikv_server" "-Wl,--gc-sections" "-pie" "-nodefaultlibs" "-L" "/data/src/github.com/pingcap/tikv/target/debug" "-L" "/data/src/github.com/pingcap/tikv/target/debug/deps" "-L" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib" "-Wl,-Bstatic" "-Wl,-Bdynamic" "/data/src/github.com/pingcap/tikv/target/debug/libtikv.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_iter-50df698bc905252c.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_complex-6a8fb72fb99d524f.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libmiow-b827bf6129dcd844.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libslab-4f8b9e9b6e35857c.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnix-f8d033ce6354dacb.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnet2-3c7950c09adfba48.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblazy_static-3a04918be71c80ee.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libbitflags-e61ad67c3301e77d.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libtest-fe3cdf61.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libcfg_if-72c1f992b13d5087.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libkernel32-df86a08647459244.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_rational-576f79d570f1b9fd.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_bigint-0f2c1e47b214fd65.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-f9fbfc75eb2416e0.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libgetopts-fe3cdf61.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libwinapi-0f5f77a5f89d9e92.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_integer-52fdddf28cd8e924.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_traits-555e3a6260c26680.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libterm-fe3cdf61.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libws2_32-c73408aa15cb2b6f.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/librustc_serialize-3561541d79c18212.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblibrocksdb_sys-f3a3d7f876a016af.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-0c57a2126eb1e7e8.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libstd-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcollections-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libpanic_unwind-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_unicode-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libunwind-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/librand-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc_jemalloc-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/liblibc-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcore-fe3cdf61.rlib" "-l" "util" "-l" "rocksdb" "-l" "c" "-l" "m" "-l" "dl" "-l" "pthread" "-l" "gcc_s" "-l" "pthread" "-l" "c" "-l" "m" "-l" "rt" "-l" "util" "-l" "compiler-rt"
    note: /bin/ld: cannot find -lrocksdb
    collect2: 错误：ld 返回 1
    
    error: aborting due to previous error
    error: Could not compile `tikv`.
    
    Caused by:
      Process didn't exit successfully: `rustc src/bin/tikv-server.rs --crate-name tikv_server --crate-type bin -g --cfg feature="default" --out-dir /data/src/github.com/pingcap/tikv/target/debug --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern getopts=/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib --extern crc=/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern rocksdb=/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib --extern num=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib --extern tipb=/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --extern kvproto=/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --extern quick_error=/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib --extern mio=/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --extern cadence=/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib --extern uuid=/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib --extern byteorder=/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib --extern bytes=/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib --extern threadpool=/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib --extern fs2=/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib --extern time=/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib --extern tikv=/data/src/github.com/pingcap/tikv/target/debug/libtikv.rlib` (exit code: 101)

merge upstream

    [vagrant@localhost tikv]$ git pull upstream master
    来自 https://github.com/pingcap/tikv
     * branch            master     -> FETCH_HEAD
    更新 b59c9df..67c112e
    Fast-forward

...

rebuild

    [vagrant@localhost tikv]$ cargo build --verbose
           Fresh num-traits v0.1.32
           Fresh log v0.3.6
           Fresh bytes v0.3.0
           Fresh num-integer v0.1.32
           Fresh bytes v0.3.0 (https://github.com/carllerche/bytes#12dfc417)
           Fresh rustc-serialize v0.3.19
           Fresh toml v0.1.28
           Fresh getopts v0.2.14 (https://github.com/rust-lang-nursery/getopts#4dfc197a)
           Fresh libc v0.2.11
           Fresh rand v0.3.14
           Fresh uuid v0.1.18
           Fresh num-iter v0.1.32
           Fresh tempdir v0.3.4
           Fresh protobuf v1.0.20 (https://github.com/stepancheg/rust-protobuf.git#668cbcd4)
       Compiling kvproto v0.0.1 (https://github.com/pingcap/kvproto#301260d7)
         Running `rustc /home/vagrant/.cargo/git/checkouts/kvproto-e694950bee90bc86/master/src/lib.rs --crate-name kvproto --crate-type lib -g -C metadata=b8a85b01f8205945 -C extra-filename=-b8a85b01f8205945 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --cap-lints allow`
           Fresh threadpool v1.3.0
           Fresh lazy_static v0.1.16
           Fresh cfg-if v0.1.0
           Fresh libc v0.1.12
       Compiling librocksdb_sys v0.1.0 (https://github.com/ngaut/rust-rocksdb.git#7de3c903)
         Running `rustc /home/vagrant/.cargo/git/checkouts/rust-rocksdb-40347f12054c8f1a/master/librocksdb_sys/src/lib.rs --crate-name librocksdb_sys --crate-type lib -g -C metadata=f3a3d7f876a016af -C extra-filename=-f3a3d7f876a016af --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-0c57a2126eb1e7e8.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --cap-lints allow`
       Compiling rocksdb v0.3.0 (https://github.com/ngaut/rust-rocksdb.git#7de3c903)
         Running `rustc /home/vagrant/.cargo/git/checkouts/rust-rocksdb-40347f12054c8f1a/master/src/lib.rs --crate-name rocksdb --crate-type lib -g --cfg feature=\"default\" -C metadata=7a44dca17654e942 -C extra-filename=-7a44dca17654e942 --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-0c57a2126eb1e7e8.rlib --extern librocksdb_sys=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibrocksdb_sys-f3a3d7f876a016af.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --cap-lints allow`
           Fresh tipb v0.0.1 (https://github.com/pingcap/tipb.git#4140a6e6)
           Fresh crc v1.2.0
           Fresh byteorder v0.5.1
           Fresh cadence v0.5.1
           Fresh quick-error v0.2.2
           Fresh num-complex v0.1.32
           Fresh winapi v0.2.6
           Fresh num-bigint v0.1.32
           Fresh num-rational v0.1.32
           Fresh num v0.1.32
           Fresh slab v0.1.3
           Fresh bitflags v0.3.3
           Fresh nix v0.5.0-pre (https://github.com/carllerche/nix-rust?rev=c4257f8a76#c4257f8a)
           Fresh winapi-build v0.1.1
           Fresh ws2_32-sys v0.2.1
           Fresh kernel32-sys v0.2.2
           Fresh net2 v0.2.23
           Fresh miow v0.1.2
           Fresh time v0.1.35
           Fresh mio v0.5.0 (https://github.com/carllerche/mio.git#7d4778f2)
           Fresh fs2 v0.2.4
       Compiling tikv v0.0.1 (file:///data/src/github.com/pingcap/tikv)
         Running `rustc src/lib.rs --crate-name tikv --crate-type lib -g --cfg feature=\"default\" --out-dir /data/src/github.com/pingcap/tikv/target/debug --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern byteorder=/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib --extern threadpool=/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib --extern getopts=/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib --extern rocksdb=/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib --extern bytes=/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern crc=/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib --extern kvproto=/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib --extern uuid=/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib --extern quick_error=/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib --extern num=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib --extern time=/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --extern fs2=/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --extern tipb=/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib --extern cadence=/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib --extern mio=/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib`
         Running `rustc src/bin/tikv-server.rs --crate-name tikv_server --crate-type bin -g --cfg feature=\"default\" --out-dir /data/src/github.com/pingcap/tikv/target/debug --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern byteorder=/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib --extern threadpool=/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib --extern getopts=/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib --extern rocksdb=/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib --extern bytes=/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern crc=/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib --extern kvproto=/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib --extern uuid=/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib --extern quick_error=/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib --extern num=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib --extern time=/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --extern fs2=/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --extern tipb=/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib --extern cadence=/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib --extern mio=/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib --extern tikv=/data/src/github.com/pingcap/tikv/target/debug/libtikv.rlib`
    error: linking with `cc` failed: exit code: 1
    note: "cc" "-Wl,--as-needed" "-Wl,-z,noexecstack" "-m64" "-L" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib" "/data/src/github.com/pingcap/tikv/target/debug/tikv_server.0.o" "-o" "/data/src/github.com/pingcap/tikv/target/debug/tikv_server" "-Wl,--gc-sections" "-pie" "-nodefaultlibs" "-L" "/data/src/github.com/pingcap/tikv/target/debug" "-L" "/data/src/github.com/pingcap/tikv/target/debug/deps" "-L" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib" "-Wl,-Bstatic" "-Wl,-Bdynamic" "/data/src/github.com/pingcap/tikv/target/debug/libtikv.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_iter-50df698bc905252c.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_complex-6a8fb72fb99d524f.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libmiow-b827bf6129dcd844.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libslab-4f8b9e9b6e35857c.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnix-f8d033ce6354dacb.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnet2-3c7950c09adfba48.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblazy_static-3a04918be71c80ee.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libbitflags-e61ad67c3301e77d.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libtest-fe3cdf61.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libcfg_if-72c1f992b13d5087.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libkernel32-df86a08647459244.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_rational-576f79d570f1b9fd.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_bigint-0f2c1e47b214fd65.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-f9fbfc75eb2416e0.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libgetopts-fe3cdf61.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libwinapi-0f5f77a5f89d9e92.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_integer-52fdddf28cd8e924.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_traits-555e3a6260c26680.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libterm-fe3cdf61.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libws2_32-c73408aa15cb2b6f.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/librustc_serialize-3561541d79c18212.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblibrocksdb_sys-f3a3d7f876a016af.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-0c57a2126eb1e7e8.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libstd-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcollections-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libpanic_unwind-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_unicode-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libunwind-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/librand-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc_jemalloc-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/liblibc-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcore-fe3cdf61.rlib" "-l" "util" "-l" "rocksdb" "-l" "c" "-l" "m" "-l" "dl" "-l" "pthread" "-l" "gcc_s" "-l" "pthread" "-l" "c" "-l" "m" "-l" "rt" "-l" "util" "-l" "compiler-rt"
    note: /bin/ld: cannot find -lrocksdb
    collect2: 错误：ld 返回 1
    
    error: aborting due to previous error
    error: Could not compile `tikv`.
    
    Caused by:
      Process didn't exit successfully: `rustc src/bin/tikv-server.rs --crate-name tikv_server --crate-type bin -g --cfg feature="default" --out-dir /data/src/github.com/pingcap/tikv/target/debug --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern byteorder=/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib --extern threadpool=/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib --extern getopts=/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib --extern rocksdb=/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib --extern bytes=/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern crc=/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib --extern kvproto=/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib --extern uuid=/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib --extern quick_error=/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib --extern num=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib --extern time=/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --extern fs2=/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --extern tipb=/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib --extern cadence=/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib --extern mio=/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib --extern tikv=/data/src/github.com/pingcap/tikv/target/debug/libtikv.rlib` (exit code: 101)

enable feature of dev

    [vagrant@localhost tikv]$ make build
    cargo build --features dev
     Downloading clippy v0.0.73
     Downloading regex-syntax v0.3.1
     Downloading semver v0.2.3
     Downloading clippy_lints v0.0.73
     Downloading quine-mc_cluskey v0.2.2
     Downloading unicode-normalization v0.1.2
     Downloading nom v1.2.2
     Downloading matches v0.1.2
       Compiling nom v1.2.2
       Compiling semver v0.2.3
       Compiling regex-syntax v0.3.1
       Compiling unicode-normalization v0.1.2
       Compiling quine-mc_cluskey v0.2.2
       Compiling matches v0.1.2
       Compiling clippy_lints v0.0.73
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/drop_ref.rs:38:64: 38:70 error: no method named `def_id` found for type `rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/drop_ref.rs:38                 let def_id = cx.tcx.def_map.borrow()[&path.id].def_id();
                                                                                                                                                                        ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/enum_glob_use.rs:47:76: 47:82 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/enum_glob_use.rs:47                     if let Some(node_id) = cx.tcx.map.as_local_node_id(def.def_id()) {
                                                                                                                                                                                         ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/enum_glob_use.rs:54:72: 54:78 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/enum_glob_use.rs:54                         let child = cx.sess().cstore.item_children(def.def_id());
                                                                                                                                                                                     ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:72:39: 72:45 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:72                 !used_in_expr(cx, def.def_id(), cond),
                                                                                                                                                 ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:64:13: 126:15 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:73:56: 73:62 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:73                 let Some(value) = check_assign(cx, def.def_id(), then),
                                                                                                                                                                  ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:71:12: 71:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:64:13: 126:15 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:74:39: 74:45 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:74                 !used_in_expr(cx, def.def_id(), value),
                                                                                                                                                 ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:71:12: 71:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:64:13: 126:15 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:80:69: 80:75 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:80                         if let Some(default) = check_assign(cx, def.def_id(), else_) {
                                                                                                                                                                               ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:71:12: 71:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:64:13: 126:15 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:142:28: 142:34 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:142             self.id == def.def_id(),
                                                                                                                                       ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:139:9: 146:11 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:158:21: 158:27 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:158         decl == def.def_id(),
                                                                                                                                ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:152:5: 175:7 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/mem_forget.rs:30:69: 30:75 error: no method named `def_id` found for type `rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/mem_forget.rs:30                 let def_id = cx.tcx.def_map.borrow()[&path_expr.id].def_id();
                                                                                                                                                                               ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/minmax.rs:58:60: 58:66 error: no method named `def_id` found for type `rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/minmax.rs:58             let def_id = cx.tcx.def_map.borrow()[&path.id].def_id();
                                                                                                                                                                  ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:16: 185:19 error: mismatched types [E0308]
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185         val == 0.0 || val == ::std::f64::INFINITY || val == ::std::f64::NEG_INFINITY
                                                                                                                     ^~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:16: 185:19 help: run `rustc --explain E0308` to see a detailed explanation
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:16: 185:19 note: expected type `rustc_const_math::ConstFloat`
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:16: 185:19 note:    found type `_`
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:30: 185:50 error: mismatched types [E0308]
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185         val == 0.0 || val == ::std::f64::INFINITY || val == ::std::f64::NEG_INFINITY
                                                                                                                                   ^~~~~~~~~~~~~~~~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:30: 185:50 help: run `rustc --explain E0308` to see a detailed explanation
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:30: 185:50 note: expected type `rustc_const_math::ConstFloat`
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:30: 185:50 note:    found type `f64`
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:61: 185:85 error: mismatched types [E0308]
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185         val == 0.0 || val == ::std::f64::INFINITY || val == ::std::f64::NEG_INFINITY
                                                                                                                                                                  ^~~~~~~~~~~~~~~~~~~~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:61: 185:85 help: run `rustc --explain E0308` to see a detailed explanation
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:61: 185:85 note: expected type `rustc_const_math::ConstFloat`
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:61: 185:85 note:    found type `f64`
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:396:39: 396:50 error: no method named `unhygienize` found for type `syntax::ast::Name` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:396                    segment != segment.unhygienize() && // not in bang macro
                                                                                                                                            ^~~~~~~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/needless_bool.rs:158:10: 158:12 error: expected an array or slice, found `&[syntax::codemap::Spanned<rustc::hir::Stmt_>]` [E0529]
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/needless_bool.rs:158         ([], Some(e)) => fetch_bool_expr(&**e),
                                                                                                                        ^~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/needless_bool.rs:158:10: 158:12 help: the semantics of slice patterns changed recently; see issue #23121
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/needless_bool.rs:159:10: 159:17 error: expected an array or slice, found `&[syntax::codemap::Spanned<rustc::hir::Stmt_>]` [E0529]
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/needless_bool.rs:159         ([ref e], None) => {
                                                                                                                        ^~~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/needless_bool.rs:159:10: 159:17 help: the semantics of slice patterns changed recently; see issue #23121
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/regex.rs:106:30: 106:36 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/regex.rs:106             let def_id = def.def_id();
                                                                                                                                    ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:71:12: 71:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/regex.rs:101:9: 120:11 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:69:39: 69:50 error: no method named `unhygienize` found for type `syntax::ast::Name` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:69             bindings.push((ident.node.unhygienize(), ident.span))
                                                                                                                                             ^~~~~~~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:123:35: 123:46 error: no method named `unhygienize` found for type `syntax::ast::Name` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:123             let name = ident.node.unhygienize();
                                                                                                                                          ^~~~~~~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:330:71: 330:82 error: no method named `unhygienize` found for type `syntax::ast::Name` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:330     !path.global && path.segments.len() == 1 && path.segments[0].name.unhygienize() == name
                                                                                                                                                                              ^~~~~~~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:340:30: 340:41 error: no method named `unhygienize` found for type `syntax::ast::Name` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:340         if self.name == name.unhygienize() {
                                                                                                                                     ^~~~~~~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/transmute.rs:63:69: 63:75 error: no method named `def_id` found for type `rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/transmute.rs:63                 let def_id = cx.tcx.def_map.borrow()[&path_expr.id].def_id();
                                                                                                                                                                              ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:59:29: 59:35 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:59                 if Some(did.def_id()) == cx.tcx.lang_items.owned_box() {
                                                                                                                                  ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:68:52: 68:58 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:68                             match_def_path(cx, did.def_id(), &paths::VEC),
                                                                                                                                                         ^~~~~~
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:60:21: 77:22 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:78:50: 78:56 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
    /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:78                 } else if match_def_path(cx, did.def_id(), &paths::LINKED_LIST) {
                                                                                                                                                       ^~~~~~
    error: aborting due to 26 previous errors
    error: Could not compile `clippy_lints`.
    
    To learn more, run the command again with --verbose.
    Makefile:8: recipe for target 'build' failed
    make: *** [build] Error 101



## branch

push local master

    [vagrant@localhost tikv]$ git status
    位于分支 master
    您的分支领先 'origin/master' 共 6 个提交。
      （使用 "git push" 来发布您的本地提交）
    无文件要提交，干净的工作区
    [vagrant@localhost tikv]$ git push
    warning: push.default 尚未设置，它的默认值在 Git 2.0 已从 'matching'
    变更为 'simple'。若要不再显示本信息并保持传统习惯，进行如下设置：
    
      git config --global push.default matching
    
    若要不再显示本信息并从现在开始采用新的使用习惯，设置：
    
      git config --global push.default simple
    
    当 push.default 设置为 'matching' 后，git 将推送和远程同名的所有
    本地分支。
    
    从 Git 2.0 开始，Git 缺省采用更为保守的 'simple' 模式，只推送当前
    分支到远程关联的同名分支，即 'git push' 推送当前分支。
    
    参见 'git help config' 并查找 'push.default' 以获取更多信息。
    （'simple' 模式由 Git 1.7.11 版本引入。如果您有时要使用老版本的 Git，
    为保持兼容，请用 'current' 代替 'simple'）
    
    Username for 'https://github.com': tangfeixiong
    Password for 'https://tangfeixiong@github.com': 
    对象计数中: 106, 完成.
    压缩对象中: 100% (44/44), 完成.
    写入对象中: 100% (106/106), 20.00 KiB | 0 bytes/s, 完成.
    Total 106 (delta 86), reused 79 (delta 60)
    To https://github.com/tangfeixiong/tikv
       b59c9df..67c112e  master -> master


### CI

june-21

[vagrant@localhost tikv]$ cargo build -v --features dev
       Fresh regex-syntax v0.3.1
       Fresh libc v0.1.12
       Fresh bitflags v0.3.3
       Fresh winapi v0.2.6
       Fresh matches v0.1.2
       Fresh quick-error v0.2.2
       Fresh cfg-if v0.1.0
       Fresh unicode-normalization v0.1.2
       Fresh log v0.3.6
       Fresh cadence v0.5.1
       Fresh bytes v0.3.0 (https://github.com/carllerche/bytes#12dfc417)
       Fresh quine-mc_cluskey v0.2.2
       Fresh threadpool v1.3.0
       Fresh nom v1.2.2
       Fresh semver v0.2.3
       Fresh winapi-build v0.1.1
       Fresh kernel32-sys v0.2.2
       Fresh ws2_32-sys v0.2.1
       Fresh bytes v0.3.0
       Fresh num-traits v0.1.32
       Fresh num-integer v0.1.32
       Fresh lazy_static v0.1.16
       Fresh crc v1.2.0
       Fresh libc v0.2.11
       Fresh rand v0.3.14
       Fresh fs2 v0.2.4
       Fresh tempdir v0.3.4
       Fresh librocksdb_sys v0.1.0 (https://github.com/ngaut/rust-rocksdb.git#7de3c903)
       Fresh rocksdb v0.3.0 (https://github.com/ngaut/rust-rocksdb.git#7de3c903)
       Fresh net2 v0.2.23
       Fresh miow v0.1.2
       Fresh byteorder v0.5.1
       Fresh rustc-serialize v0.3.19
       Fresh uuid v0.1.18
       Fresh toml v0.1.28
       Fresh num-bigint v0.1.32
       Fresh num-rational v0.1.32
       Fresh num-complex v0.1.32
   Compiling clippy_lints v0.0.73
     Running `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/lib.rs --crate-name clippy_lints --crate-type lib -g -C metadata=822afa863b4a505f -C extra-filename=-822afa863b4a505f --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern unicode_normalization=/data/src/github.com/pingcap/tikv/target/debug/deps/libunicode_normalization-5de3a9c8fd8ddf4e.rlib --extern rustc_serialize=/data/src/github.com/pingcap/tikv/target/debug/deps/librustc_serialize-3561541d79c18212.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern quine_mc_cluskey=/data/src/github.com/pingcap/tikv/target/debug/deps/libquine_mc_cluskey-bfa4946e33f51135.rlib --extern semver=/data/src/github.com/pingcap/tikv/target/debug/deps/libsemver-60d1aa0e68346373.rlib --extern matches=/data/src/github.com/pingcap/tikv/target/debug/deps/libmatches-5d5580ffd528031c.rlib --extern regex_syntax=/data/src/github.com/pingcap/tikv/target/debug/deps/libregex_syntax-88bc923660c879aa.rlib --cap-lints allow`
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/drop_ref.rs:38:64: 38:70 error: no method named `def_id` found for type `rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/drop_ref.rs:38                 let def_id = cx.tcx.def_map.borrow()[&path.id].def_id();
                                                                                                                                                                    ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/enum_glob_use.rs:47:76: 47:82 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/enum_glob_use.rs:47                     if let Some(node_id) = cx.tcx.map.as_local_node_id(def.def_id()) {
                                                                                                                                                                                     ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/enum_glob_use.rs:54:72: 54:78 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/enum_glob_use.rs:54                         let child = cx.sess().cstore.item_children(def.def_id());
                                                                                                                                                                                 ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:72:39: 72:45 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:72                 !used_in_expr(cx, def.def_id(), cond),
                                                                                                                                             ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:64:13: 126:15 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:73:56: 73:62 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:73                 let Some(value) = check_assign(cx, def.def_id(), then),
                                                                                                                                                              ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:71:12: 71:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:64:13: 126:15 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:74:39: 74:45 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:74                 !used_in_expr(cx, def.def_id(), value),
                                                                                                                                             ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:71:12: 71:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:64:13: 126:15 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:80:69: 80:75 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:80                         if let Some(default) = check_assign(cx, def.def_id(), else_) {
                                                                                                                                                                           ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:71:12: 71:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:64:13: 126:15 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:142:28: 142:34 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:142             self.id == def.def_id(),
                                                                                                                                   ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:139:9: 146:11 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:158:21: 158:27 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:158         decl == def.def_id(),
                                                                                                                            ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/let_if_seq.rs:152:5: 175:7 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/mem_forget.rs:30:69: 30:75 error: no method named `def_id` found for type `rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/mem_forget.rs:30                 let def_id = cx.tcx.def_map.borrow()[&path_expr.id].def_id();
                                                                                                                                                                           ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/minmax.rs:58:60: 58:66 error: no method named `def_id` found for type `rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/minmax.rs:58             let def_id = cx.tcx.def_map.borrow()[&path.id].def_id();
                                                                                                                                                              ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:16: 185:19 error: mismatched types [E0308]
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185         val == 0.0 || val == ::std::f64::INFINITY || val == ::std::f64::NEG_INFINITY
                                                                                                                 ^~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:16: 185:19 help: run `rustc --explain E0308` to see a detailed explanation
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:16: 185:19 note: expected type `rustc_const_math::ConstFloat`
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:16: 185:19 note:    found type `_`
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:30: 185:50 error: mismatched types [E0308]
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185         val == 0.0 || val == ::std::f64::INFINITY || val == ::std::f64::NEG_INFINITY
                                                                                                                               ^~~~~~~~~~~~~~~~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:30: 185:50 help: run `rustc --explain E0308` to see a detailed explanation
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:30: 185:50 note: expected type `rustc_const_math::ConstFloat`
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:30: 185:50 note:    found type `f64`
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:61: 185:85 error: mismatched types [E0308]
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185         val == 0.0 || val == ::std::f64::INFINITY || val == ::std::f64::NEG_INFINITY
                                                                                                                                                              ^~~~~~~~~~~~~~~~~~~~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:61: 185:85 help: run `rustc --explain E0308` to see a detailed explanation
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:61: 185:85 note: expected type `rustc_const_math::ConstFloat`
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:185:61: 185:85 note:    found type `f64`
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:396:39: 396:50 error: no method named `unhygienize` found for type `syntax::ast::Name` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/misc.rs:396                    segment != segment.unhygienize() && // not in bang macro
                                                                                                                                        ^~~~~~~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/needless_bool.rs:158:10: 158:12 error: expected an array or slice, found `&[syntax::codemap::Spanned<rustc::hir::Stmt_>]` [E0529]
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/needless_bool.rs:158         ([], Some(e)) => fetch_bool_expr(&**e),
                                                                                                                    ^~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/needless_bool.rs:158:10: 158:12 help: the semantics of slice patterns changed recently; see issue #23121
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/needless_bool.rs:159:10: 159:17 error: expected an array or slice, found `&[syntax::codemap::Spanned<rustc::hir::Stmt_>]` [E0529]
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/needless_bool.rs:159         ([ref e], None) => {
                                                                                                                    ^~~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/needless_bool.rs:159:10: 159:17 help: the semantics of slice patterns changed recently; see issue #23121
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/regex.rs:106:30: 106:36 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/regex.rs:106             let def_id = def.def_id();
                                                                                                                                ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:71:12: 71:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/regex.rs:101:9: 120:11 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:69:39: 69:50 error: no method named `unhygienize` found for type `syntax::ast::Name` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:69             bindings.push((ident.node.unhygienize(), ident.span))
                                                                                                                                         ^~~~~~~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:123:35: 123:46 error: no method named `unhygienize` found for type `syntax::ast::Name` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:123             let name = ident.node.unhygienize();
                                                                                                                                      ^~~~~~~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:330:71: 330:82 error: no method named `unhygienize` found for type `syntax::ast::Name` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:330     !path.global && path.segments.len() == 1 && path.segments[0].name.unhygienize() == name
                                                                                                                                                                          ^~~~~~~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:340:30: 340:41 error: no method named `unhygienize` found for type `syntax::ast::Name` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/shadow.rs:340         if self.name == name.unhygienize() {
                                                                                                                                 ^~~~~~~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/transmute.rs:63:69: 63:75 error: no method named `def_id` found for type `rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/transmute.rs:63                 let def_id = cx.tcx.def_map.borrow()[&path_expr.id].def_id();
                                                                                                                                                                          ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:59:29: 59:35 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:59                 if Some(did.def_id()) == cx.tcx.lang_items.owned_box() {
                                                                                                                              ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:68:52: 68:58 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:68                             match_def_path(cx, did.def_id(), &paths::VEC),
                                                                                                                                                     ^~~~~~
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs:56:12: 56:46 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:60:21: 77:22 note: in this expansion of if_let_chain! (defined in /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/utils/mod.rs)
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:78:50: 78:56 error: no method named `def_id` found for type `&rustc::hir::def::PathResolution` in the current scope
/home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/types.rs:78                 } else if match_def_path(cx, did.def_id(), &paths::LINKED_LIST) {
                                                                                                                                                   ^~~~~~
error: aborting due to 26 previous errors
error: Could not compile `clippy_lints`.

Caused by:
  Process didn't exit successfully: `rustc /home/vagrant/.cargo/registry/src/github.com-1ecc6299db9ec823/clippy_lints-0.0.73/src/lib.rs --crate-name clippy_lints --crate-type lib -g -C metadata=822afa863b4a505f -C extra-filename=-822afa863b4a505f --out-dir /data/src/github.com/pingcap/tikv/target/debug/deps --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern unicode_normalization=/data/src/github.com/pingcap/tikv/target/debug/deps/libunicode_normalization-5de3a9c8fd8ddf4e.rlib --extern rustc_serialize=/data/src/github.com/pingcap/tikv/target/debug/deps/librustc_serialize-3561541d79c18212.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern quine_mc_cluskey=/data/src/github.com/pingcap/tikv/target/debug/deps/libquine_mc_cluskey-bfa4946e33f51135.rlib --extern semver=/data/src/github.com/pingcap/tikv/target/debug/deps/libsemver-60d1aa0e68346373.rlib --extern matches=/data/src/github.com/pingcap/tikv/target/debug/deps/libmatches-5d5580ffd528031c.rlib --extern regex_syntax=/data/src/github.com/pingcap/tikv/target/debug/deps/libregex_syntax-88bc923660c879aa.rlib --cap-lints allow` (exit code: 101)
[vagrant@localhost tikv]$ cargo build -v
       Fresh cfg-if v0.1.0
       Fresh lazy_static v0.1.16
       Fresh crc v1.2.0
       Fresh winapi v0.2.6
       Fresh bitflags v0.3.3
       Fresh bytes v0.3.0
       Fresh slab v0.1.3
       Fresh libc v0.1.12
       Fresh libc v0.2.11
       Fresh nix v0.5.0-pre (https://github.com/carllerche/nix-rust?rev=c4257f8a76#c4257f8a)
       Fresh winapi-build v0.1.1
       Fresh kernel32-sys v0.2.2
       Fresh fs2 v0.2.4
       Fresh num-traits v0.1.32
       Fresh log v0.3.6
       Fresh cadence v0.5.1
       Fresh time v0.1.35
       Fresh rand v0.3.14
       Fresh bytes v0.3.0 (https://github.com/carllerche/bytes#12dfc417)
       Fresh protobuf v1.0.20 (https://github.com/stepancheg/rust-protobuf.git#668cbcd4)
       Fresh tipb v0.0.1 (https://github.com/pingcap/tipb.git#4140a6e6)
       Fresh getopts v0.2.14 (https://github.com/rust-lang-nursery/getopts#4dfc197a)
       Fresh rustc-serialize v0.3.19
       Fresh num-complex v0.1.32
       Fresh uuid v0.1.18
       Fresh toml v0.1.28
       Fresh num-integer v0.1.32
       Fresh num-iter v0.1.32
       Fresh num-bigint v0.1.32
       Fresh num-rational v0.1.32
       Fresh ws2_32-sys v0.2.1
       Fresh net2 v0.2.23
       Fresh miow v0.1.2
       Fresh mio v0.5.0 (https://github.com/carllerche/mio.git#7d4778f2)
       Fresh byteorder v0.5.1
       Fresh threadpool v1.3.0
       Fresh num v0.1.32
       Fresh tempdir v0.3.4
       Fresh librocksdb_sys v0.1.0 (https://github.com/ngaut/rust-rocksdb.git#7de3c903)
       Fresh rocksdb v0.3.0 (https://github.com/ngaut/rust-rocksdb.git#7de3c903)
       Fresh kvproto v0.0.1 (https://github.com/pingcap/kvproto#301260d7)
       Fresh quick-error v0.2.2
   Compiling tikv v0.0.1 (file:///data/src/github.com/pingcap/tikv)
     Running `rustc src/bin/tikv-server.rs --crate-name tikv_server --crate-type bin -g --cfg feature=\"default\" --out-dir /data/src/github.com/pingcap/tikv/target/debug --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern tipb=/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib --extern fs2=/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --extern getopts=/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib --extern cadence=/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib --extern kvproto=/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib --extern bytes=/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib --extern crc=/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --extern time=/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --extern mio=/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib --extern quick_error=/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib --extern num=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern rocksdb=/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib --extern byteorder=/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib --extern uuid=/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib --extern threadpool=/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib --extern tikv=/data/src/github.com/pingcap/tikv/target/debug/libtikv.rlib`
error: linking with `cc` failed: exit code: 1
note: "cc" "-Wl,--as-needed" "-Wl,-z,noexecstack" "-m64" "-L" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib" "/data/src/github.com/pingcap/tikv/target/debug/tikv_server.0.o" "-o" "/data/src/github.com/pingcap/tikv/target/debug/tikv_server" "-Wl,--gc-sections" "-pie" "-nodefaultlibs" "-L" "/data/src/github.com/pingcap/tikv/target/debug" "-L" "/data/src/github.com/pingcap/tikv/target/debug/deps" "-L" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib" "-Wl,-Bstatic" "-Wl,-Bdynamic" "/data/src/github.com/pingcap/tikv/target/debug/libtikv.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_iter-50df698bc905252c.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_complex-6a8fb72fb99d524f.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libmiow-b827bf6129dcd844.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libslab-4f8b9e9b6e35857c.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnix-f8d033ce6354dacb.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnet2-3c7950c09adfba48.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblazy_static-3a04918be71c80ee.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libbitflags-e61ad67c3301e77d.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libtest-fe3cdf61.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libcfg_if-72c1f992b13d5087.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libkernel32-df86a08647459244.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_rational-576f79d570f1b9fd.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_bigint-0f2c1e47b214fd65.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-f9fbfc75eb2416e0.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libgetopts-fe3cdf61.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libwinapi-0f5f77a5f89d9e92.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_integer-52fdddf28cd8e924.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libnum_traits-555e3a6260c26680.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libterm-fe3cdf61.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libws2_32-c73408aa15cb2b6f.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/librustc_serialize-3561541d79c18212.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblibrocksdb_sys-f3a3d7f876a016af.rlib" "/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-0c57a2126eb1e7e8.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libstd-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcollections-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libpanic_unwind-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/librustc_unicode-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libunwind-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/librand-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/liballoc_jemalloc-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/liblibc-fe3cdf61.rlib" "/usr/local/lib/rustlib/x86_64-unknown-linux-gnu/lib/libcore-fe3cdf61.rlib" "-l" "util" "-l" "rocksdb" "-l" "c" "-l" "m" "-l" "dl" "-l" "pthread" "-l" "gcc_s" "-l" "pthread" "-l" "c" "-l" "m" "-l" "rt" "-l" "util" "-l" "compiler-rt"
note: /bin/ld: /usr/local/lib/librocksdb.a(c.o): relocation R_X86_64_32S against `_ZTVZ42rocksdb_slicetransform_create_fixed_prefixE7Wrapper' can not be used when making a shared object; recompile with -fPIC
/usr/local/lib/librocksdb.a: error adding symbols: 错误的值
collect2: 错误：ld 返回 1

error: aborting due to previous error
error: Could not compile `tikv`.

Caused by:
  Process didn't exit successfully: `rustc src/bin/tikv-server.rs --crate-name tikv_server --crate-type bin -g --cfg feature="default" --out-dir /data/src/github.com/pingcap/tikv/target/debug --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern tipb=/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib --extern fs2=/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --extern getopts=/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib --extern cadence=/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib --extern kvproto=/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib --extern bytes=/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib --extern crc=/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --extern time=/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --extern mio=/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib --extern quick_error=/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib --extern num=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern rocksdb=/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib --extern byteorder=/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib --extern uuid=/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib --extern threadpool=/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib --extern tikv=/data/src/github.com/pingcap/tikv/target/debug/libtikv.rlib` (exit code: 101)

after install rocksdb shared_lib

    [vagrant@localhost tikv]$ cargo build -v
           Fresh winapi v0.2.6
           Fresh slab v0.1.3
           Fresh winapi-build v0.1.1
           Fresh num-traits v0.1.32
           Fresh bytes v0.3.0 (https://github.com/carllerche/bytes#12dfc417)
           Fresh threadpool v1.3.0
           Fresh lazy_static v0.1.16
           Fresh crc v1.2.0
           Fresh num-integer v0.1.32
           Fresh num-iter v0.1.32
           Fresh ws2_32-sys v0.2.1
           Fresh cfg-if v0.1.0
           Fresh quick-error v0.2.2
           Fresh libc v0.2.11
           Fresh rand v0.3.14
           Fresh tempdir v0.3.4
           Fresh log v0.3.6
           Fresh cadence v0.5.1
           Fresh rustc-serialize v0.3.19
           Fresh num-complex v0.1.32
           Fresh toml v0.1.28
           Fresh uuid v0.1.18
           Fresh num-bigint v0.1.32
           Fresh num-rational v0.1.32
           Fresh num v0.1.32
           Fresh bytes v0.3.0
           Fresh byteorder v0.5.1
           Fresh protobuf v1.0.20 (https://github.com/stepancheg/rust-protobuf.git#668cbcd4)
           Fresh tipb v0.0.1 (https://github.com/pingcap/tipb.git#4140a6e6)
           Fresh bitflags v0.3.3
           Fresh kvproto v0.0.1 (https://github.com/pingcap/kvproto#301260d7)
           Fresh libc v0.1.12
           Fresh librocksdb_sys v0.1.0 (https://github.com/ngaut/rust-rocksdb.git#7de3c903)
           Fresh getopts v0.2.14 (https://github.com/rust-lang-nursery/getopts#4dfc197a)
           Fresh rocksdb v0.3.0 (https://github.com/ngaut/rust-rocksdb.git#7de3c903)
           Fresh kernel32-sys v0.2.2
           Fresh net2 v0.2.23
           Fresh time v0.1.35
           Fresh miow v0.1.2
           Fresh fs2 v0.2.4
           Fresh nix v0.5.0-pre (https://github.com/carllerche/nix-rust?rev=c4257f8a76#c4257f8a)
           Fresh mio v0.5.0 (https://github.com/carllerche/mio.git#7d4778f2)
       Compiling tikv v0.0.1 (file:///data/src/github.com/pingcap/tikv)
         Running `rustc src/bin/tikv-server.rs --crate-name tikv_server --crate-type bin -g --cfg feature=\"default\" --out-dir /data/src/github.com/pingcap/tikv/target/debug --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --extern tipb=/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib --extern getopts=/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib --extern rocksdb=/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern mio=/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib --extern threadpool=/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --extern time=/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib --extern fs2=/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib --extern byteorder=/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib --extern num=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib --extern quick_error=/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib --extern cadence=/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib --extern bytes=/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --extern crc=/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --extern kvproto=/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib --extern uuid=/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern tikv=/data/src/github.com/pingcap/tikv/target/debug/libtikv.rlib`
         Running `rustc benches/bin/bench-tikv.rs --crate-name bench_tikv --crate-type bin -g --cfg feature=\"default\" --out-dir /data/src/github.com/pingcap/tikv/target/debug --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --extern tipb=/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib --extern getopts=/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib --extern rocksdb=/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern mio=/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib --extern threadpool=/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --extern time=/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib --extern fs2=/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib --extern byteorder=/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib --extern num=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib --extern quick_error=/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib --extern cadence=/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib --extern bytes=/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --extern crc=/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --extern kvproto=/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib --extern uuid=/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern tikv=/data/src/github.com/pingcap/tikv/target/debug/libtikv.rlib`
         Running `rustc src/bin/tikv-dump.rs --crate-name tikv_dump --crate-type bin -g --cfg feature=\"default\" --out-dir /data/src/github.com/pingcap/tikv/target/debug --emit=dep-info,link -L dependency=/data/src/github.com/pingcap/tikv/target/debug -L dependency=/data/src/github.com/pingcap/tikv/target/debug/deps --extern rand=/data/src/github.com/pingcap/tikv/target/debug/deps/librand-49a08859d086fffe.rlib --extern tipb=/data/src/github.com/pingcap/tikv/target/debug/deps/libtipb-6ea80d2d576da521.rlib --extern getopts=/data/src/github.com/pingcap/tikv/target/debug/deps/libgetopts-82790631581b1d43.rlib --extern rocksdb=/data/src/github.com/pingcap/tikv/target/debug/deps/librocksdb-7a44dca17654e942.rlib --extern libc=/data/src/github.com/pingcap/tikv/target/debug/deps/liblibc-6b483f9a7097e9a4.rlib --extern mio=/data/src/github.com/pingcap/tikv/target/debug/deps/libmio-2d9ad9925bada084.rlib --extern threadpool=/data/src/github.com/pingcap/tikv/target/debug/deps/libthreadpool-fad72996148b2f07.rlib --extern protobuf=/data/src/github.com/pingcap/tikv/target/debug/deps/libprotobuf-6a49bfd14e8dbeda.rlib --extern time=/data/src/github.com/pingcap/tikv/target/debug/deps/libtime-750bfdd52feafcb7.rlib --extern fs2=/data/src/github.com/pingcap/tikv/target/debug/deps/libfs2-78ccc7fbf1e1ae18.rlib --extern byteorder=/data/src/github.com/pingcap/tikv/target/debug/deps/libbyteorder-c882ee50a4e0aa4f.rlib --extern num=/data/src/github.com/pingcap/tikv/target/debug/deps/libnum-f6961afa9fa1a02b.rlib --extern quick_error=/data/src/github.com/pingcap/tikv/target/debug/deps/libquick_error-92f249c4e55991b4.rlib --extern cadence=/data/src/github.com/pingcap/tikv/target/debug/deps/libcadence-9576ac57842bcf06.rlib --extern bytes=/data/src/github.com/pingcap/tikv/target/debug/deps/libbytes-794ac4308adce037.rlib --extern tempdir=/data/src/github.com/pingcap/tikv/target/debug/deps/libtempdir-a136a238b008f647.rlib --extern crc=/data/src/github.com/pingcap/tikv/target/debug/deps/libcrc-1fceeef3c778b3f2.rlib --extern log=/data/src/github.com/pingcap/tikv/target/debug/deps/liblog-bf16bb9a4912b11d.rlib --extern kvproto=/data/src/github.com/pingcap/tikv/target/debug/deps/libkvproto-b8a85b01f8205945.rlib --extern uuid=/data/src/github.com/pingcap/tikv/target/debug/deps/libuuid-13c4decae1c5ee4b.rlib --extern toml=/data/src/github.com/pingcap/tikv/target/debug/deps/libtoml-88c2092368f85c4f.rlib --extern tikv=/data/src/github.com/pingcap/tikv/target/debug/libtikv.rlib`
        
    [vagrant@localhost tikv]$ ls -ltr target/debug/
    总用量 130308
    drwxrwxr-x. 2 vagrant vagrant     4096 6月  20 18:55 native
    drwxrwxr-x. 2 vagrant vagrant     4096 6月  20 18:55 examples
    drwxrwxr-x. 4 vagrant vagrant     4096 6月  20 18:57 build
    -rw-rw-r--. 1 vagrant vagrant 24513678 6月  20 19:11 libtikv.rlib
    drwxrwxr-x. 2 vagrant vagrant     4096 6月  20 19:27 deps
    -rwxrwxr-x. 1 vagrant vagrant 42391088 6月  21 10:24 tikv-server
    -rwxrwxr-x. 1 vagrant vagrant 43915480 6月  21 10:25 bench-tikv
    -rwxrwxr-x. 1 vagrant vagrant 22589864 6月  21 10:25 tikv-dump        
        



[vagrant@localhost tikv]$ target/debug/tikv-server 
target/debug/tikv-server: error while loading shared libraries: librocksdb.so.4.9: cannot open shared object file: No such file or directory


[vagrant@localhost tikv]$ LD_LIBRARY_PATH=/usr/local/lib ./target/debug/tikv-server 
2016-06-21 15:23:32,735 tikv-server.rs:77 - INFO  - malformed or missing server.addr, use default
2016-06-21 15:23:32,737 tikv-server.rs:448 - INFO  - Start listening on 127.0.0.1:20160...
2016-06-21 15:23:32,741 tikv-server.rs:77 - INFO  - malformed or missing server.dsn, use default
2016-06-21 15:23:32,745 tikv-server.rs:77 - INFO  - malformed or missing raft.cluster-id, use default
thread 'main' panicked at 'please specify raft.cluster-id', src/libcore/option.rs:699
note: Run with `RUST_BACKTRACE=1` for a backtrace.


[vagrant@localhost tikv]$ LD_LIBRARY_PATH=/usr/local/lib RUST_BACKTRACE=1 ./target/debug/tikv-server -I 1
2016-06-21 15:40:07,625 tikv-server.rs:77 - INFO  - malformed or missing server.addr, use default
2016-06-21 15:40:07,629 tikv-server.rs:448 - INFO  - Start listening on 127.0.0.1:20160...
2016-06-21 15:40:07,630 tikv-server.rs:77 - INFO  - malformed or missing server.dsn, use default
2016-06-21 15:40:07,633 tikv-server.rs:101 - INFO  - malformed or missing server.notify-capacity, use default
2016-06-21 15:40:07,633 tikv-server.rs:101 - INFO  - malformed or missing server.capacity, use default
2016-06-21 15:40:07,633 tikv-server.rs:77 - INFO  - malformed or missing server.advertise-addr, use default
2016-06-21 15:40:07,634 tikv-server.rs:101 - INFO  - malformed or missing raftstore.notify-capacity, use default
2016-06-21 15:40:07,634 tikv-server.rs:77 - INFO  - malformed or missing metric.addr, use default
2016-06-21 15:40:07,634 tikv-server.rs:77 - INFO  - malformed or missing metric.prefix, use default
2016-06-21 15:40:07,635 tikv-server.rs:77 - INFO  - malformed or missing server.store, use default
2016-06-21 15:40:07,641 rocksdb.rs:34 - INFO  - EngineRocksdb: creating for path 
2016-06-21 15:40:07,655 mod.rs:187 - INFO  - storage: [Rocksdb] started.
2016-06-21 15:40:07,777 mod.rs:154 - INFO  - starting working thread: end-point-worker
2016-06-21 15:40:07,779 mod.rs:154 - INFO  - starting working thread: snap-handler
^Z
[1]+  已停止               LD_LIBRARY_PATH=/usr/local/lib RUST_BACKTRACE=1 ./target/debug/tikv-server -I 1
[vagrant@localhost tikv]$ bg %1
[1]+ LD_LIBRARY_PATH=/usr/local/lib RUST_BACKTRACE=1 ./target/debug/tikv-server -I 1 &


[vagrant@localhost tikv]$ sudo netstat -tpnl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 127.0.0.1:20160         0.0.0.0:*               LISTEN      19565/./target/debu 
tcp        0      0 127.0.0.1:4001          0.0.0.0:*               LISTEN      4745/etcd           
tcp        0      0 127.0.0.1:10248         0.0.0.0:*               LISTEN      4762/kubelet        
tcp        0      0 127.0.0.1:10249         0.0.0.0:*               LISTEN      5119/hyperkube      
tcp        0      0 172.17.4.50:10250       0.0.0.0:*               LISTEN      4762/kubelet        
tcp        0      0 127.0.0.1:2379          0.0.0.0:*               LISTEN      4745/etcd           
tcp        0      0 127.0.0.1:2380          0.0.0.0:*               LISTEN      4745/etcd           
tcp        0      0 172.17.4.50:10255       0.0.0.0:*               LISTEN      4762/kubelet        
tcp        0      0 127.0.0.1:8080          0.0.0.0:*               LISTEN      5247/hyperkube      
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      707/sshd            
tcp        0      0 127.0.0.1:7001          0.0.0.0:*               LISTEN      4745/etcd           
tcp        0      0 172.17.4.50:443         0.0.0.0:*               LISTEN      5247/hyperkube      
tcp6       0      0 :::32124                :::*                    LISTEN      5119/hyperkube      
tcp6       0      0 :::32125                :::*                    LISTEN      5119/hyperkube      
tcp6       0      0 :::32126                :::*                    LISTEN      5119/hyperkube      
tcp6       0      0 :::32127                :::*                    LISTEN      5119/hyperkube      
tcp6       0      0 :::4194                 :::*                    LISTEN      4762/kubelet        
tcp6       0      0 :::10251                :::*                    LISTEN      5061/hyperkube      
tcp6       0      0 :::10252                :::*                    LISTEN      5191/hyperkube      
tcp6       0      0 :::30448                :::*                    LISTEN      5119/hyperkube      
tcp6       0      0 :::22                   :::*                    LISTEN      707/sshd            
tcp6       0      0 :::32121                :::*                    LISTEN      5119/hyperkube      
tcp6       0      0 :::32122                :::*                    LISTEN      5119/hyperkube      
tcp6       0      0 :::32123                :::*                    LISTEN      5119/hyperkube      
