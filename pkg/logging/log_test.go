package logging

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    _log := StdOutLogger{}
    if got := _log.Log(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}