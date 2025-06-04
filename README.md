
## Sobre este projeto

Este projeto demonstra um serviço gRPC.

A estrutura do projeto é organizada da seguinte forma:

- **proto/**: Contém as definições do Protocol Buffer para o serviço gRPC.
    - Esses arquivos definem os métodos do serviço e a estrutura dos dados trocados entre o cliente e o servidor.
    - Os arquivos `.proto` são usados para gerar o código do serviço gRPC usando o compilador de Protocol Buffer (`protoc`).

- **internal/**: Contém os detalhes de implementação interna do serviço gRPC.
    - Este diretório é dividido em subdiretórios com base na funcionalidade:
        - **server**: Contém a implementação do servidor gRPC.
        - **db**: Contém a lógica de acesso ao banco de dados.
        - **logic**: Contém a lógica de negócios da aplicação.

- **cmd/**: Contém o código principal da aplicação que inicia o servidor gRPC.
    - Este diretório normalmente contém um arquivo `main.go` que importa os pacotes necessários do diretório `internal/` e inicia o servidor gRPC.

- **Makefile**: Contém um conjunto de regras para automatizar tarefas como:
    - Construir o serviço gRPC.
    - Gerar o código do serviço gRPC a partir dos arquivos `.proto`.
    - Executar testes.
    - Validar o código.

## gRPC (gRPC Remote Procedure Call)

gRPC é um framework de chamada de procedimento remoto (RPC) moderno e de alto desempenho desenvolvido pelo Google. Ele usa Protocol Buffers como sua linguagem de definição de interface e suporta múltiplas linguagens de programação. gRPC permite que aplicações cliente e servidor se comuniquem de forma transparente e facilita a construção de sistemas conectados.

## Comparação com REST

| Característica    | gRPC                                  | REST                                     |
| ---------------- | ------------------------------------- | ---------------------------------------- |
| Protocolo         | HTTP/2                                | HTTP/1.1 ou HTTP/2                       |
| Formato de Dados | Protocol Buffers                        | JSON, XML, etc.                          |
| Interface        | Definição de Serviço (e.g., arquivo .proto) | Documentação da API (e.g., OpenAPI/Swagger) |
| Geração de Código | Integrada                              | Ferramentas externas necessárias          |
| Desempenho       | Geralmente mais rápido                  | Pode ser mais lento devido a payloads maiores |
| Streaming        | Suporte integrado                       | Suporte limitado                          |
| Casos de Uso     | Microsserviços, sistemas poliglota     | APIs Web, integrações mais simples        |

## Sugestões para uso em sistemas bancários

gRPC pode ser particularmente útil em sistemas bancários para:

- **Arquitetura de microsserviços:** Construir um sistema como uma coleção de serviços pequenos e independentes que se comunicam via gRPC.
- **Transações de alto desempenho:** Lidar com um grande número de transações com baixa latência.
- **Comunicação segura:** Usar os recursos de segurança integrados do gRPC (e.g., TLS) para proteger dados sensíveis.
- **Ambientes poliglota:** Integrar sistemas escritos em diferentes linguagens de programação.

Exemplos de casos de uso:

- Gerenciamento de contas: Criar, atualizar e excluir contas de clientes.
- Processamento de transações: Transferir fundos entre contas.
- Detecção de fraudes: Analisar transações em tempo real para identificar atividades fraudulentas.
- Relatórios: Gerar relatórios sobre a atividade da conta e o desempenho do sistema.


## Testando o serviço gRPC com Evans

Evans é uma ferramenta CLI que fornece uma maneira conveniente de interagir com serviços gRPC. Ele permite que você liste os serviços e métodos disponíveis, inspecione as mensagens e invoque os métodos com facilidade.

### Instalação do Evans

Você pode instalar o Evans usando o Go:

```bash
go install github.com/ktr0731/evans@latest
```

Certifique-se de ter o Go instalado e configurado corretamente em seu sistema.

### Usando Evans para interagir com o serviço gRPC

1.  **Navegue até o diretório do projeto:**

    ```bash
    cd /home/mauricio.silva/code/estudo/go-grpc-example
    ```

2.  **Execute o Evans e especifique o arquivo .proto:**

    ```bash
    evans -p 50051 --proto proto/seu_servico.proto repl
    ```

    Substitua `proto/seu_servico.proto` pelo caminho correto para o seu arquivo `.proto` e `50051` pela porta em que seu servidor gRPC está rodando.

3.  **Listar serviços:**

    Dentro do shell Evans, você pode listar os serviços disponíveis:

    ```
    show service
    ```

4.  **Listar métodos de um serviço:**

    Para listar os métodos de um serviço específico:

    ```
    show service SeuServico
    ```

    Substitua `SeuServico` pelo nome do serviço que você deseja explorar.

5.  **Invocar um método:**

    Para invocar um método, use o comando `call`:

    ```
    call MetodoDoServico
    ```

    Evans irá solicitar que você insira os dados necessários para a requisição. Você pode inserir os dados no formato JSON.

    Exemplo:

    ```
    call CreateAccount
    ```

    ```json
    {
      "id": "123",
      "balance": 100
    }
    ```

    Evans enviará a requisição para o servidor gRPC e exibirá a resposta.

### Dicas adicionais

*   Use a opção `-r` para recarregar o arquivo `.proto` automaticamente quando ele for alterado.
*   Explore a documentação do Evans para descobrir todos os recursos e opções disponíveis.

### Alterando o arquivo .proto

1.  **Edite o arquivo `.proto`:**

    Abra o arquivo `.proto` (por exemplo, `proto/seu_servico.proto`) e modifique a definição do serviço ou as mensagens conforme necessário.

    Exemplo:

    ```protobuf
    syntax = "proto3";

    package seu_pacote;

    service SeuServico {
      rpc CreateAccount (CreateAccountRequest) returns (CreateAccountResponse);
      rpc GetAccount (GetAccountRequest) returns (GetAccountResponse);
    }

    message CreateAccountRequest {
      string id = 1;
      double balance = 2;
    }

    message CreateAccountResponse {
      string id = 1;
      double balance = 2;
    }

    message GetAccountRequest {
      string id = 1;
    }

    message GetAccountResponse {
      string id = 1;
      double balance = 2;
    }
    ```

2.  **Compile o arquivo `.proto`:**

    Após modificar o arquivo `.proto`, você precisa compilá-lo para gerar o código gRPC correspondente. Use o seguinte comando:

    ```bash
    make proto
    ```

    Este comando usa o `protoc` (compilador de Protocol Buffer) para gerar os arquivos `.pb.go` que contêm o código gRPC. Certifique-se de ter o `protoc` e o plugin `protoc-gen-go-grpc` instalados.

### Executando o servidor gRPC

1.  **Navegue até o diretório `cmd`:**

    ```bash
    cd cmd
    ```

2.  **Execute o arquivo `main.go`:**

    Use o comando `go run` para executar o servidor gRPC:

    ```bash
    go run main.go
    ```

    Certifique-se de que o arquivo `main.go` contenha o código para iniciar o servidor gRPC e registrar os serviços.

    Exemplo de `main.go`:

    ```go
    package main

    import (
        "fmt"
        "log"
        "net"

        "google.golang.org/grpc"
        pb "seu_projeto/proto" // Importe o pacote proto gerado
        "seu_projeto/internal/server" // Importe a implementação do servidor
    )

    const (
        port = ":50051"
    )

    func main() {
        lis, err := net.Listen("tcp", port)
        if err != nil {
            log.Fatalf("failed to listen: %v", err)
        }
        s := grpc.NewServer()
        pb.RegisterSeuServicoServer(s, &server.SeuServicoServer{}) // Registre o serviço gRPC
        fmt.Printf("Server listening at %v", lis.Addr())
        if err := s.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    }
    ```

    Substitua `seu_projeto` pelo caminho correto para o seu projeto e ajuste os imports conforme necessário.