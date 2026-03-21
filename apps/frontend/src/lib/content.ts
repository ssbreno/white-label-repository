export const landingContent = {
  hero: {
    title: process.env.NEXT_PUBLIC_HERO_TITLE || 'White Label Monorepo',
    subtitle:
      process.env.NEXT_PUBLIC_HERO_SUBTITLE ||
      'A production-ready template with Go backend, Next.js frontend, and shared libraries — ready to customize for any project.',
    ctaPrimary: 'Get Started',
    ctaPrimaryHref: 'https://github.com/ssbreno/white-label-repository',
    ctaSecondary: 'View Docs',
    ctaSecondaryHref: '/docs',
  },
  features: [
    {
      title: 'Go Backend',
      description: 'High-performance REST API',
      details:
        'Built with Gin framework, PostgreSQL via pgx, structured logging, and graceful shutdown support.',
      badges: ['Go 1.21+', 'Gin', 'PostgreSQL'],
    },
    {
      title: 'Next.js Frontend',
      description: 'Modern React application',
      details:
        'Next.js 15 with App Router, TypeScript, Tailwind CSS, and shadcn/ui component library with dark mode support.',
      badges: ['Next.js 15', 'TypeScript', 'Tailwind'],
    },
    {
      title: 'AI Gateway',
      description: 'Multi-provider AI integration',
      details:
        'OpenRouter-style gateway supporting OpenAI, Anthropic, and Google with unified API, streaming, and rate limiting.',
      badges: ['OpenAI', 'Anthropic', 'Google'],
    },
    {
      title: 'Shared Libraries',
      description: 'Reusable code across projects',
      details:
        'Shared TypeScript types, utility functions, and configuration that can be imported by any app in the monorepo.',
      badges: ['Types', 'Utils', 'Config'],
    },
    {
      title: 'Docker Ready',
      description: 'Containerized development',
      details:
        'Docker Compose setup with PostgreSQL, backend, and frontend services. Multi-stage builds for production.',
      badges: ['Docker', 'Compose', 'Multi-stage'],
    },
    {
      title: 'Developer Experience',
      description: 'Quality tools built-in',
      details:
        'ESLint, Prettier, golangci-lint, Jest, Go testing, CI/CD with GitHub Actions, and Claude Code integration.',
      badges: ['ESLint', 'Jest', 'CI/CD'],
    },
  ],
  faq: [
    {
      question: 'What is this white-label template?',
      answer:
        'It is a production-ready monorepo starter kit that provides a Go REST API backend, a Next.js frontend with shadcn/ui, shared TypeScript libraries, Docker setup, CI/CD pipelines, and an AI provider gateway — all pre-configured and ready to customize for your product.',
    },
    {
      question: 'How do I customize the branding?',
      answer:
        'All branding is driven by environment variables and CSS variables. Update colors in globals.css, set your app name via NEXT_PUBLIC_APP_NAME, and configure the theme. No code changes needed for basic customization.',
    },
    {
      question: 'Which AI providers are supported?',
      answer:
        'The AI gateway supports OpenAI (GPT-4, GPT-3.5), Anthropic (Claude), and Google (Gemini) with a unified API. Providers auto-enable when their API key is configured. Streaming via SSE is supported for all providers.',
    },
    {
      question: 'How do I add a new API resource?',
      answer:
        'Follow the 6-step pattern: create a model, repository, service, handler, register routes, and add a database migration. Each layer has a clear interface. See CLAUDE.md for the detailed workflow.',
    },
    {
      question: 'Is this suitable for production use?',
      answer:
        'Yes. It includes connection pooling, graceful shutdown, health checks, rate limiting, multi-stage Docker builds with non-root users, environment-based configuration, and CI/CD pipelines. Review the security checklist in docs/ARCHITECTURE.md before deploying.',
    },
  ],
};
