# 📝 Go Fiber Blog API with Authentication

A complete **blog backend** built with [Go Fiber] and **JWT Authentication** using HS256.  
Supports:
- **Public blog browsing**
- **Admin-only blog management**
- **Secure JWT-based login**
- **Role-based access control (RBAC)**
- **Postgres db for the storag**

---

## 📖 Table of Contents
- [Overview](#-overview)
- [Features](#-features)
- [Setup Instructions](#-setup-instructions)

---

## 📌 Overview
This project is a **RESTful API** for a blog platform where:
- **Anyone** can read blog posts.
- **Only Admins** can create or delete blog posts.
- **JWT Authentication** is used to protect certain routes.
- **HS256 (HMAC-SHA256)** is used for signing tokens.

It’s perfect as a **starter template** for content-based apps with authentication and role-based permissions.

---

## ✨ Features
- Public blog listing
- Admin-only blog creation & deletion
- Secure JWT-based login
- HS256 token signing
- Role-based access middleware
- Organized & extendable folder structure

---

## ⚙️ Setup Instructions

### Clone the repository
- git clone https://github.com/avdhesh-15/go-blog.git

### Install the dependencies
- go mod tidy

### Create .env 
- PORT= Your Port
- DB_STRING= Your Postgres URL 
- JWT_SECRET= Your JWT secret 

### Generate your secret key 
- openssl rand -base64 32

### Run the project
- cd cmd
- go run main.go

---

**Now your project is ready to be tested on Postman or any other Http Client**

---
