# A sample project about implementing Domain-Driven Design and Hexagonal Architecture (Ports&Adapters) with Go!

This code is the basis for a series of blog posts:

[Implementing Domain-Driven Design and Hexagonal Architecture with Go - Part 1](https://medium.com/@TonyBologni/implementing-domain-driven-design-and-hexagonal-architecture-with-go-1-292938c0a4d4)

[Implementing Domain-Driven Design and Hexagonal Architecture with Go - Part 2](https://medium.com/@TonyBologni/implementing-domain-driven-design-and-hexagonal-architecture-with-go-2-efd432505554)

Part 3 (about Hexagonal Architecture) is ongoing

The code those blog posts are about is frozen in this branch: [freeze_blog-posts_1-3](https://github.com/AntonStoeckl/go-iddd/tree/freeze_blog-posts_1-3)

### Setup for local development

#### Start Docker container(s)

Run `docker-compose up -d` in the project root.

#### Environment configuration

##### To be able to start the service

Create local.env file in the project root (.env files is gitignored there) with following contents and replace
$PathToProjectRoot$ with the path to the root of the go-iddd project sources.

```
POSTGRES_DSN=postgresql://goiddd:password123@localhost:15432/goiddd_local?sslmode=disable
POSTGRES_MIGRATIONS_PATH_CUSTOMER=$PathToProjectRoot$/go-iddd/service/customeraccounts/infrastructure/postgres/database/migrations
GRPC_HOST_AND_PORT=localhost:5566
REST_HOST_AND_PORT=localhost:8085
REST_GRPC_DIAL_TIMEOUT=3
SWAGGER_FILE_PATH_CUSTOMER=$PathToProjectRoot$/go-iddd/src/customeraccounts/infrastructure/adapter/rest
```

##### To be able to run the tests

Create test.env file in the project root (.env files is gitignored there) with following contents and replace
$PathToProjectRoot$ with the path to the root of the go-iddd project sources.

```
POSTGRES_DSN=postgresql://goiddd:password123@localhost:15432/goiddd_test?sslmode=disable
POSTGRES_MIGRATIONS_PATH_CUSTOMER=$PathToProjectRoot$/go-iddd/src/customeraccounts/infrastructure/adapter/postgres/database/migrations
GRPC_HOST_AND_PORT=localhost:5566
REST_HOST_AND_PORT=localhost:8085
REST_GRPC_DIAL_TIMEOUT=3
SWAGGER_FILE_PATH_CUSTOMER=$PathToProjectRoot$/go-iddd/src/customeraccounts/infrastructure/adapter/rest
```

##### To run HTTP requests with GoLand's (IntelliJ) new built-in HTTP client

Create a customer.http file in the project root (.http files are gitignored there) with following contents.

```
### Register a Customer
POST http://localhost:8085/v1/customer
Accept: */*
Cache-Control: no-cache
Content-Type: application/json

{
  "emailAddress": "john@doe.com",
  "familyName": "Doe",
  "givenName": "John"
}

> {% client.global.set("id", response.body.id); %}

### Confirm a Customer's email address
PUT http://localhost:8085/v1/customer/{{id}}/emailaddress/confirm
Accept: */*
Cache-Control: no-cache
Content-Type: application/json

{
  "confirmationHash": "0acf14bbeaf0b9c6ef8e39d7f9254336"
}

### Change a Customer's email address
PUT http://localhost:8085/v1/customer/{{id}}/emailaddress
Accept: */*
Cache-Control: no-cache
Content-Type: application/json

{
  "emailAddress": "john+changed@doe.com"
}

### Change a Customer's name
PUT http://localhost:8085/v1/customer/{{id}}/name
Accept: application/json
Cache-Control: no-cache
Content-Type: application/json

{
  "givenName": "Joana",
  "familyName": "Doe"
}

### Delete a Customer
DELETE http://localhost:8085/v1/customer/{{id}}
Accept: application/json
Cache-Control: no-cache
Content-Type: application/json

### Retrieve a Customer View
GET http://localhost:8085/v1/customer/{{id}}
Accept: application/json
Cache-Control: no-cache
Content-Type: application/json

### Get the Swagger documentation
GET http://localhost:8085/v1/customer/swagger.json

###
```

**Attention**

The *ConfirmEmailAddress* request does not work without changes - the *confirmationHash* needs to be adapted.
You can find it in the *CustomerRegistered* event in the eventstore DB table.
For security reasons the response of the *Register* request does not return the hash (it **must** only be sent to the Customer via email ;-)

#### Start the service (gRPC and REST)

##### Via Terminal

1) Source the local.env file in your terminal, e.g. `source dev/local.env` or set the env vars in a different way
2) In the project root run `go run service/cmd/grpc/main.go`

##### Via GoLand

1) Create a build configuration for `service/cmd/grpc/main.go`
2) I suggest using the [EnvFile](https://plugins.jetbrains.com/plugin/7861-envfile) GoLand plugin
and add the local.env file in the build configuration
