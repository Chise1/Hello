package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q,want %q.", got, want)
	}
}
func TestProberb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q,want %q.", got, want)
	}
}
func TestHelloV2(t *testing.T) {
	want := "Hello,I'm back."
	if got := HelloV2(); got != want {
		t.Errorf("HelloV2() = %q,want %q.", got, want)
	}
}
func TestHelloV3(t *testing.T) {
	want := "Hello,V3."
	if got := HelloV3(); got != want {
		t.Errorf("HelloV3() = %q,want %q.", got, want)
	}
}
