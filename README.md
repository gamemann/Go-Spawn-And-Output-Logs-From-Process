# Go Spawn & Output Logs From Process
This is a small project to demonstrate how to spawn multiple processes from a Go program and have the Go program write their `stdout` and `stderr` lines to a log file.

In this example, we have `test` and `loader` Go programs.

The `test` program outputs a random string from the `messages` variable each second to `stdout`. This is just used as a demo program.

The `loader` program runs the `test` program five times and outputs their `stdout` and `stderr` pipes to a log file in the `logs/` directory (e.g. `logs/<pid>.log`).

## Motives
I'm working on a private project which spawns processes from a Go program. However, I wanted to write the `stdout` and `stderr` pipes from these spawned processes to a file so I knew what was going on. Since the project utilized Docker which extended build/test time, I decided to write a separate open source program to achieve this goal since I could easily test things.

## Building
You may use `make` via Makefile to build everything easily. Otherwise, you may use `go build loader.go` and `go build test.go` to build the Go programs.

## Running
Simply run the `loader` executable to test.

```bash
./loader
```

**Note** - The `remove_logs.sh` file is ran on each loader start. This Bash script simply removes all log files in the `logs/` directory so we start off from a clean slate and the loader ignores any errors from executing the file.

## Credits
* [Christian Deacon](https://github.com/gamemann)