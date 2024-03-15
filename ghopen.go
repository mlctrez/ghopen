package main

import (
	"log"
	"os/exec"
	"strings"
)

func main() {
	command := exec.Command("git", "config", "--get", "remote.origin.url")
	output, err := command.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	u := strings.ReplaceAll(string(output), "git@github.com:", "https://github.com/")
	err = exec.Command("open", u).Run()
	if err != nil {
		log.Fatal(err)
	}

}
