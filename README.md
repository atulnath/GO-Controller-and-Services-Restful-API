# GO-Controller-and-Services-Restful-API

This is a simple RESTful API built with [Gin](https://github.com/gin-gonic/gin), a popular Go web framework. The application simulates a note-taking backend where users can retrieve and create notes associated with specific courses.


---

## 🚀 Features

- Retrieve a list of pre-defined notes (`GET /note_app/`)
- Add a new sample note (`POST /note_app/`)
- Modular MVC-like structure (Controller, Service)
- Lightweight and beginner-friendly design

---

## 📦 Endpoints

| Method | Endpoint        | Description                 |
|--------|------------------|-----------------------------|
| GET    | `/note_app/`     | Fetch all available notes   |
| POST   | `/note_app/`     | Create a new sample note    |

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin Web Framework
- **Architecture:** Service-Oriented + Clean Routing

---

## 📌 Sample Response

### GET `/note_app/`

```json
{
  "notes": [
    {
      "course_name": "Math",
      "course_id": "M101",
      "note": "Algebra notes"
    },
    {
      "course_name": "Science",
      "course_id": "S101",
      "note": "Physics notes"
    }
  ]
}
```
---
POST /note_app/
```json
{
  "notes": [
    {
      "course_name": "German",
      "course_id": "DE101",
      "note": "Deutsch notes"
    }
  ]
}
```
---
## 📚 Future Improvements
Add persistent storage (e.g., PostgreSQL, SQLite)

Implement CRUD operations with real data input

Add Swagger API documentation

Add authentication/authorization
