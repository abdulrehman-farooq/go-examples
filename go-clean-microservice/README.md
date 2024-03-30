# Clean Architecture Project

This project follows the principles of Clean Architecture as described by Uncle Bob. The goal of clean architecture is to create systems that are independent of frameworks, databases, and external details, allowing for flexibility, maintainability, and testability.

## Clean Architecture Overview

Clean Architecture consists of several layers, each with distinct responsibilities:

### 1. Entities

Entities are the most abstract layer of the system and represent the core business objects. They encapsulate enterprise-wide business rules and are independent of external concerns.

### 2. Use Cases

Use Cases, also known as Interactors, contain application-specific business rules. They orchestrate the flow of data and operations between the entities and the external layers, such as the controllers or presenters.

### 3. Interface Adapters

Interface Adapters convert data from the format most convenient for the use cases and entities into the format most convenient for external frameworks and tools. This layer includes Presenters, Controllers, Gateways, and Data Mappers.

### 4. Frameworks and Drivers

Frameworks and Drivers are the most external layers and contain all the details related to the UI, databases, web frameworks, etc. They are responsible for communication with external systems and tools.

## Layers in Detail

- **Entities**: Contains business objects and enterprise-wide business rules.
- **Use Cases**: Contains application-specific business rules and orchestrates data flow.
- **Interface Adapters**: Converts data between internal and external formats.
- **Frameworks and Drivers**: Contains details related to UI, databases, and external systems.

## Project Structure 
 .
├── doc
|   ├─api
|   ├─integration-details
├── cmd
│   └── service
|       └── main.go
├── internal
│   ├── domain
│   ├── infrastructure
│   ├── interface
│   ├── usecase
│   └── registery
├── resources
├── test
├── go.mod
├── go.sum
├── Dockerfile
└── docker-compose.yaml


## Reference Links

- [Golang Standards Project Layout](https://github.com/golang-standards/project-layout)
- [The Clean Architecture by Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
