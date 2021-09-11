
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
/* eslint-disable @typescript-eslint/no-unsafe-call */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
import { config } from 'dotenv';
import Web3 from 'web3';
// import { Accounts } from "web3-eth-accounts";
import { Account } from 'web3-core';

export class AccountFactory {
  create(): Account {
    return this.createWithSpecificId();
  }

  createWithSpecificId(id?: string): Account {
    config();
    const url = process.env.HTTP_URL === undefined ? '' : process.env.HTTP_URL;
    const web3 = new Web3(new Web3.providers.HttpProvider(url));
    // alternatively
    // web3.eth.accounts.create(web3.utils.randomHex(32))
    return id === undefined ? web3.eth.accounts.create() : web3.eth.accounts.create(id);
  }
}
