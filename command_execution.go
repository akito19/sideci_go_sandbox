package main

import "os/exec"

func f(string x) {
	cmd := "ls"
	exec.Command(cmd)
}
