# movies-api

Technical assesment for backend developer possition at KenTech, S.L.U.

# project structure

/movies-api
│── /config                         # Project configuration
│── /docs                           # Swagger API documentation
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
│── main.go                         # Application entry point
│── go.mod                          # Go module
│── go.sum                          # Dependencies
