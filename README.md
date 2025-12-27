# YAWC (Yet Another WC)
A custom wc tool written in Golang.

## Usage
```
./main [options] file...
```

## Options
- `-l` line count
- `-w` word count  
- `-c` byte count
- `-m` character count

No flags defaults to lines, words, bytes.

## Build
```
go build -o main
```
