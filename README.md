# dot Microservices with Go Frontend web server

## Branches
- You can find a "ready" version on the `main` branch 
- there a more empty version with comments on the `start` branch for exploring by yourself

## Instructions for the start branch
- initialize a module (create a `go.mod` file) with `go mod init <module-name>`
for example `go mod init frontend` or `go mod init github.com/<username>/<project>`
- the `cmd/frontend/templates` folder contains html templates. Feel free to edit them.
- the main entry point is `cmd/frontend/main.go`



## How to start the web server:
Open a terminal from the outer `frontend` folder
- `go run cmd/frontend/main.go`
- if prompted, allow network permissions
- open in a browser `http://localhost`

