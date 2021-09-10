import Web3 from "web3";
// import { Accounts } from "web3-eth-accounts";
import { Account } from "web3-core";
// const web3 = new Web3('ws://localhost:8546');
import { config } from "dotenv";

export class AccountFactory {
  create(): Account | any {
    config();
    const url = process.env.HTTP_URL === undefined ? "" : process.env.HTTP_URL;
    const web3 = new Web3(new Web3.providers.HttpProvider(url));
    // let result = web3.eth.accounts.create();
    let result = web3.eth.accounts.create(web3.utils.randomHex(32));
    return result;
  }
}
