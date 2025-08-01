# 🧑‍💻 For To-Do Backend App Maintainers

## 📝 Maintainer Guidelines

Before you start contributing to the To-Do Backend App, please make sure you have the following development tools installed on your local machine, which are used to maintain a clean and maintainable codebase:

- [Make](https://www.gnu.org/software/make/manual/make.html) (a build automation tool)
  - Brief explanation of the circumstances requiring the installation of Make and the procedure for installing the `make` tool

    | Operating System | `GNU` `make` Availability | Installation Instructions if Not Available |
    | :-- | :-- | :-- |
    | **Linux** (Ubuntu, Mint, Arch, etc.) | ✅ Pre-installed | No further action needed |
    | **macOS** | ✅ Comes with an older version (`v3.81`) | To upgrade, use `brew install make`.<br>(available as `gmake` to avoid conflicts) |
    | **Windows** | ❌ Absent by default | Options for installation: Chocolatey (`choco install make`), [GnuWin32](https://gnuwin32.sourceforge.net/packages/make.htm), or [MSYS2](https://www.msys2.org) |

  - To list all available `make` commands, run `make help`
- [Golangci-lint](https://golangci-lint.run) (a fast linters runner for Go)
  > **Note:** The `golangci-lint` configuration is defined in the `.golangci.yaml` file by default.
  - Install linter using `make check` command
    - (Or) install linter according to the official [installation instructions](https://golangci-lint.run/welcome/install/)
  - To lint Golang code, run `make lint`
  - To format Golang code, run `make fmt`
  - To integrate `golangci-lint` with [VSCode](https://code.visualstudio.com), follow these steps:
    1. Open VSCode settings:
        - On macOS press `⌘ + ,` or `Ctrl + ,` on Windows and Linux
    2. Navigate to the settings section (where `User` and `Workspace` are listed), under the `search settings` bar, select `Workspace`, then click on the `Open Settings (JSON)` icon located in the top-right corner next to the `search settings` bar.
       - This will open the `settings.json` file in the `.vscode` folder at the root of the project, where you can add settings that will only be applied to files in the current project.
    3. Add to `setting.json` the following lines:

        ```json
        {
          "go.lintTool":"golangci-lint",
          "go.lintFlags": [
            "--allow-parallel-runners"
          ]
        }
        ```

- [Pre-commit](https://pre-commit.com) (manages and maintains multi-language `pre-commit` hooks)
  - Install `pre-commit` using `make check` command
    - (Or) install `pre-commit` according to the [installation instructions](https://paolozaino.wordpress.com/2023/12/15/software-development-installing-pre-commit-to-check-our-code-repositories-and-improve-consistency-across-teams/)
    - Next, execute `make pre-commit-install` to set up the `pre-commit` hooks in your local `.git` repository
- [Upx](https://upx.github.io) (an advanced executable file compressor)
- [Goose](https://pressly.github.io/goose/) (a database migration tool)
- [Sqlc](https://docs.sqlc.dev/en/stable/index.html) (generates fully type-safe idiomatic Go code from SQL)
- [Docker](https://www.docker.com) (a containerization platform)
- [Docker Compose](https://docs.docker.com/compose/) (a tool for defining and running multi-container Docker applications)

## 🔥 Launching

The **To-Do Backend App** uses [Docker Compose](https://docs.docker.com/reference/cli/docker/compose/) to manage all its dependencies (e.g., database, cache) configured in the `build/docker-compose.yaml` file and run them in separate Docker containers:

1. To build and run those containers, run the following command:

    > **TIP:** useful flags for `docker compose up` are (all flags described in [documentation](https://docs.docker.com/reference/cli/docker/compose/up/#options)):
    > - `--build` - rebuild images before starting containers
    > - `--force-recreate` - recreate containers even if their configuration and image haven't changed
    > - `--detach` - run containers in the background

    ```bash
    # Output:
    # [+] Running 3/3
    # ✔ Network    postgres-network  Created 0.0s
    # ✔ Container  postgres          Started 0.2s
    docker compose --file=build/docker-compose.yaml up --detach
    ```

2. To stop and remove containers, networks, volumes, and images created by `up`, use the following command.

    ```bash
    # Output:
    # [+] Running 2/2
    # ✔ Container  postgres          Removed 0.2s
    # ✔ Network    postgres-network  Removed 0.2s
    docker compose --file=build/docker-compose.yaml down
    ```

3. To list containers for a Docker Compose project, with current status and exposed ports, run the following command:

    > **TIP:** useful flags for `docker compose ps` are (all flags described in [documentation](https://docs.docker.com/reference/cli/docker/compose/ps/#options)):
    > - `--all` - show all containers (only running containers by default)
    > - `--format` - format output using a custom template (`table` by default)
    > - `--quiet` - only display container IDs
    > - `--no-trunc` - don't truncate output

    ```bash
    # Output:
    # NAME IMAGE COMMAND SERVICE CREATED STATUS PORTS
    # postgres postgres:17.5-alpine3.22 "docker-entrypoint.s…" database About a minute ago Up About a minute (healthy) 127.0.0.1:5432->5432/tcp
    docker compose ps --all
    ```

## 🔨 Maintaining

### 📦 Docker image tags

The **To-Do Backend App** uses Docker images with specific tags (e.g., `postgres:17.5-alpine3.22`) to ensure reproducibility and consistency across different environments.

A brief explanation of the image tags:

| Tag pattern | What it is | When to pick it |
| :-- | :-- | :-- |
| `1.24.5-bookworm`, `1.24-bookworm`, `bookworm` | Debian 12 full image (~900 MB) | Pick when you need dynamic libraries, `apt`, man pages, or you are building inside the container (e.g., `go build`, `apt install`).           |
| `1.24.5-bullseye`, `1.24-bullseye`, `bullseye` | Debian 11 full image           | Same as above, but choose it only when you need the older Debian 11 base for compatibility.                                                   |
| `1.24.5-alpine3.22`, `1.24-alpine`, `alpine`   | Alpine Linux base (~50 MB)     | Pick when you want a small image that still has a shell and `apk`;<br>Great for CI and lightweight services.                                     |
| `1.24.5-slim-bullseye`, `1.24-slim`            | Debian “slim” (~100 MB)        | Pick when you need Debian stability but can drop docs, man pages, and extra locales;<br>Good drop-in replacement for full Debian.                |
| `scratch`                                      | Empty base (0 MB)              | Pick only for the final stage in a multi-stage build when you have a **static** CGO-disabled binary and need the absolute smallest footprint. |
| `tip-…`                                        | Nightly / bleeding-edge        | Pick only for experiments or when you need the latest Go features and accept instability; **never for production**.                           |

### 📦 Docker Life-Cycle from Image to Running Container

1. **📄 Image Creation**
    - An **image** is a read-only template with instructions for creating a Docker container.
    - You can create your own images using a `Dockerfile`, which contains a series of commands to build the image.
    - Each instruction in a `Dockerfile` creates a layer in the image.
    - When you change the `Dockerfile` and rebuild the image, only the changed layers are rebuilt.
    - To build the image, you use the `docker buildx build` command:

      ```bash
      docker buildx build --file=build/Dockerfile --tag=todo-backend:latest .
      # Or
      docker build --file=build/Dockerfile --tag=todo-backend:latest .
      ```

    1.1 **📦 Image Inspection**
    - To display a list of all top-level images that have been built, along with their repository, tags, and sizes, use the following command:

      ```bash
      docker image list
      # Or
      docker images
      ```

    1.2 **🗑️ Image Removeing**
    - To remove (and un-tag) one or more images from the host node, use the `docker image remove` command.

      ```bash
      docker image remove todo-backend:latest
      # Or
      docker rmi todo-backend:latest
      ```

    - You cannot remove an image of a running container unless you use the `-f`/`--force` option.

2. **📦 Container Creation**
    - A **container** is a runnable instance of an image.
    - You can create a container using the `docker container create` command.
    - This command prepares the container but does not start it:

      ```bash
      docker container create --name=todo-app --rm --publish=8081:8081 todo-backend:latest
      # Or
      docker create --name=todo-app --rm --publish=8081:8081 todo-backend:latest
      ```

    - Alternatively, you can use `docker container run`, which combines the creation and starting of the container in one step:

      ```bash
      docker container run --name=todo-app --rm --publish=8081:8081 --detach todo-backend:latest
      # Or
      docker run --name=todo-app --rm --publish=8081:8081 --detach todo-backend:latest
      ```

3. **🔥 Container Execution**
    - When a container is running, it executes the commands specified in the image.
    - While running, the container’s main process (PID 1) is active and can handle network requests, write data to volumes, and communicate with other containers
    - You can start a created container using the `docker container start` command

      ```bash
      docker container start todo-app
      # Or
      docker start todo-app
      ```

4. **⏸️ Container Pause and Unpause**
    - You can pause a running container using the docker pause command.
    - This sends a `SIGSTOP` signal to all processes inside the container, freezing them without terminating them

      ```bash
      docker container pause todo-app
      # Or
      docker pause todo-app
      ```

    - To resume the container, use the `docker unpause` command

      ```bash
      docker container unpause todo-app
      # Or
      docker unpause todo-app
      ```

5. **🛑 Container Stop**
    - A container enters the stopped state when its main process finishes or when you manually stop it using the docker stop command.
    - This sends a `SIGTERM` signal to the main process, allowing it to shut down gracefully.
    - If the process does not stop within a grace period, Docker sends a `SIGKILL` signal.
    - Stopped containers can be restarted using the `docker start` command.

      ```bash
      docker container stop todo-app
      # Or
      docker stop todo-app
      ```

6. **🗑️ Container Removal**
    - To remove a container, use the `docker rm` command (works only on stopped containers).
    - This deletes the container’s metadata and filesystem, freeing up resources.
    - You can force-remove a running container using the `-f` flag (uses `SIGKILL`)

      ```bash
      docker container remove todo-app
      # Or
      docker rm todo-app
      ```

### 💾 Database

The **To-Do Backend App** uses [PostgreSQL](https://www.postgresql.org) as its database management system.

The current database diagram is shown below:

![Database Diagram](./static/db-diagram.png)

> **NOTE:** Database diagram is generated using [ChartDB](https://chartdb.io)

### 💾 Database migrations

The **To-Do Backend App** uses the [Goose](https://github.com/pressly/goose) package to manage database migrations.

1. To create a new migration (with `.sql` extension), run the following command:

    ```bash
    # env_file_name: path to the `.env` file
    # migration_name: name of the new migration
    goose -env <env_file_name> -s create <migration_name> sql
    ```

2. To run migrations (i.e. update the database to the most recent version available), run the following command:

    ```bash
    goose -env <env_file_name> up
    ```

3. To rollback migrations (i.e. return the database to the previous version), run the following command:

    ```bash
    # Roll back the version by 1
    goose -env <env_file_name> down
    # Roll back all migrations
    goose -env <env_file_name> reset
    ```

4. To list applied migrations and their statuses, run the following command:

    ```bash
    goose -env <env_file_name> status
    ```

## ⚙️ Sql to Go code generation

[sqlc](https://docs.sqlc.dev/en/stable/index.html) generates fully type-safe idiomatic Go code from SQL.

Here’s how it works:

1. You write SQL queries in files with `.sql` extension
2. You run `sqlc` to generate Go code that presents type-safe interfaces to those queries
3. You write app code that calls the methods `sqlc` generated

> **NOTE:** The `sqlc.yaml` file contains the configuration for `sqlc`.
>
> - The `schema` field specifies the directory containing SQL migration files, and `sqlc` has support for parsing `goose` migrations (more details are available in the [documentation](https://docs.sqlc.dev/en/latest/howto/ddl.html#goose)).

### ⚙️ Go Code Genereation

Execute the following command to generate Go code from SQL scripts located in the `src/internal/adapters/databases/postgres/queries` directory:

```bash
DB_USER=postgres DB_PASSWORD=postgres DB_HOST=0.0.0.0 DB_PORT=5432 DB_NAME=postgres DB_SSL_MODE=disable sqlc generate
```

By default, `sqlc` runs its analysis using a built-in query analysis engine.

- While fast, this engine can’t handle some complex queries and type-inference.
- That's why we configure `sqlc` to use our own database for enhanced analysis using metadata from it.
- Database-backed analysis is configured in the `sqlc.yaml` by the following fields:

  ```yaml
  database:
    managed: false
    uri: postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}
  ```

### 🔎 Queries Checking and Linting

To perform static analysis of SQL queries for syntax and type errors, use the following command:

```bash
DB_USER=postgres DB_PASSWORD=postgres DB_HOST=0.0.0.0 DB_PORT=5432 DB_NAME=postgres DB_SSL_MODE=disable sqlc compile
```

To lints queries using rules (ex: `postgresql-no-delete-without-where` or `postgresql-no-update-without-where`) defined in the `sqlc.yaml` configuration file, use the following command:

```bash
DB_USER=postgres DB_PASSWORD=postgres DB_HOST=0.0.0.0 DB_PORT=5432 DB_NAME=postgres DB_SSL_MODE=disable sqlc vet
```

## References

...
