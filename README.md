# 📊 Commit Analyzer API

Uma API REST desenvolvida em **Go (Golang)** utilizando o framework **Gin**. O objetivo deste projeto é analisar repositórios do GitHub e gerar métricas sobre os commits, categorizando-os baseados no padrão [Conventional Commits](https://www.conventionalcommits.org/) (ex: `feat`, `fix`, `chore`, `refactor`, etc.).

## 🚀 Live Demo

A API está atualmente hospedada na Railway e pode ser acessada publicamente:

**Base URL:** `https://commit-analyzer-production.up.railway.app`

---

## 🛠️ Tecnologias Utilizadas

* [Go](https://go.dev/) - Linguagem principal.
* [Gin Web Framework](https://github.com/gin-gonic/gin) - Roteamento e HTTP de alta performance.
* [Railway](https://railway.app/) - Plataforma de Deploy e Hospedagem.

---

## 📖 Documentação da API

### Obter Métricas de Commits

Retorna a contagem total de commits e a distribuição por tipo (*type*) para um repositório específico.

**Endpoint:**
`GET /metrics/:user/:repo`

| Parâmetro | Tipo   | Descrição                                      |
| :-------- | :----- | :--------------------------------------------- |
| `user`    | string | Nome do usuário ou organização dono do repo.   |
| `repo`    | string | Nome do repositório.                           |

**Exemplo de Requisição:**

```bash
curl [https://commit-analyzer-production.up.railway.app/metrics/nome-do-usuario/nome-do-repo](https://commit-analyzer-production.up.railway.app/metrics/nome-do-usuario/nome-do-repo)
```

**Exemplo de Resposta:**
```bash
{
  "total": 58,
  "type": {
    "Another's commits": 2,
    "build": 6,
    "chore": 7,
    "docs": 1,
    "feat": 17,
    "fix": 7,
    "refactor": 7,
    "style": 9,
    "test": 2
  }
}
```
---

## 💻 Como Rodar Localmente

Siga os passos abaixo para executar a API na sua máquina.

### Pré-requisitos
* [Go](https://go.dev/dl/) 1.18 ou superior instalado.
* Git instalado.

### Passo a Passo

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/SEU-USUARIO/commit-analyzer.git](https://github.com/SEU-USUARIO/commit-analyzer.git)
    cd commit-analyzer
    ```

2.  **Instale as dependências:**
    ```bash
    go mod tidy
    ```

3.  **Configure o ambiente (Opcional):**
    Por padrão a aplicação roda na porta `8080`. Se desejar mudar:
    ```bash
    export PORT=3000
    ```

4.  **Execute a aplicação:**
    ```bash
    # Execute a partir da raiz do projeto
    go run cmd/setup/main.go
    ```
    *(Nota: Se o seu arquivo main estiver dentro de uma pasta, use: `go run cmd/setup/main.go`)*

5.  **Acesse no navegador:**
    `http://localhost:8080/metrics/seu-usuario/seu-repo`

---

## 🤝 Contribuindo

Contribuições são bem-vindas!

1.  Faça um Fork do projeto.
2.  Crie uma Branch para sua Feature (`git checkout -b feature/MinhaFeature`).
3.  Adicione suas mudanças (`git add .`).
4.  Comite suas mudanças (`git commit -m 'Adiciona a feature X'`).
5.  Faça o Push (`git push origin feature/MinhaFeature`).
6.  Abra um Pull Request.
7.  Só aguardar!

---

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.