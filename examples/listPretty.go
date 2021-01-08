package examples

import (
	"fmt"
	"github.com/brittonhayes/fbi"
)

func main() {
	// Initialize fugitives
	f := new(fbi.Fugitives)

	// List the fugitives as pretty-printed json
	j, err := f.ListPretty()
	if err != nil {
		// handle error
		panic(err)
	}

	// Print out the results
	fmt.Println(string(j))
}
