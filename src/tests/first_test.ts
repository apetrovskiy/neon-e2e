import { NightwatchBrowser, NightwatchTests } from 'nightwatch';

const home: NightwatchTests = {
  'Neon Swap': (browser: NightwatchBrowser) => {
    browser.url('http://neonswap.live').pause(5000).end();
  }
};
export default home;
