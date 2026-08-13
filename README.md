<h1 align="center">🔗 Short URL</h1>

<p align="center">
  Encurtador de URLs full stack com geração automática de QR Code.
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.24-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Vue.js-3-4FC08D?style=flat-square&logo=vuedotjs&logoColor=white" alt="Vue.js">
  <img src="https://img.shields.io/badge/TypeScript-5-3178C6?style=flat-square&logo=typescript&logoColor=white" alt="TypeScript">
  <img src="https://img.shields.io/badge/SQLite-3-003B57?style=flat-square&logo=sqlite&logoColor=white" alt="SQLite">
</p>

## Sobre o projeto

O **Short URL** transforma endereços longos em links curtos e gera um QR Code para facilitar o compartilhamento. A aplicação possui frontend responsivo, API REST e persistência local com SQLite.

Além da funcionalidade, o projeto demonstra organização por camadas, separação de responsabilidades, casos de uso, gateways e inversão de dependência.

## Funcionalidades

- Encurtamento de URLs;
- geração automática de identificadores curtos;
- criação de QR Code para cada link;
- redirecionamento para a URL original;
- persistência dos dados em SQLite;
- validação de URLs;
- feedback visual durante o processamento;
- tratamento de erros no frontend e na API.

## Tecnologias

### Frontend

- Vue 3;
- TypeScript;
- Vite;
- Tailwind CSS 4;
- Vue Router;
- Axios.

### Backend

- Go 1.24;
- Gin;
- SQLite com driver `modernc.org/sqlite`;
- UUID;
- ShortID;
- Go QR Code;
- CORS.

## Arquitetura

O backend mantém as regras de negócio independentes dos detalhes externos:

```text
backend/internal/
├── application/       # DTOs e casos de uso
├── domain/            # Entidades, objetos de valor e contratos
├── infrastructure/    # HTTP, SQLite, IDs e geração de QR Code
├── presentation/      # Controllers e rotas
└── shared/            # Configuração e composição de dependências
```

O frontend também separa domínio, aplicação, infraestrutura e apresentação:

```text
frontend/src/
├── application/
├── domain/
├── infrastructure/
├── presentation/
└── shared/
```

## Pré-requisitos

Antes de começar, tenha instalado:

- [Go](https://go.dev/dl/) 1.24 ou superior;
- [Node.js](https://nodejs.org/) 20 ou superior;
- [Yarn](https://yarnpkg.com/).

## Executando o projeto

Clone o repositório:

```bash
git clone https://github.com/marcosfrancomarinho/encutador-de-url.git
cd encutador-de-url
```

Instale as dependências:

```bash
yarn install
```

Inicie frontend e backend juntos:

```bash
yarn dev
```

A API será iniciada em:

```text
http://localhost:3030
```

### Execução separada

Backend:

```bash
cd backend
go mod tidy
go run main.go
```

Frontend:

```bash
cd frontend
yarn install
yarn dev
```

## Scripts disponíveis

| Comando | Descrição |
| --- | --- |
| `yarn dev` | Inicia frontend e backend simultaneamente |
| `yarn start:frontend` | Inicia somente o frontend |
| `yarn start:backend` | Inicia somente a API |
| `yarn install:frontend` | Instala as dependências do frontend |
| `yarn install:backend` | Atualiza as dependências do backend |

## Autor

Desenvolvido por [Marcos Marinho](https://github.com/marcosfrancomarinho).
