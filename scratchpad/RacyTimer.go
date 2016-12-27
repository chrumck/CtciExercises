package scratchpad

import (
    "time"
    "fmt"
    "math/rand"
)

// RacyTimer is an exercise in race conditions
func RacyTimer() {
    start := time.Now()
    var t *time.Timer
    t = time.AfterFunc(
        randomDuration(),
        func() {
            fmt.Println(time.Since(start))
            t.Reset(randomDuration())
        })
    time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
     return time.Duration(rand.Int63n(1e9))
}
