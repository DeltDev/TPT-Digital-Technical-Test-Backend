# Product Management Mini App Backend
Backend REST API for a simple product management app for a small online shop.
## Table of Contents
 
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [How to Run](#how-to-run)
  - [Requirements](#requirements)
  - [Steps](#steps)
- [Other Commands](#other-commands)
- [Default Database Credentials](#default-database-credentials)
- [API Documentation](#api-documentation)
  - [Endpoint List](#endpoint-list)
  - [Get All Products](#get-all-products)
  - [Get Product by ID](#get-product-by-id)
  - [Create Product](#create-product)
  - [Update Product](#update-product)
  - [Delete Product](#delete-product)
---
## Tech Stack

- Go
- Gin
- PostgreSQL
- Docker
## API Documentation
## Endpoint List
 
- [Get All Products](#get-all-products)
- [Get Product by ID](#get-product-by-id)
- [Create Product](#create-product)
- [Update Product](#update-product)
- [Delete Product](#delete-product)
Retrieves a list of all products.
---
## Get All Products
**Endpoint**
 
```
GET /products
```
 
**Request**
 
No request body or query parameters required.
 
**Response**
 
### 200 OK
 
```json
{
  "data": [
    {
      "id": 1,
      "name": "Laptop",
      "description": "A powerful laptop",
      "price": 15000000,
      "stock": 10,
      "category": "Electronics",
      "is_active": true,
      "created_at": "2026-01-01T00:00:00Z",
      "updated_at": "2026-01-01T00:00:00Z"
    },
    {
      "id": 2,
      "name": "Mouse",
      "description": null,
      "price": 250000,
      "stock": 50,
      "category": "Electronics",
      "is_active": false,
      "created_at": "2026-01-02T00:00:00Z",
      "updated_at": "2026-01-03T00:00:00Z"
    }
  ]
}
```
 
### 500 Internal Server Error
 
```json
{
  "error": "failed to query products"
}
```
 
---

## Get Product by ID
 
Retrieves a single product by its ID.
 
**Endpoint**
 
```
GET /products/:id
```
 
**Path Parameters**
 
| Parameter | Type    | Required | Description       |
|-----------|---------|----------|-------------------|
| `id`      | integer | Yes      | The product's ID  |
 
**Response**
 
### 200 OK
 
```json
{
  "data": {
    "id": 1,
    "name": "Laptop",
    "description": "A powerful laptop",
    "price": 15000000,
    "stock": 10,
    "category": "Electronics",
    "is_active": true,
    "created_at": "2026-01-01T00:00:00Z",
    "updated_at": "2026-01-01T00:00:00Z"
  }
}
```
 
### 400 Bad Request
 
Returned when `id` is not a valid integer.
 
```json
{
  "error": "invalid id"
}
```
 
### 404 Not Found
 
```json
{
  "error": "product not found"
}
```

---
 
## Create Product
 
Creates a new product.
 
**Endpoint**
 
```
POST /products
```
 
**Request Body**
 
`Content-Type: application/json`
 
| Field         | Type    | Required | Description                              |
|---------------|---------|----------|------------------------------------------|
| `name`        | string  | Yes      | Name of the product                      |
| `description` | string  | No       | Description of the product, can be null  |
| `price`       | integer | Yes      | Price of the product in smallest unit    |
| `stock`       | integer | Yes      | Available stock quantity                 |
| `category`    | string  | Yes      | Product category                         |
| `is_active`   | boolean | No       | Whether the product is active (default: `true`) |
 
```json
{
  "name": "Keyboard",
  "description": "Mechanical keyboard",
  "price": 800000,
  "stock": 30,
  "category": "Electronics",
  "is_active": true
}
```
 
**Response**
 
### 201 Created
 
```json
{
  "data": {
    "id": 3,
    "name": "Keyboard",
    "description": "Mechanical keyboard",
    "price": 800000,
    "stock": 30,
    "category": "Electronics",
    "is_active": true,
    "created_at": "2026-05-07T10:00:00Z",
    "updated_at": "2026-05-07T10:00:00Z"
  }
}
```
 
### 400 Bad Request
 
Returned when required fields are missing or have invalid types.
 
```json
{
  "error": "Key: 'CreateProductRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
}
```
 
### 500 Internal Server Error
 
```json
{
  "error": "failed to insert product"
}
```
 
---


 
## Update Product
 
Updates an existing product by its ID. All body fields are replaced.
 
**Endpoint**
 
```
PUT /products/:id
```
 
**Path Parameters**
 
| Parameter | Type    | Required | Description       |
|-----------|---------|----------|-------------------|
| `id`      | integer | Yes      | The product's ID  |
 
**Request Body**
 
`Content-Type: application/json`
 
| Field         | Type    | Required | Description                              |
|---------------|---------|----------|------------------------------------------|
| `name`        | string  | Yes      | Name of the product                      |
| `description` | string  | No       | Description of the product, can be null  |
| `price`       | integer | Yes      | Price of the product in smallest unit    |
| `stock`       | integer | Yes      | Available stock quantity                 |
| `category`    | string  | Yes      | Product category                         |
| `is_active`   | boolean | No       | Whether the product is active (default: `true`) |
 
```json
{
  "name": "Keyboard",
  "description": "Mechanical keyboard - updated",
  "price": 850000,
  "stock": 25,
  "category": "Electronics",
  "is_active": false
}
```
 
**Response**
 
### 200 OK
 
```json
{
  "data": {
    "id": 3,
    "name": "Keyboard",
    "description": "Mechanical keyboard - updated",
    "price": 850000,
    "stock": 25,
    "category": "Electronics",
    "is_active": false,
    "created_at": "2026-05-07T10:00:00Z",
    "updated_at": "2026-05-07T11:00:00Z"
  }
}
```
 
### 400 Bad Request (invalid path parameter)
 
```json
{
  "error": "invalid id"
}
```
 
### 400 Bad Request (invalid request body)
 
```json
{
  "error": "Key: 'UpdateProductRequest.Price' Error:Field validation for 'Price' failed on the 'required' tag"
}
```
 
### 404 Not Found
 
```json
{
  "error": "product not found"
}
```
 
### 500 Internal Server Error
 
```json
{
  "error": "failed to update product"
}
```
 
---
 
## Delete Product
 
Deletes a product by its ID.
 
**Endpoint**
 
```
DELETE /products/:id
```
 
**Path Parameters**
 
| Parameter | Type    | Required | Description       |
|-----------|---------|----------|-------------------|
| `id`      | integer | Yes      | The product's ID  |
 
**Response**
 
### 204 No Content
 
Product was successfully deleted. Response body is empty.
 
### 400 Bad Request
 
Returned when `id` is not a valid integer.
 
```json
{
  "error": "invalid id"
}
```
 
### 404 Not Found
 
```json
{
  "error": "product not found"
}
```
 
### 500 Internal Server Error
 
```json
{
  "error": "internal server error"
}
```
 
---

## Default Database Credentials

| Variable | Value |
|---|---|
| Database | tpttechnicaltest |
| Username | tpttechnicaltest |
| Password | tpttechnicaltest |
| Port | 5432 |

Note: The database is automatically initialized with sample seed data during container startup.

## Project Structure

```
cmd/server
docker/
db/
internal/
  db/
  dto/
  handler/
  model/
  repository/
```
## How to Run

### Requirements
1. Docker Engine (Linux) or Docker for Desktop (Windows and Mac)
2. Go v1.26.2 or later

### Steps
1. Copy the ```.env.example``` file and rename it into ```.env```
2. Build PostgreSQL image
```bash
docker build -t tpt-postgres -f docker/Dockerfile .
```
3. Run PostgreSQL image
```bash
docker run -d --name tpt-db -p 5432:5432 tpt-postgres
```
4. Run Backend Server

```bash
go run cmd/server/main.go
```

## Other Commands
- Access PostgreSQL inside docker
```bash
docker exec -it tpt-db psql -U tpttechnicaltest -d tpttechnicaltest
```

- Run unit test for all endpoints
```bash
go test -v ./internal/handler/...
```