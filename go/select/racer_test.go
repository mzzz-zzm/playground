package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speed of server, returning the url of the fastest one", func(t *testing.T) {
		slowSvr := makeDelayedServer(20 * time.Millisecond)
		defer slowSvr.Close() // stop listening a port after this function done

		fastSvr := makeDelayedServer(0 * time.Millisecond)
		defer fastSvr.Close()

		slowURL := slowSvr.URL
		fastURL := fastSvr.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 5s", func(t *testing.T) {
		aSrv := makeDelayedServer(25 * time.Millisecond)
		defer aSrv.Close() // stop listening a port after this function done

		_, err := ConfigurableRacer(aSrv.URL, aSrv.URL, 20*time.Millisecond)
		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}
	return httptest.NewServer(http.HandlerFunc(handler))
}
