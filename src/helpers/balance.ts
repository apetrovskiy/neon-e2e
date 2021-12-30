/* eslint-disable @typescript-eslint/no-unsafe-call */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
/* eslint-disable @typescript-eslint/no-unsafe-return */
import { Account } from 'web3-core';
import { allure } from 'allure-mocha/runtime';

import { Web3Helper } from './web3-helper';

export class Balance {
  async getBalance(account: Account): Promise<string> {
    // return allure.step(`getting balance for ${account.address}`, async () => {
      return await Web3Helper.getWeb3().eth.getBalance(account.address);
    // });
  }
  convertToEther(balance: string): string {
    // return allure.step(`converting ${balance} to Ether`, () => {
      return Web3Helper.getWeb3().utils.fromWei(balance);
    // });
  }
}
