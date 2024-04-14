## Latest project structure
<pre>
├── README.md
├── cmd
│   └── http
│       └── main.go
├── config
│   ├── config.go
│   └── config.yaml
├── docs
│   └── api.yaml
├── internal
│   ├── app
│   │   └── app.go
│   ├── server
│   │   └── server.go
│   ├── middleware
│   │   ├── jwt.go
│   │   └── logger.go
│   ├── domain
│   │   ├── constant
│   │   ├── entity
│   │   └── model
│   └── somemodule
│       ├── controller
│       │   └── http
│       │       └── somemodule_handler.go
│       ├── repository
│       │   ├── db
│       │   │   └── xxx_repository.go
│       │   └── api
│       │       └── xxx_repository.go
│       └── service
│           ├── somemodule_service.go
│           └── somemodule_service_test.go
├── pkg
│   ├── cloudstorage
│   │   └── cloudstorage.go
│   ├── logger
│   │   └── logger.go
│   ├── postgres
│   │   └── postgres.go
│   └── redis
│       └── redis.go
├── Dockerfile
├── docker-compose.yaml
├── Makefile
├── README.md
├── go.mod
├── go.sum
</pre>