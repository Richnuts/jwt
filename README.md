# JWT CRUD API

This project is a **clean architecture-based** HTTP software built with the **Echo framework** in **Golang**. It provides a fully functional **JWT CRUD API** for authentication and authorization.

## 🛠 Tech Stack
- **Golang** – Backend language
- **Echo Framework** – Lightweight and high-performance web framework
- **JWT (JSON Web Token)** – Secure authentication
- **Clean Architecture** – Well-structured project following best practices

## 📂 Project Structure
```
├── config/         # Configuration files
├── delivery/       # HTTP handlers (controllers)
├── entities/       # Domain entities and models
├── repository/     # Data persistence (database layer)
├── usecase/        # Business logic and use cases
├── util/           # Utility functions
├── go.mod          # Go modules dependencies
├── go.sum          # Package checksums
├── server.go       # Main application entry point
└── tools.go        # Tool dependencies management
```

## 🚀 Getting Started

### 1️⃣ Clone the Repository
```sh
git clone https://github.com/Richnuts/jwt.git
cd jwt
```

### 2️⃣ Install Dependencies
Ensure Go is installed, then run:
```sh
go mod tidy
```

### 3️⃣ Run the Server
```sh
go run server.go
```
Your server should now be running and ready to handle requests.

## 🔑 API Endpoints

### Authentication
| Method | Endpoint  | Description |
|--------|----------|-------------|
| POST   | `/register` | Register a new user |
| POST   | `/login` | Authenticate and get JWT token |

### User Management
| Method | Endpoint  | Description |
|--------|----------|-------------|
| GET    | `/users` | Get all users |
| GET    | `/users/:id` | Get user by ID |
| PUT    | `/users/:id` | Update user details |
| DELETE | `/users/:id` | Delete user |

## 🛡️ Authentication & Authorization
- Upon successful login, the API returns a **JWT token**.
- Pass the token in the `Authorization` header as `Bearer <token>` to access protected routes.

## 🤝 Contributing
Contributions are welcome! Feel free to fork the repo, create a new branch, and submit a PR.

## 📜 License
This project is licensed under the **MIT License**.

---

🔥 *Built with Go and Echo to simplify JWT authentication!*
