package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

// NEXT implement snippets detail view
func GetSnipptets(url string, term string, snippets *[]string) error {
	cmd := exec.Command("rga", term, url[7:])

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error running rga (make sure to install for snippets): %s", err)
	}

	scan := bufio.NewScanner(bytes.NewReader(out))
	for scan.Scan() {
		*snippets = append(*snippets, scan.Text())
	}

	return nil
}
