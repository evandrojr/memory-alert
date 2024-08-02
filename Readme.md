# Memory Alert

Este é um utilitário simples escrito em Go para monitorar a quantidade de memória disponível em um sistema Linux. Quando a memória disponível cai abaixo de um certo limite, uma notificação é enviada.

## Funcionalidades

- Monitora a memória disponível no sistema.
- Envia uma notificação de alerta quando a memória disponível é inferior a 1500 MB.
- Usa o pacote `notify` para exibir notificações na área de trabalho.

## Pré-requisitos

- Este programa foi projetado para sistemas Linux, pois depende do arquivo `/proc/meminfo` para obter informações sobre a memória.
- Go instalado no sistema. Você pode baixar e instalar o Go a partir de [golang.org](https://golang.org/).


## Como usar
Clone o repositório ou baixe o código fonte para o seu diretório de trabalho:

Copiar código
```bash
git clone https://github.com/seu_usuario/memory-alert.git
cd memory-alert
```
Certifique-se de que você tem uma imagem chamada ram.png no mesmo diretório onde o executável será executado. Esta imagem será usada na notificação.

Compile o programa usando o comando:

```bash
Copiar código
go build -o memory-alert
```
Execute o programa:

```bash
Copiar código
./memory-alert
```
Coloque o memory-alert para inicar quando seu gerenciador de janelas iniciar

O programa irá monitorar a memória disponível continuamente a cada 2 segundos e exibirá uma notificação quando a memória disponível for inferior a 1500 MB.

## Personalização
Limite de Memória: Você pode alterar o limite de memória passando o valor por linha de comando.

Intervalo de Verificação: O intervalo entre as verificações de memória é atualmente de 2 segundos. Você pode ajustar isso alterando o valor time.Sleep(2 * time.Second) para um valor diferente.

## Licença
Este projeto está licenciado sob a MIT License - veja o arquivo LICENSE para mais detalhes.