import { expect } from 'chai';
import { Config } from 'config/default';
import { Given, Then, When } from 'cucumber';
import { expectations } from 'src/constants/assert-messages';
import { AccountFactory } from 'src/helpers/account-factory';
import { Balance } from 'src/helpers/balance';
import { requestFaucet } from 'src/helpers/faucet/faucet-requester';
import { logger } from 'src/utils/logger';
import { Account } from 'web3-core';

let userAlice: Account;
let initialBalanceAlice = "0";
const quotient: number = Config.faucetQuotient;
const balanceHelper = new Balance();
const applyQuotient = (amount: number) => amount * quotient;

Given('there is user Alice in Ethereum network with no initial balance', async () => {
  userAlice = new AccountFactory().createWithoutInitialBalance();
  logger.notice(`user A: ${userAlice.address}`);
});

When('the user requests the ERC20 faucet for {int}Ξ', async (amount: number) => {
  if (Config.faucetUrl) {
    logger.notice(`Requesting faucet...`);
    await requestFaucet(userAlice.address, amount);
  }
});

Then('the balance equals {int}Ξ', async (amount: number) => {
  initialBalanceAlice = await balanceHelper.getBalance(userAlice);
  const ethersAmount = balanceHelper.convertToEther(initialBalanceAlice);
  logger.notice(`user A balance = ${initialBalanceAlice}`);
  logger.notice(`user A balance in ETH = ${ethersAmount}`);
  expect(ethersAmount, expectations.INITIAL_AMOUNT_SHOULD_BE).to.be.equal(applyQuotient(amount).toString());
});
