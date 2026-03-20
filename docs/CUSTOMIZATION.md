# White Label Customization Guide

This guide explains how to fork and customise this repository for a new client project.

## Quick Checklist

When creating a new project from this template:

- [ ] Fork or copy the repository
- [ ] Update `APP_NAME` in `.env` / environment
- [ ] Change brand colors in `apps/frontend/src/app/globals.css`
- [ ] Replace logo/favicon in `apps/frontend/public/`
- [ ] Update `metadata` in `apps/frontend/src/app/layout.tsx`
- [ ] Configure API endpoints in `libs/shared/config/src/index.ts`
- [ ] Update database name in `.env`
- [ ] Review and enable/disable feature flags

---

## 1. Branding

### Colors

Edit `apps/frontend/src/app/globals.css` and change the CSS variables:

```css
:root {
  /* Primary brand color – most important change */
  --primary: 262.1 83.3% 57.8%;       /* purple example */
  --primary-foreground: 210 20% 98%;

  /* Accent color */
  --accent: 262.1 83.3% 95%;
  --accent-foreground: 262.1 83.3% 30%;

  /* Border radius – change visual style */
  --radius: 0.75rem;   /* rounded */
  /* --radius: 0rem;   sharp corners */
}
```

Use HSL format. Tools: [HSL Picker](https://hslpicker.com/), [Tailwind Colors](https://tailwindcss.com/docs/customization/colors).

### Fonts

1. Add font import in `apps/frontend/src/app/layout.tsx`:

```tsx
import { Poppins } from 'next/font/google';
const poppins = Poppins({ weight: ['400', '600', '700'], subsets: ['latin'], variable: '--font-sans' });
```

2. The `--font-sans` CSS variable is already configured in `tailwind.config.ts`.

### Logo / Favicon

Replace files in `apps/frontend/public/`:
- `favicon.ico` – browser tab icon
- `logo.svg` or `logo.png` – app logo

Reference in layout:

```tsx
import Image from 'next/image';
<Image src="/logo.svg" alt="My Brand" width={120} height={40} />
```

### App Name

```bash
# .env (or environment)
NEXT_PUBLIC_APP_NAME=My Client App
APP_NAME=my-client-app
```

---

## 2. API Configuration

### Endpoints

Edit `libs/shared/config/src/index.ts` to add or change endpoints:

```typescript
export function getApiEndpoints(baseUrl: string): APIEndpoints {
  const base = baseUrl.replace(/\/$/, '');
  return {
    health: `${base}/api/health`,
    users: { ... },
    // Add new resources:
    products: {
      list: `${base}/api/products`,
      detail: (id: string) => `${base}/api/products/${id}`,
    },
  };
}
```

### Backend URL

```bash
# Frontend .env.local
NEXT_PUBLIC_API_URL=https://api.myclient.com

# docker-compose.yml
NEXT_PUBLIC_API_URL: http://backend:8080
```

---

## 3. Database

### Name

```bash
# .env
DB_NAME=myclient_db
```

### Schema Changes

Add new migration files to `apps/backend/migrations/`:

```sql
-- migrations/002_add_products.sql
CREATE TABLE products (
  id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name       VARCHAR(255) NOT NULL,
  price      NUMERIC(10, 2) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

Update `apps/backend/internal/database/migrate.go` to include the new migration.

---

## 4. Feature Flags

Add feature flags to `.env`:

```bash
FEATURE_AUTH_ENABLED=true
FEATURE_PAYMENTS_ENABLED=false
FEATURE_DARK_MODE=true
```

Read in frontend:

```tsx
const isAuthEnabled = process.env.NEXT_PUBLIC_FEATURE_AUTH === 'true';
```

---

## 5. Adding New Backend Routes

1. **Model** – add to `apps/backend/internal/models/`
2. **Repository** – add to `apps/backend/internal/repositories/`
3. **Service** – add to `apps/backend/internal/services/`
4. **Handler** – add to `apps/backend/internal/handlers/`
5. **Route** – register in `apps/backend/internal/handlers/routes.go`

---

## 6. Adding New Frontend Pages

1. Create `apps/frontend/src/app/<page>/page.tsx`
2. Add navigation link in the header component
3. Add API calls using `src/lib/api.ts`

---

## 7. Adding shadcn/ui Components

```bash
cd apps/frontend
npx shadcn@latest add <component>

# Examples
npx shadcn@latest add dialog
npx shadcn@latest add table
npx shadcn@latest add form
npx shadcn@latest add select
```

---

## 8. Adding to Shared Libraries

### New Type

Add to `libs/shared/types/src/index.ts`:

```typescript
export interface Product {
  id: string;
  name: string;
  price: number;
}
```

Import in frontend:

```typescript
import type { Product } from '@white-label/shared-types';
```

### New Utility

Add to `libs/shared/utils/src/index.ts`:

```typescript
export function formatCurrency(amount: number, currency = 'USD'): string {
  return new Intl.NumberFormat('en-US', { style: 'currency', currency }).format(amount);
}
```

---

## 9. Environment Reference

All configurable values:

| Variable                     | Default                   | Where Used       |
|------------------------------|---------------------------|------------------|
| `APP_NAME`                   | `white-label-app`         | General          |
| `NEXT_PUBLIC_APP_NAME`       | `White Label App`         | Frontend header  |
| `NEXT_PUBLIC_API_URL`        | `http://localhost:8080`   | Frontend API     |
| `BACKEND_PORT`               | `8080`                    | Backend server   |
| `DB_HOST`                    | `localhost`               | Backend DB       |
| `DB_PORT`                    | `5432`                    | Backend DB       |
| `DB_NAME`                    | `whitelabel_db`           | Backend DB       |
| `DB_USER`                    | `postgres`                | Backend DB       |
| `DB_PASSWORD`                | `postgres`                | Backend DB       |
| `JWT_SECRET`                 | (none)                    | Backend auth     |
| `CORS_ALLOWED_ORIGINS`       | `http://localhost:3000`   | Backend CORS     |
