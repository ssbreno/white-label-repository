# System Architecture

## Overview

The white-label monorepo is structured as an Nx workspace containing two applications and three shared libraries.

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                        Client Browser                           │
└────────────────────────────┬────────────────────────────────────┘
                             │ HTTPS :3000
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Frontend (Next.js 14)                        │
│                    apps/frontend                                │
│                                                                 │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────────┐  │
│  │  App Router  │  │  shadcn/ui   │  │   API Client (axios) │  │
│  │   (pages)    │  │  Components  │  │   lib/api.ts         │  │
│  └──────────────┘  └──────────────┘  └──────────┬───────────┘  │
└────────────────────────────────────────────────────┼────────────┘
                                                     │ HTTP :8080
                                                     ▼
┌─────────────────────────────────────────────────────────────────┐
│                     Backend (Go + Gin)                          │
│                     apps/backend                                │
│                                                                 │
│  ┌──────────────┐  ┌──────────────┐  ┌────────────────────┐    │
│  │  Middleware  │  │   Handlers   │  │     Services       │    │
│  │ (CORS, auth) │  │  (HTTP layer)│  │  (business logic)  │    │
│  └──────────────┘  └──────────────┘  └─────────┬──────────┘    │
│                                                 │               │
│                                      ┌──────────▼──────────┐   │
│                                      │    Repositories     │   │
│                                      │   (data access)     │   │
│                                      └──────────┬──────────┘   │
└──────────────────────────────────────────────────┼─────────────┘
                                                   │ pgx :5432
                                                   ▼
┌─────────────────────────────────────────────────────────────────┐
│                        PostgreSQL 16                            │
│                        (Docker container)                       │
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│                     Shared Libraries                            │
│                     libs/shared/                                │
│                                                                 │
│  ┌──────────────┐  ┌──────────────┐  ┌────────────────────┐    │
│  │    types     │  │    utils     │  │      config        │    │
│  │ (interfaces) │  │  (helpers)   │  │   (env/endpoints)  │    │
│  └──────────────┘  └──────────────┘  └────────────────────┘    │
│          ↑                ↑                    ↑                │
│          └────────────────┴────────────────────┘               │
│               Used by frontend & (TypeScript) shared           │
└─────────────────────────────────────────────────────────────────┘
```

## Request Flow

```
Browser → Next.js (SSR/CSR) → Axios → Gin Router → Handler → Service → Repository → PostgreSQL
                                                       ↓
                                                  Middleware
                                              (CORS, logging, auth)
```

## Directory Structure

```
white-label-repository/
├── apps/
│   ├── backend/                 # Go application
│   │   ├── cmd/server/          # Entry point
│   │   └── internal/
│   │       ├── config/          # Configuration
│   │       ├── database/        # Connection & migrations
│   │       ├── handlers/        # HTTP handlers
│   │       ├── middleware/      # CORS, auth, logging
│   │       ├── models/          # Data models
│   │       ├── repositories/    # DB queries
│   │       └── services/        # Business logic
│   └── frontend/                # Next.js application
│       └── src/
│           ├── app/             # App Router (pages & layouts)
│           ├── components/
│           │   └── ui/          # shadcn/ui components
│           ├── hooks/           # React hooks
│           ├── lib/             # Utilities & API client
│           └── types/           # TypeScript types
├── libs/
│   └── shared/
│       ├── types/               # Shared type definitions
│       ├── utils/               # Shared utility functions
│       └── config/              # Shared configuration
├── docs/                        # Documentation
├── docker-compose.yml           # Local dev environment
├── nx.json                      # Nx workspace config
└── package.json                 # Root dependencies & scripts
```

## Technology Choices

### Backend – Go + Gin

- **Go** chosen for performance and simplicity
- **Gin** web framework for routing and middleware
- **pgx v5** for PostgreSQL – idiomatic, performant driver
- **godotenv** for `.env` file loading in development

### Frontend – Next.js + shadcn/ui

- **Next.js 14** with App Router for SSR/SSG capabilities
- **shadcn/ui** for accessible, composable UI components
- **Tailwind CSS** with CSS variables for themeable design
- **next-themes** for dark/light mode without flash

### Monorepo – Nx

- **Nx** provides task caching, dependency graph, and project orchestration
- Shared libraries are referenced via TypeScript path aliases
- Each project runs independently or together via `nx run-many`

## Data Flow

### API Request Example (List Users)

```
1. Browser requests /api/users
2. apiClient.get('/api/users') sends GET http://localhost:8080/api/users
3. Gin router matches GET /api/users → userHandler.list
4. CORS middleware validates origin
5. userHandler calls userService.GetAll(ctx)
6. userService calls userRepository.FindAll(ctx)
7. Repository executes SQL: SELECT id, email, name, ... FROM users
8. Result propagates back as JSON: { success: true, data: [...] }
```

## Security Considerations

- All secrets via environment variables (never hardcoded)
- CORS configured to allow only specific origins
- JWT for authenticated endpoints (extensible)
- PostgreSQL connection uses connection pool with limits
- Docker containers run as non-root users
- Production builds use Go's `CGO_ENABLED=0` for smaller attack surface
