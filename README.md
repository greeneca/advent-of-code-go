# Advent Of Code in Go

I am doing the Advent of code in go to learn go.

## Setup

Expects a `.session_token` file in the root directory with the session token from your advent of code account.

The program will automaticall download the input for the day when running a problem if it does nto already exist.

## Running
To run a problem, use the following command:

```bash
go run app.go <year> <day>-<part> (input file)
```

`input file` is optional. If not provided, the program will use the defauklt input from the advent of code website.

## Example

```bash
go run app.go 2017 1-1
```
This will run the first part of day 1 of 2017. The input will be downloaded if it does not already exist.

```bash
go run app.go 2017 1-2 1-1
```
This will run the second part of day 1 of 2017. It will try and use `inputs/2017/1-1.txt` as the input file. If it does not exist, it will fall back to the default input from the advent of code website.


