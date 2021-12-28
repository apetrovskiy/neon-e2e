import Web3 from 'web3';
import { Account } from 'web3-core';
import { Given, Then, When } from 'cucumber';
import { AccountFactory } from 'src/helpers/account-factory';
import { logger } from 'src/utils/logger';
import { expect } from 'chai';
import { expectations } from 'src/constants/assert-messages';
import { Balance } from 'src/helpers/balance';
import { Web3Helper } from 'src/helpers/web3-helper';
import { TxHelper } from 'src/helpers/tx-helper';
import { Config } from 'config/default';
import { errorMessages } from 'src/constants/error-messages';

let userAlice: Account;
let userBob: Account;
let initialBalanceAlice: string = '0';
let initialBalanceBob: string = '0';
const w3 = Web3Helper.getWeb3();
const balanceHelper = new Balance();
const ZEROES = '000000000000000000';

Given('there is user Alice in Ethereum network with the initial balance {int}Ξ', async (initialBalance: number) => {
  logger.notice('Creating a wallet for user Alice');
  userAlice = await new AccountFactory().create(initialBalance);
  logger.notice(`user A: ${userAlice.address}`);
  logger.notice(`planned balance is ${initialBalance}`);

  initialBalanceAlice = initialBalance.toString() + ZEROES;
  /*
  initialBalanceAlice = await balanceHelper.getBalance(userAlice);
  const ethersAmount = balanceHelper.convertToEther(initialBalanceAlice);
  logger.notice(`user A balance = ${initialBalanceAlice}`);
  logger.notice(`user A balance in ETH = ${ethersAmount}`);
  expect(ethersAmount, expectations.INITIAL_AMOUNT_SHOULD_BE).to.be.equal(Config.requestAmount.toString());
  */
});

Given('there is user Bob in Ethereum network with the initial balance {int}Ξ', async (initialBalance: number) => {
  logger.notice('Creating a wallet for user Bob');
  userBob = await new AccountFactory().create(initialBalance);
  logger.notice(`user B: ${userBob.address}`);
  logger.notice(`planned balance is ${initialBalance}`);

  initialBalanceBob = initialBalance.toString() + ZEROES;
  /*
  initialBalanceBob = await balanceHelper.getBalance(userBob);
  const ethersAmount = balanceHelper.convertToEther(initialBalanceBob);
  logger.notice(`user B balance = ${initialBalanceBob}`);
  logger.notice(`user B balance in ETH = ${ethersAmount}`);
  expect(ethersAmount, expectations.INITIAL_AMOUNT_SHOULD_BE).to.be.equal(Config.requestAmount.toString());
  */
});

When('user Alice sends {int}Ξ to user Bob', async (ethNumber: number) => {
  const deploy = async () => {
    logger.notice(`Attempting to send ${ethNumber}Ξ from ${userAlice.address} to ${userBob.address}`);
    const createTransaction: any = await TxHelper.signTransaction(userAlice, userBob, Config.requestAmount);

    // Deploy transaction
    const createReceipt = await TxHelper.deployTransaction(createTransaction);

    logger.notice(`Transaction successful with hash: ${createReceipt.transactionHash}`);
  };
  await deploy()
    .then(() => logger.notice('SUCCESS'))
    .catch((err) => {
      console.error(err);
      expect(err.message, expectations.ERROR_DURING_TRANSACTION).to.contain(errorMessages.INVALID_INSTRUCTION_DATA);
    });

  logger.notice('when is finished');
});

Then('the recipient has balance increased by {int}Ξ', async (ethNumber: number) => {
  const balance = await balanceHelper.getBalance(userBob);

  logger.notice(`initial user B balance ${initialBalanceBob}`);
  logger.notice(`amount ${w3.utils.toWei(ethNumber.toString())}`);
  logger.notice(`resulting user B balance ${balance}`);

  const expectedBalance = +initialBalanceBob + +w3.utils.toWei(ethNumber.toString());
  expect(balance, expectations.BALANCE_SHOULD_BE_EQUAL).to.be.equal(expectedBalance.toString());
});

Then('the sender has balance decreased more than by {int}Ξ', async (ethNumber: number) => {
  const balance = await balanceHelper.getBalance(userAlice);

  logger.notice(`initial user A balance ${initialBalanceAlice}`);
  logger.notice(`amount ${w3.utils.toWei(ethNumber.toString())}`);
  logger.notice(`resulting user A balance ${balance}`);

  const expectedBalance = +initialBalanceAlice - +w3.utils.toWei(ethNumber.toString());
  expect(+balance, expectations.BALANCE_SHOULD_BE_LESS_THAN).to.be.lessThan(expectedBalance);
});

Then('the sender has balance decreased by {int}Ξ', async (ethNumber: number) => {
  const balance = await balanceHelper.getBalance(userAlice);

  logger.notice(`initial user A balance ${initialBalanceAlice}`);
  logger.notice(`amount ${w3.utils.toWei(ethNumber.toString())}`);
  logger.notice(`resulting user A balance ${balance}`);

  const expectedBalance = +initialBalanceAlice - +w3.utils.toWei(ethNumber.toString());
  expect(+balance, expectations.BALANCE_SHOULD_BE_EQUAL).to.be.equal(expectedBalance);
});
