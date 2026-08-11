# Videoteca Backend

Go API for the Videoteca application. It provides authentication, role-based authorization, catalog management, user profile features, comments, ratings, search, database migrations, and Prometheus metrics.

## Tech Stack

- Go 1.24.1
- Gin HTTP router
- GORM with PostgreSQL
- golang-migrate for schema migrations
- JWT for authenticated routes
- bcrypt for password hashing
- Prometheus and Grafana for monitoring

## Setup

Install dependencies:

```sh
go mod download
```

Create a `.env` file in `be/`:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=videoteca
DB_PORT=5432
DB_SSLMODE=disable
JWT_SECRET=replace-with-a-secure-secret
```

Generate a JWT secret when needed:

```sh
openssl rand -hex 32
```

Make sure the PostgreSQL database exists before starting the API. The application applies migrations automatically from the `migrations/` directory during startup.

## Run

```sh
go run ./cmd
```

The API listens on:

```text
http://localhost:8080
```

## Build

```sh
go build -o videoteca-api ./cmd
```

On Windows:

```sh
go build -o videoteca-api.exe ./cmd
```

## Tests

```sh
go test ./...
```

## Main API Areas

| Area | Routes |
| --- | --- |
| Authentication | `POST /login`, `POST /forgot-password`, `POST /reset-password` |
| Users | `POST /users`, `GET /users/:id` |
| Catalog | `GET /movieseries`, `GET /movieseries/:id`, `GET /search` |
| Ratings | `POST /movieseries/:id/rate` |
| Comments | `GET /movieseries/:id/comments`, `POST /movieseries/:id/comments`, `DELETE /comments/:commentId` |
| Profile | `GET /user/profile`, `PUT /user/email`, `PUT /user/password`, `PUT /user/username`, `PUT /user/picture` |
| Admin | `/admin/movieseries`, `/admin/genres`, `/admin/actors` |
| Superadmin | `/superadmin/admins`, `/superadmin/admins/password` |
| Metrics | `GET /metrics` |

## Authorization

The API uses JWT bearer tokens and role-based access control.

- Public routes handle login, registration, password recovery, and basic user lookup.
- Authenticated user/admin routes require a valid JWT.
- Admin routes require the `admin` role.
- Superadmin routes require the `superadmin` role.

Send tokens as:

```http
Authorization: Bearer <token>
```

## Database

Migrations live in `migrations/`.

- `000001_init_rbac_schema` creates users, roles, access records, and role mappings.
- `000002_init_app_schema` creates movie/series, genres, actors, episodes, comments, ratings, and related join tables.

The migration runner is called from `internals/server.go` and uses the database settings from `.env`.

## Monitoring

The backend exposes Prometheus metrics at:

```text
http://localhost:8080/metrics
```

Prometheus and Grafana are configured in `docker-compose.yml`.

Start the monitoring stack:

```sh
docker compose up -d
```

Access the tools:

- Prometheus: `http://localhost:9090`
- Grafana: `http://localhost:3000`
- Grafana login: `admin` / `admin`

Grafana loads the Prometheus datasource and the `Videoteca Backend` dashboard from `docker/grafana/provisioning/` and `docker/grafana/dashboards/`.

### Metrics

| Metric | Type | Labels | Description |
| --- | --- | --- | --- |
| `http_requests_total` | Counter | `method`, `endpoint`, `status_code` | Total HTTP requests |
| `http_request_duration_seconds` | Histogram | `method`, `endpoint` | Request latency |
| `http_requests_in_progress` | Gauge | `method`, `endpoint` | Active requests |
| `connected_devices` | Gauge | `user_id`, `role`, `device` | Active JWT sessions by user agent |
| `login_total` | Counter | `status` | Login attempts by success/failure |
| `build_info` | Gauge | `version`, `module` | Build metadata |

Prometheus scrapes `host.docker.internal:8080/metrics` every 15 seconds, so keep the API running locally when using the monitoring stack.

## Useful Paths

```text
cmd/main.go                         # Application entry point
internals/api/api.go                # Route registration
internals/middleware/               # JWT, RBAC, CORS, Prometheus middleware
internals/handler/                  # HTTP handlers grouped by feature
internals/storage/                  # Catalog persistence
internals/rbac/                     # RBAC persistence
internals/migration/migrate.go      # Migration runner
migrations/                         # SQL migrations
docker/                             # Prometheus and Grafana provisioning
```
