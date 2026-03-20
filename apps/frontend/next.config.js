/** @type {import('next').NextConfig} */
const nextConfig = {
  // Enable React strict mode for better development experience
  reactStrictMode: true,

  // Transpile shared workspace packages
  transpilePackages: [
    '@white-label/shared-types',
    '@white-label/shared-utils',
    '@white-label/shared-config',
  ],

  // Environment variables exposed to the browser
  env: {
    NEXT_PUBLIC_API_URL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080',
    NEXT_PUBLIC_APP_NAME: process.env.NEXT_PUBLIC_APP_NAME || 'White Label App',
  },
};

module.exports = nextConfig;
