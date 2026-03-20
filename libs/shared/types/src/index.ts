/**
 * Shared TypeScript type definitions
 * Import: import type { User, APIResponse } from '@white-label/shared-types';
 */

// ─────────────────────────────────────────────────────────
// Entity Models
// ─────────────────────────────────────────────────────────

export interface User {
  id: string;
  email: string;
  name: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreateUserDTO {
  email: string;
  name: string;
}

export interface UpdateUserDTO {
  email?: string;
  name?: string;
}

// ─────────────────────────────────────────────────────────
// API Response Types
// ─────────────────────────────────────────────────────────

export interface APIResponse<T = unknown> {
  success: boolean;
  data?: T;
  message?: string;
  error?: string;
}

export interface PaginatedResponse<T = unknown> {
  items: T[];
  total: number;
  page: number;
  perPage: number;
  totalPages: number;
}

export interface HealthResponse {
  status: 'healthy' | 'degraded' | 'unhealthy';
  timestamp: string;
  version: string;
  services: Record<string, string>;
}

// ─────────────────────────────────────────────────────────
// Configuration Types
// ─────────────────────────────────────────────────────────

export interface AppConfig {
  name: string;
  version: string;
  environment: 'development' | 'staging' | 'production';
  apiBaseUrl: string;
}

export interface DatabaseConfig {
  host: string;
  port: number;
  name: string;
  user: string;
  sslMode: string;
}

export interface FeatureFlags {
  authEnabled: boolean;
  darkModeEnabled: boolean;
}
