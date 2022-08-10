# UMSystem

## About

API service for user and role management.

## REST API

Below is a table of available REST API endpoints and their description.

name endpoint description request response status codes

| Name             | Endpoint                 | Method | Auth method | Description | Request schema                               | Response schema               |
|------------------|--------------------------|--------|-------------|-------------|----------------------------------------------|-------------------------------|
| Create user      | `/user`                  | POST   | no          |             | Body: `{username: string, password: string}` | `{username: string}`          |
| Delete user      | `/user/{username}`       | DELETE | no          |             | Path: `{username: string}`                   | no content                    |
| Create role      | `/role`                  | POST   | no          |             | Body: `{role_name: string}`                  | `{role_name: string}`         |
| Delete role      | `/role/{role_name}`      | DELETE | no          |             | Path: `{role_name: string}`                  | no content                    |
| Add role to user | `/user/role`             | POST   | no          |             | Body:`{role_name: string}`                   | no content                    |
| SignIn           | `/signin`                | POST   | no          |             | Body: `{username: string, password: string}` | `{token: string}`             |
| SingOut          | `/signout`               | POST   | auth token  |             |                                              | no content                    |
| UserHasRole      | `/user/role/{role_name}` | HEAD   | auth token  |             | Body: `{role_name: string}`                  | `{result: bool}`              |
| Get User roles   | `/user/role`             | GET    | auth token  |             |                                              | `{roles: Array<string>}`      |

## How to build and run

### Requirements

In order to build and run the service you only need to have the latest version of golang on your machine.

- [go 1.19](https://go.dev/dl/)

### Dependencies

- [Echo web framework](github.com/labstack/echo/v4) - better alternative to the standard library's http server mux to
  avoid a lot of boilerplate code.
- [golang crypto](golang.org/x/crypto) - experimental crypto library for hashing the passwords with salt and comparing
  them in constant time to secure the server from time-based brute force attacks.

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

## TODO:

- [ ] Add more tests for edge cases.
- [ ] Make full test coverage for API server, though the underlying services logic is tested.
- [ ] Use `context.Context` for service and store as we might use real database in the future the fore request
  cancellation is must have.
- [ ] Implement transactional operations for in-memory database. Current version doesn't have transaction and might lead
  to data races in some cases.
- [ ] Make application configurable through CLI flags or config file.