# goBlue-back

Simple Role Based User Administration API using golans fiber with swagger Documentation

## Getting started


```
git clone --branch dev https://gitlab.com/semayerp/goblue-back.git
cd goblue-back
go mod init semay.com
go get
go build main.go
swag init
./main run
```

you will find the swagger docs at  https://localhost:5500/docs/

if swag init is does not work make sure to export the path in you terminal. as you need  the swagger docs to be generated
for the above endpoint to work.

```
export PATH=$PATH:$(go env GOPATH)/bin

```

To start the services using docker compose
```
docker-compose up -d
```

Then you will find the services on
* goblue[](https:goblue.localhost/docs/)
* prometheus[](http:prometheus.localhost)
* grafana[](fibergrafana.localhost)
    * username : admin
    * password : pass@123

Add Datastore in grafana pointing to http:prometheus.localhost 
Import Dashboard JSON file option to Use the profided goblue-Dashboard.json
The Dashboard should be configured to use the datasore you created earlier
NB: make sure you use prometheus version 8.5


## Tools utilized

* traefik proxy server[traefik](https://github.com/traefik/traefik)
* Fiber[Golang Fiber FastHttp framework](https://github.com/gofiber/fiber) 
* Promeetheus[Monitoring your app](https://github.com/prometheus/prometheus) 
* Docker Compose
* Docker CE


## Test and Deploy

Most Tests Are Currently Written in POST MAN
Deploye locally using Docker docker Compose.
upcoming 

# Upcoming Updates
* documentation
* ci/cd with render

## Visual Structure of Project 
* upcoming update

## Serviece Architecture Diagram 
* Upcoming update
## Usage
Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

## License
MIT LICENSE


