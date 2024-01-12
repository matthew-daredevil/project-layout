package main

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func main() {

	cmd := exec.Command("sh", "hubble", "observe")
	// cmd.Dir = "/test/workspace"

	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(output))
	}
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}