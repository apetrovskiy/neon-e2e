/* eslint-disable unicorn/prevent-abbreviations */
const baseUrl = process.env.BASE_URL || `https://neonswap.live/`;
const chromeDriverPath = require('chromedriver').path;
const geckoDriverPath = require('geckodriver').path;
const seleniumServerPath = require('selenium-server').path;

const chrome = {
  acceptSslCerts: true,
  applicationCacheEnabled: true,
  browserName: 'chrome',
  databaseEnabled: true,
  javascriptEnabled: true,
  loggingPrefs: { browser: 'ALL', driver: 'ALL', server: 'ALL' },
  w3c: false,
  webStorageEnabled: true
};

// usage:
/*
    default: {
      desiredCapabilities: {
      ...chrome,
      chromeOptions: headlessChromeOptions
      },
*/

const firefox = {
  acceptSslCerts: true,
  browserName: 'firefox',
  javascriptEnabled: true,
  mationette: true,
  'moz:firefoxOptions': {
    args: ['-headless']
  }
};

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const headlessChromeOptions = {
  args: ['headless', 'no-sandbox', 'disable-gpu']
};

module.exports = {
  //custom_commands_path: 'build/helpers/commands',  // Uncomment if you are going to use custom commands
  globals_path: 'build/src/globals.js',
  output_folder: 'output/reports',
  //page_objects_path: 'build/pages', // Uncomment if you are going to use pages
  selenium: {
    cli_args: {
      'webdriver.chrome.driver': chromeDriverPath,
      'webdriver.gecko.driver': geckoDriverPath
    },
    host: 'localhost',
    port: 4444, // standard selenium port
    server_path: seleniumServerPath,
    start_process: true // tells nightwatch to start/stop the selenium process
  },
  skip_testcases_on_fail: false,
  src_folders: ['build/src/tests'],
  test_settings: {
    chrome: {
      desiredCapabilities: {
        ...chrome,
        chromeOptions: {
          args: ['--window-size=1920,1280']
        }
      },
      webdriver: {
        start_process: true
      }
    },
    default: {
      desiredCapabilities: firefox,
      globals: {
        // for before/after hooks and variables, see src/globals.ts
      },
      launch_url: baseUrl,
      screenshots: {
        enabled: true,
        on_error: true,
        on_failure: true,
        path: 'output/screenshots'
      },
      silent: true,
      webdriver: {
        start_process: true
      }
    },
    firefox: {
      desiredCapabilities: {
        alwaysMatch: {
          'moz:firefoxOptions': {
            args: ['-headless']
          }
        },
        browserName: 'firefox'
      },
      webdriver: {
        host: 'localhost',
        port: 4444,
        server_path: 'node_modules/.bin/geckodriver',
        start_process: true
      }
    },
    safari: {
      desiredCapabilities: {
        acceptSslCerts: true,
        browserName: 'safari',
        javascriptEnabled: true,
        w3c: false
      },
      webdriver: {
        port: 4445,
        server_path: '/usr/bin/safaridriver'
      }
    }
  },
  test_workers: {
    enabled: false,
    workers: 'auto'
  }
};
