import { NightwatchBrowser, NightwatchTests } from 'nightwatch';

import { NeonSwapMessages } from '../../constants/neon-swap';

// import { URL } from '../constants/neon-swap ';

const home: NightwatchTests = {
  'Shows the Connect wallet button': (browser: NightwatchBrowser) => {
    browser.url(NeonSwapMessages.URL).assert.visible('#connect-wallet').end();
    // #connect-wallet
  },
  'Neon Swap 2222 temp !!! 2222': (browser: NightwatchBrowser) => {
    browser.url(NeonSwapMessages.URL).assert.visible('#connect-wallet').end();
    // #connect-wallet
  }
};
export default home;
