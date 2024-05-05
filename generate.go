package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

func generateUuids(gen int, doExit bool) {

	if gen <= 0 {
		handleAndExit("gen should be greater than 0")
	}

	for i := gen; i > 0; i-- {
		uuid, err := uuid.NewRandom()
		if err != nil {
			handleAndExit("failed to generate the requested number of UUIDs:", err)
		}
		fmt.Fprintln(os.Stdout, uuid)
	}

	checkAndExit(doExit)
}
