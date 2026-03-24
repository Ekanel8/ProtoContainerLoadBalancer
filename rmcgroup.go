package main

import (
	"fmt"
	"os/exec"
)

func rmcgroup() {
	tool := "rmdir"
	path := protoContainerPath
	cmd := exec.Command(tool, path)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("[Cgroup deleted successfully]")
}
