package maperr_test

import (
	"testing"

	"github.com/podhmo/maperr"
)

func TestUntyped(t *testing.T) {
	{
		run := func() error {
			var err *maperr.Error
			return err
		}
		if err := run(); err == nil {
			t.Errorf("typed error is expected, but untyped error is got")
		}
	}
	{
		run := func() error {
			var err *maperr.Error
			return err.Untyped()
		}
		if err := run(); err != nil {
			t.Errorf("untyped error is expected, but typed error is got")
		}
	}
}
