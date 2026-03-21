# White Label Repository

[![CI](https://github.com/ssbreno/white-label-repository/actions/workflows/ci.yml/badge.svg)](https://github.com/ssbreno/white-label-repository/actions/workflows/ci.yml)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

A production-ready, open-source white label monorepo built with **Nx**, featuring a **Go backend** with PostgreSQL and a **Next.js frontend** with shadcn/ui and Tailwind CSS.

## Architecture

```
white-label-repository/
├── apps/
│   ├── backend/          # Go API server (Gin + PostgreSQL)
│   └── frontend/         # Next.js 14 + shadcn/ui + Tailwind CSS
├── libs/
│   └── shared/
│       ├── types/        # Shared TypeScript type definitions
│       ├── utils/        # Shared utility functions
│       └── config/       # Shared configuration
├── docs/
│   ├── ARCHITECTURE.md   # System architecture
│   ├── CUSTOMIZATION.md  # White label customization guide
│   └── DEPLOYMENT.md     # Deployment guidelines
├── docker-compose.yml    # Local development environment
├── nx.json               # Nx workspace configuration
└── package.json          # Root workspace package
```

## Tech Stack

| Layer      | Technology                                  |
|------------|---------------------------------------------|
| Backend    | Go 1.21+, Gin, PostgreSQL, pgx             |
| Frontend   | Next.js 15, React 18, TypeScript 5          |
| UI         | shadcn/ui, Tailwind CSS                     |
| AI Gateway | OpenAI, Anthropic, Google (unified API)     |
| Monorepo   | Nx 20+                                      |
| Database   | PostgreSQL 16                               |
| Testing    | Go test + testify, Jest + Testing Library   |
| CI/CD      | GitHub Actions                              |
| Container  | Docker, Docker Compose                      |

## Quick Start

### Prerequisites

- [Node.js](https://nodejs.org/) 20+
- [Go](https://golang.org/) 1.21+
- [Docker](https://www.docker.com/) & Docker Compose
- [npm](https://www.npmjs.com/) 10+

### 1. Clone and Install

```bash
git clone https://github.com/ssbreno/white-label-repository.git
cd white-label-repository
npm install
```

### 2. Configure Environment

```bash
cp .env.example .env
# Edit .env with your configuration
```

### 3. Start with Docker Compose (Recommended)

```bash
# Start all services (PostgreSQL, Backend, Frontend)
docker-compose up -d

# View logs
docker-compose logs -f
```

### 4. Development Mode (Without Docker)

```bash
# Start PostgreSQL (requires Docker)
docker-compose up postgres -d

# Start backend
npm run dev:backend

# Start frontend (in another terminal)
npm run dev:frontend

# Start all services
npm run dev
```

## Available Commands

Run `make help` for the full list. Key commands:

| Command            | Description                          |
|--------------------|--------------------------------------|
| `make dev`         | Run all services in development mode |
| `make build`       | Build all applications               |
| `make test`        | Run all tests (backend + frontend)   |
| `make lint`        | Lint all projects                    |
| `make format`      | Format code with Prettier            |
| `make docker-up`   | Start Docker services                |
| `make docker-down` | Stop Docker services                 |
| `make db-shell`    | Open PostgreSQL shell                |
| `make graph`       | Visualize Nx dependency graph        |

## Development Workflow

```
1. Make changes to your code
2. Tests run automatically (with Nx caching)
3. Open PR → CI runs linting, tests, and builds
4. Merge → Deploy to staging/production
```

## White Label Customization

This repository is designed as a template for multiple client projects. See [docs/CUSTOMIZATION.md](docs/CUSTOMIZATION.md) for the full guide.

**Quick customization:**

1. **Colors & Branding** – Edit `apps/frontend/src/app/globals.css` (CSS variables)
2. **App Name** – Update `APP_NAME` in `.env` and `apps/frontend/src/app/layout.tsx`
3. **API Endpoints** – Configure `libs/shared/config/src/endpoints.ts`
4. **Database Schema** – Modify `apps/backend/migrations/`

## Documentation

- [Architecture Overview](docs/ARCHITECTURE.md)
- [Customization Guide](docs/CUSTOMIZATION.md)
- [Deployment Guide](docs/DEPLOYMENT.md)
- [Claude Code Guide](CLAUDE.md)
- [Contributing Guide](CONTRIBUTING.md)
- [Code of Conduct](CODE_OF_CONDUCT.md)
- [Backend README](apps/backend/README.md)
- [Frontend README](apps/frontend/README.md)

## Project Structure Details

### Backend (`apps/backend`)

```
apps/backend/
├── cmd/server/         # Application entry point
├── internal/
│   ├── ai/             # AI provider gateway (OpenRouter-style)
│   ├── config/         # App configuration
│   ├── database/       # DB connection & migrations
│   ├── handlers/       # HTTP request handlers
│   ├── middleware/      # HTTP middleware (CORS, rate limiting)
│   ├── models/         # Data models
│   ├── mocks/          # Test mocks
│   ├── repositories/   # Data access layer
│   ├── services/       # Business logic
│   └── testutil/       # Test helpers
├── Dockerfile
├── Makefile
├── .golangci.yml       # Go linting config
├── go.mod
└── README.md
```

### Frontend (`apps/frontend`)

```
apps/frontend/
├── src/
│   ├── app/            # Next.js App Router
│   ├── components/
│   │   └── ui/         # shadcn/ui components
│   ├── hooks/          # Custom React hooks
│   ├── lib/            # Utilities & API client
│   └── types/          # TypeScript types
├── public/             # Static assets
├── components.json     # shadcn/ui config
├── tailwind.config.ts
└── README.md
```

### Shared Libraries (`libs/shared`)

| Library       | Description                         |
|---------------|-------------------------------------|
| `shared-types`  | TypeScript interfaces & types     |
| `shared-utils`  | Date, validation, string helpers  |
| `shared-config` | API endpoints, env schemas        |

## AI Gateway

The backend includes an OpenRouter-style AI gateway that provides a unified API for multiple AI providers.

### Supported Providers

| Provider  | Models                                    |
|-----------|-------------------------------------------|
| OpenAI    | GPT-4o, GPT-4o Mini, GPT-4 Turbo, GPT-3.5 |
| Anthropic | Claude Sonnet 4, Claude 3.5 Sonnet/Haiku  |
| Google    | Gemini 2.0 Flash, Gemini 1.5 Pro/Flash    |

### API Endpoints

```
POST /api/ai/chat       # Chat completion (sync or streaming via SSE)
GET  /api/ai/models     # List all available models
GET  /api/ai/providers  # List enabled providers
```

Providers auto-enable when their API key is configured in `.env`. Streaming is supported via Server-Sent Events (SSE) by setting `"stream": true` in the request.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for the full guide.

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Commit changes: `git commit -m 'feat: add my feature'`
4. Push branch: `git push origin feature/my-feature`
5. Open a Pull Request

## License

GNU General Public License v3.0 - see [LICENSE](LICENSE) for details.
