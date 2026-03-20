/**
 * Shared configuration helpers
 * Import: import { getApiEndpoints, getEnvVar } from '@white-label/shared-config';
 */

// ─────────────────────────────────────────────────────────
// API Endpoints
// ─────────────────────────────────────────────────────────

export interface APIEndpoints {
  health: string;
  users: {
    list: string;
    detail: (id: string) => string;
    create: string;
    update: (id: string) => string;
    delete: (id: string) => string;
  };
}

/**
 * Build typed API endpoint paths for a given base URL
 */
export function getApiEndpoints(baseUrl: string): APIEndpoints {
  const base = baseUrl.replace(/\/$/, '');
  return {
    health: `${base}/api/health`,
    users: {
      list: `${base}/api/users`,
      detail: (id: string) => `${base}/api/users/${id}`,
      create: `${base}/api/users`,
      update: (id: string) => `${base}/api/users/${id}`,
      delete: (id: string) => `${base}/api/users/${id}`,
    },
  };
}

// ─────────────────────────────────────────────────────────
// Environment Schema
// ─────────────────────────────────────────────────────────

export interface EnvSchema {
  required: string[];
  optional: string[];
}

/** Frontend environment variable schema */
export const FRONTEND_ENV_SCHEMA: EnvSchema = {
  required: [],
  optional: ['NEXT_PUBLIC_API_URL', 'NEXT_PUBLIC_APP_NAME'],
};

/** Backend environment variable schema */
export const BACKEND_ENV_SCHEMA: EnvSchema = {
  required: ['DB_HOST', 'DB_NAME', 'DB_USER', 'DB_PASSWORD'],
  optional: [
    'APP_ENV',
    'BACKEND_PORT',
    'DB_PORT',
    'DB_SSL_MODE',
    'JWT_SECRET',
    'JWT_EXPIRATION',
    'CORS_ALLOWED_ORIGINS',
  ],
};

/**
 * Safely read an environment variable with a fallback
 * Works in both Node.js and browser environments
 */
export function getEnvVar(key: string, fallback = ''): string {
  if (typeof process !== 'undefined' && process.env) {
    return process.env[key] ?? fallback;
  }
  return fallback;
}

/**
 * Validate that all required environment variables are set
 * Returns a list of missing variable names
 */
export function validateEnv(schema: EnvSchema): string[] {
  return schema.required.filter(
    (key) => !getEnvVar(key)
  );
}

// ─────────────────────────────────────────────────────────
// App Defaults
// ─────────────────────────────────────────────────────────

export const APP_DEFAULTS = {
  name: 'White Label App',
  version: '1.0.0',
  apiUrl: 'http://localhost:8080',
  frontendUrl: 'http://localhost:3000',
} as const;
