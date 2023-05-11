package main

import "lab-loan/route"

func main() {
	e := route.New()
	e.Start(":8080")
}

