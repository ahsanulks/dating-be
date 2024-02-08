#  Dating BE

## Structure
In this project I'm using hexagonal architecture / ports and adapters pattern that focuses on creating loosely coupled application components. It emphasizes the separation of concerns and the independence of the application core from external concerns such as  databases, framework, and other external systems.

### Benefits
1. **Modularity**: Modular design by isolating core business logic from external dependencies. This enhances code maintainability and facilitates easier updates and modifications.
   
2. **Testability**: Components within the hexagonal architecture can be easily tested in isolation. Since business logic is decoupled from external dependencies, unit tests can be written without the need for complex mocking frameworks.
   
3. **Flexibility**: Easy replacement or upgrading of external components without impacting the core application logic. 
   
4. **Adaptability**: Supports multiple interfaces and adapters, enabling the application to interact with various external systems or user interfaces without modifying the core logic. 

### Drawbacks
1. **Complexity**: Introduce additional complexity, especially in larger projects, due to the need to define and manage interfaces between core components and external systems.
   
2. **Learning Curve**: Developers unfamiliar with Hexagonal Architecture may face a learning curve in understanding the concepts and best practices associated with this pattern.
   
4. **Performance Overhead**: The use of multiple layers and interfaces in Hexagonal Architecture can introduce a performance overhead compared to non using interface and layer.

### Diagram
```
driven         +-------------------------+    driver
adapter      /        Application          \  adapter
            /       (business logic)        \
    +--------------+                   +--------------+
    |  Driven Port |                   | Driver Port |
    +--------------+                   +--------------+
            \                               /  
             \                             /
              +--------------------------+
```
### Components
#### Driven Adapter:
- **Handler/API:** Responsible for handling incoming requests and translating them into application-specific commands or queries.
- **Extension:** It's possible to incorporate additional UI/presenter layers into this package.

#### Application (Internal):
All components inside the `internal` folder, organized by domain.
- **Entity:** Houses all entity data and business logic.
- **Param:** Contains input and output parameters for the application logic.
- **Port:** Defines interfaces that serve as contracts between driven and driver adapters.
- **Usecase:** Implements business logic that works with multiple entities.
- **Adapter:** Provides fake adapters for testing driver ports, enabling testing within the application layer without coupling with external technology.
- **Custom Error:** Wrapper for business logic errors.

#### Driver Adapter:
- **Infra:** Houses infrastructure-related components required for interaction with external systems or frameworks.

#### External Dependencies
Components outside the application core are related to external driven adapters or frameworks.

## Requirement
### Dev run
This project can several way to run: 
#### 1. Native
To run this project we need to have following installed:
1. [Go](https://golang.org/doc/install) version 1.21
2. [Docker](https://docs.docker.com/get-docker/)
3. [Docker Compose](https://docs.docker.com/compose/install/)
4. [GNU Make](https://www.gnu.org/software/make/)

#### 2. Dev Container
1. [Docker](https://docs.docker.com/get-docker/)
2. [Devcontainer](https://github.com/devcontainers/cli)
    - a. if using vscode, it can more simple using [devcontainer extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

### Other package needed
```make init```

### Build in tools
- [Goose](https://github.com/pressly/goose)
- [Kratos](https://github.com/go-kratos/kratos)
- [Air](https://github.com/cosmtrek/air)
- [OpenApi codegen](https://github.com/deepmap/oapi-codegen)

### migraiton
create migration
```
goose create <name_file> sql
```
run migration
```
goose up
goose down
```
### Run application
```
kratos run
```
or run with hot reload
```
air
```
### Run with docker
```
docker-compose up --build
```

## Access swagger
### Swagger UI
```
http://localhost:8000/q/swagger-ui
```
### OpenApi
```
http://localhost:8000/q/openapi.yaml
```

## Development Flow
### Create API Contract
```
kratos proto add api/{data}
```
### Generate request, openapi, swagger
```
make api
```
### Generate dependency injection
in this repository using [wire](https://github.com/google/wire)
```
make generate
```
### Test
it's needed to run database, when we want to run tests. because we run both unit test and integration test
```
make test
```
### Check Lint
```
make lint
```
