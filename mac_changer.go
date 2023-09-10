package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: sudo ./mac_changer <nome_da_interface> <novo_mac>")
		fmt.Println("Exemplo: sudo ./mac_changer eth0 00:11:22:33:44:55")
		os.Exit(1)
	}

	interfaceName := os.Args[1]
	newMAC := os.Args[2]

	// Desativar a interface
	if err := disableInterface(interfaceName); err != nil {
		fmt.Println("Erro ao desativar a interface:", err)
		os.Exit(1)
	}
	fmt.Printf("Interface %s desativada.\n", interfaceName)

	// Alterar o endereço MAC
	if err := changeMACAddress(interfaceName, newMAC); err != nil {
		fmt.Println("Erro ao alterar o endereço MAC:", err)
		os.Exit(1)
	}
	fmt.Printf("Endereço MAC da interface %s alterado para %s.\n", interfaceName, newMAC)
	
	// Ativar a interface
	if err := enableInterface(interfaceName); err != nil {
		fmt.Println("Erro ao ativar a interface:", err)
		os.Exit(1)
	}
	fmt.Printf("Interface %s ativada.\n", interfaceName)
}

func disableInterface(interfaceName string) error {
	cmd := exec.Command("ifconfig", interfaceName, "down")
	return cmd.Run()
}

func enableInterface(interfaceName string) error {
	cmd := exec.Command("ifconfig", interfaceName, "up")
	return cmd.Run()
}

func changeMACAddress(interfaceName, newMAC string) error {
	cmd := exec.Command("ifconfig", interfaceName, "hw", "ether", newMAC)
	return cmd.Run()
}