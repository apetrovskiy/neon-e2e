/* eslint-disable @typescript-eslint/restrict-template-expressions */
/* eslint-disable @typescript-eslint/no-unsafe-return */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
/* eslint-disable @typescript-eslint/no-unsafe-call */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
import { Config } from 'config/default';
import { config } from 'dotenv';
import { logger } from 'src/utils/logger';
import Web3 from 'web3';
import { allure } from 'allure-mocha/runtime';
// import { Accounts } from "web3-eth-accounts";
import { Account } from 'web3-core';

import { requestFaucet } from './faucet/faucet-requester';

export class AccountFactory {
  async create(amount?: number | undefined): Promise<Account> {
    // return allure.step(`creating account with balance ${amount}`, async () => {
    const account = this.createWithSpecificId();
    logger.notice(`Account created = ${account.address}`);
    // if (Config.useFaucet) {
    if (amount !== undefined && amount > 0) {
      logger.notice(`Requesting faucet...`);
      const requestAmont = amount ?? Config.requestAmount;
      await requestFaucet(account.address, requestAmont);
    }
    return account;
    // });
  }

  createWithoutInitialBalance(): Account {
    // return allure.step(`creating account with no initial balance`, async () => {
    const account = this.createWithSpecificId();
    logger.notice(`Account created = ${account.address}`);
    return account;
    // });
  }

  createWithSpecificId(id?: string): Account {
    // return allure.step(`creating with id = ${id}`, () => {
    config();
    const url = process.env.PROXY_URL === undefined ? '' : process.env.PROXY_URL;
    const web3 = new Web3(new Web3.providers.HttpProvider(url));
    // alternatively
    // web3.eth.accounts.create(web3.utils.randomHex(32))
    return id === undefined ? web3.eth.accounts.create() : web3.eth.accounts.create(id);
    // });
  }
}
