# User Coins API (Learning Go)

This is a from-scratch project to practice Go fundamentals:
- **Layered Architecture**: Separating handlers, middleware, and data.
- **Custom Middleware**: Implementing a request logger and auth checker(dummy).

## Project Structure
- `cmd/api/main.go`: Server and Router setup.
- `internal/handlers/`: HTTP logic for Tasks.
- `internal/middleware/`: Request logging and security.
- `internal/tools/`: Task data and mock database logic.

## How to Run
```bash
go mod tidy
go run ./cmd/api

## BASE URL:
Base URL: http://localhost:8000
