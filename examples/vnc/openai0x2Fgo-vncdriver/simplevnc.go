package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/openai/go-vncdriver/gymvnc"
)
import _ "net/http/pprof"

type config struct {
	address  string
	password string
	encoding string
}

var vnc_config *config

func init() {
	vnc_config = &config{
		encoding: "tight",
	}

	flag.StringVar(&vnc_config.address, "addr", "", "vnc server")
	flag.StringVar(&vnc_config.password, "pass", "", "vnc password")
	flag.StringVar(&vnc_config.encoding, "enc", "tight", "vnc encoding")
}

func main() {
	flag.Parse()
	if vnc_config.address == "" {
		panic("vnc server required")
	}
	f, err := os.Create("/tmp/profile-tight.pprof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	gymvnc.ConfigureLogging()

	batch := gymvnc.NewVNCBatch()
	err = batch.Open("conn", gymvnc.VNCSessionConfig{
		Address:          vnc_config.address, /*"127.0.0.1:5900"*/
		Password:         vnc_config.password,
		Encoding:         vnc_config.encoding, /*"tight"*/
		FineQualityLevel: 100,
	})
	if err != nil {
		panic(err)
	}

	start := time.Now()
	updates := 0
	errs := 0
	for i := 0; i < 200000; i++ {
		elapsed := time.Now().Sub(start)
		if elapsed >= time.Duration(1)*time.Second {
			delta := float64(elapsed / time.Second)
			log.Printf("Received: updates=%.2f errs=%.2f", float64(updates)/delta, float64(errs)/delta)

			start = time.Now()
			updates = 0
			errs = 0
		}

		batchEvents := map[string][]gymvnc.VNCEvent{
			"conn": []gymvnc.VNCEvent{},
		}
		_, updatesN, errN := batch.Step(batchEvents)
		if errN["conn"] != nil {
			log.Fatalf("error: %+v", errN["conn"])
		}

		updates += len(updatesN["conn"])
		time.Sleep(16 * time.Millisecond)
	}

	// f, err := os.Create("/tmp/hi.prof")
	// if err != nil {
	//     log.Fatal("could not create memory profile: ", err)
	// }
	// runtime.GC() // get up-to-date statistics
	// if err := pprof.WriteHeapProfile(f); err != nil {
	//     log.Fatal("could not write memory profile: ", err)
	// }
	// f.Close()
}
