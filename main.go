package main

import (
	"fmt"
	"os"

	"github.com/s0ders/go-torrent-organizer/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
