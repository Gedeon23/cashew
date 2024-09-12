package recoll

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)

func Collect(term string) ([]list.Item, error) {
	cmd := exec.Command("recoll", "-t", "-F", "author title file url", term)

	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	data := strings.Split(out.String(), "\n")
	entries := make([]list.Item, 0, 16)

	for i := 2; i < len(data)-1 && i <= cap(entries); i++ {

		fields := strings.Split(data[i], " ")
		entry := Entry{Query: term, Snippets: make([]Snippet, 0, 8)}

		url, err := base64.StdEncoding.DecodeString(fields[3])
		if err != nil {
			return nil, fmt.Errorf("Error: decoding url '%s', unknown url, this should never happen (%v)\n", fields[3], err)
		} else {
			entry.Url = string(url)
		}

		author, err := base64.StdEncoding.DecodeString(fields[0])
		if err != nil || string(author) == "" {
			entry.Author = "unknown"
		} else {
			entry.Author = string(author)
		}

		file, err := base64.StdEncoding.DecodeString(fields[2])
		if err != nil || string(file) == "" {
			split_url := strings.Split(entry.Url, "/")
			entry.File = split_url[len(split_url)-1]
		} else {
			entry.File = string(file)
		}

		title, err := base64.StdEncoding.DecodeString(fields[1])
		if err != nil || string(title) == "" {
			entry.DocTitle = entry.File
		} else {
			entry.DocTitle = string(title)
		}

		entry.Snippets = make([]Snippet, 0, 10)

		entries = append(entries, entry)
	}

	return entries, nil
}

func GetSnipptets(entry Entry, term string) (Entry, error) {
	query := fmt.Sprintf("%s dir:\"%s\" filename:\"%s\"", term, filepath.Dir(entry.Url[7:]), entry.File)
	cmd := exec.Command("recoll", "-t", "-A", "-p 12", query)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return entry, fmt.Errorf("Error in recoll query %s for snippets:\n %s\n %s", cmd.String(), err, out)
	}

	scan := bufio.NewScanner(bytes.NewReader(out))
	lineNumber := 0
	for scan.Scan() {
		if lineNumber >= 5 {
			splitScan := strings.SplitN(scan.Text(), ":", 2)
			if len(splitScan) == 2 {
				entry.Snippets = append(entry.Snippets, Snippet{Page: splitScan[0], Text: splitScan[1]})
			}
		}
		lineNumber++
	}

	return entry, nil
}
