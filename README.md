# GOT - Go Template

![Version](https://img.shields.io/badge/version-0.1.0--alpha-blue)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Production-Ready Go template to kickstart your next Go lang Project.

## Table of Contents

- [Features](#features)
- [Setup](#setup)
- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Migrations](#migrations)
- [Contributing](#contributing)
- [License](#license)

## Features

- Database Support
- Docker Integration
- Makefile Commands
- CLI Support
- REST API Ready
- Database Migrations
- Environment Config
- Modular Structure
- Testing Framework

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/got-template.git
cd got-template
```

### 2. Run Setup

```bash
make setup
```

## Configuration

Edit the following files:

- `config/local.yaml` – App configuration (port, log level, etc.)
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

## Contributing

Feel free to open issues and submit PRs. All contributions are welcome!

## License

Licensed under the [MIT License](LICENSE).
