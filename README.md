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
cd frontend
npm install
npm run dev

cd backend/cmd/app
go run main.go

backend/
  ├─ cmd/
  │   └─ app/
  │       └─ main.go
  ├─ internal/
  │   ├─ controllers/
  │   ├─ services/
  │   └─ database/
frontend/
  ├─ src/
  │   ├─ components/
  │   ├─ composables/
  │   └─ App.vue
  └─ package.json
