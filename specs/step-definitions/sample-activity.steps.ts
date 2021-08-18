import { defineFeature, loadFeature } from "jest-cucumber";
import { AccountFactory } from "../../src/helpers/create_account";
const feature = loadFeature("specs/features/sample-activity.feature");
import Web3 from 'web3';
// const web3 = new Web3(process.env.RPC_URL);
// const Web3 = require('web3');
// const url_evm = process.env.RPC_URL;
// const web3 = new Web3(url_evm);

defineFeature(feature, (test) => {
    let userA: any;
    let userB: any;

    function getWeb3() {
        require('dotenv').config();
        const Web3 = require('web3');
        const web3 = new Web3(process.env.RPC_URL);
        return web3;
    }

    test("a new account's activity", ({ given, when, then }) => {
        beforeAll(() => {
            console.log("BEFORE ALL");
        });
        given("there is user Alice in Ethereum", async () => {
            userA = new AccountFactory().create();
            console.log("user A");
            console.log(userA);
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
        given("there is user Bob in Ethereum", async () => {
            userB = new AccountFactory().create();
            console.log("user B");
            console.log(userB);
            const w3 = getWeb3();
            const balance = await w3.eth.getBalance(userB.address);
            console.log(`user B balanse = ${balance}`);
        });
        when("user A sends Eth to user B", () => {
            const deploy = async () => {
                console.log(`Attempting to make transaction from ${userA.toAddress} to ${userB.toAddress}`);
                const w3 = getWeb3();
                const createTransaction = await w3.eth.accounts.signTransaction(
                    {
                        from: userA.address,
                        to: userB.address,
                        value: w3.utils.toWei('100', 'ether'),
                        gas: '4294967295',
                    },
                    userA.privateKey
                );

                // Deploy transaction
                const createReceipt = await w3.eth.sendSignedTransaction(createTransaction.rawTransaction);

                console.log(`Transaction successful with hash: ${createReceipt.transactionHash}`);
            };
            // deploy()
            //     .then(() => console.log("SUCCESS"))
            //     .catch((err) => console.error(err));

            console.log('when');
        });
        then("the recipient has balance increased", () => {
            console.log('then 1');
        });
        then("the sender has balance decreased", () => {
            console.log('then 2');
        });
    });
});