package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getPID(line string) *int {
	split := strings.Split(strings.TrimSpace(line), " ")

	if len(split) > 0 {
		num, err := strconv.Atoi(split[0])
		if err != nil {
			return nil
		}
		return &num
	}

	return nil
}

func killProcesses(pids []string) {
	r := append([]string{"kill"}, pids...)
	cmd := exec.Command("sudo", r...)
	// fmt.Println(cmd.String())
	cmd.Run()
}

func contains(s string, substrings []string) bool {
	for _, ss := range substrings {
		if strings.Contains(s, ss) && !strings.Contains(s, "kall") {
			return true
		}
	}
	return false
}

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Printf("Killing processes containing: ")

	for _, arg := range argsWithoutProg {
		fmt.Printf("%s ", arg)
	}

	fmt.Println()

	cmd := exec.Command("ps", "-ax")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(out.String(), "\n")

	// fmt.Printf("Line count: %v\n", len(lines))

	pids := []string{}

	for _, line := range lines {
		if contains(line, argsWithoutProg) {
			pid := getPID(line)
			fmt.Println(line)
			if pid != nil {
				pids = append(pids, fmt.Sprintf("%v", *pid))
			}
		}
	}

	// fmt.Println(pids)

	fmt.Println("Total Processes Matching:", len(pids))
	killProcesses(pids)
}
