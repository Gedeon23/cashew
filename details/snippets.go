package details

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/Gedeon23/cashew/entry"
)

// NEXT implement snippets detail view
func GetSnipptets(entry *entry.Recoll, term string) error {
	query := fmt.Sprintf("%s dir:\"%s\" filename:\"%s\"", term, filepath.Dir(entry.Url[7:]), entry.File)
	cmd := exec.Command("recoll", "-t", "-A", "-p 12", query)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error in recoll query %s for snippets:\n %s\n %s", cmd.String(), err, out)
	}

	scan := bufio.NewScanner(bytes.NewReader(out))
	lineNumber := 0
	for scan.Scan() {
		if lineNumber >= 5 {
			entry.Snippets = append(entry.Snippets, scan.Text())
		}
		lineNumber++
	}

	return nil
}
