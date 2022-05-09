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

	// job manager metrics test
	config, err := c.JobManagerMetrics()
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
