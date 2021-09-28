package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
  t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
    slowServer := makeDelayedServer(20 * time.Millisecond)
    fastServer := makeDelayedServer(0 * time.Millisecond)

    defer slowServer.Close()
    defer fastServer.Close()

    slowUrl := slowServer.URL
    fastUrl := fastServer.URL

    want := fastUrl
    got, err := Racer(slowUrl, fastUrl)

    if err != nil {
      t.Fatalf("did not expect error but got one %v", err)
    }

    if got != want {
      t.Errorf("got %q, want %q", got, want)
    }

    slowServer.Close()
    fastServer.Close()
  })

  t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
    server := makeDelayedServer(25 * time.Millisecond)

    defer server.Close()

    _, err := ConfigurableRacer(server.URL, server.URL, 20 * time.Millisecond)

    if err == nil {
      t.Error("expected error but didn't get one")
    }
  })
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(delay)
    w.WriteHeader(http.StatusOK)
  }))
}
