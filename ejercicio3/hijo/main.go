package main

import (
	"bufio"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//wait for signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGUSR1, syscall.SIGUSR2)

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Hijo: proceso iniciado\n")
	writer.Flush()

	for {
		sig := <-sigChan

		switch sig {
		case syscall.SIGUSR1:
			writer.WriteString("Hijo: señal SIGUSR1 recibida\n")
			writer.Flush()
		case syscall.SIGUSR2:
			writer.WriteString("Hijo: señal SIGUSR2 recibida\n")
			writer.WriteString("Hijo: proceso terminado\n")
			writer.Flush()
			os.Exit(0)
		default:
			writer.WriteString("Señal desconocida\n")
			writer.Flush()
		}
	}

}
