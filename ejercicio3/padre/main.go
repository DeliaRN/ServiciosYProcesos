package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

func main() {

	cmd := exec.Command("go", "run", "../hijo/main.go")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error al obtener el pipe de salida estándar: %v", err)
	}

	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			fmt.Println("Padre recibió:", scanner.Text())
		}
	}()

	if err := cmd.Start(); err != nil {
		log.Fatalf("Error al iniciar el proceso hijo: %v", err)
	}

	comando1 := "1. Enviar la señal SIGUSR1 al proceso hijo."
	comando2 := "2. Enviar la señal SIGUSR2 al proceso hijo."
	comando3 := "3. Salir."

	for {
		fmt.Println(comando1)
		fmt.Println(comando2)
		fmt.Println(comando3)

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			cmd.Process.Signal(syscall.SIGUSR1)
		case 2:
			cmd.Process.Signal(syscall.SIGUSR2)
		case 3:
			err := cmd.Process.Signal(syscall.Signal(0))
			if err == nil {
				cmd.Process.Kill()
			}
			cmd.Wait()
			return

		default:
			fmt.Print("Opción no válida")
		}

	}

}
