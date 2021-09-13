import Web3 from 'web3';
import { Account } from 'web3-core';
import { Given, Then, When } from 'cucumber';
import { AccountFactory } from 'src/helpers/account-factory';
import { logger } from 'src/utils/logger';
import { expect } from 'chai';
import { expectations } from 'src/constants/assert-messages';
import { unit } from 'src/constants/unit';
import { Config } from 'config/default';
import { errorMessages } from 'src/constants/error-messages';

function getWeb3() {
  const web3 = new Web3(new Web3.providers.HttpProvider(Config.url));
  return web3;
}

let userAlice: Account;
let userBob: Account;
let initialBalanceAlice: number = 0;
let initialBalanceBob: number = 0;
const w3 = getWeb3();

Given('there is user Alice in Ethereum network with the initial balance {int}Ξ', async (initialBalance: number) => {
  userAlice = new AccountFactory().create();
  logger.notice(`user A: ${userAlice.address}`);
  const balance = await w3.eth.getBalance(userAlice.address);
  const ethersAmount = w3.utils.fromWei(balance, unit.ether);
  logger.notice(`user A balance = ${balance}`);
  logger.notice(`user A balance in ETH = ${ethersAmount}`);
  initialBalanceAlice = balance;
  expect(ethersAmount, expectations.INITIAL_AMOUNT_SHOULD_BE).to.be.equal(initialBalance.toString());
});

Given('there is user Bob in Ethereum network with the initial balance {int}Ξ', async (initialBalance: number) => {
  userBob = new AccountFactory().create();
  logger.notice(`user B: ${userBob.address}`);
  const balance = await w3.eth.getBalance(userBob.address);
  const ethersAmount = w3.utils.fromWei(balance, unit.ether);
  logger.notice(`user B balance = ${balance}`);
  logger.notice(`user B balance in ETH = ${ethersAmount}`);
  initialBalanceBob = balance;
  expect(ethersAmount, expectations.INITIAL_AMOUNT_SHOULD_BE).to.be.equal(initialBalance.toString());
});

When('user Alice sends {int}Ξ to user Bob', async (ethNumber: number) => {
  const deploy = async () => {
    logger.notice(`Attempting to send ${ethNumber}Ξ from ${userAlice.address} to ${userBob.address}`);
    const createTransaction = await w3.eth.accounts.signTransaction(
      {
        from: userAlice.address,
        to: userBob.address,
        value: w3.utils.toWei(ethNumber.toString(), unit.ether),
        gas: '4294967295'
      },
      userAlice.privateKey
    );

    // Deploy transaction
    const createReceipt = await w3.eth.sendSignedTransaction(createTransaction.rawTransaction!);

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
  const balance = await w3.eth.getBalance(userBob.address);

  logger.notice(`initial user B balance ${initialBalanceBob}`);
  logger.notice(`amount ${w3.utils.toWei(ethNumber.toString())}`);
  logger.notice(`resulting user B balance ${balance}`);

  const expectedBalance = +initialBalanceBob + +w3.utils.toWei(ethNumber.toString());
  expect(balance, expectations.BALANCE_SHOULD_BE_EQUAL).to.be.equal(expectedBalance.toString());
});

Then('the sender has balance decreased more than by {int}Ξ', async (ethNumber: number) => {
  const balance = await w3.eth.getBalance(userAlice.address);

  logger.notice(`initial user A balance ${initialBalanceAlice}`);
  logger.notice(`amount ${w3.utils.toWei(ethNumber.toString())}`);
  logger.notice(`resulting user A balance ${balance}`);

  const expectedBalance = +initialBalanceAlice - +w3.utils.toWei(ethNumber.toString());
  expect(+balance, expectations.BALANCE_SHOULD_BE_LESS_THAN).to.be.lessThan(expectedBalance);
});

Then('the sender has balance decreased by {int}Ξ', async (ethNumber: number) => {
  const balance = await w3.eth.getBalance(userAlice.address);

  logger.notice(`initial user A balance ${initialBalanceAlice}`);
  logger.notice(`amount ${w3.utils.toWei(ethNumber.toString())}`);
  logger.notice(`resulting user A balance ${balance}`);

  const expectedBalance = +initialBalanceAlice - +w3.utils.toWei(ethNumber.toString());
  expect(+balance, expectations.BALANCE_SHOULD_BE_EQUAL).to.be.equal(expectedBalance);
});
