package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func run() {
	fmt.Printf("Runnning %v PID %d\n", os.Args[2:], os.Getpid())
	if runtime.GOOS != "linux" {
		fmt.Println("Containers only supported on Linux")
		os.Exit(1)
	}

	if os.Geteuid() != 0 {
		fmt.Println("Need root privileges (use sudo)")
		os.Exit(1)
	}

	cmd := exec.Command("/proc/self/exe", append([]string{"subrun"}, os.Args[2:]...)...) // TODO: wiki
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{ // TODO: wiki
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS, // TODO: wiki
		Unshareflags: syscall.CLONE_NEWNS,
	}
	cmd.Run()
	rmcgroup()
}
