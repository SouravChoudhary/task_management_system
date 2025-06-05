## 🧱 Microservices Architecture Principles Demonstrated

This project is designed with key microservices concepts in mind, even though it's a single service. It can be easily extended and integrated into a broader distributed system.

---

###  1. **Single Responsibility Principle**
- The service focuses solely on **task management**—creating, listing, updating, and deleting tasks.
- It maintains separation of concerns by structuring components cleanly (handlers, services, repositories, DTOs, models).

---

###  2. **Layered Architecture with Dependency Injection**
- The codebase is divided into:
  - `handler` – HTTP layer
  - `service` – Business logic
  - `repository` – Data access logic
- Interfaces decouple implementations, making the service testable and modular.

---

###  3. **Interface-Driven Design**
- All dependencies (e.g., data store) are injected using interfaces.
- Enables easy mocking and substitution for testing and future integrations.

---

###  4. **Containerization**
- The entire service is **Dockerized**.
- Docker Compose is used to manage service dependencies like Postgres.
- Promotes reproducibility and simplifies local development and deployment.

---

###  5. **Configurability**
- Uses external config (`config.yaml`) to separate environment-specific variables (DB credentials, ports, etc.).
- Promotes 12-Factor App principles.

---

###  6. **RESTful API Design**
- All endpoints follow clear and consistent REST patterns.
- JSON is used for request and response payloads.
- Proper HTTP status codes and error messages are returned.

---

###  7. **Scalability-Ready**
- Stateless design allows the service to be horizontally scaled behind a load balancer.
- With proper API gateway and service discovery, it can easily fit into a larger microservices ecosystem.

---

###  8. **Validation and Error Handling**
- Uses centralized input validation in DTOs.
- Provides structured and informative error responses at the API layer.

---

This architecture ensures the service can be independently deployed, tested, and scaled—hallmarks of a microservices-based system.

