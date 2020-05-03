package maperr_test

import (
	"fmt"

	"github.com/podhmo/maperr"
)

func ExampleError_Format() {
	var err *maperr.Error
	err = err.
		AddSummary("💣error is occured").
		Add("name", maperr.Message{Text: "name is empty"})

	fmt.Printf("%v\n", err)
	fmt.Printf("%+v\n", err)

	// Output:
	// Error -- "💣error is occured" (1 number of errors)
	// Error -- {
	//   "summary": "💣error is occured",
	//   "messages": {
	//     "name": [
	//       {
	//         "text": "name is empty"
	//       }
	//     ]
	//   }
	// }

}
