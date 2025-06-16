# Paper-Transmission

# Origami CLI

This Go tool encrypts text messages into SVG origami crease patterns that can be printed.

## Installation

1. Make sure you have Go 1.18+ installed.
2. Go to the `cmd` folder and build the binary:
   ```bash
   cd cmd
   go build -o origami origami.go
   ```

## Usage

```bash
./origami -m "secret message" -o pattern.svg
```

- `-m` — message to encrypt
- `-o` — SVG output file name (default: `crease_pattern.svg`)

## Example

```bash
./origami -m "hello world" -o hello.svg
```

This will create a `hello.svg` file with a visual crease pattern.

## How it works

Each character of the message is encoded as a separate line radiating from the center of the sheet at a unique angle. This is a simple visualization that can be made more complex for stronger encryption.