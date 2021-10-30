/* eslint-disable @typescript-eslint/no-unsafe-call */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
/* eslint-disable @typescript-eslint/no-unsafe-return */
import { unit } from 'src/constants/unit';
import { Account } from 'web3-core';

import { Web3Helper } from './web3-helper';

export class Balance {
  async getBalance(account: Account): Promise<string> {
    return await Web3Helper.getWeb3().eth.getBalance(account.address);
  }
  convertToEther(balance: string): string {
    return Web3Helper.getWeb3().utils.fromWei(balance, unit.ether);
  }
}
