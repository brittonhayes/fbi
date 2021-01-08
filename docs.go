// fbi is a Go client for the FBI's Most Wanted REST API 🚔
//
// Installation
//
//	go get -u github.com/brittonhayes/fbi
//
//
// Usage
//
//	import (
//		"encoding/json"
//		"fmt"
//		"github.com/brittonhayes/fbi"
//	)
//
//	func main() {
//		f := new(fbi.Fugitives)
//		f.List()
//		j, err := json.MarshalIndent(&f, "", "\t")
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(string(j))
//	}
//
// Reference
//
// Source API: https://api.fbi.gov/wanted/v1/list
//
package fbi
