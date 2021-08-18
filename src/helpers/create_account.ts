// import Web3 from 'web3';
// const web3 = new Web3('ws://localhost:8546');

export class AccountFactory {
    create(): object {
        require('dotenv').config();
        const Web3 = require('web3');
        const url_evm = process.env.RPC_URL;
        const web3 = new Web3(url_evm);
        // let result = web3.eth.accounts.create();
        let result = web3.eth.accounts.create(web3.utils.randomHex(32));
        return result;
    }
}