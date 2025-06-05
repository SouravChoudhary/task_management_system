# Inter-Service Communication
If we add another service (e.g., User Service), here’s how communication can be established:

1. REST (Representational State Transfer) – Simple and Common
Services expose RESTful HTTP APIs.
Example: Task Service can call GET http://user-service/api/users/{id} to get user info.
Easy to implement using standard HTTP clients (like Go's net/http).

2. gRPC – High-Performance Binary Protocol
Strongly typed, faster than REST, ideal for internal microservices.
Example: Define GetUser RPC in user.proto and call via gRPC stub from Task Service.

1. Message Queues (Asynchronous Communication) – Decoupled and Scalable
Use a message broker like Kafka, RabbitMQ.
Services publish events (e.g., TaskCreated) and other services (like a Notification Service) subscribe to them.
Enables eventual consistency and reduces direct coupling.
Example:
Task Service → Publishes task.created event to RabbitMQ
Notification Service → Subscribes and sends email when new task is created

Recommendation:

For simple CRUD microservices --> REST.
For high-performance internal calls --> gRPC.
For event-driven workflows --> message queues.

# Scalability: Horizontal Scaling
We can scale our services horizontally by adding more instances of task-management service. This allows us to handle increased traffic and load without affecting the overall system’s performance.

we can also have read replicas  to offload read queries from the primary database instance. This can improve the overall system’s performance and reduce the load on the primary database instance.