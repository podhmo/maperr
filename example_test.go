package maperr_test

import (
	"fmt"

	"github.com/podhmo/maperr"
)

func ExampleFormat() {
	var err *maperr.Error
	err = err.AddSummary("💣error is occured")

	fmt.Printf("%v", err)
	// Output:
	// Error -- "💣error is occured" (1 number of errors)
}

func ExampleFormatVerbose() {
	var err *maperr.Error
	err = err.
		AddSummary("💣error is occured").
		Add("name", maperr.Message{Text: "name is empty"})

	fmt.Printf("%+v", err)
	// Output:
	// Error -- {
	//   "summary": "💣error is occured",
	//   "fields": {
	//     "name": [
	//       {
	//         "text": "name is empty"
	//       }
	//     ]
	//   }
	// }
}
