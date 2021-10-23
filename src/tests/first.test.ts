import { NightwatchBrowser, NightwatchTests } from 'nightwatch';

import { NeonSwapMessages } from '../constants/neon-swap';

const home: NightwatchTests = {
  'Neon Swap': (browser: NightwatchBrowser) => {
    browser.url(NeonSwapMessages.URL).assert.visible('#connect-wallet').end();
    // #connect-wallet
  }
};
export default home;
