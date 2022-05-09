package main

import (
	"os"

	"github.com/thinry/api"
)

func main() {
	c, err := api.New(os.Getenv("FLINK_API"))
	if err != nil {
		panic(err)
	}

	// shutdown test
	if err := c.Shutdown(); err != nil {
		panic(err)
	}
}
