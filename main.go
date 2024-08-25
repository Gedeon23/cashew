package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Example command
	cmd := exec.Command("recoll", "-t", "-F", "author title file url", "Wiederanlauf")

	// Capture stdout and stderr
	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Printf("Stderr: %s\n", errOut.String())
		return
	}

	// Print stdout

	data := strings.Split(out.String(), " ")

	for i, entry := range data {
		entry, _ := base64.StdEncoding.DecodeString(entry)
		data[i] = string(entry)
	}

	fmt.Printf("Output:\n%s\n", data)
}
