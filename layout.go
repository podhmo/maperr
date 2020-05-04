package maperr

import "fmt"

// FullLayout ...
type FullLayout struct {
	Summary     string               `json:"summary"`
	Messages    map[string][]Message `json:"messages,omitempty"`
	MaxPriority Priority             `json:"-"`
	Total       int                  `json:"-"`
}

// Layout ...
func (v *FullLayout) Layout(err *Error) interface{} {
	layout := &FullLayout{
		Summary:     err.Summary,
		Messages:    map[string][]Message{},
		MaxPriority: err.MaxPriority,
		Total:       err.Total,
	}

	for name, messages := range err.Messages {
		var buf []Message
		for _, m := range messages {
			if m.Error == nil {
				buf = append(buf, m)
				continue
			}

			err, ok := m.Error.(*Error)
			if !ok {
				buf = append(buf, m)
				continue
			}

			buf = append(buf, Message{
				Text:     m.Text,
				Error:    v.Layout(err),
				Priority: m.Priority,
				Kind:     m.Kind,
			})
		}
		layout.Messages[name] = buf
	}
	return layout
}

// FlattenLayout ...
type FlattenLayout struct {
	Summary     string               `json:"summary"`
	Messages    map[string][]Message `json:"messages,omitempty"`
	MaxPriority Priority             `json:"-"`
	Total       int                  `json:"-"`
}

// Layout ...
func (v *FlattenLayout) Layout(err *Error) interface{} {
	return v.layout(err)
}
func (v *FlattenLayout) layout(err *Error) *FlattenLayout {
	layout := &FlattenLayout{
		Summary:     err.Summary,
		Messages:    map[string][]Message{},
		MaxPriority: err.MaxPriority,
		Total:       err.Total,
	}

	for name, messages := range err.Messages {
		var buf []Message
		for _, m := range messages {
			if m.Error == nil {
				buf = append(buf, m)
				continue
			}

			err, ok := m.Error.(*Error)
			if !ok {
				buf = append(buf, m)
				continue
			}

			if layout.MaxPriority < err.MaxPriority || (layout.MaxPriority == err.MaxPriority && layout.Summary == "") {
				layout.Summary = fmt.Sprintf("%s, %s", name, err.Summary)
				layout.MaxPriority = err.MaxPriority
			}

			replaced := v.layout(err)
			for childName, childMessages := range replaced.Messages {
				fullname := fmt.Sprintf("%s.%s", name, childName)
				layout.Messages[fullname] = append(layout.Messages[fullname], childMessages...)
			}
		}
		if buf != nil {
			layout.Messages[name] = buf
		}
	}
	return layout
}
