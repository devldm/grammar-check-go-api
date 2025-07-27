# GrammarCheck Go API

This repository provides the backend API for the GrammarCheck frontend application. It is designed to manage grammar challenges, user accounts, and user solutions, interacting with a PostgreSQL database and exposing a RESTful interface.

## Overview

The GrammarCheck Go API delivers structured grammar challenges to users and tracks their solutions, enabling the frontend to offer an interactive experience. The API `was` containerized and hosted on Google Cloud Run, with builds managed by Google Cloud Build to ensure continuous delivery from repository updates.

## Features

- **API for grammar challenges**: Create, retrieve, and manage grammar exercises.
- **User management**: Register and request user information via Clerk authentication.
- **Solution tracking**: Submit, retrieve, and delete user solutions for grammar challenges.
- **Health check endpoint**: For service monitoring and uptime checks.
- **Cloud-native deployment**: Seamlessly runs on Google Cloud Run using Docker containers.

## Tech Stack

- **Go** (Golang)
- **go-chi** (HTTP router)
- **Docker** (containerization)
- **PostgreSQL** (database)
- **SQLC & Goose** (database access and migrations)
- **Google Cloud Run & Google Cloud Build** (hosting and CI/CD)

## API Routes

### Grammars

```http
GET    /grammars                # Returns all grammar challenges
GET    /grammars/{grammarId}    # Retrieve a grammar challenge by ID
POST   /grammars                # Create a new grammar challenge
```

### Users

```http
GET    /user/{clerk-id}         # Retrieve a user by Clerk ID
POST   /user                    # Register a new user
```

### Solutions

```http
GET    /solutions                           # Get all solutions (with optional limit)
GET    /solutions/user/{clerkUserId}        # Get all solutions for a user
GET    /solutions/{clerkUserId}/{grammarId} # Check if a user has completed a grammar
GET    /solutions/{grammarId}               # Get all solutions for a grammar (with user data)
POST   /solutions                          # Submit a new solution
POST   /solutions/delete                   # Delete a solution
```

### Health

```http
GET    /health   # Service health check
```

## Deployment

- The API is built with Docker and can be deployed using Google Cloud Build.
- Production deployment `was` managed by Google Cloud Run for scalability and reliability.

## Development

1. Clone the repository.
2. Install Go and Docker.
3. Run `docker build . -t grammar-check-api` to build the container.
4. Run locally with Docker or directly using Go for development.
