/*
import { expect } from 'chai';
import { Config } from 'config/default';
import { Given, Then, When, World } from 'cucumber';
// import { allure } from 'nightwatch-allure/lib/NightwatchAllureReporter';
import { expectations } from 'src/constants/assert-messages';
import { AccountFactory } from 'src/helpers/account-factory';
import { Balance } from 'src/helpers/balance';
import { requestFaucet } from 'src/helpers/faucet/faucet-requester';
import { logger } from 'src/utils/logger';
import { Account } from 'web3-core';
import { allure } from 'allure-mocha/runtime';

import { CucumberJSAllureFormatter } from 'allure-cucumberjs';
import { AllureRuntime } from 'allure-cucumberjs';

let userAlice: Account;
let initialBalanceAlice = '0';
const balanceHelper = new Balance();

Given('there is user Alice in Ethereum network with no initial balance', async () => {
  // it('creating an account without balance', async () => {
  // allure.step('creating an account without balance', async () => {
  userAlice = new AccountFactory().createWithoutInitialBalance();
  logger.notice(`user A: ${userAlice.address}`);
  this[this.userAlice] = userAlice;
  // await assertBalance(0);
  // });
  // });
});

// @step(`user Alice requests the Ether faucet for ${amount}Ξ`)
When('user Alice requests the Ether faucet for {int}Ξ', async (amount: number) => {
  // it('requesting faucet if there is URL in the .env file', async () => {
  // allure.step('requesting faucet if there is URL in the .env file', async () => {
  if (Config.faucetUrl) {
    logger.notice(`Requesting faucet...`);
    await requestFaucet(this[this.userAlice].address, amount);
  }
  // });
  // });
});

Then('the balance equals {int}Ξ', async (amount: number) => {
  // allure.step(`expecting user's balance to be equal to ${amount}`, async () => {
  await assertBalance(amount);
  // });
});

const assertBalance = async (amount: number) => {
  // it('assertBalance', async () => {
  // allure.step('assertBalance', async () => {
  initialBalanceAlice = await balanceHelper.getBalance(this[this.userAlice]);
  const ethersAmount = balanceHelper.convertToEther(initialBalanceAlice);
  logger.notice(`user A balance = ${initialBalanceAlice}`);
  logger.notice(`user A balance in ETH = ${ethersAmount}`);
  expect(ethersAmount, expectations.INITIAL_AMOUNT_SHOULD_BE).to.be.equal(amount.toString());
  // });
  // });
};
*/
