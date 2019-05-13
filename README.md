# Golang GRPC with RESTful endpoint (GRPC Gateway) Example

GRPC Gateway: https://github.com/grpc-ecosystem/grpc-gateway
An Example with TLS: https://github.com/philips/grpc-gateway-example

# Quick Start - Go Application
Requires [protobuf](https://github.com/protocolbuffers/protobuf) and [golang-migrate](https://github.com/golang-migrate/migrate).  
`gofmt`, `golint` and `direnv` (to use `.envrc`) recommended.  
In a separate console, first run `docker-compose up` to boot up the accompanying `mysql` database.
```bash
make tools
make compile-protobuf # requires you to install protobuf first
go get ./...
go mod vendor

# set the ENVIRONMENT variable (see below) before running main.go
make run

# alternatively you can run `make run-client` in another console
curl localhost:8080/_ah/health 
```

If you are using this as an example project to work off of, don't forget to run `go mod init` to set up go modules

# Building for Docker
```bash
make compile-protobuf
make compile-binary
docker build -t <THING> --build-arg env=<ENVIRONMENT>
docker run <THING> -p 8080:80808 ...
```

# Defining Services

Services are defined in `/proto`. When compiled, the generated interfaces need to be implemented. In this example, they're implemented in `/services` but how it's structured is completely up to you.

In the protobuf 3 language, it's very simple to define standard HTTP methods. We define the request model and view model to be returned in the proto files.

The normal approach to the services is:
 1. Define the endpoints and models
 2. Run the protobuf compiler (Provided in the Makefile). This will generate boilerplate GRPC code with interfaces that your services must implement. There is also the added benefit of Swagger docs automatically being generated with a protoc plugin.
 3. Wire up the new services in `main.go` (steps 2 and 3 can be run at the same time with `make run`)

There is a complete example with two endpoints. The `health` endpoint is a simple service that simply displays that the server is up.
The `pet` endpoint is a service designed to interact with a database at `ECOMM_CLOUD_DB_CONNECTION_STRING`. It will come with one `GET` and one `POST` request, along with validation and error handling.

# Migrations

Migrations are performed with [golang-migrate](https://github.com/golang-migrate/migrate) and are stored in the `/migrations` directory.

Simple migrations can be created by calling `make migrate-create <name-of-migration>`

Migrating up or down can be performed by calling, for example, `make migrate up` or `make migrate down 3`.

More advanced migrations will have to be run manually.

# Testing

Testing can be performed via the standard `go test` interface (wrapped by `make test`). To test the GRPC server, run `make run-client`.
The code for the client is in the `/client` directory.
Currently, it does not load from the files in the `configs` folder, so you will have to manually set the required variables before using it.
All client interactions must also be hard coded, but extensive testing should be done via `go test` anyway.

There are no current real tests that have been written as it is expected that all the implementations are going to be redone anyway.


# Configurations

Configurations are loaded from the toml files in `/configs`. The behaviour is as follows:
 
 1. .toml file gets loaded first
 2. environment variable will override what's in the toml file
 3. required environment variables do not need to be added to the toml files

# Mandatory Environment Variables
| Variable | Example | Description |
| --- | --- | --- |
| ENVIRONMENT | local | name the config toml files to be the same as the environment variable name.
| DB_CONNECTION_STRING | `root@tcp(localhost:3306)/test?charset=utf8` | URI connection string for connecting to the database  


# Optional Environment Variable Overrides
| Variable | Example | Description |
| --- | --- | --- |
| ENABLE_DEBUGGING | true | Always enable log.debug
| API_PORT | ":8080" | The RESTful endpoint port
| GRPC_PORT | ":8081" | The GRPC endpoint port
| GRPC_HOST | "localhost" | The hostname of the GRPC server
