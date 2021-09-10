/** @type {import('@ts-jest/dist/types').InitialOptionsTsJest} */
module.exports = {
  modulePaths: ["src/"],
  verbose: true,
  preset: "ts-jest",
  testEnvironment: "node",
  testMatch: ["**/*.steps.ts", "**/*.spec.ts"],
  testTimeout: 30000,
  testPathIgnorePatterns: ["fixtures"],
  globals: {},
  setupFilesAfterEnv: ["jest-allure/dist/setup",],
  reporters: ["default", "jest-allure"],
  testRunner: "jest-jasmine2",
  // setupFilesAfterEnv: ["<rootDir>/path/to/allure-report.js"],
};
/*
// jest.config.ts
import type {Config} from '@jest/types';

// Синхронно загружаемый конфиг
const config: Config.InitialOptions = {
  verbose: true,
};
export default config;

// Или с помощью асинхронной функции
export default async (): Promise<Config.InitialOptions> => {
  return {
    verbose: true,
  };
};
*/