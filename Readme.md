# My Go REST API

## Introduction

This project is a simple REST API built with Go (Golang), featuring basic GET and POST endpoints and integrated Swagger UI for API documentation and testing. It demonstrates handling HTTP requests and in-memory data management.

## Getting Started

### Prerequisites

- Go (version 1.x)
- Git (for cloning the repository)

### Installation

1. **Clone the Repository**

   ```sh
   git clone https://github.com/yourusername/my-go-rest-api.git
   cd my-go-rest-api
   ```

2. **Install Dependencies**

   ```sh
   go mod tidy
   ```

3. **Run the Application**
   ```sh
   go run main.go
   ```

### Accessing the Swagger UI

Once the application is running, you can access the Swagger UI by navigating to `http://localhost:8000/swagger/index.html` in your web browser.

## API Endpoints

### GET /users

- **Description**: Retrieves a list of users.
- **Response**: An array of user objects.
- **Example Request**: `GET http://localhost:8000/users`

### POST /users

- **Description**: Adds a new user.
- **Request Body**:
  ```json
  {
    "id": "string",
    "name": "string"
  }
  ```
- **Response**: The added user object.
- **Example Request**: `POST http://localhost:8000/users`

## Development

### Running Tests

(Include instructions on how to run any tests, if applicable.)

### Building for Production

(Include instructions on how to build the application for production, if applicable.)

## Contributing

Contributions to this project are welcome! Please fork the repository and submit a pull request with your changes.

## License

This project is licensed under the [MIT License](LICENSE).

---
