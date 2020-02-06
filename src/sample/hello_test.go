package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	got := hello2()
	if got != want {
        t.Errorf("Hello2() = %q, want %q", got, want)
    }
}