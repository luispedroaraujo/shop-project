# shop-project
A simple Golang-based API that allows users to fetch and filter products from a SQLite3 database. It supports filtering by `category` and `priceLessThan` query parameters.

## Endpoints
- **GET** http://localhost:8080/products
    - Retrieves a list of products.
    - Optional query parameters:
        - `category`: Filter products by category (e.g., `boots`, `sandals`, `sneakers`).
        - `priceLessThan`: Filter products with a price less than or equal to the specified value.
    - Example: http://localhost:8080/products?category=boots&priceLessThan=71000

### Sample Response
```json
[
    {
        "sku": "000003",
        "name": "Ashlington leather ankle boots",
        "category": "boots",
        "price": {
            "original": 71000,
            "final": 49700,
            "discount_percentage": "30%",
            "currency": "EUR"
        }
    }
]
```

## Folder structure
```
- config
    - database.go  # Database connection and configuration
- handlers
    - product_handler.go  # Handlers for product-related API endpoints
    - product_handler_test.go  # Tests for product handlers
- models
    - product.go  # Product model definition
    - product_test.go  # Tests for product model
    - testData.go  # Sample data for testing
- routes
    - product_routes.go  # Routes for product endpoints
- utils
    - utils.go  # Utility functions
```

## Technologies used
- Golang
- GORM with SQLite3
- Gin Web Framework
- Docker & Docker Compose

## How to run the API
Assuming you have docker and docker-compose installed:

`docker-compose up --build -d`

### Run tests
To run tests while the Docker container is running, execute:

`docker-compose exec app go test ./... -v`

## Decisions made
- GORM with SQLite3 was chosen for simplicity in development and local testing. SQLite is lightweight, easy to set up, and doesn't require a dedicated server, making it ideal for rapid development. Also allows for easy transition to other relational databases (e.g., PostgreSQL, MySQL) in the future, as GORM abstracts the underlying database interactions
- Using Gin because it is a fast, lightweight, and easy-to-use web framework that fits well with the performance needs of the API
- Decided to use a products.json file to load products in the database if the database is empty

## Improvements and next steps
- Implement discounts rules dynamically using a discount entity
- Implement a cache to store the requested products with discounts applied
- Implement pagination to return remaining matching products
- Use a robust RDBMS instead of SQLite3
- Use Swagger to generate API documentation