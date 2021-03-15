package main

import (
	"fmt"
	"os"

	"github.com/aleale2121/Golang-TODO-Hex-DDD/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
