import Web3 from 'web3';
import { Account } from 'web3-core';
import { config } from 'dotenv';
import { Given, Then, When } from 'cucumber';
import { AccountFactory } from 'src/helpers/create_account';
import { logger } from 'src/utils/logger';

function getWeb3() {
  config();
  const url = process.env.HTTP_URL === undefined ? '' : process.env.HTTP_URL;
  const web3 = new Web3(new Web3.providers.HttpProvider(url));
  return web3;
}

let userAlice: Account;
let userBob: Account;
const w3 = getWeb3();

Given('there is user Alice in Ethereum network', async () => {
  userAlice = new AccountFactory().create();
  logger.notice(`user A: ${userAlice.address}`);
  const balance = await w3.eth.getBalance(userAlice.address);
  logger.notice(`user A balanse = ${balance}`);
});

Given('there is user Bob in Ethereum network', async () => {
  userBob = new AccountFactory().create();
  logger.notice(`user B: ${userBob.address}`);
  const balance = await w3.eth.getBalance(userBob.address);
  logger.notice(`user B balanse = ${balance}`);
});

When('user Alice sends {int}Ξ to user Bob', async (ethNumber: number) => {
  const deploy = async () => {
    logger.notice(`Attempting to make transaction from ${userAlice.address} to ${userBob.address}`);
    const w3 = getWeb3();
    const createTransaction = await w3.eth.accounts.signTransaction(
      {
        from: userAlice.address,
        to: userBob.address,
        value: w3.utils.toWei(ethNumber.toString(), 'ether'),
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
    .catch((err) => console.error(err));

  logger.notice('when is finished');
});

Then('the recipient has balance increased', async () => {
  const balance = await w3.eth.getBalance(userBob.address);
  logger.notice(`user B balanse = ${balance}`);
});

Then('the sender has balance decreased', async () => {
  const balance = await w3.eth.getBalance(userAlice.address);
  logger.notice(`user A balanse = ${balance}`);
});
