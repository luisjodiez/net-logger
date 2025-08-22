# net-logger
Collect data on network metrics

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

## Requirements

- Go 1.21+
- `ping` command available in PATH (for ICMP)
