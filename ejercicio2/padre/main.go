package main

import (
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

func main() {

	cmd := exec.Command("go", "run", "../hijo/main.go")
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Error al iniciar el proceso hijo: %v", err)
	}
	//pidHijo := cmd.Process.Pid

	comando1 := "1. Enviar la señal SIGUSR1 al proceso hijo."
	comando2 := "2. Enviar la señal SIGUSR2 al proceso hijo."
	comando3 := "3. Salir."

	for {
		fmt.Println(comando1)
		fmt.Println(comando2)
		fmt.Println(comando3)

		var choice int
		fmt.Scan(&choice)

		/*
			if _, err := fmt.Scan(&choice); err != nil {
				fmt.Println("Error al leer: ", err)
			}
		*/

		switch choice {
		case 1:
			//syscall.Kill(pidHijo, syscall.SIGUSR1)
			cmd.Process.Signal(syscall.SIGUSR1)
		case 2:
			//syscall.Kill(pidHijo, syscall.SIGUSR2)
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
