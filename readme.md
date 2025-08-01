# NUDL

> **v0.0.1 — Release**

**NUDL** is a simple command-line tool that turns your local JSON files into a fully functional RESTful API in seconds — all without a server.

Built with **Go** and a minimalist **TUI (Terminal UI)**, NUDL is ideal for prototyping, local testing, and demoing data structures.

![NUDL Demo](./assets/demo.gif)

---

## 🔧 Features

- 🗂️ Turn any JSON file into REST-style endpoints
- 🧭 Auto-generates `/collection` and `/collection/:id`
- 💻 Cross-platform binaries (Linux, Windows, macOS)
- 📦 CLI-first with clean interactive TUI

---

## 📥 Download

| Platform | Binary |
|----------|--------|
| Windows  | [`nudl.exe`](./download/v0.0.1/windows/nudl.exe) |
| Linux    | [`nudl`](./download/v0.0.1/linux/nudl)         |
| macOS    | [`nudl`](./download/v0.0.1/mac/nudl)           |

---

## 🚀 Quick Start

```bash
# 1. Download and give execution permission
chmod +x nudl

# 2. Run the tool with your JSON file
./nudl new -p 3000 -j ./example.json

# 3. For help
./nudl -h
```

Given an input like:
```json
{
  "users": [
    { "id": 1, "name": "Alice" },
    { "id": 2, "name": "Bob" }
  ]
}
```
NUDL will start a local API and serve:
- `/users` → all users
- `/users/1` → user with ID = 1

---

## 📁 JSON Structure Notes

- Top-level keys must map to arrays.
- Each object in the array must include an `id` field (number).

Example:
```json
{
  "products": [
    { "id": 1, "title": "Chair" },
    { "id": 99999, "title": "Table" }
  ]
}
```

---

## 📜 License

[MIT](./LICENSE)

---

## 🧪 Feedback & Contributions

This is the MVP release. Feedback, ideas, and issues are welcome via [GitHub Issues](https://github.com/imzadev/nudl/issues).

---

Made with 💻 by [imzadev](https://github.com/imzadev)
