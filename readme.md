# NUDL :

> **Project Name**: NUDL  
> **Function**: Transform JSON files into RESTful APIs with an interactive terminal UI (TUI)  

---
🚧 Project under active devlopment - Not ready for production
---
## 🎯 1. Features
- ✅ Interactive TUI to load and visualize JSON API setup
- ✅ Parse JSON into dynamic REST endpoints (e.g. `/users`, `/users/1`)
- ✅ Run a local API server serving the generated endpoints
- ✅ Support GET endpoints by default, extensible later
- ✅ CLI flags support for headless mode (optional)
- ✅ .....
---

## 📊 2. Input & Output Example

### Sample JSON File:
```json
{
  "users": [
    { "id": 1, "name": "Hamza" },
    { "id": 2, "name": "Amina" }
  ],
  "products": [
    { "id": 10, "name": "Laptop" }
  ]
}
```

### Auto-Generated Endpoints:
- `GET /users`
- `GET /users/1`
- `GET /products`
- `GET /products/10`

---

## License

This project is licensed under the [MIT License](./LICENSE)
