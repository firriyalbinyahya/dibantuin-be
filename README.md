# DIBANTUIN - SOCIAL DONATION APP

This system is designed to facilitate online social donation activities, enabling users to provide assistance to groups in need. It also allows social organizations to create donation programs, receive contributions, and manage transactions, with strong security and flexibility as the foundation of a robust and modern backend API.

**Problem**

Many social organizations have difficulty reaching donors/volunteers who are willing to provide financial assistance on a large scale.

**Goal**
Create a backend system that can accommodate aid, record transactions, and make it easier for organizations and admins to manage aid programs efficiently and securely.

## Table of Contents

- [Main Features](#main-features)
- [Technology Used](#technology-used)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributions](#contributions)
- [License](#license)


---

## Main Features

* User Management (Registration, Login, Authentication using JWT)
* CRUD Donation Programs
* CRUD Categories
* Manual Donations
Allows donors to contribute by uploading proof of bank transfer or evidence of in-kind donations.
* Admin Verification
Enables administrators to manually review and approve donation programs and incoming donations. This step ensures the legitimacy of campaigns and the authenticity of transactions
* Audit Trail
Maintains a detailed record of important system activities, including user actions, donation submissions, and verification steps

---

## Technology Used

* **Backend**: Go (Gin), GORM
* **Database**: MySQL
* **Caching**: Redis
* **Authentication**: JWT

---

## Instalation

Follow these steps to run this project on your local machine.

### Prerequisite

Make sure you have [Go](https://golang.org/doc/install) version 1.22 or higher and [Docker](https://docs.docker.com/get-docker/) installed.

### Steps

1.  Clone this repository:
    ```bash
    git clone [https://github.com/feriyalbinyahya/dibantuin-be.git](https://github.com/feriyalbinyahya/dibantuin-be.git)
    cd dibantuin-be
    ```

2.  Create a `.env` file from `.env.example` and fill in the required variables (such as database and Redis credentials).

3.  Run the project using Docker Compose:
    ```bash
    docker-compose up --build
    ```
    Or run it manually:
    ```bash
    go mod tidy
    go run main.go
    ```

---

## API Usage
Here are some cURL examples to interact with the API. Ensure your server is running at http://localhost:8080.

* 1. Example `cURL` for login:
```bash
curl -X POST http://localhost:8080/api/auth/login \
-H "Content-Type: application/json" \
-d '{
    "email": "user@example.com",
    "password": "validpassword"
}'
```

Example of a successful response:
```json
{
  "code": "SUCCESS",
  "message": "Login successful",
  "data": {
    "name": "Test User",
    "email": "test@example.com",
    "role": "user",
    "access_token": "eyJhbGciOiJIUzI1Ni...",
    "access_expired": "2025-08-21T15:00:00Z",
    "refresh_token": "eyJhbGciOiJIUzI1Ni...",
    "refresh_expired": "2025-08-28T15:00:00Z"
  }
}
```

* 2. Get User Details
Use the access token you received from logging in to retrieve a specific user's data. Replace :id with a valid user ID.

```bash
# Replace <ACCESS_TOKEN> with the token you got from the previous step.
curl -X GET http://localhost:8080/api/users/1 \
-H "Authorization: Bearer <ACCESS_TOKEN>"
```

Example of a successful response:
```json
{
  "code": "SUCCESS",
  "message": "User retrieved successfully",
  "data": {
    "id": 1,
    "name": "Test User",
    "email": "test@example.com",
    "role": "user",
    "created_at": "2025-08-21T14:30:00Z"
  }
}
```
* 3. Updating User Data
To update user data, send the access token and the data you want to change in JSON format.

```bash
# Updating only username
curl -X PUT http://localhost:8080/api/users/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
-d '{
    "name": "New Test Name"
}'
```

* 4. Deleting a User
To delete a user account, send a DELETE request with a valid access token. Only the user or an administrator can perform this action.

```bash
# Deleting user account with ID 1
curl -X DELETE http://localhost:8080/api/users/1 \
-H "Authorization: Bearer <ACCESS_TOKEN>"
```

## Project Structure
This project follows a clean, layered architecture to separate concerns.
```
dibantuin-be/
├── .github/             # GitHub workflows for CI/CD
│   └── workflows/
├── config/              # Configuration files (e.g., database, Redis)
├── constants/           # Constants used throughout the application
├── controller/          # Layer for handling HTTP requests and responses
├── entity/              # Database models and DTOs
├── middleware/          # Middleware such as authentication and authorization
├── repository/          # Layer for database interaction
├── routes/              # API route definitions
├── service/             # Layer for business logic
├── uploads/             # For file uploads
├── utils/               # General utility functions
├── .env                 # Environment variables file
├── .gitignore           # List of ignored files for Git
├── Dockerfile           # Docker image definition
├── README.md            # Project overview and documentation
├── docker-compose.yml   # Docker configuration for services
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
└── main.go              # Application entry point
```

