package main

import (
	"bytes"
	"encoding/base64"
	"github.com/Gedeon23/cashew/entry"
	"github.com/charmbracelet/bubbles/list"
	"log"
	"os/exec"
	"strings"
)

func Collect(term string) []list.Item {
	// Example command
	cmd := exec.Command("recoll", "-t", "-F", "author title file url", term)

	// Capture stdout and stderr
	var out, errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	// Run the command
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error: %v\n", err)
		log.Fatalf("Stderr: %s\n", errOut.String())
		return nil
	}

	data := strings.Split(out.String(), "\n")
	entries := make([]list.Item, 0, 20)

	for i := 2; i < len(data)-1 && i <= cap(entries); i++ {

		fields := strings.Split(data[i], " ")
		entry := entry.Recoll{}

		url, err := base64.StdEncoding.DecodeString(fields[3])
		if err != nil {
			log.Fatalf("err: decoding url '%s', unknown url, this should never happen (%v)\n", fields[3], err)
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

		entries = append(entries, entry)
	}

	log.Printf("returned entries: %s", entries)
	return entries
}
