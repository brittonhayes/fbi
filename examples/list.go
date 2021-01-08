package examples

import (
	"fmt"
	"github.com/brittonhayes/fbi"
)

func main() {
	// Initialize fugitives
	f := new(fbi.Fugitives)

	// List the fugitives as pretty-printed json
	err := f.List()
	if err != nil {
		// handle error
		panic(err)
	}

	// Print out the results or parse fields from the request
	fmt.Printf("Found %d results from request", f.Total)

	// Print raw results
	fmt.Println(f)

	// Print specific items from the list of results
	fmt.Println(f.Items[0])
}
