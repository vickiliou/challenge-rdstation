# Customer Success Balancing

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
