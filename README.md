
# Project Title

This project is a simple backend API that manages movie data. It follows a clean architecture pattern and uses Golang as the main programming language, GORM for database interaction, and integrates various utilities like logging, security, and API documentation with Swagger.


## Appendix

Technical assesment for backend developer possition at KenTech, S.L.U.


## Authors

- [@Lezard82](https://www.github.com/Lezard82)


## Features

- Manage movie information (CRUD operations)
- Swagger API documentation for easy testing and interaction
- GORM ORM for database interactions
- Secure API with middlewares for authentication and validation
- Log management with a singleton logger


## Project Structure

    /movies-api
    │── /config                             # Project configuration
    │── /docs                               # Swagger API documentation
    |── /src
    │   |── /infrastructure
    │   |   |── / api
    |   |   |   |── / dto                   # Data Transfer Objects
    |   |   |   |── / handler               # Handlers
    |   |   |   |── / helpers               # Helpers
    |   |   |   |── / middleware            # Middlewares
    |   |   |   |── / router                # Routes
    |   |   |   |── server.go               # API entry point
    │   |   ├── / db
    |   |   |   |── / models                # GORM database models
    |   |   |   |── / database.go           # API database interface
    |   |   |   |── / gorm.connection.go    # GORM database connection
    |   |   |   |── / gorm.database.go      # GORM database implementation
    │   |   ├── / logger                    # Logger Singleton
    │   |   ├── / repository                # Infrastructure repositories
    │   |   ├── / security                  # Infrastructure security libraries
    │   |   ├── / utils                     # Infrastructure utility libraries
    |   │── /internal
    |   │   ├── /domain                     # Domain entities
    │   |   ├── /repository                 # Domain repositories
    │   |   ├── /usecase                    # Domain use cases, business logic
    │── /tests
    │── main.go                             # Application entry point
    │── go.mod                              # Go module
    │── go.sum                              # Dependencies

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/Lezard82/movies-api.git
    cd movies-api
    ```

2. Install the dependencies:

    ```bash
    go mod tidy
    ```

3. Set up your environment variables by copying the `.env.example` file to `.env`:

    ```bash
    cp .env.example .env
    ```

4. Build and start the project using Docker Compose:

    ```bash
    docker-compose up --build
    ```

This will start the application and a MySQL database container. The MySQL database will be accessible on port 3306.

5. Run the application:

    ```bash
    go run main.go
    ```

The API will be available at `http://localhost:8080`.
    
## Docker Configuration

The project includes a `docker-compose.yml` file to easily manage the database and application environment.

## Database Migrations

The project includes a migrations folder with a migrate.go script. After installing the database, execute this script:

    ```bash
    go run ./migrations/migrate.go
    ```

## API Documentation

The API is documented with Swagger. After running the application, you can access the Swagger UI by navigating to:

http://localhost:8080/docs

## Testing

To run tests, use the following command:

    ```bash
    go test ./...
    ```
## License

[MIT](https://choosealicense.com/licenses/mit/)

