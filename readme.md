# Email Sequence Management API

A production-ready API for managing email sequences, designed with modularity, testability, and adherence to Go best practices.

---

## **Features**

### 1. Comprehensive Documentation
- OpenAPI specification available at `/api/v1/docs/swagger-ui/index.html`.
- Visiting the root route (`/`) provides a complete guideline with critical links to documentation and key endpoints.
- The API is self-explanatory, enabling navigation and usage with minimal prior knowledge.

### 2. Metrics
real time metrics are a vital part of a production-ready API, hence I have integrated Prometheus metrics system. 
- Prometheus metrics exposed at `/metrics`.
- Provides real-time monitoring of application performance and database activity. 

### 3. Testability
All modules (services, repositories, workers) implement interfaces for easy mocking.
Comprehensive unit tests covering business logic, API routes, and database interactions.

### 4. Database Transactions
Multi-step operations, such as sequence creation, use transactions to ensure atomicity and consistency.

### 5. Standard Routes
- `/api/v1/health`: Health check for monitoring service availability.
- `/api/v1/info`: Gives information about the API usage information and some metrics specially for maintainers.
- `/`: Root route with links to documentation and critical API endpoints.

### 6. Architecture
Decoupled layers for API, services, and database interactions.
Dependency injection ensures clean, testable implementations.

### 7. Deployment
Fully containerized with Docker Compose.


---

## **Quick Start**

### 1. Run Services
```bash
docker-compose up --build
```
### 2. Running the tests
```bash
go test -v ./...
```

