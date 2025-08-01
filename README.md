# 📋 To-Do Backend App

## Overview

**The To-Do Backend App** - is a Golang web app designed to manage tasks efficiently.

- It provides a RESTful API for creating, reading, updating, and deleting to-do items.
- The app is built using Go's robust standard library and follows best practices for clean and maintainable code.

## ✨ Features

...

## ⚙️ Techonologies

| Logo | Technology | Description |
| :-- | :-- | :-- |
| <img src="https://avatars.githubusercontent.com/u/18133?s=280&v=4" alt="Git Logo" width="40"> | [Git](https://git-scm.com) | Is a free and open source distributed version control system |
| <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/c/c2/GitHub_Invertocat_Logo.svg/250px-GitHub_Invertocat_Logo.svg.png" alt="Github Logo" width="40"> | [Github](https://github.com) | Is a proprietary developer platform that allows developers to create, store, manage, and share their code. |
| <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Golang Logo" width="40"> | [Golang (Go)](https://go.dev) | An open-source programming language with a built-in concurrency and a robust standard library |
| <img src="https://golangci-lint.run/logo.png" alt="Golang CI Logo" width="40"> | [Golangci-lint](https://golangci-lint.run) | Is a fast linters runner for Go. |
| <img src="https://avatars.githubusercontent.com/u/4092?s=48&v=4" alt="Dotenv Logo" width="40"> | [Godotenv](https://github.com/joho/godotenv) | Loads env vars from a `.env` file |
| <img src="https://avatars.githubusercontent.com/u/68232?s=48&v=4" alt="Zerolog logo" width="40"> | [Zerolog](https://github.com/rs/zerolog) | A fast and simple logger dedicated to JSON output |
| <img src="https://avatars.githubusercontent.com/u/1841476?s=48&v=4" alt="Testify logo" width="40"> | [Testify](https://github.com/stretchr/testify) | A framework for writing tests in Go. |
| <img src="https://miro.medium.com/v2/resize:fit:700/1*s8I4jBW2KKP687LqWh3OtQ.png" alt="Docker Compose logo" width="40"> | [Docker Compose](https://docs.docker.com/compose/) | Is a tool for defining and running multi-container apps. |
| <img src="https://img.icons8.com/fluent/512/docker.png" alt="Docker logo" width="40"> | [Docker](https://docs.docker.com/engine/) | Is an open source containerization technology for building and containerizing your apps |
| <img src="https://github.com/pressly/goose/raw/main/assets/goose_logo.png" alt="Goose logo" width="40"> | [Goose](https://pkg.go.dev/github.com/pressly/goose/v3#section-readme) | Is a database migration tool. Both a CLI and a library |
| <img src="https://avatars.githubusercontent.com/u/94130?s=48&v=4" alt="pgx" width="40"> | [Pgx](https://github.com/jackc/pgx#pgx---postgresql-driver-and-toolkit) | Is a pure Go driver and toolkit for PostgreSQL. |
| <img src="https://static-00.iconduck.com/assets.00/postgresql-icon-1987x2048-v2fkmdaw.png" alt="Postgres logo" width="40"> | [Postgres](https://www.postgresql.org) | Open source object-relational database system |
| <img src="https://docs.sqlc.dev/en/stable/_static/logo.png" alt="Sqlc logo" width="40"> | [Sqlc](https://docs.sqlc.dev/en/stable/index.html) | Generates fully type-safe idiomatic Go code from SQL |
| <img src="https://avatars.githubusercontent.com/u/689082?s=48&v=4" alt="Squirrel logo" width="40"> | [Squirrel](https://pkg.go.dev/github.com/Masterminds/squirrel#section-readme) | A fluent SQL generator for Go |

## 💻 Installation

To install the **To-Do Backend App**, follow these steps:

1. ⬇️ Clone the To-Do Backend App [repository](https://github.com/RenZorRUS/todo-backend) to your local machine using the following command:

    ```bash
    # HTTPS
    git clone https://github.com/RenZorRUS/todo-backend.git
    # SSH
    git clone git@github.com:RenZorRUS/todo-backend.git
    ```

2. 🧭 Navigate to the project directory and install the dependencies using the following command:

    ```bash
    cd todo-backend
    ```

## 🔥 Launching

The **To-Do Backend App** uses [Docker Compose](https://docs.docker.com/reference/cli/docker/compose/) to manage all its dependencies (e.g., database, cache) configured in the `build/docker-compose.yaml` file and run them in separate [docker containers](https://www.docker.com/resources/what-container/):

1. To build and run those containers, run the following command:

    ```bash
    docker compose --file=build/docker-compose.yaml up --detach
    ```

2. To stop and remove containers, networks, volumes, and images created by `up`, use the following command.

    ```bash
    docker compose --file=build/docker-compose.yaml down
    ```

3. To list containers for a Docker Compose project, with current status and exposed ports, run the following command:

    ```bash
    docker compose ps --all
    ```

## 🧑‍💻 For Maintainers

For those maintaining the To-Do Backend App, refer to the [📃 documentation](/docs/for-maintainers.md) to understand the necessary installations and setup required for your development and contribution efforts.

## 🔗 References

### 🗣️ Languages

- [Go Style Decisions](https://google.github.io/styleguide/)
