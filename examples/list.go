package examples

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/fbi"
)

func main() {
	f := new(fbi.Fugitives)
	f.List()
	j, err := json.MarshalIndent(&f, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}
