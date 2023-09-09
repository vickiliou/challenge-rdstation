# Customer Success Balancing

Customer Success Balancing é um tech challenge proposto pela empresa RDStation, onde oferece uma solução para distribuir clientes entre Customer Success, garantindo que o mais ocupado seja identificado para fazer um plano de ação para contratação.

## Como rodar os testes

### Opção 1: Rodar localmente

#### Versão mínima

- Go 1.16

#### Terminal

No terminal, execute os comandos na raíz do projeto:

```bash
go test
```

### Opção 2: Rodar com docker

Dependências:
Ter docker instalado

#### Terminal

No terminal, execute os comandos na raíz do projeto:

```bash
docker build -t challenge-rdstation .
docker run --rm challenge-rdstation
```
