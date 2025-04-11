# go-file-vault
# 🔐 go-file-vault

Secure file uploads, encryption, and user authentication — all powered by **Go** and **Fiber**.  
Think of it as your own little vault, where you call the shots. 🔒

---

## 🚀 What is this?

A backend project built with love, logic, and late-night hustle — designed to:

- 🔐 Let users **sign up and log in** securely (JWT-powered).
- 📤 **Upload files** with AES encryption before saving.
- 📁 **List & download** files if you're legit (auth-guarded).
- 💬 Ready to plug into a frontend (like Vue) or run as an API.

---

## ⚙️ Features

✨ Simple, but solid:
- 🛡️ AES-256 **file encryption**
- 👤 **User auth** with hashed passwords (bcrypt) and JWT
- 📁 **Upload**, **List**, and **Download** files
- 🌍 CORS-enabled for local frontend dev (`localhost:5173`)
- 💾 In-memory user store (easy to switch to a DB)

---

## 📁 Folder Structure

---go-file-vault/ 
├── backend/ │ 
├── controllers/ # Upload, Auth, List, Download │ 
├── middleware/ # JWT middleware │ 
├── models/ # User model 
│ └── uploads/ # Encrypted files live here 
└── main.go # App entry point

## 🔧 Setup

### 1. Clone it

```bash
git clone https://github.com/dharu630/go-file-vault.git
cd go-file-vault/backend
go mod tidy
go run main.go

**##🤝 **Contribution & Credits****

This is a learning + portfolio project, but feel free to fork, star ⭐️, and build your own twist on it.

Made with ❤️ by @dharu630
