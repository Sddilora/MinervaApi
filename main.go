package main

import setup "api/Setup"

func main() {

	start := setup.Setup() // start is err

	if start != nil {
		panic(start)
	}
}
