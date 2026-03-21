# Contributing to White Label Repository

Thank you for your interest in contributing! This guide will help you get started.

## Prerequisites

- [Node.js](https://nodejs.org/) 20+
- [Go](https://golang.org/) 1.21+
- [Docker](https://www.docker.com/) & Docker Compose
- [npm](https://www.npmjs.com/) 10+

## Development Setup

1. **Fork and clone** the repository
2. **Install dependencies**: `make install`
3. **Copy environment**: `cp .env.example .env`
4. **Start services**: `make docker-up && make dev`

See [CLAUDE.md](CLAUDE.md) for the full command reference.

## Code Style

### Go (Backend)

- Format with `gofmt` (automatic in VS Code with recommended extensions)
- Lint with `golangci-lint` (`cd apps/backend && make lint`)
- Follow clean architecture: handler → service → repository
- Error wrapping: `fmt.Errorf("context: %w", err)`
- See [CLAUDE.md](CLAUDE.md) for detailed patterns

### TypeScript (Frontend & Libraries)

- Format with Prettier (`make format`)
- Lint with ESLint (`make lint`)
- Semi: true, single quotes, 100 char width, 2-space indent
- Use `cn()` for className composition
- Use CSS variables for colors (never hardcode)
- Server Components by default (no `'use client'` unless interactive)

## Git Workflow

### Branch Naming

```
feature/<description>   — new features
fix/<description>       — bug fixes
docs/<description>      — documentation
refactor/<description>  — code restructuring
chore/<description>     — maintenance
```

### Commit Messages (Conventional Commits)

```
<type>(<optional-scope>): <description>

Types: feat, fix, docs, refactor, test, chore, perf, style, ci, build
```

Examples:

```
feat: add product listing API endpoint
fix(frontend): resolve dark mode flash on page load
test(backend): add user service unit tests
```

### Pull Requests

1. Create a branch from `master`
2. Make your changes
3. Ensure all tests pass: `make test`
4. Ensure code is formatted: `make format` and `make lint`
5. Push and open a PR against `master`
6. Fill out the PR template
7. Request review

## Adding Features

### New Backend Resource

See the 6-step workflow in [CLAUDE.md](CLAUDE.md#adding-a-new-resource-eg-products).

### New Frontend Page

1. Create route directory: `apps/frontend/src/app/<route>/`
2. Add `page.tsx` (Server Component by default)
3. Export `metadata` for SEO
4. Use section components from `src/components/sections/`

### New Shared Type

1. Add interface to `libs/shared/types/src/index.ts`
2. Mirror Go struct in `apps/backend/internal/models/`
3. Keep JSON field names aligned

## Reporting Issues

- Use the provided issue templates
- Include steps to reproduce for bugs
- Include expected vs actual behavior
- Include environment information (OS, Node.js version, Go version)

## Code of Conduct

See [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).
