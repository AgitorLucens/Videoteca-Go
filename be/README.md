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
- swaggo/swag v2 for OpenAPI 3.1 documentation

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
| Docs | `GET /swagger/index.html`, `GET /openapi.json` |

## API Documentation (Swagger / OpenAPI 3.1)

Interactive docs are generated with [swag](https://github.com/swaggo/swag) v2 from Go annotations on the handlers and served publicly at:

```text
http://localhost:8080/swagger/index.html
```

The raw OpenAPI 3.1 spec is served at:

```text
http://localhost:8080/openapi.json
```

The Swagger UI loads the spec from `/openapi.json` instead of the usual `/swagger/doc.json`, because the spec is registered by swag v2 while the gin-swagger UI helper reads from the swag v1 registry.

Protected endpoints accept the `BearerAuth` scheme (`Authorization: Bearer <token>`): call `POST /login` for a token, then use the **Authorize** button in the UI.

After changing annotations in handlers or `cmd/main.go`, regenerate the spec with:

```sh
go install github.com/swaggo/swag/v2/cmd/swag@latest
swag init -g cmd/main.go -d ./ --v3.1
```

The generated files live in `docs/` (`docs.go`, `swagger.json`, `swagger.yaml`) and are committed to the repo.

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
cmd/main.go                         # Application entry point (general API annotations)
internals/api/api.go                # Route registration + Swagger UI route
internals/handler/                  # HTTP handlers grouped by feature (per-endpoint annotations)
internals/middleware/               # JWT, RBAC, CORS, Prometheus middleware
internals/storage/                  # Catalog persistence
internals/rbac/                     # RBAC persistence
internals/migration/migrate.go      # Migration runner
docs/                               # Generated OpenAPI 3.1 spec (swag init --v3.1)
migrations/                         # SQL migrations
docker/                             # Prometheus and Grafana provisioning
```
