package main

import (
	"fmt"
	"os"

	"github.com/thinry/api"
)

func main() {
	c, err := api.New(os.Getenv("FLINK_API"))
	if err != nil {
		panic(err)
	}

	// jobs test
	jobs, err := c.Jobs()
	if err != nil {
		panic(err)
	}
	fmt.Println(jobs)
}
