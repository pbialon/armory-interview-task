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

The solution to the problem is implemented as a command-line tool called `logs_printer`. 
The tool reads log files from a specified directory, sorts the lines based on timestamps, and prints them to the screen.

### Implementation Details

The tool uses a priority queue to sort the lines from the log files.
First, it reads the first lines from the set of given files and creates a priority queue. 
A "priority" is the timestamp of a given log line. Lower timestamps have higher priority.
Then, each time we pop a line from the priority queue, and we read the next line from the file that contained the popped line.
We fix the heap, and we repeat the process until all lines from all files are read and printed.
The solution works in _O(n * log(k) )_ time, where _n_ is the total number of lines in all files, and _k_ is the number of files.

### What is not covered in the solution

During the implementation I made the following assumptions:
1. Number of log files is less than the limit of the open file descriptors on the machine.  
2. The combined length of any set of lines from the log files (1 line from each file) is less than 16,000,000,000 (it is feasible to store them in memory).
3. All the logs files are located in a single directory on local disk, and they have the .log extension.
4. If there is any error in parsing a log line, the line is ignored, and we log the error to the console.

Apart from the first assumption, the other assumptions are not critical for the solution. 
To handle a case when the number of log files is greater than the limit of the open file descriptors,
we can modify the solution to read the batches of log files, and create an output file for each batch.
We can do it until we are left with fewer files to process than the limit of the open file descriptors.

I didn't implement this, as I didn't have enough time to do it, but it can be done if needed.

### Repository Structure
The repository structure consists of the following components:

- `src`: Contains the source code for the solution. It contains of two packages:
  - `log_files`: Contains the code for reading and sorting the log files.
  - `priority_queue`: Contains the code for the priority queue used for on-the-fly files sorting.
- `resources`: Contains the log files used for testing the solution.


## Usage

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

