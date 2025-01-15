package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// aDuration := measureResponseTime(a)
	// bDuration := measureResponseTime(b)
	// if aDuration < bDuration {
	// 	return a
	// }
	// return b

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func measureResponseTime(url string) time.Duration {
	sTime := time.Now()
	http.Get(url)
	return time.Since(sTime)
}

func ping(url string) chan any {
	ch := make(chan any)
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
