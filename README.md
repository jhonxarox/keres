# Multi-Finance Backend Service

This is the backend service for the Multi-Finance application. It is built using Golang with the Gin framework, and PostgreSQL as the database. The application includes customer, transaction, and limit management services, and supports features such as pagination and automated database migrations.

---

## Features

- Customer, transaction, and limit management services.
- Pagination support for API endpoints.
- Automated database migrations using `golang-migrate`.
- Containerized deployment with Docker and Docker Compose.
- Health check endpoint.

---

## Prerequisites

- [Go](https://golang.org/doc/install) (Version 1.23.3 or later)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker](https://www.docker.com/) (for Dockerized setup)

---

## Environment Variables

The application requires the following environment variables:

| Variable      | Description                  | Default Value       |
|---------------|------------------------------|---------------------|
| `DB_USERNAME` | PostgreSQL username          | `your_postgres_user` |
| `DB_PASSWORD` | PostgreSQL password          | `your_postgres_password` |
| `DB_NAME`     | PostgreSQL database name     | `multi_finance`     |
| `DB_HOST`     | PostgreSQL host              | `127.0.0.1`         |
| `DB_PORT`     | PostgreSQL port              | `5432`              |

Create a `.env` file in the root of the project and add these variables.

Example `.env` file:
```plaintext
DB_USERNAME=your_postgres_user
DB_PASSWORD=your_postgres_password
DB_NAME=multi_finance
DB_HOST=127.0.0.1
DB_PORT=5432
```

## Running the Application
### 1. Run Without Docker

#### Step 1: Install Dependencies
Make sure you have Go installed, then run:

```bash
go mod tidy
```

#### Step 2: Set Up the Database

1. Start a local PostgreSQL instance.
2. Create the database:
```bash
createdb -U <username> multi_finance
```

3. Run migrations:
```bash
go run cmd/main.go
```

#### Step 3: Start the Application
Run the following command:

```bash
go run cmd/main.go
```
The server will start on http://localhost:8080.

### 2. Run With Docker
#### Step 1: Build and Start Containers
Use Docker Compose to build and run the application and database containers:

```bash
docker-compose up --build
```

#### Step 2: Verify the Setup
The application will be available at `http://localhost:8080`.

The PostgreSQL database will be available at `localhost:5432` with the credentials defined in `docker-compose.yml`.

## API Endpoints
### Health Check
- URL: `GET /health`
- Response:
```json
{
    "message": "Application is running!"
}
```

### Customer Endpoints
- Get All Customers: `GET /customers?page=<page>&limit=<limit>`
- Create Customer: `POST /customers`
```json
{
    "nik": "123456789",
    "full_name": "John Doe",
    "legal_name": "Johnathan Doe",
    "place_of_birth": "New York",
    "date_of_birth": "1990-01-01",
    "salary": 50000
}
```

### Transaction Endpoints
- Get All Transactions: `GET /transactions?page=<page>&limit=<limit>`
- Create Transaction: `POST /transactions`
```json
{
    "customer_id": 1,
    "contract_number": "T12345",
    "otr": 20000,
    "admin_fee": 500,
    "installment_amount": 1000,
    "interest_amount": 200,
    "asset_name": "Car"
}
```

### Limit Endpoints
- Get Customer Limits: `GET /limits/<customer_id>`
- Create Limit: `POST /limits`
```json
{
    "customer_id": 1,
    "tenor": 12,
    "limit_amount": 10000.50
}
```

## Running Tests
To run unit tests for the application:

```bash
go test ./... -v
```

## Directory Structure
```plaintext
.
├── cmd/                     # Application entry points
│   ├── main.go              # Main application
├── config/                  # Configuration files
├── internal/                # Application business logic
│   ├── customer/            # Customer module
│   ├── transaction/         # Transaction module
│   ├── limit/               # Limit module
│   ├── database/            # Database initialization
├── migrations/              # Database migrations
├── pkg/                     # Shared utilities
├── go.mod                   # Go module dependencies
├── go.sum                   # Go dependency lock file
├── Dockerfile               # Dockerfile for building the application
├── docker-compose.yml       # Docker Compose configuration
└── .env                     # Environment variables (ignored in version control)
```

## Notes
- Ensure that Docker and docker-compose are installed if using Docker.
- The .env file is excluded from version control using .gitignore.

## Future Enhancements
- Add authentication using JWT.
- Enhance test coverage with integration tests.
- Add deployment scripts for cloud platforms (e.g., AWS, GCP).

## Feel free to reach out if you face any issues running the application!
