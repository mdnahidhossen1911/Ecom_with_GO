# 🛒 Ecom_with_GO - সম্পূর্ণ টেকনিক্যাল ডকুমেন্টেশন

## 📋 সূচিপত্র
1. [প্রজেক্ট ওভারভিউ](#প্রজেক্ট-ওভারভিউ)
2. [আর্কিটেকচার বিশ্লেষণ](#আর্কিটেকচার-বিশ্লেষণ)
3. [ডিরেক্টরি স্ট্রাকচার](#ডিরেক্টরি-স্ট্রাকচার)
4. [ডাটাবেস স্কিমা](#ডাটাবেস-স্কিমা)
5. [API এন্ডপয়েন্টস](#api-এন্ডপয়েন্টস)
6. [ফাইল বাই ফাইল বিস্তারিত বিশ্লেষণ](#ফাইল-বাই-ফাইল-বিস্তারিত-বিশ্লেষণ)
7. [সিকিউরিটি ফিচার্স](#সিকিউরিটি-ফিচার্স)
8. [মিডলওয়্যার সিস্টেম](#মিডলওয়্যার-সিস্টেম)
9. [ডেটা ফ্লো](#ডেটা-ফ্লো)
10. [ডিপেন্ডেন্সি ম্যানেজমেন্ট](#ডিপেন্ডেন্সি-ম্যানেজমেন্ট)

---

## 🎯 প্রজেক্ট ওভারভিউ

### প্রজেক্ট বিবরণ
**নাম:** Ecom_with_GO  
**ভাষা:** Go (Golang)  
**ডাটাবেস:** PostgreSQL  
**আর্কিটেকচার:** Clean Architecture / Hexagonal Architecture  
**API টাইপ:** RESTful API  
**Authentication:** JWT (JSON Web Token)  

### টেকনোলজি স্ট্যাক

| Category | Technology | Purpose |
|----------|-----------|---------|
| Backend Language | Go 1.x | Main programming language |
| Web Framework | net/http (Standard Library) | HTTP server & routing |
| Database | PostgreSQL | Data persistence |
| ORM/Query Builder | sqlx | Database operations |
| Migration Tool | sql-migrate | Database schema management |
| Environment Config | godotenv | Environment variable management |
| Password Hashing | crypto/md5 | Password security (⚠️ deprecated) |
| JWT | Custom Implementation | Authentication |

### প্রজেক্ট ফিচার্স

✅ **User Management**
- User Registration
- User Login with JWT
- Password Hashing
- User CRUD Operations (partial)

✅ **Product Management**
- Create Product (authenticated)
- List Products with Pagination
- Get Single Product
- Update Product (authenticated)
- Delete Product (authenticated)
- Product Count

✅ **Security Features**
- JWT Authentication
- Rate Limiting (6 req/sec per IP)
- IP Blocking (5 min on abuse)
- CORS Support
- Password Hashing

✅ **Advanced Features**
- Database Migration System
- Middleware Chain Management
- Concurrent Data Fetching (Goroutines)
- Request Logging
- Pagination Support
- Error Handling

---

## 🏗️ আর্কিটেকচার বিশ্লেষণ

### Clean Architecture Layers

```
┌─────────────────────────────────────────────────────────────┐
│                    PRESENTATION LAYER                        │
│  (rest/handlers/*, rest/middleware/*, rest/server.go)       │
│  - HTTP Request/Response Handling                           │
│  - JSON Serialization/Deserialization                       │
│  - Middleware Execution                                     │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                    BUSINESS LOGIC LAYER                      │
│              (user/*, product/*)                            │
│  - Business Rules & Validation                              │
│  - Domain Logic                                             │
│  - Service Interfaces                                       │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                    DATA ACCESS LAYER                         │
│                     (repo/*)                                │
│  - Database Operations (CRUD)                               │
│  - SQL Query Execution                                      │
│  - Data Mapping                                             │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                    INFRASTRUCTURE LAYER                      │
│              (infra/db/*, config/*)                         │
│  - Database Connection                                      │
│  - Configuration Management                                 │
│  - External Dependencies                                    │
└─────────────────────────────────────────────────────────────┘
```

### Dependency Flow (Dependency Inversion Principle)

```
main.go
  └─> cmd/serve.go (Dependency Injection Hub)
        ├─> config.GetConfig()
        ├─> db.NewDBConnection()
        ├─> db.MigrateDB()
        ├─> repo.NewUserRepo(dbCon)
        ├─> repo.NewProductRepo(dbCon)
        ├─> user.NewService(userRepo)
        ├─> product.NewService(productRepo)
        ├─> userHandler.NewHandler(config, usrService)
        ├─> productHandler.NewHandler(middleware, prdService)
        └─> rest.NewServer(config, handlers...)
              └─> server.Start() → HTTP Listen :8000
```

### Request Flow Diagram

```
┌───────────┐
│  Client   │
└─────┬─────┘
      │ HTTP Request
      ▼
┌─────────────────────────────────────────┐
│         Middleware Chain                │
│  1. CORS                                │
│  2. Preflight                           │
│  3. Logger                              │
│  4. RateLimit                           │
│  5. AuthenticateJWT (conditional)       │
└─────┬───────────────────────────────────┘
      │
      ▼
┌─────────────────────────────────────────┐
│         HTTP Handler                    │
│  - Parse Request                        │
│  - Validate Input                       │
│  - Call Service Layer                   │
└─────┬───────────────────────────────────┘
      │
      ▼
┌─────────────────────────────────────────┐
│         Service Layer                   │
│  - Business Logic                       │
│  - Call Repository                      │
└─────┬───────────────────────────────────┘
      │
      ▼
┌─────────────────────────────────────────┐
│         Repository Layer                │
│  - SQL Query Execution                  │
│  - Database Operations                  │
└─────┬───────────────────────────────────┘
      │
      ▼
┌─────────────────────────────────────────┐
│         PostgreSQL Database             │
└─────────────────────────────────────────┘
      │
      │ Return Data
      ▼
┌─────────────────────────────────────────┐
│         JSON Response                   │
└─────────────────────────────────────────┘
      │
      ▼
┌───────────┐
│  Client   │
└───────────┘
```

---

## 📁 ডিরেক্টরি স্ট্রাকচার

```
Ecom_with_GO/
│
├── .env                          # Environment variables
├── .idea/                        # IDE configuration (IntelliJ/GoLand)
├── .vscode/                      # VS Code configuration
├── main.go                       # Application entry point
├── go.mod                        # Go module dependencies
├── go.sum                        # Dependency checksums
│
├── cmd/                          # Command layer
│   └── serve.go                  # Server initialization & DI
│
├── config/                       # Configuration management
│   ├── config.go                 # Main config loader
│   └── db_config.go              # Database config loader
│
├── db_quarys/                    # Database query examples
│   └── create_users_table.sql    # SQL example
│
├── domain/                       # Domain models (Entities)
│   ├── user.go                   # User entity
│   └── product.go                # Product entity
│
├── infra/                        # Infrastructure layer
│   └── db/
│       ├── connection.go         # DB connection setup
│       └── migrate.go            # DB migration runner
│
├── migrations/                   # Database migrations
│   ├── 000001-create-users.up.sql
│   ├── 000001-create-users.down.sql
│   ├── 000002-create-products.up.sql
│   └── 000002-create-products.down.sql
│
├── user/                         # User service layer
│   ├── port.go                   # Service interface
│   ├── service.go                # Service implementation
│   ├── create.go                 # Create user logic
│   ├── find.go                   # Find user logic
│   ├─��� get.go                    # Get user logic (unimplemented)
│   ├── list.go                   # List users (unimplemented)
│   ├── update.go                 # Update user (unimplemented)
│   └── delete.go                 # Delete user (unimplemented)
│
├── product/                      # Product service layer
│   ├── port.go                   # Service interface
│   └── service.go                # Service implementation (all methods)
│
├── repo/                         # Repository layer (Data Access)
│   ├── users.go                  # User repository
│   └── product.go                # Product repository
│
├── rest/                         # REST API layer
│   ├── server.go                 # HTTP server setup
│   │
│   ├── handlers/                 # HTTP handlers
│   │   ├── user/
│   │   │   ├── handler.go        # User handler struct
│   │   │   ├── port.go           # Handler service interface
│   │   │   ├── create_user.go    # POST /users
│   │   │   ├── login.go          # POST /users/login
│   │   │   └── router.go         # User routes registration
│   │   │
│   │   └── product/
│   │       ├── handler.go        # Product handler struct
│   │       ├── port.go           # Handler service interface
│   │       ├── create_product.go # POST /products
���   │       ├── get_products.go   # GET /products
│   │       ├── get_product_by_id.go # GET /products/{id}
│   │       ├── update_product.go # PUT /products/{id}
│   │       ├── delete_product.go # DELETE /products/{id}
│   │       └── router.go         # Product routes registration
│   │
│   └── middleware/               # Middleware components
│       ├── manager.go            # Middleware chain manager
│       ├── config_middleware.go  # Config-aware middleware
│       ├── cors.go               # CORS headers
│       ├── preflight.go          # OPTIONS request handler
│       ├── logger.go             # Request logging
│       ├── rate_limit.go         # Rate limiting & IP blocking
│       └── authenticate_jwt.go   # JWT authentication
│
└── util/                         # Utility functions
    ├── create_jwt.go             # JWT token generator
    ├── secPass.go                # Password hashing (MD5)
    ├── sent_json_response.go     # JSON response helpers
    └── send_page.go              # Pagination response helper
```

---

## 🗄️ ডাটাবেস স্কিমা

### Users Table

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(150) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    is_owner BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

**ফিল্ড বিবরণ:**
- `id`: UUID প্রাইমারি কী (auto-generated)
- `name`: ইউজারের নাম (সর্বোচ্চ 150 অক্ষর)
- `email`: ইউজার ইমেইল (unique constraint)
- `password`: হ্যাশকৃত পাসওয়ার্ড (MD5)
- `is_owner`: মালিক কিনা (boolean, default: false)
- `created_at`: একাউন্ট তৈরির সময়
- `updated_at`: শেষ আপডেটের সময়

### Products Table

```sql
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    image_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**ফিল্ড বিবরণ:**
- `id`: UUID প্রাইমারি কী (auto-generated)
- `title`: পণ্যের শিরোনাম (সর্বোচ্চ 255 অক্ষর)
- `description`: পণ্যের বিবরণ (unlimited text)
- `price`: পণ্যের মূল্য (10 ডিজিট, 2 দশমিক)
- `image_url`: পণ্যের ছবির URL
- `created_at`: পণ্য যুক্ত করার সময়
- `updated_at`: শেষ আপডেটের সময়

### ER Diagram

```
┌─────────────────────────────────┐
│           USERS                 │
├─────────────────────────────────┤
│ • id (UUID, PK)                 │
│ • name (VARCHAR)                │
│ • email (VARCHAR, UNIQUE)       │
│ • password (TEXT)               │
│ • is_owner (BOOLEAN)            │
│ • created_at (TIMESTAMPTZ)      │
│ • updated_at (TIMESTAMPTZ)      │
└─────────────────────────────────┘
         │
         │ (Future: orders relationship)
         │
         ���
┌─────────────────────────────────┐
│         PRODUCTS                │
├─────────────────────────────────┤
│ • id (UUID, PK)                 │
│ • title (VARCHAR)               │
│ • description (TEXT)            │
│ • price (NUMERIC)               │
│ • image_url (TEXT)              │
│ • created_at (TIMESTAMPTZ)      │
│ • updated_at (TIMESTAMPTZ)      │
└─────────────────────────────────┘
```

---

## 🌐 API এন্ডপয়েন্টস

### Base URL
```
http://localhost:8000
```

### Authentication
JWT token প্রয়োজন এমন endpoints এর জন্য:
```
Authorization: Bearer <JWT_TOKEN>
```

---

### 👤 User Endpoints

#### 1. Create User (Registration)
**Endpoint:** `POST /users`  
**Authentication:** ❌ Not Required  
**Description:** নতুন ইউজার রেজিস্ট্রেশন

**Request Body:**
```json
{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "mypassword123",
    "is_owner": false
}
```

**Response (200):**
```json
{
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "John Doe",
    "email": "john@example.com",
    "is_owner": false,
    "created_at": "2026-02-10T10:30:00Z",
    "updated_at": "2026-02-10T10:30:00Z"
}
```

**Password Hashing:**
- Password টি MD5 hash করে ডাটাবেসে সংরক্ষিত হয়
- Original password কখনো সংরক্ষিত হয় না

---

#### 2. User Login
**Endpoint:** `POST /users/login`  
**Authentication:** ❌ Not Required  
**Description:** ইউজার লগইন এবং JWT token প্রাপ্তি

**Request Body:**
```json
{
    "email": "john@example.com",
    "password": "mypassword123"
}
```

**Response (200):**
```json
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI1NTBlODQwMC1lMjliLTQxZDQtYTcxNi00NDY2NTU0NDAwMDAiLCJuYW1lIjoiSm9obiBEb2UiLCJlbWFpbCI6ImpvaG5AZXhhbXBsZS5jb20iLCJpc19vd25lciI6ZmFsc2UsImNyZWF0ZWRfYXQiOiIyMDI2LTAyLTEwVDEwOjMwOjAwWiIsInVwZGF0ZWRfYXQiOiIyMDI2LTAyLTEwVDEwOjMwOjAwWiJ9.signature"
```

**JWT Token Structure:**
```
Header.Payload.Signature
```

**Payload Contents:**
```json
{
    "sub": "user_id",
    "name": "user_name",
    "email": "user_email",
    "is_owner": true/false,
    "created_at": "timestamp",
    "updated_at": "timestamp"
}
```

**Error Response (404):**
```json
"Invalid Credential: user not found"
```

---

### 📦 Product Endpoints

#### 1. List Products (with Pagination)
**Endpoint:** `GET /products?page=1&limit=10`  
**Authentication:** ❌ Not Required  
**Description:** পণ্য তালিকা pagination সহ

**Query Parameters:**
- `page` (optional): পৃ��্ঠা নম্বর (default: 1)
- `limit` (optional): পৃষ্ঠা প্রতি আইটেম (default: 10)

**Response (200):**
```json
{
    "data": [
        {
            "id": "660e8400-e29b-41d4-a716-446655440000",
            "title": "iPhone 15 Pro",
            "description": "Latest Apple smartphone",
            "price": 999.99,
            "image_url": "https://example.com/iphone15.jpg",
            "created_at": "2026-02-01T10:00:00Z",
            "updated_at": "2026-02-01T10:00:00Z"
        },
        {
            "id": "770e8400-e29b-41d4-a716-446655440001",
            "title": "MacBook Pro M3",
            "description": "Powerful laptop",
            "price": 2499.99,
            "image_url": "https://example.com/macbook.jpg",
            "created_at": "2026-02-02T11:00:00Z",
            "updated_at": "2026-02-02T11:00:00Z"
        }
    ],
    "pagination": {
        "page": 1,
        "limit": 10,
        "total_item": 25,
        "total_page": 3
    }
}
```

**Implementation Note:**
- Concurrent data fetching using Goroutines
- Two goroutines run parallel:
  1. Fetch product list
  2. Fetch total count
- Results merged via channels

---

#### 2. Get Single Product
**Endpoint:** `GET /products/{productID}`  
**Authentication:** ❌ Not Required  
**Description:** নির্দিষ্ট পণ্যের বিস্তারিত তথ্য

**Response (200):**
```json
{
    "id": "660e8400-e29b-41d4-a716-446655440000",
    "title": "iPhone 15 Pro",
    "description": "Latest Apple smartphone with A17 Pro chip",
    "price": 999.99,
    "image_url": "https://example.com/iphone15.jpg",
    "created_at": "2026-02-01T10:00:00Z",
    "updated_at": "2026-02-01T10:00:00Z"
}
```

**Error Response (404):**
```json
"Product not found"
```

---

#### 3. Create Product
**Endpoint:** `POST /products`  
**Authentication:** ✅ JWT Required  
**Description:** নতুন পণ্য যুক্ত করা

**Request Headers:**
```
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json
```

**Request Body:**
```json
{
    "title": "Samsung Galaxy S24",
    "description": "Latest Android flagship",
    "price": 899.99,
    "image_url": "https://example.com/galaxy-s24.jpg"
}
```

**Response (201):**
```json
{
    "id": "880e8400-e29b-41d4-a716-446655440002",
    "title": "Samsung Galaxy S24",
    "description": "Latest Android flagship",
    "price": 899.99,
    "image_url": "https://example.com/galaxy-s24.jpg",
    "created_at": "2026-02-10T12:00:00Z",
    "updated_at": "2026-02-10T12:00:00Z"
}
```

**Error Response (401):**
```json
"Unauthorized"
```

---

#### 4. Update Product
**Endpoint:** `PUT /products/{productID}`  
**Authentication:** ✅ JWT Required  
**Description:** বিদ্যমান পণ্য আপডেট করা

**Request Headers:**
```
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json
```

**Request Body:**
```json
{
    "title": "iPhone 15 Pro Max",
    "description": "Updated description",
    "price": 1099.99,
    "image_url": "https://example.com/iphone15-pro-max.jpg"
}
```

**Response (201):**
```json
{
    "id": "660e8400-e29b-41d4-a716-446655440000",
    "title": "iPhone 15 Pro Max",
    "description": "Updated description",
    "price": 1099.99,
    "image_url": "https://example.com/iphone15-pro-max.jpg",
    "created_at": "2026-02-01T10:00:00Z",
    "updated_at": "2026-02-10T13:00:00Z"
}
```

**Error Responses:**
- `400`: Bad Request (invalid JSON)
- `401`: Unauthorized (missing/invalid JWT)
- `404`: Product not found
- `500`: Internal server error

---

#### 5. Delete Product
**Endpoint:** `DELETE /products/{productID}`  
**Authentication:** ✅ JWT Required  
**Description:** পণ্য মুছে ফেলা

**Request Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Response (200):**
```json
"Product deleted successfully"
```

**Error Responses:**
- `400`: Product ID required
- `401`: Unauthorized
- `404`: Product not found

---

### Error Response Format

সব endpoints একই error format অনুসরণ করে:

**Client Error (4xx):**
```json
"Error message describing the problem"
```

**Server Error (5xx):**
```json
"Internal server error message"
```

---

### Rate Limiting

**Rules:**
- Maximum 6 requests per second per IP
- Exceeded requests → 429 Too Many Requests
- IP blocked for 5 minutes on abuse

**Rate Limit Response (429):**
```json
"Too many requests! IP blocked for 5 minutes."
```

---

## 📄 ফাইল বাই ফাইল বিস্তারিত বিশ্লেষণ

### 🔹 Root Level Files

#### main.go
```go
package main

import "ecom_project/cmd"

func main() {
	cmd.Serve()
}
```

**বিস্তারিত:**
- **উদ্দেশ্য:** Application এর entry point
- **কাজ:** শুধুমাত্র `cmd.Serve()` function কল করে
- **ডিজাইন প্যাটার্ন:** Separation of concerns - main থেকে initialization logic আলাদা
- **Execution Flow:** 
  1. Program শুরু হয় main() থেকে
  2. cmd.Serve() কল হয়
  3. Serve() সব initialization করে server চালু করে

---

#### .env
```env
VERSION=1.0.0
SERVICE_NAME=ecom_project
PORT=8000
JWT_SECURE_KEY=Mdn@hid1911

DB_HOST=127.0.0.1
DB_PORT=5432    
DB_USER=postgres
DB_PASSWORD=2003
DB_NAME=ecom
DB_SSLMODE=disable
```

**বিস্তারিত:**
- **উদ্দেশ্য:** Environment-specific configuration সংরক্ষণ
- **Security Note:** Production এ এই ফাইল .gitignore এ থাকা উচিত
- **Configuration Categories:**
  - **Application Config:** VERSION, SERVICE_NAME, PORT
  - **Security Config:** JWT_SECURE_KEY
  - **Database Config:** All DB_* variables

**Best Practices:**
- Different .env files for dev/staging/production
- Never commit sensitive credentials to git
- Use environment variable managers (e.g., AWS Secrets Manager) in production

---

### 🔹 cmd/serve.go

```go
package cmd

import (
	"ecom_project/config"
	"ecom_project/infra/db"
	"ecom_project/product"
	"ecom_project/repo"
	"ecom_project/rest"
	productHandler "ecom_project/rest/handlers/product"
	userHandler "ecom_project/rest/handlers/user"
	"ecom_project/rest/middleware"
	"ecom_project/user"
	"fmt"
	"os"
)

func Serve() {
	// 1. Load Configuration
	config := config.GetConfig()

	// 2. Database Connection
	dbCon, err := db.NewDBConnection(config.DBConfig)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		os.Exit(1)
		return
	}

	// 3. Database Migration
	error := db.MigrateDB(dbCon, "./migrations")
	if error != nil {
		fmt.Println("Failed to migrate the database:", error)
		os.Exit(1)
		return
	}

	// 4. Repository Layer Initialization
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	// 5. Service Layer Initialization
	usrService := user.NewService(userRepo)
	prdService := product.NewService(productRepo)
	middleware := middleware.NewConfigMiddleware(config)

	// 6. Handler Layer Initialization
	productsHandler := productHandler.NewHandler(middleware, prdService)
	userHandler := userHandler.NewHandler(config, usrService)

	// 7. HTTP Server Initialization & Start
	server := rest.NewServer(config, productsHandler, userHandler)
	server.Start()
}
```

**বিস্তারিত:**

**1. Configuration Loading**
```go
config := config.GetConfig()
```
- Singleton pattern ব্যবহার করে config load
- First call এ .env থেকে পড়ে
- Subsequent calls একই instance return করে
- Validation সহ load (missing values = os.Exit(1))

**2. Database Connection**
```go
dbCon, err := db.NewDBConnection(config.DBConfig)
```
- PostgreSQL এ connection establish করে
- Connection string তৈরি করে config থেকে
- sqlx.DB instance return করে
- Connection failure = graceful shutdown

**3. Database Migration**
```go
error := db.MigrateDB(dbCon, "./migrations")
```
- migrations/ folder এর SQL files execute করে
- Up migrations run করে (table creation/updates)
- Version tracking for idempotency
- Migration failure = application exit

**4. Dependency Injection - Repository Layer**
```go
userRepo := repo.NewUserRepo(dbCon)
productRepo := repo.NewProductRepo(dbCon)
```
- Database connection inject করা হয় repositories তে
- Each repository gets its own instance
- Follows Repository Pattern

**5. Dependency Injection - Service Layer**
```go
usrService := user.NewService(userRepo)
prdService := product.NewService(productRepo)
```
- Repository instances inject করা হয় services এ
- Business logic layer তৈরি
- Decoupled from data access details

**6. Dependency Injection - Handler Layer**
```go
productsHandler := productHandler.NewHandler(middleware, prdService)
userHandler := userHandler.NewHandler(config, usrService)
```
- Service instances inject করা হয় handlers এ
- HTTP request handling layer তৈরি
- Middleware configuration included

**7. Server Start**
```go
server := rest.NewServer(config, productsHandler, userHandler)
server.Start()
```
- HTTP server তৈরি
- Routes register করা
- Port 8000 তে listen শুরু
- Blocking call - application runs until stopped

**ডিজাইন প্যাটার্ন:**
- ✅ Dependency Injection
- ✅ Constructor Injection
- ✅ Dependency Inversion Principle
- ✅ Single Responsibility (each layer has one job)

**Error Handling Strategy:**
- Critical errors → os.Exit(1)
- Prevents partial initialization
- Fail-fast approach

---

### 🔹 config/config.go

```go
package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

var configuration *Config

type Config struct {
	Version      string
	ServiceName  string
	Port         string
	JwtSecureKey string
	DBConfig     DBConfig
}

func loadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}

	// Load VERSION
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION not set in environment")
		os.Exit(1)
	}

	// Load SERVICE_NAME
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("SERVICE_NAME not set in environment")
		os.Exit(1)
	}

	// Load PORT
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT not set in environment")
		os.Exit(1)
	}

	// Validate PORT is numeric
	if _, err := fmt.Sscanf(port, "%d", new(int)); err != nil {
		fmt.Println("PORT must be a valid number")
		os.Exit(1)
	}

	// Load JWT_SECURE_KEY
	jwtSecurekey := os.Getenv("JWT_SECURE_KEY")
	if jwtSecurekey == "" {
		fmt.Println("JWT Secure Key not set in environment")
		os.Exit(1)
	}

	fmt.Printf("Configuration loaded: \nVersion=%s,\nServiceName=%s,\nPort=%s\n", 
		version, serviceName, port)

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		Port:         port,
		JwtSecureKey: jwtSecurekey,
	}

	loadDBConfig()
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
```

**বিস্তারিত:**

**Singleton Pattern Implementation:**
```go
var configuration *Config  // Package-level variable

func GetConfig() *Config {
	if configuration == nil {  // Lazy initialization
		loadConfig()
	}
	return configuration
}
```
- একবার মাত্র load হয়
- Memory efficient
- Thread-safe না (simple use case এর জন্য OK)

**Validation Strategy:**
```go
if version == "" {
	fmt.Println("VERSION not set in environment")
	os.Exit(1)
}
```
- প্রতিটি required field check করা হয়
- Missing value = immediate exit
- Prevents runtime errors

**Port Validation:**
```go
if _, err := fmt.Sscanf(port, "%d", new(int)); err != nil {
	fmt.Println("PORT must be a valid number")
	os.Exit(1)
}
```
- Port numeric কিনা verify করা
- Invalid port = application won't start
- Prevents bind errors

**Configuration Print:**
```go
fmt.Printf("Configuration loaded: \nVersion=%s,\nServiceName=%s,\nPort=%s\n", 
	version, serviceName, port)
```
- Startup time এ configuration দেখানো
- Debugging এর জন্য helpful
- Production এ sensitive data print করা উচিত না

---

### 🔹 config/db_config.go

```go
package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func loadDBConfig() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB_HOST not set in environment")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("DB_PORT not set in environment")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB_USER not set in environment")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB_PASSWORD not set in environment")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB_NAME not set in environment")
		os.Exit(1)
	}

	dbSSLMode := os.Getenv("DB_SSLMODE")
	if dbSSLMode == "" {
		fmt.Println("DB_SSLMODE not set in environment")
		os.Exit(1)
	}

	configuration.DBConfig = DBConfig{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		DBName:   dbName,
		SSLMode:  dbSSLMode,
	}
}
```

**বিস্তারিত:**

**Database Configuration Structure:**
```go
type DBConfig struct {
	Host     string  // Database server address (e.g., localhost, IP, domain)
	Port     string  // Database port (PostgreSQL default: 5432)
	User     string  // Database username
	Password string  // Database password
	DBName   string  // Database name to connect
	SSLMode  string  // SSL connection mode (disable, require, verify-ca, verify-full)
}
```

**SSL Mode Options:**
- `disable`: No SSL (development only)
- `require`: SSL required but don't verify
- `verify-ca`: Verify Certificate Authority
- `verify-full`: Full SSL verification (production recommended)

**Security Considerations:**
- Password plain text তে load হয় (acceptable for config loading)
- Never log password values
- Use environment variables or secret managers
- .env file should be in .gitignore

**Configuration Validation:**
- প্রতিটি field required
- Missing any field = application exit
- Fail-fast approach prevents connection errors

---

### 🔹 domain/user.go

```go
package domain

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	IsOwner   bool      `json:"is_owner" db:"is_owner"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
```

