package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func displayPrompt() {
	fmt.Print("> ")
}

func readInput(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
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
	case "cd":
		err = execChangeDir(args)
	case "exit":
		os.Exit(0)
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
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		loop(reader)
	}
}
