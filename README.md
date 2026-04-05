# Kitchen Microservice

A **Golang microservice** example using **gRPC** for inter-service communication, designed with scalability, modularity, and clean separation of concerns in mind.

---

## Architecture Overview

The project follows a **microservice architecture** with multiple layers:

- **gRPC Server**: Listens on a port and registers handlers  
- **Handler Layer**: Implements gRPC methods, converts requests, and calls the service layer  
- **Service Layer**: Contains business logic independent of transport  
- **Generated gRPC Code**: Auto-generated from `.proto` files, defining service interfaces and message types  

---

## Features

- Microservice-based design with independent services  
- gRPC-based communication for high performance and type safety  
- Dependency injection for flexibility and testability  
- Clean separation of transport, business logic, and data layers  

---

## Microservice Principles

- Service-to-service communication using gRPC  
- Clear separation of concerns  
- Scalable and maintainable architecture  
- Designed to support retries, timeouts, circuit breakers, and service discovery  

---

## Folder Structure

- `protobuf/` - Proto definitions  
- `services/orders/` - Orders microservice  
  - `grpc.go` - gRPC server bootstrap  
  - `handler/` - gRPC handlers  
  - `service/` - Business logic  
- `services/common/genproto/` - Generated protobuf code  

---

## Next Steps

- Add additional services (Kitchen, Inventory)  
- Implement retries, timeouts, and circuit breakers  
- Integrate service discovery  
- Add observability (logging, metrics, tracing)  

---

## License

MIT © Nithesh KV  
