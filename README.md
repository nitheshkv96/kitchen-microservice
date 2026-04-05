# kitchen-microservice

A production-ready microservice architecture example in Golang using gRPC for inter-service communication.

This project demonstrates a clean and modular structure for building scalable and maintainable microservices.

🏗️ Architecture Overview

This project implements a microservice architecture with clear separation of concerns:

Kitchen-Microservice/
├── services/
│   ├── orders/         # Orders microservice
│   │   ├── handler/    # gRPC handlers (transport layer)
│   │   ├── service/    # Business logic layer
│   │   └── grpc.go     # gRPC server bootstrap
│   └── common/         # Shared proto definitions and utilities
│       └── genproto/   # Generated gRPC code
├── protobuf/           # .proto files for gRPC definitions
└── README.md
Layers
gRPC Server (grpc.go)
Listens on a specific port
Registers service handlers
Acts as the entry point for incoming requests
Handler (handler/grpc.go)
Implements gRPC RPC methods
Converts gRPC requests into domain objects
Calls the service layer
Service (service/orders.go)
Contains business logic (e.g., creating orders)
Independent of transport (gRPC, HTTP, etc.)
Testable and easily replaceable
Generated gRPC code (genproto)
Auto-generated from .proto files
Defines the service interface and message types
⚡ Features
Microservice-based design
Orders service is independent and can be scaled separately
Can communicate with other services (e.g., Kitchen, Inventory) via gRPC
gRPC communication
Fast, strongly-typed RPC framework
Automatically generates Go client/server code from .proto definitions
Dependency injection
Handlers depend on service interfaces, not concrete implementations
Makes testing and swapping implementations easy
Clean separation of concerns
Transport layer, business logic, and data layer are fully decoupled
🚀 Getting Started
Prerequisites
Golang >= 1.20
protoc compiler for generating gRPC code
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
Install Dependencies
go mod tidy
Generate gRPC Code
make gen

(or run the protoc commands directly on your .proto files)

Run the Orders Service
go run services/orders/grpc.go

This starts the gRPC server on the configured port (e.g., :50051).

📝 Usage
Sample RPC Call
conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
client := pb.NewOrderServiceClient(conn)

res, err := client.CreateOrder(ctx, &pb.CreateOrderRequest{
    Order: &pb.Order{
        OrderID: 42,
        ProductID: 12,
        CustomerID: 1,
        Quantity: 10,
    },
})
The request is handled by OrdersGrpcHandler
Business logic executed in OrderService
Response returned to client
🌐 Microservice & Distributed System Principles
Service-to-service communication: Uses gRPC for fast, strongly-typed RPC calls.
Separation of concerns: Transport layer (handler) is independent of business logic (service).
Scalability: Each microservice can run independently and be horizontally scaled.
Testability: Service interfaces allow mocking and unit testing without a database.
Extensible: Easy to add retries, timeouts, or service discovery mechanisms in production.
📂 Folder Structure
kitchen-microservice/
├── protobuf/               # Proto definitions
├── services/
│   ├── orders/
│   │   ├── grpc.go         # gRPC server entry
│   │   ├── handler/        # gRPC handler methods
│   │   └── service/        # Business logic
│   └── common/genproto/    # Generated protobuf code
├── go.mod
└── README.md
📈 Next Steps
Implement Kitchen service and connect it with Orders via gRPC
Add timeouts, retries, and circuit breakers for production-grade service-to-service communication
Add service discovery (DNS or Consul) for dynamic service resolution
Integrate logging, metrics, and tracing (e.g., OpenTelemetry)