**বিস্তারিত:**

**Struct Tags Explanation:**

**JSON Tags:**
```go
`json:"id"`      // Serialization name in JSON
`json:"-"`       // Exclude from JSON (password security)
```

**DB Tags:**
```go
`db:"id"`        // Column name in database
`db:"password"`  // Maps to password column
```

**Field-by-Field Analysis:**

1. **ID (UUID)**
   - Type: string (UUID stored as string)
   - Database: UUID type
   - JSON: "id"
   - Purpose: Unique identifier

2. **Name**
   - Type: string
   - Max length: 150 chars (DB constraint)
   - Required: Yes
   - Purpose: User's full name

3. **Email**
   - Type: string
   - Max length: 255 chars
   - Unique: Yes (DB constraint)
   - Required: Yes
   - Purpose: Login credential & contact

4. **Password**
   - Type: string
   - JSON: `-` (excluded from JSON responses)
   - Hashed: MD5 (⚠️ should use bcrypt)
   - Purpose: Authentication

5. **IsOwner**
   - Type: bool
   - Default: false
   - Purpose: Role-based access control
   - Future: Can determine admin privileges

6. **CreatedAt**
   - Type: time.Time
   - Auto-set: Database default NOW()
   - Not in JSON: Can be added if needed
   - Purpose: Audit trail

