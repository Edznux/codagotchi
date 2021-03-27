// +build wasm

package metrics

var (
	Tags = []string{"environment:dev"}
)

func Gauge(key string, value float64, tags []string, rate float64) {
	// noop in wasm (no monitoring client side)
}
func Count(key string, value float64, tags []string, rate float64) {
	// noop in wasm (no monitoring client side)
}
