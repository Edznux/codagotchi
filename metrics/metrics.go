// +build !wasm

package metrics

import (
	"log"

	"github.com/DataDog/datadog-go/statsd"
)

var (
	Tags         = []string{"environment:dev"}
	statsdClient *statsd.Client
)

func init() {
	var err error
	statsdClient, err = statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Println(err)
		log.Println("Continuing anyway...")
	}
}

func Gauge(key string, value float64, tags []string, rate float64) {
	statsdClient.Gauge(key, value, tags, rate)
}

func Count(key string, value int64, tags []string, rate float64) {
	statsdClient.Count(key, value, tags, rate)
}
