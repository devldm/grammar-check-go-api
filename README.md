# GrammerCheck Go API
This repository serves the GrammarCheck frontend application with data from a postgres database. 

## Tech Stack
- Go
- go-chi
- Docker
- postgreSQL
- SQLC & Goose

## Routes
As this is still underdevlopment this will be updated and 
### Grammars
```bash
GET "/grammars" # Returns all grammars

POST "/grammars" # Create new grammar challenge

GET "/grammars/{grammarId}" # Get one grammar by Id
```
### Users
```bash
POST "/user" # Create user

GET "/user/clerk-id" # Get user by clerk-id
```
### Solutions
```bash
POST "/solutions" # Create new solution

GET "/solutions/{clerkUserId}" # Get user submitted solutions
```
