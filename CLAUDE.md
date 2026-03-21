# CLAUDE.md

## Project Overview

White-label Nx monorepo — a production-ready, open-source template with Go backend, Next.js frontend, and shared TypeScript libraries. Designed as a starter kit for building client-specific products.

- **Repository**: github.com/ssbreno/white-label-repository
- **License**: GPL-3.0
- **Default branch**: master

## Architecture

```
apps/backend/       — Go 1.21+ REST API (Gin + pgx v5 + PostgreSQL 16)
apps/frontend/      — Next.js 15 + React 18 + shadcn/ui + Tailwind CSS
libs/shared/types/  — @white-label/shared-types (interfaces, DTOs)
libs/shared/utils/  — @white-label/shared-utils (date, validation, string helpers)
libs/shared/config/ — @white-label/shared-config (endpoints, env schemas)
```

## Quick Start

### With Docker (recommended)

```bash
cp .env.example .env
docker-compose up -d
```

### Without Docker

```bash
npm install
docker-compose up postgres -d   # still need PostgreSQL
npm run dev:backend              # terminal 1
npm run dev:frontend             # terminal 2
```

### Full parallel dev

```bash
npm run dev
```

## Common Commands

```bash
# Development
make dev                    # run all services in parallel
make dev-frontend           # frontend only
make dev-backend            # backend only (requires PostgreSQL)
make install                # install all dependencies (npm + Go modules)

# Build & Quality
make build                  # build all apps
make test                   # run all tests
make test-backend           # Go tests only
make test-frontend          # Jest tests only
make lint                   # lint all projects
make format                 # format with Prettier

# Docker
make docker-up              # start all containers
make docker-down            # stop all containers
make docker-build           # build and start containers
make docker-logs            # follow container logs

# Database
make db-shell               # open PostgreSQL shell
make db-reset               # reset database (drop + recreate)

# Nx
make graph                  # visualize dependency graph
npx nx affected --target=test   # test only affected projects
```

## Go Backend Conventions (`apps/backend/`)

### Architecture (Clean Architecture)

```
cmd/server/main.go       — entry point, wiring only
internal/config/         — configuration structs, Load() from env
internal/database/       — DB connection pool (pgxpool) and migrations
internal/handlers/       — HTTP handlers (Gin), route registration
internal/middleware/      — Gin middleware (CORS, recovery, logging, rate limiting)
internal/models/         — Data models and request/response types
internal/repositories/   — Data access layer (SQL queries via pgx)
internal/services/       — Business logic layer
internal/ai/             — AI provider gateway (OpenRouter-style)
internal/mocks/          — Test mocks (manual, function-field pattern)
internal/testutil/       — Test helpers (httptest, testdb)
```

### Adding a New Resource (e.g., "products")

1. Create model: `internal/models/product.go`
2. Create repository interface + impl: `internal/repositories/product_repository.go`
3. Create service interface + impl: `internal/services/product_service.go`
4. Create handler: `internal/handlers/product_handler.go`
5. Register routes: `internal/handlers/routes.go`
6. Add migration: `internal/database/migrate.go`

### Naming Conventions

- Files: `snake_case` (`user_handler.go`, `user_repository.go`)
- Types: `PascalCase` (`UserService`, `CreateUserRequest`)
- Functions: `PascalCase` exported, `camelCase` unexported
- Handler methods: lowercase `camelCase` (`list`, `getByID`, `create`, `update`, `delete`)
- Constructor pattern: `NewXxx(deps) *Xxx`
- Interfaces: `XxxInterface` (e.g., `UserRepositoryInterface`, `UserServiceInterface`)

### Error Handling

- Wrap errors: `fmt.Errorf("context: %w", err)`
- HTTP errors: `models.NewError[any](err.Error())`
- HTTP success: `models.NewSuccess(data, "message")`
- All responses use `APIResponse[T]` generic wrapper

### Import Order

1. Standard library
2. Third-party packages
3. Internal packages (`github.com/ssbreno/white-label-repository/backend/internal/...`)

### Database

- pgx v5 with pgxpool for connection pooling
- SQL queries inline in repository methods
- Migrations tracked via `schema_migrations` table
- PostgreSQL UUIDs via `gen_random_uuid()`

### Testing

- Test files: `*_test.go` alongside source files
- Pattern: table-driven tests with `testify/assert`
- Mocks: manual function-field pattern in `internal/mocks/`
- HTTP tests: `httptest.NewRecorder()` + Gin test mode
- Integration tests: use `internal/testutil/testdb.go` with test database
- Run: `cd apps/backend && go test ./...`

