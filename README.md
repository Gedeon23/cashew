# Cashew

**Cashew** (name not final) is a terminal user interface for [recoll](https://recoll.org).
The goal of this project is to offer a keyboard driven UI for recoll and integrate well into workflows centered around the terminal.

![basic usage](./media/cashew.gif)


## Current Features

- Search recoll index
- Display of Title (file name if no title is found in metadata) and Author
- easy keyboard driven navigation
- Detail View for each entry
- easily open the entry in your preferred pdf viewer
- Snippets View
- Open pdf at snippet location (currently only supporting zathura)

## Planned Features

- easily open possible source files (latex, typst) in your editor 
- Surrounding Directory View
- flags to integrate into bash script


## Installation

Assuming you have Go installed on your system, the installation process is as follows.
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

### Needed Applications

Right now the only requirement is recoll itself. I might add features in the future which are not accessible within the recoll cli.

---

## Credits
I'm using [bubbletea](https://github.com/charmbracelet/bubbletea) for the TUI and [recoll](https://recoll.org) for indexing and querying PDFs
