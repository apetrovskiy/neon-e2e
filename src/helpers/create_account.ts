export class AccountCreator {
    create(): object {
        require('dotenv').config();
        let Web3 = require('web3');
        let url_evm = process.env.RPC_URL;
        let web3 = new Web3(url_evm);
        let result = web3.eth.accounts.create();
        return result;
    }
}