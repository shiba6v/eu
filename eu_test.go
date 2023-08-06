package eu_test

import (
	"errors"
	"testing"

	"github.com/shiba6v/eu"
)

func TestWrap(t *testing.T) {
	error1 := errors.New("error1")
	f := func() error {
		return eu.Wrap(error1)
	}
	err := f()
	if errors.Is(err, error1) {
		// success
	} else {
		t.Log("fail to catch error")
		t.Fail()
	}
}
