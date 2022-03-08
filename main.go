package main

import (
	"fmt"
	"github.com/Ladence/aerospike-tweet/internal/storage"
	aero "github.com/aerospike/aerospike-client-go"
)

// Please handle errors properly.
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	aeroClient := storage.NewAero()
	defer aeroClient.Close()

	key, err := aero.NewKey("test", "aerospike", "key")
	panicOnError(err)

	// define some bins with data
	bins := aero.BinMap{
		"bin1": 42,
		"bin2": "An elephant is a mouse with an operating system",
		"bin3": []interface{}{"Go", 2009},
	}

	// write the bins
	err = aeroClient.Put(nil, key, bins)
	panicOnError(err)

	// read it back!
	rec, err := aeroClient.Get(nil, key)
	panicOnError(err)

	fmt.Printf("Rec: %+v", *rec)
}
