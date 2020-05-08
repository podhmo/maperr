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

func TestFirstErrorSideEffect(t *testing.T) {
	{
		run := func() error {
			return (&maperr.Error{}).AddSummary("summary").Untyped()
		}
		if err := run(); err == nil {
			t.Errorf("want err, but nil")
		}
	}
	{
		run := func() error {
			var err *maperr.Error
			return err.AddSummary("summary").Untyped()
		}
		if err := run(); err == nil {
			t.Errorf("want err, but nil")
		}
	}
	{
		// FIXME
		run := func() error {
			var err *maperr.Error
			err.AddSummary("summary") // not err = err.AddSummary()
			return err.Untyped()
		}
		if err := run(); err != nil {
			t.Errorf("want nil, but err")
		}
	}
}
