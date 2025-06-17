# Effective Octo Memory - REST API

A Go-based REST API with a domain-oriented architecture inspired by Domain-Driven Design principles.

## Project Structure

This project uses a flat, domain-oriented architecture:

```
server/
├── accounts/            # Accounts domain
│   ├── handler.go       # HTTP handlers
│   ├── interfaces.go    # Repository interfaces
│   ├── model.go         # Domain model
│   ├── service.go       # Business logic
│   └── store.go         # Data access implementation
├── config/              # Application configuration
│   ├── config.go        # Environment configuration
│   └── dependencies.go  # Dependency initialization
├── db/                  # Database utilities
├── routes/              # Routing utilities
├── server/              # Server setup and initialization
├── users/               # Users domain
│   ├── handler.go       # HTTP handlers
│   ├── interfaces.go    # Repository interfaces
│   ├── model.go         # Domain model
│   ├── service.go       # Business logic
│   └── store.go         # Data access implementation
├── utils/               # Utility functions
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
└── main.go              # Application entry point
```

## Getting Started

### Prerequisites

- Go 1.24.4 or higher
- PostgreSQL

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/effective-octo-memory.git
   cd effective-octo-memory/server
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Create a `.env` file with your configuration:

   ```bash
   cp .env.example .env
   # Edit .env with your settings
   ```

4. Run database migrations:

   ```bash
   go run main.go -migrate
   ```

5. Start the server:
   ```bash
   go run main.go
   ```

## Configuration

The application uses environment variables for configuration. You can set these directly or use a `.env` file.

### Environment Variables

#### Server Configuration

- `SERVER_HOST`: Server host address (default: `0.0.0.0`)
- `SERVER_PORT`: Server port (default: `1234`)
- `SERVER_ENV`: Environment (`development`, `production`, `test`) (default: `development`)

#### Database Configuration

- `DB_HOST`: Database host (default: `localhost`)
- `DB_PORT`: Database port (default: `5432`)
- `DB_USERNAME`: Database username (default: `user`)
- `DB_PASSWORD`: Database password (default: `password`) _(required)_
- `DB_NAME`: Database name (default: `mydb`)
- `DB_SSL_MODE`: Database SSL mode (default: `disable`)

#### Security Configuration

- `JWT_SECRET_KEY`: Secret key for JWT token signing _(required)_
- `UNPROTECTED_ROUTES`: Comma-separated list of routes that don't require authentication (e.g., `login,register,health`)

### Example .env File

```
# Server settings
SERVER_HOST=0.0.0.0
SERVER_PORT=1234
SERVER_ENV=development

# Database settings
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=secure_password
DB_NAME=octodb
DB_SSL_MODE=disable

# Security settings
JWT_SECRET_KEY=your-super-secret-key-here
UNPROTECTED_ROUTES=login,register,health
```

## API Endpoints

### Authentication

- `POST /login`: User login
- `POST /register`: User registration

### Users

- `GET /users`: Get all users
- `POST /users`: Create a new user

### Accounts

- `GET /accounts`: Get all accounts
- `POST /accounts`: Create a new account

## Development

### Running Database Migrations

```bash
go run main.go -migrate
# or the short form
go run main.go -m
```

### Development Mode

The application runs in development mode by default. To switch to production:

```
SERVER_ENV=production go run main.go
```

## License

[MIT](LICENSE)
