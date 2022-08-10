# UMSystem

## About

API service for user and role management.

## REST API

Below is a table of available REST API endpoints and their description.

name endpoint description request response status codes

| Name             | Endpoint                 | Method | Auth method | Description | Request schema                               | Response schema          | Status codes |     |
|------------------|--------------------------|--------|-------------|-------------|----------------------------------------------|--------------------------|--------------|-----|
| Create user      | `/user`                  | POST   | no          |             | Body: `{username: string, password: string}` | `{username: string}`     |              |     |
| Delete user      | `/user/{username}`       | DELETE | no          |             | Path: `{username: string}`                   | no content               |              |     |
| Create role      | `/role`                  | POST   | no          |             | Body: `{role_name: string}`                  | `{role_name: string}`    |              |     |
| Delete role      | `/role/{role_name}`      | DELETE | no          |             | Path: `{role_name: string}`                  | no content               |              |     |
| Add role to user | `/user/role`             | POST   | no          |             | Body:`{role_name: string}`                   | no content               |              |     |
| SignIn           | `/signin`                | POST   | no          |             | Body: `{username: string, password: string}` | `{token: string}`        |              |     |
| SingOut          | `/signout`               | POST   | auth token  |             |                                              | no content               |              |     |
| UserHasRole      | `/user/role/{role_name}` | HEAD   | auth token  |             | Body: `{role_name: string}`                  | `{result: bool}`         |              |     |
| Get User roles   | `/user/role`             | GET    | auth token  |             |                                              | `{roles: Array<string>}` |              |     |


## How to build and run

### Requirements

The project doesn't have any external dependencies and relies only on the standard library provided by golang.

In order to build and run the service you only need to have the latest version of golang on your machine.

- [go 1.19](https://go.dev/dl/)

### Build

Run the command below to build the app:

```shell
go build -o app ./cmd/app/main.go
```

### Run

You can run the executable built earlier or use `go run` command:

```shell
go run ./cmd/app/main.go
```

### Test

Run the command below to run the tests:

```shell
go test -v -count=1 -race ./...
```