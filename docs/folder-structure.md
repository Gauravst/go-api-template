# Folder Structure (GOT - Go Template)

### Root Level:

- `cmd/`
  Main application entry point. For example, `cmd/got-api/` contains the code that starts your Go API server.

- `configs/`
  Configuration files (JSON). You can define different files for different environments like dev, prod, etc.

- `docs/`
  Project documentation (you are probably reading this from here). Good place to put extra docs like folder structure, setup, usage tips, etc.

- `internal/`
  Application-specific packages. This code is private to your app and cannot be imported by others.

  - `internal/api/` – HTTP routes, handlers, and middleware.
  - `internal/cli/` – CLI commands and utilities.
  - `internal/config/` – Load and manage app configs.
  - `internal/database/` – DB connection and setup logic.
  - `internal/models/` – App data models (e.g. User, Token) for ORM/SQL.
  - `internal/dto/` – Request and response data structs.
  - `internal/repositories/` – DB queries and data access.
  - `internal/services/` – Business logic and processing.
  - `internal/utils/` – Helpers like JWT, email, hashing.

- `migrations`  
  SQL migration files used for database schema changes. Managed via `make migrate-up` and `make migrate-down` commands.

- `test`  
  Contains end-to-end (e2e) and integration tests.  
  (Note: Unit tests are placed alongside the logic files using `_test.go` suffix, like `user_test.go`.)

- `Makefile`
  Automation commands for local development (`make setup`, `make dev`, etc.)

- `Dockerfile`  
  Defines how to build and run the app inside a Docker container.

- `.dockerignore`  
  Lists files and folders to exclude from the Docker build context.

- `.env`  
  Environment-specific variables used during local development.

- `.env.example`  
  A sample environment file to show required variables. Used as a reference.

- `go.mod` / `go.sum`
  Go module definition and dependencies.

- `README.md`
  Main documentation file with instructions, setup, features, etc.

Of course Gaurav bhai! Here's the improved version of your summary:

### Summary:

This folder structure follows Go best practices. After running the `make setup` command, only the folders relevant to your selected features will remain — giving you a clean, minimal, and ready-to-use project structure.
