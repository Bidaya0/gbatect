package main

import (
	"github.com/bidaya0/gbatect/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
