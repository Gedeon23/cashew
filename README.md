# Cashew

**Cashew** (name not final) is a terminal user interface for [recoll](recoll.org).
The goal of this project is to offer a keyboard driven UI for recoll and integrate well into workflows centered around the terminal.


## Current Features

- Search recoll index
- Display of Title (file name if no title is found in metadata) and Author
- easy keyboard driven navigation
- Detail View for each entry
- easily open the entry in your preferred pdf viewer

## Planned Features

- easily open possible source files (latex, typst) in your editor 
- Surrounding Directory View
- flags to integrate into bash script


## Installation

Assuming you have Go installed on your system, the installation process is rather straight forward.
Just clone this repo:
```sh
git clone https://github.com/Gedeon23/cashew.git  
```

`cd` into the directory:
```
cd cashew
```

build the binary:
```
go build .
```

---

## Credits
I'm using [bubbletea](github.com/charmbracelet/bubbletea) for the TUI and [recoll](reoll.org) for indexing and querying PDFs
