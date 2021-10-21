import { NightwatchBrowser, NightwatchTests } from 'nightwatch';

const home: NightwatchTests = {
  '01 test 01': (browser: NightwatchBrowser) => {
    browser.url('http://neonswap.live').pause(5000).end();
  }
};
export default home;
