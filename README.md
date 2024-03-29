# HotelPro REST API

This is a sample REST API for managing hotels and reviews.

## Features

- CRUD operations for hotels and reviews
- User registration and login with JWT authentication
- Swagger documentation for API endpoints

## Installation

### Prerequisites

- Go (v1.13 or later)
- Docker (optional)

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/furkankulaber/hp-hotel-rest.git

2. Navigate to the project directory:
   ```bash
   cd hp-hotel-rest

## Usage

### Running with Docker

You can run the project with Docker. Make sure you have Docker and Docker Compose installed.

1. Build and Start the Docker image:
   ```
   docker-compose up -d

### Running without Docker

If you prefer not to use Docker, you can run the project directly:

1. Install dependencies:
   ```
   go mod tidy
2. Start the application:
   ```
   go run cmd/main.go or air

## API Documentation

You can access the Swagger documentation for the API at http://localhost:8080/swagger/index.html after starting the application.

### Endpoints

* /auth/register: Register a new user.
* /auth/login: Log in an existing user.
* /auth/protected: Access a protected route.


* /hotels: Get all hotels.
* /hotel/{id}: Get a hotel by ID.
* /hotel/{hotelID}/reviews Get reviews by hotel ID
* /hotel/{hotelID}/reviews Add a new review
* /hotel/reviews/{reviewID} Update a review

> [!IMPORTANT]
> Please note that the authentication endpoints (register and login) work independently and do not have any direct connection with the hotel and review endpoints. They are designed solely for user authentication purposes.

### Usage

- Register a new user by making a POST request to /auth/register.
- in an existing user by making a POST request to /auth/login.
- Use the received JWT token in the Authorization header to access protected routes.

### Technologies Used

* Go
* Fiber
* GORM
* JWT


### Testing

To run tests, execute the following command:
   ```
go test ./internal/test/
