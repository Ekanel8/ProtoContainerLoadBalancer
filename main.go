package main

import (
	"os"
)

// TODO:убрать рут контейнеру

// docker         run <image> <cmd> <params>
// go run main.go run         <cmd> <params>

const ( // TODO: <env>
	protoContainerPath = "/sys/fs/cgroup/container"
	fs                 = "Debian-fs"
	user               = "doc"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "subrun":
		subrun()
	default:
		panic("Unknown command")
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