7. **UpdatedAt**
   - Type: time.Time
   - Auto-updated: ON UPDATE NOW()
   - Not in JSON: Can be added if needed
   - Purpose: Audit trail

**Design Patterns:**
- ✅ Domain-Driven Design (DDD)
- ✅ Rich Domain Model
- ✅ Separation of Concerns (domain logic separate)

**Security:**
- Password excluded from JSON with `json:"-"`
- Never exposed in API responses
- Hashed before storage

---

### 🔹 domain/product.go

```go
package domain

type Product struct {
	ID          string  `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImageURL    string  `json:"image_url" db:"image_url"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
	UpdatedAt   string  `json:"updated_at" db:"updated_at"`
}
```

**বিস্তারিত:**

**Field-by-Field Analysis:**

1. **ID (UUID)**
   - Type: string
   - Generated: Database (uuid_generate_v4())
   - Immutable: Never changes
   - Purpose: Unique product identifier

2. **Title**
   - Type: string
   - Max length: 255 chars
   - Required: Yes
   - Example: "iPhone 15 Pro Max"
   - Purpose: Product name

3. **Description**
   - Type: string
   - Max length: Unlimited (TEXT type)
   - Optional: Can be null
   - Example: "Latest flagship smartphone..."
   - Purpose: Detailed product information