## TypeScript/React Conventions (`apps/frontend/`)

### File Structure

```
src/app/              — Next.js App Router pages and layouts
src/components/       — React components
src/components/ui/    — shadcn/ui components (auto-generated, do not manually edit)
src/components/layout/ — Header, Footer
src/components/sections/ — Landing page sections (Hero, Features, CTA, FAQ)
src/hooks/            — Custom React hooks
src/lib/              — Utilities (api.ts, utils.ts, site-config.ts, content.ts)
src/types/            — TypeScript types (re-exports from shared)
```

### File Naming

- Components: `kebab-case` files (`theme-toggle.tsx`), `PascalCase` exports
- Pages: `page.tsx` in route directory (Next.js convention)
- Hooks: `use-xxx.ts` (e.g., `use-auth.ts`)
- Utils: `kebab-case` (`api.ts`, `utils.ts`)

### Component Patterns

- Server Components by default (no `'use client'` unless needed)
- `'use client'` only for components with state, effects, or event handlers
- Use `cn()` from `@/lib/utils` for className merging
- Use CSS variables for theming (never hardcode colors)
- shadcn/ui for UI primitives

### Import Order

1. React/Next.js
2. Third-party libraries
3. `@/components`, `@/hooks`, `@/lib`, `@/types` (path aliases)
4. `@white-label/shared-*` (workspace packages)
5. Relative imports
6. CSS/styles

### Path Aliases

```
@/*                         → ./src/*
@white-label/shared-types   → libs/shared/types/src/index.ts
@white-label/shared-utils   → libs/shared/utils/src/index.ts
@white-label/shared-config  → libs/shared/config/src/index.ts
```

### SEO

- All page content must be server-rendered (no `'use client'` on pages)
- Metadata via Next.js `Metadata` export on each page
- JSON-LD structured data via `<script type="application/ld+json">`
- Site config centralized in `src/lib/site-config.ts`
- robots.ts and sitemap.ts at app root

### Testing

- Jest + React Testing Library
- Test files in `__tests__/` directories alongside components
- Config: `apps/frontend/jest.config.ts`
- Run: `cd apps/frontend && npm test`

## Shared Libraries (`libs/shared/`)

### Adding a New Type

1. Add interface to `libs/shared/types/src/index.ts`
2. Import in frontend: `import type { MyType } from '@white-label/shared-types'`
3. Mirror Go struct in `apps/backend/internal/models/` (keep JSON tags aligned)

### Adding a New Utility

1. Add function to `libs/shared/utils/src/index.ts`
2. Export from barrel file
3. Import: `import { myUtil } from '@white-label/shared-utils'`

### Adding a New Config Entry

1. Add to `libs/shared/config/src/index.ts`
2. Update env schemas if adding env vars

## Code Formatting

- **TypeScript/JSON**: Prettier (semi: true, singleQuote: true, printWidth: 100, tabWidth: 2)
- **Go**: gofmt (enforced by editor + golangci-lint)
- **Linting Go**: golangci-lint with govet, errcheck, staticcheck, revive
- **Linting TS**: ESLint with next/core-web-vitals, jsx-a11y, import ordering

## Git Conventions

### Commit Messages (Conventional Commits)

```
<type>(<optional-scope>): <description>

Types: feat, fix, docs, refactor, test, chore, perf, style, ci, build

Examples:
  feat: add product listing API endpoint
  fix(frontend): resolve dark mode flash on page load
  docs: update API endpoint documentation
  test(backend): add user service unit tests
```

### Branch Naming

```
feature/<description>   — new features
fix/<description>       — bug fixes
docs/<description>      — documentation
refactor/<description>  — code restructuring
chore/<description>     — maintenance
```

## Environment Variables

- Root `.env.example` has all variables
- NEVER commit `.env` files (covered by .gitignore)
- Secrets (`JWT_SECRET`, `DB_PASSWORD`, API keys) must be changed in production
- AI provider keys (`OPENAI_API_KEY`, `ANTHROPIC_API_KEY`, `GOOGLE_API_KEY`) are optional — providers auto-enable when key is set

## Docker

- `docker-compose.yml` at root: postgres, backend, frontend
- Backend: multi-stage `golang:1.21-alpine` → `alpine:3.19`
- Frontend: multi-stage `node:20-alpine`
- Network: `whitelabel-network`
- Volume: `postgres_data` for database persistence
