// fbi is a Go client for the FBI's Most Wanted REST API 🚔
//
// Installation
//
//		go get -u github.com/brittonhayes/fbi
//
//
// Usage
//
//	func main() {
//		// Initialize fugitives
//		f := new(fbi.Fugitives)
//
//		// List the fugitives as pretty-printed json
//		err := f.List()
//		if err != nil {
//		    // handle error
//			panic(err)
//		}
//
//		// Print raw results
//		fmt.Println(f)
//
//		// Print specific items from the list of results
//		fmt.Println(f.Items[0])
//
//	}
//
//
// Pretty Print Results
//
//	func main() {
//		// Initialize fugitives
//		f := new(fbi.Fugitives)
//
//		// Pretty print the results as JSON
//		// List the fugitives as pretty-printed json
//		j, err := f.ListPretty()
//		if err != nil {
//			panic(err)
//		}
//
//		// Print out the results
//		fmt.Println(string(j))
//
//	}
//
// Reference
//
// Source API: https://api.fbi.gov/wanted/v1/list
//
package fbi
