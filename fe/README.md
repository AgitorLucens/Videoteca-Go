# Videoteca Frontend

React/Vite frontend for Videoteca. It handles authentication, protected navigation, movie and series browsing, profile management, comments, ratings, and admin screens for catalog maintenance.

## Tech Stack

- React 19
- Vite 7
- React Router
- Axios
- Tailwind CSS
- React Slick
- Vitest and Testing Library
- ESLint

## Setup

Install dependencies:

```sh
npm install
```

Start the development server:

```sh
npm run dev
```

Vite prints the local app URL, usually:

```text
http://localhost:5173
```

The API client currently uses:

```text
http://localhost:8080
```

Start the backend before using authenticated or data-driven screens.

## Scripts

| Command | Description |
| --- | --- |
| `npm run dev` | Start the Vite development server |
| `npm run build` | Create a production build in `dist/` |
| `npm run preview` | Preview the production build locally |
| `npm run lint` | Run ESLint |
| `npm test` | Run Vitest once |
| `npm run test:watch` | Run Vitest in watch mode |

## Routes

| Route | Purpose |
| --- | --- |
| `/` | Login and account entry |
| `/forgot-password` | Password recovery |
| `/home` | Authenticated home page |
| `/main` | Main browsing page |
| `/movies/:id` | Movie or series details |
| `/profile` | User profile |
| `/admin/movies` | Admin movie/series list |
| `/admin/movies/new` | Create a movie or series |
| `/admin/movies/:id/edit` | Edit a movie or series |
| `/admin/tags` | Manage genres/tags |
| `/admin/actors` | Manage actors |
| `/superadmin` | Superadmin tools |

Protected routes are wrapped by `src/routes/ProtectedRoutes.jsx` and use the auth state from `src/hooks/useAuth.jsx`.

## API Client

Shared HTTP behavior lives in:

```text
src/services/http.service.js
```

It creates one Axios instance, adds the JWT bearer token from `localStorage.user`, and redirects to `/` on `401` responses.

Feature-specific service files live in `src/services/`, including:

- `login.service.js`
- `movieserie.service.js`
- `search.service.js`
- `rating.service.js`
- `comment.service.js`
- `profile.service.js`
- `genre.service.js`
- `actor.service.js`
- `superadmin.service.js`

## Project Structure

```text
src/
+-- components/   # Shared UI components
+-- hooks/        # Auth and local storage hooks
+-- pages/        # Route-level screens
+-- routes/       # Protected route wrapper
+-- services/     # API service modules
+-- test/         # Test setup
+-- utils/        # Utility helpers
+-- App.jsx       # Router and route registration
+-- main.jsx      # React entry point
```

## Tests

Run all frontend tests:

```sh
npm test
```

Run tests while developing:

```sh
npm run test:watch
```

The test environment is configured in `vite.config.js` with `jsdom`, globals enabled, and `src/test/setup.js`.

## Build

```sh
npm run build
```

Preview the built app:

```sh
npm run preview
```

## Development Notes

- The app stores the authenticated user in `localStorage` under `user`.
- The JWT token is attached automatically to requests when present.
- Route protection checks the current auth state before rendering nested pages.
- Carousel UI uses `react-slick` and `slick-carousel`.
- SVG assets used as React components must be valid JSX.
