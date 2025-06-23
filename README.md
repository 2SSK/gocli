# gocli

A simple CLI application written in Go that allows you to execute and install scripts using a command-line interface. This project uses [Cobra](https://github.com/spf13/cobra) for building the CLI.

## Features

- Run custom scripts via subcommands
- Install dependencies or components (e.g., Neofetch) using subcommands
- Easily extensible with new commands and scripts

## Project Structure

```
.
├── cmd/                # Go source files for CLI commands
│   ├── hello.go        # Runs the hello script
│   ├── install.go      # Handles install subcommands
│   ├── neofetch.go     # Installs neofetch via script
│   ├── root.go         # Root command definition
│   └── run.go          # Handles run subcommands
├── scripts/            # Shell scripts executed by the CLI
│   ├── hello.sh
│   └── neofetch.sh
├── main.go             # Entry point for the CLI
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/gocli.git
   cd gocli
   ```

2. **Build the CLI:**
   ```sh
   go build -o gocli
   ```

## Usage

### Run the CLI

```sh
./gocli [command]
```

### Available Commands

- `run`  
  Run different scripts using subcommands.

    - `run hello`  
      Executes the `hello.sh` script.  
      You can pass an argument with `-a` or `--arg`:
      ```sh
      ./gocli run hello -a "your_argument"
      ```

- `install`  
  Install components or dependencies using subcommands.

    - `install neofetch`  
      Installs [Neofetch](https://github.com/dylanaraps/neofetch) using the appropriate package manager via `scripts/neofetch.sh`:
      ```sh
      ./gocli install neofetch
      ```

### Example

```sh
./gocli run hello -a "World"
./gocli install neofetch
```

## Adding New Scripts

1. Add your script to the `scripts/` directory.
2. Create a new command file in `cmd/` (see [`cmd/hello.go`](cmd/hello.go) or [`cmd/neofetch.go`](cmd/neofetch.go) for examples).
3. Register your command in the appropriate parent command in its `init()` function.

## Requirements

- Go 1.18+
- Bash (for running scripts)
- Supported package manager for install scripts

## License

See