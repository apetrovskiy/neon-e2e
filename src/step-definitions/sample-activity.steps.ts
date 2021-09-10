import { defineFeature, loadFeature } from "jest-cucumber";
import { AccountFactory } from "../helpers/create_account";
const feature = loadFeature("features/sample-activity.feature");
import Web3 from "web3";
// import { Accounts} from "web3-eth-accounts";
import { Account } from "web3-core";
import { config } from "dotenv";
// const web3 = new Web3(process.env.HTTP_URL);
// const Web3 = require('web3');
// const url_evm = process.env.HTTP_URL;
// const web3 = new Web3(url_evm);

defineFeature(feature, (test) => {
  let userA: Account | any;
  let userB: Account | any;
  ``;

  function getWeb3() {
    config();
    const url = process.env.HTTP_URL === undefined ? "" : process.env.HTTP_URL;
    const web3 = new Web3(new Web3.providers.HttpProvider(url));
    return web3;
  }

  test("a new account's activity", ({ given, when, then }) => {
    // beforeAll(() => {
    //   console.log("BEFORE ALL");
    // });
    given("there is user Alice in Ethereum network", async () => {
      userA = new AccountFactory().create();
      console.log(`user A: ${userA}`);
      const w3 = getWeb3();
      const balance = await w3.eth.getBalance(userA.address);
      console.log(`user A balanse = ${balance}`);

      /*
            var getBalance = async function () {
                await w3.eth.getBalance(userA.address).then(console.log);
            };
            await getBalance();
            */
    });
    given("there is user Bob in Ethereum network", async () => {
      userB = new AccountFactory().create();
      console.log(`user B: ${userB}`);
      const w3 = getWeb3();
      const balance = await w3.eth.getBalance(userB.address);
      console.log(`user B balanse = ${balance}`);
    });
    when(/^user Alice sends (\d+) Eth to user Bob$/, (ethNumber: number) => {
      const deploy = async () => {
        console.log(
          `Attempting to make transaction from ${userA.address} to ${userB.address}`
        );
        const w3 = getWeb3();
        const createTransaction = await w3.eth.accounts.signTransaction(
          {
            from: userA.address,
            to: userB.address,
            value: w3.utils.toWei(ethNumber.toString(), "ether"),
            gas: "4294967295",
          },
          userA.privateKey
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
    then("the recipient has balance increased", () => {
      console.log("then 1");
    });
    then("the sender has balance decreased", () => {
      console.log("then 2");
    });
  });
});
