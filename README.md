# Short Url - Encurtador de URLs

Um aplicativo web para encurtar URLs de forma rápida, segura e eficiente, com geração de QR Code para cada link encurtado.

---

## 💻 Tecnologias

- **Frontend:** Vue 3, TypeScript, Tailwind CSS  
- **Backend:**  Go
- **Banco de dados:** SQLite 
- **Outras libs:** Axios, Gin, Shortid

---

## 🚀 Funcionalidades

- Encurtar URLs longas.  
- Gerar QR Code para cada URL encurtada.  
- Visualizar link encurtado e QR Code na interface.  
- Feedback de carregamento (loading) durante o processo.  
- Tratamento de erros e URLs inválidas.  

---

## ⚡ Instalação

```bash
yarn install
yarn dev

OU

cd frontend/
yarn install
yarn dev



cd backend/
go mod tidy
go run main.go
````

## 🧱  Estrutura do projeto
``` bash
backend/
  ├─ main.go
  ├─ internal/
  │   ├─ application/
  │   ├─ domain/
  │   ├─ infrastructure/
  │   ├─ presentation/
  │   └─ shared/
frontend/
  ├─ src/
  │   ├─ application/
  │   ├─ domain/
  │   ├─ infrastructure/
  │   ├─ presentation/
  │   ├─ shared/
  │   └─ main.ts
  └─ package.json
````