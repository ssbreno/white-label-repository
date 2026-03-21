export const siteConfig = {
  name: process.env.NEXT_PUBLIC_APP_NAME || 'White Label App',
  description:
    process.env.NEXT_PUBLIC_APP_DESCRIPTION ||
    'A production-ready white-label template with Go backend, Next.js frontend, and shared libraries — ready to customize for any project.',
  url: process.env.NEXT_PUBLIC_SITE_URL || 'http://localhost:3000',
  ogImage: process.env.NEXT_PUBLIC_OG_IMAGE || '/og-image.png',
  locale: process.env.NEXT_PUBLIC_LOCALE || 'en_US',
  twitterHandle: process.env.NEXT_PUBLIC_TWITTER_HANDLE || '',
  author: process.env.NEXT_PUBLIC_AUTHOR || 'White Label Team',
  keywords: (process.env.NEXT_PUBLIC_KEYWORDS || 'white-label,template,monorepo,nextjs,go,typescript')
    .split(',')
    .map((k) => k.trim()),
  themeColor: process.env.NEXT_PUBLIC_THEME_COLOR || '#2563eb',
};
