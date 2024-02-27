# GrammerCheck Go API
This repository serves the GrammarCheck frontend application with data from a postgres database. 

It is hosted on Google Cloud Run (containerised) and uses Google Cloud build to keep up to date with the repository changes.

## Tech Stack
- Go
- go-chi
- Docker
- postgreSQL
- SQLC & Goose
- Google Cloud Run & Google Cloud Build

## Routes
As this is still underdevlopment this will be updated and 
### Grammars
```bash
GET "/grammars" # Returns all grammars

GET "/grammars/{grammarId}" # Get one grammar by Id

POST "/grammars" # Create new grammar challenge
```
### Users
```bash
GET "/user/clerk-id" # Get user by clerk-id

POST "/user" # Create user
```
### Solutions
```bash
GET "/solutions" # Get all solutions

GET "/solutions/user/{clerkUserId}" # Get solutions by user id

GET "/solutions/{clerkUserId}/{grammarId}" # Get users solutions per grammar (check if they have completed a grammar)

GET "/solutions/{grammarId}" # Get solutions by grammar (with user data)

POST "/solutions" # Create new solution

POST "/solutions/delete" # Delete solution
```
### Health
```bash
GET "/health" # Check health
```
