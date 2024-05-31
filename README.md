# dot Microservices with Go Frontend web server

## Branches
- Start with the `step1-add-templates-start` branch, 
- Then you can check `step1-add-templates-solution`
- Then `step2-connect-broker-start`, `step-2-connect-broker-solution`
- and so on...

## TODO for the step1-add-templates-start branch
- initialize a module (create a `go.mod` file) with `go mod init <module-name>`
for example `go mod init frontend` or `go mod init github.com/<username>/<project>`
- the `cmd/frontend/templates` folder contains html templates. 
- the main entry point is `cmd/frontend/main.go`
- Try to implement the TODO comments in the `main.go` file


## How to start the web server:
Open a terminal from the outer `frontend` folder
- `go run cmd/frontend/main.go`
- if prompted, allow network permissions
- open in a browser `http://localhost`

