# GOT - Go Template

![Version](https://img.shields.io/badge/version-0.2.0--alpha-blue)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Production-Ready Go template to kickstart your next Go lang Project.

## Table of Contents

- [Features](#features)
- [Setup](#setup)
- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Migrations](#migrations)
- [Documentation](#Documentation)
- [Contributing](#contributing)
- [License](#license)

## Features

- Database Support
- ORM Support
- Docker Integration
- Makefile Commands
- CLI Support
- REST API Ready
- Database Migrations
- Environment Config
- Modular Structure
- Testing e2e/integration/unit Support

## Setup

### 1. Clone the Repository in Your Project Directory

```bash
git clone https://github.com/gauravst/got.git .
```

### 2. Create .env file from .env.example

```bash
cp .env.example .env
```

### 3. Run Setup

```bash
make setup
```

## Configuration

Edit the following files:

- `config/local.yaml` – App configuration (port, host etc.)
- `.env` – Secrets and environment variables

## Dependencies

```bash
go mod download
```

## Migrations

```bash
make migrate-up    # Apply migrations
make migrate-down  # Rollback migrations
```

## Documentation

- [Folder Structure](docs/folder-structure.md)

## Contributing

Feel free to open issues and submit PRs. All contributions are welcome!

## License

Licensed under the [MIT License](LICENSE).
