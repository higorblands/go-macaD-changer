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
6. 
## Exemplo

Para modificar o endereço MAC da interface de rede eth0 para "00:11:22:33:44:55", execute o seguinte comando:

sudo ./mac_changer eth0 00:11:22:33:44:55

## Notas

- Lembre-se de que a alteração do endereço MAC de uma interface de rede pode afetar a conectividade de rede. Use com cuidado e apenas para fins legais e autorizados.
- Este programa foi projetado para sistemas Linux. Pode não funcionar em outros sistemas operacionais.
- Este é um projeto simples e não é garantido que funcione em todas as configurações de rede ou sistemas. Use-o por sua conta e risco.

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir problemas (issues) ou enviar solicitações de pull (pull requests) para melhorar este projeto.

## Licença

Este projeto é distribuído sob a licença [MIT](LICENSE). Consulte o arquivo `LICENSE` para obter mais detalhes.
