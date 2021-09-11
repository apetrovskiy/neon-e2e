// import { defineFeature, loadFeature } from "jest-cucumber";
// import { AccountFactory } from "../helpers/create_account";
// const feature = loadFeature("features/sample-activity.feature");
import Web3 from 'web3';
// import { Accounts} from "web3-eth-accounts";
import { Account } from 'web3-core';
import { config } from 'dotenv';

// ? Given there is user Alice in Ethereum network
// Undefined. Implement with the following snippet:

import { Given, Then, When } from 'cucumber';
import { AccountFactory } from 'src/helpers/create_account';

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
  // // Write code here that turns the phrase above into concrete actions
  // return 'pending';
  userAlice = new AccountFactory().create();
  console.log(`user A: ${userAlice}`);
  const balance = await w3.eth.getBalance(userAlice.address);
  console.log(`user A balanse = ${balance}`);
});

// ? And there is user Bob in Ethereum network
// Undefined. Implement with the following snippet:

Given('there is user Bob in Ethereum network', async () => {
  // // Write code here that turns the phrase above into concrete actions
  // return 'pending';
  userBob = new AccountFactory().create();
  console.log(`user B: ${userBob}`);
  const balance = await w3.eth.getBalance(userBob.address);
  console.log(`user B balanse = ${balance}`);
});

// ? When user Alice sends 1 Eth to user Bob
// Undefined. Implement with the following snippet:

When('user Alice sends {int} Eth to user Bob', function (ethNumber: number) {
  // // When('{ProductName} {ProductName} {ProductName} {float} {ProductName} {ProductName} {ProductName} {ProductName}', function (ProductName, ProductName2, ProductName3, float, ProductName4, ProductName5, ProductName6, ProductName7) {
  // // When('{ProductName} {ProductName} {ProductName} {ProductName} {ProductName} {ProductName} {ProductName} {ProductName}', function (ProductName, ProductName2, ProductName3, ProductName4, ProductName5, ProductName6, ProductName7, ProductName8) {
  // // Write code here that turns the phrase above into concrete actions
  // // eslint-disable-next-line no-console
  // console.log(ethNumber);
  // return 'pending';
  const deploy = async () => {
    console.log(
      `Attempting to make transaction from ${userAlice.address} to ${userBob.address}`
    );
    const w3 = getWeb3();
    const createTransaction = await w3.eth.accounts.signTransaction(
      {
        from: userAlice.address,
        to: userBob.address,
        value: w3.utils.toWei(ethNumber.toString(), "ether"),
        gas: "4294967295",
      },
      userAlice.privateKey
    );

    // Deploy transaction
    const createReceipt = await w3.eth.sendSignedTransaction(
      createTransaction.rawTransaction!
    );

    console.log(
      `Transaction successful with hash: ${createReceipt.transactionHash}`
    );
  };
  deploy()
      .then(() => console.log("SUCCESS"))
      .catch((err) => console.error(err));

  console.log("when");
});

// ? Then the recipient has balance increased
// Undefined. Implement with the following snippet:

Then('the recipient has balance increased', async () => {
  // // Write code here that turns the phrase above into concrete actions
  // return 'pending';
  console.log("then 1");
  const balance = await w3.eth.getBalance(userBob.address);
  console.log(`user B balanse = ${balance}`);
});

// ? And the sender has balance decreased
// Undefined. Implement with the following snippet:

Then('the sender has balance decreased', async () => {
  // // Write code here that turns the phrase above into concrete actions
  // return 'pending';
  console.log("then 2");
  const balance = await w3.eth.getBalance(userAlice.address);
  console.log(`user A balanse = ${balance}`);
});
