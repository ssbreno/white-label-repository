const nextJest = require('next/jest');

const createJestConfig = nextJest({ dir: './' });

/** @type {import('jest').Config} */
const config = {
  displayName: 'frontend',
  testEnvironment: 'jsdom',
  moduleNameMapper: {
    '^@/(.*)$': '<rootDir>/src/$1',
    '^@white-label/shared-types$': '<rootDir>/../../libs/shared/types/src/index.ts',
    '^@white-label/shared-utils$': '<rootDir>/../../libs/shared/utils/src/index.ts',
    '^@white-label/shared-config$': '<rootDir>/../../libs/shared/config/src/index.ts',
  },
  collectCoverageFrom: [
    'src/**/*.{ts,tsx}',
    '!src/**/*.d.ts',
    '!src/components/ui/**',
  ],
  testPathIgnorePatterns: ['<rootDir>/.next/', '<rootDir>/node_modules/'],
};

module.exports = createJestConfig(config);
