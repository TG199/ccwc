# ccwc `wc` Command Line Tool

A lightweight command line tool written in Go that mimics the Unix `wc` (word count) command. It allows users to count lines, words, and characters in a file.

## Features
- Count lines in a file using the `-l` flag.
- Count words in a file using the `-w` flag.
- Count characters in a file using the `-c` flag.
- Provide a summary (lines, words, characters) for multiple files.

## Prerequisites
- [Go](https://golang.org/doc/install) installed on your system.

## Usage
1. Clone this repository or download the source file.
2. Build the program using the provided `Makefile`.

### Example Commands
1. Build the program:
   ```bash
   make build

2. Count lines in a program:
    Count lines in a file:
    ./bin/wc-mini -l test.txt

    Count words in a file:
    ./bin/wc-mini -w test.txt
    
    Count characters in a file:
    ./bin/wc-mini -c test.txt

Get a summary of lines, words, and characters for multiple files:
./bin/wc-mini file1.txt file2.txt

Clean build artifacts:
make clean

Testing:
You can run the program against a sample file using:
make test
Contributing
Contributions are welcome! Feel free to submit a pull request or open an issue for bugs and features.

License
This project is open-source and available under the MIT License.

### How to Use These Files
1. Save the `Makefile` in the same directory as your `main.go` file.
2. Save the `README.md` in the root of your project directory.
3. Run `make build` to compile your program, and follow the instructions in the README for usage.
