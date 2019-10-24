package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/dghubble/sling"
	"github.com/pkg/errors"
	"github.com/tylertreat/bench"
)

type requesterFactory struct {
	ids []string
}


var (
	targetUrl = "http://localhost:3077"
	memprofile = ""
)

type requester struct {
	client *sling.Sling
	ids    []string
}

func (r *requester) Setup() error {
	r.client = sling.New().Client(&http.Client{
		Timeout: 15 * time.Second,
	}).Base(targetUrl)
	return nil
}

func (r *requester) Request() error {
	id := r.ids[rand.Intn(len(r.ids))]
	var result, errMsg json.RawMessage
	if _, err := r.client.New().Get("/tests/"+id).Receive(&result, &errMsg); err != nil {
		log.Printf("Error while making request to Alex: %v\n", err)
		return errors.Wrap(err, "Failed request to Alexandria")
	}
	if len(errMsg) > 0 {
		log.Printf("Error from Alex: %s\n", errMsg)
		return errors.Errorf("while doing GET %s", errMsg)
	}
	return nil
}

func (r *requester) Teardown() error {
	return nil
}

func (rf *requesterFactory) GetRequester(uint64) bench.Requester {
	return &requester{
		ids: rf.ids,
	}
}


func main() {
	r := &requesterFactory{ ids: []string{"hhh", "test", "ttss", "12312"} }

	benchmark := bench.NewBenchmark(r, 10, 10, 30*time.Second, 0)
	summary, err := benchmark.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(summary)
	err = summary.GenerateLatencyDistribution(nil, "summary.txt")
	if err != nil {
		log.Printf("something happended: %s", err)
	}

	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

	fmt.Printf("%+v\n", summary)
}