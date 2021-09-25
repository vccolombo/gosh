package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
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

func execCommand(command string, args []string) *bytes.Buffer {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return &out
}

func displayOutput(out *bytes.Buffer) {
	fmt.Print(out.String())
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		displayPrompt()
		input := readInput(reader)
		command, args := parseInput(input)
		out := execCommand(command, args)
		displayOutput(out)
	}
}
