package main

import "github.com/dsk52/vote/route"

func main() {
	router := route.Init()

	router.Logger.Fatal(router.Start(":1323"))
}
