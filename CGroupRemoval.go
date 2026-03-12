package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func rmcgrp() {
	cgProcs := filepath.Join(protoContainerPath, "cgroup.procs")

	if data, err := os.ReadFile(cgProcs); err == nil {
		if len(data) > 0 {
			fmt.Printf("В cgroup ещё есть процессы: %s\n", data)
			fmt.Println("Попробуй завершить процессы перед удалением.")
			return
		}
	} else if !os.IsNotExist(err) {
		fmt.Printf("Ошибка чтения cgroup.procs: %v\n", err)
		return
	}

	if err := os.RemoveAll(protoContainerPath); err != nil {
		fmt.Printf("Не удалось удалить cgroup: %v\n", err)
		return
	}

	fmt.Println("Cgroup directory removed successfully!")
}
