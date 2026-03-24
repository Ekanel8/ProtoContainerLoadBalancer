package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func cgroupctrl() { // TODO: %CPU %MEM
	err := os.MkdirAll(protoContainerPath, 0700)
	if err != nil {
		panic("Failed to create cgroup dir")
	}
	pidStr := strconv.Itoa(int(os.Getpid()))
	limiterPath := filepath.Join(protoContainerPath, "cgroup.procs")

	must(os.WriteFile(limiterPath, []byte(pidStr), 0644))

	pidsMaxPath := filepath.Join(protoContainerPath, "pids.max")

	must(os.WriteFile(pidsMaxPath, []byte("20\n"), 0644))

	fmt.Printf("Set PID %d to custom cgroup: %s\n", os.Getpid(), protoContainerPath)
}
