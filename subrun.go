package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func subrun() {
	fmt.Printf("SubRunnning %v as PID %d\n", os.Args[2:], os.Getpid())

	cgroupctrl()

	syscall.Sethostname([]byte("container"))
	syscall.Chroot("/home/" + user + "/" + fs)
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "") /// TODO: wiki
	os.Setenv("PATH", "/bin:/usr/bin:/sbin:/usr/sbin")

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

	must(syscall.Unmount("/proc", 0))
}
