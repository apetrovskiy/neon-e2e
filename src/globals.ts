import { NightwatchBrowser } from 'nightwatch';
import { NightwatchAllureReporter } from 'nightwatch-allure';

module.exports = {
  afterEach: (_browser: never, done: () => void) => {
    done();
  },
  beforeEach: (browser: NightwatchBrowser, done: () => void) => {
    done();
  },
  reporter: (results: never, done: () => void) => {
    const reporter = new NightwatchAllureReporter({ folder: './report/allure-results', sendData: true });
    reporter.write(results, done);
  }
};
