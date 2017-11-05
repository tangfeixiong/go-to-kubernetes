# Instruction

## Fedora23

bz2
```
[vagrant@localhost gorocksdb]$ sudo dnf install bzip2-devel
```

zstd
```
[vagrant@localhost gorocksdb]$ sudo dnf install libzstd-devel
```

lz4
```
[vagrant@localhost gorocksdb]$ sudo dnf install lz4-devel
```

snappy
```
[vagrant@localhost gorocksdb]$ sudo dnf install snappy-devel
```

__build__

```
[vagrant@localhost gorocksdb]$ CGO_FLAGS="-I/Users/fanhongling/Downloads/workspace/src/github.com/facebook/rocksdb/include" CGO_LDFLAGS="-L/Users/fanhongling/Downloads/workspace/src/github.com/facebook/rocksdb/ -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" go install -v .
github.com/tecbot/gorocksdb
```

or
```
[vagrant@localhost gorocksdb]$ CGO_FLAGS="-I/usr/local/include/rocksdb" CGO_LDFLAGS="-L/usr/local/lib/ -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" go install -v .
```

lib
```
[vagrant@localhost gorocksdb]$ ls ~/go/pkg/linux_amd64/github.com/tecbot/
gorocksdb.a
```