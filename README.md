  # MAC Changer em GoLang

Este é um aplicativo simples escrito em GoLang para modificar o endereço MAC de uma interface de rede em um sistema Linux. O endereço MAC (Media Access Control) é uma identificação única atribuída a dispositivos de rede.
Modificar o endereço MAC pode ser útil para fins de segurança, teste ou solução de problemas de rede.

## Pré-requisitos

- GoLang instalado na sua máquina ([Instalação GoLang](https://golang.org/doc/install))
- Permissões de superusuário (root) para alterar o endereço MAC.

## Uso

1. Clone este repositório para a sua máquina local:

git clone https://github.com/seu-usuario/mac-changer-golang.git

2. Navegue até o diretório do projeto:

cd mac-changer-golang

3. Compile o código Go:

go build mac_changer.go

4. Execute o programa com privilégios de superusuário (root) e forneça o nome da interface de rede e o novo endereço MAC que você deseja definir:

Substitua `<nome_da_interface>` pelo nome da interface de rede que deseja modificar (por exemplo, eth0) e `<novo_mac>` pelo novo endereço MAC desejado (no formato XX:XX:XX:XX:XX:XX).

5. O programa exibirá mensagens de sucesso ou erro com base na operação.

## Exemplo

Para modificar o endereço MAC da interface de rede eth0 para "00:11:22:33:44:55", execute o seguinte comando:

sudo ./mac_changer eth0 00:11:22:33:44:55

## Notas

- Lembre-se de que a alteração do endereço MAC de uma interface de rede pode afetar a conectividade de rede. Use com cuidado e apenas para fins legais e autorizados.
- Este programa foi projetado para sistemas Linux. Pode não funcionar em outros sistemas operacionais.
- Este é um projeto simples e não é garantido que funcione em todas as configurações de rede ou sistemas. Use-o por sua conta e risco.

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir problemas (issues) ou enviar solicitações de pull (pull requests) para melhorar este projeto.


----- ENG

# MAC Changer in GoLang

This is a simple application written in GoLang to modify the MAC address of a network interface on a Linux system. MAC (Media Access Control) address is a unique identification assigned to network devices.
Modifying the MAC address can be useful for security, testing, or network troubleshooting purposes.

## Prerequisites

- GoLang installed on your machine ([GoLang Installation](https://golang.org/doc/install))
- Superuser (root) permissions to change MAC address.

## Usage

1. Clone this repository to your local machine:

git clone https://github.com/seu-usuario/mac-changer-golang.git

2. Navigate to the project directory:

cd mac-changer-golang

3. Compile the Go code:

go build mac_changer.go

4. Run the program with superuser (root) privileges and provide the network interface name and new MAC address you want to set:

Replace `<interface_name>` with the name of the network interface you want to modify (for example, eth0) and `<new_mac>` with the new desired MAC address (in the format XX:XX:XX:XX:XX:XX).

5. The program will display success or error messages based on the operation.

## Example

To modify the MAC address of the eth0 network interface to "00:11:22:33:44:55", run the following command:

sudo ./mac_changer eth0 00:11:22:33:44:55

## Grades

- Remember that changing the MAC address of a network interface may affect network connectivity. Use with caution and only for legal and authorized purposes.
- This program is designed for Linux systems. May not work on other operating systems.
- This is a simple project and is not guaranteed to work on all network configurations or systems. Use it at your own risk.

## Contributions

Contributions are welcome! Feel free to open issues or submit pull requests to improve this project.

## License

This project is distributed under the [MIT] license (LICENSE). See the `LICENSE` file for more details.

## Licença

Este projeto é distribuído sob a licença [MIT](LICENSE). Consulte o arquivo `LICENSE` para obter mais detalhes.
