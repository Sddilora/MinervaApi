package main

import (
	setup "api/Http"
	"log"
)

func main() {

	app, appFire := setup.Setup()

	defer appFire.Close()

	log.Fatal(app.Listen(":7334"))

}
