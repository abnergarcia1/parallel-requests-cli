package main

import (
	"flag"
	"log"
	"strconv"

	"abnergarcia1/parallel-requests-cli/cmd"
)

func main() {
	maxProcessesStr := flag.String("parallel", "10", "max parallel requests")
	flag.Parse()

	maxProcess, err := strconv.Atoi(*maxProcessesStr)
	if err != nil {
		log.Fatal("parallel arg was provided with invalid or null value")
	}
	urls := flag.Args()

	cmd.InitProcessor(maxProcess, urls)
}
