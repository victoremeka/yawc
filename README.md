# YAWC (Yet Another WC)
A custom wc tool written in Golang.

## Install
Download from [Releases](https://github.com/victoremeka/yawc/releases) or build from source.

## Usage
```
yawc [options] file...
```

## Options
- `-l` line count
- `-w` word count  
- `-c` byte count
- `-m` character count

No flags defaults to lines, words, bytes.

## Build
```
go build -o yawc
```
