package main

import (
	setup "api/Http"
	"log"
)

func main() {

	app, appFire := setup.Setup()

	log.Fatal(app.Listen(":7334"))

	appFire.Close()
}
