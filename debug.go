package dddebug

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// DD dumps the given data and terminates the program.
func DD(data interface{}) {
	// Convert the data to a pretty-printed JSON string for readability
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	// Print the pretty-printed JSON string
	fmt.Println(string(bytes))

	// Exit the program
	os.Exit(1)
}
