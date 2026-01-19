package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//wait for signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Signal(syscall.SIGUSR1), os.Signal(syscall.SIGUSR2))

	for {
		sig := <-sigChan

		switch sig {
		case syscall.SIGUSR1:
			_ = sig
		case syscall.SIGUSR2:
			pid := os.Getpid()
			syscall.Kill(pid, syscall.SIGTERM)
		default:
			fmt.Println("Señal desconocida")
		}
	}

}
