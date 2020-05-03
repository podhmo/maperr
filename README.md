# maperr

- https://godoc.org/github.com/podhmo/maperr

## install

```
go get github.com/podhmo/maperr
```

## how to use

```go
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
```

result (verbose output)

```
2020/05/03 17:16:34 Error -- {
  "summary": "name, required",
  "messages": {
    "name": [
      {
        "text": "required"
      }
    ]
  }
}
exit status 1
```

## short version

If your code is something like following

- `panic(err)`
- `log.Fatalf("%v", err)`

Output is here.

```
Error -- "name, required" (1 number of errors)
```
