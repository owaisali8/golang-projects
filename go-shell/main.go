package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func shell() {
	reader := bufio.NewReader(os.Stdin)
	for {

		currentDir, err := os.Getwd()
		host, err := os.Hostname()
		user, err := user.Current()

		fmt.Print("<", user.Name, "@", host, ">\n")
		fmt.Print(currentDir, "> ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println()

	}
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":

		if len(args) < 2 {
			return errors.New("Path Required")
		}

		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
		wg.Done()
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()

}

func main() {
	wg.Add(1)
	go shell()
	wg.Wait()
}
