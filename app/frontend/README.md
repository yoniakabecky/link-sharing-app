# Frontend

SvelteKit 2 + Svelte 5 frontend for the link-sharing app.

## Stack

- **SvelteKit 2** with Svelte 5 runes
- **TypeScript**
- **Valibot** for form validation
- **Vitest** for unit tests, **Playwright** for E2E tests
- **Storybook** for component documentation

## Prerequisites

- Node.js 18+
- npm

## Setup

1. Install dependencies:
   ```bash
   npm install
   ```

2. Configure environment variables:
   ```bash
   cp .env.example .env
   ```

3. Start the dev server:
   ```bash
   npm run dev
   ```

The app will be available at `http://localhost:5173`.

## Environment Variables

See [`.env.example`](.env.example).

## Commands

| Command | Description |
|---|---|
| `npm run dev` | Start development server |
| `npm run build` | Production build |
| `npm run check` | Type checking |
| `npm run lint` | Run Prettier + ESLint |
| `npm run format` | Auto-format code |
| `npm run test:unit` | Run unit tests (Vitest) |
| `npm run test:e2e` | Run E2E tests (Playwright) |
| `npm run storybook` | Start Storybook (port 6006) |

## Routes

| Path | Description | Access |
|---|---|---|
| `/` | Home / landing page | Public |
| `/shared/{id}` | View a shared profile | Public |
| `/register` | Sign up | Public |
| `/login` | Sign in | Public |
| `/dashboard` | Manage profiles | Authenticated |
| `/links` | Edit links for a profile | Authenticated |
| `/profile` | Edit profile info and avatar | Authenticated |
| `/preview` | Live preview of a profile | Authenticated |

## Project Structure

```
src/
├── routes/
│   ├── (auth)/         # Login and register (redirects away if logged in)
│   ├── (app)/          # Protected routes (enforces login)
│   │   ├── dashboard/
│   │   └── (editor)/   # Links and profile editor
│   └── (public)/       # Public-facing pages
│
└── lib/
    ├── components/     # Reusable UI components
    ├── remote/         # API client functions (one file per resource)
    ├── models/         # TypeScript types and Valibot validation schemas
    ├── fetcher.ts      # HTTP client (attaches auth token to requests)
    ├── state.svelte.ts # Global state (active profile, etc.)
    └── require-auth.ts # Auth guard utility
```

## Auth Flow

- Tokens (access + refresh) are stored in cookies.
- `hooks.server.ts` validates the session on every request and automatically refreshes an expired access token.
- `fetcher.ts` attaches `Authorization: Bearer <token>` to all API requests.
- The active profile ID is persisted in `sessionStorage`.
