package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

func main() {

	// Recuerda compilar el archivo hijo antes de ejecutar el padre
	// go build -o hijo ../hijo/hijo.go
	// Luego sólo -> go run main.go

	cmd := exec.Command("../hijo/hijo")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error creando Pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("Error al iniciar el proceso hijo: %v", err)
	}

	reader := bufio.NewReader(stdout)

	msg, _ := reader.ReadString('\n')
	fmt.Print(msg)

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
			msg, _ := reader.ReadString('\n')
			fmt.Print(msg)

		case 2:
			cmd.Process.Signal(syscall.SIGUSR2)
			msg, _ := reader.ReadString('\n')
			fmt.Print(msg)

			msg, _ = reader.ReadString('\n')
			fmt.Print(msg)

			cmd.Wait()
			return

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
