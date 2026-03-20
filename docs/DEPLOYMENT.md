# Deployment Guide

This guide covers deploying the white-label monorepo to production environments.

## Overview

The monorepo contains:
- **Backend** – Go binary (statically compiled)
- **Frontend** – Next.js application
- **Database** – PostgreSQL

---

## 1. Docker Compose (Simple / VPS)

The easiest path for single-server deployments.

### Build & Start

```bash
# Copy and configure environment
cp .env.example .env
vim .env   # Set DB passwords, JWT secret, etc.

# Build and start all services
docker-compose up -d --build

# Check logs
docker-compose logs -f

# Check service health
curl http://localhost:8080/api/health
```

### Update Deployment

```bash
git pull
docker-compose up -d --build
```

---

## 2. Manual (Backend)

### Build Go Binary

```bash
cd apps/backend
CGO_ENABLED=0 GOOS=linux go build -o dist/server ./cmd/server

# Run
./dist/server
```

### Systemd Service (Linux)

```ini
# /etc/systemd/system/whitelabel-backend.service
[Unit]
Description=White Label Backend
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/whitelabel
ExecStart=/opt/whitelabel/server
EnvironmentFile=/opt/whitelabel/.env
Restart=always

[Install]
WantedBy=multi-user.target
```

```bash
systemctl enable whitelabel-backend
systemctl start whitelabel-backend
```

---

## 3. Manual (Frontend)

### Build Next.js

```bash
cd apps/frontend
npm ci
npm run build
npm run start   # Starts on port 3000
```

### Process Manager (PM2)

```bash
npm install -g pm2
cd apps/frontend
pm2 start npm --name "frontend" -- start
pm2 save
pm2 startup
```

---

## 4. Cloud Platforms

### Vercel (Frontend)

1. Connect your GitHub repository to Vercel
2. Set **Root Directory** to `apps/frontend`
3. Add environment variables:
   - `NEXT_PUBLIC_API_URL` → your backend URL
   - `NEXT_PUBLIC_APP_NAME` → your app name
4. Deploy

### Railway / Render (Backend)

1. Create a new service pointed at the repository
2. Set **Root Directory** to `apps/backend`
3. Build command: `go build -o server ./cmd/server`
4. Start command: `./server`
5. Add all environment variables from `.env.example`

### AWS / GCP / Azure (Docker)

```bash
# Build and push backend image
docker build -t your-registry/backend:latest apps/backend
docker push your-registry/backend:latest

# Build and push frontend image
docker build -t your-registry/frontend:latest apps/frontend
docker push your-registry/frontend:latest
```

Use Kubernetes, ECS, Cloud Run, or similar to deploy.

---

## 5. Environment Variables for Production

| Variable              | Production Requirement                        |
|-----------------------|-----------------------------------------------|
| `DB_PASSWORD`         | Strong random password                        |
| `JWT_SECRET`          | 32+ character random string                   |
| `DB_SSL_MODE`         | Set to `require` for managed databases        |
| `CORS_ALLOWED_ORIGINS`| Your production frontend URL                  |
| `APP_ENV`             | Set to `production`                           |
| `NEXT_PUBLIC_API_URL` | Full HTTPS backend URL                        |

Generate a strong JWT secret:

```bash
openssl rand -hex 32
```

---

## 6. Database

### PostgreSQL Managed Services

Compatible with:
- **AWS RDS** for PostgreSQL
- **Google Cloud SQL**
- **Azure Database for PostgreSQL**
- **Supabase**
- **Neon**

Set `DB_SSL_MODE=require` when connecting to managed databases.

### Backups

```bash
# Backup
pg_dump -U postgres whitelabel_db > backup_$(date +%Y%m%d).sql

# Restore
psql -U postgres whitelabel_db < backup_20240101.sql
```

---

## 7. Nginx Reverse Proxy (Optional)

```nginx
# /etc/nginx/sites-available/whitelabel
server {
    listen 80;
    server_name yourdomain.com;
    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl;
    server_name yourdomain.com;

    ssl_certificate /etc/letsencrypt/live/yourdomain.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/yourdomain.com/privkey.pem;

    # Frontend
    location / {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }

    # Backend API
    location /api/ {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

---

## 8. Health Checks

```bash
# Backend health
curl https://api.yourdomain.com/api/health

# Expected response
{
  "status": "healthy",
  "timestamp": "...",
  "version": "1.0.0",
  "services": {
    "api": "healthy",
    "database": "healthy"
  }
}
```

---

## 9. CI/CD (GitHub Actions)

Create `.github/workflows/deploy.yml`:

```yaml
name: Deploy

on:
  push:
    branches: [main]

jobs:
  deploy-backend:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Build and push Docker image
        run: |
          docker build -t backend apps/backend
          # Push to your registry

  deploy-frontend:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: 20
      - run: cd apps/frontend && npm ci && npm run build
      # Deploy to Vercel/Netlify/etc
```
