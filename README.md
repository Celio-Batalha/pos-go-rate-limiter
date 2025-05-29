Desenvolver um rate limiter em Go que possa ser utilizado para controlar o tráfego de requisições para um serviço web. O rate limiter deve ser capaz de limitar o número de requisições com base em dois critérios:

1. Endereço IP: O rate limiter deve restringir o número de requisições recebidas de um único endereço IP dentro de um intervalo de tempo definido.

2. Token de Acesso: O rate limiter poderá limitar as requisições baseadas em um token de acesso único, permitindo diferentes limites de tempo de expiração para diferentes tokens. O Token deve ser informado no header no seguinte formato:
API_KEY: <TOKEN>

As configurações de limite do token de acesso devem se sobrepor as do IP. Ex: Se o limite por IP é de 10 req/s e a de um determinado token é de 100 req/s, o rate limiter deve utilizar as informações do token.

Requisitos:

- O rate limiter deve poder trabalhar como um middleware que é injetado ao servidor web
- O rate limiter deve permitir a configuração do número máximo de requisições permitidas por segundo.
- O rate limiter deve ter ter a opção de escolher o tempo de bloqueio do IP ou do Token caso a quantidade de requisições tenha sido excedida.
- As configurações de limite devem ser realizadas via variáveis de ambiente ou em um arquivo “.env” na pasta raiz.
- Deve ser possível configurar o rate limiter tanto para limitação por IP quanto por token de acesso.
- O sistema deve responder adequadamente quando o limite é excedido:
  - Código HTTP: 429
  - Mensagem: you have reached the maximum number of requests or actions allowed within a certain time frame
- Todas as informações de "limiter” devem ser armazenadas e consultadas de um banco de dados Redis. - Você pode utilizar docker-compose para subir o Redis.
- Crie uma “strategy” que permita trocar facilmente o Redis por outro mecanismo de persistência.
- A lógica do limiter deve estar separada do middleware.

Exemplos:

1. Limitação por IP: Suponha que o rate limiter esteja configurado para permitir no máximo 5 requisições por segundo por IP. Se o IP 192.168.1.1 enviar 6 requisições em um segundo, a sexta requisição deve ser bloqueada.
2. Limitação por Token: Se um token abc123 tiver um limite configurado de 10 requisições por segundo e enviar 11 requisições nesse intervalo, a décima primeira deve ser bloqueada.
3. Nos dois casos acima, as próximas requisições poderão ser realizadas somente quando o tempo total de expiração ocorrer. Ex: Se o tempo de expiração é de 5 minutos, determinado IP poderá realizar novas requisições somente após os 5 minutos.

Dicas:

- Teste seu rate limiter sob diferentes condições de carga para garantir que ele funcione conforme esperado em situações de alto tráfego.

Entrega:

- O código-fonte completo da implementação.
- Documentação explicando como o rate limiter funciona e como ele pode ser configurado.
- Testes automatizados demonstrando a eficácia e a robustez do rate limiter.
- Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
- O servidor web deve responder na porta 8080.


# pos-go-rate-limite

## Overview
Este projeto implementa um limitador de taxa em Go para controlar o tráfego de solicitações a um serviço web. O limitador de taxa pode limitar solicitações com base em endereços IP e tokens de acesso, proporcionando flexibilidade no gerenciamento de limites de solicitação.

## Features
- Integração de middleware para uso fácil com servidores web.
- Máximo de solicitações configuráveis permitidas por segundo.
- Opção para definir o tempo de bloqueio para IPs ou tokens quando os limites são excedidos.
- Configuração através de variáveis de ambiente ou um arquivo `.env'.
- Suporta limitação de taxa baseada em IP e baseada em token.
- Responde com código de status HTTP 429 quando os limites são excedidos, juntamente com uma mensagem descritiva.
- Utiliza o Redis para armazenar e consultar informações de limite de taxa, com suporte para o Docker.
- Padrão de estratégia implementado para permitir fácil troca do mecanismo de persistência.


## Setup Instructions
1. Clone the repository:
   ```
   git clone <repository-url>
   cd pos-go-rate-limite
   ```

3. Start the Redis service using Docker:
   ```
   docker-compose up -d
   ```
4. voce pode fazer o teste usando um postman ou extesion do vscode api.http

```
GET http://localhost:8080
Headers:
API_KEY: my-test-token

```

GET http://localhost:8080
Headers:
API_KEY:
```

## Usage
- Para testar o limitador de taxa, envie solicitações ao serviço web usando uma ferramenta como `curl' ou Postman.
- Incluir o token de acesso no cabeçalho da solicitação como se segue:
  ```
  API_KEY: <TOKEN>
  ```