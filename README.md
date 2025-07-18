# gocli

A simple, cross-platform CLI tool written in Go for running custom scripts and installing packages from the terminal. Built with the Cobra CLI framework, `gocli` is designed to be easily extensible and user-friendly.

## Features

- Run custom shell scripts with subcommands
- Install tools like Neofetch via CLI
- Easily extensible for new scripts and commands
- Cross-platform support

## Installation

To install `gocli` globally, run:

```bash
bash <(curl -fsSL https://raw.githubusercontent.com/2SSK/gocli/main/install.sh)
```

Or, build from source:

```bash
git clone https://github.com/2SSK/gocli.git
cd gocli
go build -o gocli
```

## Usage

### Run a script

```bash
gocli run hello -a "World"
```

### Install Neofetch

```bash
gocli install neofetch
```

### Show help

```bash
gocli --help
```

## Project Structure

```
.
├── cmd/         # CLI command implementations
├── scripts/     # Shell scripts
├── install.sh   # Installer script
├── main.go      # Entry point
```

## Requirements

- Go 1.18+
- Bash
- Internet connection (for some scripts)

## Contributing

Contributions are welcome! Please open issues or submit pull requests for new features, bug fixes, or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
