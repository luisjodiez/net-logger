# net-logger

## Table of Contents
- [Net Logger CLI](#net-logger-cli)
- [Usage](#usage)
- [Output](#output)
- [Requirements](#requirements)
- [Build Instructions](#build-instructions)
  - [Prerequisites](#prerequisites)
  - [Build Steps](#build-steps)
- [Runtime Requirements](#runtime-requirements)
- [Build Dependencies](#build-dependencies)

# Net Logger CLI

A simple Go CLI tool to concurrently probe hosts (ICMP or TCP) and log results with timestamps. Generates a Markdown report on exit.

## Usage

```
go run ./cmd/netlogger --for 10s 8.8.8.8 example.com:80
```

- `--for` duration: Run for a specific time (e.g., 10s, 1m)
- `--until` datetime: Run until a specific RFC3339 datetime
- Neither `--for` or `--until`: until program is stopped
- Arguments: List of targets (host for ICMP, host:port for TCP)

## Output

A `report.md` file will be generated with a Markdown table of all probe results.

## Build Instructions

To manually build the `net-logger` application, follow these steps:

### Prerequisites
- **Go (1.21+)**: Ensure Go is installed and available in your PATH.
- **Git**: Required to fetch dependencies.

### Build Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/luisjodiez/net-logger.git
   cd net-logger
   ```
2. Build the application:
   ```bash
   go build -o netlogger ./cmd/netlogger
   ```
3. The binary `netlogger` will be created in the current directory.

## Runtime Requirements
- **ICMP Support**: Ensure the application has the necessary permissions to send ICMP packets (e.g., run with elevated privileges if required).
- **Network Access**: The application requires access to the internet or the specified targets.

## Build Dependencies
- **Go Modules**: The project uses Go modules to manage dependencies. Ensure you have internet access to fetch them during the build process.
