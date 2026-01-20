package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Recuerda compilar el archivo hijo antes de ejecutar el padre
	// go build -o hijo ../hijo/hijo.go

	fmt.Println("Hijo: proceso iniciado")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGUSR1, syscall.SIGUSR2)

	for {
		sig := <-sigChan

		switch sig {
		case syscall.SIGUSR1:
			fmt.Println("Hijo: señal SIGUSR1 recibida")

		case syscall.SIGUSR2:
			fmt.Println("Hijo: señal SIGUSR2 recibida")
			fmt.Println("Hijo: proceso terminado")
			os.Exit(0)

		default:
			fmt.Println("Señal desconocida")
		}
	}

}
