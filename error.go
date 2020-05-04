package maperr

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Priority ...
type Priority int

const (
	// PriorityHigh is high priority
	PriorityHigh Priority = 10
	// PriorityNormal is normal priority
	PriorityNormal = 0
	// PriorityLow is low priority
	PriorityLow = -10
)

// Error ...
type Error struct {
	Summary     string               `json:"summary"`
	Messages    map[string][]Message `json:"messages,omitempty"`
	MaxPriority Priority             `json:"-"`
	Total       int                  `json:"-"`
}

// Message ...
type Message struct {
	Text     string      `json:"text,omitempty"`
	Error    interface{} `json:"error,omitempty"`
	Priority Priority    `json:"-"`
	Kind     string      `json:"-"`
}

// AddSummary ...
func (e *Error) AddSummary(summary string) *Error {
	if e == nil {
		e = &Error{
			Summary:  summary,
			Messages: map[string][]Message{},
		}
	}
	e.Summary = summary
	e.MaxPriority = PriorityHigh
	return e
}

// Add ...
func (e *Error) Add(name string, message Message) *Error {
	if e == nil {
		e = &Error{
			Messages:    map[string][]Message{},
			MaxPriority: message.Priority,
		}
		if message.Text != "" {
			e.Summary = fmt.Sprintf("%s, %s", name, message.Text)
		}
	}

	e.Messages[name] = append(e.Messages[name], message)
	if e.MaxPriority < message.Priority && message.Text != "" {
		e.MaxPriority = message.Priority
		e.Summary = fmt.Sprintf("%s, %s", name, message.Text)
	}
	e.Total++
	return e
}

// Untyped ...
func (e *Error) Untyped() error {
	if e == nil {
		return nil
	}
	return e
}

// Error ...
func (e *Error) Error() string {
	var b strings.Builder
	if prettyPrint {
		e.writeMultiline(&b)
	} else {
		e.writeSingleline(&b)
	}
	return b.String()
}

// Format ...
func (e *Error) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, "Error -- ")
	if c == 'v' && f.Flag('+') {
		e.writeMultiline(f)
	} else {
		e.writeSingleline(f)
	}
}

func (e *Error) writeSingleline(w io.Writer) {
	if e == nil {
		fmt.Fprintf(w, `nil`)
		return
	}
	fmt.Fprintf(w, `%q (%d number of errors)`, e.Summary, e.Total)
}

func (e *Error) writeMultiline(w io.Writer) {
	if e == nil {
		fmt.Fprintf(w, `{}`)
		return
	}

	// the following errors occurred?
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(e); err != nil {
		fmt.Fprintf(w, `{"error": %q}`, err)
	}
}

// MarshalJSON ...
func (e *Error) MarshalJSON() ([]byte, error) {
	layout := DefaultLayout.Layout(e)
	return json.Marshal(layout)
}

// Layout ...
type Layout interface {
	Layout(*Error) interface{}
}

var prettyPrint bool
var DefaultLayout Layout

func init() {
	prettyPrint, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	DefaultLayout = &FlattenLayout{}
}
