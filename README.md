# Backoffice for a textile manufacturer and retailer

## Employee, product and clients administration

---

## 🧱 Project Architecture

This project follows a layered architecture to promote scalability, maintainability, and clear separation of concerns.

## 📁 Folder Structure by Module

Each domain or module (e.g., User) is organized into its own directory under the internal/ folder. Inside each module, the code is separated into the following layers:

handler/ – Handles HTTP requests and responses (e.g., UserHandler)

service/ – Contains business logic (e.g., UserService)

repository/ – Manages database interactions (e.g., UserRepository)

model/ – Defines data structures and DTOs (e.g., User, UserDTO)

This structure ensures that all functionality related to a feature is self-contained and easy to maintain or scale.

```go
internal/
└── user/
    ├── handler/
    │   └── user_handler.go
    ├── service/
    │   └── user_service.go
    ├── repository/
    │   └── user_repository.go
    └── model/
        ├── user.go
        └── user_dto.go
```

## 🧩 Other Key Components

`cmd/` – Application entry point (e.g., main.go)

`deployments/` – Infrastructure and deployment configurations (e.g., Dockerfile, Terraform)

`internal/config/` – Application configuration

`internal/util/` – Shared utility functions