4. **Price**
   - Type: float64
   - Database: NUMERIC(10, 2)
   - Format: 10 digits total, 2 decimal places
   - Example: 999.99
   - Purpose: Product price in currency

5. **ImageURL**
   - Type: string
   - Format: URL string
   - Optional: Can be null
   - Example: "https://cdn.example.com/product.jpg"
   - Purpose: Product image reference

6. **CreatedAt**
   - Type: string (⚠️ should be time.Time)
   - Auto-generated: Database
   - Format: ISO 8601 timestamp
   - Purpose: Track when product was added

7. **UpdatedAt**
   - Type: string (⚠️ should be time.Time)
   - Auto-updated: Database
   - Format: ISO 8601 timestamp
   - Purpose: Track last modification

**Type Inconsistencies:**
```go
// Current (Product)
CreatedAt string
UpdatedAt string

// Better Practice (User)
CreatedAt time.Time
UpdatedAt time.Time
```
- User struct correctly uses time.Time
- Product should also use time.Time
- String conversion should happen at JSON layer

**Price Handling:**
- float64 suitable for display
- For financial calculations, consider decimal library
- Prevents floating-point precision issues

---

### 🔹 infra/db/connection.go

```go
package db

import (
	"ecom_project/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDBConnectionString(config *config.DBConfig) string {
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		config.User, 
		config.Password, 
		config.Host, 
		config.Port, 
		config.DBName, 
		config.SSLMode,
	)
	return connectionString
}

func NewDBConnection(config config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetDBConnectionString(&config)
	
	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
	return db, nil
}
```

