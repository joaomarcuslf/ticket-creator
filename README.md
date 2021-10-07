
# ticket-creator

[![go](https://github.com/joaomarcuslf/ticket-creator/actions/workflows/go.yml/badge.svg)](https://github.com/joaomarcuslf/ticket-creator/actions/workflows/go.yml)

This is a service that creates Tickets to be shared for others.

I used clean-arch concepts, and TDD.

![image](https://raw.githubusercontent.com/joaomarcuslf/ticket-creator/main/static/ticket-for-you.png)

## Getting Started

1. Copy ```sample.env``` to ```.env``` and rename the variables if you need
2. You can run this repo on vscode

![image](https://raw.githubusercontent.com/joaomarcuslf/ticket-creator/main/static/run-application.png)

Or, you can run by doing this:

```sh
PORT=8080 go run main.go
```

### Using other clients

This application was built to answer as 3 kinds of client.

#### App

This is the default, it answer HTML, like a normal webapp.

To set APP mode:

```sh
PORT=8080 client=app go run main.go
```

#### Rest

This is the Rest client, it answer JSON, and TBD XML.

To set REST mode:

```sh
PORT=8080 client=app go run main.go
```

#### gRPC

This is the gRPC client, it answer on rpc clients, like Bloomrpc.

To set GRPC mode:

```sh
PORT=8080 client=grpc go run main.go
```

## Running Tests

```sh
go test -cover ./...
```

## Deploying

On pushing to `main`, it will automatically push to heroku

## Collaborators

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/joaomarcuslf">
        <img src="https://avatars.githubusercontent.com/u/53450523?v=4" width="100px;" alt="Joaomarcuslf's Github picture"/><br>
        <sub>
          <b>joaomarcuslf</b>
        </sub>
      </a>
    </td>
  </tr>
</table>

[⬆ Scroll top](#ticket-creator)<br>
