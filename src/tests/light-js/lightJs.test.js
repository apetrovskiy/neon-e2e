import Api from '@parity/api';
import light, { balanceOf$ } from '@parity/light.js';
import { skip, suite, test } from '@testdeck/mocha';
import { epic, feature } from 'allure-decorators';
import { expect } from 'chai';

import { logger } from '../../utils/logger';

const EpicName = 'light.js';
const SuiteName = EpicName;

/*
// Wait for web3.currentProvider to be injected
window.addEventListener('load', () => {
  // Tell light.js to use MetaMask provider
  light.setProvider(window.web3.currentProvider);
  balanceOf$('0x123').subscribe(console.log); // Logs the balance
});
*/

@suite(SuiteName)
class LightJsTests {
  @epic(EpicName)
  @feature('Account test')
  @test
  public async shouldConnectUsingLightJs() {
    const provider = new Api.Provider.Http(process.env.PROXY_URL);
    light.setProvider(provider);
    const api = new Api(provider);
    expect(api, 'Should not be null').to.not.be.null;
    logger.notice(`api = ${api}`);
  }
}