**বিস্তারিত:**

**Connection String Format:**
```
user=postgres password=2003 host=127.0.0.1 port=5432 dbname=ecom sslmode=disable
```

**GetDBConnectionString() Analysis:**
- **Purpose:** Config থেকে connection string তৈরি
- **Format:** PostgreSQL standard connection string
- **Security:** Password plain text (connection string এ normal)

**NewDBConnection() Analysis:**

1. **Connection String Creation:**
```go
dbSource := GetDBConnectionString(&config)
```
- Config struct থেকে formatted string তৈরি

2. **Database Connection:**
```go
db, err := sqlx.Connect("postgres", dbSource)
```
- `sqlx.Connect()` - Connection pool তৈরি করে
- `"postgres"` - Driver name (lib/pq)
- `dbSource` - Connection configuration

3. **Error Handling:**
```go
if err != nil {
	fmt.Println(err)
	return nil, err
}
```
- Connection failure print করা
- Error propagate করা caller এ

**sqlx Library Benefits:**
- Named parameters support (`:name` syntax)
- Struct scanning
- Better error messages
- Connection pooling
- Compatible with database/sql

**Connection Pooling:**
- sqlx automatically creates connection pool
- Default settings used
- Can be configured with db.SetMaxOpenConns(), db.SetMaxIdleConns()

