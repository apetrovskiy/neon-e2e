import { NightwatchBrowser } from 'nightwatch';

module.exports = {
  afterEach: (_browser: never, done: () => void) => {
    done();
  },
  beforeEach: (browser: NightwatchBrowser, done: () => void) => {
    done();
  }
};
