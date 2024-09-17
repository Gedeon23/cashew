# Cashew

**Cashew** (name not final) is a terminal user interface for [recoll](https://recoll.org).
The goal of this project is to offer a keyboard driven UI for recoll and integrate well into workflows centered around the terminal.

![basic usage](./media/cashew.gif)


## Current Features

- Search recoll index
- easy keyboard driven navigation
- Detail View for each entry
- easily open the entry in your preferred pdf viewer
- Snippets View
- Open pdf at snippet location (currently only supporting zathura)

## Features Currently Working On

- flags to integrate into bash script

## Planned Features

- determine preferred document viewers
- easily open possible source files (latex, typst) in your editor 
- Surrounding Directory View


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


### Flags

`cashew -h`
```bash
Usage of cashew:
  -log
    	turn on error logging, usage: --log=true, file: /tmp/cashew_debug.log, default: false
  -search string
    	predefine search query, usage: --search='<query>'
```


---

## Credits
I'm using [bubbletea](https://github.com/charmbracelet/bubbletea) for the TUI and [recoll](https://recoll.org) for indexing and querying PDFs
