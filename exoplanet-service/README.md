# Exoplanet Service

## Description

This microservice provides functionalities to manage exoplanets including CRUD operations and fuel estimation for trips to the exoplanets.

## How to Run

### Using Docker

1. Build the Docker image:
    ```sh
    docker build -t exoplanet-service .
    ```

### Using Go

1. Run the application:
    ```sh
    go run main.go
    ```

### Endpoints

- `POST /exoplanets`: Create a new exoplanet.
- `GET /exoplanets`: Retrieve all ex
