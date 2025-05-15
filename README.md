# Coupon System MVP

A backend service for managing coupon codes in a medicine ordering platform, built with Go, Gin, and GORM with SQLite database.

---

## Features so far

- Create and manage coupons with attributes like coupon code, expiry date, usage type, applicable medicines/categories, minimum order value, discount details, and usage limits.
- Database schema designed with normalized tables for coupons, medicines, categories, and their relationships.
- API endpoint to create coupons (`POST /coupons`).
- Automatic database migration using GORM’s `AutoMigrate`.
- Basic project structure following Go best practices.

---

## Tech Stack

- **Go** (Golang)  
- **Gin** (HTTP web framework)  
- **GORM** (ORM for Go)  
- **SQLite** (lightweight SQL database)  

---

## Setup Instructions

1. **Clone the repo**

   ```bash
   git clone https://github.com/yourusername/coupon-system.git
   cd coupon-system
````

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Run the app**

   ```bash
   go run ./cmd/main.go
   ```

   The app will start on default port 8080.

4. **Database**

   On start, the app will automatically create (or migrate) the SQLite database file `coupons.db` with the necessary tables.

5. **Create coupon API**

   Use POST `/coupons` with JSON body:

   ```json
   {
       "coupon_code": "SAVE20",
       "expiry_date": "2025-12-31T23:59:59Z",
       "usage_type": "one_time",
       "applicable_medicine_ids": "med_1,med_2",
       "applicable_categories": "painkiller,antibiotic",
       "min_order_value": 500,
       "valid_time_window_start": "2025-05-01T00:00:00Z",
       "valid_time_window_end": "2025-05-31T23:59:59Z",
       "terms_and_conditions": "Valid for May only",
       "discount_type": "percent",
       "discount_value": 20,
       "max_usage_per_user": 1
   }
   ```

---

## Project Structure

```
coupon-system/
│
├── cmd/                  # Entry point (main.go)
├── internal/
│   ├── handlers/              # HTTP handlers
│   ├── service/          # Business logic (planned)
│   ├── repository/       # DB connection and queries
│   ├── models/           # Database models
│   └── utils/            # Utility functions (planned)
├── migrations/           # Database migration files (planned)
├── configs/              # Config files (planned)
├── docs/                 # Swagger/OpenAPI docs (planned)
├── Dockerfile            # Docker setup (planned)
├── go.mod
└── README.md
```

---

## Next Steps

* Implement coupon validation logic.
* Add more API endpoints (e.g., GET applicable coupons, POST validate coupon).
* Add concurrency-safe usage count updates.
* Add caching layer.
* Write Swagger documentation.
* Dockerize the app.
* Deploy to cloud.

---

## Notes on Concurrency and DB Migrations

* Using GORM's `AutoMigrate` for simple database schema management during development.
* Careful locking and transactional updates will be added for concurrent coupon usage handling.

