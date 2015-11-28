# gencontext

gencontext is a golang tool that allows the creation of context storage and recall

## Install

	go install github.com/sheenobu/gencontext/cmd/...

## Usage

mydata/mydata.go:
	package mydata

	//go:generate gencontext -type int .

main.go:
	package main

	import (
		"golang.org/x/net/context"
		"mydata"
	)

	func main() {

		ctx := context.Background()

		ctx = mydata.NewContext(ctx, 1)

		// get data
		val, ok := mydata.FromContext(ctx)

		// get data with default
		val2 := mydata.FromContextDefault(ctx, 3)

		// get data and panic on missing
		val3 := mydata.Must(mydata.FromContext(ctx))
	}


## License

Code is based on stringer, therefore has the go BSD-style license

## TODO:

 * Import detection
 * Custom prefixes


