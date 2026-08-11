# Videoteca Go

Videoteca Go is a full-stack movie and series catalog application. The backend is a Go API backed by PostgreSQL, and the frontend is a React/Vite single-page app for authentication, browsing titles, managing profiles, and administering catalog data.

## Project Structure

```text
.
+-- be/   # Go API, PostgreSQL migrations, RBAC, Prometheus/Grafana config
+-- fe/   # React + Vite frontend
```

## Features

- User authentication with JWT.
- Role-based access control for users, admins, and superadmins.
- Movie and series catalog browsing.
- Search, ratings, comments, and profile management.
- Admin tools for movies, series, genres, and actors.
- Superadmin tools for admin account management.
- Backend metrics exposed to Prometheus and visualized in Grafana.

## Tech Stack

| Area | Tools |
| --- | --- |
| Backend | Go, Gin, GORM, PostgreSQL, golang-migrate |
| Auth | JWT, bcrypt, RBAC |
| Monitoring | Prometheus, Grafana |
| Frontend | React, Vite, React Router, Axios, Tailwind CSS |
| Testing | Go test, Vitest, Testing Library |

## Getting Started

### Prerequisites

- Go 1.24.1 or newer.
- Node.js and npm.
- PostgreSQL.
- Docker, if you want to run Prometheus and Grafana.

### Backend

```sh
cd be
go mod download
go run ./cmd
```

The API runs on `http://localhost:8080`. See [be/README.md](be/README.md) for environment variables, migrations, routes, tests, and monitoring.

### Frontend

```sh
cd fe
npm install
npm run dev
```

The Vite dev server prints the local frontend URL, usually `http://localhost:5173`. The frontend expects the backend at `http://localhost:8080`. See [fe/README.md](fe/README.md) for scripts, routing, and development notes.

## Common Workflow

1. Start PostgreSQL and create the database expected by the backend `.env`.
2. Start the backend from `be/`; migrations run automatically on startup.
3. Start the frontend from `fe/`.
4. Open the frontend and sign in or create users through the available API/UI flows.

## Validation

Run backend tests:

```sh
cd be
go test ./...
```

Run frontend checks:

```sh
cd fe
npm run lint
npm test
npm run build
```

## Monitoring

From `be/`, start the monitoring stack:

```sh
docker compose up -d
```

- Prometheus: `http://localhost:9090`
- Grafana: `http://localhost:3000` using `admin` / `admin`

Grafana is preconfigured with a Prometheus datasource and a Videoteca backend dashboard.
