package main

import "fmt"

//go:generate json2go -p main -n example -i example.json -o example.go

func main() {
	e := Example{}
	fmt.Printf("%+v\n", e)
}
