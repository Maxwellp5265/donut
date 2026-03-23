# Donut

A terminal [donut](https://www.a1k0n.net/2011/07/20/donut-math.html) animation and ASCII art CLI written in Go, built with Charm's Bubble Tea and Lip Gloss TUI libraries.

## Demo

![Donut Demo](demo.gif)

## Note on GPU-Accelerated Terminals

For the best animation experience, use a GPU-accelerated terminal emulator such as [Alacritty](https://github.com/alacritty/alacritty), [Kitty](https://github.com/kovidgoyal/kitty), or similar. While not required, these terminals can provide smoother and more vibrant visuals for the donut animation.

## Table of Contents

- [Requirements](#requirements)
- [Features](#features)
- [Usage](#usage)
- [Dependencies](#dependencies)
- [Project Structure](#project-structure)
- [How it Works](#how-it-works)
- [Customization](#customization)
- [License](#license)

## Requirements

- [Go 1.26+](https://go.dev/dl/)

## Features

- Real-time optimized 3D simulation based on sloane's rotating [donut](https://www.a1k0n.net/2011/07/20/donut-math.html)
- Colorful, animated terminal output for modern terminal emulators
- 🚀 Emoji render mode
- Hardwired 60FPS ASCII rendering on most computers
- 300μs allocation free core loop (Apple Silicon M1)

## Usage

### Clone

```sh
git clone https://github.com/erik-adelbert/donut.git
cd donut
```

### Run with Makefile

```sh
make run
```

Build the executable:

```sh
make build
./bin/donut
```

### Install with go install

```sh
go install github.com/erik-adelbert/donut/cmd/donut@latest
```

### Test and benchmark

```sh
make test
make bench
```

### Run without Makefile

```sh
go run ./cmd/donut
```

Build a binary:

```sh
mkdir -p bin
go build -o bin/donut ./cmd/donut
./bin/donut
```

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- [Go terminal support](https://pkg.go.dev/golang.org/x/term)

## Project Structure

- `cmd/donut/` — CLI entry point (`main` package)
- `donut/` — Core simulation and rendering logic

## How it Works

- The model rotate and project a donut onto the screen. This work is a careful port of sloane's [2006 IOCCC entry](https://www.ioccc.org/2006/sloane/index.html),  [donut.c](https://www.a1k0n.net/2011/07/20/donut-math.html)

## Customization

- Adjust the palette in `donut/palette.go` to change the donut appearance.
- Change the simulation parameters as needed.

## License

MIT. See [LICENSE](LICENSE.TXT).

## Author

Erik Adelbert

Note: I don't need to vibe my code. This project is crafted.
