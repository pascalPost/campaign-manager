# campaign-manager

Project to manage campaigns based on go templates and a YAML based configuration file.
The project is based on the [Wails](https://wails.io) framework and offers a web interface.
The backend is written in Go.

## Architecture

### Overview
                   +--------------------------+
                   |                          |
                   |         Frontend         |
                   |         (React)          |
                   |          Wails           |
                   +------------+-------------+
                                |
                                | HTTP Requests (REST)
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

### Backend

For server development: use air to update (re-build) on change
```shell
air
```
The REST API is available at http://localhost:3000/api/v1.



## Dependencies
- [Wails](https://wails.io)

## Wails

### About

[Wails template for Nextjs with app router](https://github.com/thisisvk-in/wails-template-nextjs-app-router).
For more details [Nextjs-Template](https://github.com/thisisvk-in/wails-template-nextjs-app-router).

### Live Development

To run in live development mode, run `wails dev` in the project directory. In another terminal, go into the `frontend`
directory and run `npm run dev`. The frontend dev server will run on http://localhost:3000. Connect to this in your
browser and connect to your application.

### Building

To build a redistributable, production mode package, use `wails build`. Static asset directory will be `frontend/dist`.
