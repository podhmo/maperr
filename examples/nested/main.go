package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/podhmo/maperr"
)

type Person struct {
	Name string `json:"name"`
	Info *Info  `json:"info"`
}

func (p *Person) UnmarshalJSON(b []byte) error {
	var err *maperr.Error

	// loading internal data
	var inner struct {
		Name *string          `json:"name"` // required
		Info *json.RawMessage `json:"info"`
	}
	if rawErr := json.Unmarshal(b, &inner); rawErr != nil {
		return err.AddSummary(rawErr.Error())
	}

	// binding field value and required check
	if inner.Name != nil {
		p.Name = *inner.Name
	} else {
		err = err.Add("name", maperr.Message{Text: "required"})
	}
	if inner.Info != nil {
		p.Info = &Info{}

		if rawerr := json.Unmarshal(*inner.Info, p.Info); rawerr != nil {
			err = err.Add("info", maperr.Message{Error: rawerr})
		}
	}
	return err.Untyped()
}

type Info struct {
	Name string `json:"name"`
}

func (i *Info) UnmarshalJSON(b []byte) error {
	var err *maperr.Error

	// loading internal data
	var inner struct {
		Name *string `json:"name"` // required
	}
	if rawErr := json.Unmarshal(b, &inner); rawErr != nil {
		return err.AddSummary(rawErr.Error())
	}

	// binding field value and required check
	if inner.Name != nil {
		i.Name = *inner.Name
	} else {
		err = err.Add("name", maperr.Message{Text: "required"})
	}
	return err.Untyped()
}

func main() {
	b := bytes.NewBufferString(`{"name": "foo", "info": {}}`)
	decoder := json.NewDecoder(b)

	var ob Person
	if err := decoder.Decode(&ob); err != nil {
		log.Fatalf("!! %+v", err)
	}
	fmt.Printf("%+v", ob)
}
