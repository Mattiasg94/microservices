> **⚠️ This project is not yet finished. ⚠️**

# What it demostrates right now:
- Ports and Adapters (Hexagonal Architecture)
- Database access with sqlc for type-safe SQL queries and managed migrations

### Proto Files & gRPC Communication
- Loose Coupling: Keeping API contracts separate allows each microservice to evolve independently.
- Independent Versioning: Versioning proto files separately means changes can be managed without breaking services.
- Automated Generation: GitHub Actions auto-generate Go stubs on tag pushes, keeping code in sync.
- Scalable Architecture: Clean separation between the external APIs and our internal domain logic.

See more details at [microservices-proto](https://github.com/Mattiasg94/microservices-proto)


# TODO
### This is what I plan to do next:
- e2e test with docker compose, unit and integration tests
- Tracing with jaeger and opentelemetry
- Improved logging
- Automatic database migrations
- Local kubernetes setup
- Skaffold for improved local development experience
