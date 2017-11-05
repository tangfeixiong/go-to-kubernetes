package main

import (
	"github.com/tecbot/gorocksdb"
)

func main() {
	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(gorocksdb.NewLRUCache(3 << 30))
	opts := gorocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, "/tmp/rocksdb-data")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	println(db.Name())

	ro := gorocksdb.NewDefaultReadOptions()
	wo := gorocksdb.NewDefaultWriteOptions()
	// if ro and wo are not used again, be sure to Close them.
	err = db.Put(wo, []byte("foo"), []byte("bar"))
	if err != nil {
		panic(err)
	}

	value, err := db.Get(ro, []byte("foo"))
	if err != nil {
		panic(err)
	}
	defer value.Free()
	println(string(value.Data()))

	err = db.Delete(wo, []byte("foo"))
	if err != nil {
		panic(err)
	}

}
