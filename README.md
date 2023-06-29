# Logs Printer

This repository contains a solution to the task of printing individual lines from log files, sorted by timestamp. The solution is implemented in Go.

## Problem Description

The problem assumes that there are multiple servers generating log files for a distributed application. Each log file can vary in size from _100MB_ to _512GB_, and they are copied to a local machine with limited resources. The local directory structure consists of log files from different servers.

Example log file paths:
- `/temp/server-ac329xbv.log`
- `/temp/server-buyew12x.log`
- `/temp/server-cnw293z2.log`

The goal is to print the individual lines from the log files, sorted by timestamp.

## Solution

The solution to the problem is implemented as a command-line tool called `logs_printer`. The tool reads log files from a specified directory, sorts the lines based on timestamps, and prints them to the screen.

### Usage

To use the `logs_printer` tool, follow these steps:

1. Clone this repository to your local machine.
2. Navigate to the repository directory.
3. Ensure that you have Go installed.
4. Open a terminal or command prompt in the repository directory.
5. Run the following command to compile the source code and create the executable:
```bash
make compile
./build/logs_printer <path_to_logs>
```
Replace `<path_to_logs>` with the path to the directory containing the log files. The tool assumes that all log files in the specified directory have the .log extension.

### Makefile Commands

The repository includes a `Makefile` that provides convenient commands for building and testing the solution.

- `make compile`: Compiles the source code and creates the executable inside the `build` directory.
- `make test`: Runs the tests for the solution.

