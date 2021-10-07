/* eslint-disable @typescript-eslint/no-unsafe-member-access */
/* eslint-disable @typescript-eslint/no-unsafe-call */
/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
/* eslint-disable @typescript-eslint/no-unsafe-return */
import { Config } from 'config/default';
import Web3 from 'web3';

export class Web3Helper {
  public static getWeb3() {
    return new Web3(new Web3.providers.HttpProvider(Config.url));
  }
}
