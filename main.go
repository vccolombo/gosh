package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func displayPrompt() {
	path, _ := os.Getwd()
	fmt.Printf("%s > ", path)
}

func exitShell(code int) {
	os.Exit(code)
}

func readInput(reader *bufio.Reader) string {
	line, err := reader.ReadString('\n')
	if err == io.EOF {
		exitShell(0)
	}

	line = strings.TrimSuffix(line, "\n")

	return line
}

func parseInput(input string) (string, []string) {
	splitted := strings.Split(input, " ")
	command := splitted[0]
	args := splitted[1:]

	return command, args
}

func execChangeDir(args []string) error {
	if len(args) != 1 {
		return errors.New("cd error: no path provided")
	}
	return os.Chdir(args[0])
}

func execRealCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	return err
}

func execCommand(command string, args []string) error {
	var err error = nil
	switch command {
	case "":
		// do nothing
	case "cd":
		err = execChangeDir(args)
	case "exit":
		exitShell(0)
	default:
		err = execRealCommand(command, args)
	}

	return err
}

func loop(reader *bufio.Reader) {
	displayPrompt()
	input := readInput(reader)
	command, args := parseInput(input)
	err := execCommand(command, args)
	if err != nil {
		// exitError, _ := err.(*exec.ExitError)
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func setupSignals() {
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGINT,
		syscall.SIGTSTP)

	go func() {
		for {
			s := <-signalChanel
			switch s {
			case syscall.SIGINT: // Ctrl + c
				fmt.Println()
				displayPrompt()
			case syscall.SIGTSTP: // Ctrl + z
				fmt.Println()
				displayPrompt()
			}
		}
	}()
}

func main() {
	setupSignals()

	reader := bufio.NewReader(os.Stdin)
	for {
		loop(reader)
	}
}
