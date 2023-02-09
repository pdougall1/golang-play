package dismisser_test

import (
	"testing"

	"github.com/pdougall1/golang-play/dismisser"
)

func TestDismisser(t *testing.T) {
	t.Parallel()

	got := dismisser.Bye()
	want := "Not Later bra"

	if got != want {
		t.Errorf("Got %s but wanted %s", got, want)
	}
}
