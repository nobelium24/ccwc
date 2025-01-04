# ccwc - A Custom Implementation of the `wc` Command in Go

## Overview

`ccwc` (Coding Challenges Word Count) is a custom implementation of the Unix `wc` command-line tool, written in Go. This program demonstrates foundational concepts of Unix Philosophy:

- Writing simple, modular tools with clean interfaces.
- Designing programs to connect seamlessly with other tools.

This project provided a great way to deepen my understanding of Go programming, file handling, and working with command-line interfaces.

## Features

`ccwc` supports the following functionalities:

1. **Byte count** (`-c`): Counts the number of bytes in a file.
2. **Line count** (`-l`): Counts the number of lines in a file.
3. **Word count** (`-w`): Counts the number of words in a file.
4. **Character count** (`-m`): Counts the number of characters in a file, including support for multibyte characters.
5. **Default behavior**: When no options are specified, outputs the line, word, and byte counts.
6. **Standard input support**: Reads from standard input if no filename is provided.

## Usage

### Build

To compile the program, ensure you have Go installed on your system, then run:

```bash
go build -o ccwc
```

This will generate an executable named `ccwc`.

### Run

You can use the `ccwc` command with the following options:

#### Count Bytes
```bash
./ccwc -c test.txt
```
Output:
```
  342190 test.txt
```

#### Count Lines
```bash
./ccwc -l test.txt
```
Output:
```
    7145 test.txt
```

#### Count Words
```bash
./ccwc -w test.txt
```
Output:
```
   58164 test.txt
```

#### Count Characters
```bash
./ccwc -m test.txt
```
Output (may vary based on locale):
```
  339292 test.txt
```

#### Default Behavior (Lines, Words, and Bytes)
```bash
./ccwc test.txt
```
Output:
```
    7145   58164  342190 test.txt
```

#### Read from Standard Input
```bash
cat test.txt | ./ccwc -l
```
Output:
```
    7145
```

## Testing

The program was tested against the GNU `wc` command using `test.txt`. Ensure that your locale settings match for consistent results when using the `-m` option.

## File Structure

```
.
├── main.go        # Main source file
├── countByte/
│   └── countbyte.go
├── countLines/
│   └── countLines.go
├── countWords/
│   └── countWords.go
├── characterCount/
│   └── characterCount.go
├── test.txt       # Sample input file for testing
├── README.md      # Documentation
```

## Concepts Covered

- File I/O in Go
- Command-line argument parsing
- Handling standard input
- Text processing
- Character encoding and locale considerations

## Acknowledgments

Inspired by the Unix Philosophy and `wc` utility. Special thanks to the [Coding Challenges Shared Solutions GitHub repo](https://github.com/CodingChallengesFYI) for fostering collaboration and learning.

## License

This project is licensed under the MIT License.

---

Feel free to fork and improve `ccwc` or share your solutions to help others learn!
