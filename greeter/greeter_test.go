package greeter_test

import (
	"testing"

	greeter "github.com/pdougall1/golang-play/greeter"
)

func TestGreeter(t *testing.T) {
	t.Parallel()

	got := greeter.Greet()
	want := "Hello World"

	if got != want {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}
