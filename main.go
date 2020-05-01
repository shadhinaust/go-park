package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./executor"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		if cmd[0] == "exit" {
			return
		}
		executor.ExecCmd(cmd)
		fmt.Print(">")
	}
}
