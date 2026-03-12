package main

import (
	"fmt"
	"os"
)

// TODO: Cgroup ()
// TODO:убрать рут контейнеру

// docker         run <image> <cmd> <params>
// go run main.go run         <cmd> <params>

const (
	protoContainerPath = "/sys/fs/cgroup/container"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "subrun":
		if subrun() == 0 {
			rmcgrp()
		} else {
			fmt.Println("subrun не вернул 0...")
		}
	default:
		panic("Unknown command")
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
