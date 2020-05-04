package maperr_test

import (
	"encoding/json"
	"testing"

	"github.com/podhmo/maperr"
)

func TestFullLayout(t *testing.T) {
	child := func() error {
		var err *maperr.Error
		return err.AddSummary("child error").
			Add("name", maperr.Message{Text: "required"}).
			Untyped()
	}
	parent := func() *maperr.Error {
		var err *maperr.Error
		return err.AddSummary("parent error").
			Add("name", maperr.Message{Text: "required"}).
			Add("info", maperr.Message{Error: child()})
	}

	b, err := json.MarshalIndent((&maperr.FullLayout{}).Layout(parent()), "", "  ")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	got := string(b)
	t.Logf("got: %v", got)

	want := `{
  "summary": "parent error",
  "messages": {
    "info": [
      {
        "error": {
          "summary": "child error",
          "messages": {
            "name": [
              {
                "text": "required"
              }
            ]
          }
        }
      }
    ],
    "name": [
      {
        "text": "required"
      }
    ]
  }
}`
	if want != got {
		t.Errorf("want\n\t%q\n\nbut got\n\t%q", want, got)
	}
}

func TestFlattenLayout(t *testing.T) {
	child := func() error {
		var err *maperr.Error
		return err.AddSummary("child error").
			Add("name", maperr.Message{Text: "required"}).
			Untyped()
	}
	parent := func() *maperr.Error {
		var err *maperr.Error
		return err.AddSummary("parent error").
			Add("name", maperr.Message{Text: "required"}).
			Add("info", maperr.Message{Error: child()})
	}

	b, err := json.MarshalIndent((&maperr.FlattenLayout{}).Layout(parent()), "", "  ")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	got := string(b)
	t.Logf("got: %v", got)

	want := `{
  "summary": "parent error",
  "messages": {
    "info.name": [
      {
        "text": "required"
      }
    ],
    "name": [
      {
        "text": "required"
      }
    ]
  }
}`
	if want != got {
		t.Errorf("want\n\t%q\n\nbut got\n\t%q", want, got)
	}
}
