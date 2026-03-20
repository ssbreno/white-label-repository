# White Label Repository

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
| Frontend   | Next.js 14+, React 18, TypeScript 5         |
| UI         | shadcn/ui, Tailwind CSS                     |
| Monorepo   | Nx 20+                                      |
| Database   | PostgreSQL 16                               |
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

## Available Scripts

| Script             | Description                          |
|--------------------|--------------------------------------|
| `npm run dev`      | Run all services in development mode |
| `npm run dev:frontend` | Run only the frontend            |
| `npm run dev:backend`  | Run only the backend             |
| `npm run build`    | Build all applications               |
| `npm run test`     | Run all tests                        |
| `npm run lint`     | Lint all projects                    |
| `npm run graph`    | Visualize project dependency graph   |

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
- [Backend README](apps/backend/README.md)
- [Frontend README](apps/frontend/README.md)

## Project Structure Details

### Backend (`apps/backend`)

```
apps/backend/
├── cmd/server/         # Application entry point
├── internal/
│   ├── config/         # App configuration
│   ├── database/       # DB connection & migrations
│   ├── handlers/       # HTTP request handlers
│   ├── middleware/     # HTTP middleware (CORS, auth, logging)
│   ├── models/         # Data models
│   ├── repositories/   # Data access layer
│   └── services/       # Business logic
├── migrations/         # SQL migrations
├── Dockerfile
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

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Commit changes: `git commit -m 'feat: add my feature'`
4. Push branch: `git push origin feature/my-feature`
5. Open a Pull Request

## License

GNU General Public License v3.0 - see [LICENSE](LICENSE) for details.
