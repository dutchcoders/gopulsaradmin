package pulsaradmin

import (
	"fmt"

	pulsaradmin "github.com/dutchcoders/gopulsaradmin"
)

func ExampleExamples_output() {
	client, err := pulsaradmin.New(
		misp.WithURL("{url}"),
		misp.WithKey("{key}"),
	)
	if err != nil {
		panic(err.Error)
	}

	result, err := client.Search(ip)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Search results: %s\n", result)
}
