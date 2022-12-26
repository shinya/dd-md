# dd-md

## Installation

### Install

```bash
go install github.com/shinya/dd-md@latest
```

### Uninstall

* Confirm(dry run)

```bash
go clean -i -n github.com/shinya/dd-md
```

* Execution

```bash
go clean -i github.com/shinya/dd-md
```

## Usage

```
Usage:
  dd-md [OPTIONS]

Application Options:
  -i  Import template read file (If not specified, read the file name template.md)
  -c  Configuration read file  (If not specified, read the file name setting.ini)
  -o  Output file name (If not specified, output to standard output)
```
