
# 🔄 API Multithreading - Busca de CEP Mais Rápida 🏃‍♂️

Bem-vindo ao repositório do desafio da **Pós GoExpert 2024**! Este projeto foi desenvolvido por **Paulo Nunes** para demonstrar conhecimentos em **Go**, abrangendo conceitos como **APIs**, **multithreading**, **goroutines**, **channels**, e **contextos com timeout**.

---

## 📋 Desafio Proposto

O desafio consiste em buscar o resultado mais rápido entre duas APIs distintas para consulta de CEP, descartando a resposta mais lenta. As APIs utilizadas são:

1. **BrasilAPI**: [https://brasilapi.com.br/api/cep/v1/](https://brasilapi.com.br/api/cep/v1/)
2. **ViaCEP**: [http://viacep.com.br/ws/](http://viacep.com.br/ws/)

### Requisitos:

1. Realizar requisições simultâneas às duas APIs para um CEP fornecido.
2. Aceitar a resposta mais rápida e descartar a resposta mais lenta.
3. Exibir no terminal:
   - O resultado do endereço (CEP, logradouro, bairro, localidade, UF).
   - Qual API retornou a resposta.
4. Limitar o tempo de resposta a **1 segundo**. Caso nenhuma API responda nesse tempo, exibir um erro de timeout.

---

## 🚀 Implementação

### Ferramentas e Técnicas Utilizadas

1. **Multithreading em Go**:
   - Utilização de **goroutines** para realizar chamadas simultâneas às duas APIs.
   - Comunicação entre threads usando **channels**.

2. **Controle de Timeout**:
   - Uso do pacote **`context`** para limitar o tempo total de execução a **1 segundo**.

3. **Estrutura do Projeto**:
   - Organização em diretórios seguindo boas práticas.

---

## ⚙️ Como Executar o Projeto

### Pré-requisitos

- [Go instalado](https://golang.org/dl/) (versão 1.20 ou superior).
- Conexão com a internet para acessar as APIs.

### Passo a Passo

1. Clone o repositório:
   ```bash
   git clone https://github.com/paulnune/goexpert-api-multithreading-go.git
   cd goexpert-api-multithreading-go
   ```

2. Inicialize o módulo Go:
   ```bash
   go mod init api-multithreading-go
   ```

3. Instale dependências e organize o projeto:
   ```bash
   go mod tidy
   ```

4. Execute o projeto:
   ```bash
   go run cmd/main.go
   ```

---

## 📂 Estrutura do Projeto

```
goexpert-api-multithreading-go/
├── configs/
│   └── config.go
├── dto/
│   └── address.go
├── internal/
│   ├── api/
│   │   ├── brasilapi.go
│   │   └── viacep.go
│   └── handler/
│       └── handler.go
├── cmd/
│   └── main.go
├── go.mod
└── go.sum
```

---

## 📝 Exemplos de Saída

### Caso Bem-sucedido
```
Resposta ViaCEP: {CEP:01153-000 Logradouro:Rua Vitorino Carmilo Bairro:Barra Funda Localidade:São Paulo UF:SP Provider:ViaCEP}
```

---

## 👨‍💻 Autor

**Paulo Henrique Nunes Vanderley**  
- 🌐 [Site Pessoal](https://www.paulonunes.dev/)  
- 🌐 [GitHub](https://github.com/paulnune)  
- ✉️ Email: [paulo.nunes@live.de](mailto:paulo.nunes@live.de)  
- 🚀 Aluno da Pós **GoExpert 2024** pela [FullCycle](https://fullcycle.com.br)

---

## 🎉 Agradecimentos

Este repositório foi desenvolvido com muita dedicação para a **Pós GoExpert 2024**. Agradeço à equipe da **FullCycle** por proporcionar uma experiência incrível de aprendizado! 🚀