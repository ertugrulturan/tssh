package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	defaultUser := "root"
	defaultPort := "22"

	args := os.Args[1:]

	if len(args) > 0 {
		user := defaultUser
		hostIP := args[0]
		port := defaultPort

		if len(args) > 1 {
			user = args[1]
		}

		if len(args) > 2 {
			port = args[2]
		}

		sshCommand := fmt.Sprintf("ssh -p %s %s@%s", port, user, hostIP)

		for {
			if ping(hostIP) {
				fmt.Println("Ping is successful, SSH connection is being made...")
				time.Sleep(2 * time.Second)

				cmd := exec.Command("bash", "-c", sshCommand)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Stdin = os.Stdin
				cmd.Run()
				break
			} else {
				fmt.Println("Ping waiting...")
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func ping(hostIP string) bool {
	cmd := exec.Command("ping", "-c", "1", hostIP)
	err := cmd.Run()
	return err == nil
}
