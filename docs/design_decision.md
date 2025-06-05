
##  Design Decisions

### Microservices-Oriented Architecture
- Built as an **independent service** that communicates over HTTP.
- Designed to be plugged into a larger ecosystem of services.

### Layered Structure (Clean Architecture Inspired)
- **Handler Layer**: Deals with HTTP requests and responses.
- **Service Layer**: Business logic is encapsulated here.
- **Repository Layer**: Abstracts database operations via interfaces, making it easy to test and swap implementations (e.g., Postgres, Redis, etc.).

###  DTO-Based Request Validation
- Request payloads are validated using custom logic within `Validate()` methods in DTO structs.
- Prevents invalid data from reaching business logic.

###  Structured Error Handling
- All errors are returned as structured JSON responses with meaningful messages.
- Utilizes specific HTTP status codes (e.g., `422 Unprocessable Entity`, `404 Not Found`, etc.).

###  Configurable and Environment-Aware
- Uses a `config.yaml` file for database credentials, server ports, and other environment settings.
- Supports separation of development, staging, and production configurations.

###  Pagination and Filtering
- Query parameters like `page`, `limit`, and `status` are supported to enhance frontend performance and scalability.

### Containerization Ready
- Fully Docker-compatible to ensure easy deployment and consistent environments across dev, test, and production.