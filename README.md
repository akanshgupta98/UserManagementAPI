# 🧑‍💼 User Management API (Go HTTP Server)

This is a basic User Management API built using Go’s standard `net/http` library. It demonstrates foundational web server development in Go including route handling, middleware, structured JSON responses, query parameter parsing, and JSON request/response handling.

---

## ✨ Features

- ✅ Health check at `/`
- 📄 Fetches user by email at `GET /user`
- 🧑 Creates a user from JSON payload at `POST /createUser`
- 💬 Echoes query parameter at `GET /echo?name=xyz`
- 🧾 Middleware for logging request method and URI

---

## 🧪 API Endpoints

| Method | Endpoint                  | Description                                           |
|--------|---------------------------|-------------------------------------------------------|
| GET    | `/`                       | Health check (plain text response)                   |
| GET    | `/user`                   | Fetches by email query     |
| POST   | `/createUser`             | Accepts a user JSON and returns confirmation         |
| GET    | `/echo?name=xyz`          | Echoes back the `name` from the query param          |

---

### 🧾 Example Usage

#### 🔹 GET `/user?email=akansh.gupta@example.com`
- Fetches the user by email and returns user data:
```json
{
  "firstname": "Akansh",
  "lastname": "Gupta",
  "email": "akansh.gupta@example.com"
}
```
#### 🔹 POST /user
- Creates the user:
  Payload:
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com"
}
```
  Response:
```json
{
  "message": "User created successfully",
  "data": {
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com"
  }
}
```

#### 🔹 Get /echo?msg=hello
- Returns the same message
- Response:
```json
{
  "message": "hello"
}
```





  
