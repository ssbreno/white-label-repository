-- Initial database setup
-- This file is auto-executed by PostgreSQL on first startup via docker-entrypoint-initdb.d

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Users table (created by migrations, this is for documentation only)
-- See internal/database/migrate.go for the actual migration code
