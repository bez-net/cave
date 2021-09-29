# Cave


Cave is a golang restful api and a video/audio streaming server for the Adullam bible seminary of the Remnant Christian Network (RCN).

### Features!

  - Used mux for routing
  - Used salt to hash the password
  - Used JWT for authentication
  - Follows repository, service and model structure
  - Used interface to hide implementation of repositories and services
  - Error handling
  
### Project structure
```bash
cave
├── api
│   ├── api.go
│   ├── db
│   │   ├── connection.go
│   │   └── indexes.go
│   ├── handler
│   │   ├── middleware.go
│   │   ├── person.go
│   │   ├── person_test.go
│   │   └── response_writer.go
│   └── model
│       ├── person.go
│       └── response.go
├── config
│   └── config.go
├── docker-compose.yml
├── dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── README.md
└── swagger.json
```
  
### Installation

Cave requires [Go](https://golang.org/) 1.10+ to run.

Install the dependencies and devDependencies and start the server.

```sh
git clone https://github.com/onos9/cave.git
cd cave
go run main.go
```
#### Building for source
For production release:
```sh
go build
```
  ### Todos

  - Email verification
  - Containerized - Docker + Kubernetes
  - Write Tests
  - Write scripts
  
### Tech

Project uses a number of open source projects to work properly:

* [Go] - Awesome programing language by Google
* [mux] - Implements a request router and dispatcher in Go
* [MongoDB] - document-based, big community, database
* [Docker] - Build, Share, and Run Any App, Anywhere
* [Kubernetes] - Automating deployment, scaling, and management of containerized applications


API endpoint - http://localhost:8181
Swagger endpoint - http://localhost:8181/swagger/index.html


   [mux]: <https://www.gorillatoolkit.org/pkg/mux>
   [Go]: <https://golang.org/>
   [MongoDB]: <https://www.mongodb.com/>
   [Docker]: <https://www.docker.com/>
   [Kubernetes]: <https://kubernetes.io/>
