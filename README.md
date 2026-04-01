# Link Sharing App

A full-stack web application for creating and sharing profile pages with social/platform links. Built as a skills demonstration project.

Inspired by the [Frontend Mentor - Link Sharing App Challenge](https://www.frontendmentor.io/challenges/linksharing-app-Fbt7yweGsT).

<p align="center">
  <img src="/app/frontend/src/lib/assets/lp_image.webp" alt="App preview" width="700" />
</p>

## Tech Stack

| Layer            | Technology                        |
| ---------------- | --------------------------------- |
| Frontend         | SvelteKit 2, Svelte 5, TypeScript |
| Backend          | Go 1.25, Chi v5                   |
| Database         | MySQL 8                           |
| Containerization | Docker / Docker Compose           |

## Features

- **Multi-profile management** — create, edit, and delete multiple profiles per account
- **Link management** — add, edit, and reorder social/platform links per profile
- **Avatar uploads** — PNG/JPEG/BMP, up to 2MB
- **Public shareable profiles** — each profile has a public URL accessible without authentication
- **JWT authentication** — HS256 access tokens (15 min) with opaque refresh token rotation (7 days), stored hashed in DB
- **Responsive UI** — live preview with mobile mockup display
- **Component library** — Storybook documentation for all UI components

## Architecture

```
app/frontend/    # SvelteKit 2 + Svelte 5 frontend
backend/         # Go REST API
docker-compose.yml
```

**Backend** follows a layered architecture: `Handler → Service → Repository`. Chi middleware injects the authenticated `userId` into request context on protected routes.

**Frontend** uses Svelte 5 runes (`$state`, `$derived`, `$effect`) throughout. All API calls go through remote functions (`src/lib/remote/*.remote.ts`) via a central `fetcher.ts` that attaches the JWT. Tokens are stored in cookies and validated (with automatic refresh) in `hooks.server.ts` on every request.

## Getting Started

**Prerequisites:** Docker, Go, [Task](https://taskfile.dev), [Air](https://github.com/air-verse/air), [dbmate](https://github.com/amacneil/dbmate), Node.js

```bash
# 1. Start the database
docker compose up -d

# 2. Start the backend
cd backend
cp .env.example .env   # fill in values
task migrate_up
task run_server        # hot reload via Air, default port 8080

# 3. Start the frontend (new terminal)
cd app/frontend
cp .env.example .env   # set API_BASE_URL=http://localhost:8080
npm install
npm run dev            # default port 5173
```

## Commands

### Frontend (`app/frontend/`)

```bash
npm run dev           # Dev server
npm run build         # Production build
npm run check         # Type checking
npm run lint          # Prettier + ESLint
npm run format        # Auto-format
npm run test:unit     # Vitest unit tests
npm run test:e2e      # Playwright E2E tests
npm run storybook     # Component explorer (port 6006)
```

### Backend (`backend/`)

```bash
task run_server       # Start with hot reload (Air)
task migrate_new      # Create new migration
task migrate_up       # Apply migrations
task migrate_down     # Rollback last migration
task reset_db         # Drop + recreate + migrate
go test ./...         # Run all tests
```
