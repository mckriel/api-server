# Multi-API Database Demo

A comprehensive service demonstrating different API paradigms (REST, SOAP, gRPC, GraphQL, WebSockets, WebRTC, and Webhooks) accessing shared data across multiple database types.

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                    Multi-API Service (Go)                   │
├─────────────────────────────────────────────────────────────┤
│ REST    │ SOAP    │ gRPC    │ GraphQL │ WebSockets │ WebRTC │
├─────────────────────────────────────────────────────────────┤
│                   Shared Service Layer                      │
├─────────────────────────────────────────────────────────────┤
│  MongoDB        │     Redis Cache     │     MySQL          │
│  (Documents)    │     (Sessions)      │     (Relations)    │
└─────────────────────────────────────────────────────────────┘
```

## Project Structure

```
multi-api-demo/
├── cmd/
│   └── server/
│       └── main.go           # Main server entry point
├── internal/
│   ├── api/
│   │   ├── rest/             # REST endpoints (Gin/Gorilla)
│   │   ├── soap/             # SOAP service
│   │   ├── grpc/             # gRPC service & protobuf
│   │   ├── graphql/          # GraphQL schema & resolvers
│   │   ├── websocket/        # WebSocket handlers
│   │   ├── webhook/          # Webhook handlers
│   │   └── webrtc/           # WebRTC signaling
│   ├── service/              # Business logic layer
│   ├── repository/           # Database access layer
│   │   ├── mysql/
│   │   ├── mongodb/
│   │   └── redis/
│   └── models/               # Shared data models
├── proto/                    # Protocol buffer definitions
├── schema/                   # Database schemas & migrations
├── docker-compose.yml        # Database services
└── README.md
```

## Database Strategy

### MySQL (Relational Data)
```sql
-- Normalized relational data
users: id, name, email, created_at
products: id, name, price, category_id  
orders: id, user_id, total, status, created_at
order_items: order_id, product_id, quantity, price
categories: id, name
```

### MongoDB (Document Store)
```json
// Rich nested user profiles
{
  "_id": "user_123",
  "profile": {
    "name": "John Doe",
    "email": "john@example.com",
    "preferences": {
      "theme": "dark",
      "notifications": true
    },
    "activity_log": [
      {"action": "login", "timestamp": "2024-01-01T10:00:00Z"},
      {"action": "view_product", "product_id": "456", "timestamp": "..."}
    ],
    "addresses": [
      {"type": "home", "street": "123 Main St", "city": "..."}
    ]
  }
}
```

### Redis (Cache & Real-time Data)
```
// Session management & caching
user:123:session → {"token": "...", "expires": "..."}
product:456:views → 1250
active_users → ["123", "456", "789"]
orders:real_time → pub/sub channel for live updates
user:123:cart → {"items": [...], "total": 99.99}
```

## API Demonstrations

### Same Data, Different Access Patterns:

**Get User Profile:**
- **REST**: `GET /api/v1/users/123`
- **SOAP**: `<getUserProfile><userId>123</userId></getUserProfile>`  
- **gRPC**: `GetUser(UserRequest{id: 123})`
- **GraphQL**: 
  ```graphql
  query {
    user(id: 123) {
      name
      email
      orders {
        total
        status
      }
    }
  }
  ```

**Real-time Features:**
- **WebSockets**: Live order status updates, chat
- **Webhooks**: Notify external systems when orders complete
- **WebRTC**: Real-time video customer support

## Go Dependencies

```go
// go.mod
module multi-api-demo

go 1.21

require (
    // HTTP/REST
    github.com/gin-gonic/gin v1.9.1
    github.com/gorilla/mux v1.8.0
    
    // SOAP
    github.com/hooklift/gowsdl v0.5.0
    
    // gRPC
    google.golang.org/grpc v1.59.0
    google.golang.org/protobuf v1.31.0
    
    // GraphQL
    github.com/99designs/gqlgen v0.17.40
    
    // WebSockets
    github.com/gorilla/websocket v1.5.0
    
    // WebRTC
    github.com/pion/webrtc/v3 v3.2.21
    
    // Databases
    github.com/go-sql-driver/mysql v1.7.1
    github.com/go-redis/redis/v8 v8.11.5
    go.mongodb.org/mongo-driver v1.12.1
    
    // Utilities
    github.com/joho/godotenv v1.4.0
    github.com/sirupsen/logrus v1.9.3
)
```

## Service Ports

```
REST API:     :8080
SOAP:         :8081  
gRPC:         :8082
GraphQL:      :8080/graphql
WebSockets:   :8080/ws
WebRTC:       :8080/webrtc
Webhooks:     :8080/webhook/*
```

## Database Access Patterns

### REST → MySQL
- Direct SQL queries for CRUD operations
- Demonstrates traditional relational data access
- Good for standard business operations

### SOAP → All Databases  
- Complex business logic involving multiple data sources
- Transaction handling across databases
- Enterprise-style data processing

### gRPC → Redis + MySQL
- High-performance cached queries
- Demonstrates efficient data access patterns
- Good for microservice communication

### GraphQL → MongoDB + Joins
- Flexible data fetching with custom queries
- Demonstrates document-based access with relational joins
- Single endpoint, multiple data sources

### WebSockets → Redis Pub/Sub
- Real-time data streaming
- Live updates and notifications
- Session management

### WebRTC → Redis + Real-time signaling
- Peer connection management
- Real-time communication setup

## Docker Setup

```yaml
# docker-compose.yml
version: '3.8'
services:
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: demo
    ports:
      - "3306:3306"
      
  mongodb:
    image: mongo:6.0
    ports:
      - "27017:27017"
      
  redis:
    image: redis:7.0
    ports:
      - "6379:6379"
```

## Getting Started

1. **Setup databases**: `docker-compose up -d`
2. **Install dependencies**: `go mod tidy`
3. **Run migrations**: `go run cmd/migrate/main.go`
4. **Seed data**: `go run cmd/seed/main.go`
5. **Start server**: `go run cmd/server/main.go`

## Learning Objectives

- **API Paradigms**: Compare REST, SOAP, gRPC, GraphQL approaches
- **Database Patterns**: Understand when to use SQL, NoSQL, Cache
- **Real-time Communication**: WebSockets, WebRTC implementation
- **Performance**: Measure and compare different approaches
- **Architecture**: Clean separation of concerns in Go

## Example Use Cases

1. **E-commerce Platform**: Products, Orders, Users across all APIs
2. **User Management**: Authentication, profiles, sessions
3. **Real-time Features**: Live chat, notifications, video calls
4. **Analytics**: View tracking, user behavior, performance metrics
5. **Integration**: Webhooks for external system notifications

This project serves as a comprehensive learning platform for understanding different API architectures and database access patterns in a real-world context.