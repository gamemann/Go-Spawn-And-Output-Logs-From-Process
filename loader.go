package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

type TestProcess struct {
	Process *exec.Cmd
	Stdout  *bufio.Writer
	Stderr  *bufio.Writer
}

const RMLOGS_EXEC = "./remove_logs.sh"
const EXEC = "./test"
const NUM_PROCESSES = 5

func main() {
	// Remove logs by using shell script.
	rmLogs := exec.Command(RMLOGS_EXEC)

	err := rmLogs.Run()

	if err != nil {
		fmt.Println("Failed to remove logs via shell script.")
		fmt.Println(err)
	}

	var processes []*TestProcess

	for i := 0; i < NUM_PROCESSES; i++ {
		process := exec.Command(EXEC)

		stdoutPipe, _ := process.StdoutPipe()
		stderrPipe, _ := process.StderrPipe()

		err := process.Start()

		if err != nil {
			fmt.Printf("[%d] Error creating process.\n", i)
			fmt.Println(err)

			continue
		}

		if process.Process == nil {
			fmt.Printf("[%d] Could not find process.\n", process.Process.Pid)

			continue
		}

		// Create log file.
		fileName := fmt.Sprintf("logs/%d.log", process.Process.Pid)
		logFile, err := os.Create(fileName)

		if err != nil {
			fmt.Printf("[%d] Error creating log file :: %s.\n", process.Process.Pid, fileName)

			continue
		}

		stdoutWriter := bufio.NewWriter(logFile)
		stderrWriter := bufio.NewWriter(logFile)

		// Create goroutines to capture stdout and stderr output and write them our log files.
		go func() {
			scanner := bufio.NewScanner(stdoutPipe)
			for scanner.Scan() {
				line := scanner.Text()
				stdoutWriter.WriteString(line + "\n")
				stdoutWriter.Flush()
			}
		}()

		go func() {
			scanner := bufio.NewScanner(stderrPipe)
			for scanner.Scan() {
				line := scanner.Text()
				stderrWriter.WriteString(line + "\n")
				stderrWriter.Flush()
			}
		}()

		processes = append(processes, &TestProcess{Process: process, Stdout: stdoutWriter, Stderr: stderrWriter})
	}

	fmt.Println("Started loader. Check logs/ directory for outputs.")

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	<-sigc

	os.Exit(0)
}
