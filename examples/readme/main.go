package main

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/podhmo/maperr"
)

type Person struct {
	Name string `json:"name"` // required
	Age  int    `json:"age"`
}

func (p *Person) UnmarshalJSON(b []byte) error {
	var err *maperr.Error
	var inner struct {
		Name *string `json:"name"`
		Age  *int    `json:"age"`
	}

	if rawerr := json.Unmarshal(b, &inner); rawerr != nil {
		err = err.AddSummary(rawerr.Error())
	}

	if inner.Name != nil {
		p.Name = *inner.Name
	} else {
		err = err.Add("name", maperr.Message{Text: "required"})
	}
	if inner.Age != nil {
		p.Age = *inner.Age
	}

	return err.Untyped()
}

func main() {
	b := bytes.NewBufferString(`{"age": 20}`)
	decoder := json.NewDecoder(b)

	var p Person

	if err := decoder.Decode(&p); err != nil {
		log.Fatalf("%+v", err)
	}
}