**Driver Import:**
```go
_ "github.com/lib/pq"
```
- Blank import
- Registers PostgreSQL driver
- Side-effect only (no direct usage)

---

### 🔹 infra/db/migrate.go

```go
package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, dir string) error {
	migratements := &migrate.FileMigrationSource{
		Dir: dir,
	}
	
	_, err := migrate.Exec(db.DB, "postgres", migratements, migrate.Up)
	if err != nil {
		return err
	}
	
	fmt.Println("Database migration completed successfully")
	return nil
}
```

**বিস্তারিত:**

**Migration System:**

1. **File-Based Migrations:**
```go
migratements := &migrate.FileMigrationSource{
	Dir: dir,  // "./migrations"
}
```
- SQL files থেকে migrations load করে
- Directory structure based
- Version-ordered execution

2. **Migration Execution:**
```go
_, err := migrate.Exec(db.DB, "postgres", migratements, migrate.Up)
```
- `db.DB` - Underlying *sql.DB (from sqlx)
- `"postgres"` - Database dialect
- `migratements` - Migration source
- `migrate.Up` - Direction (up = apply, down = rollback)

**Migration File Naming Convention:**
```
000001-create-users.up.sql      # Up migration
000001-create-users.down.sql    # Down migration (rollback)
000002-create-products.up.sql
000002-create-products.down.sql
```

