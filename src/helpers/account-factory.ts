/* eslint-disable @typescript-eslint/restrict-template-expressions */
/* eslint-disable @typescript-eslint/no-unsafe-return */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
/* eslint-disable @typescript-eslint/no-unsafe-call */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
import { Config } from 'config/default';
import { config } from 'dotenv';
import { logger } from 'src/utils/logger';
import Web3 from 'web3';
// import { Accounts } from "web3-eth-accounts";
import { Account } from 'web3-core';

import { requestFaucet } from './faucet/faucet-requester';

export class AccountFactory {
  async create(): Promise<Account> {
    const account = this.createWithSpecificId();
    logger.notice(`Account created = ${account.address}`);
    if (Config.faucetUrl) {
      logger.notice(`Requesting faucet...`);
      await requestFaucet(account.address, Config.faucetQuotient * 10);
    }
    return account;
  }

  createWithoutInitialBalance(): Account {
    const account = this.createWithSpecificId();
    logger.notice(`Account created = ${account.address}`);
    // if (Config.faucetUrl) {
    //   logger.notice(`Requesting faucet...`);
    //   await requestFaucet(account.address, Config.faucetQuotient * 10);
    // }
    return account;
  }

  createWithSpecificId(id?: string): Account {
    config();
    const url = process.env.PROXY_URL === undefined ? '' : process.env.PROXY_URL;
    const web3 = new Web3(new Web3.providers.HttpProvider(url));
    // alternatively
    // web3.eth.accounts.create(web3.utils.randomHex(32))
    return id === undefined ? web3.eth.accounts.create() : web3.eth.accounts.create(id);
  }
}
