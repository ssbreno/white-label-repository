# Frontend – Next.js + shadcn/ui + Tailwind CSS

Modern, themeable React application built with Next.js 14, shadcn/ui, and Tailwind CSS.

## Tech Stack

- **Next.js 14** with App Router
- **React 18** with TypeScript 5
- **Tailwind CSS 3** with CSS variables theming
- **shadcn/ui** component library
- **next-themes** for dark/light mode
- **Axios** for API communication

## Project Structure

```
apps/frontend/
├── src/
│   ├── app/                  # Next.js App Router pages
│   │   ├── globals.css       # Global styles & CSS variables (theme)
│   │   ├── layout.tsx        # Root layout
│   │   └── page.tsx          # Home page
│   ├── components/
│   │   ├── ui/               # shadcn/ui components
│   │   │   ├── badge.tsx
│   │   │   ├── button.tsx
│   │   │   ├── card.tsx
│   │   │   └── input.tsx
│   │   ├── theme-provider.tsx
│   │   └── theme-toggle.tsx
│   ├── hooks/                # Custom React hooks
│   ├── lib/
│   │   ├── api.ts            # API client (axios)
│   │   └── utils.ts          # Utility functions (cn)
│   └── types/                # TypeScript types
├── public/                   # Static assets
├── components.json           # shadcn/ui configuration
├── tailwind.config.ts        # Tailwind configuration
├── next.config.js            # Next.js configuration
└── tsconfig.json             # TypeScript configuration
```

## Quick Start

### Prerequisites

- Node.js 20+
- npm 10+

### 1. Install Dependencies

```bash
cd apps/frontend
npm install
```

### 2. Configure Environment

```bash
cp .env.example .env.local
# Edit .env.local
```

### 3. Start Development Server

```bash
# From workspace root
npm run dev:frontend

# Or directly
cd apps/frontend
npm run dev
```

The app runs on `http://localhost:3000`.

## Available Scripts

| Script          | Description              |
|-----------------|--------------------------|
| `npm run dev`   | Start development server |
| `npm run build` | Production build         |
| `npm run start` | Start production server  |
| `npm run lint`  | Run ESLint               |

## Theming (White Label)

All colors are CSS variables in `src/app/globals.css`. To rebrand:

```css
:root {
  /* Change primary color (brand color) */
  --primary: 262.1 83.3% 57.8%;  /* e.g., purple */
  --primary-foreground: 210 20% 98%;

  /* Change border radius */
  --radius: 0.75rem;
}
```

### Adding shadcn/ui Components

```bash
# Install a new component
npx shadcn@latest add <component-name>

# Examples
npx shadcn@latest add dialog
npx shadcn@latest add table
npx shadcn@latest add form
```

## Environment Variables

| Variable                 | Default                     | Description             |
|--------------------------|-----------------------------|-------------------------|
| `NEXT_PUBLIC_API_URL`    | `http://localhost:8080`     | Backend API URL         |
| `NEXT_PUBLIC_APP_NAME`   | `White Label App`           | Application display name |
