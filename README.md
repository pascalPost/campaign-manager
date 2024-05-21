# campaign-manager

campaign-manager is a project designed to manage computation campaigns (execution of a lot of potentially
compute intensive, long-running tasks) with a user-friendly frontend for creating, editing, monitoring and interaction.
The campaigns can be generated based on Go templates and YAML-based configuration files.

## Project Structure

The project consists of two main building blocks:

1. File explorer and file editing
2. Launch and managing of computations via workload manager

## Architecture

The architecture of the project is as follows:

```plaintext
                   +--------------------------+
                   |                          |
                   |         Frontend         |
                   |         (React)          |
                   |                          |
                   +------------+-------------+
                                |
                                | HTTP Requests (REST API)
                                |
                   +------------v-------------+
                   |                          |
                   |         Backend          |
                   |       (GoLang API)       |
                   |                          |
                   +------------+-------------+
                                |
                                | Workload Manager
                                |
                   +------------v-------------+
                   |                          |
                   |     High-Performance     |
                   |     Computing Cluster    |
                   |                          |
                   +--------------------------+
```

### Backend

The backend is written in Go and uses [oapi-codegen](https://github.com/deepmap/oapi-codegen) to generate REST API
boilerplate from `./api/openapi.yaml`.
The router used is chi (https://github.com/go-chi/chi).

#### Development

For auto-update (re-build) on change, you can use air, running (
after [installation](https://github.com/cosmtrek/air#installation))
```shell
air
```
The air setup will also run the restapi code generation, which can be also run manually
```shell
go generate ./...
```

### Frontend

The frontend is built with React based on the shadcn-ui components and communicates with the backend through HTTP
requests (REST API).
The calls are generate based on the OpenAPI documentation `./api/openapi.yaml`.
The frontend is a git submodule of https://github.com/pascalPost/campaign-manager-frontend.git.

#### Development

The dev server can be started with `pnpm run dev`.
