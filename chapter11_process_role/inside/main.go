package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	path, _ := os.Executable()
	fmt.Printf("file name:  %s\n", os.Args[0])
	fmt.Printf("file path:  %s\n", path)
	fmt.Printf("process id: %d\n", os.Getpid())
	fmt.Printf("ppid:       %d\n", os.Getppid())

	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Fprintf(os.Stdout, "group ID: %d session ID: %d\n", syscall.Getpgrp(), sid)

	fmt.Printf("user ID:      %d\n", os.Getuid())
	fmt.Printf("group ID:     %d\n", os.Getgid())
	fmt.Printf("effective user ID:  %d\n", os.Geteuid())
	fmt.Printf("effective group ID: %d\n", os.Getegid())
	groups, _ := os.Getgroups()
	fmt.Printf("sub group ID: %v\n", groups)

	wd, _ := os.Getwd()
	fmt.Printf("working directory: %v\n", wd)
}