**Version Tracking:**
- Migration library tracks applied versions
- Creates `gorp_migrations` table
- Prevents duplicate execution
- Idempotent by design

**Migration Flow:**
1. Read all `.up.sql` files from directory
2. Check which versions already applied
3. Execute pending migrations in order
4. Update version tracking table
5. Return success/error

**Error Handling:**
- Migration failure stops execution
- Partial migrations rolled back
- Error propagated to caller
- Application exits on failure (in serve.go)

**Best Practices:**
- Use sequential numbering (000001, 000002, ...)
- Write reversible migrations (both up and down)
- Test migrations in development first
- Keep migrations small and focused

---

### 🔹 migrations/ SQL Files

#### 000001-create-users.up.sql

```sql
-- +migrate Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
     name VARCHAR(150) NOT NULL,
     email VARCHAR(255) UNIQUE NOT NULL,
     password TEXT NOT NULL,
     is_owner BOOLEAN NOT NULL DEFAULT FALSE,
     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

**বিস্তারিত:**

**Migration Directive:**
```sql
-- +migrate Up
```
- Migration library এর special comment
- Marks file as "up" migration
- Tells migrate tool to execute this on upgrade

**UUID Extension:**
```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```
- PostgreSQL UUID functions enable করে
- `uuid_generate_v4()` function available করে
- Idempotent (`IF NOT EXISTS`)

**Table Creation - Users:**

**Primary Key:**
```sql
id UUID PRIMARY KEY DEFAULT uuid_generate_v4()
```
- UUID type (128-bit unique identifier)
- Auto-generated (no need to provide on INSERT)
- Better than SERIAL for distributed systems

**Name Field:**
```sql
name VARCHAR(150) NOT NULL
```
- Variable character field (max 150)
- Required field
- Stores user's full name

**Email Field:**
```sql
email VARCHAR(255) UNIQUE NOT NULL
```
- Max 255 characters
- UNIQUE constraint (index automatically created)
- Required field
- Used for login

**Password Field:**
```sql
password TEXT NOT NULL
```
- Unlimited length (TEXT type)
- Stores hashed password (MD5 in this case)
- Required field

**IsOwner Field:**
```sql
is_owner BOOLEAN NOT NULL DEFAULT FALSE
```
- Boolean type (true/false)
- Default: false (regular user)
- Role-based access control

**Timestamps:**
```sql
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
```
- TIMESTAMPTZ = Timestamp with timezone
- Auto-populated on INSERT
- Audit trail

---

#### 000001-create-users.down.sql

```sql
-- +migrate Down
DROP TABLE IF EXISTS users;
```

**বিস্তারিত:**

**Rollback Migration:**
- Reverses changes made by up migration
- Safe with `IF EXISTS`
- Deletes table and all data

**When Used:**
- Migration rollback
- Development testing
- Fixing migration errors

---

#### 000002-create-products.up.sql

```sql
-- +migrate Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS products (
     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    image_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

**বিস্তারিত:**

**Products Table Structure:**

**Title Field:**
```sql
title VARCHAR(255) NOT NULL
```
- Product name
- Max 255 characters
- Required

**Description Field:**
```sql
description TEXT
```
- Optional field (can be NULL)
- Unlimited length
- Detailed product information

**Price Field:**
```sql
price NUMERIC(10, 2) NOT NULL DEFAULT 0
```
- NUMERIC type (exact decimal)
- Total 10 digits, 2 after decimal
- Example: 12345678.90
- Default: 0
- Better than FLOAT for money

**Image URL Field:**
```sql
image_url TEXT
```
- Optional field
- Stores URL string
- Reference to image file

**Timestamps:**
```sql
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
```
- Same as users table
- Audit trail

---

## 🔐 সিকিউরিটি ফিচার্স

### 1. JWT Authentication

**Implementation Location:** `rest/middleware/authenticate_jwt.go`

**Token Structure:**
```
Header.Payload.Signature
```

**Header:**
```json
{
    "alg": "HS256",
    "typ": "JWT"
}
```

**Payload:**
```json
{
    "sub": "user_id",
    "name": "User Name",
    "email": "user@example.com",
    "is_owner": false,
    "created_at": "2026-02-10T10:00:00Z",
    "updated_at": "2026-02-10T10:00:00Z"
}
```

**Signature Generation:**
```go
message := base64(header) + "." + base64(payload)
signature := HMAC-SHA256(message, JWT_SECURE_KEY)
token := message + "." + base64(signature)
```

**Verification Process:**
1. Extract token from Authorization header
2. Split into parts (header, payload, signature)
3. Recreate signature using secret key
4. Compare signatures
5. Allow/deny request

**Protected Endpoints:**
- POST /products
- PUT /products/{id}
- DELETE /products/{id}

---

### 2. Password Hashing

**Implementation:** `util/secPass.go`

```go
func SecPass(pass string) string {
	hash := md5.Sum([]byte(pass))
	return hex.EncodeToString(hash[:])
}
```

**⚠️ Security Warning:**
- MD5 হ্যাশিং deprecated এবং দুর্বল
- Rainbow table attacks এর শিকার হতে পারে
- Recommended: bcrypt, argon2, scrypt

**Better Implementation (Recommendation):**
```go
import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
```

---

### 3. Rate Limiting

**Implementation:** `rest/middleware/rate_limit.go`

**Configuration:**
```go
const (
	maxReqPerSec   = 6              // Maximum requests per second
	blockDuration  = 5 * time.Minute // Block duration
	windowDuration = 1 * time.Second // Time window
)
```

**Algorithm:**

```
For each incoming request:
  1. Get client IP
  2. Check if IP is blocked
     → If blocked and block time not expired: Reject (429)
  3. Check request count in current time window
     → If window expired: Reset counter
  4. Increment request counter
  5. If counter > maxReqPerSec:
     → Block IP for blockDuration
     → Reject request (429)
  6. Otherwise: Allow request
```

**Data Structure:**
```go
type clientInfo struct {
	requests int       // Request count in current window
	window   time.Time // Window start time
	blocked  time.Time // Block end time
}

var reqData = make(map[string]*clientInfo)
```

**IP Extraction:**
```go
func getIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
```

**Concurrency Safety:**
```go
var mu sync.Mutex  // Protects reqData map

mu.Lock()
// ... access reqData ...
mu.Unlock()
```

**Response:**
```
HTTP/1.1 429 Too Many Requests
Content-Type: text/plain

Too many requests! IP blocked for 5 minutes.
```

---

### 4. CORS (Cross-Origin Resource Sharing)

**Implementation:** `rest/middleware/cors.go`

```go
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		next.ServeHTTP(w, r)
	})
}
```

**Headers Explanation:**

**Content-Type:**
```
Content-Type: application/json
```
- সব response JSON format এ

**Allow-Origin:**
```
Access-Control-Allow-Origin: *
```
- যেকোনো domain থেকে request allow
- Production এ specific domains দেওয়া উচিত
- Example: `Access-Control-Allow-Origin: https://myapp.com`

**Allow-Methods:**
```
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, PATCH, OPTIONS
```
- Allowed HTTP methods
- REST operations support

**Allow-Headers:**
```
Access-Control-Allow-Headers: Content-Type
```
- Client কোন headers পাঠাতে পারবে
- Should include: Authorization (for JWT)

**Preflight Requests:**
```go
func Preflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOPTIONS {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
```
- OPTIONS requests handle করে
- Browser CORS preflight check এর জন্য

---

### 5. Request Logging

**Implementation:** `rest/middleware/logger.go`

```go
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.RemoteAddr
		}

		next.ServeHTTP(w, r)

		fmt.Printf(
			"User IP: %s | Method: %s | URL: %s | Duration: %v\n",
			ip,
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
```

**Log Format:**
```
User IP: 192.168.1.100 | Method: GET | URL: /products | Duration: 15ms
User IP: 127.0.0.1 | Method: POST | URL: /users/login | Duration: 234ms
```

**IP Detection:**
1. Check `X-Forwarded-For` header (proxy/load balancer)
2. Fallback to `RemoteAddr` (direct connection)

**Metrics Collected:**
- Client IP address
- HTTP method
- Request path
- Response time (duration)

**Use Cases:**
- Performance monitoring
- Debugging
- Security auditing
- Traffic analysis

---

## 🔄 মিডলওয়্যার সিস্টেম

### Middleware Chain Architecture

```
HTTP Request
     │
     ▼
┌─────────────────┐
│   CORS          │ ← Set CORS headers
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   Preflight     │ ← Handle OPTIONS requests
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   Logger        │ ← Log request details
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   RateLimit     │ ← Check rate limits
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ AuthenticateJWT │ ← Verify JWT (conditional)
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   Handler       │ ← Execute business logic
└─────────────────┘
         │
         ▼
   HTTP Response
```

### Middleware Manager

**Implementation:** `rest/middleware/manager.go`

```go
type Middleware func(http.Handler) http.Handler

type Manager struct {
	middlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		middlewares: []Middleware{},
	}
}

func (m *Manager) Use(mw Middleware) {
	m.middlewares = append(m.middlewares, mw)
}

func (m *Manager) Apply(handler http.Handler, middleware ...Middleware) http.Handler {
	h := handler
	for i := len(middleware) - 1; i >= 0; i-- {
