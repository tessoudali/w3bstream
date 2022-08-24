package main

import (
	"fmt"
	"math/big"

	"cuelang.org/go/cue"
)

func main() {
	const config = `
TimeSeries: {
  "2019-09-01T07:00:00Z": 36
}
TimeSeries: {
  "2019-09-01T07:10:59Z": 200
}
`
	var r cue.Runtime

	instance, err := r.Compile("test", config)
	if err != nil {
		fmt.Println(err)
	}

	var bigInt big.Int
	instance.Lookup("TimeSeries").Lookup("2019-09-01T07:10:59Z").Int(&bigInt)
	fmt.Println(bigInt.String())
}
