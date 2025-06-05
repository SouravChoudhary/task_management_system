
# Task Management Microservice
# [ Detailed Documentation is available under docs/ directory ]
## Instructions to Run the Service

### Prerequisites
- Docker and Docker Compose installed on your machine
- (Optional) Go installed if you want to run the service locally without Docker

---

### Running with Docker Compose

1. **Clone the repository:**

```bash
git clone https://github.com/SouravChoudhary/task_management_system.git
cd task_management_system
````

2. **Start the service with Docker Compose:**

```bash
docker-compose up --build
```
---

This process will:

* Build and run the `task-service` container
* Start the PostgreSQL database container
* Run the service accessible on port `8080`
* Create tasks table in the database if it doesn't exist

---

### Running Locally (Without Docker)

1. **Set up PostgreSQL database** locally or use a hosted instance.

2. **Update your database credentials** in the `config/config.yaml` file.

3. **Create the database table manually** 
3. **Create the required database table manually:**

Once the PostgreSQL container is running, connect to the database and execute the following SQL to create the `tasks` table:

```sql
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

4. **Run the service:**

```bash
go run ./cmd/main.go
```

5. **Access the API at:** `http://localhost:8080`

---

### Notes

* Ensure the database service is running and accessible before starting the task service.
* When using Docker Compose, the service waits for the database to be ready before starting.
* To stop the running containers via Docker Compose, press `Ctrl+C` and then run:

```bash
docker-compose down
```
