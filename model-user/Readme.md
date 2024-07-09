# User Management API
## Overview
This project implements a simple User Management API using Go and the Gorilla Mux package. The API allows you to perform CRUD (Create, Read, Update, Delete) operations on user data.

## Features
- Get Users: Retrieve a list of all users.
- Get User: Retrieve a specific user by ID.
- Create User: Add a new user.
- Update User: Update an existing user by ID.
- Delete User: Remove a user by ID.

## Installation
- Clone the repository:

```bash
git clone https://github.com/anjaliBaditya/cli-tools
cd cli-tools/model-user
```
- Install dependencies:
```bash 
go get -u github.com/gorilla/mux
```

- Build and run the application:
```bash 
go build -o user-api
./user-api
```

## Usage
The API runs on localhost:8080 and provides the following endpoints:

- Get Users
Retrieve a list of all users.

- Endpoint: GET /users
Response: JSON array of users.

```bash 
curl -X GET http://localhost:8080/users
```

Get User
Retrieve a specific user by ID.

Endpoint: GET /users/{id}
Response: JSON object of the user.
Example:
```bash
curl -X GET http://localhost:8080/users/1
```
### Create User
Add a new user.

Endpoint: POST /users
Request Body: JSON object containing name and email.
Example:
```bash
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john.doe@example.com"}'
```

### Update User
Update an existing user by ID.

Endpoint: PUT /users/{id}
Request Body: JSON object containing name and email.
Example:
```bash
curl -X PUT http://localhost:8080/users/1 -H "Content-Type: application/json" -d '{"name": "Jane Doe", "email": "jane.doe@example.com"}'
```

### Delete User
Remove a user by ID.

Endpoint: DELETE /users/{id}
Example:
```bash 
curl -X DELETE http://localhost:8080/users/1
```

## Code Structure
- main.go: The main file containing the API endpoints and their implementations.
- User struct: Represents a user with ID, Name, and Email fields.
- Global variables:
- users: A map to store user data.
- nextID: An integer to track the next user ID.
- mu: A mutex to handle concurrent access to user data.
