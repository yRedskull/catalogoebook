package main

import (
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	r, _err := ConfigServer()

	if _err != nil {
		panic(_err)
	}

	Routes(r)

	r.Run(":"+ port)
}
